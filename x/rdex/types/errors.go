package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rdex module sentinel errors
var (
	ErrSample                           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSwapPoolAlreadyExist             = sdkerrors.Register(ModuleName, 1101, "swap pool already exist error")
	ErrInvalidAddress                   = sdkerrors.Register(ModuleName, 1102, "invalid address error")
	ErrInsufficientTokenBalance         = sdkerrors.Register(ModuleName, 1103, "insufficient token balance error")
	ErrInsufficientFisBalance           = sdkerrors.Register(ModuleName, 1104, "insufficient fis balance error")
	ErrSwapPoolNotExit                  = sdkerrors.Register(ModuleName, 1105, "swap pool not exist")
	ErrAddLpUnitZero                    = sdkerrors.Register(ModuleName, 1106, "add lp unit zero error")
	ErrSwapAmountTooFew                 = sdkerrors.Register(ModuleName, 1107, "swap amount too few")
	ErrLessThanMinOutAmount             = sdkerrors.Register(ModuleName, 1108, "less than min out amount")
	ErrInsufficientLpTokenBalance       = sdkerrors.Register(ModuleName, 1109, "insufficient lp balance error")
	ErrUnitAmount                       = sdkerrors.Register(ModuleName, 1110, "unit amount error")
	ErrPoolTokenBalanceInsufficient     = sdkerrors.Register(ModuleName, 1111, "pool token balance insufficient")
	ErrPoolBaseTokenBalanceInsufficient = sdkerrors.Register(ModuleName, 1112, "pool base token balance insufficient")
	ErrPoolOneSideZero                  = sdkerrors.Register(ModuleName, 1113, "pool one side zero")
	ErrProviderNotExist                 = sdkerrors.Register(ModuleName, 1114, "provider not exist")
)
