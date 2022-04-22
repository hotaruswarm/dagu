package config_test

import (
	"fmt"
	"jobctl/internal/config"
	"jobctl/internal/settings"
	"jobctl/internal/utils"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testDir     = path.Join(utils.MustGetwd(), "../../tests/testdata")
	testHomeDir = path.Join(utils.MustGetwd(), "../../tests/config")
	testConfig  = path.Join(testDir, "all.yaml")
	testEnv     = []string{}
)

func TestMain(m *testing.M) {
	settings.InitTest(testHomeDir)
	testEnv = []string{
		fmt.Sprintf("LOG_DIR=%s", path.Join(testHomeDir, "/logs")),
		fmt.Sprintf("PATH=%s", os.ExpandEnv("${PATH}")),
	}
	code := m.Run()
	os.Exit(code)
}

func TestAssertDefinition(t *testing.T) {
	loader := config.NewConfigLoader()

	_, err := loader.Load(path.Join(testDir, "err_no_name.yaml"), "")
	require.Equal(t, err, fmt.Errorf("job name must be specified."))

	_, err = loader.Load(path.Join(testDir, "err_no_steps.yaml"), "")
	require.Equal(t, err, fmt.Errorf("at least one step must be specified."))
}

func TestAssertStepDefinition(t *testing.T) {
	loader := config.NewConfigLoader()

	_, err := loader.Load(path.Join(testDir, "err_step_no_name.yaml"), "")
	require.Equal(t, err, fmt.Errorf("step name must be specified."))

	_, err = loader.Load(path.Join(testDir, "err_step_no_command.yaml"), "")
	require.Equal(t, err, fmt.Errorf("step command must be specified."))
}

func TestReadConfig(t *testing.T) {
	f, err := config.ReadConfig(testConfig)
	require.NoError(t, err)
	if len(f) == 0 {
		t.Error("reading yaml file failed")
	}
}
