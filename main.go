package main

import (
	_ "github.com/DanielShoomoN/hiddify-split-tunnel-extension/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
