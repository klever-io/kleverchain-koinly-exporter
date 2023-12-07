package main

type TXContract_ContractType int32

const (
	TXContract_TransferContractType                TXContract_ContractType = 0
	TXContract_CreateAssetContractType             TXContract_ContractType = 1
	TXContract_CreateValidatorContractType         TXContract_ContractType = 2
	TXContract_ValidatorConfigContractType         TXContract_ContractType = 3
	TXContract_FreezeContractType                  TXContract_ContractType = 4
	TXContract_UnfreezeContractType                TXContract_ContractType = 5
	TXContract_DelegateContractType                TXContract_ContractType = 6
	TXContract_UndelegateContractType              TXContract_ContractType = 7
	TXContract_WithdrawContractType                TXContract_ContractType = 8
	TXContract_ClaimContractType                   TXContract_ContractType = 9
	TXContract_UnjailContractType                  TXContract_ContractType = 10
	TXContract_AssetTriggerContractType            TXContract_ContractType = 11
	TXContract_SetAccountNameContractType          TXContract_ContractType = 12
	TXContract_ProposalContractType                TXContract_ContractType = 13
	TXContract_VoteContractType                    TXContract_ContractType = 14
	TXContract_ConfigITOContractType               TXContract_ContractType = 15
	TXContract_SetITOPricesContractType            TXContract_ContractType = 16
	TXContract_BuyContractType                     TXContract_ContractType = 17
	TXContract_SellContractType                    TXContract_ContractType = 18
	TXContract_CancelMarketOrderContractType       TXContract_ContractType = 19
	TXContract_CreateMarketplaceContractType       TXContract_ContractType = 20
	TXContract_ConfigMarketplaceContractType       TXContract_ContractType = 21
	TXContract_UpdateAccountPermissionContractType TXContract_ContractType = 22
	TXContract_DepositContractType                 TXContract_ContractType = 23
	TXContract_ITOTriggerContractType              TXContract_ContractType = 24
)

// Enum value maps for TXContract_ContractType.
var (
	TXContract_ContractType_name = map[int32]string{
		0:  "TransferContractType",
		1:  "CreateAssetContractType",
		2:  "CreateValidatorContractType",
		3:  "ValidatorConfigContractType",
		4:  "FreezeContractType",
		5:  "UnfreezeContractType",
		6:  "DelegateContractType",
		7:  "UndelegateContractType",
		8:  "WithdrawContractType",
		9:  "ClaimContractType",
		10: "UnjailContractType",
		11: "AssetTriggerContractType",
		12: "SetAccountNameContractType",
		13: "ProposalContractType",
		14: "VoteContractType",
		15: "ConfigITOContractType",
		16: "SetITOPricesContractType",
		17: "BuyContractType",
		18: "SellContractType",
		19: "CancelMarketOrderContractType",
		20: "CreateMarketplaceContractType",
		21: "ConfigMarketplaceContractType",
		22: "UpdateAccountPermissionContractType",
		23: "DepositContractType",
		24: "ITOTriggerContractType",
	}
	TXContract_ContractType_value = map[string]int32{
		"TransferContractType":                0,
		"CreateAssetContractType":             1,
		"CreateValidatorContractType":         2,
		"ValidatorConfigContractType":         3,
		"FreezeContractType":                  4,
		"UnfreezeContractType":                5,
		"DelegateContractType":                6,
		"UndelegateContractType":              7,
		"WithdrawContractType":                8,
		"ClaimContractType":                   9,
		"UnjailContractType":                  10,
		"AssetTriggerContractType":            11,
		"SetAccountNameContractType":          12,
		"ProposalContractType":                13,
		"VoteContractType":                    14,
		"ConfigITOContractType":               15,
		"SetITOPricesContractType":            16,
		"BuyContractType":                     17,
		"SellContractType":                    18,
		"CancelMarketOrderContractType":       19,
		"CreateMarketplaceContractType":       20,
		"ConfigMarketplaceContractType":       21,
		"UpdateAccountPermissionContractType": 22,
		"DepositContractType":                 23,
		"ITOTriggerContractType":              24,
	}
)

type Transaction_TXResult int32

const (
	Transaction_SUCCESS Transaction_TXResult = 0
	Transaction_FAILED  Transaction_TXResult = 1
)

