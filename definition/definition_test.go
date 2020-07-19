package definiton

import (
	"testing"
)

func TestDefinition(t *testing.T) {

	want := "1.2.3"
	def := New("1.2.3")
	if got := def.Version(); got != want {
		t.Errorf("Version() = %q, want %q", got, want)
	}
}
