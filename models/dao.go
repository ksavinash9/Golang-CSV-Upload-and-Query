package models
import (
    "gopkg.in/redis.v3"
    "fmt"
)

const (
//redis
    redis_db = 0
    redis_password = ""

    redis_port = "6379"
    redis_host = "localhost"
)


func SetupRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        
        Addr:     redis_host+":"+redis_port,
        Password: redis_password, // no password set
        DB:       redis_db,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)

    return client
}

type Response struct {
    Status int `json:"status"`
    Msg    string `json:"msg"`
    Data   interface{} `json:data`
}