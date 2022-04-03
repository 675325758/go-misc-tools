package main

import (
	"fmt"
	"github.com/675325758/go-misc-tools/mysql/sqlite3"
	"github.com/675325758/go-misc-tools/mysql/mysql"
	"github.com/675325758/go-misc-tools/redis"
	"github.com/675325758/go-misc-tools/websocket"
	"time"
)

func main() {
	fmt.Println("xxxxx")
	redis.Test()
	websocket.Test()
	sqlite3.Test()
	mysql.Test();
	for {
		time.Sleep(time.Second)
	}
}
