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

// NewStorageProxyReadRepairAttemptedGetParams creates a new StorageProxyReadRepairAttemptedGetParams object
// with the default values initialized.
func NewStorageProxyReadRepairAttemptedGetParams() *StorageProxyReadRepairAttemptedGetParams {

	return &StorageProxyReadRepairAttemptedGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyReadRepairAttemptedGetParamsWithTimeout creates a new StorageProxyReadRepairAttemptedGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyReadRepairAttemptedGetParamsWithTimeout(timeout time.Duration) *StorageProxyReadRepairAttemptedGetParams {

	return &StorageProxyReadRepairAttemptedGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyReadRepairAttemptedGetParamsWithContext creates a new StorageProxyReadRepairAttemptedGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyReadRepairAttemptedGetParamsWithContext(ctx context.Context) *StorageProxyReadRepairAttemptedGetParams {

	return &StorageProxyReadRepairAttemptedGetParams{

		Context: ctx,
	}
}

// NewStorageProxyReadRepairAttemptedGetParamsWithHTTPClient creates a new StorageProxyReadRepairAttemptedGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyReadRepairAttemptedGetParamsWithHTTPClient(client *http.Client) *StorageProxyReadRepairAttemptedGetParams {

	return &StorageProxyReadRepairAttemptedGetParams{
		HTTPClient: client,
	}
}

/*StorageProxyReadRepairAttemptedGetParams contains all the parameters to send to the API endpoint
for the storage proxy read repair attempted get operation typically these are written to a http.Request
*/
type StorageProxyReadRepairAttemptedGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy read repair attempted get params
func (o *StorageProxyReadRepairAttemptedGetParams) WithTimeout(timeout time.Duration) *StorageProxyReadRepairAttemptedGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy read repair attempted get params
func (o *StorageProxyReadRepairAttemptedGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy read repair attempted get params
func (o *StorageProxyReadRepairAttemptedGetParams) WithContext(ctx context.Context) *StorageProxyReadRepairAttemptedGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy read repair attempted get params
func (o *StorageProxyReadRepairAttemptedGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy read repair attempted get params
func (o *StorageProxyReadRepairAttemptedGetParams) WithHTTPClient(client *http.Client) *StorageProxyReadRepairAttemptedGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy read repair attempted get params
func (o *StorageProxyReadRepairAttemptedGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyReadRepairAttemptedGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
