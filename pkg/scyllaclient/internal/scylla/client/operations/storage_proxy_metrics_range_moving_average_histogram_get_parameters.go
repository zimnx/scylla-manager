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

// NewStorageProxyMetricsRangeMovingAverageHistogramGetParams creates a new StorageProxyMetricsRangeMovingAverageHistogramGetParams object
// with the default values initialized.
func NewStorageProxyMetricsRangeMovingAverageHistogramGetParams() *StorageProxyMetricsRangeMovingAverageHistogramGetParams {

	return &StorageProxyMetricsRangeMovingAverageHistogramGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsRangeMovingAverageHistogramGetParamsWithTimeout creates a new StorageProxyMetricsRangeMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsRangeMovingAverageHistogramGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsRangeMovingAverageHistogramGetParams {

	return &StorageProxyMetricsRangeMovingAverageHistogramGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsRangeMovingAverageHistogramGetParamsWithContext creates a new StorageProxyMetricsRangeMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsRangeMovingAverageHistogramGetParamsWithContext(ctx context.Context) *StorageProxyMetricsRangeMovingAverageHistogramGetParams {

	return &StorageProxyMetricsRangeMovingAverageHistogramGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsRangeMovingAverageHistogramGetParamsWithHTTPClient creates a new StorageProxyMetricsRangeMovingAverageHistogramGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsRangeMovingAverageHistogramGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsRangeMovingAverageHistogramGetParams {

	return &StorageProxyMetricsRangeMovingAverageHistogramGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsRangeMovingAverageHistogramGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics range moving average histogram get operation typically these are written to a http.Request
*/
type StorageProxyMetricsRangeMovingAverageHistogramGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics range moving average histogram get params
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsRangeMovingAverageHistogramGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics range moving average histogram get params
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics range moving average histogram get params
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) WithContext(ctx context.Context) *StorageProxyMetricsRangeMovingAverageHistogramGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics range moving average histogram get params
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics range moving average histogram get params
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsRangeMovingAverageHistogramGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics range moving average histogram get params
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsRangeMovingAverageHistogramGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
