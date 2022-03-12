package gtree

import (
	"github.com/fbaube/gtoken"
	// "github.com/dimchansky/utfbom"
	L "github.com/fbaube/mlog"
)

func MakeGTagsFromGTokens(GTs []*gtoken.GToken) (GEs []*GTag, err error) {
	var pGE *GTag
	var GT gtoken.GToken
	var e error

	// fmt.Printf("gtree.MakeGTagsFromGTokens: nGTokens: %d \n", len(GTs))
	for _, pGT := range GTs {
		if pGT == nil {
			// println("MakeGTagsFromGTokens: nil")
			continue
		}
		GT = *pGT
		pGE, e = NewGTagFromGToken(GT)
		// println("new.GTag:", pGE.String())
		isEmptyCDATA := ((pGE == nil) && (e == nil))
		// If pTag is nil, discard the GToken
		if pGE == nil {
			GTs = nil
		}
		if isEmptyCDATA {
			continue
		}
		if e != nil { // pTag is nil
			// fmt.Sprintf()
			L.L.Error("NewGTreeFromXMLtoken: failed: %s", e.Error())
			continue
		}
		// pTag is not nil
		// pTag.Depth = gxml.NrOpenTags
		GEs = append(GEs, pGE)
	}
	return GEs, nil
}
