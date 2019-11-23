package urls

import (
	"fmt"
	"net/url"
	"strconv"
)

// генерирует url к steam api. использует токен, выданый ему при создании для более удобной генерации ссылок
type APIUrls struct {
	key string
}

func New(key string) APIUrls {
	return APIUrls{key: key}
}

func (a APIUrls) APITradeOffers(lang string, received, sent, descriptions, onlyActive, onlyHistorical bool, historicalCutoff int) string {
	return APITradeOffers(a.key, lang, received, sent, descriptions, onlyActive, onlyHistorical, historicalCutoff)
}

// APITradeOffers генерирует ссылку на получение списка торговых предложений. Есть возможность фильтровать торговые
// предложения по допустимым параметрам
func APITradeOffers(key, lang string, received, sent, descriptions, onlyActive, onlyHistorical bool, historicalCutoff int) string {
	query := url.Values{}
	query.Add("key", key)
	query.Add("language", lang)
	if received {
		query.Add("get_received_offers", "1")
	} else {
		query.Add("get_received_offers", "0")
	}
	if sent {
		query.Add("get_sent_offers", "1")
	} else {
		query.Add("get_sent_offers", "0")
	}
	if descriptions {
		query.Add("get_descriptions", "1")
	} else {
		query.Add("get_descriptions", "0")
	}
	if onlyActive {
		query.Add("active_only", "1")
	} else {
		query.Add("active_only", "0")
	}
	if onlyHistorical {
		query.Add("historical_only", "1")
	} else {
		query.Add("historical_only", "0")
	}
	query.Add("time_historical_cutoff", strconv.Itoa(historicalCutoff))
	u := url.URL{
		Scheme:   "https",
		Host:     "api.steampowered.com",
		Path:     "IEconService/GetTradeOffers/v1",
		RawQuery: query.Encode(),
	}
	return u.String()
}

func ItemImageByClassID(appID, classID, h, w uint) string {
	return fmt.Sprintf("https://steamcommunity-a.akamaihd.net/economy/image/class/%d/%d/%dx%d", appID, classID, h, w)
}
