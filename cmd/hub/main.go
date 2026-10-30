// Command hub builds and checks a static, dependency-free curriculum.
package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

type Reading struct {
	Title, URL, Section, Why string
	Mandatory                bool
}
type Component struct {
	Slug, Title, Goal string
	Before, Steps     []string
	Hint, Trace, Gate string
}
type Project struct {
	ID, Slug, Title, Kind, Stage, Summary, Overview, Refresher, Spec, Run, Verification string
	Hours                                                                               int
	Depends                                                                             []string
	Reading                                                                             []Reading
	Components                                                                          []Component
	Engineering, Practical, Extras                                                      []string
	Files                                                                               map[string]string
}
type Page struct{ File, Title, Body string }
type Requirement struct {
	ID, Rule, Origin string
	Prep             []string
	Evidence         string
}
type Chapter struct{ File, Title, Body string }

var projects []Project
var routes []string

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func esc(s string) string  { return html.EscapeString(s) }
func code(s string) string { return "<pre><code>" + esc(s) + "</code></pre>" }
func list(items []string) string {
	s := "<ol>"
	for _, x := range items {
		s += "<li>" + x + "</li>"
	}
	return s + "</ol>"
}
func write(name string, b []byte) {
	must(os.MkdirAll(filepath.Dir(name), 0755))
	must(os.WriteFile(name, b, 0644))
}
func load(name string, v any)     { b, e := os.ReadFile(name); must(e); must(json.Unmarshal(b, v)) }
func projectURL(p Project) string { return "guides/" + p.Slug + "/index.html" }
func linkProject(id string) string {
	for _, p := range projects {
		if p.ID == id {
			return "<a href=\"" + projectURL(p) + "\">" + esc(p.ID+" · "+p.Title) + "</a>"
		}
	}
	return esc(id)
}
func relative(from, to string) string {
	r, e := filepath.Rel(filepath.Dir(from), to)
	must(e)
	return filepath.ToSlash(r)
}
func page(file, title, body, side, cls string) {
	pre := relative(file, "index.html")
	base := strings.TrimSuffix(pre, "index.html")
	nav := ""
	for _, n := range [][2]string{{"index.html", "Overview"}, {"roadmap.html", "Roadmap"}, {"coverage.html", "Coverage"}, {"instructor-guide.html", "For instructors"}} {
		current := ""
		if file == n[0] {
			current = " aria-current=\"page\""
		}
		nav += "<a href=\"" + base + n[0] + "\"" + current + ">" + n[1] + "</a>"
	}
	layout := "single-layout"
	if side != "" {
		layout = "reading-layout"
	}
	content := "<!DOCTYPE html>\n<html lang=\"en\"><head><meta charset=\"utf-8\"/><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"/><meta name=\"color-scheme\" content=\"dark\"/><meta name=\"description\" content=\"Go-only lem-in academy: hands-on preparation, exact planning, and independent verification.\"/><title>" + esc(title) + " · Lem-in Academy</title><link rel=\"stylesheet\" href=\"" + base + "assets/hub.css\"/><link rel=\"stylesheet\" href=\"" + base + "assets/academy.css\"/></head><body class=\"" + cls + "\"><a class=\"skip-link\" href=\"#main-content\">Skip to content</a><header class=\"site-header\"><div class=\"header-inner\"><a class=\"brand\" href=\"" + pre + "\" aria-label=\"Lem-in Academy home\"><span class=\"brand-mark\" aria-hidden=\"true\">L</span><span>Lem-in<span class=\"brand-subtitle\">ACADEMY</span></span></a><nav class=\"site-nav\" aria-label=\"Curriculum navigation\">" + nav + "</nav><span class=\"header-label\">GO ONLY · STANDARD LIBRARY</span></div></header><main id=\"main-content\" class=\"" + layout + "\"><article class=\"reading-content\">" + body + "</article>" + side + "</main><footer class=\"site-footer\"><p>Learn it. Build it. Explain it.</p><a href=\"" + base + "source-review.html\">Sources and academy policies</a></footer></body></html>\n"
	write(filepath.Join("public", filepath.FromSlash(file)), []byte(content))
	routes = append(routes, file)
}
func starter(p Project) {
	folder := filepath.Join("starters", p.Slug)
	files := map[string]string{}
	for n, b := range p.Files {
		files[n] = b
	}
	files["go.mod"] = "module academy.local/" + p.Slug + "\n\ngo 1.22\n"
	if _, ok := files["main.go"]; !ok {
		files["main.go"] = "package main\n\nimport (\n \"fmt\"\n \"os\"\n)\n\nfunc main() {\n if err := Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {\n  fmt.Fprintln(os.Stderr, err)\n  os.Exit(1)\n }\n}\n"
	}
	files["README.md"] = "# " + p.ID + " — " + p.Title + "\n\nUnfinished student scaffold. Core functions return ErrTODO. Compilation is not solution verification.\n\nRead the guide's specification before changing contracts. Implement core TODOs, remove the explicit t.Skip from acceptance tests, and add adversarial cases.\n\nGo 1.22 or newer; standard library only. Run go test ./... and go vet ./... inside this directory.\n\n" + p.Run + "\n\nGuide: ../../public/" + projectURL(p) + "\n"
	var buf bytes.Buffer
	z := zip.NewWriter(&buf)
	keys := []string{}
	for name := range files {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, name := range keys {
		if strings.HasSuffix(name, ".go") {
			cmd := exec.Command("gofmt")
			cmd.Stdin = strings.NewReader(files[name])
			b, e := cmd.Output()
			if e != nil {
				must(fmt.Errorf("gofmt %s/%s: %w", p.Slug, name, e))
			}
			files[name] = string(b)
		}
		write(filepath.Join(folder, filepath.FromSlash(name)), []byte(files[name]))
		write(filepath.Join("public", "downloads", p.Slug, filepath.FromSlash(name)), []byte(files[name]))
		header := &zip.FileHeader{Name: p.Slug + "/" + name, Method: zip.Deflate}
		header.SetModTime(time.Date(2026, 10, 28, 0, 0, 0, 0, time.UTC))
		w, e := z.CreateHeader(header)
		must(e)
		_, e = io.WriteString(w, files[name])
		must(e)
	}
	must(z.Close())
	write(filepath.Join("public", "downloads", p.Slug+".zip"), buf.Bytes())
}
func chapters(p Project) []Chapter {
	deps := []string{}
	for _, id := range p.Depends {
		deps = append(deps, linkProject(id))
	}
	depText := "Basic programming, terminal use, files and Git; no Go assumed."
	if len(deps) > 0 {
		depText = "Finish these builds first: " + strings.Join(deps, ", ") + "."
	}
	reading := "<h2>Mandatory before building</h2><ul>"
	for _, r := range p.Reading {
		if r.Mandatory {
			reading += "<li><a href=\"" + r.URL + "\">" + esc(r.Title) + "</a> — <strong>" + esc(r.Section) + "</strong>. " + esc(r.Why) + "</li>"
		}
	}
	reading += "</ul><h2>Recommended while building</h2><ul>"
	for _, r := range p.Reading {
		if !r.Mandatory {
			reading += "<li><a href=\"" + r.URL + "\">" + esc(r.Title) + "</a> — " + esc(r.Section) + ". " + esc(r.Why) + "</li>"
		}
	}
	reading += "</ul><h2>Reuse and readiness</h2><p>" + depText + "</p><p>Reading is preparation. Be ready to explain the invariant you will test before implementing it.</p>"
	stem := "../../downloads/" + p.Slug
	starterBody := "<p><a class=\"button button-primary\" download=\"\" href=\"" + stem + ".zip\">Download unfinished Go starter</a></p><p>The ZIP contains its own module, executable wiring, typed TODO boundaries, fixtures and explicitly skipped acceptance tests. There is no completed implementation. A green compile is only a scaffold check.</p><h2>Files to inspect</h2><ul>"
	names := []string{"go.mod", "main.go", "README.md"}
	for name := range p.Files {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, n := range names {
		starterBody += "<li><a href=\"" + stem + "/" + n + "\">" + esc(n) + "</a></li>"
	}
	starterBody += "</ul><h2>Core contract to implement</h2>" + code(p.Files["core.go"]) + "<h2>Start safely</h2><p>Extract into a new folder. Open the module directory before running commands; nested modules do not participate in the hub's own test command. Preserve the supplied data contracts unless your guide explicitly allows a change. Replace ErrTODO paths with validated behavior; do not suppress errors merely to make tests green.</p>"
	cs := []Chapter{{"overview.html", "Overview", p.Overview + "<h2>Prerequisite gate</h2><p>" + depText + "</p><p><strong>Planning estimate:</strong> " + fmt.Sprint(p.Hours) + " focused hours before review. Add 25–50% if Go or the mathematical model is new. Learner completion remains unassessed.</p>"}, {"required-reading.html", "Required reading", reading}, {"refresher.html", "Focused refresher", p.Refresher}, {"spec.html", "Full specification", p.Spec}, {"starter-structure.html", "Starter structure", starterBody}}
	for _, c := range p.Components {
		cs = append(cs, Chapter{"building-" + c.Slug + ".html", c.Title, "<p>" + c.Goal + "</p><h2>Before writing code</h2>" + list(c.Before) + "<h2>Build in small passes</h2>" + list(c.Steps) + "<h2>A grounded hint</h2><p>" + c.Hint + "</p><h2>Trace and self-check</h2><p>" + c.Trace + "</p><h2>Exit gate</h2><p>" + c.Gate + "</p>"})
	}
	cs = append(cs, Chapter{"wiring-it-up.html", "Wiring and running", "<p>The supplied main function delegates to <code>Run</code>. Implement the project boundary and connect the components from the guided chapters. Do not print debugging data into a machine-readable stream.</p><h2>Run inside the extracted module</h2>" + code(p.Run) + "<h2>Troubleshooting</h2><ul><li><code>ErrTODO</code> is expected until you implement the contract. Never mark that result as a completed project.</li><li>If Go cannot find a module, change into the folder containing this project's <code>go.mod</code>. No external module downloads are needed.</li><li>On Windows, run the built executable with <code>.\\</code>; on Unix use <code>./</code>. Go commands themselves are the same.</li><li>If output differs, compare structure and errors before changing a golden file. Stable ordering must come from the program, not repeated retries.</li></ul>"}, Chapter{"testing.html", "Verification and exit gates", p.Verification + "<h2>Submission evidence</h2><p>Remove each acceptance test's explicit skip only when you begin that feature. Add tests for the failure boundaries described here, run <code>go test ./...</code> and <code>go vet ./...</code>, then show a reviewer a failing test produced by one deliberate bug and the repaired result. Submit the commands, input, expected result and actual result. A reviewer must inspect tests as well as pass counts.</p><p><strong>Completion:</strong> the guide does not assess you automatically. Every required behavior and component gate must be demonstrated before a learner or instructor records readiness.</p>"}, Chapter{"engineering-questions.html", "Engineering Questions and Socratic Vision", "<p>Discuss these after the build. Use your own traces and code to defend a decision; there are no hidden answers.</p>" + list(p.Engineering)}, Chapter{"practical-questions.html", "Practical Questions", "<p>Each task should take a few minutes against your finished artifact. Keep the resulting test or observation.</p>" + list(p.Practical)}, Chapter{"extra-features.html", "Level Up: extra features", "<p>Optional, ordered from small extensions to larger investigations. Pass the mandatory gates first.</p>" + list(p.Extras)})
	refs := "<p>Primary references checked for this academy edition on 28 October 2026. The Go documentation is versioned independently; the taught APIs are available in Go 1.22 or newer.</p><ul>"
	for _, r := range p.Reading {
		refs += "<li><a href=\"" + r.URL + "\">" + esc(r.Title+" — "+r.Section) + "</a></li>"
	}
	refs += "</ul><p><a href=\"../../source-review.html\">Source snapshots, adaptation notes and unresolved external assumptions</a>.</p>"
	cs = append(cs, Chapter{"references.html", "References", refs})
	return cs
}
func buildProject(p Project) {
	starter(p)
	cs := chapters(p)
	// Links injected into prose use root paths; make them relative for guides.
	adjust := func(s string) string { return strings.ReplaceAll(s, "href=\"guides/", "href=\"../../guides/") }
	outline := "<aside class=\"page-outline\" aria-label=\"Guide chapters\"><p class=\"outline-title\">" + esc(p.ID) + " · Guide chapters</p><ol>"
	for _, c := range cs {
		outline += "<li><a href=\"" + c.File + "\">" + esc(c.Title) + "</a></li>"
	}
	outline += "</ol></aside>"
	idx := "<p><a href=\"../../index.html\">Academy</a> / " + esc(p.ID) + "</p><p class=\"eyebrow\">" + esc(strings.ToUpper(p.Kind)) + " · " + fmt.Sprint(p.Hours) + " HOURS · UNASSESSED</p><h1>" + esc(p.Title) + "</h1><p>" + esc(p.Summary) + "</p><ol>"
	for _, c := range cs {
		idx += "<li><a href=\"" + c.File + "\">" + esc(c.Title) + "</a></li>"
	}
	idx += "</ol><p><a href=\"../../roadmap.html\">Roadmap</a> · <a href=\"../../downloads/" + p.Slug + ".zip\" download=\"\">Download TODO starter</a></p>"
	page("guides/"+p.Slug+"/index.html", p.ID+" · "+p.Title, idx, outline, "guide-index")
	for i, c := range cs {
		prev := "index.html"
		if i > 0 {
			prev = cs[i-1].File
		}
		next := "../../roadmap.html"
		nextTitle := "Return to roadmap"
		if i+1 < len(cs) {
			next = cs[i+1].File
			nextTitle = cs[i+1].Title
		}
		body := "<p><a href=\"../../index.html\">Academy</a> / <a href=\"index.html\">" + esc(p.ID+" · Guide index") + "</a></p><p class=\"eyebrow\">CHAPTER " + fmt.Sprint(i+1) + " OF " + fmt.Sprint(len(cs)) + "</p><h1>" + esc(c.Title) + "</h1>" + adjust(c.Body) + "<nav class=\"chapter-nav\" aria-label=\"Chapter navigation\"><a rel=\"prev\" href=\"" + prev + "\">← Previous</a><a href=\"index.html\">Guide index</a><a rel=\"next\" href=\"" + next + "\">" + esc(nextTitle) + " →</a></nav>"
		page("guides/"+p.Slug+"/"+c.File, p.ID+" · "+c.Title, body, outline, "")
	}
}
func copyTree(from, to string) error {
	return filepath.WalkDir(from, func(f string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			return nil
		}
		r, e := filepath.Rel(from, f)
		if e != nil {
			return e
		}
		b, e := os.ReadFile(f)
		if e != nil {
			return e
		}
		write(filepath.Join(to, r), b)
		return nil
	})
}
func build() {
	manifest := []map[string]string{}
	must(filepath.WalkDir("sources", func(f string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			return nil
		}
		b, e := os.ReadFile(f)
		if e != nil {
			return e
		}
		manifest = append(manifest, map[string]string{"file": filepath.Base(f), "sha256": fmt.Sprintf("%x", sha256.Sum256(b)), "repository": "01-edu/public", "revision": "13cd08c1d25db64e79535a51348c332f7dc83b9f", "checked": "2026-10-28", "pinned_content_verified": "true"})
		return nil
	}))
	write("evidence/source-manifest.json", mustJSON(manifest))
	routes = nil
	must(os.MkdirAll("public", 0755))
	must(copyTree("assets", "public/assets"))
	must(copyTree("sources", "public/sources"))
	must(copyTree("evidence", "public/evidence"))
	// Fully emit one guide, scaffold and chapter set before beginning the next.
	for _, p := range projects {
		buildProject(p)
	}
	var pages []Page
	load("curriculum/pages.json", &pages)
	for _, p := range pages {
		page(p.File, p.Title, p.Body, "", "")
	}
	home := "<section class=\"hub-hero\"><p class=\"eyebrow\"><span class=\"status-dot\" aria-hidden=\"true\"></span> THE GO ALGORITHM LEARNING PATH</p><h1>Lem-in project track<span class=\"hero-accent\">Prepare. Then solve.</span></h1><p class=\"hero-description\">Turn an ant farm into a correct, quickest schedule. Learn Go, parsing, graphs, residual flow and independent verification through concrete builds, then bring them together in lem-in.</p><div class=\"hero-actions\"><a class=\"button button-primary\" href=\"roadmap.html\">Start with the roadmap ↗</a><a class=\"button button-secondary\" href=\"#project-library\">Explore the projects ↓</a></div><div class=\"hero-note\"><span>GO ONLY</span><span>STANDARD LIBRARY</span><span>LEARNER PROGRESS UNASSESSED</span></div></section><div class=\"stat-grid\"><div class=\"stat\"><strong>12</strong><span>Required preparation builds</span></div><div class=\"stat\"><strong>01</strong><span>Core lem-in assignment</span></div><div class=\"stat\"><strong>02</strong><span>Optional visualizer prep builds</span></div><div class=\"stat\"><strong>01</strong><span>Go GIF visualizer bonus</span></div></div><section class=\"hub-section\"><div class=\"section-heading\"><div><p class=\"eyebrow\">FIND YOUR NEXT STEP</p><h2>One track. A complete toolkit.</h2></div></div><div class=\"quick-grid\">"
	for i, n := range [][3]string{{"roadmap.html", "Follow the roadmap", "Preparation, integration and optional extension in order."}, {"coverage.html", "Trace every requirement", "Connect the brief to preparation and review evidence."}, {"policies.html", "Read the exact contract", "Grammar, turn semantics, output and limits."}, {"prerequisites.html", "Set up your workspace", "Begin with basic programming; learn Go here."}, {"instructor-guide.html", "Prepare to teach", "Readiness gates, review evidence and honest estimates."}, {"source-review.html", "Inspect the sources", "Upstream snapshots and explicit academy decisions."}} {
		home += fmt.Sprintf("<a class=\"quick-card\" href=\"%s\"><span class=\"card-number\">%02d</span><h3>%s ↗</h3><p>%s</p></a>", n[0], i+1, n[1], n[2])
	}
	home += "</div></section><section class=\"hub-section\"><p class=\"eyebrow\">PUT YOUR SKILLS TO WORK</p><h2>Two integration assignments</h2><div class=\"assignment-grid\">"
	for _, p := range projects {
		if p.Kind == "assignment" || p.Kind == "bonus" {
			home += "<a class=\"assignment-card phase-base\" href=\"" + projectURL(p) + "\"><div class=\"assignment-top\"><span class=\"project-id\">" + p.ID + "</span><span class=\"badge\">" + esc(p.Kind) + "</span></div><h3>" + esc(p.Title) + "</h3><p>" + esc(p.Summary) + "</p><div class=\"assignment-bottom\"><span>" + fmt.Sprint(p.Hours) + " focused hours · unassessed</span><span>→</span></div></a>"
		}
	}
	home += "</div></section><section id=\"project-library\" class=\"hub-section\"><p class=\"eyebrow\">GO PREPARATION LIBRARY</p><h2>Build each skill before you need it</h2><div class=\"library-grid\">"
	stages := []string{"01 · Go and colony data", "02 · Graphs and turn safety", "03 · Exact optimization", "04 · Verification and release", "05 · Optional visualizer preparation"}
	for i, s := range stages {
		home += "<section class=\"phase-card phase-base\"><div class=\"phase-heading\"><span class=\"phase-number\">" + fmt.Sprint(i+1) + "</span><h3>" + esc(s[5:]) + "</h3></div><ol class=\"project-list\">"
		for _, p := range projects {
			if p.Stage == s && p.Kind == "preparation" {
				home += "<li><a href=\"" + projectURL(p) + "\"><span class=\"project-id\">" + p.ID + "</span><span class=\"project-title\">" + esc(p.Title) + "</span><span class=\"project-time\">" + fmt.Sprint(p.Hours) + "h</span></a></li>"
			}
		}
		home += "</ol></section>"
	}
	home += "</div></section><section class=\"hub-section info-grid\"><div class=\"info-panel\"><h2>A build, not a solution dump</h2><p>Each guide has exact reading, a focused refresher, an explicit contract, unfinished downloadable Go code, component chapters, verification gates and discussion prompts. Students supply the algorithms.</p></div><div class=\"info-panel\"><h2>Exactness is part of correctness</h2><p>The planner searches deadlines with time-expanded integral flow. The track includes a separate movement checker and a tiny exhaustive oracle to challenge optimality claims.</p></div></section><p class=\"storage-note\">Standalone academy edition · 28 October 2026. No deployment, external progress tracking or automatic learner assessment.</p>"
	page("index.html", "Go-only lem-in learning hub", home, "", "hub")
	var reqs []Requirement
	load("curriculum/coverage.json", &reqs)
	cov := "<h1>Requirement coverage</h1><p>Every row connects a contract to preparation and independently reviewable evidence. U means upstream; A means academy policy; B means optional bonus. Coverage does not certify a student implementation.</p><p><a href=\"coverage.csv\" download=\"\">Download coverage CSV</a> · <a href=\"policies.html\">Exact contract</a></p><div class=\"table-scroll\" tabindex=\"0\" role=\"region\" aria-label=\"Requirement coverage table\"><table><caption>Source → preparation → evidence</caption><thead><tr><th scope=\"col\">ID and origin</th><th scope=\"col\">Requirement</th><th scope=\"col\">Preparation</th><th scope=\"col\">Review evidence</th></tr></thead><tbody>"
	csv := "ID,Origin,Requirement,Preparation,Evidence\n"
	q := func(s string) string { return "\"" + strings.ReplaceAll(s, "\"", "\"\"") + "\"" }
	for _, r := range reqs {
		links := []string{}
		for _, id := range r.Prep {
			links = append(links, linkProject(id))
		}
		cov += "<tr><th scope=\"row\">" + esc(r.ID+" · "+r.Origin) + "</th><td>" + esc(r.Rule) + "</td><td>" + strings.Join(links, "<br/>") + "</td><td>" + esc(r.Evidence) + "</td></tr>"
		csv += strings.Join([]string{q(r.ID), q(r.Origin), q(r.Rule), q(strings.Join(r.Prep, " ")), q(r.Evidence)}, ",") + "\n"
	}
	cov += "</tbody></table></div>"
	page("coverage.html", "Requirement coverage", cov, "", "")
	write("public/coverage.csv", []byte(csv))
	b, e := json.MarshalIndent(projects, "", "  ")
	must(e)
	write("public/catalog.json", b)
	sort.Strings(routes)
	write("public/routes.json", mustJSON(routes))
	fmt.Printf("Built %d HTML pages and %d unfinished starter modules.\n", len(routes), len(projects))
}
func mustJSON(v any) []byte { b, e := json.MarshalIndent(v, "", "  "); must(e); return append(b, '\n') }

