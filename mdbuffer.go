package gtree

import (
	"fmt"

	CT "github.com/fbaube/ctoken"
	"github.com/fbaube/gtoken"
	MU "github.com/fbaube/miscutils"
	SU "github.com/fbaube/stringutils"
	AST "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Markdown Node Type strings
var mdnt map[AST.NodeKind]string

func init() {
	mdnt = make(map[AST.NodeKind]string)
	mdnt[AST.KindDocument] = "Document"
	mdnt[AST.KindBlockquote] = "BlkQuote"
	mdnt[AST.KindList] = "List"
	// mdnt[AST.KindItem] = "Item"
	mdnt[AST.KindParagraph] = "Para"
	mdnt[AST.KindHeading] = "Heading"
	// mdnt[AST.KindHorizontalRule] = "HorzRule"
	mdnt[AST.KindEmphasis] = "Emph"
	// mdnt[AST.KindStrong] = "Strong"
	// mdnt[AST.KindDel] = "Del"
	mdnt[AST.KindLink] = "Link"
	mdnt[AST.KindImage] = "Image"
	mdnt[AST.KindText] = "Text"
	mdnt[AST.KindHTMLBlock] = "HtmlBlk"
	mdnt[AST.KindCodeBlock] = "CodeBlk"
	// mdnt[AST.KindSoftbreak] = "SoftBrk"
	// mdnt[AST.KindHardbreak] = "HardBrk"
	mdnt[AST.KindCodeBlock] = "CodeBlock"
	mdnt[AST.KindCodeSpan] = "CodeSpan"
	// mdnt[AST.KindHTMLSpan] = "HtmlSpan"
	/* old code
	mdnt[AST.KindTable] = "Table"
	mdnt[AST.KindTableCell] = "TblCell"
	mdnt[AST.KindTableHead] = "TblHead"
	mdnt[AST.KindTableBody] = "TblBody"
	mdnt[AST.KindTableRow] = "TblRow"
	*/
}

// GTokenizeMDbuffer takes the raw XML and parses it into a slice of `GToken`s.
// It takes a string, not an `io.Reader`, so we know that the caller already had
// access to the full file contents and verified that it is in fact an XML file.
func GTokenizeMDbuffer(inString string) (GTokzn []*gtoken.GToken, err error) {

	// var e error
	if true {
		tt := MU.Into("GTokenizeMDbuffer")
		defer MU.Outa("GTokenizeMDbuffer", tt)
	}

	// NOTE BF v2 MD stuff
	// Get the AST from BlackFriday
	var theMD parser.Parser // BF2.Markdown
	var RootBFnode AST.Node
	// theMD = BF2.New()
	theMD = parser.NewParser()
	// RootBFnode = theMD.Parse([]byte(inString))
	var theRdr text.Reader
	theRdr = text.NewReader([]byte(inString))
	// Parse(reader text.Reader, opts ...ParseOption) ast.Node
	RootBFnode = theMD.Parse(theRdr)
	println("==BEG== DumpNode:BF:Root")
	DumpBFnode(RootBFnode, 0)
	println("==END== DumpNode:BF:Root")

	// NOTE:810 own.GTag's
	// Convert the BF AST to our own GTag's
	var RootGTag *GRootTag
	RootGTag = NewGTagTreeFromBFtree(RootBFnode)
	println("==BEG== DumpNode:GTag:Root")
	println(GTag(*RootGTag).String())
	println("==END== DumpNode:GTag:Root")
	return GTokzn, nil
}

func NewGTagTreeFromBFtree(p AST.Node) *GRootTag {
	println("NewGTagTreeFromBFtree ENTRY")
	var pp *GTag
	pp = NewGTagFromBFnode(p)
	var Kids []AST.Node
	Kids = KidsAsSlice(p)
	for i, c := range Kids {
		fmt.Printf("[%d]INN ", i)
		DumpBFnode(c, 0)
		fmt.Printf("[%d]OUT ", i)
		cc := NewGTagFromBFnode(c)
		if pp == nil {
			panic("nil Parent")
		}
		if cc == nil {
			panic("nil Kid")
		}
		pp.AddKid(cc)
	}
	var pRoot *GRootTag
	pRoot = (*GRootTag)(pp)
	return pRoot
}

// NewGTagFromBFnode basically just assigns to this field:
// - gparse.GToken
// which comprises:
// - GTagTokType
// - XName
// - GAttList
func NewGTagFromBFnode(p AST.Node) *GTag {
	var NT AST.NodeType
	var NK AST.NodeKind
	NT = p.Type()
	NK = p.Kind()
	var pp *GTag
	pp = new(GTag)
	pp.TDType = CT.TD_type_ELMNT
	sKids := ListKids(p)
	fmt.Printf("New GTag :: type %d :: kind %d :: %s \n", NT, NK, sKids)

	switch NK.String() { // NT.String() {
	case "Document":
		println("START OF DOCUMENT")
		pp.CName = *CT.NewCName("", "markdown")
		return pp
	case "List", "Item":
		var lst = p.(*AST.List) // p.List
		println(DumpList(*lst))
		return pp
	case "Heading":
		var hdg = p.(*AST.Heading)
		println(DumpHdg(*hdg))
		pp.CName = *CT.NewCName("", fmt.Sprintf("H%d", hdg.Level))
		return pp
	case "Link", "Image":
		var lnk = p.(*AST.Link)
		println(DumpLink(*lnk))
		return pp
	case "CodeBlock":
		var cbd = p.(*AST.CodeBlock)
		println(DumpCdBlk(*cbd))
		return pp
		/* old code
		case "Table", "TableCell", "TableHead", "TableBody", "TableRow":
			var tcd = p.(*AST.TableCell)
			println(DumpTableCell(*tcd))
			return pp
		*/
	case "BlockQuote":
		return pp
	case "Paragraph":
		return pp
	case "HorizontalRule":
		return pp
	case "Emph":
		return pp
	case "Strong":
		return pp
	case "Del":
		return pp
	case "Text":
		return pp
	case "HTMLBlock":
		return pp
	case "Softbreak":
		return pp
	case "Hardbreak":
		return pp
	case "Code":
		return pp
	case "HTMLSpan":
		return pp
	}
	return pp
}

// =================
// ==== BF node ====
// == (BF library) =
// =================

func KidsAsSlice(p AST.Node) []AST.Node {
	var pp []AST.Node
	var c AST.Node
	c = p.FirstChild()
	for c != nil {
		pp = append(pp, c)
		c = c.NextSibling()
	}
	return pp
}

func ListKids(p AST.Node) string {
	var pp []AST.Node
	pp = KidsAsSlice(p)
	if len(pp) == 0 {
		return "<0-kids>"
	}
	var s string
	for i, c := range pp {
		s += fmt.Sprintf("[%d:%s]", i, c.Type)
	}
	return s
}

func DumpBFnode(p AST.Node, iLvl int) {
	var s string
	s = fmt.Sprintf("%s<%s>: ", SU.GetIndent(iLvl), mdnt[p.Kind()])
	// s += "<|" + string(p.Literal) + "|> "
	s += DumpHdg(*p.(*AST.Heading)) + " "
	s += DumpList(*p.(*AST.List)) + " "
	s += DumpCdBlk(*p.(*AST.CodeBlock)) + " "
	s += DumpLink(*p.(*AST.Link)) + " "
	// s += DumpTableCell(p.TableCellData) + " "
	println(s)
	var c AST.Node
	c = p.FirstChild()
	for c != nil {
		DumpBFnode(c, iLvl+1)
		c = c.NextSibling()
	}
}

// type myBFnodeVisitor func(node *Node, entering bool) WalkStatus
func myBFnodeVisitor(N AST.Node, entering bool) AST.WalkStatus {
	if !entering {
		return AST.WalkContinue // GoToNext
	}
	// fmt.Printf("%s \n",DumpBFnode(N, 0)) // SU.GetIndent(lvl)
	return AST.WalkContinue // GoToNext
}

// ==================
// ==== AST node ====
// = (found object) =
// ==================

/* old code
Type     string
Literal  string      `json:",omitempty"`
Attr     interface{} `json:"-"`
Children []*ASTNode  `json:",omitempty"`

func (p *ASTNode) DumpASTnode() string {

for child := range p.Children { // p.FirstChild; child != nil; child = child.Next {
	a.Children = append(a.Children, NewASTNode(child))
}
}

// type myASTnodeVisitor func(node *Node, entering bool) WalkStatus
func myASTnodeVisitor(N *ASTnode, entering bool) BF.WalkStatus {
	if !entering {
		return BF.GoToNext
	}
	fmt.Printf("%s \n", DumpASTnode(N)) // SU.GetIndent(lvl)
	return BF.GoToNext
}
*/

// ================
// ===== GTag =====
// ================

func DumpGTag(p AST.Node) string {
	var s string
	s = fmt.Sprintf("BFnode<%s>: ", mdnt[p.Kind()])
	// s += "<|" + SU.NormalizeWhitespace(string(p.Literal)) + "|> "
	s += DumpHdg(*(p.(*AST.Heading)))
	s += DumpList(*(p.(*AST.List)))
	s += DumpCdBlk(*(p.(*AST.CodeBlock)))
	s += DumpLink(*(p.(*AST.Link)))
	// s += DumpTableCell(p.TableCellData)
	return s
}

// type myGTagVisitor func(node *Node, entering bool) WalkStatus
func myGTagVisitor(N AST.Node, entering bool) AST.WalkStatus {
	if !entering {
		return AST.WalkContinue // GoToNext
	}
	fmt.Printf("%s \n", DumpGTag(N)) // SU.GetIndent(lvl)
	return AST.WalkContinue          // GoToNext
}

// =======================
// ==== BF node stuff ====
// =======================

// Level        int    // This holds the heading level number
// HeadingID    string // This might hold heading ID, if present
// IsTitleblock bool   // Specifies whether it's a title block
func DumpHdg(h AST.Heading) string {
	// var ttl string
	// if h.IsTitleblock {
	//    ttl = "IsTtl:"
	// }
	if h.Level == 0 { // && h.HeadingID == "" {
		return ""
	}
	return fmt.Sprintf("<H%d::%s%s> ", h.Level) // , ttl, h.HeadingID)
}

// ListFlags   ListType
// Tight       bool   // Skip <p>s around list item data if true
// BulletChar  byte   // '*', '+' or '-' in bullet lists
// Delimiter   byte   // '.' or ')' after the number in ordered lists
// RefLink   []byte   // If not nil, turns this list item into a footnote item and triggers different rendering
// IsFootnotesList bool   // This is a list of footnotes
func DumpList(L AST.List) string {
	var tight string // ftnts
	if L.IsTight {
		tight = "IsTight:"
	}
	/* old code
	if L.IsFootnotesList {
		ftnts = "IsFtnts:"
	}
	if L.ListFlags == 0 && L.BulletChar == 0 &&
		L.Delimiter == 0 && len(L.RefLink) == 0 {
		return ""
	}
	*/
	return fmt.Sprintf("<List:%s%sBult:%c:Delim:%c:RefLink:%s> ",
		// tight, ftnts, L.BulletChar, L.Delimiter, L.RefLink)
		tight, L.Marker)
}

// IsFenced   bool  // Fenced code block, or else an indented one
// Info     []byte  // This holds the info string
// FenceChar  byte
// FenceLength int
// FenceOffset int
func DumpCdBlk(cb AST.CodeBlock) string {
	return "CodeBlock?!?!"
	/* old code
	var fenced string
	if cb.IsFenced {
		fenced = "IsFenced"
	}
	if len(cb.Info) == 0 && cb.FenceChar == 0 &&
		cb.FenceLength == 0 && cb.FenceOffset == 0 {
		return ""
	}
	return fmt.Sprintf("<CdBlk:%s:ch:%c:len:%d:ofs:%d:Info:%s> ",
		fenced, cb.FenceChar, cb.FenceLength, cb.FenceOffset, string(cb.Info))
	*/
}

// Destination []byte // Destination is what goes into a href
// Title       []byte // The tooltip thing that goes in a title attribute
// NoteID      int    // The S/N of a footnote, or 0 if not a footnote
// Footnote    *Node  // If footnote, a direct link to the FN Node, else nil.
func DumpLink(L AST.Link) string {
	return "Link?!?!"
	/* old code
	var isFN bool
	isFN = (L.NoteID != 0) && (L.Footnote == nil)
	if len(L.Destination) == 0 && len(L.Title) == 0 &&
		L.NoteID == 0 && L.Footnote == nil {
		return ""
	}
	if !isFN {
		return fmt.Sprintf("<Link:Ttl:%s:Dest:%s> ",
			string(L.Title), string(L.Destination))
	}
	return fmt.Sprintf("<FN-link:#%d:Ttl:%s:Dest:%s> ",
		L.NoteID, string(L.Title), string(L.Destination))
	*/
}

/* old code
// IsHeader  bool       // This tells if it's under the header row
// Align CellAlignFlags // This holds the value for align attribute
func DumpTableCell(tc AST.TableCellData) string {
	if tc.Align == 0 && !tc.IsHeader {
		return ""
	}
	if tc.IsHeader {
		return "<TblCell:IsHdr> "
	}
	return "<TblCell:notHdr> "
}
*/
