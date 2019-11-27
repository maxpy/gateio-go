/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FundingBookItem struct {
	// Loan rate
	Rate string `json:"rate,omitempty"`
	// Borrowable amount
	Amount string `json:"amount,omitempty"`
	// How long the loan should be repaid
	Days int32 `json:"days,omitempty"`
}