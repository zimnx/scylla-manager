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
	"github.com/go-openapi/strfmt"
)

// NewColumnFamilyMetricsMemtableOffHeapSizeGetParams creates a new ColumnFamilyMetricsMemtableOffHeapSizeGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsMemtableOffHeapSizeGetParams() *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {

	return &ColumnFamilyMetricsMemtableOffHeapSizeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsMemtableOffHeapSizeGetParamsWithTimeout creates a new ColumnFamilyMetricsMemtableOffHeapSizeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsMemtableOffHeapSizeGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {

	return &ColumnFamilyMetricsMemtableOffHeapSizeGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsMemtableOffHeapSizeGetParamsWithContext creates a new ColumnFamilyMetricsMemtableOffHeapSizeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsMemtableOffHeapSizeGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {

	return &ColumnFamilyMetricsMemtableOffHeapSizeGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsMemtableOffHeapSizeGetParamsWithHTTPClient creates a new ColumnFamilyMetricsMemtableOffHeapSizeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsMemtableOffHeapSizeGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {

	return &ColumnFamilyMetricsMemtableOffHeapSizeGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsMemtableOffHeapSizeGetParams contains all the parameters to send to the API endpoint
for the column family metrics memtable off heap size get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsMemtableOffHeapSizeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics memtable off heap size get params
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics memtable off heap size get params
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics memtable off heap size get params
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics memtable off heap size get params
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics memtable off heap size get params
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsMemtableOffHeapSizeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics memtable off heap size get params
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsMemtableOffHeapSizeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
