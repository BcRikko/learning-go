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
		filepath.Join(dir, "dir"),
		filepath.Join(dir, "dir.tar.gz"),
	); err != nil {
		panic(err)
	}

	if err := Decompress(
		filepath.Join(dir, "dir2.tar.gz"),
		filepath.Join(dir, "dir2"),
	); err != nil {
		panic(err)
	}
}
