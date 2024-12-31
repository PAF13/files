package files

import (
	"strings"
)

/*
╭───────────┬─────────────────╮
│ prefix    │ C:              │
│ parent    │ C:\Users\viking │
│ stem      │ spam            │
│ extension │ txt             │
╰───────────┴─────────────────╯
*/
type Path struct {
	Prefix    string `json:"prefix"`
	Parent    string `json:"parent"`
	Stem      string `json:"stem"`
	Extension string `json:"extension"`
}

func ParsePath(path string) Path {
	p := Path{}
	if path != "" {
		//prefix
		prefix := strings.Split(path, ":")[0] + ":"
		p.Prefix = prefix

		//parent
		parentParse := strings.Split(path, "\\")
		parentLast := len(parentParse) - 1
		stemext := parentParse[parentLast]
		parent := strings.ReplaceAll(path, stemext, "")
		p.Parent = parent

		//stem - extension
		stemParse := strings.Split(stemext, `.`)
		stemLast := len(stemParse) - 1
		ext := stemParse[stemLast]
		if stemLast < 1 {
			p.Stem = stemext
		} else {
			stem := strings.ReplaceAll(stemext, "."+ext, "")
			p.Stem = stem
			p.Extension = ext
		}
	}
	return p
}
