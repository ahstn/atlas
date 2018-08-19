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

func TestInitaliseCommand(t *testing.T) {
	pathenv := os.Getenv("PATH")
	defer os.Setenv("PATH", pathenv)

	path := SetupMvnInPath(t)

	var mvn Maven
	mvn.initialiseCommand()
	if mvn.cmd.Path != path {
		t.Fatal("Expected cmd path to be mvn. Got:", mvn.cmd.Path)
	}
}

func TestShorthandTasks(t *testing.T) {
	pathenv := os.Getenv("PATH")
	defer os.Setenv("PATH", pathenv)

	SetupMvnInPath(t)

	var mvn Maven
	mvn.initialiseCommand()
	mvn.Clean()
	mvn.Build()
	mvn.Package()
	mvn.SkipTests()

	cmd := strings.Join(mvn.cmd.Args, " ")
	if !strings.Contains(cmd, "clean") {
		t.Fatal("Expected cmd to contain arg 'clean'. Got:", cmd)
	} else if !strings.Contains(cmd, "install") {
		t.Fatal("Expected cmd to contain arg 'install'. Got:", cmd)
	} else if !strings.Contains(cmd, "package") {
		t.Fatal("Expected cmd to contain arg 'package'. Got:", cmd)
	} else if !strings.Contains(cmd, "-DskipTests") {
		t.Fatal("Expected cmd to contain arg '-DskipTests'. Got:", cmd)
	}
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

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)
	fmt.Print(fakeOutput)
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
