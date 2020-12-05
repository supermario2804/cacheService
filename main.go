package main

import (
	"cacheDataService/handlers"
	"cacheDataService/utils"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/api/setPage", handlers.SetPageCache)
	http.HandleFunc("/api/set", handlers.SetTableCache)
	http.HandleFunc("/api/get", handlers.GetTableCache)
	http.HandleFunc("/api/getPage", handlers.GetPageCache)
	go backupDB()
	fmt.Println("The server started...")
	http.ListenAndServe(":8090", nil)
}

func backupDB() {

	for {
		utils.PrintInfo("Taking backup of redis DB")
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		defer client.Close()

		_, connErr := client.Ping().Result()
		if connErr != nil {
			utils.HandleError(connErr, "Redis DB backup failed!!")
		}
		saveErr := client.BgSave().Err()
		if saveErr != nil {
			utils.HandleError(saveErr, "Redis DB backup has failed!")
		}

		utils.PrintInfo("Redis backup completed successfully!!")
		time.Sleep(30 * time.Minute)
	}

}
