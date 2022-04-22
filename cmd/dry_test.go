package main

import (
	"testing"
)

func Test_dryCommand(t *testing.T) {
	tests := []appTest{
		{
			args: []string{"", "dry", testConfig("basic_success.yaml")}, errored: false,
			output: []string{"Starting DRY-RUN"},
		},
	}

	for _, v := range tests {
		app := makeApp()
		runAppTestOutput(app, v, t)
	}
}
