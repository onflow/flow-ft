// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../src/contracts/ExampleToken.cdc (7.12kB)
// ../src/contracts/FlowToken.cdc (7.092kB)
// ../src/contracts/FungibleToken.cdc (7.306kB)
// ../src/contracts/TokenForwarding.cdc (1.972kB)

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

var _exampletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\xdf\x73\x1b\xb7\x11\x7e\xd7\x5f\xb1\xcd\x43\x4b\x4d\x24\x8a\x9d\xe9\xf4\x41\x23\x27\x96\x6b\xb9\xe3\x69\xeb\x76\x6c\xa7\x7d\x15\xee\x6e\x49\xa2\xbe\x03\x38\x00\x8e\x14\xe3\xd1\xff\x9e\xd9\xc5\x8f\x03\x8e\x47\x4a\x8a\x13\x3f\x24\xe6\xdd\x61\x77\xb1\xfb\xed\xb7\x1f\x60\xd9\x6d\xb4\x71\xf0\xae\x57\x2b\x59\xb5\xf8\x59\x7f\x41\x05\x4b\xa3\x3b\x58\x3c\xbc\xfb\xe9\xc3\xdf\xdf\xbf\xf9\xe7\xdd\xe7\x7f\xff\xe3\xee\xc3\xed\xdb\xb7\x1f\xef\x3e\x7d\x3a\x3b\xdb\xf4\x15\xd4\x5a\x39\x23\x6a\x07\x77\x0f\xa2\xdb\x84\x65\xd7\x23\x2b\x5f\xcf\xce\x00\x00\xae\xae\xe0\xb3\x76\xa2\x05\xdb\x6f\x36\xed\x1e\xf4\xb2\x58\x65\x41\x2a\xc0\x07\x69\x1d\xaa\x1a\x79\x05\x79\xd8\x0a\x03\x8e\x96\x7d\xe2\x55\xd7\xf0\xd3\x3b\xf9\xf0\xd7\xbf\x24\x93\x77\x5b\x54\x0e\xdc\x5a\x38\x90\x16\xb0\x93\xce\x61\x03\xbb\x35\x2a\x70\x6b\x1c\x02\x94\x16\x6a\x83\xc2\x61\x93\x4c\x23\x2f\xf5\xce\xdf\x2b\xe9\xa4\x68\xe5\xcf\xd8\xcc\xa4\xff\x7b\xe9\xf0\xfc\x59\x1e\xfd\x46\x84\x41\xd8\x49\xb7\x6e\x8c\xd8\x85\x24\x0a\xf8\xaf\xe8\x5b\x37\xe9\xfb\x7f\xf1\xd3\x99\xe8\x74\xaf\x5c\x74\x79\xc1\x4b\xaf\xe1\xb6\x69\x0c\x5a\xfb\xe3\x4b\x43\x68\x70\xa3\xad\xa4\x37\x4e\x9f\x0c\xe0\x6d\xfc\xf0\x20\x00\xa7\x5f\xe8\x5e\xe1\x2e\x0f\xa1\x93\xea\x58\xc6\xff\xc5\xaf\x46\x1e\x5f\xbe\x45\xeb\x8c\xde\x1f\x71\xf1\xa6\x37\xea\xd7\xb9\x10\xbc\x11\x8e\xde\x80\x41\xab\x7b\x53\xe3\x71\x0c\xf1\x5e\xcc\xdf\xfc\xbb\x99\x68\x5b\xbd\xc3\xe6\xf6\xd7\xba\xad\x28\xec\xe7\xb8\xe5\xfd\x25\xb7\x83\x87\xa1\xd2\x57\x57\xc9\xab\xa8\xd7\xd0\x5b\x34\x60\x9d\x36\x68\x41\x28\x90\xca\x3a\xa1\x6a\xa4\x56\xd4\xaa\xdd\x73\xc7\xf0\x62\xea\x45\xb7\x46\xe9\xbf\x16\x2b\x4c\x1d\xbc\x46\x58\xf6\xaa\x76\x52\xfb\x8e\x1d\x96\x08\xd5\xc0\x4a\x6f\x91\x72\x0e\x95\x37\xb6\x31\xc8\xcf\x37\xda\x3a\xea\xc5\x46\xf2\xc2\x68\x4d\xaa\x11\x55\xc4\xbe\xdd\x73\x75\x6b\xd1\xb6\xd8\xcc\x73\xdf\xf5\x1a\xeb\x2f\x16\xd6\x62\xb3\xa1\x7c\x39\x30\xbd\x72\xb2\x43\x5e\x89\x5b\x34\x20\x52\x7c\x9c\xb8\xc2\x44\xb4\xf4\x31\xa4\x96\xde\x2b\xbf\xf5\x0a\x63\x92\xe3\xae\x88\x3a\xf0\xc1\x51\x72\x0a\x26\xe1\xca\x51\x8c\xd1\x9a\x07\xe2\x52\x2a\x5e\x7b\x01\x56\xd3\x6b\xc3\x85\x53\x1a\x76\x62\x0f\x4b\x4d\x81\x75\xa2\x95\xb5\xd4\xbd\xf5\x85\x70\x3a\xb8\xf4\x09\x4c\x59\xd1\x7d\x70\x2a\x15\x08\x69\xe6\x70\x0b\x76\x83\xb5\x14\x6d\x00\xda\x00\x0d\x85\xd8\x58\x32\x54\x0d\x21\x38\xcd\xc0\x8d\xd6\x86\x8e\x2c\xb2\x40\x28\x4a\x66\xd8\xff\x88\xb6\xe7\xff\x31\x7a\x2b\x1b\x34\x17\xa3\xe7\x1f\xb1\x46\xb9\x3d\x7c\xfe\x46\xb4\x0c\xa6\x40\xf7\xc1\xfd\x5a\xb7\x14\xe1\x1a\xa1\x0a\xef\xf5\x12\x04\x27\xc0\x86\xb8\xd2\xe7\x91\xef\xc3\x97\x25\xd7\x27\xc8\x44\xa2\x2e\x8c\x12\x12\xe2\x6e\x38\xa9\x54\x7f\x02\x46\x5a\x4b\x0b\x67\x23\xcb\xe7\xf0\x35\xbd\xa7\x3f\x16\xdb\xe5\x3c\x9a\x7c\x15\x8d\xa7\x4f\x1e\x8b\x48\x22\xc3\x67\xcf\xf2\xd7\xef\x22\x0a\x3d\x5e\xc4\x97\xd8\x73\x0e\x57\x04\x53\x66\x07\x10\xfc\x50\x98\x55\xdf\x61\xa8\x59\x44\x95\x6a\x92\x0b\xeb\x8d\x84\x35\x3c\x50\x52\xdf\xcd\xf3\x45\xef\x5d\x80\x94\x0d\x5c\xe2\x90\xe6\xba\x30\xfb\xd0\xa4\x91\x76\x7a\xeb\x91\x42\xe5\xc9\x0d\x90\xd9\x4e\x2b\xdc\xa7\x2f\x2b\x94\x6a\x05\xce\x08\x65\x97\x68\x0c\x36\x73\xf2\x62\xd0\xf5\x46\xf9\xc2\x2a\xdc\xb5\xfb\xdc\x48\x6c\xa4\xe0\x52\x17\xed\xc4\x76\x7d\x5b\x52\xa7\x48\xc7\x3d\x58\x65\xe3\x2a\x37\x85\xad\xc5\x1d\x35\xd3\x7c\x2a\xcd\x04\x98\x65\xaf\x52\x9e\xc6\x54\x7f\x0d\xaf\x4b\x8c\xfa\x88\x4e\x16\xbd\xf8\x79\x19\x72\x5e\x2c\x20\xca\x3e\x3a\xba\xfd\xff\xe3\xe8\x66\x63\x7a\xa7\xd0\xfc\x38\x17\x7e\x8e\x9e\x17\xb6\x7c\x1e\xe1\xe6\x32\x67\x82\x01\xa6\xde\xda\xf9\x11\x04\x86\x8c\xbd\x04\x80\xa1\x26\xba\xfa\x3f\xd6\x63\xf4\x31\xe4\x44\xd3\xd8\xa2\xdf\x9c\x4d\x4d\x16\x2a\x99\x35\x32\xfd\xe4\xed\xd9\x69\x30\x4a\x0b\x61\x1a\xd2\xe2\x30\xad\x79\x95\x25\x87\x3e\x98\x0a\x6b\xd1\x5b\x1c\x20\x5d\x74\x19\xc5\x98\xc1\x98\x00\x8b\x26\xfa\x0e\xc4\xc6\xa3\x81\x97\xfe\x69\x88\x76\x2d\x8a\x8d\x54\x88\x8a\x40\x68\xfb\x0e\x1b\xde\x2a\x93\xf4\x52\xf3\xa0\x09\x08\x0c\x6a\x62\x7e\x80\xb0\x90\xea\x99\x2f\xeb\x14\xaa\xc6\x5c\xd2\xa2\x83\x2d\xef\xef\xe6\x32\x68\x40\xfb\x07\x78\x9d\x4b\xde\x79\xb9\xdb\xa7\xc0\xf8\xbd\xb7\x37\x1f\xd3\xd2\x08\x93\x87\x6a\xae\x58\xe6\x45\xdd\x93\xc0\x2c\xd6\xc0\x2b\x58\xcc\x17\xc5\xfb\x58\xcb\x6d\xb1\x85\x0c\x9f\xe1\x83\xd9\x38\x2f\x45\x02\x32\x65\x0f\xaf\x8e\xbf\xba\x2c\x12\x91\x79\xcb\x7c\x26\xde\xb9\xeb\x36\x6e\x3f\x25\x81\xca\x86\x28\x79\xd2\x23\x91\x78\x04\x44\x0e\xf0\x9f\xd1\xe8\x34\xe7\x55\x93\x78\x4f\x0e\xbc\x26\xda\x96\x18\x32\xf0\x1b\x4d\x6b\x9e\xee\x5d\x6f\x3d\xcf\xd1\x20\xb7\x49\x94\xe4\xc6\x58\x89\xb1\x11\x6f\x36\x51\xe6\x58\x7d\xd1\x03\x6d\x1a\xaf\x19\xb8\xa1\xfc\xfb\x64\xac\xae\x79\x32\x78\x21\x20\xaa\x96\x7b\xd5\xf8\x41\x1d\xb1\x6b\x83\xa8\x08\x63\x17\xdc\x7e\x83\x07\x92\x80\xb0\x3e\x4e\xe3\xec\x69\x1e\x7d\x82\xc6\x16\xf3\xc5\x79\x5e\xab\x42\x7c\xdc\x36\x9d\x54\xd2\x3a\x23\x9c\x36\x99\xcd\x54\xd0\x0f\xb8\xf3\xba\xe7\x59\x44\x97\xea\x9a\x55\x6b\x52\xce\x9f\x1a\x28\x23\xc7\x47\x24\xfd\x35\xbc\x0e\x82\xec\xeb\x61\x33\x9e\x3c\x13\x14\x3f\x4f\x4f\x84\xe9\x08\x8e\x18\x28\xe7\x43\xda\x85\x3f\x28\x7c\x63\xfa\x46\xc7\x92\x67\xa5\xcf\x3b\x66\x00\xf9\xbf\x4e\x65\x6a\x7c\x8c\x39\x95\x8d\x68\xf0\x28\x07\x64\x48\x39\xd4\xfb\x71\xee\xf9\x89\xc8\x5d\x20\x08\x7d\xb1\x81\xfc\x79\x80\x46\x4c\x54\xd1\xcf\x92\xcf\x09\x04\x63\x25\x15\x04\x1b\xf5\x9d\x3f\xb3\xc6\xb3\x43\xc4\x62\x39\x1f\x93\x6c\x87\x4c\x0c\x4f\x42\xaf\xf0\x44\xcb\x3e\x97\x4a\xfa\x54\x85\xe9\x73\x9b\xed\xeb\x82\x67\x3e\x45\xd5\x45\x4e\x73\xd9\x3d\xcd\xc5\x58\x94\x66\xd2\xaf\x3b\x46\x82\xa7\xc0\x31\x84\x3b\x21\xd7\x0e\xc7\xe3\x08\x31\x74\xa2\x2c\x9f\xd0\x9f\x90\xe8\x1f\x82\xa5\xd9\xe2\xfc\x1a\xbe\xf3\x29\x0b\x77\x10\x9e\x8f\x2b\x84\x15\x03\xc9\x50\x2e\x14\xd3\xfb\x77\xc7\xac\xdd\x84\xf9\x3b\xaa\xc0\x11\xbb\x2d\x5a\xeb\x8d\x72\xe5\x43\x55\xbd\xa9\xd2\xc5\xe3\x37\xcf\xc3\xef\xa7\x64\xe9\x61\xac\x30\xb5\x81\x27\x35\xed\xe8\x62\x66\x2c\x41\xe1\x9b\x54\x2b\x9f\xc2\xa6\xd9\x74\x4a\x96\x8f\xb7\x53\xfc\x3e\xe0\x00\xfa\x6f\x6c\xf9\x8c\xf2\xbe\x95\x07\x88\xf8\x9e\xe4\x80\x44\x6f\x85\xe6\xec\x8d\x7a\x41\x67\x06\xc9\x34\xc8\xf4\x78\x41\x73\x01\xb8\x5c\x62\xed\xe4\x16\xdb\x3d\x5b\xe5\x33\xd9\x20\x80\x8f\x98\xff\xa0\x1d\x5e\x7b\xcd\xee\xe5\x45\x76\x75\x26\x7a\xa7\x3b\xe1\x24\xb5\xee\x1e\x6c\x5f\xf1\xed\x06\x36\xc3\xf1\xb2\x60\xb3\xfc\xe6\xb6\xb8\xf6\xe1\xa0\xfb\xda\x69\x73\xb2\xeb\x87\x54\xfc\xee\x1a\x9a\x56\x89\x88\x98\xe3\x92\x79\x5a\xc1\x8e\x9a\x61\x74\x85\x78\x88\xec\x0c\xdf\x8c\xed\x7c\x0b\x0c\xe1\xb2\xa5\xff\xbc\x58\x90\x92\xce\x53\xeb\x07\x5f\x96\x51\x56\xa1\x63\x1a\xce\x86\x08\x4b\x47\xb1\x45\x12\xa1\x52\x15\x37\x74\xa3\xe4\x17\xf9\x9b\xee\xd5\x71\x8c\xe7\x65\xf4\xa1\x21\xe6\xe4\x6f\x76\x73\xc9\xc6\xfc\x09\xe2\x2a\xf8\xbd\xc2\xac\x1a\xbe\x8a\x53\xdb\x13\x84\x84\x56\xd6\x50\x8b\x8d\xa8\x64\x2b\xdd\x3e\x0e\x0f\x96\xc1\x4d\x7e\x41\xc1\x37\x72\xf8\xb0\xd1\x16\xed\x78\xa6\xde\x07\x39\x7b\x0f\x1d\xba\xb5\xa6\x63\x9c\xd1\xfd\xca\x67\xec\x3e\x5e\x4e\xdd\xf3\x45\x8b\x59\x8a\x69\xa1\x52\xec\xad\x95\xea\xcb\xcd\x1f\x0f\x41\xf5\x75\xfa\xde\xeb\xf1\x87\x59\x81\x96\x2b\xbf\xb1\x22\x0f\xe9\x8e\xac\xf8\xd2\x09\xb3\x42\x77\x2a\x75\xe9\xf3\xdf\x39\x87\xa1\xfc\xf7\xb0\x94\xd8\x8e\x52\xf8\x26\xbe\xfb\xcd\x33\x18\x2c\x3f\x27\x81\xe1\xd3\xdf\x22\x7f\x4c\x07\xcc\xee\x43\x13\x14\x67\x8e\xd9\x69\xcc\xf3\xda\x13\x98\x67\x5b\x65\xbd\xee\x88\x40\x84\x0a\xd7\xf4\x5c\x0e\xbb\xd6\xbb\x4c\xff\xa5\xdb\xe4\x9d\xb0\xd9\x9d\xe6\x70\xf7\x95\x71\xd0\x89\x7f\x9b\x9a\xee\xde\xc7\xb3\xc7\xb3\x5f\x02\x00\x00\xff\xff\x18\x1f\x2f\x65\xd0\x1b\x00\x00"

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfc, 0x7b, 0xce, 0x2a, 0x2e, 0x77, 0x0, 0x73, 0x9f, 0x82, 0x36, 0xa6, 0x90, 0x17, 0x93, 0x40, 0xfd, 0x25, 0x56, 0x2f, 0xe2, 0x4f, 0xfd, 0x1d, 0xcd, 0x13, 0x61, 0x6b, 0xd5, 0xf, 0x42, 0x69}}
	return a, nil
}

