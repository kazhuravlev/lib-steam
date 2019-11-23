package tradeurl

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/kazhuravlev/lib-steam/steamid"
)

var (
	ErrBadURL = errors.New("bad url")
)

func Build(steamID steamid.SteamID, tradeToken string) (string, error) {
	return fmt.Sprintf("https://steamcommunity.com/tradeoffer/new/?partner=%s&token=%s", steamID.GetAccountIdString(), tradeToken), nil
}

// Parse парсит переданный tradeURL. возвращает идентификатор пользователя и токен
func Parse(tradeURL string) (string, string, error) {
	u, err := url.Parse(tradeURL)
	if err != nil {
		return "", "", err
	}

	partner := u.Query().Get("partner")
	token := u.Query().Get("token")
	if partner == "" || token == "" {
		return "", "", ErrBadURL
	}

	return partner, token, nil
}

// CheckToken проверяет токен на корректность
func CheckToken(token string) bool {
	if len(token) != 8 {
		return false
	}

	return true
}
