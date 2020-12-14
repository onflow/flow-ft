// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../transactions/burn_tokens.cdc (1.102kB)
// ../../../transactions/create_forwarder.cdc (1.8kB)
// ../../../transactions/mint_tokens.cdc (890B)
// ../../../transactions/scripts/get_balance.cdc (460B)
// ../../../transactions/scripts/get_supply.cdc (229B)
// ../../../transactions/setup_account.cdc (1.186kB)
// ../../../transactions/transfer_tokens.cdc (1.376kB)

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

var _burn_tokensCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xc1\x6e\xd3\x40\x10\xbd\xfb\x2b\x1e\x3d\xa0\xe4\x50\xbb\x48\x88\x43\x54\x28\x69\x9b\x20\x04\x0a\x12\x4d\xe1\xbc\xb6\x27\xf1\x0a\x7b\xd7\x9a\x1d\x93\x54\x55\xff\x1d\xed\xae\xe3\xda\x48\x6d\x0e\xd9\x64\xf7\xcd\x9b\xf7\xde\x4c\x96\x61\x5b\x69\x07\x61\x65\x9c\x2a\x44\x5b\x03\xed\xa0\x20\xd4\xb4\xb5\x12\xc2\xce\xb2\xff\x3b\x7a\x97\x4a\x49\x92\x65\x28\x6c\x57\x97\xc8\x09\x9d\xa3\x12\xf9\x03\xa4\x22\xa8\xb2\xd1\x06\xaa\x28\x6c\x67\x04\x62\x91\x77\x6c\x20\xf6\x0f\x19\xe7\x8b\x76\x6c\x1b\x0f\xd4\x0c\x27\x96\xa9\xc4\x2f\xd5\xd5\x9e\x2f\x09\x5a\x28\x14\x68\xb3\x87\x6a\x02\xc5\xe1\xd4\x45\xa1\x55\xac\x1a\x12\x62\xcf\xeb\x9b\x8d\x54\x25\x89\x6e\x5a\xcb\x82\x75\x67\xf6\x3a\xaf\x69\xeb\x5b\xc6\x76\x17\xc7\xf5\xfd\xe6\xcb\xd7\xeb\xef\xab\xed\x8f\x6f\xab\xcd\xf2\xf6\xf6\xe7\xea\xee\xee\x54\xb0\x3a\xaa\xa6\xfd\x0f\x3f\xc1\x25\xa3\x36\xb3\xa8\x6a\x81\xfb\xb5\x3e\x7e\x78\x3f\xc7\x63\x92\x00\x40\x96\x45\x1f\x60\x72\xb6\xe3\x82\x42\x4a\xa8\x6c\x5d\xba\x28\x35\x24\x10\x6f\x15\x13\x72\xf2\x1e\xbd\x57\x2a\x03\x43\x4d\x82\xbf\x9e\x62\x81\xcf\x13\x0f\x69\x0c\x68\x00\x85\x84\x17\x78\x3b\xd6\x9d\x2e\xfd\xa5\x76\xc2\x4a\x2c\x47\x6c\xcb\xd4\x2a\xa6\x99\xd3\x7b\x43\xbc\xc0\xb2\x93\x6a\x19\xe7\x32\xc8\xee\xa5\xff\xd6\x52\x95\xac\x0e\x78\x77\x71\x12\x7a\x9a\x53\x3f\xd0\xa0\x0c\xda\x84\xa1\xa9\x3d\x0d\xd5\x8e\xea\x5d\x1a\x5f\x2f\xcf\x11\x7b\xa5\xb9\x65\xb6\x87\xcb\xa9\xc4\x60\xe3\xd3\xcc\x13\x2f\x90\xf5\x3c\x19\x8d\x20\x01\x31\x7f\x33\x70\xfb\x4f\x7a\xe8\xb5\x0d\xc9\xc7\x73\x3e\x31\x70\xc3\xe4\x77\x55\x81\x69\x47\x4c\xc6\xe7\x6f\xc7\xfb\x18\xbe\x87\xd9\xbc\xe4\x23\xc2\x3e\xbe\x6a\x63\x92\xf4\xab\x76\x02\x72\x3e\x71\x73\x75\x85\x56\x19\x5d\xcc\xce\x6e\xc2\x5a\x1b\x2b\x88\x5d\x5e\xd6\x7e\x52\x7d\x16\xa9\x9e\xa2\x71\x3a\x52\xd1\x09\xe1\x71\xe0\xf7\xbb\x11\xf6\x89\xc3\x24\x06\x3f\x69\x11\xc2\xd9\xd0\xe1\x3a\xbc\xce\x9e\x25\x0d\x3f\x62\x5d\xea\x8f\x20\xdd\xf5\xbe\x2e\xcf\x9f\xe7\x3b\x4a\xbc\x24\x27\x6c\x1f\xfa\xb2\x5e\xd6\x53\x82\x7f\x01\x00\x00\xff\xff\x1f\x03\x0a\x45\x4e\x04\x00\x00"

