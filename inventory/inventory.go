package inventory

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/franela/goreq"
	"github.com/kazhuravlev/lib-steam/meta"
)

var (
	// ErrSteamError ошибка steam. На данный момент любой ответ считается
	// ошибкой если вернул статус, отличный от 200
	ErrSteamError = errors.New("steam error")
)

// GetInventory получает инвентарь указанного пользователя по указанной игре
func GetInventory(steamCommunityID int, appID meta.AppID, tradable bool, language meta.Lang) (*Inventory, error) {
	query := url.Values{}
	//	query.Set("donotcache", time.Now().Unix())
	if tradable {
		query.Set("trading", "1")
	}

	r := goreq.Request{
		Method:      "GET",
		Uri:         fmt.Sprintf("http://steamcommunity.com/profiles/%d/inventory/json/%d/2", steamCommunityID, meta.Apps[appID]),
		QueryString: query,
	}.WithCookie(&http.Cookie{Name: "Steam_Language", Value: meta.Langs[language]})
	resp, err := r.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrSteamError
	}
	inv, err := ParseInventory(resp.Body)
	if err != nil {
		return nil, err
	}

	return inv, nil
}

// BuildRequest создает объект запроса к steam для получения информации об
// инвентаре пользователя. Может использоваться без функции GetIngentory для
// работы через пул прокси или выполнения запроса другими способами
func BuildRequest(steamCommunityID int, appID meta.AppID, tradable bool, language meta.Lang) (*http.Request, error) {
	query := url.Values{}
	//	query.Set("donotcache", time.Now().Unix())
	if tradable {
		query.Set("trading", "1")
	}
	urlStr := fmt.Sprintf("http://steamcommunity.com/profiles/%d/inventory/json/%d/2?%s", steamCommunityID, meta.Apps[appID], query.Encode())

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "Steam_Language", Value: meta.Langs[language]})
	return req, nil
}
