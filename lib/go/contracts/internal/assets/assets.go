// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/ExampleToken.cdc (7.899kB)
// ../../../contracts/FungibleToken.cdc (7.27kB)
// ../../../contracts/utilityContracts/PrivateReceiverForwarder.cdc (2.601kB)
// ../../../contracts/utilityContracts/TokenForwarding.cdc (2.353kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _exampletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x51\x6f\xe3\xb8\x11\x7e\xcf\xaf\x98\xe6\xa1\x75\x70\x89\x93\x02\x45\x1f\x82\xec\xdd\xee\xb6\xbb\xc0\x3d\xf4\xb0\xb8\xbb\xb6\xaf\xa1\xa5\xb1\xcd\xae\x44\x1a\x24\x65\xc7\x17\xe4\xbf\x17\x33\x24\x25\x92\x92\x9c\x64\x83\xcd\x4b\x6c\x8b\x33\x1c\x0e\xbf\xf9\xe6\x23\x25\xdb\x9d\x36\x0e\x3e\x77\x6a\x23\x57\x0d\xfe\xae\xbf\xa2\x82\xb5\xd1\x2d\x9c\x2f\xaf\xb3\x5f\x97\x55\x5d\x9d\x9f\x9d\xed\xba\x15\x54\x5a\x39\x23\x2a\x07\x9f\x1e\x44\xbb\x0b\xcf\x6f\x0b\x27\x8f\x67\x67\x00\x00\xd7\xd7\xd7\xf0\xbb\x76\xa2\x01\xdb\xed\x76\xcd\x11\xf4\x3a\x33\xb3\x20\x15\xe0\x83\xb4\x0e\x55\x85\x6c\x42\x53\xec\x85\x01\x47\x66\xbf\xb1\xd5\x2d\xfc\xfb\xb3\x7c\xf8\xfb\xdf\x06\x9f\xbf\x39\x6d\xc4\x06\x41\xa8\x1a\xbe\x74\xab\x46\x56\xf0\x45\xb8\xad\xed\x3d\x34\xe8\xe0\x3f\xa2\x6b\x5c\x18\x49\x4f\x6f\x21\xf9\x92\x8d\xfc\x15\x2b\x94\x7b\x34\xde\x95\x1f\x3b\x7c\xce\x86\x7e\x14\x8d\x50\x15\xbe\x60\xe4\x87\xba\x95\x6a\x76\xfa\x24\x3d\x94\x87\x9f\x95\x74\x52\x34\xf2\x0f\xac\xe3\x93\x61\xc4\x16\x01\xf7\xa8\x1c\xb8\xad\x70\x20\x2d\x60\x2b\x9d\xc3\x1a\x0e\x5b\x54\xe0\xb6\x38\xec\x89\xb4\x50\x19\x14\x2e\xb8\xa1\x58\xbc\xe9\x68\x9a\x85\xf4\x9f\xf3\x14\x5f\x94\x81\xfd\x57\xba\x6d\x6d\xc4\x41\xbd\x3a\x2c\xbf\xbf\xc2\x20\x1c\xa2\x0f\x8f\x2d\xe1\x77\x66\x32\xc0\x7e\xba\x85\x68\x75\xa7\x5c\x8c\xeb\x92\x4d\x6f\xe1\x43\x5d\x1b\xb4\xf6\xa7\x51\x9c\xff\xc4\x9d\xb6\xd2\x7d\x43\xfa\x86\x38\xeb\xe8\x03\x9c\x3e\x19\x65\x3f\xd9\x28\x4a\xa7\x4f\xc4\xf8\x2f\xa9\xbe\x21\x40\x85\x87\x34\xc8\x76\x70\x52\x86\xe5\xfd\x17\x31\x8d\xa2\xf8\xd8\x19\xf5\xc6\x34\x59\x67\xf4\x71\x26\x08\xef\x7e\x3e\x08\x0e\xd2\xfc\x23\x01\xe9\x2b\xa2\x10\x9c\x0d\x4e\x81\x01\x83\x56\x77\xa6\xc2\x79\xd0\x67\x73\x2d\x44\xd3\xe8\x03\xd6\x1f\xe6\x22\xe3\xc8\xdf\x16\xd9\x8a\x5d\xbc\x20\xb2\x6c\xae\x45\x12\xc4\x00\xba\x74\xf2\x4f\xa2\xda\x42\x67\xd1\x80\x75\xda\xa0\x05\xa1\x40\x2a\xeb\x88\x8a\x88\x53\xb5\x6a\x8e\x4c\x04\x6c\x4e\xa4\xea\xb6\x28\xfd\x68\xb1\xc1\x6c\x11\xeb\x4e\x55\x4e\x6a\xcf\xbd\x83\x0d\x51\xe9\x46\xef\x91\x76\x0f\x56\xde\xdb\xce\x78\x8a\xdd\x69\xeb\x88\x63\x6a\xc9\x86\xbd\x3b\xa9\x0a\xda\x8f\x84\x74\x64\xa0\x54\xa2\x69\xb0\x5e\x66\xb3\x57\x5b\xac\xbe\x5a\xd8\x8a\xdd\x8e\xb2\xe6\xc0\x74\xca\xc9\x16\xd9\x14\xf7\x68\x40\xf4\x11\x72\xfa\x72\x1f\xbd\xaf\x5f\x43\x8a\x69\x84\xf2\xeb\x5f\x61\x4c\x76\x5c\x19\xd1\x22\x3e\x38\xca\x50\xc6\x92\xbc\x83\x14\x66\xef\xce\xe3\x7a\x2d\x15\x1b\x5f\x82\xd5\xf4\xdc\xf0\x0e\x2a\x0d\x07\x71\x84\xb5\xa6\xd8\x5a\xd1\xc8\x4a\xea\xce\xfa\xed\x70\x3a\xcc\xe9\xb3\x38\xa4\x46\x77\x61\x5a\xa9\x40\x48\xb3\x84\x0f\x60\x77\x58\x49\xd1\x04\x54\x0e\x20\x51\x88\xb5\x25\x4f\xab\x21\x06\xa7\x19\xe5\xbd\xbb\x81\x04\xf2\x54\x10\xa2\x7a\x47\x1c\x42\xd1\x89\x97\x5f\x8c\xde\xcb\x1a\xcd\x65\xf1\x7b\xec\x79\xe5\xef\xa1\xc1\xc5\x0e\x9e\xee\x1d\xb7\x64\x58\x85\x01\x7e\x75\x16\xf6\x3d\x62\xd3\xf6\x1d\x46\xe5\xad\xdb\x3b\x03\xd9\x77\x21\xde\x96\xe8\x90\xc0\x10\x97\xc2\x49\x25\x08\x10\x36\x7a\x5b\x32\x5c\x14\x9e\x2f\xe0\xb1\x7f\x4e\x7f\x16\x9b\xf5\x32\xba\x7c\x17\x9d\xf7\x43\x9e\xf2\x65\xc5\xd6\x94\xfe\x98\x0d\xf8\x1c\xb1\xe8\x31\x23\xbe\xfa\xe2\xf3\xf4\x06\xc2\x7f\x31\x9b\xae\x45\xe5\x32\x43\xaa\x9b\xe8\xdd\x7a\xeb\x60\xc4\x4d\xb0\x2f\xbc\xe5\xec\xd4\x3f\xbb\x80\x2d\x1b\xd8\xc5\x21\xe9\x35\x61\x8e\xa1\x64\x23\x11\x75\xd6\x23\x66\xab\x9b\x3a\xf3\x40\x93\xb4\x5a\xe1\xb1\x1f\xba\x42\xa9\x36\xe0\x8c\x50\x76\x8d\xc6\x60\xbd\xa4\x69\x0c\xba\xce\x28\xcb\xe3\x15\x1e\x9a\x63\xe6\x25\x16\x55\x98\x54\x67\xa5\xc5\x8e\x7d\x91\x52\xd1\x48\xc7\xf5\xb8\x4a\x9a\x69\xe6\x0b\x1b\x8b\x07\x2a\xac\xe9\x65\x13\x7a\xd6\x9d\xea\x13\x57\xb6\x91\x5b\x78\x9f\xa3\xd5\xc7\x74\x12\x01\xd9\xd7\xab\xb0\x09\x99\x01\x11\xf9\xac\xfe\xf0\xff\xa3\xfe\x60\x67\xfa\xa0\xd0\xfc\xb4\x14\xbe\xcf\x5f\x64\xbe\x7c\x2a\xe1\xee\x2a\xa5\x85\x01\xb3\xde\xdb\xc5\x1c\x1c\x43\xd2\x5e\x87\xc6\xb0\x31\x7a\xf5\x3f\xac\x4a\x48\x32\x0c\x45\x5d\xdb\xcc\x8d\x74\xb6\xaf\xba\xb0\x9f\x59\x55\x23\xf0\x12\xed\x0b\x10\x2a\x2d\x84\xbe\x4a\x9e\x82\x34\x60\x17\x96\xa6\xf7\xa1\xad\xb0\x12\x9d\xc5\x01\xf4\x79\x0d\x52\xc8\x09\xb8\x09\xc6\x68\x62\x24\x81\xf5\x98\x80\xd8\xf6\x2f\x43\xec\x5b\x91\xaf\x6b\x85\xa8\x08\x99\xb6\x6b\xb1\xe6\xa5\x33\x89\xaf\x35\x37\xa3\x00\xcb\x20\x5e\x4e\x03\x30\x6c\xc4\xc2\xef\xfa\x14\xe8\x4a\xde\x21\xc9\xcf\x54\x08\x77\x57\x41\xe7\xda\x3f\xc1\xfb\xf4\xb4\xb3\xcc\xd7\xfe\x1c\x56\x7f\xf0\xfe\x96\x25\x85\x15\x90\x1d\x8b\xd1\xcc\xcc\x6b\xd2\x67\x71\x9b\xd9\xc0\x3b\xb8\x59\xde\x64\xcf\xe3\xce\xe6\x6c\x9f\xc0\x37\x0c\x58\x94\x79\x91\xeb\x7c\x55\x3f\x92\xeb\x62\x0c\xfd\x65\x89\x4a\x0e\x7f\xf0\x6e\xfe\xd1\x55\xe6\x3a\x73\xf9\x74\x96\x7f\x7a\x1a\x24\x96\xaf\xcc\x4f\xed\xce\x1d\xa7\xd5\x56\x5e\x65\x39\x07\x7b\x40\x13\x3f\x81\x48\x8b\xe6\x0f\x34\x7a\x50\x13\xaa\xee\x39\x55\x0e\x94\x29\x9a\x86\xd8\x37\x50\x27\x49\x02\xd6\x10\x6d\x67\x3d\x85\xfa\x7e\x1a\xd5\x4f\xe6\x8d\x65\x1f\x7b\xf1\x7e\x7b\x3a\x2e\xa5\x1e\xfd\xa0\x4d\xed\xa5\x09\x57\xa6\x7f\x3e\x78\xab\x2a\xee\x42\x5e\x6f\x88\x55\xc3\x14\x60\xbc\x1a\x88\xb8\xb7\x7d\x77\xe7\xf2\x03\x77\xdc\xe1\x58\x78\x50\xa1\x94\xc9\x5c\x10\x47\x97\xac\xfc\x0c\x29\xde\x2c\x6f\x2e\xd2\x4d\xca\x44\x0d\x1f\xa3\xa5\x75\x46\x38\x6d\x4a\x55\xe2\xfd\xfd\x82\x07\xaf\xa9\x5e\xc8\x9b\xfd\x8e\x26\xdb\x34\x79\xb2\x38\x49\x11\xc5\xdc\x33\xc7\x8b\x5b\x78\x1f\xf4\xde\xe3\xb8\x80\x4f\x9e\x4f\xb2\xaf\xa7\x9b\xcc\x74\x04\x33\x0e\x9e\x66\x52\xe8\x8f\x24\x6f\x4e\x61\x71\x04\x7a\x59\x0a\xfd\xdc\x8c\x1d\xff\x71\x2a\x5b\xe5\x99\xe9\x54\x46\xa2\xc3\x79\x16\x48\x10\x33\x75\xae\x88\xed\xd4\x37\x5a\x2e\x02\x41\x48\x8c\xf5\xe3\xcf\x1d\xd4\xaa\xa2\x56\x7f\x99\x46\xef\xc1\x30\x52\xd7\x41\x1d\x52\xe1\xf9\xb3\x76\x3c\xa5\x44\x54\xe6\xad\xb6\x3f\x1e\x40\xa2\xba\x27\x31\x98\x4f\x45\x76\xbe\x71\xbc\x70\xab\xc9\xc0\x26\x8b\xbb\x64\x3d\x41\x81\xb5\x91\xd9\x5c\x72\xc5\x77\x39\x52\xc1\x89\xba\x6c\xe7\xb8\xf0\x24\x4c\x86\x90\x27\xf4\xe0\xb8\xc1\x16\xd8\xa1\x43\xec\xb8\xdd\x84\x6c\x73\x37\xba\x85\x73\x9f\xb1\x70\xb9\xe2\x19\x79\x85\xb0\x61\x30\x19\xca\x83\x62\x86\x3f\x9f\xf3\x73\x17\x7a\x77\xb1\x01\x33\x7e\x1b\xb4\xd6\x3b\xa5\x5c\xc4\x4d\xf5\xae\xce\x67\xda\x18\x7c\x63\x8f\xfc\x61\x4a\xf1\x8e\x63\x85\xa9\x05\x3c\x2b\x97\x8b\x1b\xa7\x52\xdd\xc2\x9b\x04\x31\x9f\xf6\xa6\x59\x75\x4a\xf1\x97\xcb\xc9\xbe\xcf\xf3\x40\x42\x7b\x6f\xe7\x01\x22\xbf\xe7\x39\xa0\xa7\xb8\x5c\xbc\x76\x46\xbd\xaa\x30\x83\xe2\x1a\x4e\x00\xf1\x46\xe8\x12\x70\xbd\xc6\xca\xc9\x3d\x36\x47\xf6\xcb\x87\xbe\x41\x4c\xcf\x4e\xf0\x8b\x76\x78\xeb\xcf\x03\x5e\x64\x24\xd7\x7e\xa2\x73\xba\x15\x4e\x52\xe9\x1e\xc1\x76\x2b\xbe\x4b\xc1\xba\x3f\xcf\xe6\x47\xcf\xf4\xd6\x3f\xbb\x68\xe2\xb0\xbb\xca\x69\x73\xba\xea\x87\x7c\x7c\x77\x15\x4e\x56\x22\xe2\x66\x5e\x74\x4f\x6b\xe0\xa2\x24\x8a\xfb\xcf\x31\xbe\x13\xfc\x31\xc2\xd3\x25\x30\x90\xf3\xc2\xfe\xeb\xcd\x0d\x69\xf1\x7c\x48\xf9\x6a\x03\xde\xc1\x75\x10\x80\xd7\x98\xac\x35\x5f\x2a\x9b\x8e\xdf\x75\x90\xf1\x8e\xbf\x65\xb6\x71\x60\x6e\x3e\x7a\xff\x31\x63\xfd\xb1\xc8\x1f\x1b\x97\xaf\x44\xe6\xc2\xe6\x71\xd9\x95\x91\xef\xfa\x09\x8a\x58\x81\x97\xbd\x27\x69\x9e\x2c\x9a\xc5\x1e\x49\x7f\x4b\x95\x5d\x84\x7a\x97\x67\x93\x90\x99\x26\xa9\x72\x5b\x2e\xf2\x65\x05\x2a\x58\xd2\x7c\x8b\xbb\x2b\x76\x96\x1c\xbb\xca\xcd\xba\x98\x5a\x99\x00\x9f\x44\xa8\xc4\x4e\xac\x64\x23\xdd\x31\xf6\x4a\xd6\xfe\x75\x7a\xe7\xc3\xd7\x9d\xf8\xb0\xd3\x16\x53\xb2\xe0\xd1\xf7\x41\xc2\xdf\x43\x8b\x6e\xab\xe9\x08\x6c\x74\xb7\xf1\xc9\xba\x8f\x9b\x7a\x0f\xac\x29\xd6\xa2\x9a\xcc\x49\xb6\xac\x46\xaa\xaf\x77\x7f\x7e\x9c\xbe\x3e\x7c\xfa\x71\x31\xa6\xe2\x31\xc6\x2e\xb3\x41\x4e\x98\x0d\xba\x99\xf4\xf4\x23\xbf\x73\x9e\xc2\xee\xde\xc3\x5a\x62\x53\xa4\xe9\x63\x7c\xf6\xda\x2c\x8d\x89\xe6\x71\xf2\x7a\x75\x32\x6d\xa3\xda\x7a\x63\xd6\x98\xd6\xb8\x59\x0d\xc8\xce\x8e\x53\x8b\xd3\x40\x66\xdb\x04\xc8\x65\xf9\xe6\x1b\xf4\x89\x38\x50\xa8\xf4\x25\x89\xdd\xea\x43\xa2\x63\xfb\xfb\xf7\x83\xb0\xc9\x25\x70\x3d\x95\xdb\x84\x51\x4f\xbc\xb4\x9c\x2e\xcc\xa7\xb3\xa7\xb3\xff\x07\x00\x00\xff\xff\x28\x6f\x23\x8e\xdb\x1e\x00\x00"

func exampletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_exampletokenCdc,
		"ExampleToken.cdc",
	)
}

func exampletokenCdc() (*asset, error) {
	bytes, err := exampletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6d, 0xc4, 0x91, 0xfc, 0x89, 0xb1, 0x18, 0xc2, 0xf4, 0xda, 0xc4, 0x6, 0x3a, 0x40, 0x24, 0xad, 0xfe, 0xff, 0xd, 0xe9, 0x9a, 0xeb, 0x67, 0x9b, 0xb7, 0x3, 0xd9, 0x7e, 0x6c, 0x92, 0xd3, 0xa3}}
	return a, nil
}

var _fungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x59\x4d\x73\xdb\xba\xd5\xde\xf3\x57\x9c\x49\x16\xb1\xf3\x2a\xf2\x5d\xbc\xd3\x85\x67\x6e\xdb\xdc\xf6\x66\x26\x9b\x4e\xa7\x75\x7b\xb7\x82\xc8\x43\x09\x63\x10\xe0\x05\x40\xc9\xcc\x9d\xfc\xf7\xce\x39\xf8\x22\x29\xda\x56\x12\x6f\x2c\x91\xc0\x83\xf3\xf9\xe0\x01\x74\xf7\xfe\x7d\x55\xbd\x85\x87\x23\xc2\x27\x65\xce\xf0\x69\xd0\x07\xb9\x57\x08\x0f\xe6\x11\x35\x38\x2f\x74\x23\x6c\x53\x55\x6f\xdf\xc2\x2e\xbd\xe4\x77\x3b\xa8\x8d\xf6\x56\xd4\x1e\xa4\xf6\x68\x5b\x51\x63\x55\x11\x50\xfe\x0a\xfe\x28\x3c\x08\xa5\x96\xb0\x69\xa6\x83\xb3\x19\x54\x03\x47\x71\x42\xf0\x86\x9e\xb7\xc6\x76\xe0\xcd\xb6\xfa\xdc\x82\x80\xc1\xa1\x75\x70\x16\xda\x3b\x7a\xdf\x60\xaf\xcc\x08\x02\x34\x9e\xc1\xcf\xa0\x36\xe0\x8f\x28\x6d\xfe\x5e\x05\x64\x8d\xd8\xd0\x4c\xd9\xf5\x0a\x3b\xd4\x9e\x86\xc1\xcc\x91\x62\xef\x96\xed\x9f\x80\x2c\xcc\x6b\x8d\xa2\x18\x91\x43\x84\x62\x07\x85\x0e\x84\x6e\x40\x8b\x4e\xea\x43\xc5\xee\xfa\x59\x04\x5c\x8f\xb5\x6c\x25\xba\x6d\x08\xe1\x7f\xc5\xa0\xfc\x0e\x2c\x3a\x33\x58\x0a\xd8\xaf\xa2\x3e\x82\xa8\x6b\x33\xb0\x6d\xc2\x83\x39\x6b\x17\x9c\x4b\xe1\x49\x4e\xb0\x1d\x82\x0c\xa6\xbc\xd4\x58\x99\x96\x97\x63\xd0\x8c\x09\xce\x1b\x8b\x0d\x48\x1d\x43\x92\xd0\xe9\xb9\x38\x44\x2f\x97\x93\x8e\xc2\x41\x87\xfe\x68\x1a\x07\xd9\x0f\x73\xd6\x68\xd9\x43\xe3\x8f\x68\x63\x3a\x6a\xa1\xa1\x16\x4a\x45\x97\xfe\x69\xcd\x49\x36\x68\x77\x1b\xd8\xfd\x0b\x6b\x94\x27\xfe\x4c\xb3\x76\xbf\x08\x45\x86\x16\x87\x4b\x68\x1c\x9b\xe1\xa6\x4f\xa0\xc1\x5a\x09\x8b\xd0\x5b\xfc\x50\x1b\xdd\x48\x2f\x8d\x0e\x21\xee\x8d\xf3\xd3\x67\x6c\xa3\x45\xe7\xad\xac\x7d\x45\xc6\xe2\x13\xd6\x03\xbd\x84\x18\x96\x76\xd0\x75\x18\x1c\x42\x11\x5c\x0e\xee\x8f\x40\xeb\x38\xec\x85\x15\x1e\x61\x8f\xb5\x18\xc8\x16\x0f\x07\x79\x42\xc7\xc3\xc9\x5b\xfe\x20\xf6\x52\x49\x3f\x52\x0a\xdc\x51\x58\xac\x04\x58\x6c\xd1\xa2\xae\xb9\x2e\x42\x98\x43\x40\x43\x0a\xb5\x1a\x01\x9f\x7a\xe3\x22\x54\x2b\x51\x35\xae\x58\x54\x49\x0d\x46\x23\x18\x0b\x9d\xb1\x98\x2c\x2e\xa1\xd8\x56\xd5\x67\x6a\x1d\x67\xa2\x41\x21\xf4\x0b\x6b\x3a\xf1\x88\x50\x0f\xce\x9b\x2e\x47\x38\x86\x26\x17\x3c\xc5\x66\x1e\x65\x6a\x24\x03\x27\x61\xa5\x19\x68\xb4\xd4\x07\x07\x67\xe9\x8f\x0c\x1f\x2a\x6f\x5b\x7d\x32\x16\xf0\x49\x10\xcc\x06\x04\xb4\x62\xa8\xd1\x73\xee\xf7\x58\xd0\xb1\x81\xfd\x98\xfa\x96\x7b\x80\xc3\x01\xa9\x28\x66\xcd\xf5\xcb\x08\x83\x93\xfa\x30\xb1\x95\x52\x5b\x4c\xdb\x44\x37\x4d\xfb\x2c\x63\x54\x64\x81\x43\xdd\xf0\x54\x1b\xea\x2d\xb5\x4b\x8f\x68\x3f\x78\xf3\x81\xfe\x6f\xd8\x25\x33\x78\x6a\x1b\x5a\x94\x58\x80\x56\x62\x72\x20\x6f\x05\xd4\x48\xa8\x0a\x14\x36\x07\xb4\xe0\x3a\x61\x7d\x5e\x6a\x0b\x0f\x26\xac\x14\xd1\xbd\x01\xa1\x4b\x23\x6c\xaa\xc0\x4f\xb1\x49\x1d\xc5\x64\xe4\x45\x1b\x2b\xce\x93\x58\x42\x6b\x4d\x37\x2d\x12\xe6\xaa\xd0\x43\x5c\xb9\x0d\xf6\xc6\x49\x9f\xcb\x03\x8c\x9e\xad\xf4\xce\xa5\xe2\x22\x8a\xa4\xd0\x7b\x0c\xf8\x56\x68\xd7\xa2\xdd\x56\xd5\xfb\xbb\xaa\xba\xbb\xbb\x9b\x53\x1b\x3d\xe1\xa7\x2b\xb4\xfc\x2c\x25\xe7\xdc\x6e\x79\x7a\x3f\xec\x57\x98\x7e\x41\xa1\x7f\x54\x15\x00\x40\x5a\xca\x1b\x2f\x14\xe8\xa1\xdb\xa3\xe5\xda\x0e\x71\x90\x1a\xf0\x49\x3a\x4f\x7d\xb3\xcd\x13\x3e\x7b\x90\x0e\x86\x3e\x76\xd2\xa4\xb6\x2c\x3d\x42\xed\x06\x8b\x85\x93\x02\xb6\x1b\xfa\x5e\x8d\x19\xc3\x79\x31\x3a\x22\xba\x81\xdb\x99\x4a\x23\x00\x36\xc2\x63\x1a\xc5\xff\xc9\x9d\x93\xb0\x01\xe6\xdf\x8c\x72\x0f\xff\xf9\x24\x9f\xfe\xf4\xff\x13\x1f\xd8\xde\xcf\x5a\x7a\x29\x94\xfc\x82\xcd\x0c\x22\x79\x89\x27\x4c\x9c\x2d\x1d\x60\x27\x3d\xb5\xc3\x99\x52\x4b\x86\x96\xa0\x39\xa8\x2d\x0a\xbf\x80\x21\x4b\x02\xc4\xc5\x72\x37\x32\x7c\x9e\xdb\x77\xbb\x34\xf0\xb7\x58\x6b\xfa\x9b\xcd\x0b\xf9\x20\x0a\x4c\xf5\xaa\x43\x95\x8a\x50\x69\x2f\x1a\x9a\x97\xbd\x11\x1d\x6d\x2c\xc9\xbe\x0d\x43\xdc\xc3\xc7\xa6\xb1\xe8\xdc\x5f\x2e\xec\xfd\x7b\xa8\xf3\xef\x08\x67\xb1\xb7\x49\x18\x54\x8b\xe6\x2a\x7b\xf3\xb2\x17\xf6\x7a\xb3\x6a\x6d\x22\xaf\x55\x33\x17\x6d\x84\xc4\x7c\x75\xa4\x79\x8b\xbf\x0f\xd2\x72\xf1\x3a\x68\x8d\xcd\xd1\x25\x66\x4c\x20\x0b\x52\x28\xf5\xce\x24\x35\xf6\xa5\x35\xa6\x2d\xd2\x18\x74\xa0\x4d\x5e\x70\xbe\x96\xd1\xb0\xdb\xa7\xbd\xf6\x88\x16\x37\x79\xee\x64\x6b\x53\x28\x68\x2b\x31\x7d\xac\xd0\xde\x38\x27\xe3\x6e\x62\xda\x50\xa4\x64\x44\xdc\x51\xfa\x18\x06\x57\x4c\x27\x8f\x1b\xc3\x76\x68\xac\xd1\x39\x61\xa5\x1a\xa3\x40\x61\x82\x33\x67\x0d\xd1\x92\xed\x45\x56\x2e\x55\x40\xd9\x28\x22\x85\xa4\xa5\x32\x8f\xba\x61\x1f\x89\x69\x19\x38\x56\x27\x89\x1b\x67\x93\xc3\xd6\xe0\x07\x4b\x45\x13\xb9\x33\x6f\x70\x16\x3b\x73\xc2\x26\x6f\x74\x93\x89\x33\x90\x87\x89\x84\x78\xc7\xe4\x82\xce\x81\xc2\x13\x2a\x2a\xd0\x7e\xd8\x2b\x59\x6f\x60\x3f\x50\xd1\x4a\x47\xcf\x28\x2e\x82\xe2\xb6\x57\xd8\xcd\xc0\x52\x16\x58\x19\x14\x69\x45\x92\x8c\xd3\xce\x76\xe5\xe0\xcc\x85\xdb\x0c\xa8\x66\xfd\xc7\xec\xa0\x46\xde\x42\xc2\xea\xc9\xd2\x97\xfd\x09\xab\x76\x62\x84\x83\x15\xda\x47\x59\x17\xd7\xc9\x3e\xd2\x8e\x9e\x6a\x81\xdc\x91\xa7\xc4\xa2\xc5\x8a\x3e\xcb\x90\xa8\xf1\xcd\xd9\x25\xb5\x5b\xcf\xe4\x22\x75\x29\xe3\xce\x10\xb8\xfe\x52\xee\xb3\xeb\xfe\x68\xcd\x70\xa0\xad\x39\x0b\xac\x6b\x1d\x0a\x5a\x89\xbd\xa2\xa0\xbc\xe2\x13\x27\xef\x1a\x97\x08\x6b\xe1\xc7\xcc\xf6\x19\xc6\xb7\xfb\x41\x5d\xd1\x0e\x3a\x97\xfb\x82\xa2\x6e\xef\xe1\xaf\xa1\x7c\xff\xc8\x53\x78\x9a\x71\xcb\x47\x01\x19\x76\x16\x5d\x3c\x62\xb4\xd1\xea\x50\x5c\xd4\x0d\x70\x12\x6a\xc0\x8b\x69\x61\xca\x36\xb6\x2d\xfc\xfc\x33\x44\x2b\x2e\x46\xd2\xdf\x9b\xc4\xff\x42\xc5\x71\xd0\x0d\xce\x93\x2c\xa4\x95\x9c\xe8\x10\x44\x08\x52\x42\x8c\xf2\xb6\xec\x35\xec\xd3\x9b\x19\xfc\xd7\x6a\xfe\xe9\x6b\xe1\xe3\x74\xaa\xf8\x71\x3e\x8e\xbb\xc7\x0a\x1d\xf3\x6e\x72\x25\x1d\xff\x86\x89\x04\xa5\xae\xd5\xd0\x20\x49\xc9\x74\x34\x09\x66\xd4\x47\xac\x1f\xe7\x41\x88\x14\x90\x51\xce\xc8\x07\x5b\xca\x10\x49\xfc\x6b\x14\x7e\x08\x43\x50\xf8\xd5\x94\x11\x1a\x93\x06\xad\xcb\xf9\x0d\x28\xf9\x48\xa7\x51\x25\x59\x45\x75\x24\x8f\x84\x6e\x8a\x80\x62\x9d\x4b\x2f\x48\x34\xc9\x96\x8b\xd6\x43\xaf\xc2\x61\x04\x5e\x27\xf2\x94\xa4\x25\x91\x27\x71\xeb\xc5\x23\x16\x36\x26\x86\x8e\x6f\x1c\x6d\x4d\xeb\xe1\x2f\xfd\x34\xf6\xf8\x62\xff\x44\xac\x9b\xa0\x40\x42\xcf\xdc\x2e\xeb\x28\x9e\x46\xaf\x29\x23\x12\x6f\x42\xea\x90\x8f\xb2\xb5\xf2\x39\x0e\xa6\xc7\xee\x0c\x42\x1e\x4d\x8a\x4f\xf8\x20\x5d\x34\x9e\xc3\xc0\x20\x5f\xa2\x10\xdc\x4c\x2b\x23\x43\xd0\x26\x52\x44\x20\xd4\xc6\x5a\xac\xbd\x1a\xaf\x8a\x7f\x74\x6e\x19\xfe\x22\xc7\x27\xcd\x28\xe0\xb4\xdc\x33\x67\x11\x25\x81\x1c\x87\xcf\xc5\x31\xfd\x91\x89\x37\x8b\xb7\xb7\xd7\xf1\x93\x43\xd5\x4e\x69\x26\xa1\xac\xf3\x4c\xf2\x28\xb1\xcb\x34\x36\xa9\x5a\xc2\xa3\x04\x74\x35\xa3\x5c\x8a\xc6\x14\xab\x09\x85\x2f\xcb\xa0\xdc\x27\x78\xf3\xdc\x11\xf4\x85\x54\xf1\x9a\xf7\x59\xf0\x6c\x72\xc7\x6c\xd6\x73\xc7\xe6\x84\x1b\x11\x91\xae\x35\x98\x67\x6a\xcb\xe7\xbf\xb1\x67\xa9\x20\xd6\x4e\x67\x1d\x0a\x3d\xa1\x89\x08\x88\x27\xb4\xe3\x73\x07\xbf\xc5\xb5\x81\x7b\xe9\xa2\x6c\x0a\xca\xd9\x69\xb0\x95\x1a\xa7\xe6\x2d\x6f\xba\x72\x3c\x5b\x63\xbb\xbc\x2d\x3d\x73\x79\x34\xc5\x9f\xdf\x23\x4d\xef\x0a\x02\x87\xf0\x8d\x91\x8b\x8a\x29\x12\x7e\x93\x2e\x5c\x68\x48\xb9\x74\x79\xbd\x31\xc8\xa6\x1f\x68\x8d\x08\x5b\xae\x43\x42\x96\x62\x88\xc2\xdd\x56\xd1\x6f\xf2\xcb\x4c\x3e\xcc\x64\x47\x6f\x25\x05\x26\x69\xc3\x45\x9d\x5f\x32\x50\x80\x78\xb9\x47\x5f\x15\xd8\xbb\xb0\x9d\xef\x8a\xc4\xe6\x05\xde\xb9\x19\x53\xc1\xaa\xc8\xce\x3c\x57\xb6\x9e\x04\x8c\xcd\xda\xfc\x1f\x96\x40\x16\x5f\x63\x98\x3f\xbf\x22\x64\x3e\x06\xf5\x52\x64\x49\x62\x1a\x15\x54\x9e\xd0\x60\x2c\xe0\xef\x83\x50\xe1\xdb\x8a\xa6\x79\x51\xc9\xc0\x8b\x52\x8d\xce\x03\x1c\x27\x52\xcd\x42\x95\xeb\x9f\xdd\x1e\x5b\x63\x71\xc7\xd2\x00\x7d\xac\x4a\x35\xe4\x45\x17\x1b\xd2\x1a\x78\xbc\x2d\xd9\xe3\x41\x6a\x4d\x65\xb4\xb8\x14\x2d\xd7\xa5\x2b\xb3\x5f\x27\x6e\x36\xf0\x66\xfa\xf8\x16\x3e\xbc\x1c\xed\x7f\xe4\x0a\xd9\x2f\x88\x9d\xef\xc0\xa2\xe6\x28\x91\xed\x2d\x9e\xf8\x86\x32\x0d\x17\x41\xa2\x5c\x2f\x23\xaf\xd4\x21\xa2\x69\x48\x83\x94\x85\x22\x39\xcd\x32\x2d\x57\xce\x99\xd7\xa9\x90\x45\xf2\xef\xee\xe0\xa3\x73\x68\x7d\xb9\xd2\x9a\x73\x7a\x74\xbf\x5c\x74\x30\x21\x91\x38\x48\xf2\x7a\x89\x17\xd5\xf6\xa9\x5c\x40\xcb\x70\xee\xe9\x7d\x22\x90\x88\x76\x45\x07\x91\xed\x5b\xe9\x3e\xc7\xdf\x18\x42\x8e\x0f\xe8\x1f\xc6\x1e\x6f\x6e\x6f\xef\x61\x3d\xbb\x7f\x13\x9a\x04\x71\x8a\x32\xb3\x5c\x6d\xba\x5e\x78\xde\x6b\xc2\x8f\x35\xe4\xdf\x77\xf4\xca\x55\xd5\xf7\x7f\xe9\x31\x3b\x90\x1e\x7f\x57\x2d\xba\xa1\x7b\xb5\x08\x4b\x7a\xbe\xed\x2c\x13\x04\xe0\xaf\x5d\xef\xc7\x58\x81\xf1\x9c\xa9\xc7\xf8\x83\x83\x89\x63\x66\xa4\xca\x59\x3d\x0a\x2a\xdc\x2f\x68\xcd\x52\x3a\x56\xd3\x2a\x5c\x2e\x71\xb3\x46\xa1\x2b\xa1\xbe\x3c\x06\xfe\xb4\xfd\xe9\x1e\xde\xd0\x96\xa6\xf1\xac\xc6\xa4\x5e\xa3\x4d\x1c\x32\xfe\x4d\x6a\x6a\xd2\x9b\x0b\xdf\xbf\x56\xff\x0b\x00\x00\xff\xff\xe7\x1a\xc5\x29\x66\x1c\x00\x00"

func fungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_fungibletokenCdc,
		"FungibleToken.cdc",
	)
}

func fungibletokenCdc() (*asset, error) {
	bytes, err := fungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "FungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x25, 0x9c, 0x1d, 0xaf, 0x56, 0xca, 0x66, 0xdd, 0xbe, 0x5, 0x14, 0x40, 0xee, 0xae, 0xd1, 0xf3, 0x63, 0x1d, 0x6a, 0x32, 0x37, 0x36, 0x8a, 0x96, 0xd1, 0x8, 0x7c, 0x53, 0x4, 0xab, 0xf0, 0xbb}}
	return a, nil
}

var _utilitycontractsPrivatereceiverforwarderCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\xc1\x8e\xdb\x36\x10\xbd\xeb\x2b\x26\x2e\xd0\x4a\xc1\x46\xbe\x14\x3d\x18\xeb\x6c\x83\x6d\xf7\x58\x2c\x92\xb4\xd7\x62\x44\x8d\x2d\x36\x32\x29\x90\x23\xab\x8b\x85\xff\x3d\x20\x25\x51\xa2\x64\x07\xd8\xbd\xac\x28\x71\x66\xde\x7b\x7c\x33\xf4\xf6\x7d\x92\xfc\x04\x4f\xad\x3a\xca\xa2\x26\xf8\xaa\xbf\x91\x82\x67\x23\xcf\xc8\x04\x9f\x49\x90\x3c\x93\x81\x47\xad\xd8\xa0\xe0\x24\xf9\x5a\x49\x0b\x62\x58\x82\x3c\x35\x35\x9d\x48\xb1\x05\x04\xdb\x90\x90\x58\x83\x21\xab\x5b\x23\x08\x50\x95\x60\xc6\x14\x52\x31\x99\x03\x0a\x82\xa4\xab\xb4\x25\x28\xa9\xd1\x56\x32\x1c\x5a\x25\x58\x6a\x05\xd2\x82\x56\xf5\x0b\x08\xac\x6b\x74\x60\x8a\x17\x40\x05\x58\x9e\xa4\x02\xae\x8c\x6e\x8f\x15\x20\x34\x6d\x51\x4b\x01\x02\x1b\x2c\x64\x2d\xf9\x25\x4f\x92\xf7\xdb\x24\x91\xa7\x46\x1b\x0e\x54\x7a\x26\x07\xa3\x4f\xb0\xc9\xb7\x79\xbe\x8d\x3e\xe4\xa2\x14\x9b\x24\x69\xda\x62\x22\x33\xb0\x1e\x49\x3f\x69\xd3\xa1\x29\xc9\xc0\x6b\x92\x00\x00\x6c\xb7\xf0\xe7\x99\x14\x03\x57\xc8\x0e\x2d\x9d\x24\x33\x95\xd0\x55\xa4\x80\x5d\x5a\x0b\x68\x02\x33\x2a\x81\x35\x70\x45\xc0\x68\x8e\xc4\x41\x0b\x9f\xcd\x95\x26\x9f\x6e\xa8\xfb\x47\x1f\x95\xe2\x49\xb7\x8a\x77\xf0\xf7\x93\xfc\xff\xb7\x5f\xef\x80\xf5\x0e\x3e\x95\xa5\x21\x6b\x1f\xb2\x24\xc4\xd6\xc4\xf0\x85\x54\x49\xe6\x0b\x6b\x83\x47\x7a\x46\xae\x76\x30\x5b\xc4\x7b\x17\xec\x6e\x06\xfd\x20\xe6\xd9\x2b\xdf\x87\x4c\xcf\x53\x99\x70\xf0\x2b\xe9\x06\xf9\xbc\x79\xa4\x75\x82\x19\xf2\xca\xcc\xa5\xf2\xfa\x75\xb2\xae\xa1\x20\xb0\xa4\x38\x8f\x63\x09\xf8\xa5\x21\x90\xaa\x94\x02\x99\xec\x70\x0e\xfe\x28\x10\x0c\x1d\xc8\x90\x12\xe4\x44\xc7\x58\xeb\x3e\x45\x78\x44\x21\xc8\xda\xd4\x52\x7d\xc8\xe0\x8c\xc6\x6d\x96\x8d\x24\xa7\xfa\x63\xb0\xd5\xfd\xcf\xaf\xb1\x65\x46\x19\x2e\x1f\x23\x52\x03\x85\x6b\x85\xb6\x5b\x67\xc7\xde\xdd\x1e\x2c\xe3\x37\x72\x60\xff\xc1\xb6\x66\xd0\xc5\x7f\x24\x18\xd0\x7a\x9b\x9b\x63\xeb\x3a\xc9\x77\xcd\xa1\x17\xd0\xce\x33\x49\x1e\xed\x14\xe0\xfe\x62\x87\x4c\xad\x95\xea\xe8\xbf\x59\xd6\x86\xca\x49\x8d\x1f\xf0\x1f\x8d\x9f\xb9\x16\x1c\x69\xa4\xae\x63\x76\xf0\x7b\x4c\xdd\x57\xc9\xe0\x35\xa4\x70\x7f\xf5\xcc\xd2\x9f\xe9\x00\x7b\x70\x8a\xe6\x01\x5d\x5e\x68\x63\x74\x97\x66\xef\x92\x55\x5c\x81\x35\xba\xb3\xda\xfb\x0e\xcd\x87\x65\xbc\x6f\x96\x3b\x8f\xd1\xdd\x7f\x70\xff\xb3\x78\xbb\xeb\xc6\x5b\xbd\x34\xe4\xef\x9b\xc9\xa3\xd4\x9d\x22\xf3\x90\x63\xdf\x58\x59\xc8\x74\x99\x92\x4a\x25\x39\x7d\xab\x35\x96\x22\x35\x86\x16\x6f\x06\x6a\x0b\x8d\xe0\xdd\x1e\x94\xac\x77\xb0\x79\xd4\x6d\x5d\x82\xd2\x0c\xfd\xb7\x69\x0a\x4f\x16\xf7\x63\xcd\x1d\xf7\x84\x69\x13\x15\xb9\x44\xab\xf8\x5c\x60\x3f\xd5\x4f\xe2\x80\x4b\x98\x74\xc2\x10\x32\xfd\x45\xdd\xd4\xcb\xfd\x2b\x67\x5f\x45\xdd\xac\xc7\x27\x58\x9d\xe4\xca\xc3\x6a\x8c\x3e\xcb\xd2\xfb\x70\x5e\x68\xf0\xa0\x9b\x15\xce\x72\xeb\x1a\x6f\x97\xdb\x59\x75\x36\x6d\x26\x81\xb9\x35\x0a\xee\x3f\xf4\x35\xe0\x6a\x85\xf0\x98\x8d\xe4\xd7\xa3\xac\x1f\xb1\xb3\xcc\x23\x78\x4b\xaa\x1c\xdc\xe6\x41\xd9\xf4\x5f\x18\xdc\x14\xe6\xf5\xdd\x30\xd5\x6e\xf7\xd3\xaa\x31\x50\x08\x67\x59\xd8\xc3\x91\xf8\x53\xbf\x48\x83\x4b\x57\xdb\x9b\x78\x42\xc3\x7e\x4c\x90\x1f\x89\xe7\x0a\xde\xba\xdc\xf2\xf0\xf4\x31\xbd\xb9\xe7\xe6\x3d\x90\xad\x9c\x3d\x19\xfa\xe1\x01\x1a\x54\x52\xa4\x6b\x47\x47\xb3\x7a\xa0\x30\xce\x3c\x32\x9b\x05\xcf\x05\xc7\xd5\x2c\xe8\x35\x8e\xa1\x5c\xf7\xb5\xef\x68\xeb\x4f\x74\x75\xf1\xdd\xf9\xd1\x79\xed\x4a\xbc\x1b\x7e\x72\x2c\x2f\xbe\xe8\xfc\x7c\x8b\xad\xee\x63\x3f\x13\xc7\x72\x8b\xcd\xb7\x2f\x64\x17\xb5\xb8\x91\x6f\x45\x4d\x68\x60\x3f\x83\xb9\x28\x35\x7a\xc2\xe2\x99\xd2\xd0\x13\x3d\xda\x34\x9b\x4d\xc5\x15\x81\xe1\x28\x2e\xc9\xe5\x7b\x00\x00\x00\xff\xff\x76\x9e\xa6\x51\x29\x0a\x00\x00"

