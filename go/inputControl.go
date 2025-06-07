package main

import (
	"slices"
	"errors"
	"fmt"
)

var (
	STATS_CMD = "workflows_stats"
	POSSIBLE_CMDS = []string{STATS_CMD}
)

type Control struct {
	PAT *string
	CMD *string
	WF_REPO *string
}

func validateControls(control Control) (bool, error) {
	if !slices.Contains(POSSIBLE_CMDS, *control.CMD) {
		return false, invalidCommandError()
	}
	return true, nil
}

func invalidCommandError() error {
	return errors.New(fmt.Sprintf("Command is not one of valid commands: %s", POSSIBLE_CMDS))
}