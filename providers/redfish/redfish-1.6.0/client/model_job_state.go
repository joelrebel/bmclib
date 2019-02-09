/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type JobState string

// List of JobState
const (
	NEW JobState = "New"
	STARTING JobState = "Starting"
	RUNNING JobState = "Running"
	SUSPENDED JobState = "Suspended"
	INTERRUPTED JobState = "Interrupted"
	PENDING JobState = "Pending"
	STOPPING JobState = "Stopping"
	COMPLETED JobState = "Completed"
	CANCELLED JobState = "Cancelled"
	EXCEPTION JobState = "Exception"
	SERVICE JobState = "Service"
	USER_INTERVENTION JobState = "UserIntervention"
	CONTINUE JobState = "Continue"
)