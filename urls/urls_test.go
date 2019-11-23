package urls

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func urlEquals(t *testing.T, a, b string) {
	u1, err := url.Parse(a)
	assert.Nil(t, err)

	u2, err := url.Parse(b)
	assert.Nil(t, err)

	assert.Equal(t, u1.Scheme, u2.Scheme)
	assert.Equal(t, u1.Opaque, u2.Opaque)
	assert.Equal(t, u1.User, u2.User)
	assert.Equal(t, u1.Host, u2.Host)
	assert.Equal(t, u1.Path, u2.Path)
	assert.Equal(t, u1.RawPath, u2.RawPath)
	assert.Equal(t, u1.Query().Encode(), u2.Query().Encode())
	assert.Equal(t, u1.Fragment, u2.Fragment)
}

func TestAPIUrlsNew(t *testing.T) {
	assert.Equal(t, "MyKey123", New("MyKey123").key)
}

func TestAPIUrlsBindings(t *testing.T) {
	u := New("MyKey")

	assert.Equal(t,
		APITradeOffers("MyKey", "RussIaN", true, true, true, true, true, 10),
		u.APITradeOffers("RussIaN", true, true, true, true, true, 10))

}

func TestAPITradeOffers(t *testing.T) {
	u := APITradeOffers("MyKey", "RussIaN", true, true, true, true, true, 10)
	expected := "https://api.steampowered.com/IEconService/GetTradeOffers/v1?key=MyKey&get_received_offers=1&get_sent_offers=1&get_descriptions=1&language=RussIaN&active_only=1&historical_only=1&time_historical_cutoff=10"
	urlEquals(t, expected, u)
}
