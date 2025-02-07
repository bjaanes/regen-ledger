package simulation

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	basketsims "github.com/regen-network/regen-ledger/x/ecocredit/basket/simulation"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	marketsims "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/simulation"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateClass              = "op_weight_msg_create_class"                //nolint:gosec
	OpWeightMsgCreateBatch              = "op_weight_msg_create_batch"                //nolint:gosec
	OpWeightMsgSend                     = "op_weight_msg_send"                        //nolint:gosec
	OpWeightMsgRetire                   = "op_weight_msg_retire"                      //nolint:gosec
	OpWeightMsgCancel                   = "op_weight_msg_cancel"                      //nolint:gosec
	OpWeightMsgUpdateClassAdmin         = "op_weight_msg_update_class_admin"          //nolint:gosec
	OpWeightMsgUpdateClassMetadata      = "op_weight_msg_update_class_metadata"       //nolint:gosec
	OpWeightMsgUpdateClassIssuers       = "op_weight_msg_update_class_issuers"        //nolint:gosec
	OpWeightMsgCreateProject            = "op_weight_msg_create_project"              //nolint:gosec
	OpWeightMsgUpdateProjectAdmin       = "op_weight_msg_update_project_admin"        //nolint:gosec
	OpWeightMsgUpdateProjectMetadata    = "op_weight_msg_update_project_metadata"     //nolint:gosec
	OpWeightMsgMintBatchCredits         = "op_weight_msg_mint_batch_credits"          //nolint:gosec
	OpWeightMsgSealBatch                = "op_weight_msg_seal_batch"                  //nolint:gosec
	OpWeightMsgBridge                   = "op_weight_msg_bridge"                      //nolint:gosec
	OpWeightMsgAddCreditType            = "op_weight_msg_add_credit_type"             //nolint:gosec
	OpWeightMsgAddClassCreator          = "op_weight_msg_add_class_creator"           //nolint:gosec
	OpWeightMsgRemoveClassCreator       = "op_weight_msg_remove_class_creator"        //nolint:gosec
	OpWeightMsgSetClassCreatorAllowlist = "op_weight_msg_set_class_creator_allowlist" //nolint:gosec
	OpWeightMsgUpdateClassFee           = "op_weight_msg_update_class_fee"            //nolint:gosec
)

// ecocredit operations weights
const (
	WeightCreateClass           = 10
	WeightCreateProject         = 20
	WeightCreateBatch           = 50
	WeightSend                  = 100
	WeightRetire                = 80
	WeightCancel                = 30
	WeightUpdateClass           = 30
	WeightUpdateProjectAdmin    = 30
	WeightUpdateProjectMetadata = 30
	WeightMintBatchCredits      = 33
	WeightSealBatch             = 33
	WeightBridge                = 33
)

