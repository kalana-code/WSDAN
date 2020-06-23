package config

import (
	"errors"
	"log"
	"sync"
)

// Settings is used to store shared data
type Settings struct {
	ControllerMAC string `json:"ControllerMAC"`
	ControllerIP  string `json:"ControllerIP"`
}

var (
	once     sync.Once
	instance *Settings
	infoLog  string = "INFO: [CF]:"
	errorLog string = "ERROR: [CF]:"
)

// GetSettings is used to get settings
func GetSettings() *Settings {
	log.Println(infoLog, "Invoke GetSettings")
	once.Do(func() {
		instance = new(Settings)
	})
	return instance
}

// SetControllerMAC is used to set controller MAC
func (obj *Settings) SetControllerMAC(mac string) error {
	if instance != nil {
		instance.ControllerMAC = mac
		return nil
	}
	return errors.New("Settings are not initiated")
}

// GetControllerMAC is used to get controller MAC
func (obj *Settings) GetControllerMAC() (string, error) {
	if instance != nil {
		return instance.ControllerMAC, nil
	}
	return "", errors.New("Settings are not initiated")
}