func burn_tokensCdcBytes() ([]byte, error) {
	return bindataRead(
		_burn_tokensCdc,
		"burn_tokens.cdc",
	)
}

func burn_tokensCdc() (*asset, error) {
	bytes, err := burn_tokensCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "burn_tokens.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x67, 0xdd, 0xb2, 0x81, 0x70, 0xa8, 0xa2, 0xe0, 0xaf, 0x3b, 0x70, 0x29, 0x4, 0x2c, 0x61, 0x40, 0x3c, 0x23, 0x9d, 0x71, 0xf6, 0xa4, 0x7c, 0x57, 0x29, 0x2b, 0x59, 0xa4, 0x46, 0xb4, 0x21}}
	return a, nil
}

var _create_forwarderCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6f\xe3\x36\x10\xbd\xf3\x57\xcc\x5e\x5a\x3b\x70\xe4\x7e\x5e\x8c\x6c\x01\x77\x93\x2c\x82\x16\x59\x20\x49\xdb\x63\x77\x4c\x8e\xcd\xe9\x4a\xa4\x40\x8e\xac\x35\x16\xf9\xef\x05\x49\x59\x91\x8d\x62\x03\xd4\x27\x4b\x9c\x79\xf3\xde\xbc\x27\x2e\x2f\x2e\x94\x7a\xb2\x1c\x41\x02\xba\x88\x5a\xd8\x3b\xe0\x08\x08\x42\x4d\x5b\xa3\x10\x6c\x7d\x48\x8f\x93\x73\xb1\x28\xa0\x7d\x57\x1b\xd8\x10\x74\x91\x8c\x12\x0f\x91\x04\xba\x16\xd0\x01\x6a\xed\x3b\x27\x20\x3e\x35\xf7\x18\x0c\x18\x6a\x7d\x64\x21\x03\xe2\x3f\x91\x8b\xe9\x0c\x9d\x17\x4b\x01\x02\x69\xe2\x3d\x85\x4a\xa9\xbb\x2d\xa0\x3b\x78\x47\x10\xc9\x99\x38\x2d\x4e\x73\xc2\xb7\x11\x6e\x0b\x22\x05\x78\x18\xfa\x16\x4a\x2c\x8d\x4f\xd0\x73\x5d\xc3\x3f\x5d\x94\x71\xb8\x58\x1f\x69\x82\x95\xca\xff\xc4\xae\x96\xa2\xc4\x62\x84\x0d\x91\x53\x49\x01\xc6\x7c\x1c\x48\x73\xcb\xe4\x04\xd0\x19\xa0\x86\xd3\x1f\xa0\x7d\x7a\x93\x9b\xd8\x19\xd6\x28\x14\x55\x6f\x59\xdb\xcc\xee\x38\x30\xa9\xb4\xc7\x81\xd5\xb0\xe0\x1e\x0f\x0b\xe0\xa4\x0f\xfc\x76\x7b\xa9\x2d\xb2\x83\x48\x61\xcf\x9a\xa0\x47\x27\x99\x5a\xe3\x1d\x8b\x0f\xd0\x5b\x9f\x6c\x18\x00\xd9\xed\xd4\x0b\x7d\x96\x05\xb0\x80\x46\x07\x3d\x8a\xb6\x85\x56\x3e\x8a\x44\xd0\x5b\x0a\x34\x21\x00\x1a\x1b\x82\x6d\xf0\x4d\xa5\xd4\xa3\x50\x3b\x54\x16\xb7\x8a\x55\x11\x7a\x16\x5b\x1a\x46\x15\x61\xa5\xd4\xf7\x15\x3c\x59\x82\xdb\xce\xed\x78\x53\x13\x3c\xe5\x0a\xed\x9d\x04\xd4\x69\x0b\x42\x61\x8b\x9a\x20\xda\x9c\x07\xac\x03\xa1\x39\xa4\x5c\x18\x6a\x6b\x7f\x20\x03\xd1\x37\x94\x49\xa9\x1f\x0a\x1a\xb6\x6d\xcd\x1a\x13\x9e\x9c\xe2\x0d\x28\x93\xee\x4a\xfd\x58\x9a\x26\x8e\x0c\xf1\x1a\x8a\x2d\xee\x09\x70\x30\x34\x85\x55\x72\x9e\x0b\x70\x20\x14\x32\x0a\x00\xb2\x91\x51\x7c\x20\x03\xec\x80\x25\xe6\x27\xdc\x51\xd1\x8e\xd0\x76\x9b\x9a\xa3\x25\x33\x66\x49\xfd\x54\xc1\x75\x26\x92\xf7\xf9\x31\xab\xbf\x1d\x3d\xa9\xb4\xd1\x1f\x5f\xc8\xe7\x94\x1a\xde\x6e\x29\x4c\x68\xaa\x9f\xab\x94\x59\x40\x70\xd4\xc3\xba\xbc\x5c\xc1\xbb\xcc\x2c\xc3\x1e\xf5\x38\x1f\x1a\xac\xeb\xc3\x22\xd3\x15\x4b\x0e\x42\xe7\xca\xe4\x22\xe4\xef\xd1\x9a\x32\x7a\xf2\x51\x96\xa6\x1d\x89\xb0\xdb\xc1\xc9\x07\x91\xac\x3f\x19\x54\x02\x7c\x16\xf4\x4a\x5d\x2c\x95\xe2\xa6\xf5\x41\x46\xbf\x8b\xdd\x19\xe0\xbb\xcf\xb7\x7f\xdc\xbf\xbf\xfb\xf5\xf7\x9b\xa7\x0f\xbf\xdd\xdc\xaf\xaf\xaf\x1f\x6e\x1e\x1f\x8f\x0d\x37\x9f\xb1\x69\xcf\xea\xff\xab\xee\x6c\x83\x23\xf4\x87\x87\xbf\xd6\x0f\xd7\x77\xf7\xef\x8f\xf5\x6a\xa2\x6d\x76\xbc\x21\x56\xb0\x36\x26\x50\x8c\x73\xf8\xa2\xb2\xe0\x36\x50\x8b\x81\x66\xa8\xb5\xac\x60\xdd\x89\x1d\x36\x9c\x2a\x60\xf8\xd5\x24\x93\xf8\xbc\x4d\x5b\x1a\xaa\x46\xe4\x79\xb5\x23\x79\x87\x2d\x6e\xb8\x66\x39\x5c\x7d\xf3\xe5\x64\x05\xd5\x71\x99\xcf\xbf\xcc\x96\x39\x27\x7a\x49\x13\xc9\xc7\xe3\xf9\x1b\x75\x32\x75\x9f\x43\x79\x75\x79\xae\xbb\x2a\x7e\xde\x53\x3f\x5e\x67\xb3\x91\xe1\xea\x85\xec\x7c\x44\x4b\x02\xab\x88\x7b\x9a\x5d\x5d\x66\xd4\x05\x88\x5f\xc1\x72\xc8\xf0\x09\x9b\x11\x73\xfe\xc2\x26\xdd\x3c\x09\xe2\x44\xe6\x2b\x5a\xaa\x8d\x0f\xc1\xf7\x5f\x5b\xc6\x1c\xde\xbc\x05\xc7\xf5\x64\xdb\x23\xdb\xce\xd5\xec\x3e\x7d\x7d\xc6\xd8\xf5\x7c\xaa\x34\x75\xfe\x5f\x13\x16\x20\x18\x76\x24\xaf\x6f\xa7\x0c\x7e\x56\xf0\x6f\x00\x00\x00\xff\xff\xf1\x62\x12\x90\x08\x07\x00\x00"

