/*
 * @File: databases.databases.go
 * @Description: Creates global database instance
 * @Author: Hua Van Thong (huavanthong14@gmail.com)
 */
package databases

import (
	"fmt"
	"sync"
)

// Database shares global database instance
var (
	Database MongoDB
)

/*********************************************************************
Design Pattern: Singleton

Feature:
	+ Get instance DB for MongoDB

Purpose:
	+ to get only a instance for DB,
	+ and other clients can't access to DB if instance is exist

Refer: https://golangbyexample.com/singleton-design-pattern-go/
*********************************************************************/
var singleInstance *MongoDB

var lock = &sync.Mutex{}

func getInstanceDB() *MongoDB {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creting Single Instance Now")
			singleInstance = &Database
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}

	return singleInstance
}
