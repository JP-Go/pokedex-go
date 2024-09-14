package commands

import "os"

func CommandExitHandler() error {
	os.Exit(0)
	return nil
}
