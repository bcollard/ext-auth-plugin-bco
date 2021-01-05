package main

import (
	impl "github.com/solo-io/ext-auth-plugin-examples/plugins/example_plugin/pkg"
	"github.com/solo-io/ext-auth-plugins/api"
)

func main() {}

// Compile-time assertion
var _ api.ExtAuthPlugin = new(impl.RequiredHeaderPlugin)

// This is the exported symbol that Gloo will look for.
//noinspection GoUnusedGlobalVariable
var Plugin impl.RequiredHeaderPlugin
