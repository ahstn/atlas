package builder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
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

func TestArgs(t *testing.T) {
	mvn := NewClient("./", nil, []string{"clean install"}, nil)

	assert.Equal(t, "mvn --batch-mode clean install", mvn.Args())
}
func TestModifyArgs(t *testing.T) {
	mvn := NewClient("./", nil, []string{"clean install"}, nil)
	mvn.ModifyArgs([]string{"package"})

	assert.Equal(t, "mvn --batch-mode package", mvn.Args())
}

func TestRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommand("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(false)

	if !strings.Contains(buf.String(), "✔  test-app 1.0.0-SNAPSHOT Complete") {
		t.Fatal("Expected output to have 'app Complete'. Got:", buf.String())
	}
	if !strings.Contains(buf.String(), "✔  jar: target/test-app.jar Complete") {
		t.Fatal("Expected output to have '.jar Complete'. Got:", buf.String())
	}
}

func TestErrorRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommandError("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(false)

	if !strings.Contains(buf.String(), "[ERROR] No goals have been specified for this build.") {
		t.Fatal("Expected output to have 'app Complete'. Got:", buf.String())
	}
}

func TestVerboseRun(t *testing.T) {
	var buf bytes.Buffer

	mvn := Maven{
		cmd: *fakeExecCommand("mvn", "clean", "install"),
		out: &buf,
	}

	mvn.Run(true)

	if !strings.Contains(buf.String(), "[INFO] Building test-app 1.0.0-SNAPSHOT") {
		t.Fatal("Expected output to have 'app Complete'. Got:", buf.String())
	}
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
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
