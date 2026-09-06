package memefish

import (
	"reflect"
	"testing"

	"github.com/cloudspannerecosystem/memefish/ast"
	"github.com/cloudspannerecosystem/memefish/token"
)

func newLookaheadTestParser(sql string) *Parser {
	p := &Parser{Lexer: &Lexer{File: &token.File{Buffer: sql}}}
	p.nextToken()
	return p
}

func assertParserCheckpoint(t *testing.T, p *Parser, lexer *Lexer, errors []*Error) {
	t.Helper()
	if !reflect.DeepEqual(p.Lexer, lexer) {
		t.Errorf("lexer state = %#v, want %#v", p.Lexer, lexer)
	}
	if !reflect.DeepEqual(p.errors, errors) {
		t.Errorf("diagnostics = %v, want %v", p.errors, errors)
	}
}

func TestParserLookaheadOutcomes(t *testing.T) {
	for _, method := range []string{"lookahead", "tryParse"} {
		for _, result := range []bool{false, true} {
			for _, advance := range []bool{false, true} {
				for _, existing := range []bool{false, true} {
					t.Run(method+"/"+boolName(result)+"/advance="+boolName(advance)+"/errors="+boolName(existing), func(t *testing.T) {
						p := newLookaheadTestParser("a.b /* comment */ + 1")
						if existing {
							// Spare capacity exercises diagnostic appends sharing storage.
							p.errors = make([]*Error, 1, 4)
							p.errors[0] = &Error{Message: "existing"}
						}
						beforeLexer, beforeErrors := p.cloneLexer(), p.errors
						var afterLexer *Lexer
						var afterErrors []*Error
						callback := func() bool {
							if advance {
								p.nextToken()
							}
							p.errors = append(p.errors, &Error{Message: "trial"})
							afterLexer, afterErrors = p.cloneLexer(), p.errors
							return result
						}
						run := p.lookahead
						if method == "tryParse" {
							run = p.tryParse
						}
						if got := run(callback); got != result {
							t.Fatalf("result = %v, want %v", got, result)
						}
						if method == "tryParse" && result {
							assertParserCheckpoint(t, p, afterLexer, afterErrors)
						} else {
							assertParserCheckpoint(t, p, beforeLexer, beforeErrors)
						}
					})
				}
			}
		}
	}
}

func boolName(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func TestParserLookaheadPanic(t *testing.T) {
	for _, method := range []string{"lookahead", "tryParse"} {
		t.Run(method, func(t *testing.T) {
			p := newLookaheadTestParser("a.b")
			before := p.cloneLexer()
			// Both recognized parse errors and unexpected panic values must escape.
			for _, sentinel := range []any{&Error{Message: "parse failure"}, new(int)} {
				func() {
					defer func() {
						if got := recover(); got != sentinel {
							t.Errorf("panic = %v, want original %v", got, sentinel)
						}
					}()
					run := p.lookahead
					if method == "tryParse" {
						run = p.tryParse
					}
					run(func() bool {
						p.nextToken()
						p.errors = append(p.errors, &Error{Message: "trial"})
						panic(sentinel)
					})
				}()
				assertParserCheckpoint(t, p, before, nil)
			}
		})
	}
}

func TestParserLookaheadNested(t *testing.T) {
	for _, outer := range []string{"lookahead", "tryParse"} {
		for _, inner := range []string{"lookahead", "tryParse"} {
			for _, outerResult := range []bool{false, true} {
				for _, innerResult := range []bool{false, true} {
					t.Run(outer+"/"+inner+"/"+boolName(outerResult)+"/"+boolName(innerResult), func(t *testing.T) {
						p := newLookaheadTestParser("a + b")
						before := p.cloneLexer()
						var parentLexer *Lexer
						var parentErrors []*Error
						runOuter := p.lookahead
						if outer == "tryParse" {
							runOuter = p.tryParse
						}
						runOuter(func() bool {
							p.nextToken()
							p.errors = append(p.errors, &Error{Message: "outer"})
							innerLexer, innerErrors := p.cloneLexer(), p.errors
							runInner := p.lookahead
							if inner == "tryParse" {
								runInner = p.tryParse
							}
							runInner(func() bool {
								p.nextToken()
								p.errors = append(p.errors, &Error{Message: "inner"})
								return innerResult
							})
							if inner != "tryParse" || !innerResult {
								assertParserCheckpoint(t, p, innerLexer, innerErrors)
							}
							parentLexer, parentErrors = p.cloneLexer(), p.errors
							return outerResult
						})
						if outer == "tryParse" && outerResult {
							assertParserCheckpoint(t, p, parentLexer, parentErrors)
						} else {
							assertParserCheckpoint(t, p, before, nil)
						}
					})
				}
			}
		}
	}
}

func TestParserLookaheadRecoveredParse(t *testing.T) {
	p := newLookaheadTestParser("1 + )")
	before := p.cloneLexer()
	if p.tryParse(func() bool {
		expr := p.parseExpr()
		if _, ok := expr.(*ast.BadExpr); !ok || len(p.errors) == 0 {
			t.Fatalf("expected recovered expression with diagnostics, got %T", expr)
		}
		// A normal return with a recovery node does not accept the trial.
		return false
	}) {
		t.Fatal("recovered trial unexpectedly accepted")
	}
	assertParserCheckpoint(t, p, before, nil)
	p.parseExpr()
	if len(p.errors) != 1 {
		t.Fatalf("reparse produced %d diagnostics, want 1", len(p.errors))
	}
}

func TestParserLookaheadNestedLexicalPanic(t *testing.T) {
	p := newLookaheadTestParser("a + 'unterminated")
	before := p.cloneLexer()
	func() {
		defer func() {
			if _, ok := recover().(*Error); !ok {
				t.Error("expected lexical error to escape both checkpoints")
			}
		}()
		p.lookahead(func() bool {
			p.nextToken()
			p.errors = append(p.errors, &Error{Message: "outer"})
			innerLexer, innerErrors := p.cloneLexer(), p.errors
			defer func() {
				// Inner restoration must run before outer restoration.
				assertParserCheckpoint(t, p, innerLexer, innerErrors)
			}()
			return p.tryParse(func() bool {
				p.errors = append(p.errors, &Error{Message: "inner"})
				p.nextToken()
				return true
			})
		})
	}()
	assertParserCheckpoint(t, p, before, nil)
}

func TestParserLookaheadTokenReplay(t *testing.T) {
	for _, sql := range []string{"a.SELECT /* comment */ + 1", "/* first */ `名前`.field", "a", "a /* unterminated"} {
		t.Run(sql, func(t *testing.T) {
			p := newLookaheadTestParser(sql)
			p.Lexer.nextToken(true) // Includes dot-identifier mode and bad tokens.
			before := p.cloneLexer()
			p.lookahead(func() bool {
				for range 4 {
					p.Lexer.nextToken(true)
				}
				return true
			})
			assertParserCheckpoint(t, p, before, nil)
			for range 4 {
				p.Lexer.nextToken(true)
				before.nextToken(true)
				assertParserCheckpoint(t, p, before, nil)
			}
		})
	}
}
