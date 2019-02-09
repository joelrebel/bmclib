/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type EncryptionTypes string

// List of EncryptionTypes
const (
	NATIVE_DRIVE_ENCRYPTION EncryptionTypes = "NativeDriveEncryption"
	CONTROLLER_ASSISTED EncryptionTypes = "ControllerAssisted"
	SOFTWARE_ASSISTED EncryptionTypes = "SoftwareAssisted"
)