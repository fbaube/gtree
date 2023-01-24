package gtree

import (
	"io"
	// "github.com/dimchansky/utfbom"
)

// GTree is the workspace for AND the results of parsing a hierarchically
// organised markup file. The file should *not* have to be XML, but *should*
// have (or be modelable with) hierarchical (tree) structure.
//
// Currently an GTree contains XML. There is not (yet) any higher-order
// semantics imposed or added, but it is entirely possible that an GTree
// could instead be base don (say) the Pandoc AST.
//
// GTree maintains a 1-to-1 mapping btwn the tokens returned by the Golang
// XML parser, and its own "Tag" elements. This makes it easy to sort out
// errors, and to provide meaningful error messages that directly quote inputs.
//
// NOTE:1050 that the file does not have to be a well-formed XML file (or
// other markup file) with a single root element. It can also be
// - A DTD file (*.dtd, *.mod)
// - An XML data file that happens not to have a single top-level
// root element (this makes it an "XML fragment")
//
// Thus this function makes no assumptions about the top-down structure
// of the XML file, but it does expect that it is basically well-formed. -n-
// The file is read entirely into memory and parsed as a unit, in several
// passes, which implies that
// - an GTree is returned as the complete end result of parsing a single
// file, and no intermediate results are exposed to the caller (altho after
// the parsing function returns, the total output of each pass is available
// as fields in the GTree struct)
// - every transcluded file (i.e. external entity reference - of type general
// "&foo;" or parameter "%foo;") is processed as a separate GTree and then
// merged (i.e. transcluded) into the file's GTree
//
// Each data structure in this structure represents the results of another
// processing pass, and a further refinement of our run-time representation
// of the content.
//
// TODO:540 Make sure that comments are properly associated with markup tags
// (XML document data) and markup declarations (DTD stuff).
type GTree struct {
	// This data structure should know where its own root tag is, but
	// try not to use this field a lot because it might be redundant.
	RootTagIndex   int
	RootTagCount   int
	RootTagsDiffer bool
	// Scratch variables for matching start and end tags
	NrOpenTags int
	Tagstack

	// MMCtype is Markup & Mixed Content type, modeled after MimeType
	// in FU.InputFile, but focused on the "markup language" aspect.
	// Markdown is presumed to be MDITA, cos in any case, any flavor
	// of Markdown should be "mostly" compatible with MDITA. <br/>
	// Fields:
	// - [0] = doc family = `image/dita/lwdita/html/schema`
	// - [1] = doc format = format/`dtd`
	// - [2] = specifics
	// <br/> Common values:
	// - Textual  image files:  image /  text / (svg|eps)
	// - Binary   image files:  image /  bin  / (fmt)
	// - DITA13 content files:   dita / (tech|..) / (task|..)
	// - LwDITA content files: lwdita / (xdita|hdita|mdita[xp]) / (map|topic|..)
	// -   HTML content files:   html / (5|4) [/TBD]
	// - Parsed  schema files: schema / dtd / (root Tag)
	// MCMtype []string
}

func (T *GTree) EchoTo(w io.Writer) {
	// w.Write([]byte(T.Echo()))
}

func (et GTree) String() string {
	return "GTree!" // /* "dbg." + */ et.String()
}

/*
// Echo implements Markupper.
// TODO This func is probably actually a good def for String().
// TODO Extract from this a GTag.String(), to use in CLI cmd xmltokens .
func (pGF gfile.GFile) EchoGTree() string {
	sb := bytes.NewBufferString("")
	// println("<!-- GTree.ECHO.nTags:", pET.NrTags(), "-->")
	var peekNextToken *GTag
	var wroteNewlineSoMustIndentFollowingTag = false

	for i, pTag := range pGF.GTags {
		// pTag.GToken = nil
		// fmt.Printf("Tag[%02d]: %#v \n", i, *pTag)
		// println("Echo:", i, SU.GetIndent(pTag.Depth), pTag.String())
		TT := pTag.GTagTokType
		if wroteNewlineSoMustIndentFollowingTag {
			if TT == "SE" || TT == "EE" || TT == "CD" || TT == "Cmt" {
				sb.Write([]byte(SU.GetIndent(pTag.Depth)))
			}
		}
		// println("<<", TT, "||", pTag.Echo(), ">>")
		sb.Write([]byte(pTag.Echo()))
		//
		// NOW NEWLINE HANDLING
		//
		// If it's the end of the input, NEWLINE.
		isLastToken := (i == (len(pGF.GTags) - 1))
		if isLastToken {
			sb.Write([]byte("\n"))
			break
		}
		// Now we have to peek ahead.
		peekNextToken = pGF.GTags[i+1]
		// If this OR next token is block, NEWLINE.
		if pTag.IsBlock() || peekNextToken.IsBlock() {
			wroteNewlineSoMustIndentFollowingTag = true
			sb.Write([]byte("\n"))
			continue
		}
		// Reset the boolean.
		wroteNewlineSoMustIndentFollowingTag = false
		// If it's a start element and the element is empty,
		// NO NEWLINE, to keep the end element on the same line.
		// TODO:560 Merge the start and end tags into a self-closed tag.
		 * 		if pTag.GTagTokType == "SE" && peekNextToken.GTagTokType == "EE" {
		 * 			continue
		 * 		}
		 * If this token and the next are both inline, NO NEWLINE.
		 * 		nextIsInline := peekNextToken.TagSummary.IsInline()
		 * 		if pTag.TagSummary.IsInline() && nextIsInline {
		 * 			continue
		 * 		}
		//
		// TODO:700 Some other cases. Allow:
		 * <li> <p> Some text </p> </li>
		 * <title> Some text </title>, etc.
		//
		// NEWLINE!
		wroteNewlineSoMustIndentFollowingTag = true
		sb.Write([]byte("\n"))
	}
	return sb.String()
}
*/
