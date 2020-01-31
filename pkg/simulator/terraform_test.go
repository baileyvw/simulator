package simulator_test

import (
	sim "github.com/controlplaneio/simulator-standalone/pkg/simulator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"

	"io/ioutil"
	"os"
)

var pwd, _ = os.Getwd()
var testVarFileArg = "--var-file=" + pwd + "/" + fixture("noop-tf-dir") + "/settings/bastion.tfvars"
var tfDir = pwd + "/" + fixture("noop-tf-dir")
var noopLogger = zap.NewNop().Sugar()

var tfCommandArgumentsTests = []struct {
	prepArgs  []string
	arguments []string
}{
	{[]string{"output"}, []string{"output", "-json"}},
	{[]string{"init"}, []string{"init", "-input=false", testVarFileArg, "-backend-config=bucket=test-bucket", tfDir}},
	{[]string{"plan"}, []string{"plan", "-input=false", testVarFileArg, tfDir}},
	{[]string{"apply"}, []string{"apply", "-input=false", testVarFileArg, "-auto-approve", tfDir}},
	{[]string{"destroy"}, []string{"destroy", "-input=false", testVarFileArg, "-auto-approve"}},
}

func Test_PrepareTfArgs(t *testing.T) {

	pwd, _ := os.Getwd()
	simulator := sim.NewSimulator(
		sim.WithLogger(noopLogger),
		sim.WithBucketName("test-bucket"),
		sim.WithTfDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithTfVarsDir(pwd+"/"+fixture("noop-tf-dir")))

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
	tmpDir, _ := ioutil.TempDir("", "")
	simulator := sim.NewSimulator(
		sim.WithLogger(noopLogger),
		sim.WithTfDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithScenariosDir("test"),
		sim.WithAttackTag("latest"),
		sim.WithBucketName("test"),
		sim.WithTfVarsDir(pwd+"/"+fixture("noop-tf-dir")),
		sim.WithTmpDir(tmpDir))

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
