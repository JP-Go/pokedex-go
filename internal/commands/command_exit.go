package commands

import "os"

func commandExitHandler(_args ...string) error {
	os.Exit(0)
	return nil
}

func createExitHandler() commandCallback {
	return commandExitHandler
}
