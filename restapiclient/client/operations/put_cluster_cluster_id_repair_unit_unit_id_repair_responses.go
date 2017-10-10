// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutClusterClusterIDRepairUnitUnitIDRepairReader is a Reader for the PutClusterClusterIDRepairUnitUnitIDRepair structure.
type PutClusterClusterIDRepairUnitUnitIDRepairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDRepairUnitUnitIDRepairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutClusterClusterIDRepairUnitUnitIDRepairCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutClusterClusterIDRepairUnitUnitIDRepairCreated creates a PutClusterClusterIDRepairUnitUnitIDRepairCreated with default headers values
func NewPutClusterClusterIDRepairUnitUnitIDRepairCreated() *PutClusterClusterIDRepairUnitUnitIDRepairCreated {
	return &PutClusterClusterIDRepairUnitUnitIDRepairCreated{}
}

/*PutClusterClusterIDRepairUnitUnitIDRepairCreated handles this case with default header values.

OK
*/
type PutClusterClusterIDRepairUnitUnitIDRepairCreated struct {
	/*location of the new task
	 */
	Location string
}

func (o *PutClusterClusterIDRepairUnitUnitIDRepairCreated) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repair/unit/{unit_id}/repair][%d] putClusterClusterIdRepairUnitUnitIdRepairCreated ", 201)
}

func (o *PutClusterClusterIDRepairUnitUnitIDRepairCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}
