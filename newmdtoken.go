package gtree

import "github.com/fbaube/gtoken"

/*

func NewGTagFromMDtoken(M MD.Token) (*GTag, error) {
	switch mdtok := M.(type) {
	case *MD.CodeBlock: // pre><code
	case *MD.CodeInline: // code
	case *MD.Fence:
	case *MD.Hardbreak: // br
	case *MD.Hr: // hr
	case *MD.Image:
	case *MD.LinkClose: // a
	case *MD.LinkOpen: // a href=
	case *MD.Softbreak: // br
	case *MD.Text:

	case *MD.HTMLBlock:
		return CDataTagMD(mdtok.Content), nil
	case *MD.HTMLInline:
		return CDataTagMD(mdtok.Content), nil
	case *MD.HeadingClose: // h + level
		return EndTagMD("h" + strconv.Itoa(mdtok.HLevel)), nil
	case *MD.HeadingOpen: // h + level
		return StartTagMD("h" + strconv.Itoa(mdtok.HLevel)), nil
	case *MD.Inline:
		return CDataTagMD(mdtok.Content), nil

	case *MD.BlockquoteClose:
		return EndTagMD("blockquote"), nil
	case *MD.BlockquoteOpen:
		return StartTagMD("blockquote"), nil
	case *MD.BulletListClose:
		return EndTagMD("ul"), nil
	case *MD.BulletListOpen:
		return StartTagMD("ul"), nil
	case *MD.EmphasisClose:
		return EndTagMD("em"), nil
	case *MD.EmphasisOpen:
		return StartTagMD("em"), nil
	case *MD.ListItemClose:
		return EndTagMD("li"), nil
	case *MD.ListItemOpen:
		return StartTagMD("li"), nil
	case *MD.OrderedListClose:
		return EndTagMD("ol"), nil
	case *MD.OrderedListOpen:
		return StartTagMD("ol"), nil
	case *MD.ParagraphClose:
		return EndTagMD("p"), nil
	case *MD.ParagraphOpen:
		return StartTagMD("p"), nil
	case *MD.StrikethroughClose:
		return EndTagMD("s"), nil
	case *MD.StrikethroughOpen:
		return StartTagMD("s"), nil
	case *MD.StrongClose:
		return EndTagMD("strong"), nil
	case *MD.StrongOpen:
		return StartTagMD("strong"), nil
	case *MD.TableClose:
		return EndTagMD("table"), nil
	case *MD.TableOpen:
		return StartTagMD("table"), nil
	case *MD.TbodyClose:
		return EndTagMD("tbody"), nil
	case *MD.TbodyOpen:
		return StartTagMD("tbody"), nil
	case *MD.TdClose:
		return EndTagMD("td"), nil
	case *MD.TdOpen:
		return StartTagMD("td"), nil
	case *MD.ThClose:
		return EndTagMD("th"), nil
	case *MD.ThOpen:
		return StartTagMD("th"), nil
	case *MD.TheadClose:
		return EndTagMD("thead"), nil
	case *MD.TheadOpen:
		return StartTagMD("thead"), nil
	case *MD.TrClose:
		return EndTagMD("tr"), nil
	case *MD.TrOpen:
		return StartTagMD("tr"), nil
	}
	return nil, nil
}

*/

func StartTagMD(tag string) *GTag {
	var pTag = new(GTag)
	pTag.TTType = gtoken.TT_type_ELMNT
	pTag.TagOrPrcsrDrctv = tag
	// pTag.AsString = "<" + tag + ">"
	// nOpenTags++
	return pTag
}

func EndTagMD(tag string) *GTag {
	var pTag = new(GTag)
	pTag.TTType = gtoken.TT_type_ENDLM
	pTag.TagOrPrcsrDrctv = tag
	// pTag.AsString = "</" + tag + ">"
	// nOpenTags++
	return pTag
}

func CDataTagMD(content string) *GTag {
	var pTag = new(GTag)
	pTag.TTType = gtoken.TT_type_CDATA
	/*
		pTag.AsString = S.TrimSpace(content)
		pTag.Keyword = pTag.AsString
		if pTag.AsString == "" {
			println("WARNING: Got an all-whitespace xml.CharData")
			return nil
		}
	*/
	return pTag
}
