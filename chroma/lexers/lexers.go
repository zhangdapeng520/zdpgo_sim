// Package lexers contains the registry of all lexers.
//
// Sub-packages contain lexer implementations.
package lexers

// nolint
import (
	"github.com/zhangdapeng520/zdpgo_sim/chroma"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/a"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/b"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/c"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/circular"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/d"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/e"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/f"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/g"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/h"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/i"
	"github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/internal"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/j"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/k"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/l"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/m"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/n"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/o"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/p"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/q"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/r"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/s"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/t"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/v"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/w"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/x"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/y"
	_ "github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/z"
)

// Registry of Lexers.
var Registry = internal.Registry

// Names of all lexers, optionally including aliases.
func Names(withAliases bool) []string { return internal.Names(withAliases) }

// Get a Lexer by name, alias or file extension.
func Get(name string) chroma.Lexer { return internal.Get(name) }

// MatchMimeType attempts to find a lexer for the given MIME type.
func MatchMimeType(mimeType string) chroma.Lexer { return internal.MatchMimeType(mimeType) }

// Match returns the first lexer matching filename.
func Match(filename string) chroma.Lexer { return internal.Match(filename) }

// Analyse text content and return the "best" lexer..
func Analyse(text string) chroma.Lexer { return internal.Analyse(text) }

// Register a Lexer with the global registry.
func Register(lexer chroma.Lexer) chroma.Lexer { return internal.Register(lexer) }

// Fallback lexer if no other is found.
var Fallback = internal.Fallback
