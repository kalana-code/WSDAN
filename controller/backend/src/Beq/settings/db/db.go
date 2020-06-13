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
	})
	return &instance
}

//GetSetting get current setting of system
func (*SettingDB) GetSetting() (*map[string]interface{}, error) {
	if &instance != nil {
		setting := make(map[string]interface{})
		if setting["Mode"] = "MANUAL"; instance.Automation {
			setting["Mode"] = "AUTO"
		}
		return &setting, nil
	}
	return nil, errors.New("No Data Base Initiate")
}

//ToggleMode used for change system mode
func (*SettingDB) ToggleMode() error {
	if &instance != nil {
		instance.Automation = !instance.Automation
		return nil
	}
	return errors.New("No Data Base Initiate")
}
