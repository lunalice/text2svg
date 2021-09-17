package main

import (
	"bytes"
	"fmt"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/labstack/echo"
)

type Request struct {
	Area     string `json:"area"`
	Category string `json:"category"`
}

func main() {
	e := echo.New()
	e.GET("/text2svg", MainPage())

	e.Logger.Fatal(e.Start(":1323"))
}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var u Request
		if err := c.Bind(&u); err != nil {
			return err
		}
		fmt.Println(u.Area)
		return c.HTML(http.StatusOK, text2svg(u))
	}
}

func show(c echo.Context) error {
	u := new(Request)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.Render(http.StatusOK, "page1", "hello world")
}

func text2svg(request Request) string {
	width := 500
	height := 500
	var temp bytes.Buffer
	canvas := svg.New(&temp)
	canvas.Start(500, 500)
	canvas.Text(width/2, height/2, request.Area, "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
	return temp.String()
}
