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

// NewFindConfigPermissionsCacheMaxEntriesParams creates a new FindConfigPermissionsCacheMaxEntriesParams object
// with the default values initialized.
func NewFindConfigPermissionsCacheMaxEntriesParams() *FindConfigPermissionsCacheMaxEntriesParams {

	return &FindConfigPermissionsCacheMaxEntriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigPermissionsCacheMaxEntriesParamsWithTimeout creates a new FindConfigPermissionsCacheMaxEntriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigPermissionsCacheMaxEntriesParamsWithTimeout(timeout time.Duration) *FindConfigPermissionsCacheMaxEntriesParams {

	return &FindConfigPermissionsCacheMaxEntriesParams{

		timeout: timeout,
	}
}

// NewFindConfigPermissionsCacheMaxEntriesParamsWithContext creates a new FindConfigPermissionsCacheMaxEntriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigPermissionsCacheMaxEntriesParamsWithContext(ctx context.Context) *FindConfigPermissionsCacheMaxEntriesParams {

	return &FindConfigPermissionsCacheMaxEntriesParams{

		Context: ctx,
	}
}

// NewFindConfigPermissionsCacheMaxEntriesParamsWithHTTPClient creates a new FindConfigPermissionsCacheMaxEntriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigPermissionsCacheMaxEntriesParamsWithHTTPClient(client *http.Client) *FindConfigPermissionsCacheMaxEntriesParams {

	return &FindConfigPermissionsCacheMaxEntriesParams{
		HTTPClient: client,
	}
}

/*FindConfigPermissionsCacheMaxEntriesParams contains all the parameters to send to the API endpoint
for the find config permissions cache max entries operation typically these are written to a http.Request
*/
type FindConfigPermissionsCacheMaxEntriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config permissions cache max entries params
func (o *FindConfigPermissionsCacheMaxEntriesParams) WithTimeout(timeout time.Duration) *FindConfigPermissionsCacheMaxEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config permissions cache max entries params
func (o *FindConfigPermissionsCacheMaxEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config permissions cache max entries params
func (o *FindConfigPermissionsCacheMaxEntriesParams) WithContext(ctx context.Context) *FindConfigPermissionsCacheMaxEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config permissions cache max entries params
func (o *FindConfigPermissionsCacheMaxEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config permissions cache max entries params
func (o *FindConfigPermissionsCacheMaxEntriesParams) WithHTTPClient(client *http.Client) *FindConfigPermissionsCacheMaxEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config permissions cache max entries params
func (o *FindConfigPermissionsCacheMaxEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigPermissionsCacheMaxEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
