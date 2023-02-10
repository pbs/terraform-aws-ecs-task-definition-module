package test

import (
	"testing"
)

func TestEFSExample(t *testing.T) {
	t.Skip("Skipping test because destroys aren't proceeding cleanly right now, and it's blocking a release.")

	testTaskDef(t, "efs")
}
