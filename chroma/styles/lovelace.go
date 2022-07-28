package styles

import (
	"github.com/zhangdapeng520/zdpgo_sim/chroma"
)

// Lovelace style.
var Lovelace = Register(chroma.MustNewStyle("lovelace", chroma.StyleEntries{
	chroma.TextWhitespace:         "#a89028",
	chroma.Comment:                "italic #888888",
	chroma.CommentHashbang:        "#287088",
	chroma.CommentMultiline:       "#888888",
	chroma.CommentPreproc:         "noitalic #289870",
	chroma.Keyword:                "#2838b0",
	chroma.KeywordConstant:        "italic #444444",
	chroma.KeywordDeclaration:     "italic",
	chroma.KeywordType:            "italic",
	chroma.Operator:               "#666666",
	chroma.OperatorWord:           "#a848a8",
	chroma.Punctuation:            "#888888",
	chroma.NameAttribute:          "#388038",
	chroma.NameBuiltin:            "#388038",
	chroma.NameBuiltinPseudo:      "italic",
	chroma.NameClass:              "#287088",
	chroma.NameConstant:           "#b85820",
	chroma.NameDecorator:          "#287088",
	chroma.NameEntity:             "#709030",
	chroma.NameException:          "#908828",
	chroma.NameFunction:           "#785840",
	chroma.NameFunctionMagic:      "#b85820",
	chroma.NameLabel:              "#289870",
	chroma.NameNamespace:          "#289870",
	chroma.NameTag:                "#2838b0",
	chroma.NameVariable:           "#b04040",
	chroma.NameVariableGlobal:     "#908828",
	chroma.NameVariableMagic:      "#b85820",
	chroma.LiteralString:          "#b83838",
	chroma.LiteralStringAffix:     "#444444",
	chroma.LiteralStringChar:      "#a848a8",
	chroma.LiteralStringDelimiter: "#b85820",
	chroma.LiteralStringDoc:       "italic #b85820",
	chroma.LiteralStringEscape:    "#709030",
	chroma.LiteralStringInterpol:  "underline",
	chroma.LiteralStringOther:     "#a848a8",
	chroma.LiteralStringRegex:     "#a848a8",
	chroma.LiteralNumber:          "#444444",
	chroma.GenericDeleted:         "#c02828",
	chroma.GenericEmph:            "italic",
	chroma.GenericError:           "#c02828",
	chroma.GenericHeading:         "#666666",
	chroma.GenericSubheading:      "#444444",
	chroma.GenericInserted:        "#388038",
	chroma.GenericOutput:          "#666666",
	chroma.GenericPrompt:          "#444444",
	chroma.GenericStrong:          "bold",
	chroma.GenericTraceback:       "#2838b0",
	chroma.GenericUnderline:       "underline",
	chroma.Error:                  "bg:#a848a8",
	chroma.Background:             " bg:#ffffff",
}))