func create_forwarderCdcBytes() ([]byte, error) {
	return bindataRead(
		_create_forwarderCdc,
		"create_forwarder.cdc",
	)
}

func create_forwarderCdc() (*asset, error) {
	bytes, err := create_forwarderCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "create_forwarder.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x65, 0x7c, 0xbb, 0x72, 0xbf, 0x72, 0x8f, 0x77, 0xc4, 0xdb, 0xf8, 0x1e, 0xeb, 0x80, 0xe, 0x96, 0xd, 0x2a, 0xe6, 0x6c, 0x8c, 0x6d, 0xb3, 0xab, 0xc1, 0xf1, 0x80, 0x32, 0xb9, 0x79, 0x4e, 0xbf}}
	return a, nil
}

var _mint_tokensCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\xb6\x3e\x04\x12\xd0\xc8\x3d\x14\x3d\x18\x4e\x02\xb5\xb1\x8b\xa2\xad\x0b\xc4\x71\xef\x14\xb5\x56\x16\x95\x48\x62\xb9\xaa\x1d\x04\xf9\xf7\x82\x7a\xc4\xb2\x93\x5a\x17\x01\xe4\xcc\xec\xcc\x0e\xa9\x76\x96\x05\x96\x8d\x29\x29\xaf\xf0\xde\xfe\x41\x03\x5b\xb6\x35\x7c\xd8\x2f\x37\xab\xaf\xdf\x3e\xff\x58\xdc\xff\xfa\xbe\x58\x65\xb7\xb7\x77\x8b\xf5\x3a\xea\x09\x8b\xbd\xaa\xdd\x09\xfe\x08\x17\x09\x2b\xe3\x95\x16\xb2\x26\x66\xd4\xe4\x08\x8d\xcc\x20\x2b\x0a\x46\xef\xdf\x83\xaa\x6d\x13\x0e\x36\x4b\xda\x7f\xfa\x98\xc0\x53\x04\x00\x50\xa1\x80\x04\xd1\xac\xa8\xc9\xcc\xe0\x62\x3c\x28\x6d\x0f\xc9\x0b\x2b\xb1\x7c\x8c\xbf\x43\x8d\xf4\x17\x79\x06\x17\x4f\x47\x69\xd2\xe1\xe6\x39\x6a\x19\x8e\xd1\x29\xc6\xd8\x53\x69\x02\x3c\x6b\xe4\x21\xd3\x3a\x98\x19\x4c\x84\xcf\x63\xb5\x4d\x0f\x4e\xe0\x0a\x3a\xc2\x0b\x20\xcd\x2d\xb3\xdd\xcd\xcf\x38\xbc\x8e\xc3\x66\x66\x30\xf5\x62\x59\x95\x38\xc5\x11\xb4\x45\x26\xf0\xa2\x77\x73\x03\x4e\x19\xd2\xf1\x64\xdd\x0e\x02\xf2\x60\xac\x80\x3c\x60\x97\x10\x54\x60\x4c\x92\xe8\x0d\x8f\x43\x46\xb8\x82\x12\xa5\x8f\x73\x58\x7b\x72\x70\x5d\xa2\x7c\x51\x4e\xe5\x54\x91\x3c\xc6\x53\xd7\xe4\x15\xe9\x23\x63\x83\x56\xf2\xee\x75\xd6\xff\xad\xf6\x3a\x4e\xde\x08\xb2\x31\x2a\xaf\x82\x7b\xe8\xf8\xc0\x83\x4d\xc6\x2d\x32\x1a\x8d\x93\x8e\xd7\x77\x83\x7b\xd4\x8d\xe0\xa8\x86\xd0\x6f\x4d\x46\x90\x61\x7e\x79\x5a\x4a\xaa\x19\x95\xe0\x0a\x77\x3f\x5b\x48\xac\xaa\xca\xee\xb0\xc8\xfa\xa7\xd5\x3d\xb1\xe4\xb5\x58\xf1\x5b\x35\x95\x04\xc5\x4e\x3b\x0d\xbf\x36\x92\x8f\xd5\x09\xf9\xcc\xb6\xd3\x02\x9d\xf5\x24\x7d\xcd\xf3\xcb\x91\xf8\x88\x58\xa0\x17\xb6\x8f\xfd\xac\x3e\xef\xf3\xbf\x00\x00\x00\xff\xff\xdf\xdc\xa4\x3a\x7a\x03\x00\x00"

