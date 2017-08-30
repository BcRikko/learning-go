package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Decompress(in string, out string) error {
	if _, err := os.Stat(in); err != nil {
		return err
	}

	file, err := os.Open(in)
	if err != nil {
		return err
	}
	defer file.Close()

	gr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}

		path := filepath.Join(out, hdr.Name)

		fmt.Printf("%T %d\n", tar.TypeReg, tar.TypeReg)   // int32 48
		fmt.Printf("%T %d\n", hdr.Typeflag, hdr.Typeflag) // uint8 0

		// TODO: なぜかtar.TypeRegがint32 48になってしまうので、48を足して回避
		switch hdr.Typeflag + 48 {
		case tar.TypeDir:
			if _, err := os.Stat(path); err == nil {
				if err := os.MkdirAll(path, 0755); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(file, tr); err != nil {
				return err
			}
		default:
			fmt.Println("unsupported")
		}
	}
}
