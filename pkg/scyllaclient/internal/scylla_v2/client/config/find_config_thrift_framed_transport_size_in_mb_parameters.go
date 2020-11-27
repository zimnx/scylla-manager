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

// NewFindConfigThriftFramedTransportSizeInMbParams creates a new FindConfigThriftFramedTransportSizeInMbParams object
// with the default values initialized.
func NewFindConfigThriftFramedTransportSizeInMbParams() *FindConfigThriftFramedTransportSizeInMbParams {

	return &FindConfigThriftFramedTransportSizeInMbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigThriftFramedTransportSizeInMbParamsWithTimeout creates a new FindConfigThriftFramedTransportSizeInMbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigThriftFramedTransportSizeInMbParamsWithTimeout(timeout time.Duration) *FindConfigThriftFramedTransportSizeInMbParams {

	return &FindConfigThriftFramedTransportSizeInMbParams{

		timeout: timeout,
	}
}

// NewFindConfigThriftFramedTransportSizeInMbParamsWithContext creates a new FindConfigThriftFramedTransportSizeInMbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigThriftFramedTransportSizeInMbParamsWithContext(ctx context.Context) *FindConfigThriftFramedTransportSizeInMbParams {

	return &FindConfigThriftFramedTransportSizeInMbParams{

		Context: ctx,
	}
}

// NewFindConfigThriftFramedTransportSizeInMbParamsWithHTTPClient creates a new FindConfigThriftFramedTransportSizeInMbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigThriftFramedTransportSizeInMbParamsWithHTTPClient(client *http.Client) *FindConfigThriftFramedTransportSizeInMbParams {

	return &FindConfigThriftFramedTransportSizeInMbParams{
		HTTPClient: client,
	}
}

/*FindConfigThriftFramedTransportSizeInMbParams contains all the parameters to send to the API endpoint
for the find config thrift framed transport size in mb operation typically these are written to a http.Request
*/
type FindConfigThriftFramedTransportSizeInMbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config thrift framed transport size in mb params
func (o *FindConfigThriftFramedTransportSizeInMbParams) WithTimeout(timeout time.Duration) *FindConfigThriftFramedTransportSizeInMbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config thrift framed transport size in mb params
func (o *FindConfigThriftFramedTransportSizeInMbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config thrift framed transport size in mb params
func (o *FindConfigThriftFramedTransportSizeInMbParams) WithContext(ctx context.Context) *FindConfigThriftFramedTransportSizeInMbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config thrift framed transport size in mb params
func (o *FindConfigThriftFramedTransportSizeInMbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config thrift framed transport size in mb params
func (o *FindConfigThriftFramedTransportSizeInMbParams) WithHTTPClient(client *http.Client) *FindConfigThriftFramedTransportSizeInMbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config thrift framed transport size in mb params
func (o *FindConfigThriftFramedTransportSizeInMbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigThriftFramedTransportSizeInMbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
