package sgf

import (
	"bytes"
	"regexp"
	"strings"
	"text/scanner"
)

var (
	replacer  = strings.NewReplacer("\r", "\n", "\t", " ", "\f", " ")
	spaceList = regexp.MustCompile(`[ ]+`)
	trimLines = regexp.MustCompile(`([ ]*\n[ ]*)+`)
)

// property value scan (loop until ']')
func scanValue(s *scanner.Scanner) (v string) {
	b := bytes.Buffer{}

	for {
		r := s.Scan()
		if r == scanner.EOF || r == ']' {
			break // this is the end
		}
		switch r {
		case scanner.Ident:
			b.WriteString(s.TokenText())
		case scanner.Comment:
			b.WriteString(s.TokenText())
		case '\\':
			b.WriteRune(r)       // write \
			if s.Peek() == ']' { // check escaped end "\]"
				s.Scan() // write ]
				b.WriteString(s.TokenText())
			}
		default:
			b.WriteString(s.TokenText())
		}
	}

	return b.String()
}

// Parse SGF
func Parse(sgf string) Collection {
	var (
		collection Collection
		tree       *Tree
		node       Node
		ident      string
	)

	// setup scanner (only indent and comments match)
	s := &scanner.Scanner{}
	s.Init(strings.NewReader(replacer.Replace(sgf)))
	s.Mode = scanner.ScanIdents | scanner.ScanComments

	var r rune
	for {
		if r == scanner.EOF || s.Peek() == scanner.EOF {
			break // this is the end
		}

		r = s.Scan()
		switch r {

		case '(': // sequence
			if tree == nil { // root
				tree = &Tree{}
				collection = append(collection, tree)
			} else { // subtree > go down
				t := &Tree{Parent: tree}
				tree.Collection = append(tree.Collection, t)
				tree = t
			}

		case ';': // node
			node = Node{}
			tree.Sequence = append(tree.Sequence, node)

		case ')': // end > go one up
			node = nil
			tree = tree.Parent
		}

		// property ident
		if node != nil && r == scanner.Ident {
			ident = s.TokenText()

			// skip until [
			for {
				if s.Peek() == scanner.EOF {
					break // this is the end
				}
				if s.Scan() == '[' {
					break
				}
			}

			// scan property value
			if "C" == ident {
				s.Whitespace = 0 // skip nothing in comments
				v := scanValue(s)
				v = strings.Trim(v, " \n")
				v = spaceList.ReplaceAllString(v, " ")
				v = trimLines.ReplaceAllString(v, "\n")
				node[ident] = v
			} else {
				s.Whitespace = 1 << '\n' // skip line feeds in value
				v := scanValue(s)
				v = strings.Trim(v, " ")
				v = spaceList.ReplaceAllString(v, " ")
				node[ident] = v
			}

			// reset white spaces ignore
			s.Whitespace = scanner.GoWhitespace
		}
	}

	return collection
}
