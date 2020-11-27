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
	"github.com/go-openapi/swag"
)

// NewStorageServiceRepairStatusParams creates a new StorageServiceRepairStatusParams object
// with the default values initialized.
func NewStorageServiceRepairStatusParams() *StorageServiceRepairStatusParams {
	var ()
	return &StorageServiceRepairStatusParams{

		requestTimeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceRepairStatusParamsWithTimeout creates a new StorageServiceRepairStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceRepairStatusParamsWithTimeout(timeout time.Duration) *StorageServiceRepairStatusParams {
	var ()
	return &StorageServiceRepairStatusParams{

		requestTimeout: timeout,
	}
}

// NewStorageServiceRepairStatusParamsWithContext creates a new StorageServiceRepairStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceRepairStatusParamsWithContext(ctx context.Context) *StorageServiceRepairStatusParams {
	var ()
	return &StorageServiceRepairStatusParams{

		Context: ctx,
	}
}

// NewStorageServiceRepairStatusParamsWithHTTPClient creates a new StorageServiceRepairStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceRepairStatusParamsWithHTTPClient(client *http.Client) *StorageServiceRepairStatusParams {
	var ()
	return &StorageServiceRepairStatusParams{
		HTTPClient: client,
	}
}

/*StorageServiceRepairStatusParams contains all the parameters to send to the API endpoint
for the storage service repair status operation typically these are written to a http.Request
*/
type StorageServiceRepairStatusParams struct {

	/*ID
	  The repair ID to check for status

	*/
	ID int32
	/*Timeout
	  Seconds to wait before the query returns even if the repair is not finished. The value -1 or not providing this parameter means no timeout

	*/
	Timeout *int64

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithRequestTimeout adds the timeout to the storage service repair status params
func (o *StorageServiceRepairStatusParams) WithRequestTimeout(timeout time.Duration) *StorageServiceRepairStatusParams {
	o.SetRequestTimeout(timeout)
	return o
}

// SetRequestTimeout adds the timeout to the storage service repair status params
func (o *StorageServiceRepairStatusParams) SetRequestTimeout(timeout time.Duration) {
	o.requestTimeout = timeout
}

// WithContext adds the context to the storage service repair status params
func (o *StorageServiceRepairStatusParams) WithContext(ctx context.Context) *StorageServiceRepairStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service repair status params
func (o *StorageServiceRepairStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service repair status params
func (o *StorageServiceRepairStatusParams) WithHTTPClient(client *http.Client) *StorageServiceRepairStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service repair status params
func (o *StorageServiceRepairStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the storage service repair status params
func (o *StorageServiceRepairStatusParams) WithID(id int32) *StorageServiceRepairStatusParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the storage service repair status params
func (o *StorageServiceRepairStatusParams) SetID(id int32) {
	o.ID = id
}

// WithTimeout adds the timeout to the storage service repair status params
func (o *StorageServiceRepairStatusParams) WithTimeout(timeout *int64) *StorageServiceRepairStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service repair status params
func (o *StorageServiceRepairStatusParams) SetTimeout(timeout *int64) {
	o.Timeout = timeout
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceRepairStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := swag.FormatInt32(qrID)
	if qID != "" {
		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout int64
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := swag.FormatInt64(qrTimeout)
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
