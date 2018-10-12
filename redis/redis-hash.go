package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		ReadTimeout:  1000 * 1000 * 1000 * 60 * 60 * 24,
		WriteTimeout: 1000 * 1000 * 1000 * 60 * 60 * 24,
	})

	maxLength := int64(10000 * 100)

	for n := int64(1); n <= maxLength; n *= 10 {
		fmt.Println("Set个数", n)
		TestDelBigSet(client, n)
		TestExpireBigSet(client, n)
		TestUnlinkBigSet(client, n)
		fmt.Println()
	}
}

func TestDelBigSet(client *redis.Client, count int64) {
	redisKey := fmt.Sprintf("%s%d", "del:", time.Now().Nanosecond())

	for n := int64(0); n < count; n++ {
		err := client.SAdd(redisKey, fmt.Sprintf("%d", n)).Err()
		if err != nil {
			panic(err)
		}
	}

	startTime := CurrentTimestampInMicroSecond()
	client.Del(redisKey)
	endTime := CurrentTimestampInMicroSecond()

	fmt.Println("Del", endTime-startTime)
}

func TestUnlinkBigSet(client *redis.Client, count int64) {
	redisKey := fmt.Sprintf("%s%d", "unlink:", time.Now().Nanosecond())

	for n := int64(0); n < count; n++ {
		err := client.SAdd(redisKey, fmt.Sprintf("%d", n)).Err()
		if err != nil {
			panic(err)
		}
	}
	startTime := CurrentTimestampInMicroSecond()
	client.Unlink(redisKey)
	endTime := CurrentTimestampInMicroSecond()

	fmt.Println("Unlink", endTime-startTime)
}

func TestExpireBigSet(client *redis.Client, count int64) {
	redisKey := fmt.Sprintf("%s%d", "expire:", time.Now().Nanosecond())

	for n := int64(0); n < count; n++ {
		err := client.SAdd(redisKey, fmt.Sprintf("%d", n)).Err()
		if err != nil {
			panic(err)
		}
	}
	startTime := CurrentTimestampInMicroSecond()
	client.Expire(redisKey, 0)
	endTime := CurrentTimestampInMicroSecond()

	fmt.Println("Expire", endTime-startTime)
}

func CurrentTimestampInMicroSecond() int64 {
	return time.Now().UnixNano() / 1000
}
