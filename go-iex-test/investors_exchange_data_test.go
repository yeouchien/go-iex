package goiextest

import (
	"testing"
)

func TestTOPS(t *testing.T) {
	tops, err := iexSandboxClient.TOPS(nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(tops) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	tops, err = iexSandboxClient.TOPS(struct {
		Symbols string `url:"symbols,omitempty"`
	}{"SNAP,fb,AIG+"})
	if err != nil {
		t.Error(err)
	}
	expected = "SNAP"
	actual = tops[0].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "FB"
	actual = tops[1].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "AIG+"
	actual = tops[2].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}

func TestLast(t *testing.T) {
	last, err := iexSandboxClient.Last(nil)
	if err != nil {
		t.Error(err)
	}
	expected = false
	actual = len(last) == 0
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}

	last, err = iexSandboxClient.Last(struct {
		Symbols string `url:"symbols,omitempty"`
	}{"SNAP,fb,AIG+"})
	if err != nil {
		t.Error(err)
	}
	expected = "SNAP"
	actual = last[0].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "FB"
	actual = last[1].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
	expected = "AIG+"
	actual = last[2].Symbol
	if expected != actual {
		t.Errorf("\nExpected: %v\nActual: %v\n", expected, actual)
	}
}
