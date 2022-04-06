#!/bin/sh

mkfifo myfifo
<myfifo go run Zad7a.go | go run Zad7b.go | tee myfifo
rm myfifo