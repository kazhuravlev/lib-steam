package steamid

import (
	"errors"
	"strconv"
)

const (
	// CommunityIDLength кол-во знаков в идентификтаоре стим айди
	CommunityIDLength = 17
)

var (
	nullID = SteamID(0)
	// ErrBadSteamID - некорректный идентфиикатор стим
	ErrBadSteamID = errors.New("Bad/Invalid SteamID")
)

// SteamID is usually returned and used as a string, this just insures consistant handling
type SteamID uint64

// NewFromCommunityIDString новый объект идентификатора на основе строки -
// идентификатора пользователя в коммьюнити
func NewFromCommunityIDString(communityID string) (SteamID, error) {
	if len(communityID) != 17 {
		return nullID, ErrBadSteamID
	}

	id, err := strconv.ParseUint(communityID, 10, 64)
	if err != nil {
		return nullID, ErrBadSteamID
	}

	return SteamID(id), nil
}

// NewFromCommunityIDuint новый объект идентификатора на основе uint числа -
// идентификатора пользователя в коммьюнити
func NewFromCommunityIDuint(communityID uint) (SteamID, error) {
	if len(strconv.FormatUint(uint64(communityID), 10)) != 17 {
		return nullID, ErrBadSteamID
	}

	return SteamID(communityID), nil
}

// NewFromAccountIDuint новый объект идентификатора на основе uint числа -
// идентификатора пользователя в steam
func NewFromAccountIDuint(accountID uint) (SteamID, error) {
	return NewFromAccountIDString(strconv.FormatUint(uint64(accountID), 10))
}

// NewFromAccountIDString новый объект идентификатора на основе строки -
// идентификатора аккаунта стим
func NewFromAccountIDString(accountID string) (SteamID, error) {
	accId, err := strconv.ParseUint(accountID, 10, 32)
	if err != nil {
		return nullID, ErrBadSteamID
	}

	instance := uint32(1)
	universe := int32(1)
	accountType := int32(1)

	ss := uint64(0)
	set(&ss, 0, 0xFFFFFFFF, uint64(accId))
	set(&ss, 32, 0xFFFFF, uint64(instance))
	set(&ss, 56, 0xF, uint64(universe))
	set(&ss, 52, 0xF, uint64(accountType))

	return SteamID(ss), nil
}

func (s SteamID) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

func (s SteamID) GetAccountId() uint32 {
	return uint32(s.get(0, 0xFFFFFFFF))
}

func (s SteamID) GetAccountIdString() string {
	return strconv.FormatUint(uint64(s.get(0, 0xFFFFFFFF)), 10)
}

func (s *SteamID) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.String() + "\""), nil
}

func (s *SteamID) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		b = b[1 : len(b)-1]
	}
	v, err := strconv.ParseUint(string(b), 10, 64)

	*s = SteamID(v)
	return err
}

func (s SteamID) get(offset uint, mask uint64) uint64 {
	return (uint64(s) >> offset) & mask
}

func set(i *uint64, offset uint, mask, value uint64) {
	*i = (*i & ^(mask << offset)) | (value&mask)<<offset
}
