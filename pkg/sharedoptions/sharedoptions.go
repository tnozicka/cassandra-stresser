package sharedoptions

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type TruncateType string

var TruncateTypes = []TruncateType{"never", "always", "once"}

type ConsistencyLevelType string

var ConsistencyLevels = []ConsistencyLevelType{
	"LOCAL_ONE",
}

type SerialConsistencyLevelType string

var SerialConsistencyLevels = []SerialConsistencyLevelType{
	"SERIAL",
}

type Suite struct {
	CassandraStressPath    string
	MaxErrorThreshold      float32
	MinIterations          int
	MaxIterations          int
	Warmup                 bool
	Truncate               string
	ConsistencyLevel       string
	SerialConsistencyLevel SerialConsistencyLevelType
	OperationsCount        int
	Timeout                time.Duration
}

func NewSuite() *Suite {
	return &Suite{
		CassandraStressPath:    "cassandra-stress",
		MaxErrorThreshold:      0.02,
		MinIterations:          30,
		MaxIterations:          200,
		Warmup:                 true,
		Truncate:               string(TruncateTypes[0]),
		ConsistencyLevel:       string(ConsistencyLevels[0]),
		SerialConsistencyLevel: SerialConsistencyLevels[0],
		OperationsCount:        5,
		Timeout:                0,
	}
}

func (so *Suite) AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&so.CassandraStressPath, "cassandra-stress", "", so.CassandraStressPath, "Path to cassandra-stress binary.")
	cmd.PersistentFlags().Float32VarP(&so.MaxErrorThreshold, "max-error-threshold", "", so.MaxErrorThreshold, "Run until the standard error of the mean is below this fraction.")
	cmd.PersistentFlags().IntVarP(&so.MinIterations, "min-iterations", "", so.MinIterations, "Run at least this many iterations before accepting uncertainty convergence.")
	cmd.PersistentFlags().IntVarP(&so.MaxIterations, "max-iterations", "", so.MaxIterations, "Run at most this many iterations before accepting uncertainty convergence.")
	cmd.PersistentFlags().BoolVarP(&so.Warmup, "warmup", "", so.Warmup, "Warm up the process.")
	cmd.PersistentFlags().StringVarP(&so.Truncate, "truncate", "", so.Truncate, fmt.Sprintf("Truncate the table. Valid options are %q.", TruncateTypes))
	cmd.PersistentFlags().StringVarP(&so.ConsistencyLevel, "consistency", "", so.ConsistencyLevel, fmt.Sprintf("Consistency level to use. Valid options are %q.", ConsistencyLevels))
	cmd.PersistentFlags().StringVarP((*string)(&so.SerialConsistencyLevel), "serial-consistency", "", string(so.SerialConsistencyLevel), fmt.Sprintf("Serial consistency level to use. Valid options are %q.", SerialConsistencyLevels))
	cmd.PersistentFlags().IntVarP(&so.OperationsCount, "operations", "", so.OperationsCount, "Number of operations to perform.")
	cmd.PersistentFlags().DurationVarP(&so.Timeout, "suite-timeout", "", so.Timeout, "Time limit.")
}

func (so *Suite) Validate() error {
	return nil
}

func (so *Suite) Complete() error {
	return nil
}