// ecocredit message types
var (
	TypeMsgCreateClass              = sdk.MsgTypeURL(&types.MsgCreateClass{})
	TypeMsgCreateProject            = sdk.MsgTypeURL(&types.MsgCreateProject{})
	TypeMsgCreateBatch              = sdk.MsgTypeURL(&types.MsgCreateBatch{})
	TypeMsgSend                     = sdk.MsgTypeURL(&types.MsgSend{})
	TypeMsgRetire                   = sdk.MsgTypeURL(&types.MsgRetire{})
	TypeMsgCancel                   = sdk.MsgTypeURL(&types.MsgCancel{})
	TypeMsgUpdateClassAdmin         = sdk.MsgTypeURL(&types.MsgUpdateClassAdmin{})
	TypeMsgUpdateClassIssuers       = sdk.MsgTypeURL(&types.MsgUpdateClassIssuers{})
	TypeMsgUpdateClassMetadata      = sdk.MsgTypeURL(&types.MsgUpdateClassMetadata{})
	TypeMsgUpdateProjectMetadata    = sdk.MsgTypeURL(&types.MsgUpdateProjectMetadata{})
	TypeMsgUpdateProjectAdmin       = sdk.MsgTypeURL(&types.MsgUpdateProjectAdmin{})
	TypeMsgBridge                   = sdk.MsgTypeURL(&types.MsgBridge{})
	TypeMsgMintBatchCredits         = sdk.MsgTypeURL(&types.MsgMintBatchCredits{})
	TypeMsgSealBatch                = sdk.MsgTypeURL(&types.MsgSealBatch{})
	TypeMsgAddCreditType            = sdk.MsgTypeURL(&types.MsgAddCreditType{})
	TypeMsgAddClassCreator          = sdk.MsgTypeURL(&types.MsgAddClassCreator{})
	TypeMsgRemoveClassCreator       = sdk.MsgTypeURL(&types.MsgRemoveClassCreator{})
	TypeMsgSetClassCreatorAllowlist = sdk.MsgTypeURL(&types.MsgSetClassCreatorAllowlist{})
	TypeMsgUpdateClassFee           = sdk.MsgTypeURL(&types.MsgUpdateClassFee{})
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	govk ecocredit.GovKeeper,
	qryClient types.QueryServer, basketQryClient baskettypes.QueryServer,
	mktQryClient markettypes.QueryServer, authority sdk.AccAddress) simulation.WeightedOperations {

	var (
		weightMsgCreateClass              int
		weightMsgCreateBatch              int
		weightMsgSend                     int
		weightMsgRetire                   int
		weightMsgCancel                   int
		weightMsgUpdateClassAdmin         int
		weightMsgUpdateClassIssuers       int
		weightMsgUpdateClassMetadata      int
		weightMsgCreateProject            int
		weightMsgUpdateProjectMetadata    int
		weightMsgUpdateProjectAdmin       int
		weightMsgSealBatch                int
		weightMsgMintBatchCredits         int
		weightMsgBridge                   int
		weightMsgAddCreditType            int
		weightMsgAddClassCreator          int
		weightMsgRemoveClassCreator       int
		weightMsgSetClassCreatorAllowlist int
		weightMsgUpdateClassFee           int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = WeightCreateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = WeightCreateProject
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateBatch, &weightMsgCreateBatch, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBatch = WeightCreateBatch
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = WeightSend
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRetire, &weightMsgRetire, nil,
		func(_ *rand.Rand) {
			weightMsgRetire = WeightRetire
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = WeightCancel
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassAdmin, &weightMsgUpdateClassAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassAdmin = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassIssuers, &weightMsgUpdateClassIssuers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassIssuers = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassMetadata, &weightMsgUpdateClassMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassMetadata = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateProjectAdmin, &weightMsgUpdateProjectAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProjectAdmin = WeightUpdateProjectAdmin
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateProjectMetadata, &weightMsgUpdateProjectMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProjectMetadata = WeightUpdateProjectMetadata
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgMintBatchCredits, &weightMsgMintBatchCredits, nil,
		func(_ *rand.Rand) {
			weightMsgMintBatchCredits = WeightMintBatchCredits
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSealBatch, &weightMsgSealBatch, nil,
		func(_ *rand.Rand) {
			weightMsgSealBatch = WeightSealBatch
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBridge, &weightMsgBridge, nil,
		func(_ *rand.Rand) {
			weightMsgBridge = WeightBridge
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAddCreditType, &weightMsgAddCreditType, nil,
		func(_ *rand.Rand) {
			weightMsgAddCreditType = WeightBridge
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAddClassCreator, &weightMsgAddClassCreator, nil,
		func(_ *rand.Rand) {
			weightMsgAddClassCreator = WeightBridge
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveClassCreator, &weightMsgRemoveClassCreator, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveClassCreator = WeightBridge
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSetClassCreatorAllowlist, &weightMsgSetClassCreatorAllowlist, nil,
		func(_ *rand.Rand) {
			weightMsgSetClassCreatorAllowlist = WeightBridge
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassFee, &weightMsgUpdateClassFee, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassFee = WeightBridge
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateClass,
			SimulateMsgCreateClass(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProject,
			SimulateMsgCreateProject(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateBatch,
			SimulateMsgCreateBatch(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSend,
			SimulateMsgSend(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgRetire,
			SimulateMsgRetire(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancel,
			SimulateMsgCancel(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateClassAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassIssuers,
			SimulateMsgUpdateClassIssuers(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassMetadata,
			SimulateMsgUpdateClassMetadata(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateProjectAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateProjectMetadata,
			SimulateMsgUpdateProjectMetadata(ak, bk, qryClient),
		),

		simulation.NewWeightedOperation(
			weightMsgMintBatchCredits,
			SimulateMsgMintBatchCredits(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSealBatch,
			SimulateMsgSealBatch(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgBridge,
			SimulateMsgBridge(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAddCreditType,
			SimulateMsgAddCreditType(ak, bk, govk, qryClient, authority),
		),

		simulation.NewWeightedOperation(
			weightMsgAddClassCreator,
			SimulateMsgAddClassCreator(ak, bk, govk, qryClient, authority),
		),

		simulation.NewWeightedOperation(
			weightMsgRemoveClassCreator,
			SimulateMsgRemoveClassCreator(ak, bk, govk, qryClient, authority),
		),

		simulation.NewWeightedOperation(
			weightMsgSetClassCreatorAllowlist,
			SimulateMsgSetClassCreatorAllowlist(ak, bk, govk, qryClient, authority),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassFee,
			SimulateMsgUpdateClassFee(ak, bk, govk, qryClient, authority),
		),
	}

	basketOps := basketsims.WeightedOperations(appParams, cdc, ak, bk, govk, qryClient, basketQryClient, authority)
	ops = append(ops, basketOps...)
	marketplaceOps := marketsims.WeightedOperations(appParams, cdc, ak, bk, qryClient, mktQryClient, govk, authority)

	return append(ops, marketplaceOps...)
}

// SimulateMsgUpdateProjectMetadata generates a MsgUpdateProjectMetadata with random values.
func SimulateMsgUpdateProjectMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateProjectMetadata)
		if err != nil {
			return op, nil, err
		}

		ctx := sdk.WrapSDKContext(sdkCtx)
		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgUpdateProjectMetadata, class.Id)
		if project == nil {
			return op, nil, err
		}

		admin, err := sdk.AccAddressFromBech32(project.Admin)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateProjectMetadata, err.Error()), nil, err
		}

		msg := &types.MsgUpdateProjectMetadata{
			Admin:       admin.String(),
			ProjectId:   project.Id,
			NewMetadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, base.MaxMetadataLength)),
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateProjectMetadata)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgUpdateProjectAdmin generates a MsgUpdateProjectAdmin with random values.
func SimulateMsgUpdateProjectAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateProjectAdmin)
		if err != nil {
			return op, nil, err
		}

		ctx := sdk.WrapSDKContext(sdkCtx)
		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgUpdateProjectAdmin, class.Id)
		if project == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if project.Admin == newAdmin.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateProjectAdmin, "old and new admin are same"), nil, nil
		}

		msg := &types.MsgUpdateProjectAdmin{
			Admin:     project.Admin,
			NewAdmin:  newAdmin.Address.String(),
			ProjectId: project.Id,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, project.Admin, TypeMsgUpdateProjectAdmin)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		admin, _ := simtypes.RandomAcc(r, accs)
		issuers := randomIssuers(r, accs)

		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := qryClient.Params(ctx, &types.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !utils.Contains(params.AllowedClassCreators, admin.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendable, neg := bk.SpendableCoins(sdkCtx, admin.Address).SafeSub(params.CreditClassFee...)
		if neg {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		creditTypes := []string{"C", "BIO"}
		msg := &types.MsgCreateClass{
			Admin:            admin.Address.String(),
			Issuers:          issuers,
			Metadata:         simtypes.RandStringOfLength(r, 10),
			CreditTypeAbbrev: creditTypes[r.Intn(len(creditTypes))],
			Fee:              &params.CreditClassFee[0],
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      admin,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgCreateProject generates a MsgCreateProject with random values.
func SimulateMsgCreateProject(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCreateProject)
		if class == nil {
			return op, nil, err
		}

		issuers, op, err := getClassIssuers(sdkCtx, qryClient, class.Id, TypeMsgCreateProject)
		if len(issuers) == 0 {
			return op, nil, err
		}

		admin := issuers[r.Intn(len(issuers))]
		adminAddr, err := sdk.AccAddressFromBech32(admin)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, err.Error()), nil, err
		}

		adminAcc, found := simtypes.FindAccount(accs, adminAddr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "not a simulation account"), nil, nil
		}

		spendable := bk.SpendableCoins(sdkCtx, adminAddr)

		msg := &types.MsgCreateProject{
			Admin:        admin,
			ClassId:      class.Id,
			Metadata:     simtypes.RandStringOfLength(r, 100),
			Jurisdiction: "AB-CDE FG1 345",
		}
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      adminAcc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgCreateBatch generates a MsgCreateBatch with random values.
func SimulateMsgCreateBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer, _ := simtypes.RandomAcc(r, accs)

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCreateBatch)
		if class == nil {
			return op, nil, err
		}

		result, err := qryClient.ClassIssuers(ctx, &types.QueryClassIssuersRequest{ClassId: class.Id})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, err.Error()), nil, err
		}

		classIssuers := result.Issuers
		if len(classIssuers) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "no issuers"), nil, nil
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgCreateBatch, class.Id)
		if project == nil {
			return op, nil, err
		}

		if !utils.Contains(classIssuers, issuer.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "don't have permission to create credit batch"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		now := sdkCtx.BlockTime()
		tenHours := now.Add(10 * time.Hour)

		msg := &types.MsgCreateBatch{
			Issuer:    issuer.Address.String(),
			ProjectId: project.Id,
			Issuance:  generateBatchIssuance(r, accs),
			StartDate: &now,
			EndDate:   &tenHours,
			Metadata:  simtypes.RandStringOfLength(r, 10),
			Open:      r.Float32() < 0.3, // 30% chance of credit batch being dynamic batch
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      issuer,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgSend generates a MsgSend with random values.
func SimulateMsgSend(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgSend)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgSend, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgSend, class.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balres, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balres.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		retiredBalance, err := math.NewNonNegativeDecFromString(balres.Balance.RetiredAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		if tradableBalance.IsZero() || retiredBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "balance is zero"), nil, nil
		}

		recipient, _ := simtypes.RandomAcc(r, accs)
		if admin == recipient.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "sender & recipient are same"), nil, nil
		}

		addr := sdk.AccAddress(project.Admin)
		acc, found := simtypes.FindAccount(accs, addr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "account not found"), nil, nil
		}

		issuer := ak.GetAccount(sdkCtx, acc.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuer.GetAddress())

		var tradable int
		var retired int
		var retirementJurisdiction string
		if !tradableBalance.IsZero() {
			i64, err := tradableBalance.Int64()
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, nil
			}
			if i64 > 1 {
				tradable = simtypes.RandIntBetween(r, 1, int(i64))
				retired = simtypes.RandIntBetween(r, 0, tradable)
				if retired != 0 {
					retirementJurisdiction = "AQ"
				}
			} else {
				tradable = int(i64)
			}
		}

		if retired+tradable > tradable {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "insufficient credit balance"), nil, nil
		}

		msg := &types.MsgSend{
			Sender:    admin,
			Recipient: recipient.Address.String(),
			Credits: []*types.MsgSend_SendCredits{
				{
					BatchDenom:             batch.Denom,
					TradableAmount:         fmt.Sprintf("%d", tradable),
					RetiredAmount:          fmt.Sprintf("%d", retired),
					RetirementJurisdiction: retirementJurisdiction,
				},
			},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      acc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgRetire generates a MsgRetire with random values.
func SimulateMsgRetire(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgRetire)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgRetire, project.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balanceRes, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "balance is zero"), nil, nil
		}

		randSub := math.NewDecFromInt64(int64(simtypes.RandIntBetween(r, 1, 10)))
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin, TypeMsgRetire)
		if spendable == nil {
			return op, nil, err
		}

		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		if tradableBalance.Cmp(randSub) != 1 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		msg := &types.MsgRetire{
			Owner: account.Address.String(),
			Credits: []*types.Credits{
				{
					BatchDenom: batch.Denom,
					Amount:     randSub.String(),
				},
			},
			Jurisdiction: "ST-UVW XY Z12",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgCancel generates a MsgCancel with random values.
func SimulateMsgCancel(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgCancel)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgCancel, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgCancel, project.Id)
		if batch == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(project.Admin).String()
		balanceRes, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    admin,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "balance is zero"), nil, nil
		}

		msg := &types.MsgCancel{
			Owner: admin,
			Credits: []*types.Credits{
				{
					BatchDenom: batch.Denom,
					Amount:     balanceRes.Balance.TradableAmount,
				},
			},
			Reason: simtypes.RandStringOfLength(r, 5),
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin, TypeMsgCancel)
		if spendable == nil {
			return op, nil, err
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgUpdateClassAdmin generates a MsgUpdateClassAdmin with random values
func SimulateMsgUpdateClassAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassAdmin)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassAdmin)
		if spendable == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if newAdmin.Address.String() == admin.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassAdmin, "old and new account is same"), nil, nil // skip
		}

		msg := &types.MsgUpdateClassAdmin{
			Admin:    admin.String(),
			ClassId:  class.Id,
			NewAdmin: newAdmin.Address.String(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgUpdateClassMetadata generates a MsgUpdateClassMetadata with random metadata
func SimulateMsgUpdateClassMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassMetadata)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassMetadata)
		if spendable == nil {
			return op, nil, err
		}

		msg := &types.MsgUpdateClassMetadata{
			Admin:       admin.String(),
			ClassId:     class.Id,
			NewMetadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 256)),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgUpdateClassIssuers generates a MsgUpdateClassMetaData with random values
