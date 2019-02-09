/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ReportUpdatesEnum : How to what to do with subsequent metric reports when a metric report already exists.
type ReportUpdatesEnum string

// List of ReportUpdatesEnum
const (
	OVERWRITE ReportUpdatesEnum = "Overwrite"
	APPEND_WRAPS_WHEN_FULL ReportUpdatesEnum = "AppendWrapsWhenFull"
	APPEND_STOPS_WHEN_FULL ReportUpdatesEnum = "AppendStopsWhenFull"
	NEW_REPORT ReportUpdatesEnum = "NewReport"
)