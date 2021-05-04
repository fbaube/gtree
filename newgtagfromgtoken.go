package gtree

import (
	"fmt"

	"github.com/fbaube/gtoken"
	"github.com/fbaube/lwdx"
	// "github.com/dimchansky/utfbom"
	// "github.com/pkg/errors"
)

// NewGTagFromGToken embeds the GToken and processes it.
// NOTE:390 Returns (`nil,nil`) if the token is valid but useless, and
// should be skipped, i.e. an `xml.CharData` that is all whitespace.
func NewGTagFromGToken(GT gtoken.GToken) (pTag *GTag, e error) {
	pTag = new(GTag)
	pTag.GToken = GT

	if "" == pTag.TTType {
		// println("NewGTagFromGToken: NIL")
		return nil, nil
	}

	switch pTag.TTType {

	case "Elm":
		// pTag.Depth = NrOpenTags
		// NrOpenTags++
		var TT lwdx.TagSummary
		var ok bool
		if TT, ok = lwdx.TagSummaries[pTag.GToken.GName.Local]; !ok {
			ilog.Printf("Unrecognized tag: <" + pTag.Keyword + ">")
			// TODO:1210 reinstate this next error, and change the above "ilog"
			//      to "elong", when we can simply warn for DITA 1.3 tags
			// return pTag, errors.New("Unrecognized tag: <" + pTag.Keytext + ">")
		}
		pTag.TagSummary = TT
		return pTag, nil

	case "end":
		// NrOpenTags--
		// pTag.Depth = NrOpenTags
		var TT lwdx.TagSummary
		var ok bool
		if TT, ok = lwdx.TagSummaries[pTag.Keyword]; !ok {
			// TODO
		}
		pTag.TagSummary = TT
		return pTag, nil

	case "PrI":
		pTag.TagSummary = lwdx.TTblock
		// TODO:140 Attach this PI to its parent element in the GTree
		// newNode = parentNode.NewKid("<?", myTarget)
		// newNode.StringValue = myInst
		return pTag, nil

	case "ChD":
		if pTag.Echo() == "" {
			// ilog.Printf("PCDATA is all whitespace: \n")
			// DO NOTHING
			// NIL IT OUT
			// NOTE:550 This may do weird things to elements
			// that have text content models.
			// println("WARNING: Got an all-whitespace xml.CharData")
			return nil, nil
		}
		pTag.TagSummary = lwdx.TTinline
		return pTag, nil

	case "Cmt":
		// TODO:130 Attach this Comment to its parent element in the GTree,
		// which is in fact the first XML element that follows this XML
		// comment. So, it should be done in a later pass.
		// newNode = parentNode.NewKid("<!", "--")
		// newNode.StringValue = tokenString
		pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	case "Dir":
		pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	case "":
		print(GT.String())
		pTag.TTType = "ERR"
		pTag.TagSummary = lwdx.TTblock
		// println(fmt.Sprintf("NIL GToken.type<%s> for: %+v", GT.TTType, GT))
		return nil, fmt.Errorf("NIL GToken.type<%s> for: %+v", GT.TTType, GT)

	case "Doc":
		println("--> Made GTag for Doc GToken") // , GT.String())
		pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	default:
		print(GT.String())
		pTag.TTType = "ERR"
		pTag.TagSummary = lwdx.TTblock
		println(fmt.Sprintf("Unrecognized GToken.type<%s> for: %+v",
			GT.TTType, GT))
		return nil, fmt.Errorf("Unrecognized GToken.type<%s> for: %+v",
			GT.TTType, GT)
	}
}
