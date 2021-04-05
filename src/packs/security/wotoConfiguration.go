package security

import (
	"packs/operations"
	"packs/players"
	"strconv"
)

// wotoConfiguration struct which will be used as
// the configuration of the players in the database.
// for security reason, this struct type cannot be accessed outside
// of the package.
//
// please consider using GetWotoConfig() in order to
// obtain a value of it.
type wotoConfiguration struct {
	UIDKeyName	[MaxUIDIndex]string
	LastUID		[MaxUIDIndex]players.UID
}

const (
	UIDIndexValue 	= "Indexes"
	UIDIndexSuffix 	= "Index_"
	UIDKeyName 		= "LastUID"

)


const (
	MaxUIDIndex = 9
	BaseUIDIndex = 1
	UIDIndexOffSet = 1
)

// _wotoConfig variable which will be used as the
// configuration of the players in the database.
// fo security reason, this value cannot be accessed outside of
// the package. please consider using GetWotoConfig() is order
// of obtain it.
var _wotoConfig *wotoConfiguration

// GetWotoConfig will create a new woto configuration if and only if
// the local value of wotoConfiguration is nil;
// otherwise, it will just return it.
func GetWotoConfig() *wotoConfiguration {
	if _wotoConfig != nil {
		return _wotoConfig
	}
	_w := wotoConfiguration{
		UIDKeyName: [MaxUIDIndex]string{},
		LastUID:    [MaxUIDIndex]players.UID{},
	}
	for i := BaseUIDIndex - UIDIndexOffSet;
				i <= MaxUIDIndex - UIDIndexOffSet; i++ {
		_w.UIDKeyName[i] = UIDIndexSuffix + strconv.Itoa(i + OffSetTokenParts)
		_w.LastUID[i] = players.GetMinimumUID(uint8(i + OffSetTokenParts))
	}
	return &_w
}

func (_w *wotoConfiguration) UpdateUIDsInServer(_index uint8) operations.RESULT {
	_w.LastUID[_index - UIDIndexOffSet].GoUp()
	return _client.UpdateLastUIDConfiguration(_index, _w)
}

func (_w *wotoConfiguration) SetUIDKeys() {
	_w.LastUID = [MaxUIDIndex]players.UID{}
	for i := BaseUIDIndex - UIDIndexOffSet;
					i <= MaxUIDIndex - UIDIndexOffSet; i++ {
		_w.UIDKeyName[i] = UIDIndexSuffix + strconv.Itoa(i + OffSetTokenParts)
	}
}
