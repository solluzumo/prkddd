package utils

import (
	"strings"

	"github.com/essentialkaos/translit/v3"
)

func SlugifyName(name string) string {
	return strings.ReplaceAll(translit.ICAO(strings.ToLower(name)), " ", "-")
}
