package gtree

// This file: Generic Golang XML Tags

import (
	"fmt"
	// "github.com/fbaube/gparse"
	"github.com/fbaube/gtoken"
	"github.com/fbaube/lwdx"
	ON "github.com/fbaube/orderednodes"
)

// GTag is a generic golang XML tag, used mainly for representing XML
// tags (or their Markdown equivalents) in a mixed content document.
// Child elements (called "Kids") are referenced by the embedded [Nord].
//
// Note that this is the appropriate struct for indicating block/inline,
// via func [IsBlock].
//
// (GTag might also be useful tho for holding multi-level attribute
// info in DTDs, but then again we also define a very different [DTag].)
//
// GTag is also used to represent non-tag XML items, including PIs,
// Comments, Directives, and CDATA character data items. Therefore
// a GTag is created for every XML token (even [EndElement]s), and
// they are linked into a tree structure (a GTree).
//
// GTag uses pointer receivers, not method receivers. <br/>
// For its kids it uses a linked list, not a slice.
// .
type GTag struct {
	// Nord provides tree structure
	ON.Nord
	// [GToken] includes Name and Attribute info for XML
	// tags. For a simple tag that cannot be namespaced,
	// such as a "tag" in Markdown, the tag name is in
	// [GToken.GName.Local].
	//
	// NOTE: For LwDITA's Markdown-XP, we could use
	// the Attributes to store Pandoc-style attributes.
	// TODO: Every node needs both NAMESPACE and LANGUAGE,
	// because they are inherited.
	gtoken.GToken
	// MatchingTagsIndex is used for tags only,
	// i.e. "SE", "EE" in LwDITA, HTML5, XML.
	// MatchingTagsIndex int

	// TODO: Should TagalogEntry be moved elsewhere ?
	*lwdx.TagalogEntry
	// EntityIsParameter is a bool field used for XML ENTITYs only.
	// It indicates whether the entity defined using a "%" or not.
	// This distinguishes a parameter/DTD entity from a general/data
	// entity. This is recorded during parsing, for later use when
	// we fully process the entity declaration.
	EntityIsParameter bool
	// These two bool fields used only where entity references
	// ( &name; %name; ) are legal.
	hadEntities      bool
	stillHasEntities bool

	// NOTE Maybe add these in the future.
	// D+T of last mod
	// D+T of last mod of subtrees (propagates upward)
	// IsExpanded bool (GUI helper)
	// NrOfKids    int
	// NrOfKidsAll int (incl. subtrees)
	// Size        int (bytes, or any other resource of interest)
	// SizeAll     int (incl. subtrees)
}

// IsBlock needs to check for which schema, because some tags occur
// in multiple schemata but with differing values for block/inline.
// .
func (p *GTag) IsBlock() bool {
	fmt.Printf("gtree/gtag:L76")
	return true
	//return p.TagSummary.LwditaMode == "BLOCK" ||
	//	p.TagSummary.Html5Mode == "BLOCK"
}

// GRootTag makes sure that assignments to/from a root node are explicit.
type GRootTag GTag

// NewGTag initializes the node with parser results.
func NewGTag(aNS, aName string) *GTag {
	newGTag := new(GTag)
	newGTag.GToken.GName = gtoken.GName{Space: aNS, Local: aName}
	return newGTag
}

// NewKid initializes the node with parser results
// and adds it to N as the last kid.
func (anE *GTag) NewKid(aNS, aName string) *GTag {
	if anE == nil {
		panic("NewKid got nil parent")
	}
	return anE.AddKid(NewGTag(aNS, aName)).(*GTag)
}

// Echo implements Markupper.
func (p *GTag) Echo() string {
	return p.GToken.Echo()
}

// String implements Markupper.
func (p GTag) String() string {
	var sBlk string
	if p.IsBlock() {
		sBlk = "(BLK) "
	}
	var s = p.GToken.Echo() +
		// fmt.Sprintf(" [d%d:%d] ", p.Depth, p.MatchingTagsIndex) +
		fmt.Sprintf(" [d%d] %s", p.Depth, sBlk) + p.TagalogEntry.String()
	// p.Nord.String()
	return s
}
