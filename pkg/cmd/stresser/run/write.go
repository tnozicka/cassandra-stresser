package run

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tnozicka/cassandra-stresser/pkg/genericclioptions"
	"github.com/tnozicka/cassandra-stresser/pkg/sharedoptions"
	"github.com/tnozicka/cassandra-stresser/pkg/signals"
	"github.com/tnozicka/cassandra-stresser/pkg/version"
	apierrors "k8s.io/apimachinery/pkg/util/errors"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
)

type WriteOptions struct {
	genericclioptions.IOStreams
	sharedoptions.Suite
}

func NewWriteOptions(streams genericclioptions.IOStreams) *WriteOptions {
	return &WriteOptions{
		Suite: *sharedoptions.NewSuite(),
	}
}

func NewWriteCmd(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewWriteOptions(streams)

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run cassandra-stress suite.",
		Long:  `Run the scylla operator.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := o.Validate(args)
			if err != nil {
				return err
			}

			err = o.Complete()
			if err != nil {
				return err
			}

			err = o.Run(streams, cmd)
			if err != nil {
				return err
			}

			return nil
		},
		Args: cobra.NoArgs,

		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.PersistentFlags().StringVarP()

	o.Suite.AddFlags(cmd)

	return cmd
}

func (o *WriteOptions) Validate(args []string) error {
	var errs []error

	errs = append(errs, o.Suite.Validate())

	return apierrors.NewAggregate(errs)
}

func (o *WriteOptions) Complete() error {
	var errs []error

	errs = append(errs, o.Suite.Complete())

	return apierrors.NewAggregate(errs)
}

func (o *WriteOptions) Run(streams genericclioptions.IOStreams, cmd *cobra.Command) error {
	klog.Infof("%s version %s", cmd.CommandPath(), version.Get())
	cliflag.PrintFlags(cmd.Flags())

	stopCh := signals.StopChannel()
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-stopCh
		cancel()
	}()

	return nil
}
