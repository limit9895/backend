package main

import (
	"log"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var counter int
var mu sync.Mutex

type Response struct {
	Value int `json:"value"`
}

func main() {
	// Gin 라우터 설정
	router := gin.Default()

	// 특정 프록시만 신뢰하도록 설정
	router.SetTrustedProxies(nil)

	// CORS 설정
	router.Use(cors.Default())

	// 정적 파일 서버 설정 (React 빌드 파일 서빙)
	router.Static("/static", "../frontend/build/static")

	// API 엔드포인트 설정
	router.GET("/increment", func(c *gin.Context) {
		mu.Lock()
		counter++
		value := counter
		mu.Unlock()

		c.JSON(200, Response{Value: value})
	})

	router.GET("/decrement", func(c *gin.Context) {
		mu.Lock()
		counter--
		value := counter
		mu.Unlock()

		c.JSON(200, Response{Value: value})
	})

	// 모든 경로에서 React 빌드된 파일 서빙
	router.NoRoute(func(c *gin.Context) {
		c.File("../frontend/build/index.html")
	})

	// 서버 실행
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
