package memefish_test

import "testing"

func TestGoldenUpdateSelection(t *testing.T) {
	for _, tt := range []struct {
		name      string
		update    bool
		run, skip string
		wantError bool
	}{
		{name: "full update", update: true},
		{name: "normal filtered test", run: "TestParseExpr"},
		{name: "normal skip", skip: "slow"},
		{name: "filtered update", update: true, run: "TestParseExpr", wantError: true},
		{name: "skipped update", update: true, skip: "bad", wantError: true},
		{name: "explicit match-all", update: true, run: ".*", wantError: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkGoldenUpdateSelection(tt.update, tt.run, tt.skip); (err != nil) != tt.wantError {
				t.Fatalf("error = %v, want error = %v", err, tt.wantError)
			}
		})
	}
}
