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

// NewFindConfigEnableSstablesMcFormatParams creates a new FindConfigEnableSstablesMcFormatParams object
// with the default values initialized.
func NewFindConfigEnableSstablesMcFormatParams() *FindConfigEnableSstablesMcFormatParams {

	return &FindConfigEnableSstablesMcFormatParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigEnableSstablesMcFormatParamsWithTimeout creates a new FindConfigEnableSstablesMcFormatParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigEnableSstablesMcFormatParamsWithTimeout(timeout time.Duration) *FindConfigEnableSstablesMcFormatParams {

	return &FindConfigEnableSstablesMcFormatParams{

		timeout: timeout,
	}
}

// NewFindConfigEnableSstablesMcFormatParamsWithContext creates a new FindConfigEnableSstablesMcFormatParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigEnableSstablesMcFormatParamsWithContext(ctx context.Context) *FindConfigEnableSstablesMcFormatParams {

	return &FindConfigEnableSstablesMcFormatParams{

		Context: ctx,
	}
}

// NewFindConfigEnableSstablesMcFormatParamsWithHTTPClient creates a new FindConfigEnableSstablesMcFormatParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigEnableSstablesMcFormatParamsWithHTTPClient(client *http.Client) *FindConfigEnableSstablesMcFormatParams {

	return &FindConfigEnableSstablesMcFormatParams{
		HTTPClient: client,
	}
}

/*FindConfigEnableSstablesMcFormatParams contains all the parameters to send to the API endpoint
for the find config enable sstables mc format operation typically these are written to a http.Request
*/
type FindConfigEnableSstablesMcFormatParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config enable sstables mc format params
func (o *FindConfigEnableSstablesMcFormatParams) WithTimeout(timeout time.Duration) *FindConfigEnableSstablesMcFormatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config enable sstables mc format params
func (o *FindConfigEnableSstablesMcFormatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config enable sstables mc format params
func (o *FindConfigEnableSstablesMcFormatParams) WithContext(ctx context.Context) *FindConfigEnableSstablesMcFormatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config enable sstables mc format params
func (o *FindConfigEnableSstablesMcFormatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config enable sstables mc format params
func (o *FindConfigEnableSstablesMcFormatParams) WithHTTPClient(client *http.Client) *FindConfigEnableSstablesMcFormatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config enable sstables mc format params
func (o *FindConfigEnableSstablesMcFormatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigEnableSstablesMcFormatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
