package api

func parseAssetInfo(assetInterface interface{}) (AssetClassInfo, error) {
	asset := AssetClassInfo{}
	x, ok := assetInterface.(map[string]interface{})
	if !ok {
		return asset, ErrBadResponse
	}

	sIconURL, err := stringVal(x, "icon_url")
	if err != nil {
		return asset, err
	}
	asset.IconURL = sIconURL

	sIconURLLarge, err := stringVal(x, "icon_url_large")
	if err != nil {
		return asset, err
	}
	asset.IconURLLarge = sIconURLLarge

	sMarketHashName, err := stringVal(x, "market_hash_name")
	if err != nil {
		return asset, err
	}
	asset.MarketHashName = sMarketHashName

	sMarketName, err := stringVal(x, "market_name")
	if err != nil {
		return asset, err
	}
	asset.MarketName = sMarketName

	sMarketable, err := stringVal(x, "marketable")
	if err != nil {
		return asset, err
	}
	asset.Marketable = sMarketable

	sTradable, err := stringVal(x, "tradable")
	if err != nil {
		return asset, err
	}
	asset.Tradable = sTradable

	sCommodity, err := stringVal(x, "commodity")
	if err != nil {
		return asset, err
	}
	asset.Commodity = sCommodity

	sName, err := stringVal(x, "name")
	if err != nil {
		return asset, err
	}
	asset.Name = sName

	sType, err := stringVal(x, "type")
	if err != nil {
		return asset, err
	}
	asset.Type = sType

	sClassid, err := stringVal(x, "classid")
	if err != nil {
		return asset, err
	}
	asset.Classid = sClassid

	sTags, err := mapVal(x, "tags")
	if err != nil {
		return asset, err
	}
	tags := map[string]Tag{}
	for key, iVal := range sTags {
		tag, err := parseTag(iVal)
		if err != nil {
			return asset, err
		}
		tags[key] = tag
	}
	asset.Tags = tags

	sDescriptions, err := mapVal(x, "descriptions")
	if err != nil {
		return asset, err
	}
	descriptions := map[string]Description{}
	for key, iVal := range sDescriptions {
		description, err := parseDescription(iVal)
		if err != nil {
			return asset, err
		}
		descriptions[key] = description
	}
	asset.Descriptions = descriptions

	return asset, nil
}

func parseTag(iTag interface{}) (Tag, error) {
	tag := Tag{}
	mTag, ok := iTag.(map[string]interface{})
	if !ok {
		return tag, ErrBadRequest
	}

	sCategory, err := stringVal(mTag, "category")
	if err != nil {
		return tag, err
	}
	tag.Category = sCategory

	sCategoryName, err := stringVal(mTag, "category_name")
	if err != nil {
		return tag, err
	}
	tag.CategoryName = sCategoryName

	sInternalName, err := stringVal(mTag, "internal_name")
	if err != nil {
		return tag, err
	}
	tag.InternalName = sInternalName

	sName, err := stringVal(mTag, "name")
	if err != nil {
		return tag, err
	}
	tag.Name = sName

	return tag, nil
}

func parseDescription(iDescription interface{}) (Description, error) {
	description := Description{}
	mDescription, ok := iDescription.(map[string]interface{})
	if !ok {
		return description, ErrBadRequest
	}

	sColor, err := stringVal(mDescription, "color")
	if err != nil {
		description.Color = ""
	} else {
		description.Color = sColor
	}

	sType, err := stringVal(mDescription, "type")
	if err != nil {
		return description, err
	}
	description.Type = sType

	sValue, err := stringVal(mDescription, "value")
	if err != nil {
		return description, err
	}
	description.Value = sValue

	sAppData, err := mapVal(mDescription, "app_data")
	if err == nil {
		sDefindex, err := stringVal(sAppData, "def_index")
		if err == nil {
			description.AppData.DefInfex = sDefindex
		}

		sIsItemsetName, err := stringVal(sAppData, "is_itemset_name")
		if err == nil {
			description.AppData.IsItemsetName = sIsItemsetName
		}
	}

	return description, nil
}

func stringVal(i map[string]interface{}, key string) (string, error) {
	iVal, exists := i[key]
	if !exists {
		return "", ErrBadResponse
	}

	sVal, ok := iVal.(string)
	if !ok {
		return "", ErrBadResponse
	}

	return sVal, nil
}

func mapVal(i map[string]interface{}, key string) (map[string]interface{}, error) {
	iVal, exists := i[key]
	if !exists {
		return nil, ErrBadResponse
	}

	sVal, ok := iVal.(map[string]interface{})
	if !ok {
		return nil, ErrBadResponse
	}

	return sVal, nil
}