func mint_tokensCdcBytes() ([]byte, error) {
	return bindataRead(
		_mint_tokensCdc,
		"mint_tokens.cdc",
	)
}

func mint_tokensCdc() (*asset, error) {
	bytes, err := mint_tokensCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mint_tokens.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xaf, 0x1a, 0xe0, 0xd7, 0xca, 0xe1, 0x4b, 0x48, 0x6, 0x7a, 0x51, 0xb4, 0xb8, 0xbc, 0xdf, 0x82, 0xd5, 0x93, 0x2d, 0xa5, 0xa7, 0x58, 0xe5, 0x10, 0xf4, 0xea, 0x3b, 0xbc, 0xac, 0xe4, 0xb4, 0x70}}
	return a, nil
}

var _scriptsGet_balanceCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xcb\x4e\xf3\x30\x14\x84\xf7\x7e\x8a\xf9\xbb\xf8\x49\x36\x09\x0b\xc4\xa2\xa2\x54\xbd\x24\x08\x81\x8a\xd4\x0b\x7b\x27\x39\x69\x2d\x1c\xdb\x72\x6c\x5a\x54\xf5\xdd\x51\x2e\xad\x5a\xbc\xb2\xe4\x6f\xc6\x33\xe7\xc4\x31\xd6\x3b\x51\xa3\xce\xad\x30\x0e\x96\x78\x51\xc3\xed\x08\x19\x97\x5c\xe5\x84\x52\x90\x2c\xa0\x4b\x70\x05\x9e\xe7\xda\x2b\x77\x57\x23\x95\x7a\xbf\xd6\x5f\xa4\x30\xed\x38\xc6\x44\x65\xb4\x75\x48\xbd\xda\x8a\x4c\x52\xf7\x5a\x5a\x5d\xe1\xfe\x90\x6e\x16\x2f\xaf\xd3\xf7\x64\xfd\xf1\x96\x2c\x26\xf3\xf9\x32\x59\xad\xce\x82\xe4\xc0\x2b\xf3\x87\xbf\xe1\x98\xf1\x19\x4a\xaf\x50\x71\xa1\x82\x3e\xc3\x10\x93\xa2\xb0\x54\xd7\xe1\x10\x9b\x54\x1c\x1e\x1f\x70\x64\x00\x20\xc9\x35\x39\x1d\x46\xd8\x92\x9b\x74\xf4\x59\x15\x5e\x90\x6f\xee\xa5\x5b\x52\x89\x51\x4b\x47\x5b\x72\x33\x6e\x78\x26\xa4\x70\x3f\x41\x6c\x7c\x26\x45\x1e\xd3\x55\xb6\xbe\x68\xf8\x2f\xca\xb4\xb5\x7a\xff\xf4\xff\x3a\x79\xf4\xd9\x18\x1e\x6f\xda\x47\xbd\xe4\xf4\x1c\x74\x1f\x37\x67\x3c\x86\xe1\x4a\xe4\xc1\x60\xa6\xbd\x2c\xa0\xb4\x43\x67\x78\x1e\x25\x2c\x95\x64\xa9\xb9\x39\xdd\xee\xa2\xf5\x1e\x84\xac\x35\xb1\xe4\xbc\x55\x97\x02\x51\xbf\x28\x76\x62\xbf\x01\x00\x00\xff\xff\x98\x59\x27\xea\xcc\x01\x00\x00"

