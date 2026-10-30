# P11 — Adversarial Validator and Tiny Oracle

Unfinished student scaffold. Core functions return ErrTODO. Compilation is not solution verification.

Read the guide's specification before changing contracts. Implement core TODOs, remove the explicit t.Skip from acceptance tests, and add adversarial cases.

Go 1.22 or newer; standard library only. Run go test ./... and go vet ./... inside this directory.

go test ./...
go vet ./...
go run . testdata/oracle.json
go test -run TestDifferential -count=1 ./...
go test -fuzz=FuzzTranscript -fuzztime=10s

Guide: ../../public/guides/p11-adversarial-verification/index.html
