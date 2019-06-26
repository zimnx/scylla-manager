// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCommitlogSegmentSizeInMbReader is a Reader for the FindConfigCommitlogSegmentSizeInMb structure.
type FindConfigCommitlogSegmentSizeInMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCommitlogSegmentSizeInMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigCommitlogSegmentSizeInMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigCommitlogSegmentSizeInMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCommitlogSegmentSizeInMbOK creates a FindConfigCommitlogSegmentSizeInMbOK with default headers values
func NewFindConfigCommitlogSegmentSizeInMbOK() *FindConfigCommitlogSegmentSizeInMbOK {
	return &FindConfigCommitlogSegmentSizeInMbOK{}
}

/*FindConfigCommitlogSegmentSizeInMbOK handles this case with default header values.

Config value
*/
type FindConfigCommitlogSegmentSizeInMbOK struct {
	Payload int64
}

func (o *FindConfigCommitlogSegmentSizeInMbOK) Error() string {
	return fmt.Sprintf("[GET /config/commitlog_segment_size_in_mb][%d] findConfigCommitlogSegmentSizeInMbOK  %+v", 200, o.Payload)
}

func (o *FindConfigCommitlogSegmentSizeInMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCommitlogSegmentSizeInMbDefault creates a FindConfigCommitlogSegmentSizeInMbDefault with default headers values
func NewFindConfigCommitlogSegmentSizeInMbDefault(code int) *FindConfigCommitlogSegmentSizeInMbDefault {
	return &FindConfigCommitlogSegmentSizeInMbDefault{
		_statusCode: code,
	}
}

/*FindConfigCommitlogSegmentSizeInMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigCommitlogSegmentSizeInMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config commitlog segment size in mb default response
func (o *FindConfigCommitlogSegmentSizeInMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCommitlogSegmentSizeInMbDefault) Error() string {
	return fmt.Sprintf("[GET /config/commitlog_segment_size_in_mb][%d] find_config_commitlog_segment_size_in_mb default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigCommitlogSegmentSizeInMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
