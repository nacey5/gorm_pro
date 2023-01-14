package server

import (
	"gorm_pro/chatroom-all/logic"
	"net/http"
	"os"
	"path/filepath"
)

func RegisterHandle() {
	inferRootDir()

	// 广播消息处理
	go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
}

var rootDir string

// inferRootDir 推断出项目根目录
func inferRootDir() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	var infer func(d string) string

	infer = func(d string) string {
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + "/chatroom-all/template") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	rootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
