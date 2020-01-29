package cache

import (
	"errors"
	"sync"
)

type IOTDevice struct {
	ID        string `json:"id"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

var defaultCache sync.Map

func Set(d *IOTDevice) error {
	if d == nil {
		return errors.New("nil")
	}

	defaultCache.Store(d.ID, d)
	return nil
}

func Get(id string) *IOTDevice {
	v, ok := defaultCache.Load(id)
	if !ok {
		return nil
	}

	return v.(*IOTDevice)
}

func GetAll() (devices []*IOTDevice) {
	defaultCache.Range(func(key, value interface{}) bool {
		devices = append(devices, value.(*IOTDevice))
		return true
	})

	return
}