var _flowtokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\xd8\xde\x43\x2b\xcf\xd9\xb2\x1f\x3a\x7d\xf0\x38\x77\x71\x1a\xbb\x93\x69\x9b\x76\x92\x5c\xfb\x6a\x88\x5c\x49\x68\x48\x40\x03\x80\x92\x75\x19\x7f\xf7\xce\x2e\xfe\x10\xa0\x48\xd9\x3e\x5f\xfc\x90\x98\x24\xb0\x7f\x7f\xbb\xfb\x03\x2c\xdb\x8d\x36\x0e\xee\x3a\xb5\x92\x8b\x06\xbf\xe8\xaf\xa8\x60\x69\x74\x0b\x97\x0f\x77\xbf\x7c\xfc\xdb\x87\x77\xff\xb8\xfd\xf2\xaf\xbf\xdf\x7e\xbc\x79\xff\xfe\xd3\xed\xe7\xcf\x27\x27\x9b\x6e\x01\x95\x56\xce\x88\xca\xc1\x5d\xa3\x77\xbc\xe7\x6a\x20\xe2\xdb\xc9\x09\x00\xc0\xc5\x05\x7c\xd1\x4e\x34\x60\xbb\xcd\xa6\xd9\x83\x5e\xf2\x16\x70\xb4\xc8\x82\x54\x80\x0f\xd2\x3a\x54\x15\xf2\x7a\x12\xbe\x15\x06\x1c\x6d\xfa\xcc\x7b\xae\xe0\x97\x3b\xf9\xf0\x97\x3f\x27\x81\xb7\x5b\x54\x0e\xdc\x5a\x38\x90\x16\xb0\x95\xce\x61\x0d\xbb\x35\x2a\x70\x6b\xec\x6d\x93\x16\x2a\x83\xc2\x61\x9d\x44\x23\x6f\x65\x0b\xed\x07\x25\x9d\x14\x8d\xfc\x15\xeb\x99\xf4\xbf\x97\x0a\x4f\x9f\xa5\xd1\x3b\x22\x0c\xc2\x4e\xba\x75\x6d\xc4\x2e\xc4\x4f\xc0\x7f\x44\xd7\xb8\x51\xdd\xff\x8d\x4b\x67\xa2\xd5\x9d\x72\x51\xe5\x19\x6f\xbd\x82\x9b\xba\x36\x68\xed\xcf\x2f\x35\xa1\xc6\x8d\xb6\x92\xbe\x38\x7d\xd4\x80\xf7\x71\xe1\x81\x01\x4e\xbf\x50\xbd\xc2\x5d\x6e\x42\x2b\xd5\x54\xc4\xff\xc9\x9f\x06\x1a\x5f\xee\xa2\x75\x46\xef\x27\x54\xbc\xeb\x8c\xfa\x6d\x2a\x04\x3b\xc2\xd6\x1b\x30\x68\x75\x67\x2a\x9c\xc6\x10\xfb\x62\xfe\xea\xbf\xcd\x44\xd3\xe8\x1d\xd6\x37\xbf\x55\xed\x82\xcc\x7e\x8e\x5a\xf6\x2f\xa9\xed\x35\xf4\x99\xbe\xb8\x48\x5a\x45\xb5\x86\xce\xa2\x01\xeb\xb4\x41\x0b\x42\x81\x54\xd6\x09\x55\x21\x15\xa2\x56\xcd\x9e\x2b\x86\x37\x53\x2d\xba\x35\x4a\xbf\x5a\xac\x30\xd5\xef\x1a\x61\xd9\xa9\xca\x49\xed\x2b\xb6\xdf\x22\x54\x0d\x2b\xbd\x45\x8a\x39\x2c\xbc\xb0\x8d\x41\x7e\xbf\xd1\xd6\x51\x2d\xd6\x92\x37\x46\x69\x52\x0d\x1a\x45\xac\xdb\x3d\x67\xb7\x12\x4d\x83\xf5\x3c\xd7\x5d\xad\xb1\xfa\x6a\x61\x2d\x36\x1b\x8a\x97\x03\xd3\x29\x27\x5b\xe4\x9d\xb8\x45\x03\x22\xd9\xc7\x81\x2b\x44\x44\x49\x9f\x42\x68\xe9\xbb\xf2\xae\x2f\x30\x06\x39\x7a\x45\xad\x03\x1f\x1c\x05\xa7\xe8\x24\x9c\x39\xb2\x31\x4a\xf3\x40\x5c\x4a\xc5\x7b\xcf\xc0\x6a\xfa\x6c\x38\x71\x4a\xc3\x4e\xec\x61\xa9\xc9\xb0\x56\x34\xb2\x92\xba\xb3\x3e\x11\x4e\x07\x95\x3e\x80\x29\x2a\xba\x0b\x4a\xa5\x02\x21\xcd\x1c\x6e\xc0\x6e\xb0\x92\xa2\x09\x40\xeb\xa1\xa1\x10\x6b\x4b\x82\x16\xbd\x09\x4e\x33\x70\xa3\xb4\xbe\x22\x8b\x28\x10\x8a\x92\x18\xd6\x3f\x68\xda\xf3\x7f\x1b\xbd\x95\x35\x9a\xb3\xc1\xfb\x4f\x58\xa1\xdc\x1e\xbe\x7f\x27\x1a\x06\x53\x68\xf6\x41\xfd\x5a\x37\x64\xe1\x1a\x61\x11\xbe\xeb\x25\x08\x0e\x80\x0d\x76\xa5\xe5\xb1\xdf\x87\x95\x65\xaf\x4f\x90\x89\x8d\xba\x10\x4a\x48\x88\xde\x70\x50\x29\xff\x04\x8c\xb4\x97\x36\xce\x06\x92\x4f\xe1\x5b\xfa\x4e\x3f\x16\x9b\xe5\x3c\x8a\x7c\x13\x85\xa7\x25\x8f\x85\x25\xb1\xc3\x67\xef\xf2\xcf\x77\x11\x85\x1e\x2f\xe2\x6b\xac\x39\x87\x2b\x82\x29\x77\x07\x10\xfc\x52\x98\x55\xd7\x62\xc8\x59\x44\x95\xaa\x93\x0a\xeb\x85\x84\x3d\x3c\x50\x52\xdd\xcd\xf3\x4d\x1f\x5c\x80\x94\x0d\xbd\xc4\x21\x8d\x74\x61\xf6\xa1\x48\x63\xdb\xe9\xac\x47\x0a\xa5\x27\x17\x40\x62\x5b\xad\x70\x9f\x56\x2e\x50\xaa\x15\x38\x23\x94\x5d\xa2\x31\x58\xcf\x49\x8b\x41\xd7\x19\xe5\x13\xab\x70\xd7\xec\x73\x21\xb1\x90\x82\x4a\x5d\x94\x13\xcb\xf5\x65\x49\x95\x22\x1d\xd7\xe0\x22\x1b\x57\xb9\x28\x6c\x2c\xee\xa8\x98\xe6\x63\x61\x26\xc0\x2c\x3b\x95\xe2\x34\x6c\xf5\x57\xf0\xb6\xc4\xa8\xb7\xe8\x68\xd2\x8b\xc7\xf3\x10\xf3\x62\x03\xb5\xec\xc9\xd1\xed\xff\x8f\xa3\x9b\x85\xe9\x9d\x42\xf3\xf3\x5c\xf8\x39\x7a\x5a\xc8\xf2\x71\x84\xeb\xf3\xbc\x13\xf4\x30\xf5\xd2\x4e\x27\x10\x18\x22\xf6\x12\x00\x86\x9c\xe8\xc5\xff\xb0\x1a\xa2\x8f\x21\x27\xea\xda\x16\xf5\xe6\x6c\x2a\xb2\x90\xc9\xac\x90\xe9\x91\xdd\xb3\xe3\x60\x94\x16\xc2\x34\xa4\xcd\x61\x5a\xf3\x2e\x4b\x0a\xbd\x31\x0b\xac\x44\x67\xb1\x87\x74\x51\x65\x64\x63\x06\x63\x02\x2c\x9a\xa8\x3b\x34\x36\x1e\x0d\xbc\xf5\x4f\xbd\xb5\x6b\x51\x38\xb2\x40\x54\x04\x42\xdb\xb5\x58\xb3\xab\xdc\xa4\x97\x9a\x07\x4d\x40\x60\x60\x13\xf3\x03\x84\x85\x50\xcf\x7c\x5a\xc7\x50\x35\xec\x25\x0d\x3a\xd8\xb2\x7f\xd7\xe7\x81\x03\xda\x3f\xc0\xdb\xc4\x91\xe7\xa5\xab\x4f\x21\xf1\x47\x2f\x6c\x3e\xec\x49\x03\x40\x1e\x52\xb9\x62\x9b\x67\x74\x4f\xa2\xb2\xd8\x03\x6f\xe0\x72\x7e\x59\x7c\x8f\x89\xdc\x16\x2e\x64\xe0\x0c\x0b\x66\xc3\xa0\xf4\xde\x67\x9c\x1e\xde\x4c\xbc\x3f\x2f\x42\x90\xe9\xc9\xb4\xa5\x76\x73\xdb\x6e\xdc\x7e\x8c\xf9\x94\x75\x50\xb6\x47\x0f\x40\x6a\x1f\x20\x72\x5c\xff\x8a\x46\xa7\xf1\xae\xea\xd4\xee\x64\xdf\xce\x44\xd3\x50\x63\x0c\x6d\x8d\x86\x34\x0f\xf5\xb6\xb3\xbe\xbd\xd1\xfc\xb6\x89\x8b\xe4\xc2\x98\x80\xb1\x10\x2f\x36\x75\xca\x21\xe9\xa2\x17\xda\xd4\x9e\x2a\x70\x1d\xf9\xef\x49\x58\x55\xf1\x40\xf0\xf3\x5f\x2c\x1a\x2e\x51\xe3\xe7\x73\x84\xac\x0d\x5c\x22\x4c\x5b\x70\xfb\x0d\x1e\x30\x01\x82\xf8\x30\x8c\xb3\xa7\xdb\xe7\x13\xdd\xeb\x72\x7e\x79\x9a\xe7\xaa\xe0\x1c\x37\x75\x2b\x95\xb4\xce\x08\xa7\x4d\x26\x33\x25\xf4\x23\xee\x3c\xdd\x79\x56\x7f\x4b\x79\xcd\xb2\x35\xca\xe2\x8f\xcd\x91\x81\xe2\x09\x26\x7f\x05\x6f\x03\x0f\xfb\x76\x58\x86\x47\x8f\x02\xc5\xe3\xf1\x41\x30\x6e\xc1\x84\x80\x72\x2c\x24\x2f\xfc\xf9\xe0\x95\xe1\x1b\x9c\x46\x9e\x15\x3e\xaf\x98\x01\xe4\x7f\x1d\x8b\xd4\xf0\xf4\x72\x2c\x1a\x51\xe0\x64\x0f\xc8\x90\x72\x48\xf3\xe3\xb8\xf3\x83\x90\xab\x40\x10\xfa\x62\x01\xf9\x63\x00\x4d\x96\x48\x9e\x9f\xc5\x9a\x13\x08\x86\x04\x2a\xf0\x34\xaa\x3b\x7f\x54\x8d\x47\x86\x88\xc5\x72\x2c\x26\xb6\x0e\x19\x07\x1e\x85\x5e\xa1\x89\xb6\x7d\x29\x09\xf4\xb1\x0c\xd3\x72\x9b\xf9\x75\xc6\xa3\x9e\xac\x6a\x63\x4f\x73\xd9\xe5\xcc\xd9\x90\x8b\x66\x8c\xaf\x9d\x6a\x82\xc7\xc0\xd1\x9b\x3b\xc6\xd2\xca\xa9\x38\x80\x0b\x9d\x22\xcb\x37\xf4\x13\xa2\xfc\x53\x10\x33\xbb\x3c\xbd\x82\x1f\x7c\xbc\xc2\xbd\x83\x6f\xc6\x0b\x84\x15\xa3\xc8\x50\x20\x14\xf7\xf6\x1f\xa6\xa4\x5d\x87\xb1\x3b\x08\xff\x84\xdc\x06\xad\xf5\x42\x39\xed\x21\xa5\x5e\x54\xa9\xe2\xf1\x15\x63\xf0\xc7\x31\x12\x7a\x68\x25\x8c\x99\xfe\x24\x83\x1d\x5c\xc3\x0c\x09\x27\xbc\x8a\xa3\xf2\x99\x6b\xbc\x89\x8e\x91\xf0\xa1\x3b\xc5\xf3\x41\xe9\xd3\xbf\xb1\xd2\xb3\x4e\xf7\xda\xf2\xa7\x7e\xf7\x64\xe9\xa7\xae\x56\x30\xcc\xce\xa8\x17\x14\x64\xe0\x48\x3d\x29\x8f\xd7\x31\x67\x80\xcb\x25\x56\x4e\x6e\xb1\xd9\xb3\x54\x3e\x81\xf5\x74\x77\x42\xfc\x47\xed\xf0\xca\x33\x74\xcf\x2a\xb2\x8b\x32\xd1\x39\xdd\x0a\x27\xa9\x62\xf7\x60\xbb\x05\xdf\x65\x60\xdd\x1f\x26\x8b\x26\x96\xdf\xd2\x16\x97\x3c\x6c\x74\x57\x39\x6d\x8e\x16\x7b\x1f\x8a\xef\xcb\x98\x69\x8b\x88\x70\x99\x26\xc8\xe3\x7c\x75\x50\x09\x83\xdb\xc2\x43\x58\x67\xe0\xf6\xc0\x26\x10\xdd\x78\x0c\x5d\xc1\x4d\xe7\xd6\xe1\x21\x77\x8c\x81\x5d\x56\x39\xb1\xe9\x3c\xda\x7e\x04\x66\x41\x66\x3e\x3a\x6c\xc8\xd9\x38\x61\x12\x29\xb6\x48\x74\x54\xaa\xe2\x8a\x6e\x90\x8f\x22\xa4\xe3\xe5\x3b\x34\xb0\xf7\x38\xf7\x6f\x4e\xfa\x66\xd7\xe7\x2c\xcc\x9f\x22\x2e\x82\xde\x8b\x65\x4c\x90\xcf\xea\x98\x6f\x82\x90\xd1\xc8\x0a\x2a\xb1\x11\x0b\xd9\x48\xb7\x8f\x33\x84\xd9\x70\x9d\x5f\x4f\xf0\x7d\x1c\x3e\x6c\xb4\x45\x3b\x1c\xad\xf7\x81\xd5\xde\x43\x8b\x6e\xad\xe9\x10\x67\x74\xb7\xf2\xe1\xba\x8f\x57\x53\xf7\x7c\xcd\x62\x96\x62\x9c\xaf\x14\x8e\x35\x52\x7d\xbd\xfe\xe3\x00\x64\xdf\xc6\xaf\xbc\x1e\x7f\x9a\x15\xe8\xb9\xf0\x5e\xf5\x11\x48\x77\x63\xc5\x32\x27\xcc\x0a\xdd\x64\xc4\xd2\xda\xef\x1c\xba\x90\xf2\x7b\x58\x4a\x6c\x06\x91\x7b\x17\xbf\xfd\xbe\x81\x0b\x62\x9f\x8c\x5b\x58\xf7\xea\xb0\x71\x3f\xe0\xde\xde\xe3\xbd\x38\x68\xcc\x8e\xc3\x9b\xdf\x4d\xc1\x9b\x05\x95\x39\xba\xa5\xf6\x21\x54\xb8\x8f\xe7\x14\xd8\xb5\xde\x65\x8c\x2f\x5d\x1b\xef\x84\xcd\x2e\x2f\xfb\x4b\xae\xac\x03\x1d\xf9\x23\xd4\x78\x95\x3e\x9e\x3c\x9e\xfc\x3f\x00\x00\xff\xff\x22\x1a\x70\x0c\xb4\x1b\x00\x00"

func flowtokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_flowtokenCdc,
		"FlowToken.cdc",
	)
}

func flowtokenCdc() (*asset, error) {
	bytes, err := flowtokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "FlowToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x83, 0x7e, 0x8, 0xf8, 0xcf, 0x83, 0x81, 0xa1, 0xaf, 0x40, 0x84, 0x75, 0x1a, 0x89, 0x64, 0x9d, 0x36, 0x49, 0x4b, 0x8f, 0x88, 0x10, 0x3d, 0xc3, 0x39, 0x3d, 0x5d, 0x72, 0xe2, 0x10, 0xaf, 0x72}}
	return a, nil
}

var _fungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x4d\x73\xdc\xb8\x11\xbd\xf3\x57\x74\xd9\x07\xcb\xce\x58\xda\x43\x2a\x07\x57\x79\x13\xbb\xd6\xaa\xf2\x25\x49\x25\x4a\xf6\x3a\x18\xb2\x39\x83\x15\x08\x70\x01\x70\x46\xb4\xcb\xff\x3d\xd5\x8d\x0f\x82\x1c\x4a\x1a\x55\xd6\x17\x6b\x48\xa0\xd1\xdd\x78\xfd\xfa\x01\xbc\x79\xf7\xae\xaa\x5e\xc3\xdd\x01\xe1\x56\x99\x13\xdc\x0e\x7a\x2f\x77\x0a\xe1\xce\xdc\xa3\x06\xe7\x85\x6e\x84\x6d\xaa\xea\xf5\x6b\xd8\xa6\x97\xfc\x6e\x0b\xb5\xd1\xde\x8a\xda\x83\xd4\x1e\x6d\x2b\x6a\xac\x2a\x32\x94\x7f\x82\x3f\x08\x0f\x42\x29\x68\x93\x59\xcf\x66\xd3\x4c\x07\x27\x33\xa8\x06\x0e\xe2\x48\xaf\xe8\x79\x6b\x6c\x07\xde\x5c\x57\x5f\x5b\x10\x30\x38\xb4\x0e\x4e\x42\x7b\x47\xef\x1b\xec\x95\x19\x41\x80\xc6\xd3\xc2\xd4\x06\xfc\x01\xa5\xcd\xbf\xab\x60\x59\x23\x36\x34\x53\x76\xbd\xc2\x0e\xb5\xa7\x61\x30\x0b\x64\xf2\xf7\x9a\xfd\x2f\x8c\x2c\xdc\x6b\x8d\xa2\x1c\x51\x40\x64\xc5\x0e\x0a\x1d\x08\xdd\x80\x16\x9d\xd4\xfb\x8a\xc3\xf5\xb3\x0c\xb8\x1e\x6b\xd9\x4a\x74\xd7\x21\x85\xff\x15\x83\xf2\x5b\xb0\xe8\xcc\x60\x29\x61\x5f\x44\x7d\x00\x51\xd7\x66\x60\xdf\x84\x07\x73\xd2\x2e\x04\x97\xd2\x93\x82\x60\x3f\x04\x39\x4c\xfb\x52\x63\x65\x5a\x5e\x8e\x8d\x66\x9b\xe0\xbc\xb1\xd8\x80\xd4\x31\x25\xc9\x3a\x3d\x17\xfb\x18\xe5\x72\xd2\x41\x38\xe8\xd0\x1f\x4c\xe3\x20\xc7\x61\x4e\x1a\x2d\x47\x68\xfc\x01\x6d\xdc\x8e\x5a\x68\xa8\x85\x52\x31\xa4\x7f\x5a\x73\x94\x0d\xda\xed\x06\xb6\xff\xc2\x1a\xe5\x91\xff\xa6\x59\xdb\xcf\x42\x91\xa3\x53\xc0\x53\x6a\x1c\xbb\xe1\xca\x27\xd0\x60\xad\x84\x45\xe8\x2d\xbe\xaf\x8d\x6e\xa4\x97\x46\x87\x14\xf7\xc6\xf9\xf2\x19\xfb\x68\xd1\x79\x2b\x6b\x5f\x91\xb3\xf8\x80\xf5\x40\x2f\x21\xa6\xa5\x1d\x74\x1d\x06\x87\x54\x84\x90\x43\xf8\x23\xd0\x3a\x0e\x7b\x61\x85\x47\xd8\x61\x2d\x06\xf2\xc5\xc3\x5e\x1e\xd1\xf1\x70\x8a\x96\xff\x10\x3b\xa9\xa4\x1f\x69\x0b\xdc\x41\x58\xac\x04\x58\x6c\xd1\xa2\xae\x19\x17\x21\xcd\x21\xa1\x61\x0b\xb5\x1a\x01\x1f\x7a\xe3\xa2\xa9\x56\xa2\x6a\xdc\xe4\x51\x25\x35\x18\x8d\x60\x2c\x74\xc6\x62\xf2\x78\x4a\xc5\x75\x55\x7d\xa5\xd2\x71\x26\x3a\x14\x52\xbf\xf0\xa6\x13\xf7\x08\xf5\xe0\xbc\xe9\x72\x86\x63\x6a\x32\xe0\x29\x37\xf3\x2c\x53\x21\x19\x38\x0a\x2b\xcd\x40\xa3\xa5\xde\x3b\x38\x49\x7f\x60\xf3\x01\x79\xd7\xd5\xad\xb1\x80\x0f\x82\xcc\x6c\x40\x40\x2b\x86\x1a\x3d\xef\xfd\x0e\x27\xeb\xd8\xc0\x6e\x4c\x75\xcb\x35\xc0\xe9\x80\x04\x8a\x59\x71\x7d\x1e\x61\x70\x52\xef\x0b\x5f\x69\x6b\x27\xd7\x36\x31\x4c\xd3\x2e\x4a\x34\x13\x46\x45\x0e\x38\xd4\x0d\xcf\xb4\x01\x6e\xa9\x5a\x7a\x44\xfb\xde\x9b\xf7\xf4\xff\x86\x23\x32\x83\xa7\xaa\xa1\x35\x89\x04\x68\x21\xe6\x06\x0a\x56\x40\x8d\x64\x55\x81\xc2\x66\x8f\x16\x5c\x27\xac\xcf\x4b\x5d\xc3\x9d\x09\x2b\x45\xeb\xde\x80\xd0\x53\x1d\x6c\xaa\x40\x4f\xb1\x46\x1d\xa5\x64\xe4\x45\x1b\x2b\x4e\x45\x2a\xa1\xb5\xa6\x2b\x31\xc2\x54\x15\x4a\x88\x81\xdb\x60\x6f\x9c\xf4\x19\x1d\x60\xf4\x6c\xa5\x37\x2e\x61\x8b\x18\x92\x32\xef\x31\xd8\xb7\x42\xbb\x16\xed\x75\x55\xbd\xbb\xa9\xaa\x9b\x1b\xe6\xf1\x4e\x48\xbd\xe4\xf1\x62\x17\x6e\x6e\xe0\x1f\x6c\xfa\x71\x4e\x96\x4a\xcd\x08\x53\xba\x82\xe2\x6f\x6e\xaa\x7e\xd8\xad\x90\xff\x62\xcb\xbe\x57\x15\x00\x40\x74\xca\x1b\x2f\x14\xe8\xa1\xdb\xa1\x65\xb4\x87\xd4\x48\x0d\xf8\x20\x9d\xa7\x4a\xba\x4e\xe3\xbf\x7a\x90\x0e\x86\x3e\x96\x56\x01\x36\x4b\x8f\x50\xbb\xc1\xc6\xde\x12\xcc\xba\xa1\xef\xd5\x98\xa6\x3b\x2f\x46\x47\xa4\x37\x70\x69\x13\x4e\x82\xad\x46\x78\xe4\x41\xe4\xff\x51\xd8\x30\xfb\xdf\x3c\xf9\x03\xfc\xe7\x56\x3e\xfc\xe5\xcf\xd9\xe9\x2f\x47\x4c\x84\x2c\x1d\x60\x27\x3d\x61\xfd\x44\x1b\x47\x3e\x4d\xe1\x3b\xa8\x2d\x0a\x8f\x4d\x36\x8d\x3c\x95\xb3\xe0\xbe\x6a\xe9\xa5\x50\xf2\x1b\x36\x57\x32\xfc\x3d\x5f\xf0\xed\x45\x2b\x86\x6c\x11\x65\x25\x80\xe9\x00\x2b\x11\xa0\xb1\xba\xf6\xaf\x69\xe8\x95\xe8\xa8\x01\xa4\x25\x37\x3c\xf5\x03\x7c\x6a\x1a\x8b\xce\xfd\xf5\xa5\x2e\x44\xbc\x86\x9e\xf4\x94\x03\xbf\xa4\x81\x67\x0e\x78\xb3\xb6\x7c\x22\x8d\xf8\x3b\xe3\x61\xae\x27\x90\xb8\xa6\x8e\xc4\x6a\xf1\xf7\x41\x5a\x46\x87\x83\xd6\xd8\x9c\x1f\xe2\xa2\x38\x7f\x51\x86\x13\x9e\x98\x16\xc6\x3e\x23\x2f\x4d\xf8\x15\xa1\x31\xfa\x4d\x5e\x6a\xbe\x8a\xd1\xb0\xdd\xa5\xbe\x76\x40\x8b\x9b\x34\xaf\xe8\x22\x0a\x05\xb1\xb6\xe9\x23\x5e\x7a\xe3\x9c\x8c\xc4\x6d\xda\x00\x19\x5a\x3e\x92\x77\x1f\x23\x77\xd9\x67\x8a\x34\x38\xa1\xb1\x46\xe7\x84\x95\x6a\x8c\x4a\x80\xa9\xc4\x9c\x34\x44\x37\x66\xfe\xd3\x26\x9c\x77\xdb\x89\x90\x63\x5d\xc6\x75\x52\xba\x8a\x67\xe5\xeb\xdb\x44\x49\xec\x8f\x1b\x76\x91\x22\x96\x29\x65\xa1\x90\x78\xaa\x34\x10\x58\xda\x0f\x96\xb0\xb3\xd4\x1c\xb9\xe7\x58\xec\xcc\x11\x9b\xdc\x7b\xd6\x9d\xb9\x2b\x7a\xfa\x1b\xae\x70\x74\x0e\x14\x1e\x51\x11\x5c\xfb\x61\xa7\x64\xbd\x81\xdd\x90\x38\xcb\x51\xfa\x04\x25\x77\xa7\xb0\x2b\x4d\xa5\x9d\xe2\x46\x3d\x29\x1d\x6e\x2e\xde\x58\x06\x06\xfb\x95\xf3\x38\xd7\x52\xa5\xad\x9a\x15\x19\x97\xb6\x1a\x99\xd5\xc3\xf2\xc9\xd5\xa7\xc2\x09\xcb\x76\x62\x84\xbd\x15\xda\x47\x99\x15\x17\xc9\x21\x52\x87\x4d\x80\xa1\x70\xe4\x31\x31\x59\x76\xa1\xcf\xaa\x80\x36\x2a\xb4\x24\xc1\x6a\x35\x2a\xd0\x7a\x26\xe1\xa8\x70\xd9\x76\x69\x85\x71\x9a\x60\x92\x43\xf7\x07\x6b\x86\x3d\xb5\xcb\xac\x79\x2e\x8b\x28\x88\x17\x0e\x8b\x72\xf2\x4c\x50\xbc\x79\x97\xc6\x44\xf6\x56\xc3\x99\xc5\x50\x5a\x7b\x71\x38\x54\x46\xed\xa0\x73\x81\x2c\x28\xec\xed\x07\xf8\x5b\x40\xf3\xf7\x3c\x85\xa7\x19\xb7\x7c\x14\x3d\xd8\x5a\x74\x51\xfb\xb7\xd1\xe7\x00\x31\x2a\x0e\x38\x0a\x35\xe0\xd9\xb4\x30\xe5\x3a\x96\x39\x7c\xfc\x08\xd1\x8b\xb3\x91\xf4\xef\x55\x22\x7c\xa1\xe2\x38\xe8\x06\xe7\x49\xaf\xd1\x4a\x4e\x74\x08\x22\xa4\x28\x59\x8c\xba\x73\x6a\x2a\x1c\xd3\xab\x99\xf9\x1f\xd5\xfc\xaf\x1f\x99\xaf\x93\xda\xff\x7f\xf8\x3a\x36\x93\x73\xba\x96\x7a\xd9\xfe\x9f\xa5\x6b\xa9\x6b\x35\x34\x48\xca\x2e\x1d\x14\x82\x0b\xf5\x01\xeb\xfb\x79\xe4\x91\x01\x92\x8d\x13\xf2\x29\x93\x76\x85\xf4\xf6\x25\x72\x3b\xc4\x1e\xe4\x76\x55\x70\x41\x63\xd2\x98\x75\x69\xbd\x01\x25\xef\xe9\x64\xa8\x24\x9f\xb2\x3a\x92\x27\x42\x37\x59\xbf\xb0\xe6\xa4\xe7\xa4\x59\x64\xcb\x28\xf5\xd0\xab\x70\x2e\x80\x67\xa9\x3e\x6d\xcb\x82\xea\x63\xa6\x2f\x62\xfa\x28\xf3\x89\xcc\x42\x9b\x4f\x1a\x35\x84\x50\x4e\x5c\xdf\xa7\xa9\xde\xc6\x1e\x9f\xaa\xaf\x68\xf8\x2a\x48\x92\x50\x53\x6f\x97\x45\x65\x71\xa5\xa6\x68\x46\x2e\x8d\x9f\x63\x5d\x5e\xfd\xf4\xf6\x91\xe2\x88\x62\x24\x03\x20\x95\x46\x80\xdf\x11\x2f\x05\x7d\x3c\xd2\x3e\x8d\x79\x92\x87\x42\xea\x80\xa0\x49\x2d\xf0\x31\x10\xca\x53\x7b\x9a\x4f\x7d\xb2\x28\x14\xe2\x3b\x52\x5d\x1a\x4f\x61\xdc\x9b\x20\xbd\xa2\xd4\xdc\x94\x50\x4e\x26\x58\xa5\x67\xb5\x09\xb5\xb1\x16\x6b\xaf\xc6\x4b\x20\x13\x83\x5a\x20\x66\x12\xee\x0b\xbe\x88\x5c\xfe\xc6\x2d\xf1\x90\x94\x75\x1c\x3f\x57\xd5\xf4\x8f\x3c\xbc\x5a\xbc\x3d\xdb\xee\x75\x0e\x75\xa8\xda\x92\x0a\x93\x95\xf5\xed\xfe\xbc\xd8\xe6\x32\x35\x09\xb0\xe1\x51\x32\x74\x29\x00\xca\x7d\x2b\x8f\x38\x45\x8b\x59\x22\x60\xba\x89\xf0\xe6\xb1\xd3\xeb\x92\xd5\xee\xf8\x64\x58\x2b\x61\x45\xba\xd4\x60\x5e\xab\x2d\x1f\xff\xc6\x9e\x55\x89\x58\x3b\x88\x75\x28\xf4\x9c\x97\xf0\x88\x76\x5c\x1e\x0a\xf3\xcc\xf9\x85\x81\x5b\x9e\xf8\xa2\x0d\xce\x64\x83\xad\xd4\x58\x7a\x12\xba\xa0\xd9\xfd\x86\xd1\x52\xe6\xc2\x70\x29\x90\x3b\xdd\x65\x17\x45\xc5\xfd\x50\x51\x1a\x91\xd9\x39\x9b\x2e\x5f\xa7\xd0\x9b\xe9\x4a\xe5\x31\x94\xb3\x83\x1f\xb2\x00\xde\x64\x7e\xdc\x14\xb0\x7f\x01\xea\x5f\x0e\xfa\x68\x74\xba\x25\x09\xdb\x17\x13\x1a\xae\xbc\x26\x11\x29\xbf\xcd\xa5\x4b\xba\x44\x35\x27\x47\xda\x8f\xa2\x48\xfa\x74\x81\xe0\xe4\xe8\x71\xa1\xc2\x9f\xae\xbe\xd5\xf3\x40\xa1\xf5\xb7\x41\x4b\x6c\x27\xb5\xcf\xf6\x5d\x5e\x35\xf5\x2e\x48\x52\x2b\xc9\xfd\x63\x09\x8d\xdc\x05\x93\x65\x6c\x66\x0c\x06\x7f\x8c\xfe\x5a\x6d\x15\x33\xea\xf8\xf9\x19\x15\xf5\x29\x48\xa7\x49\x13\x25\x0a\x51\x41\x60\x0a\x0d\xc6\x02\xfe\x3e\x08\x15\x7e\xad\x08\xaa\x27\x65\x14\x3c\xa9\x13\xe9\x58\xc2\x69\x22\xdd\x2e\xd4\x74\x2b\xb4\xdd\x61\x6b\x2c\x6e\x59\xa3\xa0\x8f\x3b\xa1\x86\xbc\xe8\xa2\xcf\xac\x19\x8f\xf7\xba\x3b\xdc\x4b\xad\x09\x45\x8b\xab\xd2\xe9\x12\x75\x65\xf6\xf3\x8c\xcc\x0e\x5e\x95\x8f\xdf\xc2\xfb\xa7\xb3\xfd\xf7\xd4\xe1\xce\x1a\x33\x5f\x8d\x45\xf9\x33\x65\xb6\xb7\x78\xe4\x7b\xcb\x02\x7d\x2f\xd3\xb0\x2b\x9a\x08\xbc\xb8\xc7\x33\xc4\x0a\x7a\xd2\x0b\x2b\x3a\xf4\xf1\x1a\x5c\x34\xcd\x5c\xfc\x14\x65\x10\x69\x6e\x81\x84\x78\x19\xff\x68\x49\xbe\x48\x07\x5d\xd8\x18\xd7\xb6\xe1\x4f\xe9\x71\x29\x9b\x1e\xd3\x4a\x4f\x6f\x8a\x1b\xba\x67\x77\x63\xba\x23\x7a\xd1\x89\x22\xe8\x9b\x2f\x5d\xef\xc7\xb5\x36\xfb\x49\x8f\xe1\xce\x35\x7d\x82\x98\x9f\xb4\xf9\x82\x94\x2d\xc4\x4f\x44\x65\x6f\x9a\xdd\xac\x1c\xc4\xe4\xf2\xc7\x8f\xf0\xd3\xb2\x79\xd0\x8e\x2c\x7d\xb9\x5a\xe3\x9c\x95\x2d\x39\x3f\xb4\x4d\xd2\x14\x5e\x51\x23\xd0\x78\x52\x63\xd2\x72\xd1\x49\x4e\x30\x7f\xe0\xf9\x86\xd6\x9c\x6b\x92\x94\xa9\x1f\xf1\x9e\x57\xba\xa9\x7f\x97\x32\x41\x3a\xf8\x8d\x4d\xa1\x65\x48\x86\x03\xf3\x50\x7e\x3e\x3b\xfb\x02\x56\x85\x6e\xbd\x94\x11\x62\x67\x8e\xb8\xc9\x77\x25\xe7\x23\xf8\x53\x90\x36\x0c\x8c\x60\x1b\x1b\xb2\x65\x74\x71\x47\x15\x29\xa7\x33\xfc\x31\x62\x79\x7b\xfc\xcb\xd0\x75\x23\x7c\xff\x51\xc1\xff\x02\x00\x00\xff\xff\xb9\x57\xa9\xfb\x8a\x1c\x00\x00"

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xde, 0xd5, 0x68, 0x7c, 0x88, 0xc9, 0xa8, 0x1e, 0xd, 0xee, 0xbe, 0x78, 0xd1, 0xb4, 0x65, 0xce, 0x74, 0x82, 0x74, 0x6c, 0x7b, 0xd8, 0x25, 0x89, 0x19, 0x24, 0x99, 0x66, 0x9b, 0x3d, 0x17}}
	return a, nil
}

