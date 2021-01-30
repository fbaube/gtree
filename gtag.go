package gtree

// This file: Generic Golang XML Tags

import (
	"fmt"
	// "github.com/fbaube/gparse"
	"github.com/fbaube/gtoken"
	SU "github.com/fbaube/stringutils"
	"github.com/fbaube/lwdx"
)

// GTag is a generic golang XML tag, used mainly for representing XML
// tags (or their Markdown equivalents) in a mixed content document.
// Child elements (called "Kids") are referenced by the embedded `GNode`.
//
// (`GTag` might also be useful tho for holding multi-level attribute
// info in DTDs, but then again we also define a very different `DTag`.)
//
// `GTag` is also used to represent non-tag XML items, including PIs,
// Comments, Directives, and CDATA character data items. Therefore a
// GTag is created for every XML token (even EndElement's), and they
// are linked into a tree structure (a `GTree`).
//
// `GTag` uses pointer receivers, not method receivers. <br/>
// For its kids it uses a linked list, not a slice.
//
type GTag struct {
	// GToken includes Name and Attribute info for XML tags.
	// For a simple tag that cannot be namespaced, like any
	// "tag" in Markdown, the tag name is in GToken.GName.Local .
	// NOTE:960 that for Markdown, we could the Attributes to store
	// the Pandoc-style attributes used in LwDITA's MD-XP.
	gtoken.GToken

	// TODO:280 Every node needs both NAMESPACE and LANGUAGE,
	// because they are inherited.

	// Fields used for tags only (LwDITA, HTML5, XML; i.e. "SE", "EE")
	// Depth, MatchingTagsIndex int

	// TODO:790 This data can and should be moved elsewhere
	lwdx.TagSummary

	// Provide the tree structure
	GNode

	// This bool field is used for XML ENTITYs only. It indicates whether the
	// entity defined using a "%" or not. This distinguishes a parameter/DTD
	// entity from a general/data entity. This is recorded during parsing,
	// for later use when we fully process the entity declaration.
	EntityIsParameter bool

	// These two bool fields used only where entity references ( &name; %name; )
	// are legal.
	hadEntities      bool
	stillHasEntities bool

	// NOTE:310 Maybe add these in the future.
	// D+T of last mod
	// D+T of last mod of subtrees (propagates upward)
	// IsExpanded bool (GUI helper)
	// NrOfKids    int
	// NrOfKidsAll int (incl. subtrees)
	// Size        int (bytes, or any other resource of interest)
	// SizeAll     int (incl. subtrees)
}

// Make sure that assignments to/from root node are explicit.
type GRootTag GTag

// FirstKid provides read-only access for other packages.
func (p *GTag) FirstKid() *GTag {
	return p.firstKid
}

// NextKid provides read-only access for other packages.
func (p *GTag) NextKid() *GTag {
	return p.nextKid
}

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
	return anE.AddKid(NewGTag(aNS, aName))
}

// Echo implements Markupper.
func (p *GTag) Echo() string {
	return p.GToken.Echo()
}

// String implements Markupper.
func (p GTag) String() string {
	var s = p.GToken.Echo() +
		// fmt.Sprintf(" [d%d:%d] ", p.Depth, p.MatchingTagsIndex) +
		fmt.Sprintf(" [d%d] ", p.Depth) + p.TagSummary.String()
		// p.GNode.String()
	/*
		EntityIsParameter bool
		hadEntities      bool
		stillHasEntities bool
	*/
	return s
}

// StringRecursively is fab.
func (p GTag) StringRecursively(s string, iLvl int) string {

	s += SU.GetIndent(iLvl) + p.String() + "\n" // p.GToken.Echo() +
	// fmt.Sprintf(" d%d::[%d] ", p.Depth, p.MatchingTagsIndex) +
	// p.TagSummary.String() + "\n"
	var kids []*GTag
	kids = p.KidsAsSlice()
	for _, k := range kids {
		s += k.StringRecursively(s, iLvl+1)
	}

	/*
		EntityIsParameter bool
		hadEntities      bool
		stillHasEntities bool
	*/
	return s
}

func (p *GTag) KidsAsSlice() []*GTag {
	var pp []*GTag
	c := p.firstKid
	for c != nil {
		pp = append(pp, c)
		c = c.nextKid
	}
	return pp
}
