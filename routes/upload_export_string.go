package routes

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/types"
	"github.com/labstack/echo/v4"

	"io"
	"net/http"
)

type PobPaste struct {
	Raw       string
	Decoded   string
	Character types.Character
}

func newPobPaste(raw string, decoded string, character types.Character) *PobPaste {
	return &PobPaste{
		Raw:       raw,
		Decoded:   decoded,
		Character: character,
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

	character_xml := string(unzipped)
	character, err := character.LoadCharacter(character_xml)
	if err != nil {
		return c.String(
			http.StatusNotAcceptable,
			err.Error(),
		)
	}
	paste := newPobPaste(raw, character_xml, character)

	return c.Render(
		http.StatusOK,
		"pobpaste",
		paste,
	)
}
