package main

import (
	"fmt"
	"gorm_pro/chatroom-all/global"
	"gorm_pro/chatroom-all/server"
	"log"
	"net/http"
)

func init() {
	global.Init()
}

var (
	addr   = ":2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

   ChatRoom，start on：%s
`
)

func main() {
	fmt.Printf(banner+"\n", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
