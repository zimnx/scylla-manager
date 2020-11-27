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

// NewStorageProxyMetricsRangeTimeoutsRatesGetParams creates a new StorageProxyMetricsRangeTimeoutsRatesGetParams object
// with the default values initialized.
func NewStorageProxyMetricsRangeTimeoutsRatesGetParams() *StorageProxyMetricsRangeTimeoutsRatesGetParams {

	return &StorageProxyMetricsRangeTimeoutsRatesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsRangeTimeoutsRatesGetParamsWithTimeout creates a new StorageProxyMetricsRangeTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsRangeTimeoutsRatesGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsRangeTimeoutsRatesGetParams {

	return &StorageProxyMetricsRangeTimeoutsRatesGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsRangeTimeoutsRatesGetParamsWithContext creates a new StorageProxyMetricsRangeTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsRangeTimeoutsRatesGetParamsWithContext(ctx context.Context) *StorageProxyMetricsRangeTimeoutsRatesGetParams {

	return &StorageProxyMetricsRangeTimeoutsRatesGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsRangeTimeoutsRatesGetParamsWithHTTPClient creates a new StorageProxyMetricsRangeTimeoutsRatesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsRangeTimeoutsRatesGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsRangeTimeoutsRatesGetParams {

	return &StorageProxyMetricsRangeTimeoutsRatesGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsRangeTimeoutsRatesGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics range timeouts rates get operation typically these are written to a http.Request
*/
type StorageProxyMetricsRangeTimeoutsRatesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics range timeouts rates get params
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsRangeTimeoutsRatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics range timeouts rates get params
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics range timeouts rates get params
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) WithContext(ctx context.Context) *StorageProxyMetricsRangeTimeoutsRatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics range timeouts rates get params
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics range timeouts rates get params
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsRangeTimeoutsRatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics range timeouts rates get params
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsRangeTimeoutsRatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
