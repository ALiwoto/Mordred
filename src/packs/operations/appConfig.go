package operations

import (
	"log"
	"os"
	"packs/errors"
	"packs/security"
)

type MordredConfig struct {
	dConfig 	*DataBaseConfig
	sConfig 	*ServerConfig
	_full		*string
	_fullBytes	[]byte
	_serialized	*WotoSerialized
}

type DataBaseConfig struct {

}

type ServerConfig struct {

}

const ConfigFileName			= "config.ini"


func GetMordredConfig() (*MordredConfig, RESULT){
	_mConfig := MordredConfig{
		dConfig: nil,
		sConfig: nil,
	}
	_r := _mConfig._setValues()
	if _r != SUCCESS {
		return nil, FAILED
	}
	return &_mConfig, _r
}

func (_m *MordredConfig) _setValues() RESULT {
	_info, _err := os.Stat(ConfigFileName)
	if _err != nil {
		log.Fatal(errors.ConfigFileError, _err)
		return FAILED
	}
	_m._fullBytes, _err = os.ReadFile(_info.Name())
	_value := string(_m._fullBytes)
	if security.IsEmpty(&_value) {
		log.Fatal(errors.ConfigFileError)
	}
	_m._full = &_value
	_ser, _result := Deserialize(_m._full)
	if _result != SUCCESS {
		return _result
	}
	_m._serialized = _ser
	_result = _m._setDBConfig()
	if _result != SUCCESS {
		return _result
	}
	_result = _m._setServerConfig()
	return _result
}

func (_m *MordredConfig) _setDBConfig() RESULT {

	return SUCCESS
}

func (_m *MordredConfig) _setServerConfig() RESULT {
	return SUCCESS
}

