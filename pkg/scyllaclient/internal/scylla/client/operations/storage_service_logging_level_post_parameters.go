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

// NewStorageServiceLoggingLevelPostParams creates a new StorageServiceLoggingLevelPostParams object
// with the default values initialized.
func NewStorageServiceLoggingLevelPostParams() *StorageServiceLoggingLevelPostParams {
	var ()
	return &StorageServiceLoggingLevelPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceLoggingLevelPostParamsWithTimeout creates a new StorageServiceLoggingLevelPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceLoggingLevelPostParamsWithTimeout(timeout time.Duration) *StorageServiceLoggingLevelPostParams {
	var ()
	return &StorageServiceLoggingLevelPostParams{

		timeout: timeout,
	}
}

// NewStorageServiceLoggingLevelPostParamsWithContext creates a new StorageServiceLoggingLevelPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceLoggingLevelPostParamsWithContext(ctx context.Context) *StorageServiceLoggingLevelPostParams {
	var ()
	return &StorageServiceLoggingLevelPostParams{

		Context: ctx,
	}
}

// NewStorageServiceLoggingLevelPostParamsWithHTTPClient creates a new StorageServiceLoggingLevelPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceLoggingLevelPostParamsWithHTTPClient(client *http.Client) *StorageServiceLoggingLevelPostParams {
	var ()
	return &StorageServiceLoggingLevelPostParams{
		HTTPClient: client,
	}
}

/*StorageServiceLoggingLevelPostParams contains all the parameters to send to the API endpoint
for the storage service logging level post operation typically these are written to a http.Request
*/
type StorageServiceLoggingLevelPostParams struct {

	/*ClassQualifier
	  The logger's classQualifer

	*/
	ClassQualifier string
	/*Level
	  The log level

	*/
	Level string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) WithTimeout(timeout time.Duration) *StorageServiceLoggingLevelPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) WithContext(ctx context.Context) *StorageServiceLoggingLevelPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) WithHTTPClient(client *http.Client) *StorageServiceLoggingLevelPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassQualifier adds the classQualifier to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) WithClassQualifier(classQualifier string) *StorageServiceLoggingLevelPostParams {
	o.SetClassQualifier(classQualifier)
	return o
}

// SetClassQualifier adds the classQualifier to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) SetClassQualifier(classQualifier string) {
	o.ClassQualifier = classQualifier
}

// WithLevel adds the level to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) WithLevel(level string) *StorageServiceLoggingLevelPostParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the storage service logging level post params
func (o *StorageServiceLoggingLevelPostParams) SetLevel(level string) {
	o.Level = level
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceLoggingLevelPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param class_qualifier
	qrClassQualifier := o.ClassQualifier
	qClassQualifier := qrClassQualifier
	if qClassQualifier != "" {
		if err := r.SetQueryParam("class_qualifier", qClassQualifier); err != nil {
			return err
		}
	}

	// query param level
	qrLevel := o.Level
	qLevel := qrLevel
	if qLevel != "" {
		if err := r.SetQueryParam("level", qLevel); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
