// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package jsonscheme generated by go-bindata.
// sources:
// ../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml
package jsonscheme

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

var _ResourcesComponentDescriptorOcmV3SchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1a\xef\x6f\xdb\xb8\xf5\xbb\xfe\x8a\x87\x4b\x01\x39\x4d\x64\x37\xe9\x3a\xe0\xfc\x25\xc8\x7a\x18\x50\x6c\x77\x39\xb4\xb7\x7d\x58\xea\x15\xb4\xf4\x6c\xb3\x47\x91\x1e\x49\xb9\x71\x7b\xfd\xdf\x07\x92\xa2\x44\xc9\x92\x7f\x26\xdd\x86\x5d\xbe\xc4\xa4\xde\x2f\x3e\xbe\xdf\xd2\x33\x9a\x8d\x21\x5e\x68\xbd\x54\xe3\xd1\x68\x4e\x64\x86\x1c\xe5\x30\x65\xa2\xc8\x46\x2a\x5d\x60\x4e\xd4\x28\x15\xf9\x52\x70\xe4\x3a\xc9\x50\xa5\x92\x2e\xb5\x90\x89\x48\xf3\x64\xf5\x92\xb0\xe5\x82\x5c\xc5\xd1\x33\x07\x1b\xd0\xfa\xa8\x04\x4f\xdc\xee\x50\xc8\xf9\x28\x93\x64\xa6\x47\xd7\x2f\xae\x5f\x24\x57\xd7\x25\xe9\x38\xf2\x04\xa9\xe0\x63\x88\xef\x5e\xff\x08\xaf\x3d\x33\xf8\xa1\x62\x06\xab\x97\x50\x63\xcc\x28\xa7\x06\x41\x8d\x23\x80\x1c\x35\x31\xff\x01\xf4\x7a\x89\x63\x88\xc5\xf4\x23\xa6\x3a\xb6\x5b\x4d\xea\xd5\x31\x60\x85\x52\x51\xc1\x2d\x72\x46\x34\x71\xd0\x12\xff\x55\x50\x89\x99\x23\x07\x90\x40\xcc\x49\x8e\x71\xbd\x2c\xf1\xdc\x0e\xc9\x32\x2b\x06\x61\x3f\x4b\xb1\x44\xa9\x29\xaa\x31\xcc\x08\x53\x68\x9f\x2f\xeb\xdd\x92\x82\xa1\xe6\x7f\x03\x3c\x93\x38\x1b\x43\x7c\x36\x0a\x4e\x54\xab\xfa\xa7\x80\x73\xc9\x76\x07\xaa\x44\x46\x1e\x30\x7b\x87\xf9\x0a\xa5\x47\x65\x64\x8a\x4c\xed\xc0\x74\x40\x1e\x65\x29\xc5\x8a\x66\x28\x77\x20\x79\xb0\x38\x8a\x9a\x6c\xca\x7b\x20\x52\x92\xb5\xa3\x49\x35\xe6\x95\x0c\xfd\x12\xc4\x9e\x50\xef\x7d\xee\x71\x43\x84\x15\xe5\x7a\x97\xfe\x1d\x7d\xa5\x25\xe5\x73\xaf\x68\x83\x3d\x86\x2f\x5f\xfb\x14\xbf\x24\x5a\xa3\x34\xc6\xf4\xcf\xd5\xfd\x8b\xe4\xfb\xc9\xc5\x33\xcf\x5c\xd1\x39\x27\xba\x90\x1b\x1c\xe2\xa9\x10\x0c\xc9\x1e\x56\x13\x01\x34\xee\xbf\xa1\x07\x27\xa8\x23\x92\x93\x87\xbf\x22\x9f\xeb\xc5\x18\xae\x5f\xbd\x8a\x5a\x92\xdd\x93\xe4\xf3\xe4\x3e\x21\xc9\x67\x23\xe1\xf3\xc1\xfd\x70\xd2\xda\x3a\x7f\xee\xf7\xbe\x5c\x5f\xfe\xe1\xeb\x60\xd4\x78\xfe\xa1\x03\xe7\x83\x41\x3a\x37\xa7\x8d\x00\x68\x86\x5c\x53\xbd\xbe\xd5\x5a\xd2\x69\xa1\xf1\x2f\xb8\x76\xb2\xe6\x94\x57\x82\x75\x89\x65\xb8\x0f\xee\x93\x0f\x17\x5e\x12\xbf\x79\x7e\xe3\x48\x37\x8c\xd8\xd1\x3c\x03\x4d\x7e\x45\x0e\x33\x29\x72\x50\xf6\x81\x09\x28\x40\x78\x06\x24\xfb\x58\x28\x8d\x19\x68\x01\x84\x31\xf1\x09\x08\x07\xb1\x74\x0a\x06\x86\x24\xa3\x7c\x0e\xf1\x2a\xbe\x84\x9c\x7c\x34\x51\x8b\xb3\xf5\xa5\x45\xb5\xeb\x61\x4e\x79\xb9\xeb\x79\x2d\xa8\x82\x1c\x09\x57\xa0\x17\x08\x33\x61\xa8\x1a\x22\x4e\xff\x0a\x88\x44\xc3\xca\xd8\x0a\xcd\x9a\xf2\x2a\x2f\xf0\xd5\xf0\x7a\xf8\x32\xfc\x9d\xcc\x84\xb8\x98\x12\x59\xee\xad\x42\x80\x55\x17\xc4\xd5\xf0\xda\xff\xaa\xc0\x02\xf8\xea\x67\x03\x2d\x54\xf6\x6a\x72\x33\x78\xf1\xdb\xfd\x55\xf2\xfd\xe4\x7d\xf6\xfc\x7c\x70\x33\x7e\x3f\x0c\x37\xce\x6f\xba\xb7\x92\xc1\xe0\x66\x5c\x6f\xfe\xf6\x3e\xb3\x77\x74\x9b\xfc\x23\x99\x18\x8b\xf7\xbf\x3d\xc9\x3d\x81\xcf\x3d\xc7\x8b\x41\xf8\xe0\xc2\x12\x69\xec\x58\xc8\xd2\xab\x5a\xa6\xdf\x65\x7a\xbd\xb1\xa2\x74\xff\xb5\x71\x24\x35\x86\x2f\xdd\x81\xa7\xcb\x94\x63\xf8\xea\x4c\x71\x29\x14\xd5\x42\xae\x5f\x0b\xae\xf1\x41\x1f\x12\x96\x0c\x54\x5f\x18\xb2\x14\xda\x41\x22\x38\xa3\x48\xe9\xdb\x6e\xde\x84\xb1\xbb\x59\xcd\xa5\x27\x0d\xb4\x50\xeb\xe8\xd8\x96\xb3\x94\x75\x4a\x14\xfe\x4d\xb2\xb8\x8e\x72\x1b\x22\x9b\xbf\x12\x2c\xdc\xea\x0c\x4e\xee\xaf\x11\xc8\x7e\x24\xcb\x25\xe5\xf3\x3d\x51\x01\x90\x17\xf9\x18\xee\xe3\x42\xb2\x9f\x89\x5e\xc4\x97\x10\xab\x05\xb9\x7e\xf5\xc7\x24\xa3\x73\x54\x3a\x9e\x44\x2d\x3a\x87\x52\xb6\x3a\x9e\x53\xa5\xe5\xda\x50\xbf\x7b\xfd\xa6\x5a\x4e\xcc\x1d\x90\x34\x45\xa5\xf6\x2c\x2c\x8c\x66\x2c\x14\xcc\x84\x2c\x51\x51\xc1\xc0\xac\xf0\x41\x23\x37\x49\x44\x9d\xef\x30\x96\x08\x60\x4e\xf5\xa2\x98\xde\x6e\xe7\xbd\xd5\xda\xec\xd2\x98\x40\x70\xa1\x76\x67\x76\x94\x35\xb6\xd5\xe6\x04\xac\xd4\x5f\x32\xda\x81\x6e\xac\x74\x3b\x44\x2a\xf2\x9c\xea\x6d\x3e\xc1\x05\xc7\x53\xf4\x72\xe2\xb9\x7f\x12\x1c\x9d\x61\x28\x51\xc8\x14\x7f\xa8\x1c\xee\x00\x71\x4c\xf9\x51\x2d\xca\xd2\xa2\x5a\x1b\x0a\xd5\xc2\x99\xd0\x01\x55\xcc\x86\xe0\xfb\x07\xbb\x12\x05\x1f\xb4\x24\x6f\x4a\x80\x1d\xa5\xdf\x06\x9d\x47\x28\x54\x0f\x35\xc3\xca\x06\x8f\xa8\x70\x43\xe7\xb6\x6b\xbe\xbe\x9b\x35\x83\x62\x27\x15\x87\x17\xef\x06\x0c\xfd\x78\x0f\x70\xd3\x32\x79\xe0\x08\xc0\xc5\xb8\x77\x4b\x4c\x0f\x30\xae\x05\x51\x8b\x5b\x36\x17\x92\xea\x45\x5e\x9b\x9c\x90\x39\x61\x54\x11\xc3\x68\xf3\xb1\xad\x77\x8f\x6c\x66\x1a\x0c\xb7\x56\xd5\xdd\x42\xec\x51\x88\x77\x43\x44\x41\xad\x7d\xa0\x92\xc8\x16\x0d\x98\x55\x8e\x19\x25\xbf\x78\x4f\x3c\x5c\x27\xe4\xe4\xc3\xb9\xad\x4a\x8e\x1a\xaa\x99\x71\x7e\x59\xa0\x03\x72\x69\x47\xcc\x6c\xb1\x5a\xa9\x05\x82\x36\x68\xab\xfe\x8e\x8d\x5e\xce\x44\xab\x65\x45\xef\x48\xbd\xed\x6c\xcc\x1c\xbf\x1d\x4e\x5e\xfb\xcd\x96\x9e\xac\x13\xb3\x61\x4f\xd6\x07\x95\x4c\xdf\xfa\xb4\xb5\x33\xff\x13\x93\xe2\x50\x22\x4f\xd1\x36\x22\x30\xa8\x27\x26\x4c\xa4\x84\x9d\x97\x69\xa3\x2f\x17\xf9\x80\xfa\x0e\x19\xa6\x5a\xec\x6a\xbd\x7b\xe3\xef\x41\xb1\xd0\x96\xb8\xa5\xd8\xc7\x1e\xb4\x3a\xe7\xbe\xfd\x79\xe7\x7c\xe3\xf4\xc9\x4a\x47\xdb\xdc\x7b\xfe\x4e\x11\xb6\x25\x55\x38\x03\x92\xea\x82\x30\xb6\x1e\xd7\x9c\x12\xeb\x79\x9f\x46\xa0\x96\x98\x52\xc2\x40\xa2\x81\x4f\x2d\x93\xff\xdd\x3c\x7c\x44\x3a\x6d\x3b\xa7\xe0\xd8\x4e\xa7\xa5\x42\x79\xc1\xd8\x1e\xf9\x30\x74\x64\x6b\xa5\xce\x7b\xea\x80\x78\x60\x45\xee\x09\xa8\x43\xe7\x7c\x70\x66\xf1\xad\x0f\xd7\x54\x2e\xcb\x21\x41\xa1\x34\xe4\x44\xa7\x8b\xc0\x0d\xd4\x46\x61\xb7\x59\x9c\x33\x9b\x08\x83\xad\xb0\xae\xf8\xbd\xde\xab\x4e\xe5\x62\xb0\xda\x80\x0a\x26\x8b\xd0\x9e\x2e\xf6\x0a\xe1\x88\xd5\x2d\x89\xbb\x84\xbd\x2b\x4e\x6b\x02\xa6\x53\x34\xfd\x9c\xe4\x84\xfd\x57\xd7\x9f\x22\xa5\x7f\x62\x62\xff\x02\xd4\x9e\xee\xcf\x94\xa1\x5a\x2b\x8d\xf9\xe1\xb8\x77\x5d\x0c\x9f\x3a\x2e\x88\x94\xbe\xc9\xc9\xfc\xa4\xbe\xd0\x2e\xa9\xa1\xf2\xd6\x67\xb6\x47\x69\x18\xc3\xf9\x82\xb7\x94\x26\x9b\x1d\x13\xa0\x5a\x9d\x27\x1c\x8c\x91\xb5\xf7\xb8\xd3\xce\x03\x71\x29\x52\x0c\x75\xef\x3f\xeb\xab\x4e\x6f\xcd\x01\x9a\xa5\x82\x29\x4f\x73\xc2\xe9\x0c\x95\x6e\xd7\xa5\x2d\xa6\x47\x16\xbf\x4e\x33\x2e\x34\x3b\x47\x71\x12\x28\xd0\x62\x07\xc7\xb6\xa1\x6e\xb2\x73\x10\x9e\x95\x26\x72\x8e\x1a\x33\x48\x05\xd7\x55\xf1\xd3\x4b\x5e\xd1\xcf\x5b\xcf\x62\x9e\x03\xe5\x30\x5d\x6b\x54\x9e\xc7\xd4\x28\xbb\x4d\x97\x17\xf9\xd4\xbf\x71\xe9\x73\xd9\x13\xcc\x65\x46\x19\xd6\x99\xf0\x54\x8b\xe9\x90\xb0\xb6\x1e\xcf\xaa\x4f\x2f\xfe\x79\xa8\x0e\xd0\x0b\xa2\x81\x2a\x7b\x76\xa3\x7e\xca\xed\xb3\xef\xcc\x43\xf5\x1d\x64\x54\xda\xea\x79\xdd\x7b\x1f\x5e\x6f\x77\x8f\xe4\x5f\x4f\xa0\xb0\xbb\xb6\x9f\x6d\x37\xce\xa6\x61\x5a\x7f\x87\x4f\x54\x2f\x4a\xd5\xa4\x85\x94\xc8\x75\x5d\xa0\x40\xfd\x06\x77\x9b\x96\x7c\x68\x7d\x5b\xd6\x3c\xa7\xbc\x91\x0b\x2b\xfb\x2e\x25\xfe\x5e\xfd\xec\xce\x25\xf6\x32\x1e\xb3\xe4\xe8\x2b\x1b\x82\x84\xfa\x6d\xd2\x78\x04\x50\x8f\xbf\x4e\x70\xc5\xc2\xcf\xbb\x4f\x4c\xdc\x46\x98\x4a\xd1\xc5\x96\xd9\x76\x04\x30\x47\x8e\x92\xa6\xff\xc1\xb9\x74\x29\x81\x1b\x4d\x97\x8b\x6f\xed\xb3\x8f\x33\xee\xf9\x3f\xf3\xe9\xfa\xe2\xdc\xfe\x53\xb9\x74\xc3\x44\xbf\x55\x61\xde\xfc\x82\xe4\x50\x0b\x7c\x12\x7b\x3a\x74\x32\xa6\xb6\x0d\x96\x9b\x29\xd8\xce\x7f\x66\x34\xb5\x0d\xa5\xcf\xc4\x65\x65\x68\x96\xc1\x94\xcc\x9b\x97\x3e\xf6\xa4\xe5\x04\xe2\x91\x5a\xe2\xd6\xab\xac\xe0\x7d\x9d\x2b\xdc\x1f\x89\x8f\x6c\x76\x56\xf5\x40\xe7\x70\xfa\x1b\x9d\xf2\x96\xd7\xe0\xf5\xd0\x28\xde\x07\xa1\x5d\xf2\xec\x85\xd4\x0a\xb9\x71\x14\xb5\xcc\x25\xb4\x74\x13\x37\x97\xf4\xef\x75\x6c\x4d\x20\xfe\x95\xf2\xac\xfc\x19\x7e\x8c\x96\x38\xb3\x8a\xa3\xa6\x09\xd4\xe8\x0d\xdb\x0c\x4d\x3d\x68\xd8\xf2\x61\xeb\x7b\xbe\xea\x73\x3d\x5b\x5c\x1a\xd6\xbd\x64\x52\xc1\x95\x1e\x43\x5c\x7d\x8d\x17\x88\xed\x05\x75\xc8\x9d\x7a\x31\x20\x71\xd7\x37\x14\xfb\x7d\x23\xd6\xba\xe6\xfe\x1b\xdb\xf8\x4e\x22\x86\x33\x5f\xf4\xb2\xf5\x25\x7c\x42\x10\x9c\xad\xcb\x6f\x83\x6c\x6f\x28\x38\x36\xfc\xbb\xdb\x35\xca\x97\x08\xd5\x8b\x81\x13\xbe\x6d\xab\x68\xc4\xff\x0e\x00\x00\xff\xff\xc4\xd5\x99\x51\x6d\x29\x00\x00")

func ResourcesComponentDescriptorOcmV3SchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_ResourcesComponentDescriptorOcmV3SchemaYaml,
		"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml",
	)
}

func ResourcesComponentDescriptorOcmV3SchemaYaml() (*asset, error) {
	bytes, err := ResourcesComponentDescriptorOcmV3SchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml", size: 10605, mode: os.FileMode(436), modTime: time.Unix(1658840508, 0)}
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
	"../../../../../../../../resources/component-descriptor-ocm-v3-schema.yaml": ResourcesComponentDescriptorOcmV3SchemaYaml,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"..": &bintree{nil, map[string]*bintree{
					"..": &bintree{nil, map[string]*bintree{
						"..": &bintree{nil, map[string]*bintree{
							"..": &bintree{nil, map[string]*bintree{
								"..": &bintree{nil, map[string]*bintree{
									"resources": &bintree{nil, map[string]*bintree{
										"component-descriptor-ocm-v3-schema.yaml": &bintree{ResourcesComponentDescriptorOcmV3SchemaYaml, map[string]*bintree{}},
									}},
								}},
							}},
						}},
					}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
