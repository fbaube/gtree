package gtree

//  gopkg.in/russross/blackfriday.v2
// github.com/rhinoman/go-commonmark
// github.com/rhinoman/go-commonmark

/* old code

func GetTitleFromMarkdownTokens(tokens []MD.Token) string {
	if tokens == nil || len(tokens) == 0 {
		return ""
	}
	var title = "UNTITLED"
	if heading, ok := tokens[0].(*MD.HeadingOpen); ok {
		for i := 1; i < len(tokens); i++ {
			if tok, ok := tokens[i].(*MD.HeadingClose); ok && tok.Lvl == heading.Lvl {
				break
			}
			title += MD.ExtractText(tokens[i])
		}
		title = strings.TrimSpace(title)
	}
	return title
}

*/
