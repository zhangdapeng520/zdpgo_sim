package b

import (
	. "github.com/zhangdapeng520/zdpgo_sim/chroma" // nolint
	"github.com/zhangdapeng520/zdpgo_sim/chroma/lexers/internal"
)

// BashSession lexer.
var BashSession = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "BashSession",
		Aliases:   []string{"bash-session", "console", "shell-session"},
		Filenames: []string{".sh-session"},
		MimeTypes: []string{"text/x-sh"},
		EnsureNL:  true,
	},
	bashsessionRules,
))

func bashsessionRules() Rules {
	return Rules{
		"root": {
			{`^((?:\[[^]]+@[^]]+\]\s?)?[#$%>])(\s*)(.*\n?)`, ByGroups(GenericPrompt, Text, Using(Bash)), nil},
			{`^.+\n?`, GenericOutput, nil},
		},
	}
}