var _tokenforwardingCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\xc1\x6e\xe3\x36\x10\xbd\xf3\x2b\x06\x28\xd0\xd8\x41\x22\xf7\x50\xf4\x10\xa4\x48\xd3\xc6\x2e\x8a\x2e\xb2\x80\x93\xec\x1e\x17\x63\x72\x64\x31\x96\x49\x81\x1c\x59\x31\x82\xfc\xfb\x82\x94\x44\x4b\x89\xb3\x1b\x5f\x2c\x09\x9c\x37\xef\xbd\x79\x23\xcd\x4e\x4f\x85\xf8\x05\x16\xb5\x59\xeb\x55\x49\x70\x6f\x37\x64\x60\x61\x5d\x83\x4e\x69\xb3\x86\x7f\xac\x61\x87\x92\x85\xb8\x2f\xb4\x07\xd9\xdd\x82\x2f\x6c\xe3\xa1\xb0\x0d\xa0\x01\x94\xd2\xd6\x86\x41\xda\xba\x54\xe0\x89\xa1\xae\x00\x41\xd6\x9e\xed\x36\x81\xb7\xd8\x4b\x92\xa4\x77\xe4\x04\x5b\xc0\xb2\xb4\x0d\x70\x41\x5b\x60\x0b\x79\xdb\x15\x38\x9c\xf3\xe1\x09\x82\xd2\x79\x4e\x8e\x0c\xa7\x1e\x4d\x41\x86\x76\xe4\x42\xd9\x1e\x5c\x8b\xd6\xd5\x64\x81\x25\xed\x41\xa2\x81\xaa\x5e\x95\xda\x17\xc0\x81\x76\x27\x88\x1c\x38\xf2\xb6\x76\x92\x00\x3d\x60\x22\x03\x12\x2b\x5c\xe9\x52\xf3\x1e\x1e\x6b\xcf\x50\xea\x0d\x01\xc2\x17\xac\x4b\x3e\x13\x68\x54\x68\x07\x9e\x4c\xc0\x50\x96\xbc\x39\x61\xa0\x1d\x19\x30\x44\x81\x32\x6c\x8c\x6d\x40\x33\x68\x7f\x20\x9d\x09\xf1\xb5\x20\x33\xb4\xa8\x41\xc3\x51\x9b\x74\x84\x1c\x7a\x24\x6e\x67\xad\x24\x89\x65\x19\xbb\xb5\x27\x6e\xa9\x49\x27\x44\x5e\x1b\xc9\xda\x06\x44\x05\x95\xb3\x3b\xad\x28\x34\x6d\x34\x17\xb1\x26\x09\x72\x14\x29\x48\x02\x2e\x90\x5b\xe4\xd0\x7b\x60\xb4\xe0\x82\xb4\x3b\xd8\x9d\x09\x71\x3a\x13\x42\x6f\x2b\xeb\xf8\xd5\xd4\x72\x67\xb7\xf0\xdb\xd3\xe2\xe1\xf6\xdf\xff\xfe\xfe\x34\xbf\xff\xfc\xff\xfc\xf6\xfa\xe6\x66\x39\xbf\xbb\x13\xa2\xaa\x57\x87\x60\xc4\xf3\x83\x00\x3d\x0b\x01\x00\x30\x9b\xc1\x7c\x17\xe6\x18\xe9\x68\x0f\xb4\xd5\xcc\xa4\xe2\x3c\x7b\x0e\xe8\x08\x14\x55\xd6\x6b\x6e\x4d\x0d\x92\x18\xdd\x9a\xb8\x9f\xb4\x8b\x68\xa1\x23\x45\xb8\xde\x1b\x75\xd3\xd6\x4d\x70\x1b\x7c\xbe\x80\x87\x85\x7e\xfa\xe3\xf7\xb3\xc8\xfc\x02\xae\x95\x72\xe4\xfd\xd5\x54\xa4\xfa\x94\x84\x64\xef\xc5\x58\x74\x96\xcc\xec\x34\x74\x3a\xe2\x22\x68\x1f\x98\x3b\x8a\x14\x87\x9c\xa3\x90\x46\x97\x25\xac\x62\x60\x38\x1b\xd7\x12\xf0\xbe\x22\xd0\x46\x69\x89\x4c\xbe\x33\x24\x7a\x82\xc3\xb1\xd9\x78\x3b\x10\xdd\x42\xa4\x4b\x94\x92\xbc\x9f\x78\x2a\xf3\x29\xec\x30\x8c\x5c\xea\x4a\x53\x10\xff\xeb\xf3\x71\x25\x2f\x23\x21\x1d\xed\x63\xe0\xb3\x59\xf0\xa2\xcd\x5a\x1b\x20\xdc\x90\xef\x37\x02\xec\xea\x91\x24\xc7\x1d\x32\x80\x6e\x5d\x6f\xe3\x8a\x1a\xd5\x67\xcb\x0f\x91\x34\xf7\xb3\x4c\x14\x4f\x7c\x87\x54\xfb\x10\x92\xb8\x5c\x6c\x1d\xa9\x83\x03\xc7\x68\x85\xb9\xe5\xb5\xe9\x99\x4f\xda\xe1\xfe\x35\x16\x1b\x81\xa7\xf0\x9c\xaa\xc2\x2f\xc4\xed\xfd\xb0\x04\x9c\x6c\x85\x25\x1a\x49\x7d\x64\x82\xb1\x99\x6d\x0c\xb9\xab\x0c\xdb\xf8\x4c\x47\x90\xf1\x40\x52\x94\x8d\x39\x5d\x9e\x87\xff\x43\xc1\xd8\x79\x59\xa0\x59\xd3\xb2\xaf\xed\xee\xfd\xd8\x23\xb0\x79\x7c\x90\xa7\x97\x57\xe7\x62\xb7\xf8\xea\x70\xf4\x47\x5e\xbd\xea\x35\xf9\x06\x86\x9a\xe5\x07\xb2\xf2\xda\xc2\xb1\x5e\xf8\x73\x84\x73\x4c\xa8\x36\x9a\x27\x1f\x49\xe5\x4f\x3b\xbd\x15\xfa\x22\x06\xcd\x82\xa3\x6f\x5e\x96\xdd\xa3\x10\x5a\x43\xcd\xe8\x13\xd0\xef\x58\x7a\x6d\xbe\xe3\x68\xe7\x66\x72\xf2\x4d\x8f\x8f\xa9\x0b\x09\x4d\xdd\x0f\x4a\x1d\x71\xed\x0c\x5c\x9e\x77\x9f\x82\xa3\xa8\xe9\x72\xda\x09\x7e\x11\xdf\x03\x00\x00\xff\xff\xe1\xda\x62\xc3\xb4\x07\x00\x00"

