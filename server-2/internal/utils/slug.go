package utils

import (
	"strings"
)

func GenerateSlug(name string) string {
	// simple: lowercase + replace spaces with hyphen
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.Trim(slug, "-")
	return slug
}