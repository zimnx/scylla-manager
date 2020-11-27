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

// NewFindConfigPrometheusPrefixParams creates a new FindConfigPrometheusPrefixParams object
// with the default values initialized.
func NewFindConfigPrometheusPrefixParams() *FindConfigPrometheusPrefixParams {

	return &FindConfigPrometheusPrefixParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigPrometheusPrefixParamsWithTimeout creates a new FindConfigPrometheusPrefixParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigPrometheusPrefixParamsWithTimeout(timeout time.Duration) *FindConfigPrometheusPrefixParams {

	return &FindConfigPrometheusPrefixParams{

		timeout: timeout,
	}
}

// NewFindConfigPrometheusPrefixParamsWithContext creates a new FindConfigPrometheusPrefixParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigPrometheusPrefixParamsWithContext(ctx context.Context) *FindConfigPrometheusPrefixParams {

	return &FindConfigPrometheusPrefixParams{

		Context: ctx,
	}
}

// NewFindConfigPrometheusPrefixParamsWithHTTPClient creates a new FindConfigPrometheusPrefixParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigPrometheusPrefixParamsWithHTTPClient(client *http.Client) *FindConfigPrometheusPrefixParams {

	return &FindConfigPrometheusPrefixParams{
		HTTPClient: client,
	}
}

/*FindConfigPrometheusPrefixParams contains all the parameters to send to the API endpoint
for the find config prometheus prefix operation typically these are written to a http.Request
*/
type FindConfigPrometheusPrefixParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config prometheus prefix params
func (o *FindConfigPrometheusPrefixParams) WithTimeout(timeout time.Duration) *FindConfigPrometheusPrefixParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config prometheus prefix params
func (o *FindConfigPrometheusPrefixParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config prometheus prefix params
func (o *FindConfigPrometheusPrefixParams) WithContext(ctx context.Context) *FindConfigPrometheusPrefixParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config prometheus prefix params
func (o *FindConfigPrometheusPrefixParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config prometheus prefix params
func (o *FindConfigPrometheusPrefixParams) WithHTTPClient(client *http.Client) *FindConfigPrometheusPrefixParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config prometheus prefix params
func (o *FindConfigPrometheusPrefixParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigPrometheusPrefixParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
