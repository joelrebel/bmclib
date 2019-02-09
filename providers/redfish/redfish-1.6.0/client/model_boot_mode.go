/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type BootMode string

// List of BootMode
const (
	DISABLED BootMode = "Disabled"
	PXE BootMode = "PXE"
	I_SCSI BootMode = "iSCSI"
	FIBRE_CHANNEL BootMode = "FibreChannel"
	FIBRE_CHANNEL_OVER_ETHERNET BootMode = "FibreChannelOverEthernet"
)