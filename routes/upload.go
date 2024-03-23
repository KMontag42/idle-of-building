package routes

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func UploadFile(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
}

type PobPaste struct {
	Raw     string
	Decoded string
}

func newPobPaste(raw string, decoded string) *PobPaste {
	return &PobPaste{
		Raw:     raw,
		Decoded: decoded,
	}
}

func UploadExportString(c echo.Context) error {
	// Read form fields
	raw := c.FormValue("raw")

	// urlsafe base64 decode the raw string
	decoded, err := base64.URLEncoding.DecodeString(raw)
	if err != nil {
		return c.String(
			http.StatusNotAcceptable,
			err.Error(),
		)
	}

	// zlib decompress the decoded string
	buffer := bytes.NewReader(decoded)
	zlib_reader, err := zlib.NewReader(buffer)
	if err != nil {
		return c.String(
			http.StatusNotAcceptable,
			err.Error(),
		)
	}

	unzipped, err := io.ReadAll(zlib_reader)
	if err != nil {
		return c.String(
			http.StatusNotAcceptable,
			err.Error(),
		)
	}

	paste := newPobPaste(raw, string(unzipped))

	return c.Render(
		http.StatusOK,
		"pobpaste",
		paste,
	)
}
