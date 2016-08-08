// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// generated by build/assets.sh; DO NOT EDIT

// Code generated by go-bindata.
// sources:
// pages/assets/html/containers.html
// DO NOT EDIT!

package pages

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

var _pagesAssetsHtmlContainersHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x5a\xdd\x73\xe2\x38\x12\x7f\x26\x7f\x45\x9f\xeb\x1e\xf6\xaa\xc6\x66\xbe\x5e\x6e\x8e\x50\xc5\x92\x99\x1b\x6e\x33\x49\x2a\x24\xbb\xb5\x8f\xc2\x16\xa0\x89\xb0\xbc\x92\x1c\xc2\xa5\xf2\xbf\x5f\x4b\xb2\x41\xfe\x82\x7c\xd5\xcc\xb1\x35\x1b\x6c\xa9\x7f\xfd\xeb\x56\x77\xab\x65\x33\xf8\x5b\x18\x1e\x01\x8c\x45\xb6\x91\x6c\xb1\xd4\xf0\xfe\xed\xbb\x8f\xf0\x6f\x21\x16\x9c\xc2\x24\x8d\x23\x18\x71\x0e\x97\x66\x48\xc1\x25\x55\x54\xde\xd2\x24\x3a\x42\x91\x53\x16\xd3\x54\xd1\x04\xf2\x34\xa1\x12\xf4\x92\xc2\x28\x23\x31\xfe\x29\x46\xde\xc0\xef\x54\x2a\x26\x52\x78\x1f\xbd\x85\x5f\xcc\x84\xa0\x18\x0a\xfe\xf1\x2f\x44\xd8\x88\x1c\x56\x64\x03\xa9\xd0\x90\x2b\x8a\x10\x4c\xc1\x9c\xa1\x62\x7a\x17\xd3\x4c\x03\x4b\x21\x16\xab\x8c\x33\x92\xc6\x14\xd6\x4c\x2f\xad\x9a\x02\x24\x42\x88\x3f\x0b\x08\x31\xd3\x04\x67\x13\x9c\x9f\xe1\xd5\xdc\x9f\x07\x44\x1b\xbe\xe6\xb3\xd4\x3a\xfb\xd4\xef\xaf\xd7\xeb\x88\x58\xae\x91\x90\x8b\x3e\x77\xf3\x54\xff\x74\x32\xfe\x7c\x36\xfd\x1c\x22\x5f\x23\x71\x9d\x72\xaa\x14\x48\xfa\x57\xce\x24\x1a\x3a\xdb\x00\xc9\x90\x4d\x4c\x66\xc8\x91\x93\x35\x08\x09\x64\x21\x29\x8e\x69\x61\xd8\xae\x25\xd3\x2c\x5d\xbc\x01\x25\xe6\x7a\x4d\x24\x45\x94\x84\x29\x2d\xd9\x2c\xd7\x15\x57\x95\xdc\xd0\x62\x7f\x02\x3a\x8b\xa4\x10\x8c\xa6\x30\x99\x06\xf0\xeb\x68\x3a\x99\xbe\x41\x8c\x3f\x26\x57\x5f\xcf\xaf\xaf\xe0\x8f\xd1\xe5\xe5\xe8\xec\x6a\xf2\x79\x0a\xe7\x97\x30\x3e\x3f\x3b\x99\x5c\x4d\xce\xcf\xf0\xea\x0b\x8c\xce\xfe\x84\xdf\x26\x67\x27\x6f\x80\xa2\xa3\x50\x0d\xbd\xcb\xa4\xe1\x8f\x24\x99\x71\xa2\x59\x37\x80\x29\xa5\x15\x02\x73\xe1\x08\xa9\x8c\xc6\x6c\xce\x62\xb4\x2b\x5d\xe4\x64\x41\x61\x21\x6e\xa9\x4c\xd1\x1c\xc8\xa8\x5c\x31\x65\x96\x52\x21\xbd\x04\x51\x38\x5b\x31\x4d\xb4\xbd\xd3\x30\x2a\x3a\x0a\xc3\xe1\xd1\xd1\x60\xa9\x57\x7c\x88\x93\x07\x4b\x4a\x92\xa1\x5d\x82\x81\x66\x9a\xd3\x61\x3c\x4a\x6e\x99\x42\xcd\x21\xdc\xdf\x47\x27\x4c\x65\x9c\x6c\xce\xc8\x8a\x3e\x3c\x0c\xfa\x6e\x8a\x9b\x8e\xd1\x09\xa7\x44\x53\xa5\x6d\x24\x60\x6c\x24\x86\x01\xac\x58\x8a\x64\xf1\x62\x3c\x9d\x82\xd1\x66\x67\x73\x96\xde\xe0\x72\xf1\xe3\x40\xe9\x0d\xae\xdd\x92\x52\x1d\xc0\x52\xd2\xf9\x71\x80\x7a\x2e\x85\xd0\x0f\x0f\xca\xf0\x8e\xfb\x33\xbc\x40\xbf\x93\x2c\xfc\x10\xbd\xc3\xff\x10\x31\x8a\x95\x0a\x86\x47\x3b\xcd\xe7\x99\xb1\x90\x70\x63\xdc\x8a\xbe\x54\x8f\x05\xe9\xd0\xf6\x14\xc4\x58\xa4\x26\xd8\x31\xb7\x1a\x84\xf7\xba\xea\x3f\xe4\x96\x4c\x63\xc9\x30\xb1\xb6\x96\x28\x77\xad\x64\xdc\xd4\xf3\xfd\xaf\x9c\xca\x4d\x88\x74\xdf\x46\xef\x2d\xe3\xef\xa8\x6d\xd0\x77\x32\x8f\x00\x68\x73\x71\x37\x84\xde\x64\xf4\x38\xd0\xf4\x4e\xf7\xbf\x23\x53\x77\x37\x68\x47\x5e\xd8\xfa\x14\x7e\x57\x24\x63\x35\xc8\x67\x63\x7a\x6e\x7d\x25\x92\xf1\x92\x48\xdd\x44\x1b\xf4\xcb\x7c\x18\xcc\x44\xb2\x29\x14\x24\xec\x16\x62\x4e\x94\x3a\x0e\xb6\x4c\x5c\xdc\x85\x6a\x29\xd6\x31\xc1\xaa\x09\xc3\xa2\x8e\x0d\x48\x3d\x36\x82\x9d\x30\x0f\xd5\x2a\x7c\xf7\x3e\x00\x96\x1c\x07\x5c\x2c\x44\xb0\x15\xeb\x93\xed\xd7\x8a\xbe\x52\x64\x78\xd4\xf3\x07\x32\x2c\x03\xa1\x21\x4b\xa5\x19\x32\x99\xfc\x6e\xd8\x4c\x58\xbc\x89\x72\x7d\x14\x34\x7f\x05\x2f\xc5\x67\x12\x45\x63\x99\xaf\x66\x4e\xfa\xfe\x5e\x62\x6d\xa1\xf0\xf7\x0c\x2b\x63\xaa\xc7\x5b\x33\x3f\x1d\x43\x74\x51\xbd\xa7\x1e\x1e\xac\x42\xce\x86\x9e\xb1\x75\xc9\xe8\x14\xf3\x06\x8d\x1f\xb6\x0c\x5d\xe1\x22\x19\x76\x04\x9d\x8f\x28\x8e\x00\x4d\x13\x03\x3c\xe8\x0b\xbe\x73\x8a\x25\xee\x2e\xee\xef\xd9\x1c\xa2\x89\x72\x4e\x3d\xe0\x2b\x28\x3e\x83\xe5\xc7\x1d\xc9\x28\xea\x27\x22\xbe\x31\x1e\x3b\xb1\x7f\x61\x67\x93\x23\x83\xb3\xdb\x55\x3b\x72\x36\x91\xdb\xd8\x6c\xb3\xf6\xe9\x8c\x96\x18\xb6\x48\xe8\xab\xf9\x03\x17\x22\xd9\xcf\xc4\x32\x28\xe8\x78\x4a\x1d\x9b\x69\x3e\x8b\xfd\x45\x7a\x59\x38\x7d\x18\x56\xf0\x90\xd2\x07\x3f\x96\x3c\x61\x8e\xbb\x64\xb8\x90\x22\xcf\x6a\xc1\xa4\x3c\x00\x1b\x49\x75\x86\xbd\x4a\xbe\x54\xe6\x97\xf1\xd3\x54\x12\x32\x4d\x57\x36\xae\x2a\xf3\x77\x41\x55\x8b\x27\xdf\x7d\x6d\xab\xea\x79\xd0\x85\xc5\x14\x8b\x44\xfe\x1a\x0e\x3c\x91\x0c\xf7\x69\x70\x78\x75\x07\xe6\xfc\xa0\xff\x5c\xb4\x2a\x2b\x6e\xfd\x57\xe3\xe7\xb2\xd0\xc1\x40\x8b\x8b\x06\x2a\xc3\x96\xa5\xd0\x62\x60\x42\x4e\x66\x94\x5b\xdf\xf9\xd8\xd1\x6f\x74\x63\x5c\x67\xa6\x0f\xa1\x3e\xf8\x3b\xe1\xb9\x2d\x26\xf5\x54\xad\x7a\xcd\x19\xbb\xe3\xd6\x7b\x1e\xb5\xa9\x16\x12\x7d\x39\x98\xc9\x61\x41\xc8\x40\x75\x39\xab\xb7\xf3\x95\x55\xdf\xf0\x55\x37\xab\xa7\xfa\xcb\xc3\x6f\xfa\xcb\x1f\xac\xfa\xab\xb7\x75\x17\x2e\x7d\xce\xad\x35\xa5\x27\x8b\x1b\x5d\xd1\x5a\x09\xd0\x8a\xaf\x27\x2b\x74\xd1\xe1\x08\x85\xed\xa7\x3b\x54\xc1\xfb\x98\x98\x75\xd0\x2e\x58\xbd\x11\x9f\x97\x43\x33\x5b\x98\x8b\x93\x90\x59\x19\xb3\x95\x56\x66\x99\x25\xc4\x7f\x47\x6d\x18\x6d\xb6\xe1\x19\x46\xe4\x32\xa6\x6a\x74\x4b\x18\x37\x9d\xfc\x2b\xe4\xe0\x44\x09\x6e\xbb\xe1\x5a\xfe\x39\x95\xe3\x2c\xf7\x95\x75\x06\x9a\xe7\x89\xce\xf8\x01\x12\x6b\x0c\x03\x3c\x37\x14\x1a\x43\xdb\x2e\x03\x06\x09\xe5\xee\x7b\x30\x1c\x5f\x5c\xbb\xe5\xdf\x21\x16\xc5\x1b\x9b\x7c\x43\x07\xeb\x1e\xf6\xef\x5b\xc3\xf7\xab\xdc\x97\x47\xd8\xe1\x98\x75\x2c\x63\x34\x93\x2c\xd5\xee\x66\x53\x19\x54\x60\xf2\x94\x6d\x61\x94\x0f\xd3\x64\xee\x2f\x62\x8b\x2d\xdf\xc8\xdd\x2b\x99\x83\x48\x60\xa1\x6a\x16\x8d\x45\xd5\xa0\x9d\xc6\x6e\x9b\x62\xf1\x22\x93\xd4\xcd\xcb\xcd\xc1\x93\xbb\x58\x9b\x33\x92\x68\x2e\x92\xd1\x50\x53\x08\xf8\xff\x78\x89\xdb\xdc\x24\x9d\x8b\xe8\x2c\x5f\x59\xb9\xb2\xc6\x34\xd9\x97\xa5\x66\x7b\xed\x8c\xf8\x46\x57\x42\x6e\x7e\x6c\xc0\x3b\x9d\x7b\x62\xde\x4d\x88\xdc\x03\x0c\x0b\xf3\x72\xf7\x7a\x60\xf5\x0c\x60\xff\xa5\x7b\x14\x77\x07\x4d\x21\x7f\x8d\xb7\xf6\xc8\x3f\x27\xaa\x0a\x9c\x57\x4a\x94\xb6\x24\x69\x1a\x7d\x30\x47\x3a\xcd\x2d\x24\x5f\x60\xe8\x74\x4d\xb2\xd7\x2a\x72\x08\xd5\x5a\x16\x9a\x16\x7b\x5a\x9f\x61\xb5\x27\x7d\xc0\xf2\x7a\xea\x15\xd6\x55\xba\xd0\x67\x6f\x66\xd7\xca\xb4\x46\xdd\x9d\xb8\xcd\xbc\x22\xff\xd0\x92\x15\x91\x9b\x3d\x6d\x80\x99\x65\x34\xb0\x74\xd1\x6c\x04\xaa\xd3\x8a\x64\x3e\xc7\x2e\xe7\x96\xd1\xf5\xfe\xf6\x00\xbc\x0e\x21\x37\x8c\xc3\x05\xc9\x17\x34\xa8\x42\x9a\x03\xf6\xb6\x65\xf8\x29\xd6\x5c\x48\x81\xcd\x86\x3a\xd4\xed\xf8\xe6\x64\xa5\x48\xa8\x45\xf6\x28\x83\x3a\xfa\x8c\x1f\x68\xa6\x6d\x39\x1e\x63\x60\x8b\x35\x35\x05\x1f\x87\x57\x42\x13\x0e\x65\x1c\x7e\xb4\x91\xe9\xf9\x27\xce\x72\xf4\x0c\x4e\x09\xdd\xc2\xdb\xe7\x2c\x3b\xa7\xd8\xb9\xe6\x08\x6b\xa0\x90\x17\x9c\x0a\x92\xc0\x08\xa3\x6a\x0f\x1e\xc7\x39\x55\x20\xef\xe4\xbb\x63\x66\x39\x99\xe7\xa1\x76\x53\xed\x02\xc3\xf1\xd0\xec\xff\xad\xfc\xda\x21\x7f\x95\x94\xdc\x24\x62\x9d\x76\x61\x3a\xa8\x59\x39\xad\x13\xb4\x19\x1a\x07\x77\xe7\x1f\x18\x26\xe5\x46\xfd\x83\x22\x65\x65\xd5\x1d\x5e\x86\x99\xec\xd7\xee\x78\x04\xa4\x58\x43\xfb\x81\x67\xef\x12\xd6\xa6\x35\xcb\xf1\x3f\xed\xd9\xb2\x62\xaa\x14\x0b\xf3\xc8\xbe\xa1\xa4\xe1\x93\x62\x62\x38\x23\x12\xfc\x8b\x30\x31\x07\x55\x19\x94\x75\xc4\x0d\x2c\x85\x0e\x9d\x2b\x5a\x91\xa1\xba\x57\x29\x19\x8a\x94\xe3\xd4\xaf\x42\x43\xb9\x60\xee\x90\xdc\x22\xd9\xf4\xe6\x53\xe8\x32\x6c\x35\x6b\x64\xd1\x3b\xc9\x73\xd8\x8e\x51\xee\xb1\x74\x7b\xbd\x56\xde\xed\x37\x9b\x2b\xf7\x21\xf0\xa3\xcb\x3c\x0d\xae\x55\x9f\x27\x26\xe5\x19\xd5\x6b\x21\x6f\x9e\x98\x95\xbd\x97\xa7\x63\xa1\xb8\xd8\xec\x9f\x92\x88\xbd\xfa\x68\x22\x45\x66\x82\xbf\x99\x20\xb3\x5c\x6b\xb1\x5d\xaf\x99\x4e\x01\xff\x85\x09\x9d\x93\x9c\x6b\x28\xe5\xb0\xa2\x2f\x16\xc8\xa9\x78\xc4\xee\x84\x9c\x9f\x53\xc7\x32\x54\x94\xd3\xd8\x1e\x01\xb6\xca\x20\x21\x9a\x14\xa2\x1e\x07\x20\x92\x91\x70\x49\x54\x26\xb2\x3c\x3b\x0e\xb4\xcc\x69\x71\x93\xde\xa1\x1d\x09\x45\xd8\x39\xe1\x8a\xb6\x84\x98\x0b\xaf\x76\xc5\xe5\x5a\xb7\xc7\x57\x25\x30\x63\x3c\xd3\x7a\x73\x7b\x65\x24\x38\xcb\x1a\x5e\xc2\x23\x52\xab\xca\xa0\xee\x60\xcc\x8d\x34\x0f\x40\x0a\x63\xb1\xfb\x6e\x0d\xb3\xdd\x25\xa7\xc9\x6c\xb3\xd7\x63\xcd\x98\x2f\x1e\x0f\xed\x09\xdb\xa7\x14\xe4\x25\xb6\xd4\x8b\x65\x96\xeb\x66\x15\xdc\x96\xe5\x92\xde\x6c\xa3\xb1\xc9\x69\x6c\xdf\xcf\x50\xfb\x59\x4a\x61\x1f\x1f\x37\xb6\x80\x52\x17\xb5\x33\xba\x95\xd5\x8c\xaf\x65\xe8\x17\xf5\xd3\xb6\xcc\x2f\x8c\x53\xb5\x51\x78\x46\x79\x7c\x07\x39\xdf\xca\xb8\xbd\xaf\xb5\x89\xec\x46\xea\x28\x53\xe3\x5c\x69\xb1\xfa\x46\xb5\x64\xf1\x53\xfd\x71\xa0\x58\xf5\xf6\x79\x60\xe4\x5e\xba\x9b\x38\x86\x42\x7b\xbd\x62\xed\x8b\x95\x5a\x2f\x65\x8d\xc0\x24\xb2\x38\x07\xe3\xa1\x57\x3f\x6c\xb6\xbc\x05\xf9\x69\xa1\xd1\xf2\xee\xe4\x50\x74\x3c\xae\xa9\xca\xc0\xf4\xcd\xb6\xad\xf9\x54\xaf\x17\x2c\xc5\xe4\xae\xb4\xba\xfe\x1b\x92\x30\x71\xef\x06\x71\x1b\xcf\x53\x1d\xb4\xee\xdf\xdb\xad\xbb\x4d\xce\xc2\x77\xc8\xdd\x9a\x87\xde\xc7\xef\xde\xd6\x28\x77\x17\x9a\x56\x86\x95\x6e\xb0\x86\xd4\x5e\x00\x9f\xe9\x43\xd7\x8c\x1c\x74\x63\xd1\x46\xfc\x7f\x7a\xb2\xd2\x6a\x39\x2d\xb8\xf3\x70\x4f\xcd\x8c\x8b\xf8\xa6\xee\x81\xe6\xfe\x58\xef\xc9\x5f\x71\x59\x3a\x4a\x77\xcb\xa0\x3f\xe4\x0d\xec\x7f\xbb\x5f\x0a\x2b\x8d\xd4\x2e\x90\xe4\x2f\xf7\xf7\xd1\xf6\xa5\xae\x7b\x09\xfe\xc6\xfc\x8e\xa5\x7a\xfe\xb6\xb7\x1a\xc7\x2d\x7b\xd7\xbd\xcf\xb5\x5f\xcb\x97\xbb\xf6\x07\x51\xe6\x93\x48\xb2\x76\xaf\x47\x8c\x9a\xea\x9b\x98\x62\x52\xf5\xc7\x04\xee\x37\x04\xb8\x74\xf6\xd7\x36\xff\x0b\x00\x00\xff\xff\x24\xc6\x1d\x45\xd0\x25\x00\x00")

func pagesAssetsHtmlContainersHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pagesAssetsHtmlContainersHtml,
		"pages/assets/html/containers.html",
	)
}

func pagesAssetsHtmlContainersHtml() (*asset, error) {
	bytes, err := pagesAssetsHtmlContainersHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pages/assets/html/containers.html", size: 9680, mode: os.FileMode(420), modTime: time.Unix(1469609854, 0)}
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
	"pages/assets/html/containers.html": pagesAssetsHtmlContainersHtml,
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
	"pages": {nil, map[string]*bintree{
		"assets": {nil, map[string]*bintree{
			"html": {nil, map[string]*bintree{
				"containers.html": {pagesAssetsHtmlContainersHtml, map[string]*bintree{}},
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