func SimulateMsgUpdateClassIssuers(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassIssuers)
		if class == nil {
			return op, nil, err
		}

		admin := sdk.AccAddress(class.Admin)
		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, admin.String(), TypeMsgUpdateClassIssuers)
		if spendable == nil {
			return op, nil, err
		}

		issuersRes, err := qryClient.ClassIssuers(sdk.WrapSDKContext(sdkCtx), &types.QueryClassIssuersRequest{ClassId: class.Id})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassIssuers, err.Error()), nil, err
		}
		classIssuers := issuersRes.Issuers

		var addIssuers []string
		var removeIssuers []string

		issuers := randomIssuers(r, accs)
		if len(issuers) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassIssuers, "empty issuers"), nil, nil
		}

		for _, i := range issuers {
			if utils.Contains(classIssuers, i) {
				removeIssuers = append(removeIssuers, i)
			} else {
				addIssuers = append(addIssuers, i)
			}
		}

		msg := &types.MsgUpdateClassIssuers{
			Admin:         admin.String(),
			ClassId:       class.Id,
			AddIssuers:    addIssuers,
			RemoveIssuers: removeIssuers,
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgMintBatchCredits generates a MsgMintBatchCredits with random values.
func SimulateMsgMintBatchCredits(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuerAcc, _ := simtypes.RandomAcc(r, accs)
		issuerAddr := issuerAcc.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgMintBatchCredits)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgMintBatchCredits, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgMintBatchCredits, project.Id)
		if batch == nil {
			return op, nil, err
		}

		if !batch.Open {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgMintBatchCredits, "batch is closed"), nil, nil
		}

		if batch.Issuer != issuerAddr {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgMintBatchCredits, "only batch issuer can mint additional credits"), nil, nil
		}

		msg := &types.MsgMintBatchCredits{
			Issuer:     issuerAddr,
			BatchDenom: batch.Denom,
			Issuance:   generateBatchIssuance(r, accs),
			OriginTx: &types.OriginTx{
				Source: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 64)),
				Id:     simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 64)),
			},
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, issuerAddr, TypeMsgUpdateClassIssuers)
		if spendable == nil {
			return op, nil, err
		}
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgSealBatch generates a MsgSealBatch with random values.
func SimulateMsgSealBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuerAcc, _ := simtypes.RandomAcc(r, accs)
		issuerAddr := issuerAcc.Address.String()

		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgSealBatch)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgSealBatch, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgSealBatch, project.Id)
		if batch == nil {
			return op, nil, err
		}

		if batch.Issuer != issuerAddr {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSealBatch, "only batch issuer can seal batch"), nil, nil
		}

		if !batch.Open {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSealBatch, "batch is closed"), nil, nil
		}

		msg := &types.MsgSealBatch{
			Issuer:     issuerAddr,
			BatchDenom: batch.Denom,
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, issuerAddr, TypeMsgSealBatch)
		if spendable == nil {
			return op, nil, err
		}
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgBridge generates a MsgBridge with random values.
func SimulateMsgBridge(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, qryClient types.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := sdk.WrapSDKContext(sdkCtx)
		class, op, err := utils.GetRandomClass(sdkCtx, r, qryClient, TypeMsgBridge)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgBridge, class.Id)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgBridge, project.Id)
		if batch == nil {
			return op, nil, err
		}

		issuersRes, err := qryClient.ClassIssuers(ctx, &types.QueryClassIssuersRequest{
			ClassId: class.Id,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		issuers := issuersRes.Issuers
		owner := issuers[r.Intn(len(issuers))]
		ownerAddr, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		_, found := simtypes.FindAccount(accs, ownerAddr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "not a simulation account"), nil, nil
		}

		balanceRes, err := qryClient.Balance(ctx, &types.QueryBalanceRequest{
			Address:    owner,
			BatchDenom: batch.Denom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.Balance.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "balance is zero"), nil, nil
		}

		tradable, err := tradableBalance.Int64()
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, err.Error()), nil, nil
		}

		amount := 1
		if tradable > 1 {
			amount = simtypes.RandIntBetween(r, 1, int(tradable))
		}

		msg := &types.MsgBridge{
			Target:    "polygon",
			Recipient: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
			Owner:     owner,
			Credits: []*types.Credits{
				{
					BatchDenom: batch.Denom,
					Amount:     fmt.Sprintf("%d", amount),
				},
			},
		}

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, owner, TypeMsgBridge)
		if spendable == nil {
			return op, nil, err
		}
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "fee error"), nil, err
		}

		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		acc := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)

		tx, err := helpers.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			account.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if !strings.Contains(err.Error(), "only credits previously bridged from another chain") {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgBridge, "unable to deliver tx"), nil, err
			}
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

