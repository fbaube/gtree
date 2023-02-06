package gtree

import (
	"fmt"

	"github.com/fbaube/gtoken"
	// "github.com/fbaube/lwdx"
	SU "github.com/fbaube/stringutils"
	"golang.org/x/net/html"
	// "github.com/dimchansky/utfbom"
)

// NewGTagFromHtmltoken is TODO.
// TODO: Pass a writer for Echo.
// NOTE: Returns "nil" if the token is valid but useless, and can
// be skipped, such as an xml.CharData that is all whitespace;
// NOTE: that it might cause problems.
// .
func NewGTagFromHtmlToken(T html.Token) (pTag *GTag, e error) {
	pTag = new(GTag)
	// pTkn := &(pTag.GToken)

	var TS string
	var GT *gtoken.GToken
	GT, e = NewGTokenFromHtmlToken(T)

	switch GT.TTType {

	case gtoken.TT_type_ELMNT:
		// Create new GTag. Input:
		// type StartElement struct { Name Name ; Attr []Attr }
		// type Attr struct { Name  Name ; Value string }
		// pTag.GStartTag = *new(gparse.GToken.GName)
		// pTag.GStartTag.Name.Local = TS
		pTag.TagOrPrcsrDrctv = TS
		// pTag.AsString = pTag.GStartTag.String() // include "<" and ">"
		// pTag.Depth  = pET.NrOpenTags
		fmt.Printf("cnvt: %s<%s>\n",
			SU.GetIndent(0 /*pET.NrOpenTags*/), pTag.TagOrPrcsrDrctv)
		// pET.NrOpenTags++
		// if TT, ok = lwdx.TagTypes[pTag.keynoun]; !ok {
		// elog.Printf("Unrecognized tag: <" + pTag.keynoun + ">")
		// return pTag, errors.New("Unrecognized tag: <" + pTag.keynoun + ">")
		// }
		// pTag.TagType = TT
		return pTag, nil

	case gtoken.TT_type_ENDLM:
		// pRT.pTag = nil // !!
		// type EndElement struct { Name Name }
		pTag.TagOrPrcsrDrctv = TS
		// pTag.AsString = "</" + pTag.TagOrPrcsrDrctv + ">"
		// pET.NrOpenTags--
		// pTag.Depth  = pET.NrOpenTags
		fmt.Printf("%s</%s>\n",
			SU.GetIndent(0 /*pET.NrOpenTags*/), pTag.TagOrPrcsrDrctv)
		/* old code
		var TT lwdx.TagType
		var ok bool
		if TT, ok = xmltags.TagTypes[pTag.keynoun]; !ok {
		}
		pTag.TagType = TT
		*/
		return pTag, nil

	case gtoken.TT_type_CDATA:
		// We seem to have trouble making a genuine copy of the string.
		// So, take an extra step or two to make sure it is correct.
		// pTag.AsString = TS
		// pTag.Keytext  = pTag.AsString
		// println("AFTER TrimSpace<<" + pRT.string1 + ">>")

		// FIXME
		/* old code
		if pTag.AsString == "" {
			// ilog.Printf("PCDATA is all whitespace: \n")
			// DO NOTHING
			// NIL IT OUT
			// NOTE This may do weird things to elements
			// that have text content models.
			println("WARNING: Got an all-whitespace xml.CharData")
			return nil, nil
		} */
		// fmt.Printf("%s(cdata)|%s|\n",
		// 	SU.GetIndent(0/*pET.NrOpenTags*/), pTag.AsString)
		// !! pTag.TagSummary = lwdx.TTinline
		return pTag, nil

	case gtoken.TT_type_COMNT:
		pTag.TagOrPrcsrDrctv = TS
		// pTag.AsString = "<--" + pTag.TagOrPrcsrDrctv + "-->"
		// println("ok:", pTag.AsString) // " <--|" + pRT.string1 + "|--> \n")
		// newNode = parentNode.NewKid("<!", "--")
		// newNode.StringValue = tokenString
		// !! pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	case gtoken.TT_type_DRCTV:
		s := TS
		// pTag.AsString = "<!" + s + ">"
		pTag.TagOrPrcsrDrctv, pTag.Datastring = SU.SplitOffFirstWord(s) // pRT.string1)
		// println("ok:", pTag.AsString) // " <!|" + pRT.string1 + "|" + pRT.string2 + "|> \n")
		// newNode = parentNode.NewKid("<!", "DOCTYPE")
		// newNode.StringValue = tokenString
		// !! pTag.TagSummary = lwdx.TTblock
		return pTag, nil

	default:
		pTag.TTType = gtoken.TT_type_ERROR
		// !! pTag.TagSummary = lwdx.TTblock
		return nil, fmt.Errorf("Unrecognized token type<%T> for: %+v", T, T)
	}
}

