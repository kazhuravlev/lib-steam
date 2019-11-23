package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseURL = "https://api.steampowered.com"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrBadRequest   = errors.New("bad request")
	ErrNotFound     = errors.New("some items not found")
	ErrBadResponse  = errors.New("bad steam response")
)

type Client struct {
	apiKey string
}

func New(apiKey string) (*Client, error) {
	if len(apiKey) != 32 {
		return nil, ErrInvalidToken
	}
	return &Client{apiKey: apiKey}, nil
}

// GetTradeOfferList список торговых предложений
func (c *Client) GetTradeOfferList() (*TradeOffers, error) {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("get_received_offers", "1")
	query.Set("get_sent_offers", "1")
	query.Set("get_descriptions", "1")
	query.Set("language", "english")
	query.Set("active_only", "1")
	query.Set("historical_only", "0")
	query.Set("time_historical_cutoff", "1")

	u := baseURL + "/IEconService/GetTradeOffers/v1/?" + query.Encode()

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp := RespGetTradeOffers{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Response, nil
}

// GetTradeOffer торговое предложение по идентификатору
func (c *Client) GetTradeOffer(tradeOfferID string) (*TradeOffer, error) {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("tradeofferid", tradeOfferID)
	query.Set("language", "english")

	u := baseURL + "/IEconService/GetTradeOffer/v1/?" + query.Encode()

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp := RespGetTradeOffer{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	if apiResp.Response.Offer.Tradeofferid == "" {
		return nil, ErrNotFound
	}

	return &apiResp.Response.Offer, nil
}

// GetTradeOffersSummary счетики
func (c *Client) GetTradeOffersSummary() (*TradeOffersSummary, error) {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("time_last_visit", "1")
	query.Set("language", "english")

	u := baseURL + "/IEconService/GetTradeOffersSummary/v1/?" + query.Encode()

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp := RespGetTradeOffersSummary{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Response, nil
}

// DeclineTradeOffer отклоняет входящую сделку
func (c *Client) DeclineTradeOffer(tradeOfferID string) error {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("tradeofferid", tradeOfferID)

	u := baseURL + "/IEconService/DeclineTradeOffer/v1/?" + query.Encode()

	resp, err := http.Post(u, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// CancelTradeOffer отменяет исходящую сделку
func (c *Client) CancelTradeOffer(tradeOfferID string) error {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("tradeofferid", tradeOfferID)

	u := baseURL + "/IEconService/CancelTradeOffer/v1/?" + query.Encode()

	resp, err := http.Post(u, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) GetAssetClassInfo(appID uint, classList []uint, language string) ([]AssetClassInfo, []uint, error) {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("language", language)
	query.Set("appid", strconv.FormatUint(uint64(appID), 10))
	classCount := 0
	uniqClasses := map[uint]bool{}
	for _, classID := range classList {
		if _, exists := uniqClasses[classID]; !exists {
			query.Set(fmt.Sprintf("classid%d", classCount), strconv.FormatUint(uint64(classID), 10))
			uniqClasses[classID] = true
			classCount++
		}
	}
	if classCount == 0 {
		return nil, nil, ErrBadRequest
	}
	query.Set("class_count", strconv.FormatUint(uint64(classCount), 10))

	u := baseURL + "/ISteamEconomy/GetAssetClassInfo/v1/?" + query.Encode()
	resp, err := http.Get(u)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	apiResp := ResponseGetAssetClassInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, nil, err
	}

	result := []AssetClassInfo{}
	for k, v := range apiResp.Result {
		if k == "success" {
			success, ok := v.(bool)
			delete(apiResp.Result, "success")
			delete(apiResp.Result, "error")
			notFoundItems := []uint{}
			for id, val := range apiResp.Result {
				uintID, err := strconv.ParseUint(id, 10, 64)
				if err != nil {
					continue
				}
				if len(val.(map[string]interface{})) == 0 {
					notFoundItems = append(notFoundItems, uint(uintID))
				}
			}
			if !ok || !success {
				return nil, notFoundItems, ErrNotFound
			}
			continue
		}
		rec, err := parseAssetInfo(v)
		if err == nil {
			result = append(result, rec)
		}
	}

	return result, nil, nil
}

// GetCSGOGameServersStatus получить состояние серверов CSGO
func (c *Client) GetCSGOGameServersStatus() (*GameServersStatus, error) {
	query := url.Values{}
	query.Set("key", c.apiKey)
	query.Set("format", "json")

	u := baseURL + "/ICSGOServers_730/GetGameServersStatus/v1/?" + query.Encode()

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp := RespGetGameServersStatus{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &apiResp.Result, nil
}
