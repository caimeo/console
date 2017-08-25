package console

import (
	"fmt"
	"io"
	"os"
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
	StdOut(args)
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
	con.Always(args)
}

//writes to stdout if IsVerbose == true or IsDebug == true
func Verbose(args ...interface{}) {
	con.Verbose(args)
}

//writes to stdout if IsDebug == ture
func Debug(args ...interface{}) {
	con.Debug(args)
}

//always writes to stderr
func Error(args ...interface{}) {
	con.Error(args)
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
	var collect string
	for _, v := range args {
		collect = collect + fmt.Sprint(v)
	}
	fmt.Fprintf(stdWriter, "%s\n", collect)
}

//Outputs to standard error as a single line
func StdErr(args ...interface{}) {
	if errWriter == nil {
		initWriters()
	}
	var collect string
	for _, v := range args {
		collect = collect + fmt.Sprint(v)
	}
	fmt.Fprintf(errWriter, "%s\n", collect)
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
