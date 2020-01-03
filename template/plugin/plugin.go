// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/converter"
)

// New returns a new conversion plugin.
func New(param1, param2 string) converter.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) Convert(ctx context.Context, req *converter.Request) (*drone.Config, error) {
	// TODO this should be modified or removed. For
	// demonstration purposes we show how you can ignore
	// certain configuration files by file extension.
	if strings.HasSuffix(req.Repo.Config, ".xyz") {
		// a nil response instructs the Drone server to
		// use the configuration file as-is, without
		// modification.
		return nil, nil
	}

	// get the configuration file from the request.
	config := req.Config.Data

	// TODO this should be modified or removed. For
	// demonstration purposes we make a simple modification
	// to the configuration file and add a newline.
	config = config + "\n"

	// returns the modified configuration file.
	return &drone.Config{
		Data: config,
	}, nil
}
