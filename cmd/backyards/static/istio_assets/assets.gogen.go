// Code generated by vfsgen; DO NOT EDIT.

package istio_assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
		},
		"/istio.yaml": &vfsgen۰CompressedFileInfo{
			name:             "istio.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 596,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xcd\x6e\xdb\x30\x10\x84\xef\x7a\x8a\x45\xee\x56\xaa\x22\x97\xf2\x16\x14\x01\x9a\x43\x01\xa1\x09\x7a\x5f\x91\x5b\x77\x6b\xfe\x61\xb9\x4c\x2c\x3f\x7d\x41\xd1\x76\xdd\x5e\x7a\xa3\x46\xdf\x0c\x39\x83\x99\xbf\x93\x14\x4e\xd1\x00\x17\xe5\x34\x2e\x18\x4f\xc8\xd6\xa7\xea\x46\x4e\xf7\x6f\xd3\x42\x8a\xd3\x70\xe0\xe8\x0c\x3c\x37\x64\x08\xa4\xe8\x50\xd1\x0c\x00\x11\x03\x19\x08\x54\x7e\x0e\x25\x93\x6d\xd2\xdb\x25\xf0\x6e\x1a\x1f\xc6\x8f\x77\x03\x40\x50\x5f\x0c\xa8\x54\x1a\x00\xb0\x6a\xfa\xba\x09\x3f\xd0\x97\xa6\x34\xfb\xd3\x31\x63\xec\xbe\x33\x67\x53\x54\x49\x7e\xf6\x18\xe9\x85\x6c\x15\xd6\xf5\x29\xe2\xe2\xc9\x5d\x99\xc2\x8e\x2c\xca\x73\xfc\x45\x56\x93\xb4\xeb\x01\x84\xde\x85\x95\x1e\x73\xfe\xf2\xfa\x3a\xcf\x92\x16\xba\x1a\x38\xe0\x9e\xe6\xea\xfd\x9c\x3c\xdb\xd5\xc0\xa3\x7f\xc7\xb5\x0c\x00\x7b\x54\x6a\xc7\x9e\xc1\x71\x2f\x54\xce\x1f\x00\x01\x8f\xdf\x28\x7b\xb6\x58\x0c\x4c\x9b\x48\x7f\x01\x74\x79\xd8\xa5\x53\x66\x9f\xf4\x9c\xd5\xee\x34\x70\xb3\xec\xfd\xb6\xf5\xae\x33\x6d\xa5\x0f\xbb\xe5\x64\xdb\x12\x7c\x24\xf9\x8f\xab\x33\xb7\x2e\x80\x50\xbd\xf2\x67\x5f\x8b\x92\xbc\xd4\x9c\x93\xe8\xb5\xb2\x0a\x5a\x8e\xfb\x9e\x4a\xff\x0c\xd8\x7f\x93\x18\x38\x71\x3e\x70\xdc\xa4\x7e\xbc\x54\x43\xe7\xb6\xa6\xb0\xa0\x3d\xac\x28\xae\xec\x3a\x30\xfe\x11\xca\x5a\x94\x82\xf9\xf4\x30\x4d\xc3\xef\x00\x00\x00\xff\xff\xf8\x47\x0f\x0d\x54\x02\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/istio.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
