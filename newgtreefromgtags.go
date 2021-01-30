package gtree

// "github.com/dimchansky/utfbom"

// NewGTreeFromGTags is TBS.
//
// TODO:320 FIXME Check that root Tag matches DOCTYPE.
// TODO:430 FIXME Provide a slice of dirpaths, for resolving external IDs.
// TODO:400 FIXME Multiple root Tags, set Xml contype to Fragment
// TODO:350 FIXME If has DOCTYPE, set XML contype to document (unless is Fragment)
// TODO:380 FIXME If has LwDITA DOCTYPE, set DITA contype.
//
func NewGTreeFromGTags(GEs []*GTag) (pGT *GTree, err error) {

	// var e error
	// SETUP #1: Allocate memory.
	pGT = new(GTree)
	pGT.Tagstack = *new(Tagstack)
	var pTag *GTag
	var i int

	for i, pTag = range GEs {
		atRootLevel := (pGT.NrOpenTags == 0)

		if pTag.TTType == "SE" {
			// println("SE.kwd:", pTag.GName.String())
			pGT.Tagstack.Push(NewTagentry(pTag.GName.String(), i))
			// println("Pušt", i, pTag.GName.String())
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
					elog.Printf("Got another root element <%s>", pTag.GToken.GName)
					println("==> Got second root element <", pTag.GToken.GName.String(),
						">: XML data file is a fragment")
					if pTag.GName.String() != GEs[pGT.RootTagIndex].GName.String() {
						pGT.RootTagsDiffer = true
					}
					// pXI.xmlContype = "Fragments"
				}
			}
		}
		if pTag.TTType == "EE" {
			if atRootLevel {
				elog.Printf("Unmatched top-level end tag <%s>", pTag.Keyword)
				panic("Unmatched top-level end tag: " + pTag.Keyword)
			}
			TE := pGT.Tagstack.Pop()
			// println("Popt", i, "::", TE.Index(), TE.Tag())
			if TE.Tag() != pTag.GName.String() {
				elog.Printf("Tag mismatch: |(start-tag)|%s|v|(end-tag)|%s|>",
					TE.Tag(), pTag.GName.String())
				panic("Bad tag stack: " + TE.Tag() + " v " + pTag.GName.String())
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
