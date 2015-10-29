// 通过exec执行命令
package tsample

import (
	"fmt"
	"os"
	"os/exec"
)

func testCmd() {
	var cmd = exec.Command("git", "diff")
	var err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "command err: %s\n", err)
		return
	}

	cmd = exec.Command("git", "diff")
	buff, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "command err: %s\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "command output: %s\n", buff)
}

func TestCmd() {
	testCmd()
}