// SimulateMsgAddCreditType generates a MsgAddCreditType with random values.
func SimulateMsgAddCreditType(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgAddCreditType)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, "unable to generate deposit"), nil, err
		}

		abbrev := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 1, 3))
		abbrev = strings.ToUpper(abbrev)
		name := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 1, 10))

		_, err = qryClient.CreditType(sdkCtx, &types.QueryCreditTypeRequest{
			Abbreviation: abbrev,
		})
		if err != nil {
			if !ormerrors.NotFound.Is(err) {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, err.Error()), nil, err
			}
		}

		proposalMsg := types.MsgAddCreditType{
			Authority: authority.String(),
			CreditType: &types.CreditType{
				Abbreviation: abbrev,
				Name:         name,
				Unit:         "kg",
				Precision:    6,
			},
		}

		any, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddCreditType, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Messages:       []*codectypes.Any{any},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgAddClassCreator generates a MsgAddClassCreator with random values.
func SimulateMsgAddClassCreator(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgAddClassCreator)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddClassCreator, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddClassCreator, "unable to generate deposit"), nil, err
		}

		creatorsResult, err := qryClient.AllowedClassCreators(sdkCtx, &types.QueryAllowedClassCreatorsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddClassCreator, err.Error()), nil, err
		}

		if stringInSlice(proposerAddr, creatorsResult.ClassCreators) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddClassCreator, "class creator already exists"), nil, nil
		}

		proposalMsg := types.MsgAddClassCreator{
			Authority: authority.String(),
			Creator:   proposerAddr,
		}

		any, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgAddClassCreator, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Messages:       []*codectypes.Any{any},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgRemoveClassCreator generates a MsgRemoveClassCreator with random values.
