// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: workflowcontract/v1/crafting_schema.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on CraftingSchema with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CraftingSchema) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CraftingSchema with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CraftingSchemaMultiError,
// or nil if none found.
func (m *CraftingSchema) ValidateAll() error {
	return m.validate(true)
}

func (m *CraftingSchema) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSchemaVersion() != "v1" {
		err := CraftingSchemaValidationError{
			field:  "SchemaVersion",
			reason: "value must equal v1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetMaterials() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CraftingSchemaValidationError{
						field:  fmt.Sprintf("Materials[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CraftingSchemaValidationError{
						field:  fmt.Sprintf("Materials[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CraftingSchemaValidationError{
					field:  fmt.Sprintf("Materials[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetRunner()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CraftingSchemaValidationError{
					field:  "Runner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CraftingSchemaValidationError{
					field:  "Runner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRunner()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CraftingSchemaValidationError{
				field:  "Runner",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAnnotations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CraftingSchemaValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CraftingSchemaValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CraftingSchemaValidationError{
					field:  fmt.Sprintf("Annotations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CraftingSchemaMultiError(errors)
	}

	return nil
}

// CraftingSchemaMultiError is an error wrapping multiple validation errors
// returned by CraftingSchema.ValidateAll() if the designated constraints
// aren't met.
type CraftingSchemaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CraftingSchemaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CraftingSchemaMultiError) AllErrors() []error { return m }

// CraftingSchemaValidationError is the validation error returned by
// CraftingSchema.Validate if the designated constraints aren't met.
type CraftingSchemaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CraftingSchemaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CraftingSchemaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CraftingSchemaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CraftingSchemaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CraftingSchemaValidationError) ErrorName() string { return "CraftingSchemaValidationError" }

// Error satisfies the builtin error interface
func (e CraftingSchemaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCraftingSchema.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CraftingSchemaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CraftingSchemaValidationError{}

// Validate checks the field values on Annotation with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Annotation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Annotation with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AnnotationMultiError, or
// nil if none found.
func (m *Annotation) ValidateAll() error {
	return m.validate(true)
}

func (m *Annotation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_Annotation_Name_Pattern.MatchString(m.GetName()) {
		err := AnnotationValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[\\\\w]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Value

	if len(errors) > 0 {
		return AnnotationMultiError(errors)
	}

	return nil
}

// AnnotationMultiError is an error wrapping multiple validation errors
// returned by Annotation.ValidateAll() if the designated constraints aren't met.
type AnnotationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AnnotationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AnnotationMultiError) AllErrors() []error { return m }

// AnnotationValidationError is the validation error returned by
// Annotation.Validate if the designated constraints aren't met.
type AnnotationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AnnotationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AnnotationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AnnotationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AnnotationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AnnotationValidationError) ErrorName() string { return "AnnotationValidationError" }

// Error satisfies the builtin error interface
func (e AnnotationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAnnotation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AnnotationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AnnotationValidationError{}

var _Annotation_Name_Pattern = regexp.MustCompile("^[\\w]+$")

// Validate checks the field values on CraftingSchema_Runner with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CraftingSchema_Runner) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CraftingSchema_Runner with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CraftingSchema_RunnerMultiError, or nil if none found.
func (m *CraftingSchema_Runner) ValidateAll() error {
	return m.validate(true)
}

func (m *CraftingSchema_Runner) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _CraftingSchema_Runner_Type_NotInLookup[m.GetType()]; ok {
		err := CraftingSchema_RunnerValidationError{
			field:  "Type",
			reason: "value must not be in list [RUNNER_TYPE_UNSPECIFIED]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CraftingSchema_RunnerMultiError(errors)
	}

	return nil
}

// CraftingSchema_RunnerMultiError is an error wrapping multiple validation
// errors returned by CraftingSchema_Runner.ValidateAll() if the designated
// constraints aren't met.
type CraftingSchema_RunnerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CraftingSchema_RunnerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CraftingSchema_RunnerMultiError) AllErrors() []error { return m }

// CraftingSchema_RunnerValidationError is the validation error returned by
// CraftingSchema_Runner.Validate if the designated constraints aren't met.
type CraftingSchema_RunnerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CraftingSchema_RunnerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CraftingSchema_RunnerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CraftingSchema_RunnerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CraftingSchema_RunnerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CraftingSchema_RunnerValidationError) ErrorName() string {
	return "CraftingSchema_RunnerValidationError"
}

// Error satisfies the builtin error interface
func (e CraftingSchema_RunnerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCraftingSchema_Runner.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CraftingSchema_RunnerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CraftingSchema_RunnerValidationError{}

var _CraftingSchema_Runner_Type_NotInLookup = map[CraftingSchema_Runner_RunnerType]struct{}{
	0: {},
}

// Validate checks the field values on CraftingSchema_Material with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CraftingSchema_Material) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CraftingSchema_Material with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CraftingSchema_MaterialMultiError, or nil if none found.
func (m *CraftingSchema_Material) ValidateAll() error {
	return m.validate(true)
}

func (m *CraftingSchema_Material) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _CraftingSchema_Material_Type_NotInLookup[m.GetType()]; ok {
		err := CraftingSchema_MaterialValidationError{
			field:  "Type",
			reason: "value must not be in list [MATERIAL_TYPE_UNSPECIFIED]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := CraftingSchema_Material_MaterialType_name[int32(m.GetType())]; !ok {
		err := CraftingSchema_MaterialValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CraftingSchema_Material_Name_Pattern.MatchString(m.GetName()) {
		err := CraftingSchema_MaterialValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[\\\\w|-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Optional

	// no validation rules for Output

	for idx, item := range m.GetAnnotations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CraftingSchema_MaterialValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CraftingSchema_MaterialValidationError{
						field:  fmt.Sprintf("Annotations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CraftingSchema_MaterialValidationError{
					field:  fmt.Sprintf("Annotations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CraftingSchema_MaterialMultiError(errors)
	}

	return nil
}

// CraftingSchema_MaterialMultiError is an error wrapping multiple validation
// errors returned by CraftingSchema_Material.ValidateAll() if the designated
// constraints aren't met.
type CraftingSchema_MaterialMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CraftingSchema_MaterialMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CraftingSchema_MaterialMultiError) AllErrors() []error { return m }

// CraftingSchema_MaterialValidationError is the validation error returned by
// CraftingSchema_Material.Validate if the designated constraints aren't met.
type CraftingSchema_MaterialValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CraftingSchema_MaterialValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CraftingSchema_MaterialValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CraftingSchema_MaterialValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CraftingSchema_MaterialValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CraftingSchema_MaterialValidationError) ErrorName() string {
	return "CraftingSchema_MaterialValidationError"
}

// Error satisfies the builtin error interface
func (e CraftingSchema_MaterialValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCraftingSchema_Material.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CraftingSchema_MaterialValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CraftingSchema_MaterialValidationError{}

var _CraftingSchema_Material_Type_NotInLookup = map[CraftingSchema_Material_MaterialType]struct{}{
	0: {},
}

var _CraftingSchema_Material_Name_Pattern = regexp.MustCompile("^[\\w|-]+$")
