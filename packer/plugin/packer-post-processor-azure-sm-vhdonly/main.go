// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in the project root for license information.

package main

import (
	"github.com/Azure/packer-azure/packer/post-processor/azure-sm-vhdonly"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(azuresmvhdonly.PostProcessor))
	server.Serve()
}
