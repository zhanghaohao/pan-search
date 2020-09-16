package clean

import (
	"testing"
)

func TestClean(t *testing.T) {
	err := Clean()
	if err != nil {
		t.Error(err)
	}
}
