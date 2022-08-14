package helpers

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"math/big"
	"time"

	"github.com/itchyny/base58-go"
)

func Sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))

	return algorithm.Sum(nil)
}

func Base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(encoded)
}

func GenerateUrl(expiry string, text string, time time.Time) string {
	data := expiry + text + time.String()
	byt := Sha256Of(data)
	generatedNumber := new(big.Int).SetBytes(byt).Uint64()
	url := Base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return url
}

func GenerateIfTime(timely string) (sql.NullTime, error) {
	if timely == "" {
		return sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}, nil
	}
	etime, err := time.Parse("2006-01-02", timely)
	if err != nil {
		return sql.NullTime{}, err
	}
	return sql.NullTime{
		Time:  etime,
		Valid: true,
	}, nil
}
