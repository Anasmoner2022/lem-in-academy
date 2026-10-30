package main

import (
	"net/http"
	"testing"
)

func TestDocumentRejectsDuplicateAnchor(t *testing.T) {
	_, err := parseDoc([]byte(`<html><body><h1 id="same">A</h1><p id="same">B</p></body></html>`))
	if err == nil {
		t.Fatal("duplicate target would make navigation ambiguous")
	}
}
func TestDocumentRejectsBrokenStructure(t *testing.T) {
	_, err := parseDoc([]byte(`<html><body><p>unclosed</body></html>`))
	if err == nil {
		t.Fatal("broken generated markup must fail")
	}
}
func TestDocumentRetainsEscapedLinkAndFragment(t *testing.T) {
	doc, err := parseDoc([]byte(`<html><body><a href="next.html?x=1&amp;y=2#gate">Next</a><h1 id="gate">Gate</h1></body></html>`))
	if err != nil {
		t.Fatal(err)
	}
	if !doc.IDs["gate"] || len(doc.Links) != 1 || doc.Links[0] != "next.html?x=1&y=2#gate" {
		t.Fatalf("lost navigation information: %+v", doc)
	}
}
func TestStaticRouteFollowsLocalRedirect(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/index.html" {
			http.Redirect(w, r, "./", http.StatusMovedPermanently)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	status, hops, err := staticStatus(h, "/index.html")
	if err != nil || status != 200 || hops != 1 {
		t.Fatalf("%d %d %v", status, hops, err)
	}
}
func TestStaticRouteRejectsRedirectLoop(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/loop", http.StatusFound) })
	if _, _, err := staticStatus(h, "/loop"); err == nil {
		t.Fatal("redirect cycle must fail")
	}
}
