package rsa

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/kazhuravlev/lib-steam/steamid"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	t.Time, err = time.Parse(time.RFC3339, s)
	if _, ok := err.(*time.ParseError); ok {
		// ok, that didn't work lets try parsing an int
		var i int64
		if i, err = strconv.ParseInt(s, 10, 64); err != nil {
			return err
		}
		t.Time = time.Unix(i, 0)
	}
	return err
}

func (t *Timestamp) String() string {
	return strconv.FormatInt(t.Unix(), 10)
}

type RSAResponse struct {
	Success   bool `json:"success"`
	PublicKey *rsa.PublicKey
	Timestamp Timestamp       `json:"timestamp"`
	TokenGID  string          `json:"token_gid"`
	SteamID   steamid.SteamID `json:"steamid"`
}

func (r *RSAResponse) UnmarshalJSON(b []byte) error {
	type localRSAResponse RSAResponse
	localData := struct {
		localRSAResponse
		Modulus  string `json:"publickey_mod"`
		Exponent string `json:"publickey_exp"`
	}{}

	err := json.Unmarshal(b, &localData)
	if err != nil {
		return err
	}
	*r = RSAResponse(localData.localRSAResponse)

	// Absolutly no point progressing beyond this point if there is no success
	if !r.Success {
		return nil
	}

	exponent, err := strconv.ParseInt(localData.Exponent, 16, 0)
	modulus := big.Int{}
	if _, ok := modulus.SetString(localData.Modulus, 16); !ok {
		return errors.New("invalid modulus")
	}

	r.PublicKey = &rsa.PublicKey{
		N: &modulus,
		E: int(exponent),
	}

	return nil
}
