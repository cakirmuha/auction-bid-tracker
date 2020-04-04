// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/cakirmuha/auction-bid-tracker/cmd/bid.tracker/assets"
)

//go:generate go run -tags=dev gen.go
func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		Filename:     "assets_vfsdata.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