func tokenforwardingCdcBytes() ([]byte, error) {
	return bindataRead(
		_tokenforwardingCdc,
		"TokenForwarding.cdc",
	)
}

func tokenforwardingCdc() (*asset, error) {
	bytes, err := tokenforwardingCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "TokenForwarding.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x86, 0x2f, 0xdb, 0x6e, 0x3, 0x2e, 0x56, 0x40, 0x81, 0xa8, 0x6d, 0x2, 0x2b, 0x7b, 0x8, 0x86, 0x69, 0x85, 0x46, 0x7d, 0xa0, 0x6d, 0xa, 0xe6, 0x8f, 0xeb, 0x99, 0xdc, 0x66, 0x72, 0xe0, 0xd7}}
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
	"ExampleToken.cdc":    exampletokenCdc,
	"FlowToken.cdc":       flowtokenCdc,
	"FungibleToken.cdc":   fungibletokenCdc,
	"TokenForwarding.cdc": tokenforwardingCdc,
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
	"ExampleToken.cdc":    &bintree{exampletokenCdc, map[string]*bintree{}},
	"FlowToken.cdc":       &bintree{flowtokenCdc, map[string]*bintree{}},
	"FungibleToken.cdc":   &bintree{fungibletokenCdc, map[string]*bintree{}},
	"TokenForwarding.cdc": &bintree{tokenforwardingCdc, map[string]*bintree{}},
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
