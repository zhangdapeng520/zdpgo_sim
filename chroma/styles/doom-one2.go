package styles

import (
	"github.com/zhangdapeng520/zdpgo_sim/chroma"
)

// Doom One 2 style. Inspired by Atom One and Doom Emacs's Atom One theme
var DoomOne2 = Register(chroma.MustNewStyle("doom-one2", chroma.StyleEntries{
	chroma.Text:                 "#b0c4de",
	chroma.Error:                "#b0c4de",
	chroma.Comment:              "italic #8a93a5",
	chroma.CommentHashbang:      "bold",
	chroma.Keyword:              "#76a9f9",
	chroma.KeywordConstant:      "#e5c07b",
	chroma.KeywordType:          "#e5c07b",
	chroma.Operator:             "#54b1c7",
	chroma.OperatorWord:         "bold #b756ff",
	chroma.Punctuation:          "#abb2bf",
	chroma.Name:                 "#aa89ea",
	chroma.NameAttribute:        "#cebc3a",
	chroma.NameBuiltin:          "#e5c07b",
	chroma.NameClass:            "#ca72ff",
	chroma.NameConstant:         "bold",
	chroma.NameDecorator:        "#e5c07b",
	chroma.NameEntity:           "#bda26f",
	chroma.NameException:        "bold #fd7474",
	chroma.NameFunction:         "#00b1f7",
	chroma.NameProperty:         "#cebc3a",
	chroma.NameLabel:            "#f5a40d",
	chroma.NameNamespace:        "#ca72ff",
	chroma.NameTag:              "#76a9f9",
	chroma.NameVariable:         "#DCAEEA",
	chroma.NameVariableClass:    "#DCAEEA",
	chroma.NameVariableGlobal:   "bold #DCAEEA",
	chroma.NameVariableInstance: "#e06c75",
	chroma.NameVariableMagic:    "#DCAEEA",
	chroma.Literal:              "#98c379",
	chroma.LiteralDate:          "#98c379",
	chroma.Number:               "#d19a66",
	chroma.NumberBin:            "#d19a66",
	chroma.NumberFloat:          "#d19a66",
	chroma.NumberHex:            "#d19a66",
	chroma.NumberInteger:        "#d19a66",
	chroma.NumberIntegerLong:    "#d19a66",
	chroma.NumberOct:            "#d19a66",
	chroma.String:               "#98c379",
	chroma.StringAffix:          "#98c379",
	chroma.StringBacktick:       "#98c379",
	chroma.StringDelimiter:      "#98c379",
	chroma.StringDoc:            "#7e97c3",
	chroma.StringDouble:         "#63c381",
	chroma.StringEscape:         "bold #d26464",
	chroma.StringHeredoc:        "#98c379",
	chroma.StringInterpol:       "#98c379",
	chroma.StringOther:          "#70b33f",
	chroma.StringRegex:          "#56b6c2",
	chroma.StringSingle:         "#98c379",
	chroma.StringSymbol:         "#56b6c2",
	chroma.Generic:              "#b0c4de",
	chroma.GenericDeleted:       "#b0c4de",
	chroma.GenericEmph:          "italic",
	chroma.GenericHeading:       "bold #a2cbff",
	chroma.GenericInserted:      "#a6e22e",
	chroma.GenericOutput:        "#a6e22e",
	chroma.GenericUnderline:     "underline",
	chroma.GenericPrompt:        "#a6e22e",
	chroma.GenericStrong:        "bold",
	chroma.GenericSubheading:    "#a2cbff",
	chroma.GenericTraceback:     "#a2cbff",
	chroma.Background:           "#b0c4de bg:#282c34",
}))
