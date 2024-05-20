#!/bin/zsh

go build -o bin/lw-scan-brute src/lw-scan-brute.go src/common.go
go build -o bin/lw-stage-1 src/lw-stage-1.go src/common.go
go build -o bin/lw-stage-2 src/lw-stage-2.go src/common.go
go build -o bin/c2-api src/c2-api.go
go build -o bin/c2-listener src/c2-listener.go

chmod +x bin/c2*