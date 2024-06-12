package main

import (
	_ "embed"

	"github.com/aegis-dec/polygon-edge/command/root"
	"github.com/aegis-dec/polygon-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
