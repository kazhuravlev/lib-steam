package deviceid

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// GenerateDeviceID генерируем новый идентфиикатор устройства. Используется для дальнейшем подстановки в параметры бота.
// Например, после генерации этого идентфикатора, боту необъодимо передавать этот идентфиикатор в кач-ве deviceID
// (смотри спецификация ботов). Полный идентификатор устройства должен выглядеть примерно так:
// "android:d2305bb9-a690-e199-9383-017f6709cc8e". Страые версии ботов использовали рандомно сгенерированый
// идентификатор устройства, который генерировал владелец ботов при их авторизации. Этот алгоритм - новая версия
// генератора
func GenerateDeviceID(steamID string) string {
	h := sha1.New()
	io.WriteString(h, steamID)
	s := hex.EncodeToString(h.Sum(nil))
	return s[0:8] + "-" + s[8:12] + "-" + s[12:16] + "-" + s[16:20] + "-" + s[20:32]
}
