package models

import "time"

type LocalCacheStruct struct {
	LocalCacheData []byte
	ExpiryTime     map[string]interface{}
}

type ExpiryTime struct {
	Key  string
	Time time.Time
}

type KeyValueStruct struct {
	Key         string
	Value       interface{}
	TimeSeconds int
}
