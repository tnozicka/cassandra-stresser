package run

import (
	"github.com/spf13/cobra"
	"github.com/tnozicka/cassandra-stresser/pkg/genericclioptions"
)

func NewCmd(streams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "run",
		Short:         "Run cassandra-stress suite.",
		Long:          "Run cassandra-stress suite.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand()

	return cmd
}
