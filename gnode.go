package gtree

import (
	"fmt"
	"os"
)

// This file: Structures for Generic Golang XML tree Nodes.

// GNode contains references to `GTag`s but could probably be rewritten
// to point to `GNode`s. It does not implement `interface Markupper`.
//
// Note that this basic interface - the idea of a *Node* - is defined all
// over teh interwebz. Note also that if you want a `GNode` to implement
// some other conception of a Node, it would be pretty straightforward
// to write an adapter.
//
// TODO The right way to do this would be to define the methods to return
// pointers to `GNode`s, but then if embedded in another struct (i.e. `GTag`),
// then override the methods.
//
// *Implementation note:* We use a doubly-linked list, not a slice.
//
type GNode struct {
	parent            *GTag
	firstKid, lastKid *GTag
	prevKid, nextKid  *GTag
}

type NodeIfc interface {
	GetParent() NodeIfc
	AddKid(NodeIfc) NodeIfc
}

// GetParent returns the parent, duh.
func (E *GTag) GetParent() *GTag {
	return E.parent
}

// NOTE:
// https://godoc.org/golang.org/x/net/html#Node
// func (n *Node) AppendChild(c *Node)
// func (n *Node) InsertBefore(newChild, oldChild *Node)
// func (n *Node) RemoveChild(c *Node)

// AddKid adds the supplied node as the last kid,
// and returns it (i.e. the last kid).
func (anE *GTag) AddKid(aKid *GTag) *GTag {
	if aKid.prevKid != nil || aKid.nextKid != nil {
		fmt.Fprintf(os.Stdout, "FATAL in AddKid: Tag<< %+v >> kid<< %+v >>\n", anE, aKid)
		panic("AddKid(K) can't cos K has siblings")
	}
	if aKid.parent != nil && aKid.parent != anE {
		fmt.Fprintf(os.Stdout, "FATAL in AddKid: Tag<< %+v >> kid<< %+v >>\n", anE, aKid)
		panic("E.AddKid(K) can't cos K has non-P parent")
	}
	var FK = anE.firstKid
	var LK = anE.lastKid
	// Is the new kid an only kid ?
	if FK == nil && LK == nil {
		anE.firstKid, anE.lastKid = aKid, aKid
		aKid.parent = anE
		aKid.prevKid, aKid.nextKid = nil, nil
		return aKid
	}
	// So, replace the last kid
	if LK != nil {
		if LK.parent != anE {
			fmt.Fprintf(os.Stdout, "FATAL in AddKid: E<< %+v >> K<< %+v >>\n", anE, aKid)
			panic("E.AddKid: E's last kid dusnt know E")
		}
		if LK.nextKid != nil {
			fmt.Fprintf(os.Stdout, "FATAL in AddKid: E<< %+v >> K<< %+v >>\n", anE, aKid)
			panic("E.AddKid: E's last kid has a next kid")
		}
		LK.nextKid = aKid
		aKid.prevKid = LK
		anE.lastKid = aKid
		aKid.parent = anE
		return aKid
	}
	fmt.Fprintf(os.Stdout, "FATAL in AddKid: E<< %+v >> K<< %+v >>\n", anE, aKid)
	panic("AddKid: Chaos!")
}
