package inventory

import (
	"io"
	"io/ioutil"

	"github.com/pquerna/ffjson/ffjson"
)

// ParseInventory парсит инвентарь в свои структуры. может использоваться
// снаружи при получении инвентаря из другого источника
func ParseInventory(body io.Reader) (*Inventory, error) {
	inv := new(Inventory)
	bb, _ := ioutil.ReadAll(body)
	if err := ffjson.NewDecoder().Decode(bb, inv); err != nil {
		return nil, err
	}
	return inv, nil
}
