package internal

import (
	"strings"
)

func ReplaceSpecialCharWithTexSymbol(text string) string {
	r := strings.NewReplacer("_", "\\_",
		"%", "\\%",
		"#", "\\#",
		"&", "\\&",
		"$", "\\$",
		"%", "\\%",
		"{", "\\{",
		"}", "\\}",
		"~", "\\textasciitilde",
		"^", "\\textasciicircum",
		"\\", "\\textbackslash",
	)
	return r.Replace(text)
}
