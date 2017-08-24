package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getBooks(e echo.Context) error {
	return e.String(http.StatusOK, "books")
}

func getBooksByID(e echo.Context) error {
	return e.String(http.StatusOK, fmt.Sprintf("%s: %s", e.Param("id"), "book"))
}

func main() {
	e := echo.New()

	e.GET("/books", getBooks)
	e.GET("/books/:id", getBooksByID)

	e.Logger.Fatal(e.Start(":1323"))
}
