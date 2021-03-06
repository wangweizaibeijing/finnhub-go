/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
// Quote struct for Quote
type Quote struct {
	// Open price of the day
	O float32 `json:"o,omitempty"`
	// High price of the day
	H float32 `json:"h,omitempty"`
	// Low price of the day
	L float32 `json:"l,omitempty"`
	// Current price
	C float32 `json:"c,omitempty"`
	// Previous close price
	Pc float32 `json:"pc,omitempty"`
}
