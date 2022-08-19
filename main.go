package main

import (
	"example.com/m/cache"
	"example.com/m/models"
	"fmt"
	"time"
)

func main() {
	LocalCacheMain()
}

func LocalCacheMain() {
	var opp models.Operations

	for opp != models.ExitOperations {

		fmt.Println("select operations")
		_, err := fmt.Scan(&opp)
		if err != nil {
			panic(err)
			return
		}

		switch opp {
		case models.SetOperations:
			set()
		case models.GetOperations:
			get()
		case models.DeleteOperations:
			deleteKeyValue()
		case models.ExitOperations:
			return
		default:
			return
		}
	}
}

func set() {
	var key string
	var value string
	var expTime int
	fmt.Println("enter [key]")
	_, err := fmt.Scan(&key)
	if err != nil || key == "" {
		fmt.Println("invalid key ")
		return
	}
	fmt.Println("enter [value]")
	_, err = fmt.Scan(&value)
	if err != nil || value == "" {
		fmt.Println("invalid  value")
		return
	}
	fmt.Println("enter [EXP time in seconds]")
	_, err = fmt.Scan(&expTime)
	if err != nil || expTime == 0 {
		fmt.Println("invalid time")
		return
	}

	data := models.KeyValueStruct{Key: key, Value: value, TimeSeconds: expTime}
	go func() {
	err = cache.Set(data)
		if err != nil {
			fmt.Println("unable to set data")
			return
		}
	}()
	fmt.Println("value and Expiry time added at key successfully")
}

func get() {
	var key string
	fmt.Println("enter [key]")
	_, err := fmt.Scan(&key)
	if err != nil || key == "" {
		fmt.Println("invalid key ")
		return
	}
	value := cache.Get(key)
	if value == nil {
		fmt.Println("unable to get data")
		return
	}

	fmt.Println("[key]", key)
	fmt.Println("[value]", value)
}

func deleteKeyValue() {
	var key string
	fmt.Println("enter [key]")
	_, err := fmt.Scan(&key)
	if err != nil && key == "" {
		fmt.Println("invalid key ")
		return
	}

	err = cache.DeleteKeyValue(key)
	if err != nil {
		fmt.Println("unable to delete data")
		return
	}
	fmt.Println("deleted")
}
