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

// NewCacheServiceMetricsRowRequestsMovingAvrageGetParams creates a new CacheServiceMetricsRowRequestsMovingAvrageGetParams object
// with the default values initialized.
func NewCacheServiceMetricsRowRequestsMovingAvrageGetParams() *CacheServiceMetricsRowRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsRowRequestsMovingAvrageGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsRowRequestsMovingAvrageGetParamsWithTimeout creates a new CacheServiceMetricsRowRequestsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsRowRequestsMovingAvrageGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsRowRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsRowRequestsMovingAvrageGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsRowRequestsMovingAvrageGetParamsWithContext creates a new CacheServiceMetricsRowRequestsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsRowRequestsMovingAvrageGetParamsWithContext(ctx context.Context) *CacheServiceMetricsRowRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsRowRequestsMovingAvrageGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsRowRequestsMovingAvrageGetParamsWithHTTPClient creates a new CacheServiceMetricsRowRequestsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsRowRequestsMovingAvrageGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsRowRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsRowRequestsMovingAvrageGetParams{
		HTTPClient: client,
	}
}

/*CacheServiceMetricsRowRequestsMovingAvrageGetParams contains all the parameters to send to the API endpoint
for the cache service metrics row requests moving avrage get operation typically these are written to a http.Request
*/
type CacheServiceMetricsRowRequestsMovingAvrageGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics row requests moving avrage get params
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsRowRequestsMovingAvrageGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics row requests moving avrage get params
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics row requests moving avrage get params
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) WithContext(ctx context.Context) *CacheServiceMetricsRowRequestsMovingAvrageGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics row requests moving avrage get params
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics row requests moving avrage get params
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsRowRequestsMovingAvrageGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics row requests moving avrage get params
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsRowRequestsMovingAvrageGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