func scriptsGet_balanceCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_balanceCdc,
		"scripts/get_balance.cdc",
	)
}

func scriptsGet_balanceCdc() (*asset, error) {
	bytes, err := scriptsGet_balanceCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_balance.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0xea, 0x74, 0x86, 0xb3, 0xca, 0x1e, 0xe1, 0xf0, 0x6d, 0x6b, 0x81, 0xe1, 0x3, 0x5c, 0x5f, 0xd5, 0xdb, 0x17, 0x55, 0x8d, 0x59, 0x60, 0x42, 0x0, 0x79, 0xe4, 0xb1, 0x7f, 0xb7, 0x1d, 0xf3}}
	return a, nil
}

var _scriptsGet_supplyCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbd\x4a\xc5\x40\x10\x85\xfb\x79\x8a\x53\x26\x8d\x6b\x21\x16\x82\x85\x90\xd8\x08\x0a\x26\x3e\xc0\x9a\xec\x9a\xc5\xfd\x63\x76\x16\x22\x72\xdf\xfd\x42\x92\x5b\xa4\x9d\xef\x9b\x73\x8e\x52\x18\x17\x57\x50\x26\x76\x59\xc0\x46\xcf\x05\xb2\x18\x48\x12\xed\x51\x6a\xce\xfe\x0f\xd6\x19\x3f\x93\x52\x48\x76\x83\xfd\xaa\x43\xf6\x66\x4c\xbf\x26\xa2\x04\xcd\x82\x29\x45\x61\x3d\x09\x91\x0b\x39\xb1\x9c\x1d\xcb\x29\xe0\x7e\x1d\x3f\xde\xfa\xf7\x97\xae\xfb\xec\x87\x81\x28\xd7\x6f\xd8\x1a\x11\xb4\x8b\x4d\xfb\x84\xaf\x57\xb7\x3e\x3e\xe0\x9f\x08\x00\xbc\x91\x5b\xfd\xf3\x29\xec\x6e\x9b\x36\x6c\xe8\x50\xd3\x4f\xb3\xab\xed\x7e\x60\x23\x95\xe3\xf1\x4e\x97\x6b\x00\x00\x00\xff\xff\xce\x23\xa0\xa3\xe5\x00\x00\x00"

