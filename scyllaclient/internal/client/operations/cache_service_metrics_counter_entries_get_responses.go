// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CacheServiceMetricsCounterEntriesGetReader is a Reader for the CacheServiceMetricsCounterEntriesGet structure.
type CacheServiceMetricsCounterEntriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsCounterEntriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCacheServiceMetricsCounterEntriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCacheServiceMetricsCounterEntriesGetOK creates a CacheServiceMetricsCounterEntriesGetOK with default headers values
func NewCacheServiceMetricsCounterEntriesGetOK() *CacheServiceMetricsCounterEntriesGetOK {
	return &CacheServiceMetricsCounterEntriesGetOK{}
}

/*CacheServiceMetricsCounterEntriesGetOK handles this case with default header values.

CacheServiceMetricsCounterEntriesGetOK cache service metrics counter entries get o k
*/
type CacheServiceMetricsCounterEntriesGetOK struct {
	Payload int32
}

func (o *CacheServiceMetricsCounterEntriesGetOK) Error() string {
	return fmt.Sprintf("[GET /cache_service/metrics/counter/entries][%d] cacheServiceMetricsCounterEntriesGetOK  %+v", 200, o.Payload)
}

func (o *CacheServiceMetricsCounterEntriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}