#!/bin/sh
awhile=3
swag init --parseDependency --parseInternal --parseDepth 3 -g cmd/main.go
sleep $awhile && open http://localhost:8085/swagger/index.html &
make