func scriptsGet_supplyCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_supplyCdc,
		"scripts/get_supply.cdc",
	)
}

func scriptsGet_supplyCdc() (*asset, error) {
	bytes, err := scriptsGet_supplyCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_supply.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa5, 0xda, 0x29, 0xbe, 0x3f, 0x6c, 0xe7, 0xd3, 0x2b, 0x69, 0x4f, 0x10, 0xe8, 0x42, 0x81, 0x51, 0xb7, 0x6c, 0x34, 0x33, 0xc6, 0x40, 0x28, 0x93, 0x3a, 0xf8, 0x9b, 0x98, 0xa5, 0xed, 0x1d, 0x52}}
	return a, nil
}

var _setup_accountCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\xe6\xf4\xc1\x06\xf2\x59\x3d\x07\x49\x00\xa7\x71\x8a\xa2\x45\x0a\x24\x6e\xef\x6b\x7a\x25\x11\xa1\x48\x82\x5c\x26\x36\x0c\xbf\x7b\x41\xfd\xb4\x56\x9b\x06\x3d\xa4\x3c\x18\x26\x77\x76\x76\x66\x56\x45\x59\x62\xdd\xe8\x08\x09\x64\x23\x29\xd1\xce\x42\x47\x10\x84\x5b\x6f\x48\x18\x95\x0b\xf9\xfa\xb3\x9e\x7b\xc4\x81\xb6\x5b\x10\xbe\x51\x32\x82\xc0\xd1\xa5\xa0\x38\xbf\x4b\xc3\x3a\x80\x94\x72\xc9\x4a\xc6\xc6\xfc\x46\x92\x0b\x7b\x28\xb2\x48\x91\xf3\x05\xbc\xa3\xd6\x1b\x5e\xbb\x47\xb6\x45\xa1\x5b\xef\x82\xe0\x36\xd9\x5a\x6f\x86\x57\x54\xc1\xb5\x78\xb7\xbb\xfd\x7a\xf7\xe1\xe3\xf5\xe7\xd5\xfa\xcb\xa7\xd5\xdd\xf2\xe6\xe6\x7e\xf5\xf0\x30\x36\xac\x4e\x58\x46\xfc\x04\x57\x9c\x7a\x3b\x14\x05\x00\xf8\xc0\x9e\x02\xcf\xa2\xae\x2d\x87\x73\x2c\x93\x34\xcb\x5e\xf2\x7c\xc4\xe4\xa3\x2b\xf4\x90\xc5\xc6\x85\xe0\x9e\x2f\xfe\x3b\x1d\xb7\xe8\xdc\x5f\xcd\xf2\xd4\x73\x94\x51\x5c\xa0\x9a\xcb\x53\x5f\x1d\x62\x8e\xcb\x4b\x58\x6d\x70\xf8\x41\x9c\x4f\x59\xe2\x7d\xe0\x9c\x31\xc1\xf2\xf3\x24\x8f\x21\x58\xb2\x5b\xf8\x24\xd0\x02\x6d\x31\x0c\x98\x90\x0c\xf2\x22\x3d\xf1\xec\xe2\xff\x89\x3a\xd5\x91\xaf\x5a\x2f\xfb\x8e\x6d\x36\x3f\x83\xb8\x57\x85\x16\x7f\x14\xe8\xd3\xc6\x68\x05\x45\x9e\x36\xda\x68\xd9\x0f\xbb\x1e\x84\x76\x1b\x76\xd6\xec\xc1\x3b\xef\x22\xc7\x5f\x89\x32\x74\xcb\xde\x45\x2d\xa8\x92\xed\xb7\x21\x4d\x70\xa9\x6e\xba\xe2\x3d\x2b\xd6\x4f\x1c\xa0\xad\x70\xa8\x48\xbd\xe8\xd3\x68\xfb\xf8\xd2\x12\x0e\x93\xef\x66\x31\x92\x1d\xaf\x66\x13\x96\x4e\x4c\x6f\x65\xe2\x7e\xc4\x9f\xfd\x86\x16\x0a\x35\xcb\x6b\xa1\x4d\x5a\xfe\x71\x82\x1b\x32\x64\x15\xa3\xd2\x6c\xb6\x93\xf8\xae\x87\xca\x5b\xa4\x37\x70\xfd\x6d\x78\x03\xfc\x0d\xb2\x1b\xff\x1d\x8b\xfe\xf7\x58\xe0\x7b\x00\x00\x00\xff\xff\x82\x41\xe3\xa0\xa2\x04\x00\x00"

