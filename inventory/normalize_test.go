package inventory

import "testing"
import (
	"github.com/stretchr/testify/assert"

	"bytes"
)

func TestNormalizeInventory(t *testing.T) {
	dt := "type_123"
	dc := "color_123"
	src := Inventory{
		Items: Items{
			"123": &Item{
				Id:         111111111,
				ClassId:    222222222,
				InstanceId: 333333,
				Amount:     444444444,
				Pos:        1,
			},
		},
		Descriptions: Descriptions{
			"222222222_333333": &Description{
				AppId:                     730,
				ClassId:                   222222222,
				InstanceId:                333333,
				IconUrl:                   "IconUrl_1111",
				IconUrlLarge:              "IconUrlLarge_1111",
				IconDragUrl:               "IconDragUrl_1111",
				Name:                      "Name_123",
				MarketName:                "MarketName_123",
				MarketHashName:            "MarketHashName_123",
				NameColor:                 "NameColor_123",
				BackgroundColor:           "BackgroundColor_123",
				Type:                      "Type_123",
				Tradable:                  true,
				Marketable:                true,
				Commodity:                 true,
				MarketTradableRestriction: 1,
				Descriptions: DescriptionLines{
					&DescriptionLine{
						Value: "value_123",
						Type:  &dt,
						Color: &dc,
					},
				},
				Actions: []*Action{
					{
						Name: "Name_123",
						Link: "link_123",
					},
				},
			},
		},
	}
	expected := NormalizedInventory{
		Items: []NormalizedItem{
			{
				// TODO: дописать
				AppID:                     730,
				ID:                        111111111,
				ClassID:                   222222222,
				InstanceID:                333333,
				Amount:                    444444444,
				Pos:                       1,
				IconURL:                   "IconUrl_1111",
				IconURLLarge:              "IconUrlLarge_1111",
				IconDragURL:               "IconDragUrl_1111",
				Name:                      "Name_123",
				MarketName:                "MarketName_123",
				MarketHashName:            "MarketHashName_123",
				NameColor:                 "NameColor_123",
				BackgroundColor:           "BackgroundColor_123",
				Type:                      "Type_123",
				Tradable:                  true,
				Marketable:                true,
				Commodity:                 true,
				MarketTradableRestriction: 1,
				Descriptions: []NormalizedDescription{
					{
						Value: "value_123",
						Type:  "type_123",
						Color: "color_123",
					},
				},
				Actions: []NormalizedAction{
					{
						Name: "Name_123",
						Link: "link_123",
					},
				},
				AppData: map[string]string{},
				Tags:    []NormalizedTag{},
			},
		},
	}
	result := NormalizeInventory(src, -1, -1)

	assert.Equal(t, expected, result)
}

func TestNormalizeInventoryBlank(t *testing.T) {
	src := Inventory{
		Items:        Items{},
		Descriptions: Descriptions{},
	}
	expected := NormalizedInventory{Items: []NormalizedItem{}}
	result := NormalizeInventory(src, -1, -1)

	assert.Equal(t, expected, result)
}

func TestNormalizeInventoryLimit0(t *testing.T) {
	src := Inventory{
		Items:        Items{"": &Item{}},
		Descriptions: Descriptions{"": &Description{}},
	}
	expected := NormalizedInventory{Items: []NormalizedItem{}}
	result := NormalizeInventory(src, 0, 0)

	assert.Equal(t, expected, result)
	assert.Equal(t, 0, len(result.Items))
}

func TestNormalizeInventoryOffset1(t *testing.T) {
	src := Inventory{
		Items: Items{
			"1": &Item{
				ClassId:    2,
				InstanceId: 3,
			},
			"2": &Item{
				ClassId:    2,
				InstanceId: 3,
			},
		},
		Descriptions: Descriptions{
			"2_3": &Description{},
		},
	}
	result := NormalizeInventory(src, 1, 1)

	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, uint64(2), result.Items[0].ClassID)
	assert.Equal(t, uint64(3), result.Items[0].InstanceID)
}

func TestNormalizeRealData(t *testing.T) {
	res, err := ParseInventory(bytes.NewBufferString(sample))
	assert.Nil(t, err)

	norm := NormalizeInventory(*res, -1, -1)
	assert.Equal(t, len(res.Items), len(norm.Items))
}

func BenchmarkNormalizeInventory(b *testing.B) {
	dt := "type_123"
	dc := "color_123"
	src := Inventory{
		Items: Items{
			"123": &Item{
				Id:         111111111,
				ClassId:    222222222,
				InstanceId: 333333,
				Amount:     444444444,
				Pos:        1,
			},
		},
		Descriptions: Descriptions{
			"222222222_333333": &Description{
				AppId:                     730,
				ClassId:                   222222222,
				InstanceId:                333333,
				IconUrl:                   "IconUrl_1111",
				IconUrlLarge:              "IconUrlLarge_1111",
				IconDragUrl:               "IconDragUrl_1111",
				Name:                      "Name_123",
				MarketName:                "MarketName_123",
				MarketHashName:            "MarketHashName_123",
				NameColor:                 "NameColor_123",
				BackgroundColor:           "BackgroundColor_123",
				Type:                      "Type_123",
				Tradable:                  true,
				Marketable:                true,
				Commodity:                 true,
				MarketTradableRestriction: 1,
				Descriptions: DescriptionLines{
					&DescriptionLine{
						Value: "value_123",
						Type:  &dt,
						Color: &dc,
					},
				},
				Actions: []*Action{
					{
						Name: "Name_123",
						Link: "link_123",
					},
				},
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NormalizeInventory(src, -1, -1)
	}
}
