package token

import (
	"fmt"
	"strings"
	"testing"
)

func TestPositionNewlines(t *testing.T) {
	for _, newline := range []string{"\n", "\r", "\r\n"} {
		for _, trailing := range []string{"", newline} {
			t.Run(fmt.Sprintf("%q/%q", newline, trailing), func(t *testing.T) {
				f := &File{Buffer: "SELECT 1" + newline + "!" + trailing}
				offset := Pos(8 + len(newline))
				p := f.Position(offset, offset+1)
				if p.Line != 1 || p.Column != 0 || p.EndLine != 1 || p.EndColumn != 1 {
					t.Fatalf("unexpected position: %+v", p)
				}
				if p.Source != "  2|  !\n   |  ^" {
					t.Errorf("source = %q", p.Source)
				}
				first := f.Position(0, 1)
				if strings.Contains(first.Source, "\r") {
					t.Errorf("source contains CR: %q", first.Source)
				}
				line, col := f.ResolvePos(Pos(len(f.Buffer)))
				wantLine, wantColumn := 1, 1
				if trailing != "" {
					wantLine, wantColumn = 2, 0
				}
				if line != wantLine || col != wantColumn {
					t.Errorf("EOF = %d:%d, want %d:%d", line, col, wantLine, wantColumn)
				}
			})
		}
	}
}