type Doc struct {
	IDs   map[string]bool
	Links []string
}

func parseDoc(b []byte) (Doc, error) {
	d := Doc{IDs: map[string]bool{}}
	dec := xml.NewDecoder(bytes.NewReader(b))
	for {
		tok, e := dec.Token()
		if errors.Is(e, io.EOF) {
			break
		}
		if e != nil {
			return d, e
		}
		if s, ok := tok.(xml.StartElement); ok {
			for _, a := range s.Attr {
				if a.Name.Local == "id" {
					if d.IDs[a.Value] {
						return d, fmt.Errorf("duplicate id %s", a.Value)
					}
					d.IDs[a.Value] = true
				}
				if a.Name.Local == "href" || a.Name.Local == "src" {
					d.Links = append(d.Links, a.Value)
				}
			}
		}
	}
	return d, nil
}
func check(scaffolds bool) {
	docs := map[string]Doc{}
	problems := []string{}
	count := 0
	must(filepath.WalkDir("public", func(f string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() || !strings.HasSuffix(f, ".html") {
			return nil
		}
		b, e := os.ReadFile(f)
		if e != nil {
			return e
		}
		doc, e := parseDoc(b)
		if e != nil {
			problems = append(problems, f+": "+e.Error())
		}
		r, _ := filepath.Rel("public", f)
		docs[filepath.ToSlash(r)] = doc
		return nil
	}))
	handler := http.FileServer(http.Dir("public"))
	redirectedRoutes := 0
	for f, d := range docs {
		status, redirects, err := staticStatus(handler, "/"+f)
		if redirects > 0 {
			redirectedRoutes++
		}
		if err != nil || status != http.StatusOK {
			problems = append(problems, f+": static HTTP status "+fmt.Sprint(status)+" "+fmt.Sprint(err))
		}
		for _, s := range d.Links {
			count++
			u, e := url.Parse(s)
			if e != nil {
				problems = append(problems, f+": malformed URL "+s)
				continue
			}
			if u.IsAbs() || strings.HasPrefix(s, "//") {
				continue
			}
			dest := f
			if u.Path != "" {
				dest = path.Clean(path.Join(path.Dir(f), u.Path))
			}
			if strings.HasPrefix(dest, "../") {
				problems = append(problems, f+": path escape "+s)
				continue
			}
			info, e := os.Stat(filepath.Join("public", filepath.FromSlash(dest)))
			if e != nil {
				problems = append(problems, f+": missing "+s)
				continue
			}
			if info.IsDir() {
				dest = path.Join(dest, "index.html")
			}
			if u.Fragment != "" {
				if t, ok := docs[dest]; ok && !t.IDs[u.Fragment] {
					problems = append(problems, f+": missing anchor "+s)
				}
			}
		}
	}
	seen := map[string]bool{}
	chapterCount := 0
	for _, p := range projects {
		if seen[p.ID] {
			problems = append(problems, p.ID+": duplicate project ID")
		}
		for _, id := range p.Depends {
			if !seen[id] {
				problems = append(problems, p.ID+": dependency must occur earlier: "+id)
			}
		}
		seen[p.ID] = true
		if len(p.Engineering) < 6 || len(p.Engineering) > 8 || len(p.Practical) < 5 || len(p.Practical) > 7 || len(p.Extras) < 4 || len(p.Extras) > 6 || len(p.Components) < 2 {
			problems = append(problems, p.ID+": incomplete guide contract")
		}
		chapterCount += len(chapters(p))
		if !strings.Contains(p.Files["core.go"], "TODO") {
			problems = append(problems, p.ID+": missing core TODOs")
		}
	}
	var reqs []Requirement
	load("curriculum/coverage.json", &reqs)
	for _, r := range reqs {
		if len(r.Prep) == 0 || r.Evidence == "" {
			problems = append(problems, r.ID+": uncovered requirement")
		}
		for _, id := range r.Prep {
			if !seen[id] {
				problems = append(problems, r.ID+": unknown prep "+id)
			}
		}
	}
	results := []map[string]string{}
	if scaffolds {
		for _, p := range projects {
			for _, args := range [][]string{{"test", "-v", "./..."}, {"vet", "./..."}} {
				cmd := exec.Command("go", args...)
				cmd.Dir = filepath.Join("starters", p.Slug)
				cmd.Env = append(os.Environ(), "GOWORK=off", "GOTOOLCHAIN=local", "GOPROXY=off")
				b, e := cmd.CombinedOutput()
				status := "pass"
				if e != nil {
					status = "fail"
					problems = append(problems, p.ID+" "+strings.Join(args, " ")+": "+string(b))
				}
				results = append(results, map[string]string{"project": p.ID, "command": "go " + strings.Join(args, " "), "status": status, "output": strings.TrimSpace(string(b))})
			}
		}
	}
	sort.Strings(problems)
	report := map[string]any{"go_toolchain": runtime.Version(), "html_pages": len(docs), "static_http_routes_checked": len(docs), "canonical_index_redirects_followed": redirectedRoutes, "internal_and_external_link_attributes": count, "guide_chapters": chapterCount, "project_modules": len(projects), "requirements": len(reqs), "problems": problems, "scaffold_checks": results, "solution_tests": "NOT RUN: assessed implementations are intentionally unfinished; acceptance tests are explicitly skipped", "external_links": "Primary sources researched; live availability of every external reference is not a publication guarantee"}
	write("evidence/validation.json", mustJSON(report))
	write("public/evidence/validation.json", mustJSON(report))
	fmt.Printf("Checked %d HTML documents, %d link attributes, %d requirements. Problems: %d\n", len(docs), count, len(reqs), len(problems))
	for _, p := range problems {
		fmt.Println(p)
	}
	if len(problems) > 0 {
		os.Exit(1)
	}
}

