package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets90ba780e7ce0fde2c751eedf5fb441ba62ac6407 = "<!doctype html>\n<body>\n  <p>Hello, {{.Foo}}</p>\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"html"}, "/templates/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1615022059, 1615022059539296059),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1615022097, 1615022097267296061),
		Data:     nil,
	}, "/templates/html": &assets.File{
		Path:     "/templates/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1615022240, 1615022240249296065),
		Data:     nil,
	}, "/templates/html/index.tmpl": &assets.File{
		Path:     "/templates/html/index.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1615022240, 1615022240251000000),
		Data:     []byte(_Assets90ba780e7ce0fde2c751eedf5fb441ba62ac6407),
	}}, "")