func setup_accountCdcBytes() ([]byte, error) {
	return bindataRead(
		_setup_accountCdc,
		"setup_account.cdc",
	)
}

func setup_accountCdc() (*asset, error) {
	bytes, err := setup_accountCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setup_account.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x65, 0x8e, 0x6, 0xbd, 0x93, 0x70, 0xda, 0x4a, 0x7f, 0x65, 0xb6, 0xae, 0x53, 0x8a, 0x5c, 0x9c, 0x85, 0x8c, 0xb0, 0x94, 0xfe, 0x42, 0xc3, 0xff, 0xd3, 0xf2, 0x42, 0x3e, 0x4e, 0x10, 0xa0, 0x30}}
	return a, nil
}

var _transfer_tokensCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x6f\xd3\x40\x10\x3d\xd7\xbf\x62\xda\x03\x75\x24\x6a\x73\x40\x1c\xa2\x7e\x10\xda\xb4\x42\xa0\x22\xf5\x03\xce\x6b\x7b\x12\x2f\xd8\xbb\xab\xd9\x71\x93\xaa\xea\x7f\x47\xbb\xeb\x35\x76\x8a\x28\xa7\xc8\x93\x37\x6f\xde\xbc\x37\x9b\xe7\x70\x57\x4b\x0b\x4c\x42\x59\x51\xb2\xd4\x0a\xa4\x05\x01\x8c\xad\x69\x04\x23\xac\x34\xb9\xcf\xd1\xff\x5c\x0b\x4e\xf2\x1c\x4a\xdd\x35\x15\x14\x08\x9d\xc5\x0a\x8a\x47\x10\xea\x51\x2b\x04\xd6\x60\x51\x55\xc0\xfa\x17\x2a\xeb\x3e\x85\xd2\x5c\x23\x81\x28\x4b\xdd\x29\xdf\xec\x48\xa0\x16\x16\x0a\x44\x05\x16\x19\x3a\xe3\xa0\x84\x25\xca\x07\xec\x9b\xb3\x24\xcf\x13\xaf\x11\x61\x23\xb9\xae\x48\x6c\x40\xb4\x8e\x04\x84\x1b\x51\x63\x24\x85\x15\xe9\x16\xd6\xc8\x8b\x3f\x43\x36\x51\xa1\xc3\x19\x41\xa2\x45\x46\xf2\x92\x5c\x65\xb4\x54\x92\xc8\xd6\x68\x62\xb8\xec\xd4\x5a\x16\x0d\xde\xb9\xf9\x81\xf3\xdd\xf6\xf2\xfe\xfa\xea\xf3\xa7\xaf\xcb\xbb\x6f\x5f\x96\xd7\x8b\x8b\x8b\x9b\xe5\xed\x6d\x6c\x58\x6e\x45\x6b\x76\xf0\x13\x5c\x32\x1a\x93\x06\xed\x73\xb8\xbf\x94\xdb\x0f\xef\xdf\x02\xeb\x39\x2c\xaa\x8a\xd0\xda\x19\x3c\x25\x09\x00\x40\xbf\xef\x77\xd1\x35\x0c\x84\x56\x77\x54\x62\x6f\x98\x6e\x2a\x1b\xb4\xf7\xe6\xba\xaa\x20\x84\x02\xa5\x5a\x87\x8d\x56\x48\x84\x95\xa7\x6a\x90\x5d\x16\xec\xb9\xe6\xf0\x71\xb2\x5d\xe6\xab\x61\xa6\x21\x34\x82\x30\xb5\x72\xad\x90\xe6\xb0\xe8\xb8\xee\x8d\x1c\x74\xf5\xda\xae\x90\x41\x00\xe1\x0a\x09\x55\x89\xd1\xcc\xd0\x79\x68\xc1\xb2\x26\xac\xe0\xc1\x93\xc7\x3e\x27\xc4\x57\x6e\x70\x05\x27\x3d\x38\x2b\x34\x91\xde\x1c\xbf\x19\x7b\x18\x54\x9d\xa6\xce\xca\x39\xe4\x8e\x4d\xac\x31\xc7\x11\xc4\x23\x66\xc9\xde\xde\xde\xd9\x19\x18\xa1\x64\x99\x1e\x9c\xfb\xa8\x95\x66\x08\xa4\x2f\x05\xea\x4d\xd0\xe7\xbb\xf7\x0f\x66\x93\xa5\x7e\xc4\xe3\xea\x7d\xf5\x41\xbe\xbe\x96\xc5\x66\x95\x0d\x06\xc3\xf1\xd1\xb0\x64\x16\xcf\x75\x88\x3c\xfc\xce\x7c\xef\x73\x18\x8e\x5b\x2c\x3b\xc6\xbf\x18\xec\x46\x13\x96\xd2\x48\x54\x7c\x68\xc1\x74\x45\x23\xcb\xe1\xd6\x75\xf1\x13\xcb\xa9\xbb\x03\x1a\x4e\x46\xaf\x20\x65\x3d\xfb\x9f\xf4\xc6\xb3\x6e\xc2\x13\xa4\x5d\x7a\x5f\x0c\xf9\x0d\xf0\x6c\x8d\x7c\x2e\x8c\x28\x64\x23\xf9\x31\xcd\x83\xce\x49\x5a\x91\x6e\xb6\x3f\xe4\xfd\x34\xbd\xc3\x88\x78\x3e\x4d\x5f\x4f\x35\x40\xff\xbd\x81\x4f\x63\x27\xe1\x0b\x34\xda\xca\xe0\x6c\xcc\x46\xc5\xb8\xa5\x7a\xc1\x41\xbb\x2e\x8c\x1c\xc8\xaa\x40\xd6\x1f\xe9\xf1\xd1\xf4\x0e\x62\xc6\xcf\x09\xfc\x0e\x00\x00\xff\xff\xdb\x77\xc7\x8f\x60\x05\x00\x00"

