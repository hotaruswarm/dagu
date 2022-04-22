package models_test

import (
	"jobctl/internal/config"
	"jobctl/internal/models"
	"jobctl/internal/scheduler"
	"jobctl/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func makeStep(cmd string) *config.Step {
	step := &config.Step{
		Name: "test step",
	}
	step.Command, step.Args = utils.SplitCommand(cmd)
	return step
}

func TestFromNodes(t *testing.T) {
	g := testRunSteps(
		t,
		makeStep("true"),
		makeStep("false"),
	)

	ret := models.FromNodes(g.Nodes())

	assert.Equal(t, 2, len(ret))
	assert.NotEqual(t, "", ret[1].Error)
}

func TestToNode(t *testing.T) {
	g := testRunSteps(
		t,
		makeStep("true"),
		makeStep("true"),
	)
	orig := g.Nodes()
	for _, n := range orig {
		require.Equal(t, scheduler.NodeStatus_Success, n.Status)
	}
	nodes := models.FromNodes(orig)
	for i := range nodes {
		n := nodes[i].ToNode()
		require.Equal(t, n.Step, orig[i].Step)
		require.Equal(t, n.NodeState, orig[i].NodeState)
	}
}

func testRunSteps(t *testing.T, steps ...*config.Step) *scheduler.ExecutionGraph {
	g, err := scheduler.NewExecutionGraph(steps...)
	require.NoError(t, err)
	for _, n := range g.Nodes() {
		if err := n.Execute(); err != nil {
			n.Status = scheduler.NodeStatus_Error
		} else {
			n.Status = scheduler.NodeStatus_Success
		}
	}
	return g
}
