package security


import (
	"math/rand"
	"strconv"
	"strings"
)

const FixingStringValue = "0"

const (
	MaxLengthTokenPart 		= 4
	MaxIntTokenPart   		= 1000
	MaxIntFirstTokenNumPart = 9
	OffSetTokenParts 		= 1
	MaxTokenParts 			= 3
	BaseTokenPart 			= 0
	BaseConverting 			= 10
)

type wotoToken string

func GenerateNewWotoToken() *wotoToken {
	//rand.Seed(time.Now().UnixNano())
	total := EMPTY
	_rand := (rand.Intn(MaxIntFirstTokenNumPart) % MaxIntFirstTokenNumPart) + OffSetTokenParts
	total += strconv.Itoa(_rand)
	for i := BaseTokenPart; i < MaxTokenParts; i++ {
		_rand = (rand.Intn(MaxIntTokenPart) % MaxIntTokenPart) + OffSetTokenParts
		total += FixNum(int64(_rand))
	}
	_w := wotoToken(total)
	return &_w
}
func FixNum(_num int64) string {
	_str := strconv.FormatInt(_num, BaseConverting)
	if len(_str) < MaxLengthTokenPart {
		_offSet := MaxLengthTokenPart - len(_str)
		_fixingValue := strings.Repeat(FixingStringValue, _offSet)
		return _fixingValue + _str
	}
	return _str
}
func FixCustomNum(_num , _max int) string {
	_str := strconv.FormatInt(int64(_num), BaseConverting)
	if len(_str) < _max {
		_offSet := _max - len(_str)
		_fixingValue := strings.Repeat(FixingStringValue, _offSet)
		return _fixingValue + _str
	}
	return _str
}
func (_t *wotoToken) IsEqual(_w *wotoToken) bool {
	return string(*_t) == string(*_w)
}