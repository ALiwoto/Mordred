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
	/*
	 * we don't need to create a new seed for random numbers,
	 * using rand.Seed(time.Now().UnixNano()) function.
	 * I suggest don't use it anyway, because rand.Intn() function
	 * will do its work perfectly!
	 */
	total := EMPTY
	_rand := (rand.Intn(MaxIntFirstTokenNumPart) % MaxIntFirstTokenNumPart) + OffSetTokenParts
	total += strconv.Itoa(_rand)
	for i := BaseTokenPart; i < MaxTokenParts; i++ {
		_rand = (rand.Intn(MaxIntTokenPart) % MaxIntTokenPart) + OffSetTokenParts
		total += fixNum(int64(_rand))
	}
	_w := wotoToken(total)
	return &_w
}
func fixNum(_num int64) string {
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