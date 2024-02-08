// Code generated by go-bindata.
// sources:
// 1_test.down.sql
// 1_test.up.sql
// 3_test.up.sql
// 4_test.down.sql
// 4_test.up.sql
// 5_test.down.sql
// 7_test.down.sql
// 7_test.up.sql
// DO NOT EDIT!

package testdata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
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

var __1_testDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1_testDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_testDownSql,
		"1_test.down.sql",
	)
}

func _1_testDownSql() (*asset, error) {
	bytes, err := _1_testDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_test.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440324, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_testUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1_testUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_testUpSql,
		"1_test.up.sql",
	)
}

func _1_testUpSql() (*asset, error) {
	bytes, err := _1_testUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_test.up.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440319, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_testUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _3_testUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_testUpSql,
		"3_test.up.sql",
	)
}

func _3_testUpSql() (*asset, error) {
	bytes, err := _3_testUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_test.up.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440331, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_testDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _4_testDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_testDownSql,
		"4_test.down.sql",
	)
}

func _4_testDownSql() (*asset, error) {
	bytes, err := _4_testDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_test.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_testUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _4_testUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_testUpSql,
		"4_test.up.sql",
	)
}

func _4_testUpSql() (*asset, error) {
	bytes, err := _4_testUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_test.up.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_testDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _5_testDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_testDownSql,
		"5_test.down.sql",
	)
}

func _5_testDownSql() (*asset, error) {
	bytes, err := _5_testDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_test.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440340, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_testDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _7_testDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_testDownSql,
		"7_test.down.sql",
	)
}

func _7_testDownSql() (*asset, error) {
	bytes, err := _7_testDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_test.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440343, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_testUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _7_testUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_testUpSql,
		"7_test.up.sql",
	)
}

func _7_testUpSql() (*asset, error) {
	bytes, err := _7_testUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_test.up.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486440347, 0)}
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
	"1_test.down.sql": _1_testDownSql,
	"1_test.up.sql":   _1_testUpSql,
	"3_test.up.sql":   _3_testUpSql,
	"4_test.down.sql": _4_testDownSql,
	"4_test.up.sql":   _4_testUpSql,
	"5_test.down.sql": _5_testDownSql,
	"7_test.down.sql": _7_testDownSql,
	"7_test.up.sql":   _7_testUpSql,
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
	"1_test.down.sql": &bintree{_1_testDownSql, map[string]*bintree{}},
	"1_test.up.sql":   &bintree{_1_testUpSql, map[string]*bintree{}},
	"3_test.up.sql":   &bintree{_3_testUpSql, map[string]*bintree{}},
	"4_test.down.sql": &bintree{_4_testDownSql, map[string]*bintree{}},
	"4_test.up.sql":   &bintree{_4_testUpSql, map[string]*bintree{}},
	"5_test.down.sql": &bintree{_5_testDownSql, map[string]*bintree{}},
	"7_test.down.sql": &bintree{_7_testDownSql, map[string]*bintree{}},
	"7_test.up.sql":   &bintree{_7_testUpSql, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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