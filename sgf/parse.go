package sgf

import (
	"bytes"
	"regexp"
	"strings"
	"text/scanner"
)

var (
	spaceRunes = regexp.MustCompile(`[ \t\f]+`)
	spaceList  = regexp.MustCompile(`[ ]+`)
	trimLines  = regexp.MustCompile(`[ ]*\n[ ]*`)
)

func Parse(sgf string) Collection {
	c := struct {
		Collection
		*GameTree
		*Node
		PropIdent string
	}{}

	// replace all line feed combinations with a single one
	for _, r := range []string{"\r\n", "\n\r", "\r"} {
		sgf = strings.Replace(sgf, r, "\n", -1)
	}

	// combine multiple space to a single one
	sgf = spaceRunes.ReplaceAllString(sgf, " ")

	// setup
	s := scanner.Scanner{}
	s.Filename = "sgf"
	s.Init(strings.NewReader(sgf))
	s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanComments

	var t rune
	for {
		if t == scanner.EOF || s.Peek() == scanner.EOF {
			break // this is the end
		}

		t = s.Scan()
		switch t {
		case '(': // sequence
			if c.GameTree == nil { // root
				c.GameTree = &GameTree{}
				c.Collection = append(c.Collection, c.GameTree)
			} else { // subtree > go down
				g := &GameTree{Parent: c.GameTree}
				c.GameTree.Collection = append(c.GameTree.Collection, g)
				c.GameTree = g
			}

		case ';': // node
			c.Node = &Node{Properties: map[string]string{}}
			c.GameTree.Sequence = append(c.GameTree.Sequence, c.Node)

		case ')': // end > go one up
			c.Node = nil
			c.GameTree = c.GameTree.Parent
		}

		if c.Node != nil && t == scanner.Ident { // property ident
			c.PropIdent = s.TokenText()

			// skip until [
			for {
				if s.Peek() == scanner.EOF {
					break // this is the end
				}
				if s.Scan() == '[' {
					if "C" == c.PropIdent {
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
				t = s.Scan()
				if t == scanner.EOF || t == ']' {
					break // this is the end
				}
				switch t {
				case scanner.Ident:
					v.WriteString(s.TokenText())
				case scanner.Int:
					v.WriteString(s.TokenText())
				case scanner.Float:
					v.WriteString(s.TokenText())
				case scanner.Comment:
					v.WriteString(s.TokenText())
				case '\\':
					v.WriteRune(t)       // write \
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
			c.Node.Properties[c.PropIdent] = pv

			// reset white spaces ignore
			s.Whitespace = scanner.GoWhitespace
		}
	}

	return c.Collection
}
