/*
 * Gate API v4
 *
 * APIv4 futures provides all sorts of futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 1.2.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesTrade struct {
	// Trade ID
	Id int64 `json:"id,omitempty"`
	// Trading time
	CreateTime float32 `json:"create_time,omitempty"`
	// Futures contract
	Contract string `json:"contract,omitempty"`
	// Trading size
	Size int64 `json:"size,omitempty"`
	// Trading price
	Price string `json:"price,omitempty"`
}
