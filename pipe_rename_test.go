package memefish_test

import (
	"testing"

	"github.com/cloudspannerecosystem/memefish"
)

func TestParseStatementsPipeRenameTrailingCommaBeforeSemicolon(t *testing.T) {
	stmts, err := memefish.ParseStatements("", "FROM Singers |> RENAME FirstName given_name,;")
	if err != nil {
		t.Fatalf("ParseStatements() error = %v, want nil", err)
	}
	if len(stmts) != 1 {
		t.Fatalf("ParseStatements() statements = %d, want 1", len(stmts))
	}
	if got := stmts[0].SQL(); got != "FROM Singers |> RENAME FirstName given_name" {
		t.Errorf("ParseStatements() SQL() = %q, want %q", got, "FROM Singers |> RENAME FirstName given_name")
	}
}
