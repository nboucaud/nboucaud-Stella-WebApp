package utils

import (
	"regexp"
	"strings"
)

var (
	openCodeBlockRgx  = regexp.MustCompile("`{3}[ 	\\w\\d]*\\n")
	closeCodeBlockRgx = regexp.MustCompile("`{3}")

	backtickRgx      = regexp.MustCompile("`*(.*?)`*")
	asterisksRgx     = regexp.MustCompile(`\*+(.*?)\*+`)
	underscoreRgx    = regexp.MustCompile("_+(.*?)_+")
	strikethroughRgx = regexp.MustCompile("~+(.*?)~+")
	blockQuoteRgx    = regexp.MustCompile("> (.*)")
	headerRgx        = regexp.MustCompile("#+ (.*)")
	listRgx          = regexp.MustCompile(`(?m)^([ ]?[-*+]) (.*)(\n?)`)
	numberingRgx     = regexp.MustCompile(`(\d+)\. (.*)(\n?)`)
	linkRgx          = regexp.MustCompile(`!?\[(.*?)\][\[\(].*?[\]\)]`)
	newLineRgx       = regexp.MustCompile(`\n+`)
	whitespaceRgx    = regexp.MustCompile(`(\s+)`)
	tableRgx         = regexp.MustCompile(`((\r?\n){2}|^)([^\r\n]*\|[^\r\n]*(\r?\n)?)+(=(\r?\n){2}|$)`)
)

func StripMarkdown(text string) string {
	res := text
	// TODO: Code block regex
	res = openCodeBlockRgx.ReplaceAllString(res, " ")
	res = closeCodeBlockRgx.ReplaceAllString(res, "$1")

	res = tableRgx.ReplaceAllString(res, "")

	res = backtickRgx.ReplaceAllString(res, "$1")
	res = asterisksRgx.ReplaceAllString(res, "$1")
	res = underscoreRgx.ReplaceAllString(res, "$1")
	res = strikethroughRgx.ReplaceAllString(res, "$1")
	res = blockQuoteRgx.ReplaceAllString(res, "$1")
	res = headerRgx.ReplaceAllString(res, "$1")
	res = listRgx.ReplaceAllString(res, "$2$3")
	res = numberingRgx.ReplaceAllString(res, "$2$3")
	res = linkRgx.ReplaceAllString(res, "$1")

	res = newLineRgx.ReplaceAllString(res, " ")
	res = whitespaceRgx.ReplaceAllString(res, " ")

	return strings.TrimSpace(res)
}
