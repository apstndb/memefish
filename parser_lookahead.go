package memefish

// parserCheckpoint saves lexer state and the diagnostic append boundary.
// Token comments and existing errors are shared: callbacks must not modify
// them, the source file, or previously built ASTs. Source position caches may
// be populated, since they do not change parsing behavior.
type parserCheckpoint struct {
	lexer  *Lexer
	errors []*Error
}

func (p *Parser) checkpoint() parserCheckpoint {
	return parserCheckpoint{lexer: p.cloneLexer(), errors: p.errors}
}

func (p *Parser) restore(checkpoint parserCheckpoint) {
	p.Lexer = checkpoint.lexer
	p.errors = checkpoint.errors
}

// lookahead runs check through p and restores lexer state and diagnostics,
// including on panic. Panics propagate unchanged.
// Callbacks may only append diagnostics, not modify existing ones or call
// public Parse* entry points, which drain diagnostics. Captured values and
// other external side effects are not rolled back.
func (p *Parser) lookahead(check func() bool) bool {
	checkpoint := p.checkpoint()
	defer p.restore(checkpoint)
	return check()
}

// tryParse retains lexer state and diagnostics only if parse returns non-nil.
// Otherwise it restores both, including on panic; panics propagate unchanged.
// The callback has the same restrictions as lookahead and must explicitly
// decide whether to accept the interpretation: recovery is not necessarily a
// match. An outer checkpoint can still undo an inner successful tryParse.
// This contract does not apply to all optional tryParse* methods.
// The pointer result avoids ambiguous typed-nil interface values. Parsers that
// return an interface or a value plus a separate match flag need another shape.
func tryParse[T any](p *Parser, parse func() *T) (result *T) {
	checkpoint := p.checkpoint()
	defer func() {
		if result == nil {
			p.restore(checkpoint)
		}
	}()
	return parse()
}
