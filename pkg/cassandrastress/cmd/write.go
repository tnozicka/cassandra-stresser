package cmd

import (
	"context"
	"os/exec"

	"github.com/tnozicka/cassandra-stresser/pkg/cassandrastress/shared"
	"github.com/tnozicka/cassandra-stresser/pkg/genericclioptions"
)

type WriteConfig struct {
	shared.OperationConfig
}

func NewWriteConfig() *WriteConfig {
	return &WriteConfig{
		OperationConfig: *shared.NewOperationConfig(),
	}
}

func (c *WriteConfig) ToCmd(ctx context.Context, streams genericclioptions.IOStreams) *exec.Cmd {
	return c.OperationConfig.ToCmd(ctx, streams)
}
