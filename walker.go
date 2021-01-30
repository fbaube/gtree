package gtree

type WalkerData struct {
	Parent    string
	Tag       string
	Att       string
	FoundTags []*GTag
}

var WalkerParent string
var WalkerTag string
var WalkerAtt string
var WalkerFoundTags []*GTag

// type GTagWalkable func(*GTag) bool

// var theWalkFunc =
func (E *GTag) TheWalkFunc() bool {
	// Call it on every kid
	var hasKids = (E.FirstKid() != nil)
	if hasKids {
		var pKid = E.FirstKid()
		for pKid != nil {
			pKid.TheWalkFunc()
			pKid = pKid.NextKid()
		}
	}
	// Do on self
	println("Checking:", E.GToken.GName.Local)
	if WalkerTag != E.GToken.GName.Local {
		return hasKids
	}
	if WalkerParent != "" {
		if E.GetParent() == nil {
			return hasKids
		}
		if E.GetParent().GToken.GName.Local != WalkerParent {
			return hasKids
		}
	}
	if WalkerAtt != "" {
		if "" == E.GAtts.GetAttVal(WalkerAtt) {
			return hasKids
		}
	}
	WalkerFoundTags = append(WalkerFoundTags, E)
	println("Found!", E.GToken.GName.Local)
	return hasKids
}
