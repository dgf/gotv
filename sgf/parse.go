package sgf

import (
	"bytes"
	"regexp"
	"strings"
	"text/scanner"
)

var (
	lineFeeds  = []string{"\r\n", "\n\r", "\r"}
	spaceRunes = regexp.MustCompile(`[ \t\f]+`)
	spaceList  = regexp.MustCompile(`[ ]+`)
	trimLines  = regexp.MustCompile(`[ ]*\n[ ]*`)
)

func Parse(sgf string) Collection {
	ctx := struct {
		Collection
		*Tree
		*Node
		Ident string
	}{}

	// replace all line feed combinations with a single one
	for _, f := range lineFeeds {
		sgf = strings.Replace(sgf, f, "\n", -1)
	}

	// combine multiple space to a single one
	sgf = spaceRunes.ReplaceAllString(sgf, " ")

	// setup
	s := scanner.Scanner{}
	s.Init(strings.NewReader(sgf))
	s.Mode = scanner.ScanIdents | scanner.ScanComments

	var r rune
	for {
		if r == scanner.EOF || s.Peek() == scanner.EOF {
			break // this is the end
		}

		r = s.Scan()
		switch r {

		case '(': // sequence
			if ctx.Tree == nil { // root
				ctx.Tree = &Tree{}
				ctx.Collection = append(ctx.Collection, ctx.Tree)
			} else { // subtree > go down
				t := &Tree{Parent: ctx.Tree}
				ctx.Tree.Collection = append(ctx.Tree.Collection, t)
				ctx.Tree = t
			}

		case ';': // node
			ctx.Node = &Node{Properties: map[string]string{}}
			ctx.Tree.Sequence = append(ctx.Tree.Sequence, ctx.Node)

		case ')': // end > go one up
			ctx.Node = nil
			ctx.Tree = ctx.Tree.Parent
		}

		if ctx.Node != nil && r == scanner.Ident { // property ident
			ctx.Ident = s.TokenText()

			// skip until [
			for {
				if s.Peek() == scanner.EOF {
					break // this is the end
				}
				if s.Scan() == '[' {
					if "C" == ctx.Ident {
						s.Whitespace = 0 // skip nothing in comments
					} else {
						s.Whitespace = 1 << '\n' // skip line feeds in value
					}
					break
				}
			}

			// property value (loop until ']')
			v := bytes.Buffer{}
			for {
				r = s.Scan()
				if r == scanner.EOF || r == ']' {
					break // this is the end
				}
				switch r {
				case scanner.Ident:
					v.WriteString(s.TokenText())
				case scanner.Comment:
					v.WriteString(s.TokenText())
				case '\\':
					v.WriteRune(r)       // write \
					if s.Peek() == ']' { // check escaped end "\]"
						s.Scan() // write ]
						v.WriteString(s.TokenText())
					}
				default:
					v.WriteString(s.TokenText())
				}
			}

			// clean up and add property value
			pv := spaceList.ReplaceAllString(v.String(), " ")
			pv = trimLines.ReplaceAllString(pv, "\n")
			pv = strings.Trim(pv, " \n")
			ctx.Node.Properties[ctx.Ident] = pv

			// reset white spaces ignore
			s.Whitespace = scanner.GoWhitespace
		}
	}

	return ctx.Collection
}