func utilitycontractsPrivatereceiverforwarderCdcBytes() ([]byte, error) {
	return bindataRead(
		_utilitycontractsPrivatereceiverforwarderCdc,
		"utilityContracts/PrivateReceiverForwarder.cdc",
	)
}

func utilitycontractsPrivatereceiverforwarderCdc() (*asset, error) {
	bytes, err := utilitycontractsPrivatereceiverforwarderCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "utilityContracts/PrivateReceiverForwarder.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x87, 0xce, 0x6b, 0x90, 0x12, 0x2c, 0xaf, 0xe4, 0x8c, 0xfe, 0x69, 0x5c, 0x8e, 0x67, 0x97, 0x7c, 0xc2, 0xf4, 0x5d, 0x33, 0x3c, 0xd2, 0x77, 0x98, 0x6b, 0xd, 0xbf, 0xe2, 0xef, 0x17, 0x95, 0x7c}}
	return a, nil
}

var _utilitycontractsTokenforwardingCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\xc1\x6e\xe3\x36\x14\xbc\xf3\x2b\x66\x53\xa0\x6b\x07\x59\xf9\x52\xf4\x10\x24\xdd\x16\x69\x73\xec\x21\xd8\xb6\xc7\x82\x26\x9f\x2c\x6e\x64\x52\x20\x9f\xac\x06\x81\xff\xbd\x20\x45\xd1\x52\xe2\x14\xe9\x65\x7d\xb1\x25\xf3\xcd\x9b\x79\x33\x24\x37\x97\x97\x42\x7c\x87\xfb\xde\xee\xcc\xb6\x25\x7c\x71\x8f\x64\x71\xef\xfc\x20\xbd\x36\x76\x87\x3b\x67\xd9\x4b\xc5\x42\x7c\x69\x4c\x80\xca\x8f\x08\x8d\x1b\x02\x1a\x37\x40\x5a\x48\xa5\x5c\x6f\x19\xca\xf5\xad\x46\x20\x46\xdf\x41\x42\xf5\x81\xdd\xbe\x80\x8f\xd8\x0f\xa4\xc8\x1c\xc8\x0b\x76\x90\x6d\xeb\x06\x70\x43\x7b\xb0\x43\x3d\x76\x05\xc7\x75\x21\xbe\x91\xd0\xa6\xae\xc9\x93\xe5\xd2\x63\x68\xc8\xd2\x81\x7c\x2c\x7b\x82\x1f\xd1\x72\x4d\x15\x59\xd2\x13\x94\xb4\xe8\xfa\x6d\x6b\x42\x03\x8e\xb4\xb3\x20\xf2\xf0\x14\x5c\xef\x15\x41\x06\xc8\x42\x06\x4a\x76\x72\x6b\x5a\xc3\x4f\xf8\xda\x07\x46\x6b\x1e\x09\x12\x7f\xca\xbe\xe5\x2b\x21\xad\x8e\xed\x10\xc8\x46\x0c\xed\x28\xd8\x8f\x0c\x3a\x90\x85\x25\x8a\x94\xf1\x68\xdd\x00\xc3\x30\xe1\x44\xba\x12\xe2\xaf\x86\xec\x7c\x44\x83\xb4\x9c\xb4\x29\x4f\x92\x63\x8f\xc2\xed\x6a\x94\xa4\x64\xdb\xa6\x6e\xe3\x8a\xdf\x69\x28\x2b\x44\xdd\x5b\xc5\xc6\x45\x44\x8d\xce\xbb\x83\xd1\x14\x9b\x0e\x86\x9b\x54\x53\x04\x79\x4a\x14\x14\x81\x1b\xc9\x23\x72\xec\x3d\x1b\xb4\xe0\x86\x8c\x3f\x8d\xbb\x12\xe2\x72\x23\x84\xd9\x77\xce\xf3\x0b\xd7\x6a\xef\xf6\xb8\xa8\x36\x55\xb5\x59\xfc\x51\x29\xad\x2e\x84\xe8\xfa\xed\x29\x1a\xe9\x8f\x59\x84\x9e\x85\x00\x80\xcd\x06\xbf\x1d\xa2\x93\x89\x90\x09\xa0\xbd\x61\x26\x9d\x1c\x9d\x58\x48\x4f\xd0\xd4\xb9\x60\x78\x1c\x6b\x14\xc5\xd2\xef\x88\x27\xaf\x7d\x42\x8b\x1d\x29\xc1\x4d\xd3\xd1\xbf\x8e\x75\x2b\xb9\x8f\x93\xbe\xc6\x1f\xf7\xe6\x9f\x1f\x7f\xb8\x4a\xdc\xaf\xf1\x8b\xd6\x9e\x42\xf8\xbc\x16\xa5\xbe\x64\xa1\x0c\xf8\x7a\x29\xbb\x2a\xe3\xcc\x1a\xb2\x8e\xb4\x15\x4c\x88\xcc\x3d\x25\x8a\x73\xce\x49\xc8\x60\xda\x16\xdb\x14\x19\xae\x96\xb5\x04\x7e\xea\x08\xc6\x6a\xa3\x24\x53\xc8\x03\x49\x33\x91\x73\xe3\x5c\x7a\x9c\x89\x1e\x21\xca\x4f\xa9\x14\x85\xb0\x0a\xd4\xd6\x6b\x1c\x64\x34\x5d\x99\xce\x50\x14\x7f\x57\x02\xbd\x60\x9e\x79\x9e\x43\xdb\x6c\xa2\xf8\x31\x5e\x63\x66\xe4\x23\x85\x69\x13\xc0\x6d\xbf\x92\xe2\xb4\x6d\x2c\xa4\xdf\xf5\xfb\xb4\x2b\xad\x9e\xe2\x14\xe6\x48\x86\x27\xf3\x0a\xa7\x8f\x21\x23\xf5\x21\xa6\x22\xed\x27\x76\x9e\xf4\x49\xf2\x39\x5a\xd1\xa8\xba\xb7\x13\xf3\xd5\xe8\xe6\xcf\x4b\x9f\x12\xf0\x1a\xcf\xa5\x2a\x7e\xda\x59\x66\x1e\xa8\xc6\x2d\xe2\xa4\xaa\x42\xa8\xda\x3a\xef\xdd\x70\xf3\xfd\xf3\x79\xd3\x8f\x3f\xad\xd6\x1f\xc4\x2b\xc8\xad\x6c\x65\xb4\xe7\x36\x05\xab\xca\x8f\xcb\x75\xb3\xb6\xd5\x92\xf8\xcd\xa7\xf8\xbd\x5e\x2e\x8f\x3b\xe1\xed\x1c\xe7\x0e\x53\x90\x93\x08\x37\x58\xf2\x9f\x2b\x39\x86\x7a\x5d\xd0\x8e\x0b\xb7\x55\x23\xed\x8e\x1e\x26\xc1\xf9\x39\x2c\x7d\x81\xab\xd3\x8b\xba\x9c\x91\xd9\xb9\x7c\xbe\xe8\xd3\xd2\xff\xf2\xe7\x45\xaf\xd5\xdf\xb0\x34\x3c\x9c\x0b\xe4\x4b\x9f\x3a\x4f\x2f\xde\xc4\xcf\xbc\xfa\x3d\x4e\xe1\xc3\x2d\xac\x69\xaf\x71\x71\x97\x6e\x21\xeb\x18\x63\xd9\xb9\x43\x31\x9d\x67\x51\xe4\x89\xd6\xc5\x82\xc2\x71\xf1\xb4\x0c\x0e\x6e\x17\xec\xce\x0d\xdf\x58\xc3\xab\xb3\xdb\xf1\x7d\xea\xff\x57\x48\xbf\xad\xf4\xd7\x69\x18\x0b\x8e\xe5\x98\x7f\x7d\x71\xe5\x57\xf1\x34\xb1\x34\x2c\xae\xe3\x89\x56\xb9\xc2\xde\x88\x5d\x8e\x5c\x89\xdb\xab\x1e\x6f\x8c\x3b\x9e\x15\xa5\xdd\x69\xd0\x9e\xb8\xf7\x16\x37\x9f\xf2\x3d\x7c\x16\xa6\xfc\x5c\x67\x85\x47\xf1\x6f\x00\x00\x00\xff\xff\x91\x6b\x10\x08\x31\x09\x00\x00"

