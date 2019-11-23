package api

import "github.com/kazhuravlev/lib-steam/trading"

type LoadValue string
type CapacityValue string
type DatacenterValue string

const (
	LoadLow    LoadValue = "low"
	LoadNormal LoadValue = "normal"
	LoadMedium LoadValue = "medium"
	LoadIdle   LoadValue = "idle"
	LoadHight  LoadValue = "hight"

	CapacityFull CapacityValue = "full"

	DCAustralia   DatacenterValue = "Australia"
	DCBrazil      DatacenterValue = "Brazil"
	DCChile       DatacenterValue = "Chile"
	DCEUEast      DatacenterValue = "EU East"
	DCEUNorth     DatacenterValue = "EU North"
	DCEUWest      DatacenterValue = "EU West"
	DCEmirates    DatacenterValue = "Emirates"
	DCHongKong    DatacenterValue = "Hong Kong"
	DCIndia       DatacenterValue = "India"
	DCJapan       DatacenterValue = "Japan"
	DCSingapore   DatacenterValue = "Singapore"
	DCSouthAfrica DatacenterValue = "South Africa"
	DCSpain       DatacenterValue = "Spain"
	DCUSNortheast DatacenterValue = "US Northeast"
	DCUSNorthwest DatacenterValue = "US Northwest"
	DCUSSoutheast DatacenterValue = "US Southeast"
	DCUSSouthwest DatacenterValue = "US Southwest"
)

type GameServersStatus struct {
	App struct {
		Time      string `json:"time"`
		Timestamp int    `json:"timestamp"`
		Version   int    `json:"version"`
	} `json:"app"`
	Datacenters map[DatacenterValue]struct {
		Capacity CapacityValue `json:"capacity"`
		Load     LoadValue     `json:"load"`
	} `json:"datacenters"`
	Matchmaking struct {
		OnlinePlayers    int    `json:"online_players"`
		OnlineServers    int    `json:"online_servers"`
		Scheduler        string `json:"scheduler"`
		SearchSecondsAvg int    `json:"search_seconds_avg"`
		SearchingPlayers int    `json:"searching_players"`
	} `json:"matchmaking"`
	Services struct {
		IEconItems    string `json:"IEconItems"`
		Leaderboards  string `json:"Leaderboards"`
		SessionsLogon string `json:"SessionsLogon"`
		// эта балалайка может быть "normal", "delayed"
		SteamCommunity string `json:"SteamCommunity"`
	} `json:"services"`
}

type RespGetGameServersStatus struct {
	Result GameServersStatus `json:"result"`
}

