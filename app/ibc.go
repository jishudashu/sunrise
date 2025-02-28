package app

import (
	"cosmossdk.io/core/appmodule"
	storetypes "cosmossdk.io/store/types"
	govtypes "cosmossdk.io/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icamodule "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v9/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v9/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v9/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v9/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v9/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v9/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v9/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v9/modules/core"
	ibcclienttypes "github.com/cosmos/ibc-go/v9/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v9/modules/core/03-connection/types"
	porttypes "github.com/cosmos/ibc-go/v9/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v9/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v9/modules/core/keeper"
	solomachine "github.com/cosmos/ibc-go/v9/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v9/modules/light-clients/07-tendermint"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/runtime"

	swapmodule "github.com/sunriselayer/sunrise/x/swap/module"
	// this line is used by starport scaffolding # ibc/app/import
)

// registerIBCModules register IBC keepers and non dependency inject modules.
func (app *App) registerIBCModules() error {
	// set up non depinject support modules store keys
	if err := app.RegisterStores(
		storetypes.NewKVStoreKey(ibcexported.StoreKey),
		storetypes.NewKVStoreKey(ibctransfertypes.StoreKey),
		storetypes.NewKVStoreKey(ibcfeetypes.StoreKey),
		storetypes.NewKVStoreKey(icahosttypes.StoreKey),
		storetypes.NewKVStoreKey(icacontrollertypes.StoreKey),
	); err != nil {
		panic(err)
	}

	// register the key tables for legacy param subspaces
	keyTable := ibcclienttypes.ParamKeyTable()
	keyTable.RegisterParamSet(&ibcconnectiontypes.Params{})
	app.ParamsKeeper.Subspace(ibcexported.ModuleName).WithKeyTable(keyTable)
	app.ParamsKeeper.Subspace(ibctransfertypes.ModuleName).WithKeyTable(ibctransfertypes.ParamKeyTable())
	app.ParamsKeeper.Subspace(icacontrollertypes.SubModuleName).WithKeyTable(icacontrollertypes.ParamKeyTable())
	app.ParamsKeeper.Subspace(icahosttypes.SubModuleName).WithKeyTable(icahosttypes.ParamKeyTable())

	govModuleAddr, _ := app.AuthKeeper.AddressCodec().BytesToString(authtypes.NewModuleAddress(govtypes.ModuleName))

	// Create IBC keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		app.appCodec,
		runtime.NewEnvironment(
			runtime.NewKVStoreService(app.GetKey(ibcexported.StoreKey)), app.Logger().With(log.ModuleKey, "x/ibc"),
			runtime.EnvWithMsgRouterService(app.MsgServiceRouter())),
		app.GetSubspace(ibcexported.ModuleName),
		app.UpgradeKeeper,
		govModuleAddr,
	)

	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		app.appCodec,
		runtime.NewEnvironment(
			runtime.NewKVStoreService(app.GetKey(ibcfeetypes.StoreKey)), app.Logger().With(log.ModuleKey, "x/transfer"),
			runtime.EnvWithMsgRouterService(app.MsgServiceRouter())),
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		app.AuthKeeper,
		app.BankKeeper,
	)

	// Create IBC transfer keeper
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		app.appCodec,
		runtime.NewEnvironment(
			runtime.NewKVStoreService(app.GetKey(ibctransfertypes.StoreKey)), app.Logger().With(log.ModuleKey, "x/transfer"),
			runtime.EnvWithMsgRouterService(app.MsgServiceRouter())),
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCFeeKeeper,
		app.IBCKeeper.ChannelKeeper,
		app.AuthKeeper,
		app.BankKeeper,
		govModuleAddr,
	)

	// Create interchain account keepers
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		app.appCodec,
		runtime.NewEnvironment(
			runtime.NewKVStoreService(app.GetKey(icahosttypes.StoreKey)), app.Logger().With(log.ModuleKey, "x/icacontroller"),
			runtime.EnvWithMsgRouterService(app.MsgServiceRouter())),
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // ICS4Wrapper
		app.IBCKeeper.ChannelKeeper,
		app.AuthKeeper,
		govModuleAddr,
	)

	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		app.appCodec,
		runtime.NewEnvironment(
			runtime.NewKVStoreService(app.GetKey(icacontrollertypes.StoreKey)), app.Logger().With(log.ModuleKey, "x/icacontroller"),
			runtime.EnvWithMsgRouterService(app.MsgServiceRouter())),
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.IBCKeeper.ChannelKeeper,
		govModuleAddr,
	)

	// Create IBC modules with ibcfee middleware
	// transferIBCModule := ibcfee.NewIBCMiddleware(ibctransfer.NewIBCModule(app.TransferKeeper), app.IBCFeeKeeper)
	transferIBCModuleIbcFee := ibcfee.NewIBCMiddleware(ibctransfer.NewIBCModule(app.TransferKeeper), app.IBCFeeKeeper)
	transferIBCModule := swapmodule.NewIBCMiddleware(transferIBCModuleIbcFee, &app.SwapKeeper)

	// integration point for custom authentication modules
	icaControllerIBCModule := ibcfee.NewIBCMiddleware(
		icacontroller.NewIBCMiddleware(app.ICAControllerKeeper),
		app.IBCFeeKeeper,
	)

	icaHostIBCModule := ibcfee.NewIBCMiddleware(icahost.NewIBCModule(app.ICAHostKeeper), app.IBCFeeKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter().
		AddRoute(ibctransfertypes.ModuleName, transferIBCModule).
		AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule).
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule)

	// this line is used by starport scaffolding # ibc/app/module

	app.IBCKeeper.SetRouter(ibcRouter)

	clientKeeper := app.IBCKeeper.ClientKeeper
	storeProvider := app.IBCKeeper.ClientKeeper.GetStoreProvider()

	tmLightClientModule := ibctm.NewLightClientModule(app.appCodec, storeProvider)
	clientKeeper.AddRoute(ibctm.ModuleName, &tmLightClientModule)

	smLightClientModule := solomachine.NewLightClientModule(app.appCodec, storeProvider)
	clientKeeper.AddRoute(solomachine.ModuleName, &smLightClientModule)

	// register IBC modules
	if err := app.RegisterModules(
		// IBC modules
		ibc.NewAppModule(app.appCodec, app.IBCKeeper),
		ibctransfer.NewAppModule(app.appCodec, app.TransferKeeper),
		ibcfee.NewAppModule(app.appCodec, app.IBCFeeKeeper),
		icamodule.NewAppModule(app.appCodec, &app.ICAControllerKeeper, &app.ICAHostKeeper),

		// IBC light clients
		ibctm.NewAppModule(tmLightClientModule),
		solomachine.NewAppModule(smLightClientModule),
	); err != nil {
		return err
	}

	return nil
}

