package memefish_test

import (
	"testing"

	"github.com/cloudspannerecosystem/memefish"
	"github.com/cloudspannerecosystem/memefish/ast"
)

func TestPropertyGraphDerivedPropertyListRange(t *testing.T) {
	const sql = "CREATE PROPERTY GRAPH g NODE TABLES (T PROPERTIES (id))"
	node, err := memefish.ParseDDL("", sql)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	ast.Inspect(node, func(node ast.Node) bool {
		if properties, ok := node.(*ast.PropertyGraphDerivedPropertyList); ok {
			found = true
			if got := sql[properties.Pos():properties.End()]; got != "PROPERTIES (id)" {
				t.Errorf("source range = %q", got)
			}
		}
		return true
	})
	if !found {
		t.Fatal("derived property list not found")
	}
}
