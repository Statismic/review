package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1552197605, 0),

		Content: string("<!DOCTYPE html>\n<html lang=\"en\">\n\n<head>\n</head>\n\n<body>\n    <iframe src=\"/master\" style=\"height: 100vh; width: 49%;\"></iframe>\n    <iframe src=\"/pr\" style=\"height: 100vh; width: 49%;\"></iframe>\n</body>\n\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1552195606, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`public`, &embedded.EmbeddedBox{
		Name: `public`,
		Time: time.Unix(1552195606, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.html": file2,
		},
	})
}
