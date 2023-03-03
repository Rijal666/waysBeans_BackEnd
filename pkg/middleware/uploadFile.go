package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("photo")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		tempfile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest,err)
		}
		defer tempfile.Close()

		if _, err := io.Copy(tempfile, src); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		data := tempfile.Name()
		filename:= data[:8]

		c.Set("dataFile", filename)
		return next(c)
	}
}