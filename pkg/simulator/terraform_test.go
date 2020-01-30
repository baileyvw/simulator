package simulator_test

import (
	sim "github.com/controlplaneio/simulator-standalone/pkg/simulator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"

	"os"
)

var pwd, _ = os.Getwd()
var testVarFileArg = "--var-file=" + pwd + "/" + fixture("noop-tf-dir") + "/settings/bastion.tfvars"
var noopLogger = zap.NewNop().Sugar()
var testModuleArg = "-from-module=" + pwd + "/" + fixture("noop-tf-dir")

var tfCommandArgumentsTests = []struct {
	prepArgs  []string
	arguments []string
}{
	{[]string{"output"}, []string{"output", "-json"}},
	{[]string{"init"}, []string{"init", "-input=false", testVarFileArg, testModuleArg, "-backend-config=bucket=test-bucket"}},
	{[]string{"plan"}, []string{"plan", "-input=false", testVarFileArg}},
	{[]string{"apply"}, []string{"apply", "-input=false", testVarFileArg, "-auto-approve"}},
	{[]string{"destroy"}, []string{"destroy", "-input=false", testVarFileArg, "-auto-approve"}},
}

func Test_PrepareTfArgs(t *testing.T) {

	pwd, _ := os.Getwd()
	simulator := sim.NewSimulator(
		sim.WithLogger(noopLogger),
		sim.WithBucketName("test-bucket"),
		sim.WithTfVarsDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithTfDir(pwd+"/"+fixture("noop-tf-dir")))

	for _, tt := range tfCommandArgumentsTests {
		t.Run("Test arguments for "+tt.prepArgs[0], func(t *testing.T) {
			assert.Equal(t, simulator.PrepareTfArgs(tt.prepArgs[0]), tt.arguments)
		})
	}
}

func Test_Status(t *testing.T) {
	pwd, _ := os.Getwd()
	simulator := sim.NewSimulator(
		sim.WithLogger(noopLogger),
		sim.WithTfDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithScenariosDir("test"),
		sim.WithAttackTag("latest"),
		sim.WithBucketName("test"),
		sim.WithTfVarsDir(pwd+"/"+fixture("noop-tf-dir")))

	tfo, err := simulator.Status()

	assert.Nil(t, err, "Got an error")
	assert.NotNil(t, tfo, "Got no terraform output")
}

func Test_Create(t *testing.T) {

	pwd, _ := os.Getwd()
	simulator := sim.NewSimulator(
		sim.WithLogger(noopLogger),
		sim.WithTfDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithScenariosDir("test"),
		sim.WithAttackTag("latest"),
		sim.WithBucketName("test"),
		sim.WithTfVarsDir(pwd+"/"+fixture("noop-tf-dir")))

	err := simulator.Create()
	assert.Nil(t, err)
}

func Test_Destroy(t *testing.T) {

	pwd, _ := os.Getwd()
	simulator := sim.NewSimulator(
		sim.WithLogger(noopLogger),
		sim.WithTfDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithAttackTag("latest"),
		sim.WithBucketName("test"),
		sim.WithTfVarsDir(pwd+"/"+fixture("noop-tf-dir")))

	err := simulator.Destroy()

	assert.Nil(t, err)
}
