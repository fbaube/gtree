package gtree

// "github.com/dimchansky/utfbom"

import (
	L "github.com/fbaube/mlog"
	XU "github.com/fbaube/xmlutils"
)

// NewGTreeFromGTags is TBS.
//
// TODO: FIXME Check that root Tag matches DOCTYPE.
// TODO: FIXME Provide a slice of dirpaths, for resolving external IDs.
// TODO: FIXME Multiple root Tags, set Xml contype to Fragment
// TODO: FIXME If has DOCTYPE, set XML contype to document (unless is Fragment)
// TODO: FIXME If has LwDITA DOCTYPE, set DITA contype.
func NewGTreeFromGTags(GEs []*GTag) (pGT *GTree, err error) {

	// var e error
	// SETUP #1: Allocate memory.
	pGT = new(GTree)
	pGT.Tagstack = *new(Tagstack)
	var pTag *GTag
	var i int

	for i, pTag = range GEs {
		atRootLevel := (pGT.NrOpenTags == 0)

		if pTag.TDType == XU.TD_type_ELMNT {
			// println("SE.kwd:", pTag.XName.String())
			pGT.Tagstack.Push(NewTagentry(pTag.XName.Echo(), i))
			// println("Pušt", i, pTag.XName.String())
			// pTag.Depth = pET.NrOpenTags
			pGT.NrOpenTags++

			if atRootLevel {
				// Is it the first root element we've detected ?
				if pGT.RootTagCount == 0 {
					// No problem
					pGT.RootTagCount = 1
					pGT.RootTagIndex = i
				} else {
					// Problem
					pGT.RootTagCount++
					L.L.Error("Got another root element <%s>", pTag.GToken.XName)
					println("==> Got second root element <", pTag.GToken.XName.Echo(),
						">: XML data file is a fragment")
					if pTag.XName.Echo() != GEs[pGT.RootTagIndex].XName.Echo() {
						pGT.RootTagsDiffer = true
					}
					// pXI.xmlContype = "Fragments"
				}
			}
		}
		if pTag.TDType == XU.TD_type_ENDLM {
			if atRootLevel {
				L.L.Error("Unmatched top-level end tag <%s>", pTag.TagOrPrcsrDrctv)
				panic("Unmatched top-level end tag: " + pTag.TagOrPrcsrDrctv)
			}
			TE := pGT.Tagstack.Pop()
			// println("Popt", i, "::", TE.Index(), TE.Tag())
			if TE.Tag() != pTag.XName.Echo() {
				L.L.Error("Tag mismatch: |(start-tag)|%s|v|(end-tag)|%s|>",
					TE.Tag(), pTag.XName.Echo())
				panic("Bad tag stack: " + TE.Tag() + " v " + pTag.XName.Echo())
			}
			pGT.NrOpenTags--
			// Point Start and End at each other
			// println("\t(DD) OK tag <", TE.tag, "> starts[", TE.index, "] ends[", i, "]")
			GEs[i] = nil
			// pTag.MatchingTagsIndex = TE.Index()
			// GEs[TE.Index()].MatchingTagsIndex = i
		}
	}
	// if !quietConsole { println("==> Made Micodo tokens, matched start/end tags") }
	return pGT, nil
}
