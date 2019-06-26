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

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigCommitlogUseODsyncParams creates a new FindConfigCommitlogUseODsyncParams object
// with the default values initialized.
func NewFindConfigCommitlogUseODsyncParams() *FindConfigCommitlogUseODsyncParams {

	return &FindConfigCommitlogUseODsyncParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogUseODsyncParamsWithTimeout creates a new FindConfigCommitlogUseODsyncParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCommitlogUseODsyncParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogUseODsyncParams {

	return &FindConfigCommitlogUseODsyncParams{

		timeout: timeout,
	}
}

// NewFindConfigCommitlogUseODsyncParamsWithContext creates a new FindConfigCommitlogUseODsyncParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCommitlogUseODsyncParamsWithContext(ctx context.Context) *FindConfigCommitlogUseODsyncParams {

	return &FindConfigCommitlogUseODsyncParams{

		Context: ctx,
	}
}

// NewFindConfigCommitlogUseODsyncParamsWithHTTPClient creates a new FindConfigCommitlogUseODsyncParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCommitlogUseODsyncParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogUseODsyncParams {

	return &FindConfigCommitlogUseODsyncParams{
		HTTPClient: client,
	}
}

/*FindConfigCommitlogUseODsyncParams contains all the parameters to send to the API endpoint
for the find config commitlog use o dsync operation typically these are written to a http.Request
*/
type FindConfigCommitlogUseODsyncParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseODsyncParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogUseODsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseODsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseODsyncParams) WithContext(ctx context.Context) *FindConfigCommitlogUseODsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseODsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseODsyncParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogUseODsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog use o dsync params
func (o *FindConfigCommitlogUseODsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogUseODsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
