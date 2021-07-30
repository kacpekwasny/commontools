package commontools

import (
	"bytes"
	"os/exec"
)

// From:
// https://stackoverflow.com/a/43246464/12069832

// ShellToUse set your shell
var ShellToUse = "bash"

// Shellout - to run a string as command in your preffered shell,
// set ShellToUse to ex.: "sh", default is "bash"
func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
