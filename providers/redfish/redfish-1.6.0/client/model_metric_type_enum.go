/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MetricTypeEnum : Specifies the type of metric for which the trigger is configured.
type MetricTypeEnum string

// List of MetricTypeEnum
const (
	NUMERIC MetricTypeEnum = "Numeric"
	DISCRETE MetricTypeEnum = "Discrete"
)