package db

import (
	"Beq/settings/model"
	"errors"
	"sync"
)

//SettingDB syatem setting
type SettingDB model.Settings

var instance SettingDB

var once sync.Once

//GetSystemSetting Initiating rule database
func GetSystemSetting() *SettingDB {
	once.Do(func() {
		instance = SettingDB{}
		instance.Automation = false
		instance.ForceDispurse = false
	})
	return &instance
}

//GetSetting get current setting of system
func (*SettingDB) GetSetting() (*map[string]interface{}, error) {
	if &instance != nil {
		setting := make(map[string]interface{})
		setting["Automation"] = instance.Automation
		setting["ForceDispursed"] = instance.ForceDispurse
		setting["IP"] = instance.IPAddress
		setting["MAC"] = instance.MacAddress

		return &setting, nil
	}
	return nil, errors.New("No Data Base Initiate")
}

//SetMACandIP used for adding system MAC and IP addresses
func (*SettingDB) SetMACandIP(MAC string, IP string) error {
	if &instance != nil {
		instance.MacAddress = MAC
		instance.IPAddress = IP
		return nil
	}
	return errors.New("No Data Base Initiate")
}

//GetMAC used for adding system MAC and IP addresses
func (*SettingDB) GetMAC() (string, error) {
	if &instance != nil {
		return instance.MacAddress, nil
	}
	return "", errors.New("No Data Base Initiate")
}

//ToggleMode used for change system mode
func (*SettingDB) ToggleMode() error {
	if &instance != nil {
		instance.Automation = !instance.Automation
		return nil
	}
	return errors.New("No Data Base Initiate")
}

//ToggleForceDispurserMode used for change system mode
func (*SettingDB) ToggleForceDispurserMode() error {
	if &instance != nil {
		instance.ForceDispurse = !instance.ForceDispurse
		return nil
	}
	return errors.New("No Data Base Initiate")
}

//IsForceDisposed used for forced dispursed  setting state
func (obj *SettingDB) IsForceDisposed() (bool, error) {
	if &instance != nil {
		return obj.ForceDispurse, nil
	}
	return false, errors.New("No Data Base Initiate")
}
