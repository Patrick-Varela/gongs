package wrapper

import (
	"io"
	"os"
)

func CopyFile(src, dst string, overwrite bool) error {
	s, e := os.Open(src)
	if e != nil {
		return e
	}
	defer s.Close()
	d, e := os.Create(dst)
	if e != nil {
		return e
	}
	defer d.Close()
	_, e = io.Copy(d, s)
	if e != nil {
		return e
	}
	e = d.Sync()
	if e != nil {
		return e
	}
	return nil
}
