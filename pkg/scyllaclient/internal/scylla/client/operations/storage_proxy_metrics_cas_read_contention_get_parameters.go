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

// NewStorageProxyMetricsCasReadContentionGetParams creates a new StorageProxyMetricsCasReadContentionGetParams object
// with the default values initialized.
func NewStorageProxyMetricsCasReadContentionGetParams() *StorageProxyMetricsCasReadContentionGetParams {

	return &StorageProxyMetricsCasReadContentionGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsCasReadContentionGetParamsWithTimeout creates a new StorageProxyMetricsCasReadContentionGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsCasReadContentionGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsCasReadContentionGetParams {

	return &StorageProxyMetricsCasReadContentionGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsCasReadContentionGetParamsWithContext creates a new StorageProxyMetricsCasReadContentionGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsCasReadContentionGetParamsWithContext(ctx context.Context) *StorageProxyMetricsCasReadContentionGetParams {

	return &StorageProxyMetricsCasReadContentionGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsCasReadContentionGetParamsWithHTTPClient creates a new StorageProxyMetricsCasReadContentionGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsCasReadContentionGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsCasReadContentionGetParams {

	return &StorageProxyMetricsCasReadContentionGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyMetricsCasReadContentionGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics cas read contention get operation typically these are written to a http.Request
*/
type StorageProxyMetricsCasReadContentionGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics cas read contention get params
func (o *StorageProxyMetricsCasReadContentionGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsCasReadContentionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics cas read contention get params
func (o *StorageProxyMetricsCasReadContentionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics cas read contention get params
func (o *StorageProxyMetricsCasReadContentionGetParams) WithContext(ctx context.Context) *StorageProxyMetricsCasReadContentionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics cas read contention get params
func (o *StorageProxyMetricsCasReadContentionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics cas read contention get params
func (o *StorageProxyMetricsCasReadContentionGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsCasReadContentionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics cas read contention get params
func (o *StorageProxyMetricsCasReadContentionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsCasReadContentionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
