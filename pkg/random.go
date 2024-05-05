package pkg

import (
	"errors"
	"github.com/Masterminds/goutils"
	"github.com/google/uuid"
	"math"
	"math/rand"
)

func generateRandomValue(req generateRequest) (interface{}, error) {
	l := req.Length
	if l == 0 {
		return nil, nil
	}
	switch req.Type {
	case "int":
		return RandomInt(l)
	case "float":
		return RandomFloat(l)
	case "string":
		return RandomString(l)
	case "guid":
		return RandomGUID(l)
	case "alphanum":
		return RandomAlphaNumeric(l)
	default:
		return nil, errors.New("unknown type")
	}
}

func RandomInt(len int) (int, error) {
	if len > 19 {
		return 0, errors.New("too large length for int")
	}
	maxInt := int(math.Pow10(len)) - 1
	minInt := int(math.Pow10(len - 1))
	res := rand.Intn(maxInt-minInt+1) + minInt
	if rand.Intn(2) == 0 {
		if res != math.MinInt { // if res is -9223372036854775808 can`t convert to +int
			res *= -1
		}
	}
	return res, nil
}

func RandomFloat(len int) (float64, error) {
	maxInt := math.Pow10(len) - 1
	minInt := math.Pow10(len - 1)
	res := rand.Float64()*(maxInt-minInt+1) + minInt
	if rand.Intn(2) == 0 {
		res *= -1
	}
	return res, nil
}

func RandomGUID(len int) (interface{}, error) {
	if len < 36 {
		return uuid.Nil, errors.New("too small length for GUID (must be 36)")
	}
	if len > 36 {
		return uuid.Nil, errors.New("too large length for GUID")
	}
	res := uuid.New()
	return res, nil
}

func RandomString(len int) (string, error) {
	res, err := goutils.RandomAlphabetic(len)
	return res, err
}

func RandomAlphaNumeric(len int) (string, error) {
	res, err := goutils.RandomAlphaNumeric(len)
	return res, err
}
