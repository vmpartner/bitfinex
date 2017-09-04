package src

import (
	"strconv"
	"time"
	"gitlab.com/vitams/trade/bitfinex/config"
)

type Bitfinex struct {
	APIKey          string
	APISecret       string
	TickerChannelId float64
	Channels map[int]Channel
}

func (b *Bitfinex) GetAuth() (map[string]interface{}) {
	request := make(map[string]interface{})
	payload := "AUTH" + strconv.FormatInt(time.Now().UnixNano(), 10)[:13]
	request["event"] = "auth"
	request["apiKey"] = config.APIKey
	request["authSig"] = HexEncodeToString(GetHMAC(HashSHA512_384, []byte(payload), []byte(config.APISecret)))
	request["authPayload"] = payload
	request["authNonce"] = time.Now().UnixNano()

	return request
}
