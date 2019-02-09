/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Contains a list of wildcards and their substitution values.
type Wildcard2 struct {
	// An array of Key values to substitute for the wildcard.
	Keys []string `json:"Keys,omitempty"`
	// The name of Wildcard.
	Name string `json:"Name,omitempty"`
}