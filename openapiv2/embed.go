package openapiv2

import (
	"embed"
)

//go:embed v1/*
var OpenAPIDocsV1 embed.FS

//go:embed v2/*
var OpenAPIDocsV2 embed.FS