// Enum value maps for Transaction_TXResult.
var (
	Transaction_TXResult_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	Transaction_TXResult_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

type Transaction_TXResultCode int32

const (
	Transaction_Ok                       Transaction_TXResultCode = 0
	Transaction_OutOfFunds               Transaction_TXResultCode = 1
	Transaction_AccountError             Transaction_TXResultCode = 2
	Transaction_AssetError               Transaction_TXResultCode = 3
	Transaction_ContractInvalid          Transaction_TXResultCode = 4
	Transaction_ContractNotFound         Transaction_TXResultCode = 5
	Transaction_FeeInvalid               Transaction_TXResultCode = 6
	Transaction_ParameterInvalid         Transaction_TXResultCode = 7
	Transaction_APRInvalid               Transaction_TXResultCode = 8
	Transaction_AssetIDInvalid           Transaction_TXResultCode = 9
	Transaction_AssetTypeInvalid         Transaction_TXResultCode = 10
	Transaction_AssetCantBeMinted        Transaction_TXResultCode = 11
	Transaction_AssetCantBeBurned        Transaction_TXResultCode = 12
	Transaction_AssetCantBePaused        Transaction_TXResultCode = 13
	Transaction_AssetCantBeDelegated     Transaction_TXResultCode = 14
	Transaction_AssetOwnerCantBeChanged  Transaction_TXResultCode = 15
	Transaction_AccountNotOwner          Transaction_TXResultCode = 16
	Transaction_CommissionTooHigh        Transaction_TXResultCode = 17
	Transaction_DelegationAmountInvalid  Transaction_TXResultCode = 18
	Transaction_ProposalNotActive        Transaction_TXResultCode = 19
	Transaction_ValueInvalid             Transaction_TXResultCode = 20
	Transaction_AmountInvalid            Transaction_TXResultCode = 21
	Transaction_BucketIDInvalid          Transaction_TXResultCode = 22
	Transaction_KeyConflict              Transaction_TXResultCode = 23
	Transaction_MaxDelegationAmount      Transaction_TXResultCode = 24
	Transaction_InvalidPeerKey           Transaction_TXResultCode = 25
	Transaction_MinKFIStakedUnreached    Transaction_TXResultCode = 26
	Transaction_MaxSupplyExeeced         Transaction_TXResultCode = 27
	Transaction_SaveAccountError         Transaction_TXResultCode = 28
	Transaction_LoadAccountError         Transaction_TXResultCode = 29
	Transaction_SameAccountError         Transaction_TXResultCode = 30
	Transaction_AssetPaused              Transaction_TXResultCode = 31
	Transaction_DeletegateError          Transaction_TXResultCode = 32
	Transaction_WithdrawNotAvailable     Transaction_TXResultCode = 33
	Transaction_ErrOverflow              Transaction_TXResultCode = 34
	Transaction_SetStakingErr            Transaction_TXResultCode = 35
	Transaction_SetMarketOrderErr        Transaction_TXResultCode = 36
	Transaction_BalanceError             Transaction_TXResultCode = 37
	Transaction_KAPPError                Transaction_TXResultCode = 38
	Transaction_UnfreezeError            Transaction_TXResultCode = 39
	Transaction_UndeletegateError        Transaction_TXResultCode = 40
	Transaction_WithdrawError            Transaction_TXResultCode = 41
	Transaction_ClaimError               Transaction_TXResultCode = 42
	Transaction_BucketsExceded           Transaction_TXResultCode = 43
	Transaction_AssetCantBeWiped         Transaction_TXResultCode = 44
	Transaction_AssetCantAddRoles        Transaction_TXResultCode = 45
	Transaction_FreezeError              Transaction_TXResultCode = 46
	Transaction_ITONotActive             Transaction_TXResultCode = 47
	Transaction_NFTMintStopped           Transaction_TXResultCode = 48
	Transaction_RoyaltiesChangeStopped   Transaction_TXResultCode = 49
	Transaction_ITOKAPPError             Transaction_TXResultCode = 50
	Transaction_ITOWhiteListError        Transaction_TXResultCode = 51
	Transaction_NFTMetadataChangeStopped Transaction_TXResultCode = 52
	Transaction_AlreadyExists            Transaction_TXResultCode = 53
	Transaction_Fail                     Transaction_TXResultCode = 99
)

// Enum value maps for Transaction_TXResultCode.
var (
	Transaction_TXResultCode_name = map[int32]string{
		0:  "Ok",
		1:  "OutOfFunds",
		2:  "AccountError",
		3:  "AssetError",
		4:  "ContractInvalid",
		5:  "ContractNotFound",
		6:  "FeeInvalid",
		7:  "ParameterInvalid",
		8:  "APRInvalid",
		9:  "AssetIDInvalid",
		10: "AssetTypeInvalid",
		11: "AssetCantBeMinted",
		12: "AssetCantBeBurned",
		13: "AssetCantBePaused",
		14: "AssetCantBeDelegated",
		15: "AssetOwnerCantBeChanged",
		16: "AccountNotOwner",
		17: "CommissionTooHigh",
		18: "DelegationAmountInvalid",
		19: "ProposalNotActive",
		20: "ValueInvalid",
		21: "AmountInvalid",
		22: "BucketIDInvalid",
		23: "KeyConflict",
		24: "MaxDelegationAmount",
		25: "InvalidPeerKey",
		26: "MinKFIStakedUnreached",
		27: "MaxSupplyExeeced",
		28: "SaveAccountError",
		29: "LoadAccountError",
		30: "SameAccountError",
		31: "AssetPaused",
		32: "DeletegateError",
		33: "WithdrawNotAvailable",
		34: "ErrOverflow",
		35: "SetStakingErr",
		36: "SetMarketOrderErr",
		37: "BalanceError",
		38: "KAPPError",
		39: "UnfreezeError",
		40: "UndeletegateError",
		41: "WithdrawError",
		42: "ClaimError",
		43: "BucketsExceded",
		44: "AssetCantBeWiped",
		45: "AssetCantAddRoles",
		46: "FreezeError",
		47: "ITONotActive",
		48: "NFTMintStopped",
		49: "RoyaltiesChangeStopped",
		50: "ITOKAPPError",
		51: "ITOWhiteListError",
		52: "NFTMetadataChangeStopped",
		53: "AlreadyExists",
		99: "Fail",
	}
	Transaction_TXResultCode_value = map[string]int32{
		"Ok":                       0,
		"OutOfFunds":               1,
		"AccountError":             2,
		"AssetError":               3,
		"ContractInvalid":          4,
		"ContractNotFound":         5,
		"FeeInvalid":               6,
		"ParameterInvalid":         7,
		"APRInvalid":               8,
		"AssetIDInvalid":           9,
		"AssetTypeInvalid":         10,
		"AssetCantBeMinted":        11,
		"AssetCantBeBurned":        12,
		"AssetCantBePaused":        13,
		"AssetCantBeDelegated":     14,
		"AssetOwnerCantBeChanged":  15,
		"AccountNotOwner":          16,
		"CommissionTooHigh":        17,
		"DelegationAmountInvalid":  18,
		"ProposalNotActive":        19,
		"ValueInvalid":             20,
		"AmountInvalid":            21,
		"BucketIDInvalid":          22,
		"KeyConflict":              23,
		"MaxDelegationAmount":      24,
		"InvalidPeerKey":           25,
		"MinKFIStakedUnreached":    26,
		"MaxSupplyExeeced":         27,
		"SaveAccountError":         28,
		"LoadAccountError":         29,
		"SameAccountError":         30,
		"AssetPaused":              31,
		"DeletegateError":          32,
		"WithdrawNotAvailable":     33,
		"ErrOverflow":              34,
		"SetStakingErr":            35,
		"SetMarketOrderErr":        36,
		"BalanceError":             37,
		"KAPPError":                38,
		"UnfreezeError":            39,
		"UndeletegateError":        40,
		"WithdrawError":            41,
		"ClaimError":               42,
		"BucketsExceded":           43,
		"AssetCantBeWiped":         44,
		"AssetCantAddRoles":        45,
		"FreezeError":              46,
		"ITONotActive":             47,
		"NFTMintStopped":           48,
		"RoyaltiesChangeStopped":   49,
		"ITOKAPPError":             50,
		"ITOWhiteListError":        51,
		"NFTMetadataChangeStopped": 52,
		"AlreadyExists":            53,
		"Fail":                     99,
	}
)
