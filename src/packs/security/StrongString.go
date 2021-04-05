package security

import (
	"strconv"
	"strings"
)

const (
	// the LineSeparator used in separating the byte values in the
	// StrongString.
	LineSeparator = "-"
	// EMPTY is an empty string.
	EMPTY = ""
)

const (
	// the StrongBase is Hexadecimal, which is base 16.
	StrongBase 		= 0x010
	// the StrongOffSet is the offset value of each character in the
	// strong string.
	StrongOffSet 	= 0x00D
)

// the StrongString used in the program for High-security!
type StrongString struct {
	_value []rune
}

// IsEmpty function will check if the passed-by
// string value is empty or not.
func IsEmpty(_s *string) bool {
	return *_s == EMPTY
}

// generateStrongString will generate a new StrongString
// with the specified non-encoded string value.
func GenerateStrongString(_s string) StrongString {
	_strong := StrongString{}
	_strong._setValue(_s)
	return _strong
}

// getStrong will give you an already encoded StrongString
// which provided by method _getString() method and saved in the
// data base.
func GetStrong(_s string) StrongString {
	_strong := StrongString{}
	myBytes := toByteArrayS(_s, LineSeparator)
	_strong._setValueByBytes(myBytes)
	return _strong
}


// convert the specified string array (encoded) to byte array.
func _toByteArray(myStrings []string) []rune{
	myBytes := make([]rune, 0)
	for _, _current := range myStrings {
		_myInt, _ := strconv.Atoi(_current)
		myBytes = append(myBytes, rune(_myInt + StrongOffSet))
	}
	return myBytes
}

// convert the specified string (encoded) to byte array.
func toByteArrayS(theString string, separator string) []rune {
	return _toByteArray(FixSplit(strings.Split(theString, separator)))
}

// convertToBytes will convert an ordinaryString (non-encoded) to byte array.
func convertToBytes(ordinaryString string) []rune {
	_runes := []rune(ordinaryString)
	finalRune := make([]rune, len(_runes), cap(_runes))
	for _i, _current := range _runes {
		finalRune[_i] = _current + StrongOffSet
	}
	return finalRune
}

// ConvertToString will convert the specified byte arrays to string.
func ConvertToString(_b []rune) string {
	_total := EMPTY
	for _, _current := range _b {
		_total += string(_current - StrongOffSet)
	}
	return _total
}

// FixSplit will fix the bullshit bug in the
// Split function (which is not ignoring the spaces between strings).
func FixSplit(_myStrings []string) []string{
	final := make([]string, 0, cap(_myStrings))
	for _, _current := range _myStrings {
		if !IsEmpty(&_current) {
			final = append(final, _current)
		}
	}
	return final
}

// _getString will give you an encoded string with the High-security
// level which you should use it in the database.
func (_s *StrongString) _getString() string {
	var _current int
	var total = EMPTY
	for _, b := range _s._value{
		_current = int(b)
		total += strconv.Itoa(_current)
		total += LineSeparator
	}
	return total
}

// _setValue will set the bytes value of the StrongString.
func (_s *StrongString) _setValue(theString string) {
	_s._setValueByBytes(convertToBytes(theString))
}

// _setValueByBytes will set the bytes value directly.
func (_s *StrongString) _setValueByBytes(_b []rune) {
	_s._value = _b
}

// GetValue will give you the real value of this StrongString.
func (_s *StrongString) GetValue() *string {
	realString := ConvertToString(_s._value)
	return &realString
}

// length method, will give you the length-as-int of this StrongString.
func (_s *StrongString) Length() int {
	return len(*_s.GetValue())
}

// isEmpty will check if this StrongString is empty or not.
func (_s *StrongString) IsEmpty() bool {
	return IsEmpty(_s.GetValue())
}

// isEqual will check if the passed-by-value in the arg is equal to this
// StrongString or not.
func (_s *StrongString) IsEqual(_strong *StrongString) bool {
	// check if the length of them are equal or not.
	if len(_s._value) != len(_strong._value) {
		//fmt.Println(len(_s._value), len(_strong._value))
		return false
	}
	for i := 0; i < len(_s._value); i++ {
		if _s._value[i] != _strong._value[i] {
			//fmt.Println(_s._value[i], _strong._value[i])
			return false
		}
	}
	return true
}
