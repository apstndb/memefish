package memefish_test

import (
	"testing"

	"github.com/cloudspannerecosystem/memefish"
)

func TestParseStatementsPipeAggregateTrailingCommaBeforeSemicolon(t *testing.T) {
	for _, test := range []struct {
		input string
		want  string
	}{
		{
			input: "FROM Singers |> AGGREGATE COUNT(*) AS count,;",
			want:  "FROM Singers |> AGGREGATE COUNT(*) AS count",
		},
		{
			input: "FROM Singers |> AGGREGATE COUNT(*) AS count GROUP BY FirstName,;",
			want:  "FROM Singers |> AGGREGATE COUNT(*) AS count GROUP BY FirstName",
		},
	} {
		stmts, err := memefish.ParseStatements("", test.input)
		if err != nil {
			t.Fatalf("ParseStatements(%q) error = %v, want nil", test.input, err)
		}
		if len(stmts) != 1 {
			t.Fatalf("ParseStatements(%q) statements = %d, want 1", test.input, len(stmts))
		}
		if got := stmts[0].SQL(); got != test.want {
			t.Errorf("ParseStatements(%q) SQL() = %q, want %q", test.input, got, test.want)
		}
	}
}
