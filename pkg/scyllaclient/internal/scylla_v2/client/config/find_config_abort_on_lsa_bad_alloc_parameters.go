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

// NewFindConfigAbortOnLsaBadAllocParams creates a new FindConfigAbortOnLsaBadAllocParams object
// with the default values initialized.
func NewFindConfigAbortOnLsaBadAllocParams() *FindConfigAbortOnLsaBadAllocParams {

	return &FindConfigAbortOnLsaBadAllocParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigAbortOnLsaBadAllocParamsWithTimeout creates a new FindConfigAbortOnLsaBadAllocParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigAbortOnLsaBadAllocParamsWithTimeout(timeout time.Duration) *FindConfigAbortOnLsaBadAllocParams {

	return &FindConfigAbortOnLsaBadAllocParams{

		timeout: timeout,
	}
}

// NewFindConfigAbortOnLsaBadAllocParamsWithContext creates a new FindConfigAbortOnLsaBadAllocParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigAbortOnLsaBadAllocParamsWithContext(ctx context.Context) *FindConfigAbortOnLsaBadAllocParams {

	return &FindConfigAbortOnLsaBadAllocParams{

		Context: ctx,
	}
}

// NewFindConfigAbortOnLsaBadAllocParamsWithHTTPClient creates a new FindConfigAbortOnLsaBadAllocParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigAbortOnLsaBadAllocParamsWithHTTPClient(client *http.Client) *FindConfigAbortOnLsaBadAllocParams {

	return &FindConfigAbortOnLsaBadAllocParams{
		HTTPClient: client,
	}
}

/*FindConfigAbortOnLsaBadAllocParams contains all the parameters to send to the API endpoint
for the find config abort on lsa bad alloc operation typically these are written to a http.Request
*/
type FindConfigAbortOnLsaBadAllocParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config abort on lsa bad alloc params
func (o *FindConfigAbortOnLsaBadAllocParams) WithTimeout(timeout time.Duration) *FindConfigAbortOnLsaBadAllocParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config abort on lsa bad alloc params
func (o *FindConfigAbortOnLsaBadAllocParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config abort on lsa bad alloc params
func (o *FindConfigAbortOnLsaBadAllocParams) WithContext(ctx context.Context) *FindConfigAbortOnLsaBadAllocParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config abort on lsa bad alloc params
func (o *FindConfigAbortOnLsaBadAllocParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config abort on lsa bad alloc params
func (o *FindConfigAbortOnLsaBadAllocParams) WithHTTPClient(client *http.Client) *FindConfigAbortOnLsaBadAllocParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config abort on lsa bad alloc params
func (o *FindConfigAbortOnLsaBadAllocParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigAbortOnLsaBadAllocParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