func SimulateMsgRemoveClassCreator(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgRemoveClassCreator)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveClassCreator, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveClassCreator, "unable to generate deposit"), nil, err
		}

		creatorsResult, err := qryClient.AllowedClassCreators(sdkCtx, &types.QueryAllowedClassCreatorsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveClassCreator, err.Error()), nil, err
		}

		if !stringInSlice(proposerAddr, creatorsResult.ClassCreators) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveClassCreator, "unknown class creator"), nil, nil
		}

		proposalMsg := types.MsgRemoveClassCreator{
			Authority: authority.String(),
			Creator:   proposerAddr,
		}

		any, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRemoveClassCreator, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Messages:       []*codectypes.Any{any},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgSetClassCreatorAllowlist generates a MsgSetClassCreatorAllowlist with random values.
func SimulateMsgSetClassCreatorAllowlist(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgSetClassCreatorAllowlist)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSetClassCreatorAllowlist, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSetClassCreatorAllowlist, "unable to generate deposit"), nil, err
		}

		proposalMsg := types.MsgSetClassCreatorAllowlist{
			Authority: authority.String(),
			Enabled:   r.Float32() < 0.3, // 30% chance of allowlist being enabled,
		}

		any, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSetClassCreatorAllowlist, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Messages:       []*codectypes.Any{any},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

