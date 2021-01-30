package gtree

import (
	"io"
	// FIXME (SetTagType) "github.com/fbaube/lwdx"
)

// NOTE: https://www.w3.org/TR/DOM-Level-2-Core/introduction.html
//
// "The DOM does not specify that documents must be implemented as a tree
// or a grove, nor does it specify how the relationships among objects be
// implemented. The DOM is a logical model that may be implemented in any
// convenient manner. In this spec, we use the term structure model to
// describe the tree-like representation of a document. We also use the
// term "tree" when referring to the arrangement of those information items
// that can be reached by using "tree-walking" methods (not incl. attributes)."

// MakeBlockTree builds a block/inline-aware mixed-content element tree from
// a RichTokenization, and saves it into the RT.
//
// By "block/inline-aware mixed-content" we mean that block elements are
// formed into the tree but inline elements (including interspersed text
// elements) are left as-is. The goal is that any leaf element may be
// passed to an editor widget. -n-
// TODO:640 Pass in a way to distinguish block and inline tags.
//
// For the case of an LwDITA XML/HTML file, this function may safely assume
// that the prolog ( "<?xml ...") and DOCTYPE have already been sniffed.
// This info might have to be added to the calling signature.
//
// Procedure: :ul:
// :: There is always the "current node"
// :: If we get a Block SE, it's a new kid
// ::
// TODO:810 This is actually forming subtrees out of the GNoes
func (pGT *GTree) MakeBlockTree(W io.Writer) error {
	println("MakeBlockTree: TODO")
	return nil
}

/*
	if W == nil {
		W = io.Discard
	}
	if pGT.RootTagIndex == -1 {
		return fmt.Errorf("MakeBlockTree: no root element")
	}
	var rootTag = pGT.GTags[pGT.RootTagIndex]
	var openTag = rootTag
	var nextTag *GTag
	for i, pTag := range pGT.GTags {
		if i <= pGT.RootTagIndex {
			continue
		}
		// Can be nil
		nextTag = pGT.GTags[i]
		switch pTag.GTagTokType {

		case "SE": // StartElement  // could be: <t
			openTag.AddKid(nextTag)
			openTag = nextTag
		case "EE": // EndElement     // could be: t>
			openTag = openTag.GetParent()
		case "CD": // CDATA          // could be: CD
			// openTag.AddKid()
		case "PI": // Proc. Instr.   // could be: <?
		case "Cmt": // XML comment   // could be: --
		case "Dir": // XML directive // could be: <!
			// The following are actually DIRECTIVE SUBTYPES, but they
			// are put in this list so that they can be assigned freely.
		case "DOCTYPE":
		case "ELEMENT":
		case "ATTLIST":
		case "ENTITY":
		case "NOTATION":
		}
	}
	return nil
}
*/
