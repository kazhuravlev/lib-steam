package inventory

import "testing"
import "github.com/kazhuravlev/lib-steam/meta"
import "github.com/stretchr/testify/assert"

func TestInv(t *testing.T) {
	inv, err := GetInventory(76561198261239223, 730, true, meta.LangRussian)
	assert.Nil(t, err)
	assert.NotNil(t, inv)
}
