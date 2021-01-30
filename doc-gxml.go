// Package mmmc contains generic Golang XML stuff:
// names, attributes, tags, elements, trees, files, documents.
//
// Files in this directory use Markdown, so use `godoc2md` on 'em.
//
// We make our own versions of Golang XML structures so that we can give
// them sensible new names and we can define our own methods for them.
//
// We also use a couple of shortened names (Att *Attribute*, Elm *Element*)
// to keep code readable. )
//
// ### Method naming
//
// - `NewFoo(..)` always allocates memory.
// - `Echo()` echoes an object back in source XML form, but normailzed.
// - `EchoCommented()` also outputs XML source form, but possibly with
// additional annotations added by processing.
// - `String()`` outputs a form that is useful for development and
// debugging but cannot be processed by an XML parser.
//
// ### About encoding/decoding and XML mixed content
//
// When working with XML we can generally distinguish between two types of files:
// - Files containing record-oriented data - expressed using XML elements
// - Files containing natural language documents - also expressed using
// XML elements
// - Files containing validation rules - generally expressed as XSD,
// RNG, or DTD. It is interesting to note that DTDs actually obey the
// same syntax rules as the other two; the typical file extensions
// (`.dtd .mod`) are helpful to humans but are not required by a
// parser that fully understands XML syntax.
//
package gtree
