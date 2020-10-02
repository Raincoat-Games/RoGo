@echo off
cd tests
go test -bench=.
cd ..