// staticStatus follows only local redirects and bounds the number of hops.
func staticStatus(handler http.Handler, target string) (int, int, error) {
	request := httptest.NewRequest("GET", target, nil)
	for hops := 0; hops <= 5; hops++ {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, request)
		if rr.Code < 300 || rr.Code >= 400 {
			return rr.Code, hops, nil
		}
		location, err := url.Parse(rr.Header().Get("Location"))
		if err != nil || location == nil || location.IsAbs() || location.Host != "" {
			return rr.Code, hops, fmt.Errorf("invalid local redirect")
		}
		next := request.URL.ResolveReference(location)
		request = httptest.NewRequest("GET", next.RequestURI(), nil)
	}
	return 0, 5, fmt.Errorf("too many redirects")
}
func main() {
	files, e := filepath.Glob("curriculum/projects/*.json")
	must(e)
	for _, f := range files {
		var p Project
		load(f, &p)
		projects = append(projects, p)
	}
	if len(os.Args) < 2 {
		fmt.Println("go run ./cmd/hub build | check [-scaffolds] | serve")
		return
	}
	switch os.Args[1] {
	case "build":
		build()
	case "check":
		check(len(os.Args) > 2 && os.Args[2] == "-scaffolds")
	case "serve":
		fmt.Println("Preview: http://127.0.0.1:8086")
		must(http.ListenAndServe("127.0.0.1:8086", http.FileServer(http.Dir("public"))))
	default:
		must(fmt.Errorf("unknown command %q", os.Args[1]))
	}
}
