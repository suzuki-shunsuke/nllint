package main

import (
	"github.com/suzuki-shunsuke/nllint/pkg/cli"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
)

var version = ""

func main() {
	urfave.Main("nllint", version, cli.Run)
}
