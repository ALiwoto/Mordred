package Players

import (
	"math/rand"
	"packs"
	"packs/security"
	"strconv"
	"time"
)

// UID type is used to count the players.
// originally it means: User ID.
// it will be a mandatory value is PlayerInfo struct.
type UID string

// MinimumUID represents the possible minimum UID ever.
// please consider that this UID is not valid alone,
// thus you should add an index to it in order to use it.
// so please don't use it directly; consider using GetMinimumUID()
// function in order to obtain a valid minimum UID value.
const MinimumUID = "00000000"
// UIDLength is the valid length of a UID.
const UIDLength  = 9

// getMinimumUID will give you the minimum UID
// with the specified non-zero-based index.
func GetMinimumUID(_index uint8) UID {
	_s := strconv.Itoa(int(_index))
	return UID(_s + MinimumUID)
}

// getUID will give you a random-indexed
// UID which is the last-uid of that specific index.
// WARNING:	before calling this method, be sure that
// the _client and _wotoConfig values are defined and are not nil.
func getUID() (packs.RESULT, *UID){
	rand.Seed(time.Now().UnixNano())
	_w := security.GetWotoConfig()
	_index := (rand.Intn(security.MaxUIDIndex) % security.MaxUIDIndex) + 1
	_last := _w.LastUID[_index - security.UIDIndexOffSet]
	_re := _w.UpdateUIDsInServer(uint8(_index))
	return _re, &_last
}

// IsValidUID function will check if an UID is valid or not.
func (_uid *UID) IsValidUID() bool {
	return _uid.Length() == UIDLength
}

func (_uid *UID) IsEqual(uid *UID) bool {
	_s1 := string(*_uid)
	_s2 := string(*uid)
	return _s1 == _s2
}

func (_uid *UID) Length() int {
	_s := string(*_uid)
	return len(_s)
}

func (_uid *UID) GoUp() {
	_s := string(*_uid)
	_index := _s[:security.UIDIndexOffSet]
	_sArray := _s[security.UIDIndexOffSet:]
	_sInt, _err := strconv.Atoi(_sArray)
	if _err != nil {
		panic(_err)
	}
	_sInt++
	_sArray = security.FixCustomNum(_sInt, len(MinimumUID))
	*_uid = UID(_index + _sArray)
}
