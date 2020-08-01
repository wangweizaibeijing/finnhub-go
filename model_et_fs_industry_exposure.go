/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
// EtFsIndustryExposure struct for EtFsIndustryExposure
type EtFsIndustryExposure struct {
	// ETF symbol.
	Symbol string `json:"symbol,omitempty"`
	// Array of industries and exposure levels.
	SectorExposure []map[string]interface{} `json:"sectorExposure,omitempty"`
}
