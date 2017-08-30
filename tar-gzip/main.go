package main

import (
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := Compress(
		filepath.Join(dir, "in"),
		filepath.Join(dir, "golang.tar.gz"),
	); err != nil {
		panic(err)
	}

	if err := Decompress(
		filepath.Join(dir, "golang.tar.gz"),
		filepath.Join(dir, "out"),
	); err != nil {
		panic(err)
	}
}
