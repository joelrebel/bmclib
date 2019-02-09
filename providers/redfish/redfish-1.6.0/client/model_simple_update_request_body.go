/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This action is used to update software components.
type SimpleUpdateRequestBody struct {
	// The URI of the software image to be installed.
	ImageURI string `json:"ImageURI"`
	// The array of URIs indicating where the update image is to be applied.
	Targets []string `json:"Targets,omitempty"`
	TransferProtocol TransferProtocolType `json:"TransferProtocol,omitempty"`
}