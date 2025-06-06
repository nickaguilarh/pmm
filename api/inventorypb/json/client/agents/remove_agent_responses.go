// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoveAgentReader is a Reader for the RemoveAgent structure.
type RemoveAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveAgentOK creates a RemoveAgentOK with default headers values
func NewRemoveAgentOK() *RemoveAgentOK {
	return &RemoveAgentOK{}
}

/*
RemoveAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type RemoveAgentOK struct {
	Payload interface{}
}

func (o *RemoveAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/Remove][%d] removeAgentOk  %+v", 200, o.Payload)
}

func (o *RemoveAgentOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAgentDefault creates a RemoveAgentDefault with default headers values
func NewRemoveAgentDefault(code int) *RemoveAgentDefault {
	return &RemoveAgentDefault{
		_statusCode: code,
	}
}

/*
RemoveAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RemoveAgentDefault struct {
	_statusCode int

	Payload *RemoveAgentDefaultBody
}

// Code gets the status code for the remove agent default response
func (o *RemoveAgentDefault) Code() int {
	return o._statusCode
}

func (o *RemoveAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/Remove][%d] RemoveAgent default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveAgentDefault) GetPayload() *RemoveAgentDefaultBody {
	return o.Payload
}

func (o *RemoveAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(RemoveAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RemoveAgentBody remove agent body
swagger:model RemoveAgentBody
*/
type RemoveAgentBody struct {
	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// Remove agent with all dependencies.
	Force bool `json:"force,omitempty"`
}

// Validate validates this remove agent body
func (o *RemoveAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove agent body based on context it is used
func (o *RemoveAgentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveAgentBody) UnmarshalBinary(b []byte) error {
	var res RemoveAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RemoveAgentDefaultBody remove agent default body
swagger:model RemoveAgentDefaultBody
*/
type RemoveAgentDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*RemoveAgentDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this remove agent default body
func (o *RemoveAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveAgentDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this remove agent default body based on the context it is used
func (o *RemoveAgentDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveAgentDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoveAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoveAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RemoveAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RemoveAgentDefaultBodyDetailsItems0 remove agent default body details items0
swagger:model RemoveAgentDefaultBodyDetailsItems0
*/
type RemoveAgentDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this remove agent default body details items0
func (o *RemoveAgentDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this remove agent default body details items0 based on context it is used
func (o *RemoveAgentDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveAgentDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveAgentDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res RemoveAgentDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
