package gtree

import (
	"fmt"

	"github.com/fbaube/gtoken"
	"github.com/fbaube/lwdx"
	// "github.com/dimchansky/utfbom"
	L "github.com/fbaube/mlog"
	XU "github.com/fbaube/xmlutils"
)

// NewGTagFromGToken embeds the GToken and processes it.
// NOTE: Returns (`nil,nil`) if the token is valid but useless, and
// should be skipped, i.e. an `xml.CharData` that is all whitespace.
func NewGTagFromGToken(inGTkn gtoken.GToken) (pTag *GTag, e error) {
	pTag = new(GTag)
	pTag.GToken = inGTkn

	if "" == inGTkn.TDType {
		L.L.Info("NewGTagFromGToken: EMPTY TDType")
		return nil, nil
	}

	switch inGTkn.TDType {

	case XU.TD_type_ELMNT:
		// pTag.Depth = NrOpenTags
		// NrOpenTags++
		// var TT lwdx.TagSummary
		var pTE *lwdx.TagalogEntry

		var theTag string
		theTag = pTag.GToken.XName.Local

		// if p,ok = lwdx.TagInfo[theTag]; !ok {
		pTE = lwdx.GetTEbyTagAndMarkupType(theTag, pTag.GToken.MarkupType)
		// if found OK
		if pTE != nil {
			pTag.TagalogEntry = pTE
			return pTag, nil
		}
		L.L.Dbg("GToken: %+v", inGTkn)
		// L.L.Dbg("GTag: %+v", *pTag)
		if pTag.TagOrPrcsrDrctv == "" {
			L.L.Warning("Missing tag")
		} else {
			L.L.Warning("Unrecognized tag: <" + pTag.TagOrPrcsrDrctv + ">")
		}
		// TODO: We will need an error here: reinstate this next error.
		// return pTag, errors.New("Unrecognized tag: <" + pTag.Keytext + ">")
		return pTag, nil

	case XU.TD_type_ENDLM:
		// NrOpenTags--
		// pTag.Depth = NrOpenTags
		// var TT lwdx.TagSummary
		var pTE *lwdx.TagalogEntry
		theTag := pTag.GToken.XName.Local
		pTE = lwdx.GetTEbyTagAndMarkupType(theTag, pTag.GToken.MarkupType)
		if pTE == nil {
			// TODO
		}
		pTag.TagalogEntry = pTE
		return pTag, nil

	case XU.TD_type_PINST:
		// !! pTag.TagSummary = lwdx.TTblock
		// TODO: Attach this PI to its parent element in the GTree
		// newNode = parentNode.NewKid("<?", myTarget)
		// newNode.StringValue = myInst
		return pTag, nil

	case XU.TD_type_CDATA:
		if pTag.Echo() == "" {
			// ilog.Printf("PCDATA is all whitespace: \n")
			// DO NOTHING
			// NIL IT OUT
			// NOTE:550 This may do weird things to elements
			// that have text content models.
			// println("WARNING: Got an all-whitespace xml.CharData")
			return nil, nil
		}
		// !! pTag.TagSummary = lwdx.TTinline
		return pTag, nil

	case XU.TD_type_COMNT:
		// TODO: Attach this Comment to its parent element in the GTree,
		// which is in fact the first XML element that follows this XML
		// comment. So, it should be done in a later pass.
		// newNode = parentNode.NewKid("<!", "--")
		// newNode.StringValue = tokenString
		// !! pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	case XU.TD_type_DRCTV:
		// !! pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	case "":
		print(inGTkn.String())
		pTag.TDType = XU.TD_type_ERROR
		// !! pTag.TagSummary = lwdx.TTblock
		// println(fmt.Sprintf("NIL GToken.type<%s> for: %+v", inGTkn.TDType, inGTkn))
		return nil, fmt.Errorf("NIL GToken.type<%s> for: %+v", inGTkn.TDType, inGTkn)

	case XU.TD_type_DOCMT:
		L.L.Dbg("Made GTag for GToken TDType <Doc>")
		// !! pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	default:
		print(inGTkn.String())
		pTag.TDType = XU.TD_type_ERROR
		// !! pTag.TagSummary = lwdx.TTblock
		println(fmt.Sprintf("Unrecognized GToken.type<%s> for: %+v",
			inGTkn.TDType, inGTkn))
		return nil, fmt.Errorf("Unrecognized GToken.type<%s> for: %+v",
			inGTkn.TDType, inGTkn)
	}
}
