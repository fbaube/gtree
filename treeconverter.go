package gtree

// We need to convert from Blackfriday Markdown AST to MMMC GTree.

// This new process has to be as general as possible, so:
// we convert a specific tree format to GTree ASAP.

// This means, first convert the tree (slapsash / quick 'n dirty),
// THEN sort out the screwed-up specifics of the tree source format.

// A TreeConvertedrtakes a specific tree source format, and returns
// a GTree.

// func ConvertBFMDtree(blackfriday.Node) gtree.GNode {..}

// The Walker from BFMD is awesome, so let's modify it a bit so that
// we can assume the walking order, and no entering/leaving distinction.
