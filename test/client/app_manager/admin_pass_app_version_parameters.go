// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewAdminPassAppVersionParams creates a new AdminPassAppVersionParams object
// with the default values initialized.
func NewAdminPassAppVersionParams() *AdminPassAppVersionParams {
	var ()
	return &AdminPassAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminPassAppVersionParamsWithTimeout creates a new AdminPassAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminPassAppVersionParamsWithTimeout(timeout time.Duration) *AdminPassAppVersionParams {
	var ()
	return &AdminPassAppVersionParams{

		timeout: timeout,
	}
}

// NewAdminPassAppVersionParamsWithContext creates a new AdminPassAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminPassAppVersionParamsWithContext(ctx context.Context) *AdminPassAppVersionParams {
	var ()
	return &AdminPassAppVersionParams{

		Context: ctx,
	}
}

// NewAdminPassAppVersionParamsWithHTTPClient creates a new AdminPassAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminPassAppVersionParamsWithHTTPClient(client *http.Client) *AdminPassAppVersionParams {
	var ()
	return &AdminPassAppVersionParams{
		HTTPClient: client,
	}
}

/*AdminPassAppVersionParams contains all the parameters to send to the API endpoint
for the admin pass app version operation typically these are written to a http.Request
*/
type AdminPassAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixPassAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin pass app version params
func (o *AdminPassAppVersionParams) WithTimeout(timeout time.Duration) *AdminPassAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin pass app version params
func (o *AdminPassAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin pass app version params
func (o *AdminPassAppVersionParams) WithContext(ctx context.Context) *AdminPassAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin pass app version params
func (o *AdminPassAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin pass app version params
func (o *AdminPassAppVersionParams) WithHTTPClient(client *http.Client) *AdminPassAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin pass app version params
func (o *AdminPassAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin pass app version params
func (o *AdminPassAppVersionParams) WithBody(body *models.OpenpitrixPassAppVersionRequest) *AdminPassAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin pass app version params
func (o *AdminPassAppVersionParams) SetBody(body *models.OpenpitrixPassAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AdminPassAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
