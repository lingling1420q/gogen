// Code generated by go-bindata.
// sources:
// golang/pkg/model/.DS_Store
// golang/pkg/model/database/db.go.tpl
// golang/pkg/model/database/table.go.tpl
// DO NOT EDIT!

package gens

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _golangInternalModelDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xcd\xaa\xd3\x40\x14\x80\xcf\xe4\x46\x9d\x70\x0b\x8e\xe0\xa2\x82\x8b\x80\x0b\x37\x5d\xd4\xd6\xda\x6d\x88\xed\x4e\x50\x4c\xb1\x8b\xaa\x75\x86\x0c\x34\x12\x92\x92\x4c\xcd\xa2\x14\xea\x0b\xf8\x00\xea\xd6\x9f\xc7\x10\xf4\x21\x7c\x11\xf7\xd2\xce\x11\x43\x5b\x97\xd6\x7a\x39\x1f\x0c\x1f\xc9\xfc\x24\x33\x03\x93\x73\x02\x00\x2c\x5c\xc4\x77\x00\x04\x00\x70\xb0\x76\x3c\x38\x08\xc7\xb2\x87\x83\x66\xb6\x08\x00\x05\x12\x4a\xd0\xaa\x2a\xe7\x87\xc7\x22\x4e\x0c\xb6\xdd\xdc\x18\x24\x18\x90\xf5\xfd\x53\x69\xae\x00\xe0\x87\x9a\xa7\x49\x69\xda\xed\xef\xcc\x39\x73\x2f\x5d\xbe\xc2\x3d\xef\xbc\xe1\x35\xae\x36\x9e\x45\xb3\xbc\x8a\x8c\x34\x8b\x32\x94\xc5\x64\x73\xf5\x48\x9a\x99\x92\xc5\x54\xdc\x8c\x92\x58\x2b\x59\x8c\x93\xd8\xcc\x46\x3a\x1b\xa6\xfa\x95\xce\x1e\x16\x0f\xa4\xd1\xb6\xe9\x28\xcf\x53\x85\xdd\x46\x52\x3d\x49\x74\x35\x15\xd7\xef\xe7\x99\x91\x49\xa6\x8b\xed\xd8\x76\x8c\xa7\xe3\x24\x8b\xf3\x2a\xcc\x17\x59\x5c\x4e\x6a\x15\x9c\xdf\x0a\xe4\x97\xed\x2c\x3c\xee\x4d\x45\x73\xb9\xec\xf4\x7b\x2d\xbf\xd3\xeb\xad\x5a\xfe\xb2\xdf\x6f\xb7\xfc\xbb\xdd\x7b\xab\x95\xc7\x6f\xdc\xee\x3e\x7e\xfe\xf2\xf5\x9b\xb7\xef\xde\x7f\xf8\xf8\xe9\xf3\x57\x9c\x3b\xc3\x45\xb8\xb6\xb3\x28\xdf\x8e\xbb\x07\x04\x41\x10\xc7\x06\x8f\x3f\x7e\xfe\xaf\x5f\x84\x20\x88\x93\x63\x73\x3e\xf8\xe8\x00\xbd\xb6\x66\x58\xef\xa0\xdd\x5a\x1f\x81\xf6\xd1\x01\x7a\x6d\xcd\xb0\x9d\x83\x76\xd1\x1c\x2d\xd0\x3e\x3a\x40\xaf\xad\xf1\xd0\x62\x98\x7c\x30\x7c\x32\xc3\x0c\x85\x09\xb4\x8f\x0e\xfe\xce\xda\x10\xc4\xff\xce\x99\x95\xd8\x7c\xff\x87\x7f\xce\xff\x09\x82\xb8\xc0\x30\x77\x10\x0d\xc2\xdf\x09\xc1\x1e\x0e\x06\x02\x2f\x7e\x75\xd8\x09\x04\xa0\x16\x04\x38\xf6\x67\x61\xb3\x76\x9f\x02\x01\x82\x38\x31\x7e\x06\x00\x00\xff\xff\xe5\x46\xbb\xcd\x04\x18\x00\x00")

func golangInternalModelDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_golangInternalModelDs_store,
		"golang/pkg/model/.DS_Store",
	)
}

