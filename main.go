package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// レスポンス構造体を定義
type Member struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/members", func(c echo.Context) error {
		// Member型のリスト
		members := []Member{
			{Id: 1, Name: "fujie"},
			{Id: 2, Name: "katayama"},
			{Id: 3, Name: "ito"},
			{Id: 4, Name: "shimizu"},
			{Id: 5, Name: "hosoda"},
			{Id: 6, Name: "kamizato"},
			{Id: 7, Name: "motegi"},
			{Id: 8, Name: "fujiki"},
			{Id: 9, Name: "toyoda"},
			{Id: 10, Name: "tanaka"},
			{Id: 11, Name: "kondo"},
		}

		return c.JSON(http.StatusOK, members)
	})

	// Start server
	// ポート番号は環境変数から取得
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
