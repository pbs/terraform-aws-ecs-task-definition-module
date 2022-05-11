package test

import (
	"net/http"
	"testing"
)

func TestEFSExample(t *testing.T) {
	// Checks to see if the EFS repo is open sourced yet
	resp, err := http.Get("https://github.com/pbs/terraform-aws-efs-module")

	if err != nil || resp.StatusCode != 200 {
		t.Skip("Skipping test because EFS repo is not open sourced yet")
	}

	testTaskDef(t, "efs")
}