func golangInternalModelDs_store() (*asset, error) {
	bytes, err := golangInternalModelDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang/pkg/model/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1586257494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangInternalModelDatabaseDbGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xc1\x6e\xd4\x30\x10\x3d\xc7\x5f\x31\x8d\xb4\x55\x82\x4a\xc2\x11\x55\x8a\x40\xdb\x54\xa5\xd2\xb6\x5b\x69\xcb\x19\x39\xf1\x24\x35\x8d\xed\x60\x3b\x85\x12\xe5\xc0\x0f\x94\x1b\x47\x24\x0e\x9c\x39\x72\xec\xcf\x40\xd5\xcf\x40\xde\x4d\xd3\xcd\x6a\x85\x7a\xf4\xcc\x9b\x37\xef\x3d\xdb\x71\x0c\x07\x8a\x21\x94\x28\x51\x53\x8b\x0c\xb2\x6b\xb0\x56\xa9\x2a\x22\x71\x0c\xe9\x1c\x4e\xe7\xe7\x70\x98\x1e\x9f\xef\x10\x52\xd3\xfc\x92\x96\x08\x6d\xdb\x48\x86\xba\xe2\x12\x21\x3a\xbb\x2c\xbb\x8e\x10\x2e\x6a\xa5\x2d\x04\xc4\xf3\x0b\x61\x7d\x42\x3c\xbf\xe4\x36\x52\x34\xca\x95\x88\x4b\xf5\xdc\xa0\xbe\xe2\x39\x9a\x58\x98\xc2\xff\x6f\x37\x66\x59\x0f\xb8\x68\xb2\x25\xe0\x3d\x97\x9f\x2f\x9a\xb8\x54\x5a\xf8\x24\x24\xe4\x8a\x6a\xb7\x89\x65\xa7\x54\x20\x24\xe0\xb7\x6d\x94\x52\x4b\x33\x6a\xb0\xeb\xdc\x72\x96\x1d\x28\x59\xf0\x12\x9e\xb1\x2c\x9a\xd7\x96\x2b\x69\x5c\xf5\x58\x16\x0a\x60\x59\x4d\xa7\xee\xe0\xe8\xe2\xf8\x08\x6d\x3a\x85\xfb\x9b\xdf\x7f\xbf\x7e\x63\xd9\xfd\xed\xf7\xbb\x9b\x9f\x7f\x6e\x7f\xdc\x7d\xf9\x45\x8a\x46\xe6\xb0\xec\x07\xe1\xda\x1c\xb4\xc4\xe3\x05\xf4\x8c\x3b\x09\x48\x5e\xb9\x9a\xa7\xd1\x36\x5a\xf6\x0d\xe2\x75\x3d\xac\x97\x93\x3c\x02\xdf\x3d\x14\xf7\x00\xb5\x86\xfd\x04\x84\x29\xa2\x99\xa2\x2c\x9d\xae\x1a\xc1\xca\x60\x48\x3c\xc7\xe1\x40\x6b\x7b\xbc\x9a\x4a\x9e\x07\xa8\xb5\xeb\x77\xc4\x7b\xf4\x9c\xc0\xc0\xed\x04\x8c\x14\x2c\x73\x5a\xa8\x46\xe7\xe8\xc4\xf8\xfe\x92\x6c\x6b\x17\x0a\x61\xa3\x45\xad\xb9\xb4\x45\xe0\x4f\xcc\xfe\xc4\xbc\xb6\x79\x1d\x4c\x4c\x18\x4f\xcc\xab\x9a\x6a\x83\xe7\x5c\x60\x62\x75\x83\xbb\xf9\x85\x3b\xdb\xa4\xb1\xc5\x4b\x7f\xcf\x09\x1c\x48\xdf\x1a\xd4\x92\x0a\x1c\x57\xcf\xa8\x31\x1f\x95\x66\xe3\xea\x1b\x65\xec\xb8\x92\x4e\xfb\x10\x9c\x13\xd6\xdf\xf2\x90\x99\x7b\x13\xd1\xbc\x46\x19\xf8\xe2\xda\x7c\xa8\xfc\xbd\x6d\x56\x43\xb2\x25\xc1\xf5\x00\xbb\x47\xea\x68\xa6\xca\x13\xc5\x30\x70\xb6\xc2\x71\x78\x27\xf4\xd3\x31\xab\xf0\x40\x49\x69\x1c\xd5\x8b\x55\x7a\x0f\x93\xee\x8d\x44\x0b\xb4\xeb\xb0\x60\xeb\x70\xb8\xf9\x32\x5c\xd7\xd9\x78\x02\xf5\x00\x0b\xb6\x0e\xf7\x49\xc5\xb1\xd5\xd7\xc4\x73\x9e\x13\x18\xd3\x9c\x71\x59\x06\x4f\x88\xc4\x91\xe4\x8d\xb1\x4a\x40\xa5\xca\x12\xf5\x58\xf1\xa1\xa4\x59\x85\x33\x55\x8e\x95\x6e\xc4\xd7\x01\x56\x06\xb7\x43\x0a\x5a\x19\xdc\x48\x7f\x81\x76\xb6\x5c\x16\xec\xb2\x2c\x3a\x52\x5a\xac\x8e\x6d\x17\x92\xe1\x07\x27\xb0\x3b\xfc\xc4\x36\x9d\xee\x0f\x06\x3b\xb2\xf1\x01\x3b\x42\xfe\x05\x00\x00\xff\xff\x91\x52\x39\x71\xe4\x04\x00\x00")

func golangInternalModelDatabaseDbGoTplBytes() ([]byte, error) {
	return bindataRead(
		_golangInternalModelDatabaseDbGoTpl,
		"golang/pkg/model/database/db.go.tpl",
	)
}

func golangInternalModelDatabaseDbGoTpl() (*asset, error) {
	bytes, err := golangInternalModelDatabaseDbGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang/pkg/model/database/db.go.tpl", size: 1252, mode: os.FileMode(420), modTime: time.Unix(1586257494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangInternalModelDatabaseTableGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5b\x6f\xdb\x46\x16\x7e\x16\x7f\xc5\x31\xe1\x18\x54\xe0\x25\x37\xc0\x62\x1f\x1c\x08\x8b\x58\x92\xbd\xc6\x3a\x72\xd6\x92\x91\x87\xc5\xc2\xe6\x65\x28\xd3\xa6\x86\xf2\x70\x28\x47\x21\x08\x64\x17\x08\x36\x46\xec\x24\x58\x34\x46\xd3\x26\x41\x0c\x34\x69\xda\xa2\xaa\x1f\x8a\x36\xcd\xf5\xcf\x98\x8c\xf3\xd4\xbf\x50\x0c\x87\x92\x29\xdf\x44\x21\xbd\xf8\xc9\x21\xe7\x9c\x6f\xbe\xef\x9c\x8f\x67\x46\x51\x14\x28\x3a\x06\x82\x3a\xc2\x88\xa8\x14\x19\xa0\xb5\x81\x52\xc7\xb1\x65\x41\x51\xa0\x34\x07\x95\xb9\x1a\x94\x4b\x33\xb5\x11\x41\x68\xaa\xfa\xaa\x5a\x47\xe0\xfb\x1e\x36\x10\xb1\x2d\x8c\x40\xbe\xb2\x5a\x0f\x02\x41\xb0\x1a\x4d\x87\x50\x90\x04\x00\x00\xd1\x6c\x50\x51\xc8\x89\xd4\x6a\x20\x51\xc8\xf9\xbe\x65\x02\x5a\x03\x79\xc6\xad\xb6\x30\x98\xaa\xed\xa2\x20\x10\x72\x62\xdd\xa2\xb2\xa3\xca\xba\xd3\x50\xea\xce\x9f\x5c\x44\x5a\x96\x8e\x5c\xa5\xe1\x9a\x8a\xa1\xc5\x89\xa8\x1b\x4a\x1d\xef\xfa\xb2\xb5\x62\xf1\x68\xda\x17\x83\x8d\x2e\xda\xb2\xa7\xc5\xeb\x2b\x16\xbe\xbe\xec\x29\x75\x87\x34\x44\x21\x2f\x08\x42\x4b\x25\xb0\x08\x05\x60\x8c\xe4\x8a\xb3\x2e\xe5\x05\x41\x51\x7c\xdf\x6b\x36\x11\x29\xaa\x0d\x64\x83\x7c\x85\x20\x4a\xdb\x35\x55\xb3\x51\x45\x6d\xa0\x20\x00\xca\xfe\x0d\x08\x53\x8b\xb6\x05\xda\x6e\xc6\xd2\x4f\xcd\x70\x29\xf1\x74\x0a\xbe\x00\x39\xdf\x07\xa2\xe2\x3a\x82\x51\x6b\x1c\x46\x5b\x30\x51\x00\xb9\xe8\xd8\x5e\x03\xbb\xc0\xf8\xf6\x41\x8d\xb6\xe4\x29\x0b\xd9\x46\x10\x80\xef\x8f\xb6\xe4\x69\xa7\xd6\x6e\x32\xc0\x25\xa6\x61\x42\xd4\xe3\xc4\x89\x78\x2d\x09\xbc\xd8\x2d\xeb\xa8\x7c\x85\x58\x0d\x95\xb4\xff\x81\xda\x29\xa0\x26\x7f\xb9\xb8\x8a\xda\x17\x93\x22\xc5\x19\xa3\x2d\xf9\x92\x47\x9d\x19\xac\x13\xd4\x40\x98\x42\x10\x5c\x5a\xa8\xcd\x2d\xce\x54\x8a\xf3\xe5\xcb\xe5\x4a\xad\x2f\x9a\xe1\xb7\xe4\x8a\x67\xdb\x20\x56\xe6\xc4\x20\xc0\x0e\x05\xec\xd9\xf6\x61\xcc\x05\x6c\xad\x79\x28\x08\xbc\xf8\x6f\x77\x95\x15\x8d\xb3\xe6\x7a\x44\x58\x71\x1d\x3c\x21\xa6\x85\x88\x4b\xd0\x6d\xa2\x90\x5b\x34\x34\x38\x6f\x68\x72\x69\x72\x06\x9b\x8e\x90\x5b\xa4\xd7\xe2\xe7\x1a\x51\xb1\xab\xea\xd4\x72\xb0\x10\xb0\xde\xf5\xaa\x0e\xfb\x3b\xcf\xc3\x7b\x5b\x82\xe9\x61\x1d\x24\x47\x5b\x81\xf3\x03\xba\x94\x87\xde\x83\x94\x67\x2d\xb3\x70\x1d\x7c\x21\x47\x10\xf5\x08\x06\xd1\xf7\xe5\x54\xb4\xc8\xf7\xab\xa0\xf5\x41\xcd\x0f\x3b\x8f\xf7\xde\xde\x0e\x37\xb7\x39\x97\xc1\x19\x52\x7e\x20\x57\xc6\x8b\x69\x9a\x28\x64\xc0\xbb\x6a\xd1\x65\xc7\xa3\x25\x4d\xca\xc7\x59\x72\x15\x51\x5e\x48\xf6\x22\x91\xe7\x68\x2b\x59\x15\xf5\xf0\xe0\x40\xdb\xde\x8b\xad\x70\xf3\x66\x78\xef\x1b\x43\xdb\x7b\xb7\x13\xfd\xe7\xbb\xac\x62\x53\xe4\x32\xa9\x56\x14\x5d\xc5\xe0\x22\x0a\x06\x32\x55\xcf\xa6\x60\x32\xbf\xf4\x74\x8c\x0d\xc0\xf0\x03\x2e\xb3\x57\x03\xd8\xef\xbc\x7d\xff\xa6\x63\x68\xfb\xef\x1e\x47\x77\x9e\xa6\xd9\x67\xb4\x4d\xaa\x9c\x8c\xa1\xa1\x52\x55\x53\x5d\xc4\x9a\x33\xcd\x96\xba\x65\x67\x26\x2e\xc0\x58\xcf\xc6\xbe\x90\xcb\x95\x26\x27\xa0\x9b\x20\x97\x26\xc7\x85\xdc\x01\xbf\x94\xb9\x13\x92\x7b\x2f\x6f\xbf\xdf\xf8\x6a\x58\x72\x29\x1c\xe9\xe8\x77\x93\x4f\xbc\x24\xb3\x6f\xaa\x00\xf4\x1a\xdf\x7f\x1a\xd1\xea\x3f\x67\xcb\xd7\x90\xee\x51\x87\xc0\xfe\x9d\x1f\xc3\xbb\xdb\xee\x9a\x1d\x6d\x7c\xb9\xbf\xb3\x19\x3e\x78\x3e\x14\x8b\x7e\x34\x29\x0f\x86\x26\xa7\xe1\x7d\x21\x67\x99\xd0\x65\x31\x52\x00\x6c\xd9\xec\x65\xca\x9d\x6c\x85\x55\xa7\xef\x95\xa1\x31\xb6\x42\x3c\x6d\xe2\x63\xe7\x60\xe4\x89\x62\x10\x08\x8a\x32\x83\x5d\x44\x28\x44\x77\xff\x1f\xde\x7c\xba\xf7\xe2\x46\xf4\x68\x27\xba\xbf\x1b\x6d\x75\x86\xe2\xcf\x51\x8e\xf6\x97\x91\x38\xac\x2d\x96\x82\x08\x61\xeb\xbd\xd6\x16\x09\x52\x29\x62\xbb\xe5\xe5\x32\x21\x0e\xb9\x18\x87\xa4\x94\x36\x55\x6c\xe9\x12\x22\x24\xcf\x3d\xd0\x1b\x7e\x8a\xb2\xd0\x34\x54\x8a\x20\x7a\xf2\x53\xb4\xd5\xd9\x7b\xf1\xea\xc3\x27\x9d\xe8\xf3\xef\xa3\xed\xdd\xa1\x34\x70\x94\x8f\xd0\x50\x55\x5b\x43\x29\x60\x2e\x56\x5b\x08\xc2\x67\xff\x8d\x1e\x3f\x6c\xae\x86\xdf\x7e\x1a\x3e\x7c\x1e\xde\x7a\xc0\xc9\xff\xfc\x7a\x33\xbc\xf7\x8c\x3d\xa6\x7b\xb3\xbd\xfb\xfe\xe5\xd7\xe1\x9b\xfb\xc3\x99\x9c\x31\xfb\x1d\x85\x9d\xe2\xb7\x12\xb2\x11\x45\x60\xf0\x3f\x5a\x1b\x9a\xab\x43\x49\xe1\xf9\x52\x1e\x2c\x4c\xff\xfa\x97\x6c\x92\x08\x72\xd9\x20\x4c\x4b\x4a\x60\x98\xa8\xb4\x64\x1e\x79\x82\xcc\x24\x6c\xa4\x00\xec\x76\xc1\x82\xe6\x91\xee\x10\xa3\xe2\xd0\x29\xc7\xc3\x46\x1c\xd4\x57\x0c\xf6\x39\x1e\x7c\x91\x09\xf8\xbc\xb3\xee\x5e\x32\x4d\xa4\x53\x64\xf4\xfb\x98\x93\x1a\x50\x82\xc9\x76\x62\xf4\xe8\xd1\xce\xde\xab\x1f\xc2\x5b\x4f\x3e\x3c\xf8\x82\x97\x30\x63\xbe\xb4\xbe\x8c\x08\x4a\x4e\xf0\x71\x50\x49\xdd\x05\x59\x96\x2d\x4c\x11\x31\x55\x1d\xf9\xc1\x09\xd5\xed\x4d\x6b\x77\xcd\x66\xcf\x62\xa9\x3c\x5b\xae\x95\x61\x6a\x7e\xee\x32\x2c\xf5\x9f\xff\x4b\x20\xc6\x85\xe5\x7b\x8d\x14\x40\x14\xe3\xfa\xb0\xd4\x02\x98\x0d\x2a\x57\x9b\xc4\xc2\xd4\x94\xc4\x73\x6e\x12\x75\xce\x15\xc7\xc1\x5d\xb3\xc7\xf9\x73\xec\xa6\x63\xbb\xc7\x9a\x2b\xc5\x81\x8c\xbc\x2c\xcb\x7f\x74\x13\x4f\x73\xfc\x34\xa2\x03\x5b\x72\x38\xa0\x8b\x11\x04\x7d\x53\x8d\x9f\x32\xe1\x8d\xd7\xbc\xdf\x1f\x87\x2c\xf9\x7e\xff\x4e\xdc\x10\x43\x5d\xae\x32\xdc\x26\x4e\x77\x50\xb5\x3c\x5b\x2e\xd6\xc0\xf7\xf9\xb5\xb6\x1a\x53\x98\xb5\x5c\x1a\x04\x27\xd8\xea\xea\xdf\xcb\xf3\xe5\xf8\x75\x9a\xfc\x52\xe1\x6f\x60\x5b\x0d\x8b\xc2\x05\xf1\xd8\x19\x36\xaf\xae\x73\xc7\x1c\x4a\xcc\xcb\x55\x5d\xc5\xa7\x4d\xb7\x04\xac\x70\xba\x63\x12\x7b\x60\xcb\xe6\x8e\x39\x34\x14\x07\xdc\x2b\x7b\x43\x20\x93\x5f\xa6\x2c\xe2\xd2\xbe\x29\xc0\x8d\xc1\x8f\x88\xfd\xce\x6e\xef\x7c\xc8\x88\x95\x65\x22\x9c\x6d\x57\x0c\x33\x6c\xb8\x83\x8e\x1b\x36\x27\x24\x74\x8d\x15\xc7\x1f\x7f\x44\xf6\xec\xd5\x1d\x48\x67\xc2\x56\x99\xdc\x74\x35\x2e\xda\x51\x37\x45\x1b\x37\xa2\x87\x1b\xc3\xb9\x29\xc6\xca\xe2\xa6\x7f\xfd\x3b\xa3\x9f\x5c\x56\xe6\xc1\xd1\x67\xd7\x52\xd9\xcd\x32\xc6\xe4\xfe\xb6\x76\x31\x1d\x02\x8b\xe3\x90\x7c\xa7\xfc\x3f\x5a\xe2\x22\xb3\xec\x23\x5e\xea\xfb\x31\xe1\x66\xf5\x53\xd1\xf1\x30\x4d\x7e\x10\x45\x9b\xff\x0b\x3b\x9f\x71\x0f\xed\xef\x6c\x46\xf7\x77\xb3\x3a\x29\x46\xf9\xf8\x9b\x4a\x4b\x25\xa0\xc7\x84\xe2\xb8\x23\x7e\x28\xce\x2d\x54\x6a\xd2\x85\xfc\x19\xf0\x00\x57\x3c\x16\xb3\xfd\x55\x5c\xf0\xe7\xe3\x3d\x90\xac\xc6\xfb\xb0\x9e\xfe\x12\x00\x00\xff\xff\xfd\x91\x67\x29\xcd\x14\x00\x00")

func golangInternalModelDatabaseTableGoTplBytes() ([]byte, error) {
	return bindataRead(
		_golangInternalModelDatabaseTableGoTpl,
		"golang/pkg/model/database/table.go.tpl",
	)
}

func golangInternalModelDatabaseTableGoTpl() (*asset, error) {
	bytes, err := golangInternalModelDatabaseTableGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang/pkg/model/database/table.go.tpl", size: 5325, mode: os.FileMode(420), modTime: time.Unix(1586257494, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"golang/pkg/model/.DS_Store":             golangInternalModelDs_store,
	"golang/pkg/model/database/db.go.tpl":    golangInternalModelDatabaseDbGoTpl,
	"golang/pkg/model/database/table.go.tpl": golangInternalModelDatabaseTableGoTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"golang": &bintree{nil, map[string]*bintree{
		"pkg": &bintree{nil, map[string]*bintree{
			"model": &bintree{nil, map[string]*bintree{
				".DS_Store": &bintree{golangInternalModelDs_store, map[string]*bintree{}},
				"database": &bintree{nil, map[string]*bintree{
					"db.go.tpl":    &bintree{golangInternalModelDatabaseDbGoTpl, map[string]*bintree{}},
					"table.go.tpl": &bintree{golangInternalModelDatabaseTableGoTpl, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
