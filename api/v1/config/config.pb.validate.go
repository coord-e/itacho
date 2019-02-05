// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/config.proto

package config

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on ServiceDefinition with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ServiceDefinition) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetVersion() != 1 {
		return ServiceDefinitionValidationError{
			field:  "Version",
			reason: "value must equal 1",
		}
	}

	if len(m.GetDependencies()) < 1 {
		return ServiceDefinitionValidationError{
			field:  "Dependencies",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetDependencies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServiceDefinitionValidationError{
					field:  fmt.Sprintf("Dependencies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ServiceDefinitionValidationError is the validation error returned by
// ServiceDefinition.Validate if the designated constraints aren't met.
type ServiceDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceDefinitionValidationError) ErrorName() string {
	return "ServiceDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceDefinitionValidationError{}

// Validate checks the field values on Dependency with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Dependency) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return DependencyValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetClusterName()) < 1 {
		return DependencyValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for HostHeader

	if v, ok := interface{}(m.GetTls()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DependencyValidationError{
				field:  "Tls",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetConnectTimeoutMs() <= 0 {
		return DependencyValidationError{
			field:  "ConnectTimeoutMs",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetCircuitBreaker()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DependencyValidationError{
				field:  "CircuitBreaker",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetRoutes()) < 1 {
		return DependencyValidationError{
			field:  "Routes",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DependencyValidationError{
					field:  fmt.Sprintf("Routes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetOutlierDetection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DependencyValidationError{
				field:  "OutlierDetection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HealthCheckPath

	if v, ok := interface{}(m.GetFaultFilterConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DependencyValidationError{
				field:  "FaultFilterConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DependencyValidationError{
				field:  "Headers",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.DiscoveryOption.(type) {

	case *Dependency_Lb:
		// no validation rules for Lb

	case *Dependency_Sds:

		if v, ok := interface{}(m.GetSds()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DependencyValidationError{
					field:  "Sds",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return DependencyValidationError{
			field:  "DiscoveryOption",
			reason: "value is required",
		}

	}

	return nil
}

// DependencyValidationError is the validation error returned by
// Dependency.Validate if the designated constraints aren't met.
type DependencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DependencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DependencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DependencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DependencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DependencyValidationError) ErrorName() string { return "DependencyValidationError" }

// Error satisfies the builtin error interface
func (e DependencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDependency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DependencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DependencyValidationError{}

// Validate checks the field values on CircuitBreaker with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CircuitBreaker) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMaxConnections() <= 0 {
		return CircuitBreakerValidationError{
			field:  "MaxConnections",
			reason: "value must be greater than 0",
		}
	}

	if m.GetMaxPendingRequests() <= 0 {
		return CircuitBreakerValidationError{
			field:  "MaxPendingRequests",
			reason: "value must be greater than 0",
		}
	}

	if m.GetMaxRetries() <= 0 {
		return CircuitBreakerValidationError{
			field:  "MaxRetries",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CircuitBreakerValidationError is the validation error returned by
// CircuitBreaker.Validate if the designated constraints aren't met.
type CircuitBreakerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerValidationError) ErrorName() string { return "CircuitBreakerValidationError" }

// Error satisfies the builtin error interface
func (e CircuitBreakerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreaker.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Method

	if m.GetTimeoutMs() <= 0 {
		return RouteValidationError{
			field:  "TimeoutMs",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetRetryPolicy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "RetryPolicy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.PathSpecifier.(type) {

	case *Route_Prefix:
		// no validation rules for Prefix

	case *Route_Path:
		// no validation rules for Path

	default:
		return RouteValidationError{
			field:  "PathSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on RetryPolicy with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RetryPolicy) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetRetryOn()) < 1 {
		return RetryPolicyValidationError{
			field:  "RetryOn",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetNumRetries() <= 0 {
		return RetryPolicyValidationError{
			field:  "NumRetries",
			reason: "value must be greater than 0",
		}
	}

	if m.GetPerTryTimeoutMs() <= 0 {
		return RetryPolicyValidationError{
			field:  "PerTryTimeoutMs",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RetryPolicyValidationError is the validation error returned by
// RetryPolicy.Validate if the designated constraints aren't met.
type RetryPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetryPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetryPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetryPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetryPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetryPolicyValidationError) ErrorName() string { return "RetryPolicyValidationError" }

// Error satisfies the builtin error interface
func (e RetryPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetryPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetryPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetryPolicyValidationError{}

// Validate checks the field values on OutlierDetection with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OutlierDetection) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Consecutive_5Xx

	return nil
}

// OutlierDetectionValidationError is the validation error returned by
// OutlierDetection.Validate if the designated constraints aren't met.
type OutlierDetectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutlierDetectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutlierDetectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutlierDetectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutlierDetectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutlierDetectionValidationError) ErrorName() string { return "OutlierDetectionValidationError" }

// Error satisfies the builtin error interface
func (e OutlierDetectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutlierDetection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutlierDetectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutlierDetectionValidationError{}

// Validate checks the field values on FaultFilterConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FaultFilterConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultFilterConfigValidationError{
				field:  "Delay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultFilterConfigValidationError{
				field:  "Abort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FaultFilterConfigValidationError is the validation error returned by
// FaultFilterConfig.Validate if the designated constraints aren't met.
type FaultFilterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultFilterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultFilterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultFilterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultFilterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultFilterConfigValidationError) ErrorName() string {
	return "FaultFilterConfigValidationError"
}

// Error satisfies the builtin error interface
func (e FaultFilterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultFilterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultFilterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultFilterConfigValidationError{}

// Validate checks the field values on FaultDelay with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FaultDelay) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetType() != "fixed" {
		return FaultDelayValidationError{
			field:  "Type",
			reason: "value must equal fixed",
		}
	}

	if m.GetFixedDelay() == nil {
		return FaultDelayValidationError{
			field:  "FixedDelay",
			reason: "value is required",
		}
	}

	if m.GetPercentage() == nil {
		return FaultDelayValidationError{
			field:  "Percentage",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultDelayValidationError{
				field:  "Percentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FaultDelayValidationError is the validation error returned by
// FaultDelay.Validate if the designated constraints aren't met.
type FaultDelayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultDelayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultDelayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultDelayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultDelayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultDelayValidationError) ErrorName() string { return "FaultDelayValidationError" }

// Error satisfies the builtin error interface
func (e FaultDelayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultDelay.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultDelayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultDelayValidationError{}

// Validate checks the field values on FaultAbort with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FaultAbort) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHttpStatus() >= 600 {
		return FaultAbortValidationError{
			field:  "HttpStatus",
			reason: "value must be less than 600",
		}
	}

	if m.GetPercentage() == nil {
		return FaultAbortValidationError{
			field:  "Percentage",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultAbortValidationError{
				field:  "Percentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FaultAbortValidationError is the validation error returned by
// FaultAbort.Validate if the designated constraints aren't met.
type FaultAbortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultAbortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultAbortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultAbortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultAbortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultAbortValidationError) ErrorName() string { return "FaultAbortValidationError" }

// Error satisfies the builtin error interface
func (e FaultAbortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultAbort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultAbortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultAbortValidationError{}

// Validate checks the field values on FaultFractionalPercent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultFractionalPercent) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNumerator() <= 0 {
		return FaultFractionalPercentValidationError{
			field:  "Numerator",
			reason: "value must be greater than 0",
		}
	}

	if _, ok := _FaultFractionalPercent_Denominator_InLookup[m.GetDenominator()]; !ok {
		return FaultFractionalPercentValidationError{
			field:  "Denominator",
			reason: "value must be in list [HUNDRED TEN_THOUSAND MILLION]",
		}
	}

	return nil
}

// FaultFractionalPercentValidationError is the validation error returned by
// FaultFractionalPercent.Validate if the designated constraints aren't met.
type FaultFractionalPercentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultFractionalPercentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultFractionalPercentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultFractionalPercentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultFractionalPercentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultFractionalPercentValidationError) ErrorName() string {
	return "FaultFractionalPercentValidationError"
}

// Error satisfies the builtin error interface
func (e FaultFractionalPercentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultFractionalPercent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultFractionalPercentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultFractionalPercentValidationError{}

var _FaultFractionalPercent_Denominator_InLookup = map[string]struct{}{
	"HUNDRED":      {},
	"TEN_THOUSAND": {},
	"MILLION":      {},
}

// Validate checks the field values on Headers with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Headers) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequestHeadersToAdd() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HeadersValidationError{
					field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HeadersValidationError is the validation error returned by Headers.Validate
// if the designated constraints aren't met.
type HeadersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeadersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeadersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeadersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeadersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeadersValidationError) ErrorName() string { return "HeadersValidationError" }

// Error satisfies the builtin error interface
func (e HeadersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaders.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeadersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeadersValidationError{}

// Validate checks the field values on HeaderValueOption with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeaderValueOption) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHeader() == nil {
		return HeaderValueOptionValidationError{
			field:  "Header",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderValueOptionValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAppend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderValueOptionValidationError{
				field:  "Append",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HeaderValueOptionValidationError is the validation error returned by
// HeaderValueOption.Validate if the designated constraints aren't met.
type HeaderValueOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValueOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValueOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValueOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValueOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValueOptionValidationError) ErrorName() string {
	return "HeaderValueOptionValidationError"
}

// Error satisfies the builtin error interface
func (e HeaderValueOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderValueOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValueOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValueOptionValidationError{}

// Validate checks the field values on HeaderValue with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HeaderValue) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		return HeaderValueValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for Value

	return nil
}

// HeaderValueValidationError is the validation error returned by
// HeaderValue.Validate if the designated constraints aren't met.
type HeaderValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValueValidationError) ErrorName() string { return "HeaderValueValidationError" }

// Error satisfies the builtin error interface
func (e HeaderValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValueValidationError{}
