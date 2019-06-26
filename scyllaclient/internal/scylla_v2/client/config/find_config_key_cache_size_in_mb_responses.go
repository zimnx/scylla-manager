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

// FindConfigKeyCacheSizeInMbReader is a Reader for the FindConfigKeyCacheSizeInMb structure.
type FindConfigKeyCacheSizeInMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigKeyCacheSizeInMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindConfigKeyCacheSizeInMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindConfigKeyCacheSizeInMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigKeyCacheSizeInMbOK creates a FindConfigKeyCacheSizeInMbOK with default headers values
func NewFindConfigKeyCacheSizeInMbOK() *FindConfigKeyCacheSizeInMbOK {
	return &FindConfigKeyCacheSizeInMbOK{}
}

/*FindConfigKeyCacheSizeInMbOK handles this case with default header values.

Config value
*/
type FindConfigKeyCacheSizeInMbOK struct {
	Payload int64
}

func (o *FindConfigKeyCacheSizeInMbOK) Error() string {
	return fmt.Sprintf("[GET /config/key_cache_size_in_mb][%d] findConfigKeyCacheSizeInMbOK  %+v", 200, o.Payload)
}

func (o *FindConfigKeyCacheSizeInMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigKeyCacheSizeInMbDefault creates a FindConfigKeyCacheSizeInMbDefault with default headers values
func NewFindConfigKeyCacheSizeInMbDefault(code int) *FindConfigKeyCacheSizeInMbDefault {
	return &FindConfigKeyCacheSizeInMbDefault{
		_statusCode: code,
	}
}

/*FindConfigKeyCacheSizeInMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigKeyCacheSizeInMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config key cache size in mb default response
func (o *FindConfigKeyCacheSizeInMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigKeyCacheSizeInMbDefault) Error() string {
	return fmt.Sprintf("[GET /config/key_cache_size_in_mb][%d] find_config_key_cache_size_in_mb default  %+v", o._statusCode, o.Payload)
}

func (o *FindConfigKeyCacheSizeInMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
