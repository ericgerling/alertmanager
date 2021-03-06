// Code generated by go-bindata.
// sources:
// template/default.tmpl
// DO NOT EDIT!

package deftmpl

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

var _templateDefaultTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1b\x6b\x6f\xdb\xb6\xf6\xbb\x7e\xc5\x99\x86\x8b\x35\x80\x6d\x25\xe9\x56\x2c\x0f\xe7\xc2\x75\x94\xc6\xb8\x8e\x1d\xd8\x4e\xbb\x62\x18\x06\x5a\xa2\x6d\xb6\x92\xa8\x89\x54\x9c\x6c\x77\xff\xfd\x9e\x43\xc9\x0f\xc5\x72\xe2\x06\x5d\xe2\xbb\xa5\xdd\xc3\x3c\xe2\x79\xf2\xbc\x48\x51\x7f\xfc\x01\x3e\x1f\x89\x88\x83\xfd\xeb\xaf\x2c\xe0\x89\x0e\x59\xc4\xc6\x3c\xb1\xe1\xcf\x3f\x1b\x34\xbe\xc8\xc6\x38\x91\x47\x3e\x02\xad\xb5\x28\x57\xbd\x36\x61\xe1\xf3\x9a\x7b\xa3\x79\x12\xb1\x00\x41\x08\x71\xbe\x75\xcc\x3c\xf5\xef\x84\x7b\x5c\x5c\xf3\xa4\x4e\x93\x7a\xf9\x20\xc3\xc9\xa9\x17\xc9\xab\x74\xf8\x89\x7b\x9a\xc8\xfe\x4c\x28\x7d\xcd\x74\xaa\xe0\xbf\xa0\xe5\x55\x1c\xcf\x50\xc5\x08\xf8\x6f\xf3\x87\xf6\x48\x24\x22\x1a\x13\xce\x21\xe1\x18\x2d\x54\xed\xcc\x40\x11\x35\xe0\xd1\x32\xc7\x5f\x80\x26\xbd\x4b\x64\x1a\xb7\xd9\x90\x07\xaa\xd6\x97\x89\xe6\xfe\x25\x13\x89\xaa\xbd\x67\x41\xca\x89\xe1\x27\x29\x22\xb0\x81\xa8\x42\xc6\x72\xac\xe1\x15\xd1\xaa\x35\x65\x18\xca\x28\x43\xde\xc9\x61\x4b\xf4\x76\x10\xe5\x15\xa2\x4c\x85\x9e\x14\x27\xa3\x05\x42\x79\xcd\x8b\xdc\x3b\x2c\x44\x86\x99\x19\xcb\xb8\xcf\x05\xdf\x99\xff\x5a\xb3\x36\x3e\x57\x5e\x22\x62\x2d\x64\x64\xdf\x63\x63\xcd\x6f\x74\xb6\x8e\xbf\x06\x42\xe9\x7c\x6a\xc2\xa2\x31\x4a\x86\x83\x4c\xae\x43\x6b\x01\x5c\xb5\x13\x59\xa5\x6a\x0c\x49\xe2\xd3\xa8\x0e\x73\x05\x72\xc1\x32\xe6\x8d\x28\x92\xb8\x4e\x28\x53\x81\xe4\x12\xf8\x71\x74\xfb\x32\x4d\x3c\x7e\x98\x2d\x26\x8f\x78\xc2\xb4\x4c\x32\xf7\xb3\x4a\x0c\x55\xb0\x81\x0a\x98\xf7\xb9\x86\x23\x96\x06\xba\xa6\x85\x0e\x78\x6e\x05\xcd\xc3\x38\x60\xba\xe8\x8b\xb5\x75\x26\x2f\xd2\x49\x15\x85\x40\x58\x46\xaa\x18\x68\x1b\xd2\x1b\xb1\x20\x18\x22\x60\x85\x5e\xa9\xf8\x44\x14\x1d\xe7\xa1\x89\x81\x88\x3e\x6f\x2c\x41\x9c\x70\x72\x16\x7b\xb3\xd9\x4b\xf4\xef\x35\x80\x49\x1b\x1b\x4a\x20\x3c\x19\x61\xcc\x7c\x12\x9b\xca\xb0\x22\x6e\x61\xe1\x27\x22\xf6\x26\x4c\x2f\x4c\x9c\xc8\xf0\xf1\xcb\x75\x97\x1a\xc6\xb1\x42\x94\xcd\x5d\xa9\x20\x5b\x4c\xdc\xfc\x54\xdf\xce\xe9\xad\xc6\xf3\x97\xb9\xe7\x2a\x45\x2f\x10\x3c\xd2\x8f\xd7\x78\x1d\xc5\x45\x25\x78\xdc\xa2\xaf\xd2\x15\x91\xd2\x2c\xf2\xb8\x2a\xa1\xbb\x92\xc0\xee\xb1\xaa\x8c\xd5\x98\x47\x82\x7f\x35\xa3\xae\x10\x54\x26\x11\x7d\xb9\xfa\x05\x31\x79\xc8\x44\xb0\x20\xb9\xa8\x82\x5f\x2c\x5f\x91\xd2\x44\x87\x01\x91\xb1\x8e\xbf\x39\xed\x36\x07\x1f\x2f\x5d\x20\x10\x5c\x5e\xbd\x6d\xb7\x9a\x60\x57\x1d\xe7\xc3\xeb\xa6\xe3\x9c\x0e\x4e\xe1\xa7\xf3\xc1\x45\x1b\xf6\x6a\xbb\x30\xc0\x24\xad\x04\x99\x88\x05\x8e\xe3\x76\xb0\x10\x4d\xb4\x8e\x0f\x1d\x67\x3a\x9d\xd6\xa6\xaf\x6b\x32\x19\x3b\x83\x9e\x73\x43\xb4\xf6\x08\x39\xff\x59\xd5\x4b\x98\x35\x5f\xfb\xf6\x09\x72\xae\x56\xad\xbe\xbe\x0d\x38\x30\x94\xd6\x30\xf1\x79\x82\x8d\x80\x0f\x14\x7f\x40\xa4\x15\xd2\x1e\x63\xbd\x4c\x87\x35\x4f\x86\x0e\xe9\x30\x4e\x23\xc7\x90\x63\x5e\x46\xaf\x6a\x54\xab\xce\xcc\xa1\xd0\x82\x83\x09\x87\x8b\xd6\x00\xda\xc2\xe3\x91\xe2\xf0\x0a\x07\x3b\x96\xd5\x94\xf1\x6d\x22\xc6\x13\x2c\xd8\xde\x0e\xec\xef\xee\x7d\x0f\x17\x19\x45\xcb\xba\xe4\x49\x28\x94\x42\x8a\x20\x14\x4c\x78\xc2\x87\xb7\x30\x46\x3e\x58\x82\x2a\x28\x10\xe7\x20\x47\x80\x71\x9d\x8c\x79\x05\xdb\x0e\x14\xfa\x16\xb0\xf3\x50\x88\x20\x87\x9a\x89\x88\xba\x0a\x06\x1e\xf2\xb0\x70\xa6\x9e\x20\x19\x25\x47\x7a\xca\x92\x4c\x43\xa6\x94\xf4\x04\x4a\xe8\x83\x2f\xbd\x34\xc4\xf8\x30\x85\x0e\x46\x22\xc0\xe2\xfe\x4a\xa3\xd0\x76\x3f\xc7\xb0\x77\x0c\x13\x9f\xb3\xc0\xc2\x92\x4f\xcf\x66\x8f\x4c\x03\x21\x53\x0d\x09\x57\x3a\x11\xc6\x0a\x15\x10\x91\x17\xa4\x3e\xc9\x30\x7b\x1c\x88\x50\xe4\x1c\x08\xdd\x28\xae\x2c\x24\x8a\x05\xa9\x62\xe4\xac\x40\x28\x7d\x31\xa2\xff\x73\xa3\x56\x9c\x0e\x31\x72\x26\x15\xf0\x05\x91\x1e\xa6\x1a\x81\x8a\x80\xc6\x8e\x15\xd2\xc3\x91\x09\x28\x1e\x04\x16\x52\x10\x28\xb7\xd1\x75\x21\x9d\x99\x43\xa2\xc7\x64\x50\x9d\x9b\x48\x11\x64\x3a\xc1\x55\x2d\x68\x22\x94\x35\x4a\x93\x08\x59\x72\x83\xe3\x4b\x34\x99\xe1\x48\xde\x4c\x10\x9a\x3e\x92\x41\x20\xa7\xa4\x1a\x66\x7d\x5f\xe4\x3d\x83\x59\x64\x36\xa4\xbe\xc9\x9b\xaf\x2b\x36\x0f\x28\x6a\x26\x02\x2d\x40\xbc\x58\xd5\xfc\x91\x9a\x60\xf9\x84\x21\xcf\x0d\x86\x7c\xd1\xbc\x6c\x49\x9d\x84\xd8\x53\x92\xd1\x82\x05\x10\x63\x0f\x42\xfc\xee\xaa\x59\x43\xfe\xe7\x2e\xf4\xbb\x67\x83\x0f\x8d\x9e\x0b\xad\x3e\x5c\xf6\xba\xef\x5b\xa7\xee\x29\xd8\x8d\x3e\x8e\xed\x0a\x7c\x68\x0d\xce\xbb\x57\x03\xc0\x19\xbd\x46\x67\xf0\x11\xba\x67\xd0\xe8\x7c\x84\xff\xb4\x3a\xa7\x15\x70\x7f\xba\xec\xb9\xfd\x3e\x74\x7b\x56\xeb\xe2\xb2\xdd\x72\x11\xd6\xea\x34\xdb\x57\xa7\xad\xce\x3b\x78\x8b\x78\x9d\x2e\xba\x70\x0b\x7d\x17\x89\x0e\xba\x40\x0c\x73\x52\x2d\xb7\x4f\xc4\x2e\xdc\x5e\xf3\x1c\x87\x8d\xb7\xad\x76\x6b\xf0\xb1\x62\x9d\xb5\x06\x1d\xa2\x79\xd6\xed\x41\x03\x2e\x1b\xbd\x41\xab\x79\xd5\x6e\xf4\x30\xb0\x7b\x97\xdd\xbe\x8b\xec\x4f\x91\x6c\xa7\xd5\x39\xeb\x21\x17\xf7\xc2\xed\x0c\x6a\xc8\x15\x61\xe0\xbe\xc7\x01\xf4\xcf\x1b\xed\x36\xb1\xb2\x1a\x57\x28\x7d\x8f\xe4\x83\x66\xf7\xf2\x63\xaf\xf5\xee\x7c\x00\xe7\xdd\xf6\xa9\x8b\xc0\xb7\x2e\x4a\xd6\x78\xdb\x76\x33\x56\xa8\x54\xb3\xdd\x68\x5d\x54\xe0\xb4\x71\xd1\x78\xe7\x1a\xac\x2e\x52\xe9\x59\x34\x2d\x93\x0e\x3e\x9c\xbb\x04\x22\x7e\x0d\xfc\xa7\x39\x68\x75\x3b\xa4\x46\xb3\xdb\x19\xf4\x70\x58\x41\x2d\x7b\x83\x39\xea\x87\x56\xdf\xad\x40\xa3\xd7\xea\x93\x41\xce\x7a\xdd\x8b\x8a\x45\xe6\x44\x8c\xae\x21\x82\x78\x1d\x37\xa3\x42\xa6\x86\xc2\x8a\xe0\x14\x1a\x5f\xf5\xdd\x39\x41\x38\x75\x1b\x6d\xa4\xd5\x27\x64\x52\x71\x36\xb9\x66\x55\xab\x98\x91\x4c\x0a\xbc\x09\x83\x48\xd5\x4b\x12\xdb\xde\xc1\xc1\x41\x96\xcf\xec\xcd\x26\x29\x4a\x6e\x75\x7b\x24\x23\x5d\x1d\xb1\x50\x04\xb7\x87\xf0\xdd\x39\x0f\xae\x39\x7a\x22\x83\x0e\x4f\xf9\x77\x15\x98\x03\x50\xd5\x04\x5d\x0e\xdd\x1f\x93\x5b\x15\x9b\x46\x31\x3a\x82\xa1\xbc\xa9\x2a\xf1\x3b\x3a\xff\x21\xfe\x4e\x30\x41\x56\x11\x74\x04\x86\x28\x3e\xc0\x4e\x77\xef\xfb\x18\x01\x21\x26\x26\x11\x1d\xc2\xee\x11\xe5\xd6\x09\x67\xfe\x73\xf2\x0f\xb9\x66\x40\x4d\x6f\xdd\xbe\x16\x7c\x4a\x51\x64\x53\xf4\x6a\x4c\x7a\x75\x7b\x2a\x7c\x3d\xa9\xfb\xfc\x1a\x03\xb2\x6a\x06\xcf\x67\x2c\x70\x66\xe2\xd2\x62\x56\xf9\x6f\xa9\xb8\xae\xdb\xcd\x4c\xd4\xea\xe0\x36\xe6\x4b\x82\x53\x87\xe1\xd0\xe2\x1e\x99\x4a\xa0\xb8\xae\x5f\x0d\xce\xaa\x3f\x3e\xb3\xf8\xa6\xc3\x7e\xbe\xe5\xbe\xaf\x17\x39\x76\x8c\x70\x27\x96\x75\xec\x90\x53\xd2\x8f\xa1\xf4\x6f\x41\x20\x8a\xc2\x9c\x8b\x12\xdb\x66\xa0\x6f\xe9\x77\x1e\x51\xca\x9b\x60\x55\x37\x11\xe5\x52\x75\xbf\x98\xb5\xd0\x4f\xaa\x64\x75\xca\x87\x9f\x05\x32\x32\x0f\x42\x29\xb1\xa6\x10\x52\x56\x1b\x04\x53\xdc\x5f\x4c\x22\xdf\x30\xd8\x55\xe6\x7f\x4a\x95\x3e\xc4\x8a\x13\xf1\x23\x6c\x25\xa8\x32\x21\xc9\xdd\xdd\x7f\x1d\x61\x51\x8e\x78\x75\x0e\xaa\xbd\xe1\xe1\x11\x98\x08\xc8\x26\xc0\x37\x22\xa4\x60\x41\x0e\x28\x27\xee\x60\xc6\x89\x4c\x23\xbf\xea\xc9\x40\x26\x87\xf0\xed\xe8\x0d\xfd\x5d\x36\x3f\xc4\xcc\xf7\x8d\x54\xe4\x0d\xc3\xb1\x99\x59\xb7\xf3\x99\x36\xd9\x5b\xb3\xe1\x53\xbb\xc7\x92\x4a\x1b\xea\x51\x2a\x3b\xc0\xb1\x4e\x9e\x31\x8f\x01\x90\x04\x4f\x9c\x49\xaf\x71\x7b\x80\x44\x82\x2a\xba\xd8\x18\x25\xd1\x32\x2e\x1a\xea\xda\x3c\xc0\x6c\x24\x63\xfb\x04\x03\xcc\x5f\x08\x9a\x65\x56\xfb\xcd\xee\xee\x13\x87\x4a\xa9\xd0\xd8\x45\x62\x56\x40\xb6\xc3\x40\x7a\x9f\x0b\xbe\x1d\xb2\x9b\x6a\xee\x24\x28\x6c\x7c\x53\x78\xe8\x05\x9c\x25\xc4\x50\x4f\x0a\xf0\x75\x81\x32\x37\x0e\xb0\x54\xcb\x3b\x21\x51\xb0\x96\x31\x14\x9a\xca\x17\xd7\x4f\xed\x56\x45\x7d\xef\x1a\xe7\x7e\x25\x66\x72\xd3\x22\x9b\x60\xce\xd7\x99\x2c\x81\xe5\x09\xbb\xf1\x7c\x76\xdd\xde\xcd\xc6\x2a\x66\xde\x6c\xfc\xa4\x8a\xe6\x0f\x13\xe6\x8b\x54\x1d\xc2\x6b\x03\x2b\x49\x00\xa3\x51\x21\x8b\x65\x68\x48\x04\x5d\x41\xc9\x40\xf8\xf0\x2d\x3f\xa0\xbf\xc5\xc4\x30\x1a\x2d\xd9\x62\x1b\xb2\xc3\x42\x92\xa7\xcb\x12\x6f\xd6\x06\x5c\xc1\xba\x06\x65\x9a\x97\x9a\x1f\x76\xd1\xc8\xa6\x44\xe5\xf3\x71\x43\xa7\x79\x52\xb6\x5e\xe6\xdf\x5d\xb3\x28\xab\xeb\xe6\xbe\xf9\x61\x7f\xbf\x59\x5e\x80\xf6\xc9\xaf\x6d\xc8\xe3\x2d\x63\xb0\xbc\x7a\x19\x6e\x79\x44\xce\xfe\x2c\x0e\xea\xe7\x27\xf4\x60\x0e\x4c\xee\x9c\xb5\x67\x73\x76\x60\x0f\x27\xa8\xf9\x81\x07\xea\x9c\xc0\xe2\x30\x79\xcd\x61\x3e\x9d\x7b\x00\xac\xf2\xcd\x8f\x96\xeb\x85\x83\xe5\x95\x69\xf9\xd1\x4a\x61\xf1\xe7\x39\x78\x3e\x4e\x5e\xdc\x74\x93\x62\xb6\x70\x9e\xbd\xcc\x79\xee\xf3\x8d\xad\xcf\x7d\x6b\xcd\xbe\x5d\x4e\xb0\xed\xae\x80\xb9\x67\x96\x4b\xee\x73\x87\x5c\x0d\xdc\xb8\x25\x7c\x54\xb7\x37\x39\x65\x7d\x62\x7f\x98\x25\xcd\xb3\xb3\xb3\x3c\xf9\xfa\xdc\x93\x89\x39\x93\x9b\x6d\x0f\x0a\x1b\x82\x7d\xda\x0e\x14\xf2\xf6\x50\x06\x7e\x79\xe2\xf6\xd2\x44\x11\xf5\x58\x8a\x0c\x30\x6f\x28\x44\x64\x88\xe6\x7d\xc5\x9d\x04\xff\x03\x09\x66\xe8\x99\x43\x54\x4c\x98\x21\xd2\x64\xb1\xd0\x48\xff\x77\x5e\x9a\xf4\x5f\x7f\xff\x23\xf7\x59\x49\xbd\x5e\x99\x91\x83\x8d\x95\x0f\xb3\x42\x3e\x07\xce\xbb\x37\x2c\x2f\xd9\xf2\x9e\xbc\x17\x7c\x4a\xe7\x6f\x0f\xbe\x75\x38\x76\x58\xa9\x0f\xdf\x49\xbc\xe5\xe9\x37\xfb\x53\x5a\x40\xf2\xb7\xc1\x3b\xe8\x71\x25\x45\xe1\x25\x64\xff\x9a\x90\x55\x3a\x91\xd1\xf8\xf9\x4c\xfb\xf3\xfa\xeb\x00\xbf\x40\x06\x38\x76\x32\x21\xbf\x82\xd7\x95\x34\x0c\xf9\x93\xd9\x3b\xef\x82\x24\x2f\x7e\xf8\x8f\xf1\xc3\xac\x35\x9d\xbb\xda\xf1\xf0\xf9\x96\x99\xce\x11\xcb\x6c\xf4\xc0\x65\x8f\xf5\x37\x32\x9e\x59\x99\xf5\x71\x97\x6b\x55\xa8\x05\x8b\x4b\x27\x59\x25\x78\x76\xcf\x58\x92\x68\x5b\xdc\xe3\x41\x8b\x3e\x78\x83\xe7\xff\xd4\x59\x96\x3b\xcc\xbb\x57\x8a\x9e\xa9\xa1\x9c\xb5\x5b\x2b\x3d\x25\x76\x6d\x3c\xa1\xee\xaf\xe8\x4e\xd9\xa5\x28\x6a\xa2\xb6\x2f\xc7\x3c\xae\x9a\x6e\xd8\xde\xf5\x38\xf6\xa0\xd7\xdc\x5f\xd3\xe0\xbd\x74\x85\x5b\x54\x8d\xb7\xce\x33\x51\xa6\xc9\x16\xca\xb4\x75\x76\xfa\x92\x08\xbe\xaf\x23\x7e\x09\xac\xbf\x67\x9b\xbb\xbc\xdd\x9a\x25\xe4\xa5\x0d\xd7\x0c\xf4\x0c\x5b\xae\xb9\x34\x2f\xde\xf8\x8f\xf1\xc6\x97\x4d\xd7\xcb\xa6\xeb\x65\xd3\xb5\xed\xce\xf2\xb2\xe9\xda\x9a\x96\x6d\xdd\x42\xe1\x6c\x7a\x1f\x77\xf2\x05\xaf\x42\xe7\x28\x0b\xc8\x93\xdf\xc4\x28\x5c\x4d\x5a\xba\x69\xb2\x58\xe8\x83\x83\x83\xfb\x5e\x70\x17\xdf\xec\xae\xbe\x92\xdc\x8e\xa6\x61\x9b\xda\x97\xa7\x6c\x5d\xf6\xd7\xb6\x2e\xa5\x2f\xd1\x1e\x5a\xf2\xa5\xde\xe6\xce\xbd\x86\xe2\x2d\xac\xe5\x74\x55\xfc\xe8\xf1\xe9\x1c\x62\x7f\x39\x5b\x19\x8d\x36\x4e\x55\xa8\x13\x0c\x6f\x37\x7b\x0f\xb7\x9a\x3b\x56\xee\x3b\xdc\xcd\x0c\xc7\x0e\x86\xf9\x49\xf6\x5f\xab\x98\x26\xb6\xad\xad\x5d\x73\xbd\x2e\x53\x71\x91\xbf\x8e\x1d\xba\xc5\x4a\x10\xba\x0e\x7c\x62\x59\xe5\x5f\x55\xc6\xa9\x9a\x48\xe4\xf8\x15\x3e\x2a\x5c\x21\x55\xfc\xac\x2c\xff\xa0\x74\x4d\x23\x50\xfa\x05\xa9\xb5\xc1\xe1\x57\x7e\x17\x27\x83\x99\x0f\x35\xef\xff\xe2\x6a\xe5\x5d\x56\x41\x97\x87\x4f\xe8\x72\x7e\x33\xe8\xe6\x1c\x97\xb7\x72\x05\x9e\x1b\x58\x32\x4d\x82\x47\x7c\xa7\xb5\x44\x91\x4d\x95\x42\x83\x97\x7f\xa7\xf5\x57\x7c\xac\x9c\x7f\x0c\xfb\xf7\xfb\x5a\xf9\x8e\x25\x5f\x7c\xfc\x4b\x7c\xfc\x7f\x01\x00\x00\xff\xff\xcc\x0e\x3c\xae\xc5\x3f\x00\x00")

func templateDefaultTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateDefaultTmpl,
		"template/default.tmpl",
	)
}

func templateDefaultTmpl() (*asset, error) {
	bytes, err := templateDefaultTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/default.tmpl", size: 16325, mode: os.FileMode(420), modTime: time.Unix(1466186597, 0)}
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
	"template/default.tmpl": templateDefaultTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"default.tmpl": &bintree{templateDefaultTmpl, map[string]*bintree{}},
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
