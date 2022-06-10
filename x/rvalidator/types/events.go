package types

// sudo module event types
const (
	EventTypeAddRValidator    = "add_rvalidator"
	EventTypeRmRValidator     = "rm_rvalidator"
	EventTypeUpdateRValidator = "update_rvalidator"

	AttributeKeyDenom        = "denom"
	AttributeKeyAddresses    = "addresses"
	AttributeKeyAddress      = "address"
	AttributeKeyNewAddress   = "new_address"
	AttributeKeyOldAddress   = "old_address"
	AttributeKeyCycleVersion = "cycle_version"
	AttributeKeyCycleNumber  = "cycle_number"
)