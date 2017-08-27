/*
cameo/console

A simple configurable output console

Allows initialization with verbose and/or debug options.

Use throughout your program code to output simple information to the console.

Intro

While developing and debuging it is often useful to output debug information to the console that would not normally be visible to the user. It is also often desired to include a verbose mode of operation for command line utilities. This utility library provides a simple way to acomplish this as well as output to Standard Out and Standard Error.

Messages types:

- Always: messages are always output to Standard Out

- Error: messages are always sent to Standard Error

- Debug: messages are only output to Standard Out if IsDebug() == true

- Verbose: messages are only output to Standard Out if IsVerbose() == true or IsDebug() == true


Redirect StdOut and StdErr

By default messages are sent to os.Stdout and os.Stderr. If you desire to send messages to another location use RedirectIO(io.writer, io.writer)

To connect to a logging system simply provide an io.writer that outputs to your logger.

*/
package console

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

type Console interface {
	Always(args ...interface{})
	Verbose(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	IsVerbose() bool
	IsDebug() bool
}

type console struct {
	verbose bool
	debug   bool
	ready   bool
}

//always writes to stdout
func (t *console) Always(args ...interface{}) {
	if t == nil {
		return
	}
	StdOut(args...)
}

//writes to stdout if IsVerbose == true or IsDebug == true
func (t *console) Verbose(args ...interface{}) {
	if t == nil {
		return
	}
	if t.verbose || t.debug {
		StdOut(args...)
	}
}

//writes to stdout if IsDebug == ture
func (t *console) Debug(args ...interface{}) {
	if t == nil {
		return
	}
	if t.debug {
		StdOut(args...)
	}
}

//always writes to stderr
func (t *console) Error(args ...interface{}) {
	if t == nil {
		return
	}
	StdErr(args)
}

//returns true if verbose mode is turned on
func (t *console) IsVerbose() bool {
	if t == nil {
		return false
	}
	return t.verbose
}

//returns true if debug mode is turned on
func (t *console) IsDebug() bool {
	if t == nil {
		return false
	}
	return t.debug
}

//always writes to stdout
func Always(args ...interface{}) {
	con.Always(args...)
}

//writes to stdout if IsVerbose == true or IsDebug == true
func Verbose(args ...interface{}) {
	con.Verbose(args...)
}

//writes to stdout if IsDebug == ture
func Debug(args ...interface{}) {
	con.Debug(args...)
}

//always writes to stderr
func Error(args ...interface{}) {
	con.Error(args...)
}

//returns true if verbose mode is turned on
func IsVerbose() bool {
	return con.IsVerbose()
}

//returns true if debug mode is turned on
func IsDebug() bool {
	return con.IsDebug()
}

//Outputs to standard out as a single line
func StdOut(args ...interface{}) {
	if stdWriter == nil {
		initWriters()
	}
	args = byteSlicesToHex(args...)
	fmt.Fprintln(stdWriter, args...)
}

//Outputs to standard error as a single line
func StdErr(args ...interface{}) {
	if errWriter == nil {
		initWriters()
	}
	args = byteSlicesToHex(args...)
	fmt.Fprintln(errWriter, args...)
}

func argsToString(args ...interface{}) (collect string) {
	for _, v := range args {
		collect = collect + fmt.Sprint(v)
	}
	return collect
}

var con console

var stdWriter io.Writer
var errWriter io.Writer

// Allows redirection of std and err to arbitrary io writers
func RedirectIO(stdIoWriter io.Writer, errIoWriter io.Writer) {
	stdWriter = stdIoWriter
	errWriter = errIoWriter
	initWriters()
}

func initWriters() {
	if stdWriter == nil {
		stdWriter = os.Stdout
	}
	if errWriter == nil {
		errWriter = os.Stderr
	}

}

//utiltity to convert all byte slices to hex strings
//because displaying raw byte values is not real useful
func byteSlicesToHex(args ...interface{}) (ret []interface{}) {
	for _, v := range args {
		switch i := v.(type) {
		case []byte:
			ret = append(ret, "0x"+strings.ToUpper(hex.EncodeToString(i)))
		default:
			ret = append(ret, v)

		}
	}
	return ret
}

func newConsole(verbose bool, debug bool) console {
	initWriters()
	c := console{verbose: verbose, debug: debug, ready: true}
	return c
}

func New(verbose bool, debug bool) Console {
	c := newConsole(verbose, debug)
	return &c
}

func Init(verbose bool, debug bool) Console {
	if con.ready == false {
		con = newConsole(verbose, debug)
	}
	return &con
}

func Instance() Console {
	return &con
}
