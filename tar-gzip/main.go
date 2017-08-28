package main

func main() {
	if err := Compress("./dir", "./dir.tar.gz"); err != nil {
		panic(err)
	}

	if err := Decompress("./dir.tar.gz", "dir2"); err != nil {
		panic(err)
	}
}
