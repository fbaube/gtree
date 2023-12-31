package gtree

import (
	"fmt"
	"os"

	// "bytes"
	FU "github.com/fbaube/fileutils"
	MU "github.com/fbaube/miscutils"

	// SU "github.com/fbaube/stringutils"

	MD "github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	// "github.com/dimchansky/utfbom"
)

// NewGTreeFromMarkdownFile is a convenience function that reads in the
// file, which is presumed to be Markdown (MDITA flavor), then tokenizes it,
// and then passes the buffered file contents to the next function, below.
//
// TODO:670 Provide a slice of dirpaths, for resolving external IDs.
func NewGTreeFromMarkdownFile(path FU.AbsFilePath) (pET *GTree, err error) {

	var e error
	var bb []byte

	var xhtml = false
	var page = true
	var toc = true
	var css, title string
	// rendererOutput string

	xhtml = false

	bb, e = os.ReadFile(string(path))
	if e != nil {
		return nil, fmt.Errorf("gxml.GTree.NewGTreeFromMarkdownFile.ReadFile<%w>", e)
	}
	// Keep an extra copy around so that we can read and re-read the entire file.
	var theContent = string(MU.DupeByteSlice(bb))

	// set up options
	var extensions = parser.NoIntraEmphasis |
		parser.Tables |
		parser.FencedCode |
		parser.Autolink |
		parser.Strikethrough |
		parser.SpaceHeadings
	var renderer MD.Renderer
	// render the data into HTML
	var htmlFlags html.Flags
	if xhtml {
		htmlFlags |= html.UseXHTML
	}
	/* old code
	if smartypants {
		htmlFlags |= html.Smartypants
	}
	if fractions {
		htmlFlags |= html.SmartypantsFractions
	}
	if latexdashes {
		htmlFlags |= html.SmartypantsLatexDashes
	}
	*/
	title = ""
	if page {
		htmlFlags |= html.CompletePage
		// title = getTitle(input)
	}
	if toc {
		htmlFlags |= html.TOC
	}
	params := html.RendererOptions{
		Flags: htmlFlags,
		Title: title,
		CSS:   css,
	}
	renderer = html.NewRenderer(params)

	// parse and render
	var output []byte
	parser := parser.NewWithExtensions(extensions)
	output = MD.ToHTML([]byte(theContent), parser, renderer)
	// output the result
	var out *os.File
	/* old code
	if len(args) == 2 {
		if out, err = os.Create(args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating %s: %v", args[1], err)
			os.Exit(-1)
		}
		defer out.Close()
	} else { */
	out = os.Stdout
	// }
	if _, err = out.Write(output); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing output:", err)
		os.Exit(-1)
	}

	/* old code
	md := MD.New(
		MD.HTML(true),
		MD.Tables(true),
		MD.Linkify(true),
		MD.Typographer(false),
		MD.XHTMLOutput(xhtml),
	)
	tokens := md.Parse([]byte(theContent))
	title = GetTitleFromMarkdownTokens(tokens)
	println("==> Got MD title:", title)

	pET, e = NewGTreeFromMarkdownTokens(tokens, theContent)
	if e != nil {
		return nil, errors.Wrap(e, "gxml.GTree.NewFromXmlFile.newFromBuffer")
	}

	*/

	return pET, nil
}

// NewGTreeFromMarkdownTokens takes a string, not an io.Reader, so that
// we know that the caller too can keep his own copy of the input.
// Along with creating the GTree, it also sets XmlInfo and DitaInfo.
// TODO:20 A future version could have a flag to only set those two fields
// and not create the GTree.
//
// TODO:310 FIXME Check that root Tag matches DOCTYPE.
// TODO:420 FIXME Provide a slice of dirpaths, for resolving external IDs.
// TODO:390 FIXME If multiple root Tags, set Xml contype to Fragment
// TODO:340 FIXME If has DOCTYPE, set XML contype to document (unless is Fragment)
// TODO:370 FIXME If has LwDITA DOCTYPE, set DITA contype.

// =================================================================
// =================================================================
// =================================================================
// =================================================================

// NewGTagFromMDtoken is TODO.
// TODO:630 Pass a writer for Echo.
// NOTE:380 Returns "nil" if the token is valid but useless, and can
// be skipped, such as an xml.CharData that is all whitespace;
// NOTE:1010 that it might cause problems.
