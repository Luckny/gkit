package main

import (
	"github.com/alecthomas/kong"
)

var cli struct{}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("gkit"),
		kong.Description("A git utility wrapper"),
		kong.UsageOnError(),
	)

	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
