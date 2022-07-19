// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package configs generated by go-bindata.
// sources:
// configs/config.yaml
package configs

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x92\x41\x73\xda\x40\x0c\x85\xef\xfe\x15\x9a\xc9\xb9\xc6\x84\x86\x30\x7b\xa3\x09\x99\x92\x09\x2d\x53\x3b\xc3\xb1\x23\xb0\x30\xdb\xae\xad\x65\x57\x4e\x9c\xfe\xfa\xce\xae\x63\xa0\xcd\x6d\xf5\xac\x79\xd2\xf7\xe4\x2b\xc8\xc9\xbd\x90\x83\x1d\x37\x7b\x5d\x25\x7d\xa5\x12\x80\x1f\x6d\xb3\xe2\x92\x14\x94\xb4\x6d\xab\x04\xe0\xab\x88\x5d\xb3\x13\x05\xb3\x6c\x96\x85\x0e\xc2\xb2\xd0\x35\x71\x2b\x0a\xa6\x41\xd9\x38\x2d\x74\x29\x25\x57\x30\xb7\xd6\xe8\x1d\x8a\xe6\x66\x18\x32\xb7\x36\x4c\xb8\xa7\x3d\xb6\x46\xd6\x58\x51\xae\xff\x90\x82\x71\xf0\x58\x61\x77\xa9\x04\xe9\x89\xab\x1c\x5f\x68\x8d\x72\x50\xe0\x85\x1d\x56\x34\x32\x5c\xf9\xfe\xdb\x83\x36\xf4\x0d\x6b\x52\x80\xd6\x9e\xa5\x45\x27\x0a\x52\xc3\x61\xf7\x67\x6b\x18\xcb\x8f\x26\x6d\xd4\xfd\xb9\x23\xe2\x3f\x3b\xa3\xe0\x20\x62\xd5\x68\x34\xbe\xbe\x4d\xb3\x34\x4b\xc7\x2a\x50\x8f\xbc\xa0\xe8\xdd\xa9\x7f\x59\x63\x45\x2b\xec\xfa\x6d\x6f\xfe\xd5\xe7\xc6\xf0\xeb\xa2\x13\x1f\x60\x01\x3e\x41\xfa\xcb\x56\xc3\xd3\x36\xe1\x79\xc7\x8d\x50\x27\xff\x45\x76\x8f\x82\x5b\xf4\x34\xe4\x35\xd4\x31\xb4\x2f\xc5\x9b\x25\x05\xf5\x9b\x3f\x9a\x30\xcf\x93\x6b\x22\xbc\x63\x96\x04\x60\x8d\xde\xbf\xb2\x2b\x15\x84\x9b\xb1\x17\x05\x67\x86\xc9\x24\x9b\x46\x93\x3e\xaf\xad\xe1\x4a\x53\x02\x50\xe0\xd6\xd0\xda\xd1\x5e\x77\xbd\xfa\x33\x2c\x77\x40\xe7\x49\x14\xb4\xb2\x9f\x45\x67\xe7\xe3\x75\x15\x14\xae\xa5\xfe\x56\xcb\xd2\xd0\x1d\x37\x8d\x3f\x9f\xef\xbb\xa5\xe6\x5d\x9a\x44\x9e\xc7\x4d\x31\xa0\x3c\x6e\x8a\x40\x91\xd3\xce\x05\x67\x9d\x2d\x7e\x4f\x12\x80\xa5\xf7\x2d\xb9\x8b\x85\x16\x9d\xd5\x8e\x14\xdc\x5e\x67\xd1\x62\x51\xa3\x36\x83\x49\x2c\xd4\x09\xcf\xd7\x62\xd3\xe3\x31\xdd\x71\x1d\xb6\x8c\xbf\xe8\xe7\xe9\xcd\x7b\x38\x91\xf4\x32\x97\x38\x2e\xcf\x9f\x4e\x14\x0f\x8e\xeb\xa0\x16\xac\xfe\x06\x00\x00\xff\xff\xa1\x6b\xbd\xc8\x0f\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 783, mode: os.FileMode(420), modTime: time.Unix(1658147016, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"configs/config.yaml": configsConfigYaml,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