// Since the IBC modules don't support dependency injection, we need to
// manually register the modules on the client side.
// This needs to be removed after IBC supports App Wiring.
func RegisterIBC(cdc codec.Codec, registry cdctypes.InterfaceRegistry) map[string]appmodule.AppModule {
	modules := map[string]appmodule.AppModule{
		ibcexported.ModuleName:      ibc.NewAppModule(cdc, &ibckeeper.Keeper{}),
		ibctransfertypes.ModuleName: ibctransfer.NewAppModule(cdc, ibctransferkeeper.Keeper{}),
		ibcfeetypes.ModuleName:      ibcfee.NewAppModule(cdc, ibcfeekeeper.Keeper{}),
		icatypes.ModuleName:         icamodule.NewAppModule(cdc, &icacontrollerkeeper.Keeper{}, &icahostkeeper.Keeper{}),
		ibctm.ModuleName:            ibctm.NewAppModule(ibctm.NewLightClientModule(cdc, ibcclienttypes.StoreProvider{})),
		solomachine.ModuleName:      solomachine.NewAppModule(solomachine.NewLightClientModule(cdc, ibcclienttypes.StoreProvider{})),
	}

	for _, m := range modules {
		if mr, ok := m.(appmodule.HasRegisterInterfaces); ok {
			mr.RegisterInterfaces(registry)
		}
	}

	return modules
}
