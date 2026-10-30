package main

import (
	"strings"
	"testing"
)

func TestReadDocumentAcceptance(t *testing.T) {
	t.Skip("TODO: remove when implementing the reader; add limit and injected failure cases")
	got, err := ReadDocument(strings.NewReader("a\r\nb"), Limits{MaxBytes: 10, MaxLine: 3})
	if err != nil || got.Echo != "a\nb\n" || len(got.Lines) != 2 {
		t.Fatalf("%+v, %v", got, err)
	}
}
