package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techdenglei/eshop/tool"
)

func main() {
	router := gin.Default()

	cfg, err := tool.ParseConfig("./config/App.json")
	if err != nil {
		panic(err.Error())
	}
	router.Run(cfg.Host + ":" + cfg.Port)
}
