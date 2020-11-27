// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// CacheServiceMetricsRowHitsGetReader is a Reader for the CacheServiceMetricsRowHitsGet structure.
type CacheServiceMetricsRowHitsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsRowHitsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsRowHitsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsRowHitsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsRowHitsGetOK creates a CacheServiceMetricsRowHitsGetOK with default headers values
func NewCacheServiceMetricsRowHitsGetOK() *CacheServiceMetricsRowHitsGetOK {
	return &CacheServiceMetricsRowHitsGetOK{}
}

/*CacheServiceMetricsRowHitsGetOK handles this case with default header values.

CacheServiceMetricsRowHitsGetOK cache service metrics row hits get o k
*/
type CacheServiceMetricsRowHitsGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsRowHitsGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsRowHitsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsRowHitsGetDefault creates a CacheServiceMetricsRowHitsGetDefault with default headers values
func NewCacheServiceMetricsRowHitsGetDefault(code int) *CacheServiceMetricsRowHitsGetDefault {
	return &CacheServiceMetricsRowHitsGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsRowHitsGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsRowHitsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics row hits get default response
func (o *CacheServiceMetricsRowHitsGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsRowHitsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsRowHitsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsRowHitsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
