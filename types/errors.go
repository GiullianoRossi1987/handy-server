package types

import "fmt"

type UnexpectedDBChangeBehaviourError struct {
	Operation            string
	Table                string
	ExpectedChangedLines int
	ChangedLines         int
}

func (e *UnexpectedDBChangeBehaviourError) Error() string {
	return fmt.Sprintf(
		"There were unexpected changes made to the table %s at %s. %d lines changed, expected %d lines",
		e.Table, e.Operation, e.ChangedLines, e.ExpectedChangedLines)
}
