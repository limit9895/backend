package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 라우터 설정
	router := gin.Default()

	// 특정 프록시만 신뢰하도록 설정
	router.SetTrustedProxies(nil)

	// CORS 설정
	router.Use(cors.Default())

	// 정적 파일 서버 설정 (React 빌드 파일 서빙), 즉 프론트앤드 동작시 static으로 들어오는 내용에 대해 해당 경로로 안내
	router.Static("/static", "../frontend/build/static")

	// 모든 경로에서 React 빌드된 파일 서빙
	router.NoRoute(func(c *gin.Context) {
		c.File("../frontend/build/index.html")
	})
	// 서버 실행
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
