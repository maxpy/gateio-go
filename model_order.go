/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Spot order details
type Order struct {
	// Order ID
	Id string `json:"id,omitempty"`
	// User defined information. If not empty, must follow the rules below:  1. prefixed with `t-` 2. no longer than 16 bytes without `t-` prefix 3. can only include 0-9, A-Z, a-z, underscore(_), hyphen(-) or dot(.) 
	Text string `json:"text,omitempty"`
	// Order creation time
	CreateTime string `json:"create_time,omitempty"`
	// Order last modification time
	UpdateTime string `json:"update_time,omitempty"`
	// Order status  - `open`: to be filled - `closed`: filled - `cancelled`: cancelled
	Status string `json:"status,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair"`
	// Order type. limit - limit order
	Type string `json:"type,omitempty"`
	// Account type. spot - use spot account; margin - use margin account
	Account string `json:"account,omitempty"`
	// Order side
	Side string `json:"side"`
	// Trade amount
	Amount string `json:"amount"`
	// Order price
	Price string `json:"price"`
	// Time in force
	TimeInForce string `json:"time_in_force,omitempty"`
	// Amount left to fill
	Left string `json:"left,omitempty"`
	// Total filled in quote currency
	FillPrice string `json:"fill_price,omitempty"`
	// Fee deducted
	Fee string `json:"fee,omitempty"`
	// Fee currency unit
	FeeCurrency string `json:"fee_currency,omitempty"`
	// Point used to deduct fee
	PointFee string `json:"point_fee,omitempty"`
	// GT used to deduct fee
	GtFee string `json:"gt_fee,omitempty"`
}
