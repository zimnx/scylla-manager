// Copyright (C) 2017 ScyllaDB

package client

import (
	"github.com/scylladb/mermaid/command"
	"github.com/scylladb/mermaid/restapiclient"
	"github.com/scylladb/mermaid/restapiclient/client/operations"
)

// BaseClientCommand handles common CLI client behaviours.
type BaseClientCommand struct {
	command.BaseCommand
	APIHost string
}

func (cmd *BaseClientCommand) client() *operations.Client {
	return restapiclient.New(cmd.APIHost)
}
