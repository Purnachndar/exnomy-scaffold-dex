package dex_engine

import (
	"github.com/Purnachndar/exnomy-scaffold-dex/backend/models"
	"github.com/Purnachndar/exnomy-sdk-backend/sdk"
	"github.com/Purnachndar/exnomy-sdk-backend/utils"
)

func getHydroOrderFromModelOrder(orderJSON *models.OrderJSON) *sdk.Order {
	return sdk.NewOrderWithData(
		orderJSON.Trader,
		orderJSON.Relayer,
		orderJSON.BaseCurrency,
		orderJSON.QuoteCurrency,
		utils.DecimalToBigInt(orderJSON.BaseCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.QuoteCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.GasTokenHugeAmount),
		orderJSON.Data,
		orderJSON.Signature,
	)
}

func getHydroOrderHashHexFromOrderJson(orderJSON *models.OrderJSON) string {
	order := sdk.NewOrderWithData(
		orderJSON.Trader,
		orderJSON.Relayer,
		orderJSON.BaseCurrency,
		orderJSON.QuoteCurrency,
		utils.DecimalToBigInt(orderJSON.BaseCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.QuoteCurrencyHugeAmount),
		utils.DecimalToBigInt(orderJSON.GasTokenHugeAmount),
		orderJSON.Data,
		"",
	)

	return utils.Bytes2HexP(hydroProtocol.GetOrderHash(order))
}
