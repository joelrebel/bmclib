/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// This resource contains information about a specific Task scheduled by or being executed by a Redfish service's Task Service.
type Task struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The date-time stamp that the task was last completed.
	EndTime time.Time `json:"EndTime,omitempty"`
	// Indicates that the contents of the Payload should be hidden from view after the Task has been created.  When set to True, the Payload object will not be returned on GET.
	HidePayload bool `json:"HidePayload,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// This is an array of messages associated with the task.
	Messages []IdRef `json:"Messages,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	// Oem extension object.
	Oem map[string]map[string]interface{} `json:"Oem,omitempty"`
	Payload Payload `json:"Payload,omitempty"`
	// The date-time stamp that the task was last started.
	StartTime time.Time `json:"StartTime,omitempty"`
	// The URI of the Task Monitor for this task.
	TaskMonitor string `json:"TaskMonitor,omitempty"`
	TaskState TaskState `json:"TaskState,omitempty"`
	TaskStatus Health `json:"TaskStatus,omitempty"`
}