// SimulateMsgUpdateClassFee generates a MsgToggleClassAllowlist with random values.
func SimulateMsgUpdateClassFee(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, govk ecocredit.GovKeeper,
	qryClient types.QueryServer, authority sdk.AccAddress) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		proposer, _ := simtypes.RandomAcc(r, accs)
		proposerAddr := proposer.Address.String()

		spendable, account, op, err := utils.GetAccountAndSpendableCoins(sdkCtx, bk, accs, proposerAddr, TypeMsgUpdateClassFee)
		if spendable == nil {
			return op, nil, err
		}

		params := govk.GetDepositParams(sdkCtx)
		deposit, skip, err := utils.RandomDeposit(r, sdkCtx, ak, bk, params, proposer.Address)
		switch {
		case skip:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassFee, "skip deposit"), nil, nil
		case err != nil:
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassFee, "unable to generate deposit"), nil, err
		}

		fee := utils.RandomFee(r)
		if fee.Amount.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassFee, "invalid proposal message"), nil, err
		}

		proposalMsg := types.MsgUpdateClassFee{
			Authority: authority.String(),
			Fee:       &fee,
		}

		any, err := codectypes.NewAnyWithValue(&proposalMsg)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassFee, err.Error()), nil, err
		}

		msg := &govtypes.MsgSubmitProposal{
			Messages:       []*codectypes.Any{any},
			InitialDeposit: deposit,
			Proposer:       proposerAddr,
			Metadata:       simtypes.RandStringOfLength(r, 10),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(r, txCtx)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func getClassIssuers(ctx sdk.Context, qryClient types.QueryServer, className string, msgType string) ([]string, simtypes.OperationMsg, error) {
	classIssuers, err := qryClient.ClassIssuers(sdk.WrapSDKContext(ctx), &types.QueryClassIssuersRequest{ClassId: className})
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return []string{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no credit classes"), nil
		}

		return []string{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	return classIssuers.Issuers, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomProjectFromClass(ctx context.Context, r *rand.Rand, qryClient types.QueryServer, msgType, classID string) (*types.ProjectInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.ProjectsByClass(ctx, &types.QueryProjectsByClassRequest{
		ClassId: classID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	projects := res.Projects
	if len(projects) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no projects"), nil
	}

	return projects[r.Intn(len(projects))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomBatchFromProject(ctx context.Context, r *rand.Rand, qryClient types.QueryServer, msgType, projectID string) (*types.BatchInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.BatchesByProject(ctx, &types.QueryBatchesByProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		if strings.Contains(err.Error(), ormerrors.NotFound.Error()) {
			return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, fmt.Sprintf("no credit batches for %s project", projectID)), nil
		}
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	batches := res.Batches
	if len(batches) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, fmt.Sprintf("no credit batches for %s project", projectID)), nil
	}
	return batches[r.Intn(len(batches))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func randomIssuers(r *rand.Rand, accounts []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 3, 10)

	var issuers []string
	issuersMap := make(map[string]bool)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		addr := acc.Address.String()
		if _, ok := issuersMap[addr]; ok {
			continue
		}
		issuersMap[acc.Address.String()] = true
		issuers = append(issuers, addr)
	}

	return issuers
}

func generateBatchIssuance(r *rand.Rand, accs []simtypes.Account) []*types.BatchIssuance {
	numIssuances := simtypes.RandIntBetween(r, 3, 10)
	res := make([]*types.BatchIssuance, numIssuances)

	for i := 0; i < numIssuances; i++ {
		recipient := accs[i]
		retiredAmount := simtypes.RandIntBetween(r, 0, 100)
		var retirementJurisdiction string
		if retiredAmount > 0 {
			retirementJurisdiction = "AD"
		}
		res[i] = &types.BatchIssuance{
			Recipient:              recipient.Address.String(),
			TradableAmount:         fmt.Sprintf("%d", simtypes.RandIntBetween(r, 10, 1000)),
			RetiredAmount:          fmt.Sprintf("%d", retiredAmount),
			RetirementJurisdiction: retirementJurisdiction,
		}
	}

	return res
}
