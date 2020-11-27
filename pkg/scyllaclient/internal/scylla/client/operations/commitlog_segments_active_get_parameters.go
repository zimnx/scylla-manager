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

// NewCommitlogSegmentsActiveGetParams creates a new CommitlogSegmentsActiveGetParams object
// with the default values initialized.
func NewCommitlogSegmentsActiveGetParams() *CommitlogSegmentsActiveGetParams {

	return &CommitlogSegmentsActiveGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCommitlogSegmentsActiveGetParamsWithTimeout creates a new CommitlogSegmentsActiveGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCommitlogSegmentsActiveGetParamsWithTimeout(timeout time.Duration) *CommitlogSegmentsActiveGetParams {

	return &CommitlogSegmentsActiveGetParams{

		timeout: timeout,
	}
}

// NewCommitlogSegmentsActiveGetParamsWithContext creates a new CommitlogSegmentsActiveGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCommitlogSegmentsActiveGetParamsWithContext(ctx context.Context) *CommitlogSegmentsActiveGetParams {

	return &CommitlogSegmentsActiveGetParams{

		Context: ctx,
	}
}

// NewCommitlogSegmentsActiveGetParamsWithHTTPClient creates a new CommitlogSegmentsActiveGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCommitlogSegmentsActiveGetParamsWithHTTPClient(client *http.Client) *CommitlogSegmentsActiveGetParams {

	return &CommitlogSegmentsActiveGetParams{
		HTTPClient: client,
	}
}

/*CommitlogSegmentsActiveGetParams contains all the parameters to send to the API endpoint
for the commitlog segments active get operation typically these are written to a http.Request
*/
type CommitlogSegmentsActiveGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the commitlog segments active get params
func (o *CommitlogSegmentsActiveGetParams) WithTimeout(timeout time.Duration) *CommitlogSegmentsActiveGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commitlog segments active get params
func (o *CommitlogSegmentsActiveGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commitlog segments active get params
func (o *CommitlogSegmentsActiveGetParams) WithContext(ctx context.Context) *CommitlogSegmentsActiveGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commitlog segments active get params
func (o *CommitlogSegmentsActiveGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commitlog segments active get params
func (o *CommitlogSegmentsActiveGetParams) WithHTTPClient(client *http.Client) *CommitlogSegmentsActiveGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commitlog segments active get params
func (o *CommitlogSegmentsActiveGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CommitlogSegmentsActiveGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
