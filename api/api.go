package main

import (
	"github.com/RomaOnyshkiv/app/home"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	http.HandleFunc("/", home.Home)
	http.ListenAndServe("0.0.0.0:8088", nil)

}
