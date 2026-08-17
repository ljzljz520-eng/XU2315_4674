package web

import (
	"embed"
	"io/fs"
)

//go:embed index.html app.js styles.css
var assets embed.FS

func Assets() fs.FS {
	return assets
}

func Index() []byte {
	contents, err := assets.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	return contents
}
