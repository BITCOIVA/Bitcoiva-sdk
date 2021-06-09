// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/BITCOIVA/Bitcoiva-sdk/x/slashing/types
package slashing

import (
	"github.com/BITCOIVA/Bitcoiva-sdk/x/slashing/types"
)

const (
	DefaultCodespace            = types.DefaultCodespace
	CodeInvalidValidator        = types.CodeInvalidValidator
	CodeValidatorJailed         = types.CodeValidatorJailed
	CodeValidatorNotJailed      = types.CodeValidatorNotJailed
	CodeMissingSelfDelegation   = types.CodeMissingSelfDelegation
	CodeSelfDelegationTooLow    = types.CodeSelfDelegationTooLow
	CodeMissingSigningInfo      = types.CodeMissingSigningInfo
	ModuleName                  = types.ModuleName
	StoreKey                    = types.StoreKey
	RouterKey                   = types.RouterKey
	QuerierRoute                = types.QuerierRoute
	QueryParameters             = types.QueryParameters
	QuerySigningInfo            = types.QuerySigningInfo
	QuerySigningInfos           = types.QuerySigningInfos
	DefaultParamspace           = types.DefaultParamspace
	DefaultMaxEvidenceAge       = types.DefaultMaxEvidenceAge
	DefaultSignedBlocksWindow   = types.DefaultSignedBlocksWindow
	DefaultDowntimeJailDuration = types.DefaultDowntimeJailDuration
)

var (
	// functions aliases
	RegisterCodec                            = types.RegisterCodec
	ErrNoValidatorForAddress                 = types.ErrNoValidatorForAddress
	ErrBadValidatorAddr                      = types.ErrBadValidatorAddr
	ErrValidatorJailed                       = types.ErrValidatorJailed
	ErrValidatorNotJailed                    = types.ErrValidatorNotJailed
	ErrMissingSelfDelegation                 = types.ErrMissingSelfDelegation
	ErrSelfDelegationTooLowToUnjail          = types.ErrSelfDelegationTooLowToUnjail
	ErrNoSigningInfoFound                    = types.ErrNoSigningInfoFound
	NewGenesisState                          = types.NewGenesisState
	DefaultGenesisState                      = types.DefaultGenesisState
	ValidateGenesis                          = types.ValidateGenesis
	GetValidatorSigningInfoKey               = types.GetValidatorSigningInfoKey
	GetValidatorSigningInfoAddress           = types.GetValidatorSigningInfoAddress
	GetValidatorMissedBlockBitArrayPrefixKey = types.GetValidatorMissedBlockBitArrayPrefixKey
	GetValidatorMissedBlockBitArrayKey       = types.GetValidatorMissedBlockBitArrayKey
	GetAddrPubkeyRelationKey                 = types.GetAddrPubkeyRelationKey
	NewMsgUnjail                             = types.NewMsgUnjail
	ParamKeyTable                            = types.ParamKeyTable
	NewParams                                = types.NewParams
	DefaultParams                            = types.DefaultParams
	NewQuerySigningInfoParams                = types.NewQuerySigningInfoParams
	NewQuerySigningInfosParams               = types.NewQuerySigningInfosParams
	NewValidatorSigningInfo                  = types.NewValidatorSigningInfo

	// variable aliases
	ModuleCdc                       = types.ModuleCdc
	ValidatorSigningInfoKey         = types.ValidatorSigningInfoKey
	ValidatorMissedBlockBitArrayKey = types.ValidatorMissedBlockBitArrayKey
	AddrPubkeyRelationKey           = types.AddrPubkeyRelationKey
	DoubleSignJailEndTime           = types.DoubleSignJailEndTime
	DefaultMinSignedPerWindow       = types.DefaultMinSignedPerWindow
	DefaultSlashFractionDoubleSign  = types.DefaultSlashFractionDoubleSign
	DefaultSlashFractionDowntime    = types.DefaultSlashFractionDowntime
	KeyMaxEvidenceAge               = types.KeyMaxEvidenceAge
	KeySignedBlocksWindow           = types.KeySignedBlocksWindow
	KeyMinSignedPerWindow           = types.KeyMinSignedPerWindow
	KeyDowntimeJailDuration         = types.KeyDowntimeJailDuration
	KeySlashFractionDoubleSign      = types.KeySlashFractionDoubleSign
	KeySlashFractionDowntime        = types.KeySlashFractionDowntime

	EventTypeSlash                 = types.EventTypeSlash
	EventTypeLiveness              = types.EventTypeLiveness
	AttributeKeyAddress            = types.AttributeKeyAddress
	AttributeKeyHeight             = types.AttributeKeyHeight
	AttributeKeyPower              = types.AttributeKeyPower
	AttributeKeyReason             = types.AttributeKeyReason
	AttributeKeyJailed             = types.AttributeKeyJailed
	AttributeKeyMissedBlocks       = types.AttributeKeyMissedBlocks
	AttributeValueDoubleSign       = types.AttributeValueDoubleSign
	AttributeValueMissingSignature = types.AttributeValueMissingSignature
	AttributeValueCategory         = types.AttributeValueCategory
)

type (
	CodeType                = types.CodeType
	GenesisState            = types.GenesisState
	MissedBlock             = types.MissedBlock
	MsgUnjail               = types.MsgUnjail
	Params                  = types.Params
	QuerySigningInfoParams  = types.QuerySigningInfoParams
	QuerySigningInfosParams = types.QuerySigningInfosParams
	ValidatorSigningInfo    = types.ValidatorSigningInfo
)
