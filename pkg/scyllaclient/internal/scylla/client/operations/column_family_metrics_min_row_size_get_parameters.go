// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewColumnFamilyMetricsMinRowSizeGetParams creates a new ColumnFamilyMetricsMinRowSizeGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsMinRowSizeGetParams() *ColumnFamilyMetricsMinRowSizeGetParams {

	return &ColumnFamilyMetricsMinRowSizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsMinRowSizeGetParamsWithTimeout creates a new ColumnFamilyMetricsMinRowSizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsMinRowSizeGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsMinRowSizeGetParams {

	return &ColumnFamilyMetricsMinRowSizeGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsMinRowSizeGetParamsWithContext creates a new ColumnFamilyMetricsMinRowSizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsMinRowSizeGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsMinRowSizeGetParams {

	return &ColumnFamilyMetricsMinRowSizeGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsMinRowSizeGetParamsWithHTTPClient creates a new ColumnFamilyMetricsMinRowSizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsMinRowSizeGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsMinRowSizeGetParams {

	return &ColumnFamilyMetricsMinRowSizeGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsMinRowSizeGetParams contains all the parameters to send to the API endpoint
for the column family metrics min row size get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsMinRowSizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics min row size get params
func (o *ColumnFamilyMetricsMinRowSizeGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsMinRowSizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics min row size get params
func (o *ColumnFamilyMetricsMinRowSizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics min row size get params
func (o *ColumnFamilyMetricsMinRowSizeGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsMinRowSizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics min row size get params
func (o *ColumnFamilyMetricsMinRowSizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics min row size get params
func (o *ColumnFamilyMetricsMinRowSizeGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsMinRowSizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics min row size get params
func (o *ColumnFamilyMetricsMinRowSizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsMinRowSizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}