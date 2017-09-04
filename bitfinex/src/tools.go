package src

import (
	"fmt"
	"gitlab.com/vitams/trade/bitfinex/config"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBLogin, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
}
