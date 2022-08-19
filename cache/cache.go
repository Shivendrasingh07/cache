package cache

import (
	"encoding/json"
	"example.com/m/models"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var LocalCachedData models.LocalCacheStruct

func Set(data models.KeyValueStruct) error {

	tempData := map[string]interface{}{
		data.Key: data.Value,
	}

	expTimeData := time.Now().Add(time.Duration(data.TimeSeconds))
	tempTimeData := map[string]interface{}{
		data.Key: expTimeData,
	}
	metaData, err := json.Marshal(tempData)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	LocalCachedData.LocalCacheData = metaData
	LocalCachedData.ExpiryTime = tempTimeData

	go func() {
		time.Sleep(time.Second * time.Duration(data.TimeSeconds))
		err := DeleteKeyValue(data.Key)
		if err != nil {
			fmt.Println("unable to delete data")
			return
		}
	}()

	return nil
}

func Get(key string) interface{} {
	tempData := make(map[string]interface{})
	err := json.Unmarshal(LocalCachedData.LocalCacheData, &tempData)
	if err != nil {
		fmt.Println("unable to unmarshal data")
		return err
	}
	return tempData[key]
}

func DeleteKeyValue(key string) error {
	tempData := make(map[string]interface{})
	err := json.Unmarshal(LocalCachedData.LocalCacheData, &tempData)
	if err != nil {
		fmt.Println("unable to unmarshal data")
		return err
	}

	tempData[key] = nil

	metaData, err := json.Marshal(tempData)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	LocalCachedData.LocalCacheData = metaData
	return nil
}
