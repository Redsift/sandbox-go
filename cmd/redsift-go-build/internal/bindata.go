// Code generated by go-bindata.
// sources:
// main___.go.tmpl
// DO NOT EDIT!

package internal

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _main___GoTmpl = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\xdf\x72\xdb\xb6\xd2\xbf\x26\x9e\x62\xcb\xc6\x29\x99\xa1\xa9\xaf\xed\x37\xe7\x42\x89\x2f\xd2\xd8\xcd\xc9\x8c\x6b\x67\xac\x74\x3a\x1d\xd7\xe3\xc2\x04\x44\xa3\xa6\x16\x0c\x00\xca\xf6\xa8\xbc\x3a\xe7\x7d\xfa\x0e\x7d\x95\xf3\x22\x67\x16\x00\x25\xca\x51\x9c\x76\x4e\x7c\x61\x8a\x24\xb0\xbb\xf8\xed\x6f\xff\xb1\xe5\xd5\x0d\xaf\x25\x2c\xb8\x42\xc6\xd4\xa2\xd5\xc6\x41\xc6\x92\x54\x62\xa5\x85\xc2\x7a\xf2\x9b\xd5\x98\xd2\x03\x63\xb4\xb1\xf4\x6b\xbe\x70\x74\x69\x74\x4d\x17\x6d\xc3\xff\x89\x55\x35\xf2\x86\x6e\x4c\x87\x4e\x2d\xe4\x44\xc8\xab\xce\xaf\xb1\xce\x54\x1a\x97\xfe\xe7\x3d\x56\xc3\x75\xc2\x9d\x5e\xa8\x78\x6b\x2b\xde\xf8\xdd\xb4\x35\x65\x2c\x49\x6b\xe5\xae\xbb\xab\xb2\xd2\x8b\x89\x91\xc2\xaa\xb9\x9b\xd4\x7a\x7f\xc1\xb1\xd6\x56\x57\x37\xe9\xa7\x97\x4c\x90\xa3\xfe\xf8\x3a\xcb\x51\x5c\xe9\xbb\x7d\xd3\x56\xa4\x70\x32\x01\x23\xdf\x77\xd2\x3a\x29\x00\xb5\x90\x96\x25\xab\xd5\x3e\x18\x8e\xb5\x84\x27\x6d\x01\x4f\x96\x30\x3d\x80\xf2\x2d\x77\xd7\x16\xfa\x9e\x25\xe9\x6a\x05\x4f\x5a\xe8\xfb\x34\x2c\x95\x28\xe8\x79\xce\xd8\x64\x02\xee\xbe\x95\x50\xe9\x45\xdb\x39\xf9\x7d\x87\x15\xcc\x3b\xac\xb2\xa8\xd4\xb4\x55\xf9\x2a\xbc\x3b\x0b\x4a\x73\xc8\xce\x2f\x76\xbd\xb5\xad\x46\x2b\x0b\xf0\x1e\xc8\x19\xab\x34\x5a\xef\xa4\x37\x87\xc7\x47\xa0\xd0\x7d\xfb\x0d\x1c\x80\xd2\x8e\xb3\xe4\xbb\x1f\x67\x3f\xb3\xe4\xd5\xf1\xe9\xec\xe8\x90\xcc\xf0\x36\xd0\x59\xc0\x3a\xd3\x55\x0e\x56\x2c\x51\xe2\x0e\xe8\x4f\xa1\x63\x89\x90\xb6\xa2\x1b\xeb\x8c\xc2\x9a\x25\x9d\x69\x60\x7c\x1f\xed\xff\x4c\xb6\x27\xd6\x71\x27\x83\xee\x6f\xbf\x61\x09\x39\x89\xb4\x91\x9f\xca\x33\xd9\xb2\x9e\xb1\x25\x37\x70\x83\xfa\x16\x4f\xc8\x05\x70\x00\x0b\xde\x9e\x2b\x74\x17\xcf\xe8\x1c\xab\x2d\x9f\xa8\x02\x9e\xa0\xf7\x49\x58\x4c\x3e\x21\x97\x28\xe8\xfb\x29\xac\x94\xb8\x9b\xc2\x70\x5f\x00\x9d\x75\x0a\xde\x67\x58\x1e\xd2\xc1\xfb\x3e\x2d\x06\x17\x85\x95\x58\xbe\xbd\xa9\x4f\xf8\x42\x42\xdf\x0f\xc7\x80\xbe\xd8\x72\x6f\xcf\x18\xc1\x01\x95\xc6\xb9\xaa\x3b\x23\x67\xf7\xd6\xc9\xc5\xb1\xae\x6b\x69\xb2\x1c\x9e\x35\xba\x2e\xc3\x1d\xe1\x3d\x99\xc0\xbb\xd3\xc3\xd3\xcd\x72\xb0\xf7\xb6\xd1\x35\x4b\xae\x0b\xb8\x24\xeb\xb5\x2d\xff\xa9\xad\x43\xbe\x90\x59\xce\x12\x23\x5d\x67\x10\x48\xca\x89\xbc\xcd\xb4\x2d\x67\x4e\x48\x63\x0a\xb8\x2e\xfc\xd3\x63\xeb\xc4\xf7\x0d\xaf\x6d\xbe\xb6\x85\x02\x38\xcb\x49\x1d\x89\x26\xa1\x1f\xb1\x8e\xb1\xc4\x73\xbb\x00\x61\xee\xcf\x3a\xf4\xbe\xd9\x5a\x4f\x0a\x5f\x9a\xda\x9e\x7f\x3d\xbd\xc8\x59\xa2\xe6\x7e\xc5\x17\x07\x80\xaa\x21\x05\x5e\x43\xf9\xd6\x28\x74\xf3\x2c\x55\xb8\xe4\x8d\x12\x10\xbd\xbf\x16\xc3\x9d\xd2\x38\x85\x3d\xfb\x0b\xa6\x5e\x47\xce\x92\x44\xdb\xf2\xe8\x4e\xb9\xec\xeb\x9c\x25\x3d\xf3\xb2\x83\x15\x0f\xe4\x36\x98\xfd\xfa\xe7\x1f\xbf\x5c\x66\xff\xf9\xd7\xbf\xf3\xcb\xc9\x9f\x7f\xd0\x32\x30\x1d\x96\x65\xf9\x2b\x09\x0a\x10\x05\x21\x73\xee\x78\x43\x27\x58\xf0\x1b\x99\x55\xd7\x1c\x23\xd9\x57\x7d\xce\x12\xa2\xd3\x6d\x0d\x94\x6f\xca\x9f\xb8\x72\xaf\x8d\xee\x5a\xef\x94\x6a\x5e\x97\x1e\x0a\x30\x72\x2e\x8d\x05\xa7\x03\xf1\xe0\x59\x78\xac\xb1\xb9\x67\xc9\x5c\x1b\xb8\x2c\xc0\xd3\x2c\xd0\x2e\xbc\x25\x8b\x6f\xeb\xf2\xa5\x10\xfe\x38\x49\xad\x43\x90\xc4\xed\xde\x17\x49\x22\x48\x34\xdc\xd6\xe5\xa1\x46\xef\xdc\x24\x19\xa3\x67\x1d\x37\x4e\x61\x1d\x62\x74\xcf\x02\x77\x03\x64\x58\x00\x96\x9d\x69\xfc\x9e\xe8\x84\xe9\x01\x60\xd9\x28\xeb\x24\xbe\x44\x31\x93\x66\x29\xb3\xfc\xf9\x43\xff\x3c\xf0\xd0\x20\x5a\xde\xc9\xaa\x23\xaf\xc0\x9c\xab\x46\x8a\xe9\x48\x53\xf4\x4f\x12\xc1\x7c\xb1\xbf\x86\x70\xd5\xd3\x73\xfa\xd7\x67\x18\xdd\x66\x55\x6d\xb7\x01\x27\x8e\xfa\xfc\x9f\xfb\xb7\xc8\x9b\xf2\x44\x3b\x35\xbf\xcf\x68\x6d\x01\x31\xbf\x97\xb3\x37\xaf\xdf\x9c\xbc\xdb\xba\x7f\x77\x74\xf6\x43\xce\xd6\xf0\x05\xdc\xc8\x6b\x56\xd5\x1b\xb9\x44\x0f\xd9\xc8\x90\xc1\x92\x8a\x5b\xe9\xdf\x1f\xc0\x8b\x7d\xd2\x30\x7d\x08\x6c\xad\x09\xc8\xe7\xa0\x50\x39\x45\x39\xc7\x5e\x77\x4e\xe8\x5b\xf4\x27\xb6\xaa\xce\x07\x31\x2f\xf6\xfd\x99\xb7\x25\x34\x18\x71\x0b\x50\xed\x90\x93\x92\x00\x82\xe5\x31\x7e\x6c\x99\x54\x35\xda\x8e\x5c\x1d\xa0\xf7\xa8\x63\x49\xef\x02\x3d\x7a\x96\xf4\x3e\x4a\x6f\x6b\xcf\xd7\x6c\x13\xe3\x1b\x6a\xcd\x7c\x76\xce\xf2\x98\xa6\x49\x57\x4c\x1a\xf3\x85\x2b\x67\x6d\xd4\xf8\xe5\x9e\x80\xbd\xf7\xa4\xa6\x54\xe2\x8e\x2e\x94\x02\x77\x09\x7c\xc8\xa9\x90\xb0\x49\x6e\x20\xf0\xc6\xc2\x10\x51\x19\x4b\x12\x62\x9d\x5f\xc6\x92\x64\x61\x6b\x38\xbf\xb8\xba\x77\x92\x25\x79\x08\x6c\x2c\x29\xbb\x87\xf4\x42\x54\x89\x35\x99\x52\xda\x99\x6c\x67\xba\xba\x91\x6e\x07\x79\x47\xe7\x38\x22\xe1\x04\x1c\xc7\xaf\x1c\xd4\xd2\x01\xca\x5b\x30\xb2\x05\xeb\x77\x13\x81\xd7\xb9\x25\xa6\x93\xa0\x2d\xe8\x2e\x5f\x69\x44\x59\xb9\x2c\xc4\xd1\x5f\x57\x25\x14\x6f\x40\xe3\x63\xaa\x8c\x6c\x9b\xfb\x9f\x94\xbb\xf6\x3b\xc9\xf7\x9e\xbd\x32\x20\x52\x80\x75\xbc\xba\x89\x90\x8c\xd0\x4c\x8c\xb4\xed\x3a\xe5\xfa\x26\x4b\x7a\x09\x99\x0c\x36\x64\x79\xdc\x4b\x5c\xf8\x30\xf5\xee\xb2\x39\x08\x1f\x1a\xb6\x51\x70\x6f\xf2\x6e\x1f\x84\x5d\x16\xdb\xf8\xcc\x24\x8a\x8c\x2c\xda\x91\x43\x76\x81\xa3\xbb\x46\x10\x3e\x96\x0a\x61\x50\xeb\x71\xf8\x50\x59\xdc\x8d\xaa\x09\x70\x4d\x26\x9e\xf6\x5f\x51\xae\x7d\x3f\x21\x5c\x1b\xad\xdb\x90\x5c\x57\xc1\x38\xdf\x5b\x85\x8e\xb0\x3c\xd6\x5c\xbc\xa1\x16\x21\x7b\x8a\xa5\xef\x19\xf2\xe7\xb0\x84\x83\x03\x08\xad\x4d\x30\xf1\xca\x48\x7e\x13\x15\xc6\x8d\x33\xa7\x8d\xdc\xde\x59\x00\xf5\x49\x79\x20\xe9\x03\x00\xce\x64\xb5\xcc\xf2\xcf\xa1\xff\xef\xb8\xca\xc8\x4a\xaa\xa5\xf7\x95\xd1\x8b\x5d\x0c\xfb\xe4\x91\xa8\xcf\xa3\x38\x4b\x8c\x7c\xbf\xa6\x93\x90\x44\xa7\xd8\x92\x65\x0b\x5b\x3f\xce\xa1\x6d\x0a\x67\xbe\xc1\x40\xd5\x04\xf5\x94\x6f\x9d\x6e\xdf\xa9\x85\xf4\xb2\x7d\xa1\xf2\x77\x59\xa8\xbf\x36\xf2\xb4\xd8\x14\x26\x2a\x30\xb1\x65\xca\x8c\x7c\xef\xed\xdb\xa1\x5e\xcd\xc1\xc6\x3d\xbb\x4c\x08\xec\x7f\x1e\xd6\x6c\x55\xb5\x68\x37\xbd\x18\x6a\x52\x52\x69\x74\x0a\x3b\x39\x18\xbd\x23\xbe\x86\x2e\x34\x8b\x46\xc7\x53\x65\xf9\xff\x06\xcf\x26\x9e\xa6\x9f\x27\xa0\x76\x87\x52\xcf\xc6\xc1\xf4\x61\xe6\x8e\x79\x99\x74\x7c\x82\xc1\xec\x51\x46\x05\x62\x87\x7e\xcf\x33\xdd\x4f\x17\x4f\x9f\xc6\xa3\x8d\x8f\x72\x39\xca\xb0\x43\xdd\xea\x77\x98\xf6\x80\x0f\xf0\xd8\xf8\x60\xa4\x85\xc7\x47\x88\x71\x46\x0d\xb8\x87\xa1\x62\x53\xa4\x46\xfd\x83\x5c\xba\xc0\xaf\x4a\x2f\x23\x65\xc9\xd3\x4b\xb7\xe5\x92\x20\x92\x02\xe7\xaa\xab\xcb\x19\xdd\xf9\x0a\x17\x6a\xdb\x41\x50\x60\x7d\x13\x9e\xb6\x1c\x55\x95\x0e\xcd\xd8\xb2\x00\x7d\xe3\x29\xb6\x74\x65\x16\x0c\x79\x4e\x8f\x02\x53\xc3\xf6\xe5\xba\x73\x1a\x6a\xba\x67\xe0\x90\x81\xaa\x71\xa8\x0c\x1d\xed\x80\xe2\x38\xde\x86\x83\x9d\x5f\x28\x74\xff\xf8\x7f\xdf\xe4\xd3\x6b\xd2\x4f\xd3\x72\x79\xa2\x6f\x47\x73\xc3\x87\xab\x13\xb1\x5e\x3a\x53\x58\xc9\xcc\x6f\xcf\x7d\x4f\x55\x69\x14\xbe\x99\xf3\xab\x33\x51\xce\xc2\xa3\x10\x1e\x51\x64\x94\xb5\x8a\xcb\x8b\x61\x71\x0e\xfb\x10\x9f\xf5\x63\x0e\x6c\x66\x09\x6e\x6a\xef\x58\xdf\xa7\xf8\x31\xd1\x93\xa3\x80\x2b\xad\x9b\x62\xec\xc2\x38\x8d\x5c\x86\x6c\xe3\x3f\x18\x94\x6f\xb9\xb1\xf2\x3b\xad\x1b\x9a\x49\x5e\x4b\x27\x71\x99\xa5\x87\x67\x3f\xa7\x79\x6c\x32\xbc\xf8\x83\xe0\xd2\xdf\x7f\x87\x46\xa2\xd7\x98\xd3\xb3\xff\x1b\x17\x79\x54\xcd\xd6\xc4\xb3\xf6\x2b\xea\xd8\xaf\x39\x1d\x3b\x64\x99\x0e\xed\x44\x5b\x9d\x69\xed\x0a\x98\xeb\x0e\x45\x1c\xd4\x8e\xb5\xbe\xe9\xda\x23\x32\xe4\xcd\xdb\x57\x97\x67\xa7\xa7\xef\xd2\x10\x35\x5f\x84\x65\x7f\x45\xa9\xc4\xa5\x32\x1a\x17\x12\x1d\x2c\xb9\x51\xfc\xaa\x91\x30\x88\x03\xd4\x2e\xa8\x1c\x0c\xa1\x8e\x2b\x18\x19\xe1\x63\xeb\xb1\x84\x9b\x7a\xd3\x78\x7a\x34\x7c\x39\xa5\x6e\x2f\xa6\xa6\x01\xcb\x97\x4e\x2b\x02\xe7\xf1\xac\xb7\x65\xf2\xce\x5c\xd5\x92\x4f\xbc\xde\x3d\x1b\xb3\x15\x37\xf5\x56\xe5\xc2\x11\x64\x9b\x51\xfe\x5c\x89\xbb\x8b\xa0\x7c\x04\xd5\x27\x15\x6f\xbe\xc9\x38\x4d\x23\x20\x70\x84\x0e\x47\x83\x1a\x7c\xb9\x27\xd2\x02\x94\xb8\x5b\xeb\xa7\x7e\x0f\x0e\xb6\x9b\x61\xd5\x56\xd3\xc9\x64\xcf\x4e\xf6\x84\xcf\x5d\xb4\x65\x70\x70\xdc\x8b\xf1\x8b\x03\x6f\x5b\x4a\xe4\x71\x4e\xc6\x75\xcb\x17\x0c\xdd\x9e\x9e\xc7\x59\xf9\x41\xbd\x11\xdc\xf1\x4f\x24\x35\x11\x07\xe5\x21\xc2\x7c\x84\xac\x33\x5c\x0c\x0d\x72\x7f\x2b\xc8\xf7\x1f\x17\xb5\x26\x84\xd8\xd0\xc1\xeb\x27\x8c\x7d\xa6\x10\x2c\x49\x5a\xb1\x39\x5d\x2b\x0a\x78\xea\xf2\x71\x85\xf9\xcd\x6a\x2c\x7f\xe0\xc6\x5e\xf3\x66\xfc\xb1\x67\x50\x42\xb2\x4e\x3b\x37\x05\x80\x56\x14\x2c\x49\x66\x8e\x3b\x3b\xf5\x5f\x68\x42\x84\x5f\x0c\x99\x22\x35\xd2\x76\x8d\xb3\xe9\x74\x7d\xc6\xbe\x60\x49\x9f\x3f\x00\x2b\x14\xd7\x85\xb4\x96\xd7\x32\x8e\x33\x0f\xbb\xe7\x1d\x90\xfc\x0d\x83\xbd\x86\x2d\x1b\xc3\xc5\x93\x2f\x8d\x9a\xd3\x29\xc4\x5f\x85\x7f\xec\x0d\x48\xa7\xeb\x0f\x61\x59\xe8\x4a\xfa\x91\xfd\xdb\xdd\x16\x19\x68\x37\x16\x7f\xb4\xd4\x8d\x0f\x51\x85\xf8\xfc\xd8\x52\x9a\xbc\x63\x0c\xfb\x73\xfe\x88\x8b\x78\x52\xaf\xac\x80\xa7\x95\xd9\x64\xfd\xca\x78\xd1\xac\x67\xff\x0d\x00\x00\xff\xff\xfc\xd4\xbe\x58\xd4\x15\x00\x00"

func main___GoTmplBytes() ([]byte, error) {
	return bindataRead(
		_main___GoTmpl,
		"main___.go.tmpl",
	)
}

func main___GoTmpl() (*asset, error) {
	bytes, err := main___GoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main___.go.tmpl", size: 5588, mode: os.FileMode(420), modTime: time.Unix(1544205980, 0)}
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
	"main___.go.tmpl": main___GoTmpl,
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
	"main___.go.tmpl": &bintree{main___GoTmpl, map[string]*bintree{}},
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