type TradeOffers struct {
	Descriptions []struct {
		Appid           uint   `json:"appid"`
		BackgroundColor string `json:"background_color"`
		Classid         string `json:"classid"`
		Commodity       bool   `json:"commodity"`
		Currency        bool   `json:"currency"`
		Descriptions    []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"descriptions"`
		IconURL                     string `json:"icon_url"`
		IconURLLarge                string `json:"icon_url_large"`
		Instanceid                  string `json:"instanceid"`
		MarketHashName              string `json:"market_hash_name"`
		MarketMarketableRestriction int    `json:"market_marketable_restriction"`
		MarketName                  string `json:"market_name"`
		MarketTradableRestriction   int    `json:"market_tradable_restriction"`
		Name                        string `json:"name"`
		NameColor                   string `json:"name_color"`
		Tradable                    bool   `json:"tradable"`
		Type                        string `json:"type"`
	} `json:"descriptions"`
	TradeOffersReceived []TradeOffer `json:"trade_offers_received"`
	TradeOffersSent     []TradeOffer `json:"trade_offers_sent"`
}

type RespGetTradeOffers struct {
	Response TradeOffers `json:"response"`
}

type TradeOffersSummary struct {
	EscrowReceivedCount     uint `json:"escrow_received_count"`
	EscrowSentCount         uint `json:"escrow_sent_count"`
	HistoricalReceivedCount uint `json:"historical_received_count"`
	HistoricalSentCount     uint `json:"historical_sent_count"`
	NewReceivedCount        uint `json:"new_received_count"`
	NewlyAcceptedSentCount  uint `json:"newly_accepted_sent_count"`
	PendingReceivedCount    uint `json:"pending_received_count"`
	PendingSentCount        uint `json:"pending_sent_count"`
	UpdatedReceivedCount    uint `json:"updated_received_count"`
	UpdatedSentCount        uint `json:"updated_sent_count"`
}

type RespGetTradeOffersSummary struct {
	Response TradeOffersSummary `json:"response"`
}

type Item struct {
	Amount     string `json:"amount"`
	Appid      string `json:"appid"`
	Assetid    string `json:"assetid"`
	Classid    string `json:"classid"`
	Contextid  string `json:"contextid"`
	Instanceid string `json:"instanceid"`
	Missing    bool   `json:"missing"`
}

type TradeOffer struct {
	AccountidOther     uint                    `json:"accountid_other"`
	ConfirmationMethod int                     `json:"confirmation_method"`
	EscrowEndDate      int                     `json:"escrow_end_date"`
	ExpirationTime     int                     `json:"expiration_time"`
	FromRealTimeTrade  bool                    `json:"from_real_time_trade"`
	IsOurOffer         bool                    `json:"is_our_offer"`
	Message            string                  `json:"message"`
	TimeCreated        int                     `json:"time_created"`
	TimeUpdated        int                     `json:"time_updated"`
	TradeOfferState    trading.TradeOfferState `json:"trade_offer_state"`
	Tradeofferid       string                  `json:"tradeofferid"`
	TradeId            string                  `json:"tradeid"`
	ItemsToReceive     []Item                  `json:"items_to_receive"`
	ItemsToGive        []Item                  `json:"items_to_give"`
}

type RespGetTradeOffer struct {
	Response struct {
		Offer TradeOffer `json:"offer"`
	} `json:"response"`
}

type RespResolveVanityURL struct {
	Response struct {
		Steamid string `json:"steamid"`
		Success int    `json:"success"`
	} `json:"response"`
}

type ResponseGetAssetClassInfo struct {
	Result map[string]interface{} `json:"result"`
}

type AssetClassInfo struct {
	BackgroundColor             string                 `json:"background_color"`
	Classid                     string                 `json:"classid"`
	Commodity                   string                 `json:"commodity"`
	Descriptions                map[string]Description `json:"descriptions"`
	Fraudwarnings               string                 `json:"fraudwarnings"`
	IconDragURL                 string                 `json:"icon_drag_url"`
	IconURL                     string                 `json:"icon_url"`
	IconURLLarge                string                 `json:"icon_url_large"`
	MarketHashName              string                 `json:"market_hash_name"`
	MarketName                  string                 `json:"market_name"`
	MarketTradableRestriction   string                 `json:"market_tradable_restriction"`
	MarketMarketableRestriction string                 `json:"market_marketable_restriction,omitempty"`
	ItemExpiration              string                 `json:"item_expiration,omitempty"`
	Marketable                  string                 `json:"marketable"`
	Name                        string                 `json:"name"`
	NameColor                   string                 `json:"name_color"`
	OwnerDescriptions           string                 `json:"owner_descriptions"`
	Tags                        map[string]Tag         `json:"tags"`
	Tradable                    string                 `json:"tradable"`
	Type                        string                 `json:"type"`
}

type Tag struct {
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
	InternalName string `json:"internal_name"`
	Name         string `json:"name"`
	Color        string `json:"color,omitempty"`
}

type Description struct {
	Color   string `json:"color,omitempty"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	AppData struct {
		DefInfex      string `json:"def_index"`
		IsItemsetName string `json:"is_itemset_name"`
	} `json:"app_data"`
}
