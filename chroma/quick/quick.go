// Package quick provides simple, no-configuration source code highlighting.
package quick

import (
	"io"

	"github.com/zhangdapeng520/zdpgo_sim/chroma"
	"github.com/zhangdapeng520/zdpgo_sim/chroma/formatters"
	"github.com/zhangdapeng520/zdpgo_sim/chroma/lexers"
	"github.com/zhangdapeng520/zdpgo_sim/chroma/styles"
)

// Highlight some text.
//
// Lexer, formatter and style may be empty, in which case a best-effort is made.
func Highlight(w io.Writer, source, lexer, formatter, style string) error {
	// Determine lexer.
	l := lexers.Get(lexer)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	// Determine formatter.
	f := formatters.Get(formatter)
	if f == nil {
		f = formatters.Fallback
	}

	// Determine style.
	s := styles.Get(style)
	if s == nil {
		s = styles.Fallback
	}

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return f.Format(w, s, it)
}