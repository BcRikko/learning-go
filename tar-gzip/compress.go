package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func addFile(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	hdr := new(tar.Header)
	hdr.Name = filepath.Base(path)
	hdr.Size = stat.Size()
	hdr.Mode = int64(stat.Mode())
	hdr.ModTime = stat.ModTime()

	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}

	if _, err := io.Copy(tw, file); err != nil {
		return err
	}

	return nil
}

func Compress(in string, out string) error {
	var paths []string
	filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, path)
		return nil
	})

	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()

	gw := gzip.NewWriter(file)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	for _, path := range paths[1:] {
		if err := addFile(tw, path); err != nil {
			return err
		}
	}

	return nil
}
