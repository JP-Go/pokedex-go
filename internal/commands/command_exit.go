package commands

import "os"

func CommandExitHandler(_args ...string) error {
	os.Exit(0)
	return nil
}
