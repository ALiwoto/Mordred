package operations

import (
	"packs/security"
	"strings"
)

type WotoSerialized interface {
	GetValue(_s string)			(string, bool)
	GetSubTree(_s string)		(*SubTree, bool)
	GetObjectType(_s string)	ObjectType
	EncodingType()				EncodingType
	HasLicense()				bool
	IsValue(_s string)			bool
	IsSubtree(_s string)		bool
	Serialize()					string
}


type (
	EncodingType	uint8
	ObjectType		uint8
)

type SerializeVersion string

type wotoSerialization struct {
	_full		*string
	_version	*SerializeVersion
}

const (
	NotSet	EncodingType = 0
	UTF8	EncodingType = 1
)

const (
	NoneObj		ObjectType = 0
	SubtreeObj	ObjectType = 1
	ValueObj	ObjectType = 2
)
const (
	SerializeKeyValue	= "woto-serialization"
	VersionKeyValue		= "version"
	EncodingKeyValue	= "encoding"
)

// the indexed constant values.
const (
	BaseIndex		= 0
	BaseOneIndex	= 1
	InfoHeaderIndex = BaseIndex
)

// Unlimited Char Works. (UCW)
const (
	CommentChar		rune	= '#'
	HeaderOpenChar	rune	= '<'
	HeaderCloseChar rune	= '>'
	SpaceChar 		rune	= ' '
	EqualChar		rune	= '='
	StringChar		rune	= '"'
	LineCharSeparator		= "\r\n"
)


func Deserialize(_s *string) (*WotoSerialized, RESULT) {
	if *_s == security.EMPTY {
		return nil, CANCELED
	}
	_w := wotoSerialization{}
	_w._full = _s
	_myStrings := strings.Split(*_s, LineCharSeparator)
	_myStrings = security.FixSplit(_myStrings)
	// for calculating real index, we have ignored
	// comments and empty spaces.
	// that's why its name is real index.
	var _readIndex = 0
	for _, _myString := range _myStrings {
		// check if the real index is 0 or not.
		// in the real index of zero, we have to read the header.
		if _readIndex == InfoHeaderIndex {
			// check if the first line is comment or not.
			// if it's comment, it's okay to continue.
			if isComment(&_myString) {
				continue
			}
			_result := _w._setHeader(&_myString)
			if _result != SUCCESS {
				return nil, _result
			}

		}
		_readIndex++
	}
	return nil, FAILED
}

func isComment(_s *string) bool {
	return rune((*_s)[BaseIndex]) == CommentChar
}

func isHeader(_s *string) bool {
	var b1, b2 bool
	for _i, _real := BaseIndex, BaseIndex; _i < len(*_s); i++ {
		if rune((*_s)[_i]) == SpaceChar {
			continue
		}
		if _real == BaseIndex {
			b1 = rune((*_s)[_i]) == HeaderOpenChar
			_real++
			continue
		}
		// when you reach this point, it means the _real
		// SHOULD be 1.
		// why do this?? because we want to be sure that _real is 1.
		// if the _real is 1, it means that the open header is already
		// checked, so we can check the close header char.
		if _i == len(*_s) - _real {
			b2 = rune((*_s)[_i]) == HeaderCloseChar
			if b2 {
				b2 = strings.Contains(*_s, SerializeKeyValue)
			}
			break
		}
	}
	return b1 && b2
}

func (_w *wotoSerialization) _setHeader(_s *string) RESULT {
	if _s == nil {
		return FAILED
	}
	// check if the string is in right format or not.
	if !isComment(_s) {
		return FAILED
	}
	var(
		_myString	string
		_inStrings 	[]string
		_myStrings	[]string
	)
	_myString = strings.TrimLeft(*_s, security.EMPTY)
	_myString = strings.TrimRight(_myString, security.EMPTY)
	_myStrings = security.Split(_myString, string(HeaderOpenChar),
		SerializeKeyValue, string(HeaderCloseChar), string(SpaceChar))
	for _, _current := range _myStrings {
		if strings.Contains(_current, VersionKeyValue) {
			_inStrings = security.Split(_current, VersionKeyValue,
				string(EqualChar), string(StringChar), string(SpaceChar))
			if len(_inStrings) > BaseOneIndex {
				return FAILED
			}
			_w._setVersion(_inStrings[BaseIndex])
		} else if strings.Contains(_current, EncodingKeyValue){
			// TODO: implement another values in the serialization.
		}
	}
}

func (_w *wotoSerialization) _setVersion(_s string) {
	_v := SerializeVersion(_s)
	_w._version = &_v
}

func (_w wotoSerialization) GetValue(_s string) (string, bool) {

	return security.EMPTY, false
}

func (_w wotoSerialization) HasLicense() bool {

	return false
}

func (_w wotoSerialization) EncodingType() EncodingType {


	return NotSet
}

func (_w wotoSerialization) GetObjectType(_s string) ObjectType {

	return 0
}

func (_w wotoSerialization) IsValue(_s string) bool {

	return false
}

func (_w wotoSerialization) IsSubtree(_s string) bool {

	return false
}

func (_w wotoSerialization) Serialize() string {

	return security.EMPTY
}

func (_w wotoSerialization) GetSubTree(_s string) (*SubTree, bool) {

	return nil, false
}
