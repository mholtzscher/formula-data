package formuladata

import (
	"embed"
)

//go:embed sql/migrations/*.sql
var MigrationsFileSystem embed.FS
