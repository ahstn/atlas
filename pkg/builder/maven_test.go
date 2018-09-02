package builder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeOutput = `[INFO]
[INFO] ------------------------------------------------------------------------
[INFO] Building test-app 1.0.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[INFO] Downloading from central: https://repo.maven.apache.org/
[INFO] Downloading from central: https://repo.maven.apache.org/
[INFO]
[INFO] Installing mvn.pom to ~/.m2/repository
[INFO] --- maven-jar-plugin:3.0.2:jar (default-jar) @ DockerSpringVue ---
[INFO] Building jar: target/test-app.jar
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
`

var fakeErrorOutput = `
[INFO] ------------------------------------------------------------------------
[INFO] Building test-app 1.0.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[ERROR] No goals have been specified for this build.
[ERROR]
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR]
[ERROR] For more information about the errors and possible solutions, please read the following articles
`

var fakeTestOutput = `
[INFO] ------------------------------------------------------------------------
[INFO] Building test-app 1.0.0-SNAPSHOT
[INFO] ------------------------------------------------------------------------
[ERROR] TestGetUserStatusOkay(io.ahstn.usersapi.repositories.UserRepositoryTest)  Time elapsed: 1.004 s  <<< FAILURE!
[INFO]
[INFO] Results:
[INFO]
[ERROR] Failures:
[ERROR]   UserRepositoryTest.TestGetUserStatusOkay:39
[INFO]
[ERROR] Tests run: 2, Failures: 1, Errors: 0, Skipped: 0
`

func TestNewClient(t *testing.T) {
	mvn := NewClient("./", nil, []string{"install"}, []string{"-DskipTests"})

	assert.Equal(t, []string{"mvn", "--batch-mode", "install", "-DskipTests"}, mvn.cmd.Args)
}

func TestNewCustomClient(t *testing.T) {
	pathenv := os.Getenv("PATH")
	defer os.Setenv("PATH", pathenv)

	path := SetupMvnInPath(t)

	mvn := NewCustomClient(path, "./", nil, []string{"install"}, []string{"-DskipTests"})
	assert.Equal(t, []string{"mvn", "--batch-mode", "install", "-DskipTests"}, mvn.cmd.Args)
	assert.Equal(t, path, mvn.cmd.Path)
}

func TestRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommand("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(false)

	assert.Contains(t, buf.String(), "✔  test-app 1.0.0-SNAPSHOT Complete")
	assert.Contains(t, buf.String(), "✔  jar: target/test-app.jar Complete")
}

func TestJUnitFailureRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommandTestFailure("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(false)

	assert.Contains(t, buf.String(), "Failures:")
	assert.Contains(t, buf.String(), "UserRepositoryTest.TestGetUserStatusOkay:39")
	assert.Contains(t, buf.String(), "Tests run: 2, Failures: 1, Errors: 0")
}

func TestErrorRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommandError("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(false)

	assert.Contains(t, buf.String(), "[ERROR] No goals have been specified for this build.")
	assert.NotContains(t, buf.String(), "To see the full stack trace of the errors")
	assert.NotContains(t, buf.String(), "Re-run Maven using the -X switch")
}

func TestVerboseRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommand("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(true)

	assert.Contains(t, buf.String(), "[INFO] Building test-app 1.0.0-SNAPSHOT")
}

func TestArgs(t *testing.T) {
	mvn := NewClient("./", nil, []string{"clean install"}, nil)

	assert.Equal(t, "mvn --batch-mode clean install", mvn.Args())
}

func TestModifyArgs(t *testing.T) {
	mvn := NewClient("./", nil, []string{"clean install"}, nil)
	mvn.ModifyArgs([]string{"package"})

	assert.Equal(t, "mvn --batch-mode package", mvn.Args())
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func fakeExecCommandTestFailure(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "GO_WANT_TEST_FAIL=1"}
	return cmd
}

func fakeExecCommandError(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "GO_WANT_ERR=1"}
	return cmd
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	if os.Getenv("GO_WANT_ERR") == "1" {
		fmt.Print(fakeErrorOutput)
	} else if os.Getenv("GO_WANT_TEST_FAIL") == "1" {
		fmt.Print(fakeTestOutput)
	} else {
		fmt.Print(fakeOutput)
	}
}

func SetupMvnInPath(t *testing.T) string {
	tmp, err := ioutil.TempDir("", "TestInitialiseCommand")
	if err != nil {
		t.Fatal("TempDir failed: ", err)
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed: ", err)
	}
	err = os.Chdir(tmp)
	if err != nil {
		t.Fatal("Chdir failed: ", err)
	}
	defer os.Chdir(wd)

	f, err := os.OpenFile("mvn", os.O_CREATE, 0777)
	if err != nil {
		t.Fatal("OpenFile failed: ", err)
	}
	err = f.Close()
	if err != nil {
		t.Fatal("Close failed: ", err)
	}

	err = os.Setenv("PATH", fmt.Sprintf("%s:%s", tmp, os.Getenv("PATH")))
	if err != nil {
		t.Fatal("Setenv failed: ", err)
	}

	return path.Join(tmp, "mvn")
}
