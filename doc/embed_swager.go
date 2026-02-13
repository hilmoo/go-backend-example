package doc

import (
	_ "embed"
)

//go:embed swagger/swagger.json
var SwaggerJSON []byte