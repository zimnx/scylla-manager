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

// NewStorageServiceHostidLocalGetParams creates a new StorageServiceHostidLocalGetParams object
// with the default values initialized.
func NewStorageServiceHostidLocalGetParams() *StorageServiceHostidLocalGetParams {

	return &StorageServiceHostidLocalGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceHostidLocalGetParamsWithTimeout creates a new StorageServiceHostidLocalGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceHostidLocalGetParamsWithTimeout(timeout time.Duration) *StorageServiceHostidLocalGetParams {

	return &StorageServiceHostidLocalGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceHostidLocalGetParamsWithContext creates a new StorageServiceHostidLocalGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceHostidLocalGetParamsWithContext(ctx context.Context) *StorageServiceHostidLocalGetParams {

	return &StorageServiceHostidLocalGetParams{

		Context: ctx,
	}
}

// NewStorageServiceHostidLocalGetParamsWithHTTPClient creates a new StorageServiceHostidLocalGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceHostidLocalGetParamsWithHTTPClient(client *http.Client) *StorageServiceHostidLocalGetParams {

	return &StorageServiceHostidLocalGetParams{
		HTTPClient: client,
	}
}

/*StorageServiceHostidLocalGetParams contains all the parameters to send to the API endpoint
for the storage service hostid local get operation typically these are written to a http.Request
*/
type StorageServiceHostidLocalGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service hostid local get params
func (o *StorageServiceHostidLocalGetParams) WithTimeout(timeout time.Duration) *StorageServiceHostidLocalGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service hostid local get params
func (o *StorageServiceHostidLocalGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service hostid local get params
func (o *StorageServiceHostidLocalGetParams) WithContext(ctx context.Context) *StorageServiceHostidLocalGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service hostid local get params
func (o *StorageServiceHostidLocalGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service hostid local get params
func (o *StorageServiceHostidLocalGetParams) WithHTTPClient(client *http.Client) *StorageServiceHostidLocalGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service hostid local get params
func (o *StorageServiceHostidLocalGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceHostidLocalGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
