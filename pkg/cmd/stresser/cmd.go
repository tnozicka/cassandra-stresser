package stresser

import (
	"github.com/spf13/cobra"
	"github.com/tnozicka/cassandra-stresser/pkg/cmd/stresser/run"
	"github.com/tnozicka/cassandra-stresser/pkg/cmdutil"
	"github.com/tnozicka/cassandra-stresser/pkg/genericclioptions"
)

func NewCommand(streams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use: "cassandra-stresser",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return cmdutil.ReadFlagsFromEnv("CASSANDRA_STRESSER", cmd)
		},
	}

	cmd.AddCommand(run.NewCmd(streams))

	cmdutil.InstallKlog(cmd)

	return cmd
}
