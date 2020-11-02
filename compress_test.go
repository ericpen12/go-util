package go_utils

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestCompressService(t *testing.T) {
	reader, err := Compress("go")
	if err != nil {
		t.Fatal(err)
	}

	by, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("test2.zip", by, 0700)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFilepath(t *testing.T) {
	tests := []struct {
		Dir      string
		Filename string
		Result   string
	}{
		{
			Dir:      "/a",
			Filename: "bb.txt",
			Result:   "/a/bb.txt",
		},
		{
			Dir:      "a",
			Filename: "bb.txt",
			Result:   "a/bb.txt",
		},
		{
			Dir:      "a/",
			Filename: "bb.txt",
			Result:   "a/bb.txt",
		},
		{
			Dir:      "aa",
			Filename: "bb",
			Result:   "aa/bb",
		},
	}

	for _, v := range tests {
		ret := filePath(v.Dir, v.Filename)
		if ret != v.Result {
			t.Fatalf("the input dir is %s, the file name is %s, the except is %s, but the actual is %s",
				v.Dir, v.Filename, v.Result, ret)
		}
		t.Log(ret)
	}
}

func filePath(dir, filename string) string {
	p := path.Join(dir, filename)
	return p
}
