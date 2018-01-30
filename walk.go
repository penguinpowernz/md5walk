package md5walk

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Sums is a collection of MD5 sums indexed by the filename
type Sums map[string]string

func (s Sums) Write(w io.Writer) (int, error) {
	data := []byte{}
	keys := []string{}
	for k := range s {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, name := range keys {
		data = append(data, []byte(fmt.Sprintf("%s  %s\n", s[name], name))...)
	}

	return w.Write(data)
}

// Walk will recurse through the given directory, generating MD5 sums
// for any files found
func Walk(dir string) (Sums, error) {
	sums := map[string]string{}

	err := filepath.Walk(dir, md5WalkFunc(func(path, hash string) {
		path = strings.Replace(path, dir+"/", "", -1)
		sums[path] = hash
	}))

	return sums, err
}

func md5WalkFunc(cb func(string, string)) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// don't sum directories
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		h := md5.New()
		_, err = io.Copy(h, f)
		if err != nil {
			return err
		}

		cb(path, fmt.Sprintf("%x", h.Sum(nil)))

		return nil
	}
}
