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

// ColumnFamilyMetricsSpeculativeRetriesByNameGetReader is a Reader for the ColumnFamilyMetricsSpeculativeRetriesByNameGet structure.
type ColumnFamilyMetricsSpeculativeRetriesByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewColumnFamilyMetricsSpeculativeRetriesByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsSpeculativeRetriesByNameGetOK creates a ColumnFamilyMetricsSpeculativeRetriesByNameGetOK with default headers values
func NewColumnFamilyMetricsSpeculativeRetriesByNameGetOK() *ColumnFamilyMetricsSpeculativeRetriesByNameGetOK {
	return &ColumnFamilyMetricsSpeculativeRetriesByNameGetOK{}
}

/*ColumnFamilyMetricsSpeculativeRetriesByNameGetOK handles this case with default header values.

ColumnFamilyMetricsSpeculativeRetriesByNameGetOK column family metrics speculative retries by name get o k
*/
type ColumnFamilyMetricsSpeculativeRetriesByNameGetOK struct {
	Payload int32
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/speculative_retries/{name}][%d] columnFamilyMetricsSpeculativeRetriesByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsSpeculativeRetriesByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}