func transfer_tokensCdcBytes() ([]byte, error) {
	return bindataRead(
		_transfer_tokensCdc,
		"transfer_tokens.cdc",
	)
}

func transfer_tokensCdc() (*asset, error) {
	bytes, err := transfer_tokensCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transfer_tokens.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb0, 0x32, 0xbb, 0xef, 0xec, 0x8e, 0x1b, 0x29, 0x87, 0xc3, 0xbd, 0x42, 0xe, 0x81, 0xf4, 0x6c, 0x4a, 0x40, 0xc4, 0x5c, 0xa5, 0x17, 0x20, 0x39, 0x19, 0x8c, 0x2a, 0xb3, 0x48, 0x33, 0x99, 0x60}}
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
	"burn_tokens.cdc":         burn_tokensCdc,
	"create_forwarder.cdc":    create_forwarderCdc,
	"mint_tokens.cdc":         mint_tokensCdc,
	"scripts/get_balance.cdc": scriptsGet_balanceCdc,
	"scripts/get_supply.cdc":  scriptsGet_supplyCdc,
	"setup_account.cdc":       setup_accountCdc,
	"transfer_tokens.cdc":     transfer_tokensCdc,
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
	"burn_tokens.cdc": {burn_tokensCdc, map[string]*bintree{}},
	"create_forwarder.cdc": {create_forwarderCdc, map[string]*bintree{}},
	"mint_tokens.cdc": {mint_tokensCdc, map[string]*bintree{}},
	"scripts": {nil, map[string]*bintree{
		"get_balance.cdc": {scriptsGet_balanceCdc, map[string]*bintree{}},
		"get_supply.cdc": {scriptsGet_supplyCdc, map[string]*bintree{}},
	}},
	"setup_account.cdc": {setup_accountCdc, map[string]*bintree{}},
	"transfer_tokens.cdc": {transfer_tokensCdc, map[string]*bintree{}},
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
