package generator

import (
	"time"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"

	config "github.com/cookpad/itacho/api/v1/config"
)

// ConvertServiceDefinitionToRdsResources return resources of Envoy xDS API.
func ConvertServiceDefinitionToRdsResources(def *config.ServiceDefinition) (*[]proto.Message, error) {
	deps := def.GetDependencies()
	vhosts := make([]route.VirtualHost, len(deps))
	for i, dep := range deps {
		vhost, err := convertToVirtualHost(dep)
		if err != nil {
			return nil, err
		}
		vhosts[i] = *vhost
	}

	r := &api.RouteConfiguration{
		Name:         "default",
		VirtualHosts: vhosts,
	}
	return &[]proto.Message{r}, nil
}

func convertToVirtualHost(dep *config.Dependency) (*route.VirtualHost, error) {
	var routes []route.Route
	for _, r := range dep.GetRoutes() {
		rs, err := convertToRoutes(dep, r)
		if err != nil {
			return nil, err
		}
		for _, rr := range rs {
			routes = append(routes, *rr)
		}
	}

	virtualHost := &route.VirtualHost{
		Name:    dep.GetName(),
		Domains: []string{dep.GetName()},
		Routes:  routes,
	}

	if dep.GetFaultFilterConfig() != nil {
		virtualHost.PerFilterConfig = map[string]*types.Struct{
			"envoy.fault": buildFaultFilterConfig(dep.GetFaultFilterConfig()),
		}
	}

	return virtualHost, nil
}

func convertToRoutes(dep *config.Dependency, r *config.Route) ([]*route.Route, error) {
	timeout := time.Duration(r.GetTimeoutMs()*1000*1000) * time.Nanosecond
	match := route.RouteMatch{}
	if r.GetPrefix() == "" {
		match.PathSpecifier = &route.RouteMatch_Path{Path: r.GetPath()}
	} else {
		match.PathSpecifier = &route.RouteMatch_Prefix{Prefix: r.GetPrefix()}
	}

	var rs []*route.Route

	if r.GetRetryPolicy() == nil {
		if r.GetMethod() != "" {
			match.Headers = []*route.HeaderMatcher{
				&route.HeaderMatcher{
					Name: ":method",
					HeaderMatchSpecifier: &route.HeaderMatcher_ExactMatch{
						ExactMatch: r.GetMethod(),
					},
				},
			}
		}
		r := buildRoute(dep, &timeout, match, nil)
		rs = []*route.Route{&r}
	} else {
		if r.GetMethod() == "" {
			// add retriable routes for GET/HEAD requests
			mWithRetry := match
			mWithRetry.Headers = []*route.HeaderMatcher{
				&route.HeaderMatcher{
					Name: ":method",
					HeaderMatchSpecifier: &route.HeaderMatcher_RegexMatch{
						RegexMatch: "(GET|HEAD)",
					},
				},
			}
			routeWithRetry := buildRoute(dep, &timeout, mWithRetry, r.GetRetryPolicy())
			routeWoRetry := buildRoute(dep, &timeout, match, nil)
			rs = []*route.Route{&routeWithRetry, &routeWoRetry}
		} else {
			match.Headers = []*route.HeaderMatcher{
				&route.HeaderMatcher{
					Name: ":method",
					HeaderMatchSpecifier: &route.HeaderMatcher_ExactMatch{
						ExactMatch: r.GetMethod(),
					},
				},
			}
			r := buildRoute(dep, &timeout, match, r.GetRetryPolicy())
			rs = []*route.Route{&r}
		}
	}

	return rs, nil
}

func buildFaultFilterConfig(faultFilterConfig *config.FaultFilterConfig) *types.Struct {
	faultConfig := &types.Struct{
		Fields: map[string]*types.Value{},
	}

	if faultFilterConfig.GetAbort() != nil {
		faultConfig.Fields["abort"] = buildFaultAbortConfig(faultFilterConfig.GetAbort())
	}
	if faultFilterConfig.GetDelay() != nil {
		faultConfig.Fields["delay"] = buildFaultDelayConfig(faultFilterConfig.GetDelay())
	}

	return faultConfig
}

func buildFaultAbortConfig(faultAbort *config.FaultAbort) *types.Value {
	abortConfig := &types.Value{Kind: &types.Value_StructValue{StructValue: &types.Struct{
		Fields: map[string]*types.Value{
			"http_status": &types.Value{Kind: &types.Value_NumberValue{NumberValue: float64(faultAbort.GetHttpStatus())}},
			"percentage": &types.Value{Kind: &types.Value_StructValue{StructValue: &types.Struct{
				Fields: map[string]*types.Value{
					"numerator":   &types.Value{Kind: &types.Value_NumberValue{NumberValue: float64(faultAbort.GetPercentage().GetNumerator())}},
					"denominator": &types.Value{Kind: &types.Value_StringValue{StringValue: faultAbort.GetPercentage().GetDenominator()}},
				},
			}}},
		},
	}}}

	return abortConfig
}

func buildFaultDelayConfig(faultDelay *config.FaultDelay) *types.Value {
	delayConfig := &types.Value{Kind: &types.Value_StructValue{StructValue: &types.Struct{
		Fields: map[string]*types.Value{
			"type":        &types.Value{Kind: &types.Value_StringValue{StringValue: faultDelay.GetType()}},
			"fixed_delay": &types.Value{Kind: &types.Value_StringValue{StringValue: faultDelay.GetFixedDelay().String()}},
			"percentage": &types.Value{Kind: &types.Value_StructValue{StructValue: &types.Struct{
				Fields: map[string]*types.Value{
					"numerator":   &types.Value{Kind: &types.Value_NumberValue{NumberValue: float64(faultDelay.GetPercentage().GetNumerator())}},
					"denominator": &types.Value{Kind: &types.Value_StringValue{StringValue: faultDelay.GetPercentage().GetDenominator()}},
				},
			}}},
		},
	}}}

	return delayConfig
}

func buildRoute(dep *config.Dependency, timeout *time.Duration, m route.RouteMatch, retryConfig *config.RetryPolicy) route.Route {
	var policy *route.RouteAction_RetryPolicy
	if retryConfig != nil {
		perTryTimeout := time.Duration(retryConfig.GetPerTryTimeoutMs()*1000*1000) * time.Nanosecond
		policy = &route.RouteAction_RetryPolicy{
			RetryOn:       retryConfig.GetRetryOn(),
			NumRetries:    &types.UInt32Value{Value: retryConfig.GetNumRetries()},
			PerTryTimeout: &perTryTimeout,
		}
	}

	r := route.Route{
		Match: m,
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: dep.GetClusterName(),
				},
				Timeout:     timeout,
				RetryPolicy: policy,
			},
		},
	}

	if dep.GetHostHeader() != "" {
		r.GetRoute().HostRewriteSpecifier = &route.RouteAction_HostRewrite{HostRewrite: dep.GetHostHeader()}
	} else {
		r.GetRoute().HostRewriteSpecifier = &route.RouteAction_AutoHostRewrite{AutoHostRewrite: &types.BoolValue{Value: true}}
	}

	return r
}
