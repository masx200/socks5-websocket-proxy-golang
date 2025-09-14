package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
func TestRun(t *testing.T) {
	RunMainWebSocket(t)

	RunMainDOH(t)
	RunMainDEFAULT(t)
}
