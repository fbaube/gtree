// Package gtree defines low-level structures for Generic Golang XML analysis,
// structures that are built directly atop (or map'd directly to) Golang's own
// XML structures. We do this for three reasons:
// - So that we can define our own helper methods in our own golang namespace;
// - Cos golang's XML package uses lousy naming (`Name Name`, anyone ?);
// - Cos golang's XML is written for XML data records, not for mixed content.
//
// This repo implements a protocol stack that goes:
// - the input text file (i.e. XML or other markup)
// - package encoding/xml (golang's XML package)
// - package gparse (low-level stuff)
// - package gfile (deep analysis of individual markup files)
// - package mmmc (processing of heterogeneous markup files)
//
// In general, all go files in this protocol stack should be organised as: <br/>
// - struct definition()
// - constructors (named `New*`)
// - printf stuff (Raw(), Echo(), String())
//
// Some characteristic methods:
// - Raw() returns the original string passed from the golang XML parser
// - Echo() returns a string of the item in normalised form, altho the
// presence of terminating newlines is not uniform
// String() returns a string suitable for runtime nonitoring and debugging
//
// NOTE:1280 the use of shorthand in variable names: Doc, Elm, Att.
//
// NOTE:1220 that we store non-nil namespaces with a colon appended, for easy output.
//
// NOTE:1230 that we use `godoc2md`, so we can use Markdown in these code comments.
//
// NOTE:1030 that like other godoc comments, this package comment must be *right*
// above the target statement (`package`) if it is to be included by `godoc2md`.
//
package gtree
