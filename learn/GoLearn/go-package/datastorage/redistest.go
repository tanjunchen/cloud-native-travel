package main

import (
	"fmt"
	"gopkg.in/redis.v5"
	"time"
)

// Setup 初始化 redis 连接
func RedisSetup() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "a",
		DB:       0, // 使用默认 DB
	})
	_, err := client.Ping().Result()
	return client, err
}

// ExecRedis 执行一些 redis 操作
func ExecRedis() error {
	conn, err := RedisSetup()
	if err != nil {
		return err
	}
	c1 := "value"
	conn.Set("key", c1, 5*time.Second)
	var result string
	if err := conn.Get("key").Scan(&result); err != nil {
		switch err {
		case redis.Nil:
			return nil
		default:
			return err
		}
	}
	fmt.Println("result =", result)
	return nil
}

func SortRedis() error {
	conn, err := RedisSetup()
	if err != nil {
		return err
	}
	if err := conn.LPush("list", 1).Err(); err != nil {
		return err
	}
	if err := conn.LPush("list", 3).Err(); err != nil {
		return err
	}
	if err := conn.LPush("list", 2).Err(); err != nil {
		return err
	}
	res, err := conn.Sort("list", redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	conn.Del("list")
	return nil
}

func test()  {
	if err := ExecRedis(); err != nil {
		panic(err)
	}

	if err := SortRedis(); err != nil {
		panic(err)
	}
}

func main()  {
	test()
}