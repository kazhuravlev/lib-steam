package api

import (
	"os"
	"testing"

	"github.com/kazhuravlev/lib-steam/trading"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	token = os.Getenv("steam_token")
)

func TestGetTradeOfferList(t *testing.T) {
	c, err := New(token)
	require.Nil(t, err)

	offers, err := c.GetTradeOfferList()
	assert.Nil(t, err)
	require.NotNil(t, offers)

	assert.NotEqual(t, 0, len(offers.Descriptions))
	assert.NotEqual(t, 0, len(offers.TradeOffersReceived))
	assert.NotEqual(t, 0, len(offers.TradeOffersSent))

	for _, offer := range offers.TradeOffersReceived {
		assert.Contains(t, trading.TradeOfferStates, offer.TradeOfferState)
		assert.Equal(t, false, offer.IsOurOffer)
		assert.NotEqual(t, 0, len(offer.ItemsToGive)+len(offer.ItemsToReceive))
	}

	for _, offer := range offers.TradeOffersSent {
		assert.Contains(t, trading.TradeOfferStates, offer.TradeOfferState)
		assert.Equal(t, true, offer.IsOurOffer)
		assert.NotEqual(t, 0, len(offer.ItemsToGive)+len(offer.ItemsToReceive))
	}
}

func TestGetGetAssetClassInfo(t *testing.T) {
	c, err := New(token)
	require.Nil(t, err)

	result, xx, err := c.GetAssetClassInfo(570, []uint{1506969666, 123123}, "russian")
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, result)
	assert.Equal(t, 1, len(xx))
}

func TestGetTradeOffer(t *testing.T) {
	c, err := New(token)
	require.Nil(t, err)

	tradeOffer, err := c.GetTradeOffer("123")
	assert.Equal(t, ErrNotFound, err)
	assert.Nil(t, tradeOffer)
}

func TestGetGameServersStatus(t *testing.T) {
	c, err := New(token)
	require.Nil(t, err)

	status, err := c.GetCSGOGameServersStatus()
	assert.Nil(t, err)
	assert.NotNil(t, status)

	assert.Equal(t, 17, len(status.Datacenters))
	for dc, data := range status.Datacenters {
		assert.Contains(t, []DatacenterValue{
			DCAustralia,
			DCBrazil,
			DCChile,
			DCEUEast,
			DCEUNorth,
			DCEUWest,
			DCEmirates,
			DCHongKong,
			DCIndia,
			DCJapan,
			DCSingapore,
			DCSouthAfrica,
			DCSpain,
			DCUSNortheast,
			DCUSNorthwest,
			DCUSSoutheast,
			DCUSSouthwest,
		}, dc)
		assert.Contains(t, []LoadValue{
			LoadLow,
			LoadNormal,
			LoadMedium,
			LoadIdle,
			LoadHight,
		}, data.Load)
	}
}
