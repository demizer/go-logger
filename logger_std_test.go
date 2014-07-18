// Copyright 2013,2014 The go-logger Authors. All rights reserved.
// This code is MIT licensed. See the LICENSE file for more info.

// Tests for the default standard logging object

package log

import (
	"bytes"
	"testing"
	"time"
)

func TestStdTemplate(t *testing.T) {
	var buf bytes.Buffer

	SetLevel(LEVEL_DEBUG)
	SetStreams(&buf)

	SetFlags(LdebugFlags)

	SetTemplate("{{.Text}}")
	temp := Template()

	type test struct {
		Text string
	}

	err := temp.Execute(&buf, &test{"Hello, World!"})
	if err != nil {
		t.Fatal(err)
	}

	expe := "Hello, World!"

	if buf.String() != expe {
		t.Errorf("\nGot:\t%s\nExpect:\t%s\n", buf.String(), expe)
	}
}

func TestStdSetTemplate(t *testing.T) {
	var buf bytes.Buffer

	SetLevel(LEVEL_DEBUG)
	SetStreams(&buf)

	SetFlags(LdebugFlags)

	SetTemplate("{{.Text}}")

	Debugln("Hello, World!")

	expe := "Hello, World!\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expe)
	}
}

func TestStdSetTemplateBad(t *testing.T) {
	var buf bytes.Buffer

	SetLevel(LEVEL_DEBUG)
	SetStreams(&buf)

	SetFlags(LdebugFlags)

	err := SetTemplate("{{.Text")

	Debugln("template: default:1: unclosed action")

	expe := "template: default:1: unclosed action"

	if err.Error() != expe {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expe)
	}
}

func TestStdSetTemplateBadDataObjectPanic(t *testing.T) {
	var buf bytes.Buffer

	SetLevel(LEVEL_DEBUG)

	SetStreams(&buf)

	SetFlags(LnoPrefix | Lindent)

	SetIndent(1)

	type test struct {
		Test string
	}

	err := SetTemplate("{{.Tes}}")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("\nGot:\t%q\nExpect:\tPANIC\n", buf.String())
		}
	}()

	Debugln("Hello, World!")

	// Reset the standard logging object
	SetTemplate(logFmt)
	SetIndent(0)
}

func TestStdDateFormat(t *testing.T) {
	dateFormat := DateFormat()

	expect := "Mon-20060102-15:04:05"

	if dateFormat != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", dateFormat, expect)
	}
}

func TestStdSetDateFormat(t *testing.T) {
	var buf bytes.Buffer

	SetLevel(LEVEL_ALL)

	SetStreams(&buf)

	SetFlags(Ldate)

	SetDateFormat("20060102-15:04:05")

	SetTemplate("{{.Date}}")

	Debugln("Hello")

	expect := time.Now().Format(DateFormat())

	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}

	// Reset the standard logging object
	SetTemplate(logFmt)
}

func TestStdFlags(t *testing.T) {
	SetFlags(LstdFlags)

	flags := Flags()

	expect := LstdFlags

	if flags != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", flags, expect)
	}
}

func TestStdLevel(t *testing.T) {
	SetLevel(LEVEL_DEBUG)

	level := Level()

	expect := "LEVEL_DEBUG"

	if level.String() != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", level, expect)
	}
}

func TestStdPrefix(t *testing.T) {
	SetPrefix("TEST::")

	prefix := Prefix()

	expect := "TEST::"

	if prefix != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", prefix, expect)
	}
}

func TestStdStreams(t *testing.T) {
	var buf bytes.Buffer

	SetStreams(&buf)

	bufT := Streams()

	if &buf != bufT[0] {
		t.Errorf("\nGot:\t%p\nExpect:\t%p\n", &buf, bufT[0])
	}
}

func TestStdIndent(t *testing.T) {
	var buf bytes.Buffer

	SetStreams(&buf)

	SetLevel(LEVEL_DEBUG)

	SetFlags(LnoPrefix | Lindent)

	SetIndent(0).Debugln("Test 1")
	SetIndent(2).Debugln("Test 2")

	indent := Indent()

	expe := "[DEBG] Test 1\n[DEBG]         Test 2\n"
	expi := 2

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}

	if indent != expi {
		t.Errorf("\nGot:\t%d\nExpect:\t%d\n", indent, expi)
	}
}

func TestStdTabStop(t *testing.T) {
	var buf bytes.Buffer

	SetStreams(&buf)

	SetLevel(LEVEL_DEBUG)

	SetFlags(LnoPrefix | Lindent)

	// This SetIndent doesn't have to be on a separate line, but for some
	// reason go test cover wasn't registering its usage when the functions
	// below were chained together.
	SetIndent(1)
	SetTabStop(2).Debugln("Test 1")

	SetIndent(2)
	SetTabStop(4).Debugln("Test 2")

	tabStop := TabStop()

	expe := "[DEBG]   Test 1\n[DEBG]         Test 2\n"
	expt := 4

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}

	if tabStop != expt {
		t.Errorf("\nGot:\t%d\nExpect:\t%d\n", tabStop, expt)
	}
}