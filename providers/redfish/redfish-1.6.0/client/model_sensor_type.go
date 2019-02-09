/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type SensorType string

// List of SensorType
const (
	PLATFORM_SECURITY_VIOLATION_ATTEMPT SensorType = "Platform Security Violation Attempt"
	TEMPERATURE SensorType = "Temperature"
	VOLTAGE SensorType = "Voltage"
	CURRENT SensorType = "Current"
	FAN SensorType = "Fan"
	PHYSICAL_CHASSIS_SECURITY SensorType = "Physical Chassis Security"
	PROCESSOR SensorType = "Processor"
	POWER_SUPPLY__CONVERTER SensorType = "Power Supply / Converter"
	POWER_UNIT SensorType = "PowerUnit"
	COOLING_DEVICE SensorType = "CoolingDevice"
	OTHER_UNITS_BASED_SENSOR SensorType = "Other Units-based Sensor"
	MEMORY SensorType = "Memory"
	DRIVE_SLOTBAY SensorType = "Drive Slot/Bay"
	POST_MEMORY_RESIZE SensorType = "POST Memory Resize"
	SYSTEM_FIRMWARE_PROGRESS SensorType = "System Firmware Progress"
	EVENT_LOGGING_DISABLED SensorType = "Event Logging Disabled"
	SYSTEM_EVENT SensorType = "System Event"
	CRITICAL_INTERRUPT SensorType = "Critical Interrupt"
	BUTTONSWITCH SensorType = "Button/Switch"
	MODULEBOARD SensorType = "Module/Board"
	MICROCONTROLLERCOPROCESSOR SensorType = "Microcontroller/Coprocessor"
	ADD_IN_CARD SensorType = "Add-in Card"
	CHASSIS SensorType = "Chassis"
	CHIP_SET SensorType = "ChipSet"
	OTHER_FRU SensorType = "Other FRU"
	CABLEINTERCONNECT SensorType = "Cable/Interconnect"
	TERMINATOR SensorType = "Terminator"
	SYSTEM_BOOTRESTART SensorType = "SystemBoot/Restart"
	BOOT_ERROR SensorType = "Boot Error"
	BASE_OS_BOOTINSTALLATION_STATUS SensorType = "BaseOSBoot/InstallationStatus"
	OS_STOPSHUTDOWN SensorType = "OS Stop/Shutdown"
	SLOTCONNECTOR SensorType = "Slot/Connector"
	SYSTEM_ACPI_POWER_STATE SensorType = "System ACPI PowerState"
	WATCHDOG SensorType = "Watchdog"
	PLATFORM_ALERT SensorType = "Platform Alert"
	ENTITY_PRESENCE SensorType = "Entity Presence"
	MONITOR_ASICIC SensorType = "Monitor ASIC/IC"
	LAN SensorType = "LAN"
	MANAGEMENT_SUBSYSTEM_HEALTH SensorType = "Management Subsystem Health"
	BATTERY SensorType = "Battery"
	SESSION_AUDIT SensorType = "Session Audit"
	VERSION_CHANGE SensorType = "Version Change"
	FRU_STATE SensorType = "FRUState"
	OEM SensorType = "OEM"
)