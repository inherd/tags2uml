package pkg

import "testing"

func TestParseMembersMethods(t *testing.T) {
	tests := []struct {
		name string
		args string
	}{
		{"name", "../_fixtures/ctags/deadevent.ctags"},
		{"name", "../_fixtures/ctags/eventbus"},
	}

	InitDatastore()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseClass(tt.args)
			ParseMembersMethods(tt.args)
		})
	}
}
