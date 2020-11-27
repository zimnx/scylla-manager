// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindConfigViewHintsDirectoryParams creates a new FindConfigViewHintsDirectoryParams object
// with the default values initialized.
func NewFindConfigViewHintsDirectoryParams() *FindConfigViewHintsDirectoryParams {

	return &FindConfigViewHintsDirectoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigViewHintsDirectoryParamsWithTimeout creates a new FindConfigViewHintsDirectoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigViewHintsDirectoryParamsWithTimeout(timeout time.Duration) *FindConfigViewHintsDirectoryParams {

	return &FindConfigViewHintsDirectoryParams{

		timeout: timeout,
	}
}

// NewFindConfigViewHintsDirectoryParamsWithContext creates a new FindConfigViewHintsDirectoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigViewHintsDirectoryParamsWithContext(ctx context.Context) *FindConfigViewHintsDirectoryParams {

	return &FindConfigViewHintsDirectoryParams{

		Context: ctx,
	}
}

// NewFindConfigViewHintsDirectoryParamsWithHTTPClient creates a new FindConfigViewHintsDirectoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigViewHintsDirectoryParamsWithHTTPClient(client *http.Client) *FindConfigViewHintsDirectoryParams {

	return &FindConfigViewHintsDirectoryParams{
		HTTPClient: client,
	}
}

/*FindConfigViewHintsDirectoryParams contains all the parameters to send to the API endpoint
for the find config view hints directory operation typically these are written to a http.Request
*/
type FindConfigViewHintsDirectoryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config view hints directory params
func (o *FindConfigViewHintsDirectoryParams) WithTimeout(timeout time.Duration) *FindConfigViewHintsDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config view hints directory params
func (o *FindConfigViewHintsDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config view hints directory params
func (o *FindConfigViewHintsDirectoryParams) WithContext(ctx context.Context) *FindConfigViewHintsDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config view hints directory params
func (o *FindConfigViewHintsDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config view hints directory params
func (o *FindConfigViewHintsDirectoryParams) WithHTTPClient(client *http.Client) *FindConfigViewHintsDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config view hints directory params
func (o *FindConfigViewHintsDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigViewHintsDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
