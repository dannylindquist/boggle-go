package assets

import (
	"embed"
)

//go:embed css/*.css
//go:embed js/*.js
var AssetFS embed.FS
