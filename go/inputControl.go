package main

import (
	"slices"
	"errors"
	"fmt"
)

var (
	POSSIBLE_CMDS = []string{"workflows_stats"}
)

type Control struct {
	PAT *string
	COMMAND *string
}

func validateControls(control Control) (bool, error) {
	if !slices.Contains(POSSIBLE_CMDS, *control.COMMAND) {
		return false, invalidCommandError()
	}
	return true, nil
}

func invalidCommandError() error {
	return errors.New(fmt.Sprintf("Command is not one of valid commands: %s", POSSIBLE_CMDS))
}