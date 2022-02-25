package openapiv2

import (
	"embed"
)

//go:embed v1/*
var OpenAPIDocsV1 embed.FS
