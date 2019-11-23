package steamid

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSteamID(t *testing.T) {
	jsondata := `{"steam_id": "76561198263585543"}`
	jsonout := struct {
		SteamID SteamID `json:"steam_id"`
	}{}
	err := json.Unmarshal([]byte(jsondata), &jsonout)
	assert.Nil(t, err)

	assert.Equal(t, SteamID(76561198263585543), jsonout.SteamID)
}

func TestGetAccountId(t *testing.T) {
	table := []struct {
		src   uint64
		valid uint32
	}{
		{src: 76561198073818912, valid: 113553184},
		{src: 76561198086807182, valid: 126541454},
		{src: 76561198095754838, valid: 135489110},
	}

	for _, row := range table {
		s := SteamID(row.src)
		accountID := s.GetAccountId()
		assert.Equal(t, row.valid, accountID)
	}
}

func TestGetAccountIdString(t *testing.T) {
	table := []struct {
		src   uint64
		valid string
	}{
		{src: 76561198073818912, valid: "113553184"},
		{src: 76561198086807182, valid: "126541454"},
		{src: 76561198044505631, valid: "84239903"},
	}

	for _, row := range table {
		s := SteamID(row.src)
		accountID := s.GetAccountIdString()
		assert.Equal(t, row.valid, accountID)
	}
}

func TestString(t *testing.T) {
	s := SteamID(76561198073818912)

	assert.Equal(t, "76561198073818912", s.String())
}

func TestNewFromAccountIDString(t *testing.T) {
	table := []struct {
		src   string
		valid uint64
	}{
		{src: "113553184", valid: 76561198073818912},
		{src: "126541454", valid: 76561198086807182},
		{src: "84239903", valid: 76561198044505631},
	}

	for _, row := range table {
		res, err := NewFromAccountIDString(row.src)
		assert.Nil(t, err)
		assert.Equal(t, uint(row.valid), uint(res))
	}
}

func TestNewFromCommunityIDString(t *testing.T) {
	table := []struct {
		src   string
		valid uint64
	}{
		{src: "76561198261239223", valid: 76561198261239223},
		{src: "76561198073818912", valid: 76561198073818912},
	}

	for _, row := range table {
		res, err := NewFromCommunityIDString(row.src)
		assert.Nil(t, err)
		assert.Equal(t, uint(row.valid), uint(res))
	}
}

func TestFailNewFromCommunityIDString(t *testing.T) {
	table := []string{
		"-1111111111111111",
		"123",
		"123asdasd213123123123123123123123123",
		"asd",
		"sssdddsssssssssss",
	}

	for _, row := range table {
		_, err := NewFromCommunityIDString(row)
		assert.Error(t, err)
	}
}

func TestNewFromCommunityIDuint(t *testing.T) {
	table := []struct {
		src   uint
		valid uint
	}{
		{src: 76561198261239223, valid: 76561198261239223},
		{src: 76561198073818912, valid: 76561198073818912},
	}

	for _, row := range table {
		res, err := NewFromCommunityIDuint(row.src)
		assert.Nil(t, err)
		assert.Equal(t, row.valid, uint(res))
	}
}

func TestFailNewFromCommunityIDuint(t *testing.T) {
	table := []uint{
		11111111111,
		123,
	}

	for _, row := range table {
		_, err := NewFromCommunityIDuint(row)
		assert.Error(t, err)
	}
}