func utilitycontractsTokenforwardingCdcBytes() ([]byte, error) {
	return bindataRead(
		_utilitycontractsTokenforwardingCdc,
		"utilityContracts/TokenForwarding.cdc",
	)
}

func utilitycontractsTokenforwardingCdc() (*asset, error) {
	bytes, err := utilitycontractsTokenforwardingCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "utilityContracts/TokenForwarding.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1f, 0x32, 0x22, 0x16, 0x73, 0xff, 0x5a, 0xd6, 0xef, 0x6d, 0xa2, 0x31, 0x18, 0xf, 0xab, 0x51, 0xc0, 0xbf, 0xd4, 0xf6, 0x9e, 0xb8, 0xfb, 0x33, 0xe3, 0x24, 0xac, 0xb, 0xe5, 0xfb, 0xa4, 0x33}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"ExampleToken.cdc":  exampletokenCdc,
	"FungibleToken.cdc": fungibletokenCdc,
	"utilityContracts/PrivateReceiverForwarder.cdc": utilitycontractsPrivatereceiverforwarderCdc,
	"utilityContracts/TokenForwarding.cdc":          utilitycontractsTokenforwardingCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"ExampleToken.cdc": {exampletokenCdc, map[string]*bintree{}},
	"FungibleToken.cdc": {fungibletokenCdc, map[string]*bintree{}},
	"utilityContracts": {nil, map[string]*bintree{
		"PrivateReceiverForwarder.cdc": {utilitycontractsPrivatereceiverforwarderCdc, map[string]*bintree{}},
		"TokenForwarding.cdc": {utilitycontractsTokenforwardingCdc, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
