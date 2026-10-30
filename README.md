# Lem-in Academy — Go only

A standalone, static learning hub: 12 required preparation projects, the lem-in integration assignment, two optional visualizer preparation projects and a Go GIF visualizer bonus. All 16 guides include complete chapter sets and unfinished standard-library Go starters. Learner completion is unassessed.

Open `public/index.html` or serve the generated site. Begin with `public/roadmap.html`.

## Build and verify

Run at this directory using Go 1.22 or newer on a 64-bit target:

```
go run ./cmd/hub build
go run ./cmd/hub check -scaffolds
go test ./...
go vet ./...
go run ./cmd/hub serve
```

The preview URL is http://127.0.0.1:8086. No Node.js or external Go package is required. Student starters are separate nested Go modules; root tests do not run their tests, so use the scaffold check explicitly.

## Authoring and publication

- `curriculum/projects/*.json`: complete project content, reading, specifications, component steps, questions and starter files.
- `curriculum/pages.json` and `coverage.json`: academy pages and requirement mapping.
- `cmd/hub`: Go build, ZIP packaging, structural/link/dependency/coverage checker and preview server.
- `assets`: copied forum dark theme plus academy overrides.
- `sources`: preserved original brief, audit, examples and good practices at the checked upstream revision.
- `starters`: generated unfinished Go modules. Extract a ZIP into a separate working folder for student changes.
- `evidence`: actual publication/scaffold/model/browser evidence with explicit limitations.
- `public`: real publication output, including all relative assets and downloads.

The Vercel configuration publishes the already-built `public` directory. Its static output check does not assume a remote Go installation. Rebuild and verify locally before any future requested deployment. This delivery has not been pushed, published or deployed.

The algorithm curriculum teaches exact time-expanded integral flow with shared undirected tunnel capacities. A separate scratch occupancy-search experiment checked the model on 3294 small graph/ant cases and a targeted greedy trap; its result is evidence about the teaching model, not a finished student solver or a performance audit. Read `public/validation.html` for precise checks and limits.

The forum repository was a read-only design/format reference. No existing Learning Hub was changed.
