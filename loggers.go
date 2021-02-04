package gtree

import (
	"io"
	"log"
	"os"
)

// This file supports a logging model where
// - Two levels are supported, Info and Error
// - They default to os.Stdout and os.Stderr
// - Each can optionally be routed to an io.Writer by
// the calling user (such as a package main command)
// - The loggers are otherwise self-contained,
// needing no other initialization or configuration
//
// To use this file, for example for package
// `github.com/myusername/mypkg`:
// 1. Copy THIS file into the target package directory
// (`$GOPATH/src/github.com/myusername/mypkg/`)
// 2. Edit the string value for `var myPkgName` to be
// the package name ("mypkg"), or in some other way to
// uniquely identify the directory
// 3. (Optional) In the using function, to set the
// log data routing for ANY package that uses this, use
// io.Writer's in calls to anypkg.SetInfoLogDestination(w)
// and anypkg.SetErrorLogDestination(w)
// 4. Make your logging calls to ilog & elog, for example: <br/>
// ilog.Printf("Infile: %s%s.%s\n", p.DirPath, p.BaseName, p.In.Suffix)
// 5. Note that a long log message should start with a newline
//
var LoggersComment bool // Dummy variable to pick up the comment above

var myPkgName = "gtree"

// These are PACKAGE SCOPE, so do NOT capitalize these names.
var (
	ilog, elog       *log.Logger
	iwriter, ewriter io.Writer
)

func init() {
	setupInfoLogger()
	setupErrorLogger()
}

func setupInfoLogger() {
	ilog = new(log.Logger)
	ilog.SetOutput(os.Stdout)
	ilog.SetPrefix(myPkgName + " (info) ")
	ilog.SetFlags(log.Lshortfile) // log.Ldate | log.Ltime |
	// gxml.SetILog(ilog)
}

// SetInfoLoggerDestination takes an io.Writer.
func SetInfoLoggerDestination(w io.Writer) {
	ilog.SetOutput(w)
	ilog.Printf("Info logger starts here now \n")
	iwriter = w
}

// GetInfoWriter returns an io.Writer.
func GetInfoWriter() io.Writer {
	return (iwriter)
}

func setupErrorLogger() {
	elog = new(log.Logger)
	elog.SetOutput(os.Stderr)
	elog.SetPrefix(myPkgName + " (ERR!) ")
	elog.SetFlags(log.Lshortfile) // L.Ldate | L.Ltime |
	// gxml.SetELog(elog)
}

// SetErrorLoggerDestination takes an io.Writer.
func SetErrorLoggerDestination(w io.Writer) {
	elog.SetOutput(w)
	elog.Printf("Error logger starts here now \n")
	ewriter = w
}

// GetErrorWriter returns an io.Writer.
func GetErrorWriter() io.Writer {
	return (ewriter)
}
