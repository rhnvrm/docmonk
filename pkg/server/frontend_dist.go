//go:build dist

package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rhnvrm/docmonk/frontend/dist"
)

func getFS() http.FileSystem {
	// list all files in fs
	files, _ := fs.ReadDir(dist.Assets, ".")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	return http.FS(dist.Assets)
}

func init() {
	handleRoot = func() echo.HandlerFunc {
		return echo.WrapHandler(http.FileServer(getFS()))
	}

	handleRootStatic = func() echo.HandlerFunc {
		return echo.WrapHandler(http.StripPrefix("/assets/", http.FileServer(getFS())))
	}
}
