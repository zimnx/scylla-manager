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

// NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams creates a new ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams() *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParamsWithContext creates a new ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	var ()
	return &ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams{
		HTTPClient: client,
	}
}

/*ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics cas propose estimated histogram by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) WithName(name string) *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics cas propose estimated histogram by name get params
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsCasProposeEstimatedHistogramByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}