// Copyright 2017 The WL Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wl

import (
	"bufio"
	"flag"
	"fmt"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cznic/golex/lex"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "# caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# \tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "# dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string { //TODOOK
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("# TODO: %s:%d:\n", path.Base(fn), fl) //TODOOK
}

func use(...interface{}) {}

func init() {
	use(caller, dbg, TODO) //TODOOK
}

// ============================================================================

func init() {
	flag.IntVar(&yyDebug, "yydebug", 0, "")
}

func exampleAST(rule int, src string) interface{} {
	lx := newLexer(strings.NewReader(src))
	l, err := lex.New(
		token.NewFileSet().AddFile(fmt.Sprint(rule), -1, len(src)),
		lx,
		lex.BOMMode(lex.BOMIgnoreFirst),
		lex.RuneClass(runeClass),
		lex.ErrorFunc(func(token.Pos, string) {}),
	)
	if err != nil {
		return err.Error()
	}

	lx.exampleRule = rule
	lx.parse(l, false)
	return prettyString(lx.exampleAST)
}

func testScannerCorpus(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "corpus"))
	if err != nil {
		t.Log(err)
		return
	}

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	file := token.NewFileSet().AddFile(f.Name(), -1, int(fi.Size()))
	r := bufio.NewReader(f)
	lx := newLexer(r)
	l, err := lex.New(
		file,
		lx,
		lex.BOMMode(lex.BOMIgnoreFirst),
		lex.RuneClass(runeClass),
		lex.ErrorFunc(func(pos token.Pos, msg string) { t.Fatalf("%s: %s", file.Position(pos), msg) }),
	)
	if err != nil {
		t.Fatal(err)
	}

	lx.Lexer = l
	toks := 0
	for lx.Last.Rune >= 0 {
		lx.scan()
		toks++
	}
	if _, err := r.Peek(1); err != io.EOF {
		t.Fatal(err)
	}

	t.Logf("tokens: %v", toks)
}

func TestScanner(t *testing.T) {
	t.Run("corpus", testScannerCorpus)
}

func testParseCorpus(t *testing.T, interactive bool) {
	f, err := os.Open(filepath.Join("testdata", "corpus"))
	if err != nil {
		t.Log(err)
		return
	}

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}

	testFile = token.NewFileSet().AddFile(f.Name(), -1, int(fi.Size()))
	defer func() { testFile = nil }()

	r := bufio.NewReader(f)
	p, err := NewInput(r, interactive)
	if err != nil {
		t.Fatal(err)
	}

	n := 0
	for {
		_, err := r.Peek(1)
		if err != nil {
			break
		}

		if _, err = p.ParseExpression(testFile); err != nil {
			t.Fatal(err)
		}

		n++
	}
	t.Logf("%s: expressions: %v", testFile.Position(p.lex.First.Pos()), n)
}

func testParseOther(t *testing.T) {
	for i, v := range []string{
		"i;;j",
		"i;;",
		";;j",
		";;",
		"i;;j;;k",
		"i;;;;k",
		";;j;;k",
		";;;;k",

		"#",
		"#1",
		"#string",
		"##",
		"###",

		"%",
		"%%",
		"%%%",
		"%1",

		"_",
		"_42",
		"__",
		"__42",
		"___",
		"___42",
		"_.",

		"f_",
		"f_42",
		"f__",
		"f__42",
		"f___",
		"f___42",
		"f_.",

		"<< foo",
		`<< "foo bar"`,
		"42 >> foo",
		`42 >> "foo bar"`,
		"42 >>> foo",
		`42 >>> "foo bar"`,
	} {
		lx := newLexer(strings.NewReader(v))
		l, err := lex.New(
			token.NewFileSet().AddFile(fmt.Sprint(i), -1, len(v)),
			lx,
			lex.ErrorFunc(func(token.Pos, string) {}),
		)
		if err != nil {
			t.Fatal(err)
		}

		if err := lx.parse(l, false); err != nil {
			t.Errorf("#%v: %v", i, err)
		}
	}
}

func TestParser(t *testing.T) {
	t.Run("corpus bulk", func(t *testing.T) { testParseCorpus(t, false) })
	t.Run("corpus interactive", func(t *testing.T) { testParseCorpus(t, true) })
	t.Run("other", func(t *testing.T) { testParseOther(t) })
}
