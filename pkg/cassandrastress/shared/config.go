package shared

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/tnozicka/cassandra-stresser/pkg/genericclioptions"
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

type OperationConfig struct {
	CassandraStressPath    string
	MaxErrorThreshold      float32
	MinIterations          int
	MaxIterations          int
	Warmup                 bool
	Truncate               TruncateType
	ConsistencyLevel       ConsistencyLevelType
	SerialConsistencyLevel SerialConsistencyLevelType
	OperationsCount        int
	Timeout                time.Duration
}

func NewOperationConfig() *OperationConfig {
	return &OperationConfig{
		CassandraStressPath:    "cassandra-stress",
		MaxErrorThreshold:      0.02,
		MinIterations:          30,
		MaxIterations:          200,
		Warmup:                 true,
		Truncate:               TruncateTypes[0],
		ConsistencyLevel:       ConsistencyLevels[0],
		SerialConsistencyLevel: SerialConsistencyLevels[0],
		OperationsCount:        5,
		Timeout:                0,
	}
}

func (c *OperationConfig) ToCmd(ctx context.Context, streams genericclioptions.IOStreams) *exec.Cmd {
	cmd := exec.CommandContext(ctx, c.CassandraStressPath)

	cmd.Args = append(cmd.Args, fmt.Sprintf("err<%v", c.MaxErrorThreshold))
	cmd.Args = append(cmd.Args, fmt.Sprintf("n>%d", c.MinIterations))
	cmd.Args = append(cmd.Args, fmt.Sprintf("n<%d", c.MaxIterations))
	if !c.Warmup {
		cmd.Args = append(cmd.Args, "no-warmup")
	}
	cmd.Args = append(cmd.Args, fmt.Sprintf("truncate=%s", c.Truncate))
	cmd.Args = append(cmd.Args, fmt.Sprintf("cl=%s", c.ConsistencyLevel))
	cmd.Args = append(cmd.Args, fmt.Sprintf("serial-cl=%s", c.SerialConsistencyLevel))
	cmd.Args = append(cmd.Args, fmt.Sprintf("n=%d", c.OperationsCount))
	cmd.Args = append(cmd.Args, fmt.Sprintf("duration=%d", c.Timeout))

	return cmd
}
