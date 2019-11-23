package inventory

import (
	log "github.com/sirupsen/logrus"
)

// NormalizedInventory описывает нормализованный инвентарь. Каждый объект в массиве содержит детальную информацию о
// предмете
type NormalizedInventory struct {
	Items []NormalizedItem
}

// NormalizedItem предмет в инвентаре пользвоателя. нормализован
type NormalizedItem struct {
	AppID                     uint32                  `json:"app_id"`
	ID                        uint64                  `json:"id"`
	ClassID                   uint64                  `json:"class_id"`
	InstanceID                uint64                  `json:"instance_id"`
	Amount                    uint64                  `json:"amount"`
	Pos                       uint32                  `json:"pos"`
	IconURL                   string                  `json:"icon_url"`
	IconURLLarge              string                  `json:"icon_url_large"`
	IconDragURL               string                  `json:"icon_drag_url"`
	Name                      string                  `json:"name"`
	MarketName                string                  `json:"market_name"`
	MarketHashName            string                  `json:"market_hash_name"`
	NameColor                 string                  `json:"name_color"`
	BackgroundColor           string                  `json:"background_color"`
	Type                      string                  `json:"type"`
	Tradable                  bool                    `json:"tradable"`
	Marketable                bool                    `json:"marketable"`
	Commodity                 bool                    `json:"commodity"`
	MarketTradableRestriction uint32                  `json:"market_tradable_restriction"`
	Descriptions              []NormalizedDescription `json:"descriptions"`
	Actions                   []NormalizedAction      `json:"actions"`
	AppData                   map[string]string       `json:"appdata"`
	Tags                      []NormalizedTag         `json:"tags"`
}

type NormalizedDescription struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	Color string `json:"color"`
}

type NormalizedAction struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type NormalizedTag struct {
	InternalName string `json:"internal_name"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	CategoryName string `json:"category_name"`
}

func buildIndex(items Items) map[uint32]string {
	res := make(map[uint32]string, len(items))
	for key, item := range items {
		res[item.Pos] = key
	}
	return res
}

// NormalizeInventory приводит исходный инвентарь к нормализованному
func NormalizeInventory(srcInv Inventory, limit, offset int) NormalizedInventory {
	//srcItemsLength := len(srcInv.Items)
	//if limit < 0 {
	//	limit = srcItemsLength
	//}
	//if offset < 0 {
	//	offset = 0
	//}
	//if limit > srcItemsLength - offset {
	//	limit = srcItemsLength - offset
	//}

	//itemsIndex := buildIndex(srcInv.Items)
	items := make([]NormalizedItem, len(srcInv.Items))
	index := 0
	//var offsetLimit uint32 = 0

	//for x := 0; x <
	//limitPlusOffset := limit + offset
	for _, srcItem := range srcInv.Items {
		//if offsetLimit < srcItem.Pos-1 || offsetLimit >= srcItem.Pos {
		//	offsetLimit++
		//	continue
		//}
		//offsetLimit++

		items[index].ID = srcItem.Id
		items[index].ClassID = srcItem.ClassId
		items[index].InstanceID = srcItem.InstanceId
		items[index].Amount = srcItem.Amount
		items[index].Pos = srcItem.Pos

		desc, err := srcInv.Descriptions.Get(srcItem.ClassId, srcItem.InstanceId)
		if err != nil {
			log.WithFields(log.Fields{
				"err":          err,
				"srcItem":      srcItem,
				"Descriptions": srcInv.Descriptions,
			}).Error("Cannot find description for item")

			index++
			continue
		}

		items[index].AppID = desc.AppId
		items[index].IconURL = desc.IconUrl
		items[index].IconURLLarge = desc.IconUrlLarge
		items[index].IconDragURL = desc.IconDragUrl
		items[index].Name = desc.Name
		items[index].MarketName = desc.MarketName
		items[index].MarketHashName = desc.MarketHashName
		items[index].NameColor = desc.NameColor
		items[index].BackgroundColor = desc.BackgroundColor
		items[index].Type = desc.Type
		items[index].Tradable = bool(desc.Tradable)
		items[index].Marketable = bool(desc.Marketable)
		items[index].Commodity = bool(desc.Commodity)
		items[index].MarketTradableRestriction = desc.MarketTradableRestriction
		items[index].Descriptions = make([]NormalizedDescription, len(desc.Descriptions))
		for i, descItem := range desc.Descriptions {
			items[index].Descriptions[i].Value = descItem.Value
			if descItem.Type != nil {
				items[index].Descriptions[i].Type = *descItem.Type
			}
			if descItem.Color != nil {
				items[index].Descriptions[i].Color = *descItem.Color
			}
		}
		items[index].Actions = make([]NormalizedAction, len(desc.Actions))
		for i, action := range desc.Actions {
			items[index].Actions[i].Name = action.Name
			items[index].Actions[i].Link = action.Link
		}
		items[index].AppData = desc.AppData
		items[index].Tags = make([]NormalizedTag, len(desc.Tags))
		for i, tag := range desc.Tags {
			items[index].Tags[i].Name = tag.InternalName
			items[index].Tags[i].Category = tag.Name
			items[index].Tags[i].CategoryName = tag.Category
			items[index].Tags[i].InternalName = tag.CategoryName
		}
		items[index].AppData = map[string]string{}
		index++
	}
	return NormalizedInventory{Items: items}
}