// NewGTokenFromHtmlToken does not recognise and return Processing Instrucitons !
func NewGTokenFromHtmlToken(inT html.Token) (outT *gtoken.GToken, e error) {

	return nil, nil
	/* old code
	outT = new(gxml.GToken)

	var TT   html.TokenType
	var Str  string
	var Atts []html.Attribute
	TT = inT.Type
	Str = inT.Data
	Atts = inT.Attr

	var isTagType = (
		TT == html.StartTagToken ||
		TT == html.EndTagToken ||
		TT == html.SelfClosingTagToken)
	var isTextType = (
		TT == html.TextToken ||
		TT == html.CommentToken ||
		TT == html.DoctypeToken)

	if isTagType {
		outT.Keynoun = Str
		var slash1, slash2 string
		if TT == html.EndTagToken { slash1 = "/" }
		if TT == html.SelfClosingTagToken { slash2 = "/" }
		fmt.Printf("HtmlTag: <%s%s%s> @<%+v> \n", slash1, Str, slash2, Atts)
	} else if isTextType {
		fmt.Printf("Html text: |%s| \n", Str)
	} else {
		panic("gxml.html.Tagtree.L321")
	}

	switch TT {

	case html.StartTagToken: // xml.StartElement:
		outT.GTagTokType = "SE"
		outT.GTag = *new(gxml.GTag)
		outT.GTag.Name.Local = Str
		outT.AsString = outT.GTag.String() // include "<" and ">"
		return outT, nil

	case html.EndTagToken: // xml.EndElement:
		outT.GTagTokType = "EE"
		outT.AsString = "</" + outT.Keynoun + ">"
		return outT, nil

	case html.SelfClosingTagToken: // xml.EndElement:
		outT.GTagTokType = "SC"
		outT.AsString = "<" + outT.Keynoun + "/>"
		return outT, nil

	case html.TextToken: // xml.CharData: // type CharData []byte
		outT.GTagTokType = "CD"
		outT.AsString = Str
		outT.Keynoun  = outT.AsString
		// println("AFTER TrimSpace<<" + pRT.string1 + ">>")
		if outT.AsString == "" {
			// ilog.Printf("PCDATA is all whitespace: \n")
			// DO NOTHING
			// NIL IT OUT
			// NOTE:540 This may do weird things to elements
			// that have text content models.
			println("WARNING: Got an all-whitespace xml.CharData")
			return nil, nil
		}
		fmt.Printf("%s(cdata)|%s|\n", SU.GetIndent(nOpenTags), outT.asString)
		return outT, nil

	case html.CommentToken: // xml.Comment: // type Comment []byte
		outT.Type = "Cmt"
		outT.keynoun = Str
		outT.AsString = "<--" + outT.keynoun + "-->"
		println("ok:", outT.asString) // " <--|" + pRT.string1 + "|--> \n")
		return outT, nil

	case html.DoctypeToken: // xml.Directive: // type Directive []byte
		outT.Type = "Dir"
		s := Str
		outT.AsString = "<!" + s + ">"
		outT.keynoun, outT.keyargs = SU.SplitOffFirstWord(s) // pRT.string1)
		println("ok:", outT.asString) // " <!|" + pRT.string1 + "|" + pRT.string2 + "|> \n")
		return outT, nil

	default:
		outT.Type = "ERR"
		return nil, fmt.Errorf("Unrecognized token type<%T> for: %+v", inT, inT)
	}
	*/
}
