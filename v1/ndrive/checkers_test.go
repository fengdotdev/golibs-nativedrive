package ndrive_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/golibs-nativedrive/v1/ndrive"

	"github.com/fengdotdev/golibs-testing/assert"
)

func Test_Checkers(t *testing.T) {

	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	err = ndrive.CheckDrive(workingDir)

	assert.NotErrorWithMessage(t, err, "CheckDrive should not return an error for a valid drive")

}
