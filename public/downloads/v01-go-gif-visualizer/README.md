# V01 — Optional Ant-Farm GIF Visualizer

Unfinished student scaffold. Core functions return ErrTODO. Compilation is not solution verification.

Read the guide's specification before changing contracts. Implement core TODOs, remove the explicit t.Skip from acceptance tests, and add adversarial cases.

Go 1.22 or newer; standard library only. Run go test ./... and go vet ./... inside this directory.

go test ./...
go vet ./...
go build -o visualizer .
# Windows executable name: go build -o visualizer.exe .
# Unix, after building the completed solver:
# ./lem-in colony.txt | ./visualizer
# Windows PowerShell native pipeline:
# .\lem-in.exe colony.txt | .\visualizer.exe
# For byte-exact portable tests, connect processes with Go os/exec.
# POSIX shell fixture: go run . < testdata/transcript.txt
# PowerShell: use the Go os/exec test harness for exact input bytes.

Guide: ../../public/guides/v01-go-gif-visualizer/index.html
