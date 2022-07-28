package styles

import (
	"github.com/zhangdapeng520/zdpgo_sim/chroma"
)

// Abap style.
var Abap = Register(chroma.MustNewStyle("abap", chroma.StyleEntries{
	chroma.Comment:        "italic #888",
	chroma.CommentSpecial: "#888",
	chroma.Keyword:        "#00f",
	chroma.OperatorWord:   "#00f",
	chroma.Name:           "#000",
	chroma.LiteralNumber:  "#3af",
	chroma.LiteralString:  "#5a2",
	chroma.Error:          "#F00",
	chroma.Background:     " bg:#ffffff",
}))
