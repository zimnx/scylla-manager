// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StorageServiceSstablesByKeyspacePostReader is a Reader for the StorageServiceSstablesByKeyspacePost structure.
type StorageServiceSstablesByKeyspacePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceSstablesByKeyspacePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStorageServiceSstablesByKeyspacePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStorageServiceSstablesByKeyspacePostOK creates a StorageServiceSstablesByKeyspacePostOK with default headers values
func NewStorageServiceSstablesByKeyspacePostOK() *StorageServiceSstablesByKeyspacePostOK {
	return &StorageServiceSstablesByKeyspacePostOK{}
}

/*StorageServiceSstablesByKeyspacePostOK handles this case with default header values.

StorageServiceSstablesByKeyspacePostOK storage service sstables by keyspace post o k
*/
type StorageServiceSstablesByKeyspacePostOK struct {
}

func (o *StorageServiceSstablesByKeyspacePostOK) Error() string {
	return fmt.Sprintf("[POST /storage_service/sstables/{keyspace}][%d] storageServiceSstablesByKeyspacePostOK ", 200)
}

func (o *StorageServiceSstablesByKeyspacePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}