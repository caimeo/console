package console

import (
	"strings"
	"testing"
)

type wrtr struct {
	c chan []byte
}

func (w *wrtr) Write(p []byte) (n int, err error) {
	w.c <- p
	return len(p), nil
}

func newWrtr() *wrtr {
	w := new(wrtr)
	w.c = make(chan []byte, 1)
	return w
}

func TestRedirectIO(t *testing.T) {
	wo := newWrtr()
	we := newWrtr()
	RedirectIO(wo, we)

	stdOutMsg := "My text is long and beautiful, It's FAKE TEXT!"
	stdErrMsg := "OMG it's an ERROR!!!"

	StdOut(stdOutMsg)
	outstring := string(<-wo.c)

	StdErr(stdErrMsg)
	errstring := string(<-we.c)

	if strings.TrimSpace(outstring) != stdOutMsg {
		t.Error("Redirect std out failed", outstring)
	}
	if strings.TrimSpace(errstring) != stdErrMsg {
		t.Error("Redirect std err failed", errstring)
	}
}

func TestIsVerbose(t *testing.T) {
	con := New(true, false)
	if !con.IsVerbose() {
		t.Error("Expected IsVerbose to be true")
	}
	if con.IsDebug() {
		t.Error("Expected IsDebug to be false")
	}
}

func TestIsDebug(t *testing.T) {
	con := New(false, true)
	if con.IsVerbose() {
		t.Error("Expected IsVerbose to be false")
	}
	if !con.IsDebug() {
		t.Error("Expected IsDebug to be true")
	}
}
