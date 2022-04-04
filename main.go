package main

import (
	"fmt"
	"github.com/675325758/go-misc-tools/mysql/sqlite3"
	"github.com/675325758/go-misc-tools/mysql/mysql"
	"github.com/675325758/go-misc-tools/redis"
	"github.com/675325758/go-misc-tools/config"
	"github.com/675325758/go-misc-tools/websocket"
	"github.com/675325758/go-misc-tools/cmd"
	"time"
)

func main() {
	fmt.Println("xxxxx")
	cmd.Test();
	redis.Test()
	websocket.Test()
	sqlite3.Test()
	mysql.Test();
	config.Test();
	for {
		time.Sleep(time.Second)
	}
}
