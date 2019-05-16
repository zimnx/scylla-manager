// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageServiceAutoCompactionByKeyspacePostParams creates a new StorageServiceAutoCompactionByKeyspacePostParams object
// with the default values initialized.
func NewStorageServiceAutoCompactionByKeyspacePostParams() *StorageServiceAutoCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceAutoCompactionByKeyspacePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceAutoCompactionByKeyspacePostParamsWithTimeout creates a new StorageServiceAutoCompactionByKeyspacePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceAutoCompactionByKeyspacePostParamsWithTimeout(timeout time.Duration) *StorageServiceAutoCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceAutoCompactionByKeyspacePostParams{

		timeout: timeout,
	}
}

// NewStorageServiceAutoCompactionByKeyspacePostParamsWithContext creates a new StorageServiceAutoCompactionByKeyspacePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceAutoCompactionByKeyspacePostParamsWithContext(ctx context.Context) *StorageServiceAutoCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceAutoCompactionByKeyspacePostParams{

		Context: ctx,
	}
}

// NewStorageServiceAutoCompactionByKeyspacePostParamsWithHTTPClient creates a new StorageServiceAutoCompactionByKeyspacePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceAutoCompactionByKeyspacePostParamsWithHTTPClient(client *http.Client) *StorageServiceAutoCompactionByKeyspacePostParams {
	var ()
	return &StorageServiceAutoCompactionByKeyspacePostParams{
		HTTPClient: client,
	}
}

/*StorageServiceAutoCompactionByKeyspacePostParams contains all the parameters to send to the API endpoint
for the storage service auto compaction by keyspace post operation typically these are written to a http.Request
*/
type StorageServiceAutoCompactionByKeyspacePostParams struct {

	/*Cf
	  Comma seperated column family names

	*/
	Cf *string
	/*Keyspace
	  The keyspace

	*/
	Keyspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) WithTimeout(timeout time.Duration) *StorageServiceAutoCompactionByKeyspacePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) WithContext(ctx context.Context) *StorageServiceAutoCompactionByKeyspacePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) WithHTTPClient(client *http.Client) *StorageServiceAutoCompactionByKeyspacePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCf adds the cf to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) WithCf(cf *string) *StorageServiceAutoCompactionByKeyspacePostParams {
	o.SetCf(cf)
	return o
}

// SetCf adds the cf to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) SetCf(cf *string) {
	o.Cf = cf
}

// WithKeyspace adds the keyspace to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) WithKeyspace(keyspace string) *StorageServiceAutoCompactionByKeyspacePostParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service auto compaction by keyspace post params
func (o *StorageServiceAutoCompactionByKeyspacePostParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceAutoCompactionByKeyspacePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cf != nil {

		// query param cf
		var qrCf string
		if o.Cf != nil {
			qrCf = *o.Cf
		}
		qCf := qrCf
		if qCf != "" {
			if err := r.SetQueryParam("cf", qCf); err != nil {
				return err
			}
		}

	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}