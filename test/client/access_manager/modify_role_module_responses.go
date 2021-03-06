// Code generated by go-swagger; DO NOT EDIT.

package access_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ModifyRoleModuleReader is a Reader for the ModifyRoleModule structure.
type ModifyRoleModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyRoleModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyRoleModuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyRoleModuleOK creates a ModifyRoleModuleOK with default headers values
func NewModifyRoleModuleOK() *ModifyRoleModuleOK {
	return &ModifyRoleModuleOK{}
}

/*ModifyRoleModuleOK handles this case with default header values.

A successful response.
*/
type ModifyRoleModuleOK struct {
	Payload *models.OpenpitrixModifyRoleModuleResponse
}

func (o *ModifyRoleModuleOK) Error() string {
	return fmt.Sprintf("[POST /v1/roles:module][%d] modifyRoleModuleOK  %+v", 200, o.Payload)
}

func (o *ModifyRoleModuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyRoleModuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
