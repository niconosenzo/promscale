// Code generated by vfsgen; DO NOT EDIT.

package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// SqlFiles statically implements the virtual filesystem provided to vfsgen.
var SqlFiles = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 3, 5, 15, 23, 34, 201376728, time.UTC),
		},
		"/1_base_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "1_base_schema.down.sql",
			modTime:          time.Date(2020, 3, 5, 15, 23, 34, 201239935, time.UTC),
			uncompressedSize: 126,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x76\xf6\x70\xf5\x75\x54\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x88\x2f\x28\xca\xcf\x8d\x4f\x4e\x2c\x49\xcc\xc9\x4f\x57\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\xc2\xa7\x38\x33\xaf\x24\xb5\x28\x2f\x31\x87\x90\x6a\x90\x62\xb8\x1a\x40\x00\x00\x00\xff\xff\xdd\x04\x86\x65\x7e\x00\x00\x00"),
		},
		"/1_base_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "1_base_schema.up.sql",
			modTime:          time.Date(2020, 3, 9, 12, 42, 50, 565055167, time.UTC),
			uncompressedSize: 11332,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x5b\x6f\xe3\xc6\x92\x7e\xe7\xaf\xa8\x87\x2c\x2c\x6e\x44\x61\x8c\x00\x79\xb0\x77\x02\x28\x32\x6d\x73\x23\x53\x8e\x44\x65\x26\x1b\x04\x44\x8b\x2c\x49\x3d\xa6\x9a\x1c\x76\xcb\xb2\x16\xf9\xf1\x07\x7d\xe1\x9d\x92\xed\xe4\x1c\x24\x7a\x98\x31\xc9\xbe\x54\x57\x7d\x55\xf5\x55\x91\x8e\xe3\xcf\x02\x77\x61\x39\x4e\xb0\xa5\x1c\xa2\x34\x46\x20\x9c\xef\x77\xc8\x41\x6c\x89\x00\x41\x56\x09\x02\x23\xf2\x46\x44\x18\xa4\x2c\x39\xc2\x0a\xe1\xfb\xef\x20\xda\x92\x9c\x43\x92\xb2\x8d\x65\x4d\xe6\xee\x38\x70\x61\x31\xb9\x77\x1f\xc6\x10\x66\x79\xba\x0b\x23\x22\x48\x92\x6e\xae\x7b\x1f\x52\x26\x30\x67\x24\x69\x3f\x95\x0f\xaf\xc1\x71\x62\x22\x88\xde\x9c\xc3\x26\x85\x2d\xe6\x68\x59\x4e\xff\xcf\x72\x1c\x08\x94\x9c\x31\xae\x29\xa3\x82\xa6\x8c\x83\xba\xdf\x3f\xbe\xd8\x33\x18\xff\x38\x75\x9b\xd2\x8e\x38\xe6\x14\x39\x0c\x2c\x00\x00\x1a\x83\xbc\x26\xc9\x50\x5d\xee\x50\xe4\x34\x0a\x69\x0c\x94\x09\x7d\x2b\x21\x2b\x4c\xb8\xbc\xfe\xed\x77\x7d\x67\xe9\x7b\x3f\x2f\xdd\x81\x7e\x60\x83\xe7\x4f\xa6\xcb\x1b\x17\x06\x34\xb6\x2d\xbb\x3c\xae\xe7\xdf\xb8\x9f\x41\x6f\x16\xea\xb1\x72\xdd\x99\xdf\x2f\xce\x72\xe1\xf9\x77\x70\xe7\xf9\x50\x2c\x7c\x7d\xf6\x14\x6a\x50\xff\x21\x9e\xf0\x08\x81\xfb\x39\xd0\x57\xcf\x24\xd9\x23\x08\x7c\x31\xe7\x79\x9c\x7b\x0f\xe3\xf9\xaf\xf0\x93\xfb\xab\x12\xb9\x92\xff\x09\x8f\x43\x3d\xdc\xae\x1f\xb4\xf1\xa0\x73\xda\x57\x65\x0c\x9f\xf0\x18\x66\x29\x57\x46\x33\x02\x6b\x35\xd7\x84\x6a\x8a\x9c\xa5\xbc\xd2\x7f\x21\x84\x9e\x33\x94\x43\x6b\x52\x64\x29\xef\x2a\xbd\xa3\xe3\xae\x20\x67\x96\x3b\x7f\x26\x23\x7a\xa9\xf8\x85\x3b\xf7\xc6\xd3\xba\x56\x1b\x48\x92\x7e\x55\x3b\xa7\xc2\xbb\xbe\x29\xff\xe9\x3b\xa1\x7a\x5a\x13\xa9\x9a\xd2\xb0\x4a\xfd\x7e\xdd\x0e\xb3\x39\xcc\xdd\xc7\xe9\x78\xe2\xc2\xed\xd2\x9f\x04\x5e\xa9\x8c\xc2\x21\x47\x3b\xf2\x84\xa1\xd9\x4c\xad\x32\xb0\xd5\xc2\x73\x37\x58\xce\xfd\x05\x88\x9c\x6e\x36\x98\xab\x7b\xe3\x05\x7c\xb3\xde\xb3\xe8\x1b\xeb\xc6\x9d\x4c\xc7\x73\xd7\xfa\xd1\xbd\xf3\x7c\xf9\xcc\xfd\xec\x4e\x96\x81\x0b\xeb\x34\xdf\x11\x31\xb8\x68\x28\x4d\xee\x38\xfa\x2f\x6f\x20\xe8\x0e\x21\xf0\x1e\xdc\x45\x30\x7e\x78\x0c\xfe\xcf\xe0\x08\x6e\x66\x4b\x39\xec\x71\xee\x4e\xbc\x85\x37\xf3\x87\x85\xa7\xd0\x18\x3c\x3f\xb0\x2f\xf4\x59\xdb\x3f\xdf\xfd\x34\xaa\x1d\xfc\xfa\x8c\x1c\x25\x14\x8c\x28\x30\x28\x77\x18\x82\x14\xab\xa6\x63\x8d\xed\x77\xec\xa9\x55\x25\x1f\x5d\x5b\xae\x7f\x63\x19\x1d\x4d\xc7\xfe\xdd\x72\x7c\xe7\x42\x96\x64\x1b\xfe\x35\xa9\x41\x69\xee\xdd\xdd\xb9\x73\xe8\xe8\x3e\x34\xda\x06\xad\xee\xdb\xc0\x9d\x83\xe7\x2f\xdc\x79\xd0\x85\xb1\x9e\xa7\x06\xde\xce\xe6\xe0\x8e\x27\xf7\x30\x9f\x7d\x52\x37\x0a\x2d\x3c\xce\x67\x13\xf7\x66\x39\x77\xdf\x60\xf5\x6b\x4b\xfe\x4e\xc4\x50\x15\x74\x3d\x33\x1b\xe4\xf9\x5e\x8b\xba\x8e\x5c\x0b\xe6\x28\xf6\x39\x03\x52\x4b\x2c\xb0\xda\xd3\x44\xc0\x3a\x4f\x77\x40\x1a\x9e\x41\x58\x0c\x04\xf8\x7e\xbd\xa6\x2f\x23\x50\x61\x7e\x8b\x45\x78\x50\x23\x28\x07\x91\xef\x59\x44\x04\xc6\xc0\x53\x93\xb3\xb6\x68\x26\x41\x94\xee\x93\x18\xd6\x54\x00\x65\xb0\xde\x27\xc9\xe8\x3d\x9e\x50\x37\x84\xdc\x2e\x3c\x50\xb1\x0d\xf5\xd2\x83\x12\x0d\x1d\x5f\x2e\x36\x97\x17\x4d\xdf\x91\x63\xac\xca\x69\xe4\xa3\x85\x3b\x75\x27\x01\x0c\xf8\x7e\xc5\x45\x4e\xd9\xa6\xee\xe8\x12\xb5\xf0\xfd\x77\xce\x40\xe6\xda\x30\x41\xb6\x11\xdb\x81\x5e\xdd\xfe\xf6\xd2\xb6\xe1\x8f\x3f\xe0\x22\xbc\x90\xff\x99\xbb\x57\x57\x6a\x8f\x36\xe2\xf8\xd7\x04\xbc\x87\x87\xa5\x76\xbf\xc7\xf1\x7c\x3c\x9d\xba\x53\x58\x8c\x6f\xdd\xeb\xa6\x5d\x18\x1e\xa0\xdc\x9a\x94\xb1\x58\xea\x60\xa4\x2d\xa0\x74\x2e\x93\x92\x48\x61\xcf\x51\xa9\xbb\x3e\xac\x54\x35\xfc\xb8\x17\x40\xd7\x6a\x80\x9c\x5a\x37\x5c\x9c\x22\x67\x17\x42\x9a\x66\x08\x1b\x64\x98\x13\x81\xdc\xec\xbf\x67\xf4\xeb\x5e\x83\x43\xed\xe9\xa7\x02\x0b\xd3\x52\xcd\x43\xe4\xc6\xfb\x4c\x2d\xcd\xf0\x45\xc8\xb8\x01\xe9\xba\xd7\x23\xb4\xe9\xa4\x7e\x5e\xe4\x62\x3c\x05\x2a\x80\x6f\x15\x32\x56\x08\x11\x49\x12\x8c\x35\xb1\xa1\xeb\x12\x99\x52\x40\x60\xa9\x80\x23\x0a\xc0\x17\xca\xc5\x08\xde\x81\x1d\x86\x87\xb0\x83\x9f\x5e\xcc\x84\x24\xdf\x18\xdc\x34\xf8\xc5\x5b\x91\x33\x19\x2f\x5c\x28\x17\xfe\x74\xef\xfa\x50\x47\x4b\x6b\x27\x1b\xfe\x47\x72\xb7\xe0\xde\xf5\x1b\xe1\xac\x35\x4c\xc3\xa8\x5a\xd6\x9d\xd6\x37\x91\xbf\x3f\xe7\x2b\x27\xb6\xab\x1d\xfd\xea\x4a\xea\xa2\x31\xc1\xae\xc4\xe8\x09\xa6\x12\xda\xbf\xcc\xa6\xe3\xc0\xeb\x45\xf6\x24\xc7\x1a\xb2\xb4\x6d\x35\xb4\x37\xf4\x19\x59\x1d\x94\xa3\x82\x06\xef\x39\x72\x89\x2e\x9e\xee\x10\x38\x7e\xdd\x23\x8b\x90\x4b\xe4\x18\xd8\x14\x2c\xd8\x60\xc7\x72\x1c\x4f\xc3\xfc\x24\x76\xde\x01\x9d\x48\x49\xdc\x0c\xc6\x67\x81\xd3\x44\x8a\x72\xf1\x81\xe1\xa8\x6d\x4a\x61\xb7\x30\xf4\xc9\x0b\xee\xa5\x62\x24\xe6\xc6\x0b\xa8\xf6\x31\xe0\x92\xbe\xf5\x4c\x92\x41\xb6\x09\x37\x28\x42\xcd\x23\xc3\x42\x25\x83\x8b\x3e\x7f\xbb\x18\x5e\xd0\xf8\xc2\xb6\xaf\xae\x28\x13\x40\x38\xd0\x58\x2d\x6b\x08\x0a\x65\xbc\xb9\x95\xc9\x67\x9e\x1f\xcc\x4e\xb0\x29\x99\x93\x6b\x07\xaf\x9f\xca\xee\x4a\x2c\x0f\x33\xa2\xf1\xf0\x1c\xbe\x87\x67\xc1\xdc\xef\xbc\x9d\x25\xca\xad\x2a\x21\x6e\xe7\xb3\x07\x73\xbf\xbc\xa7\x0d\x23\xf9\xbb\xe2\x16\xe5\x82\x56\x01\x6d\x23\x79\xf3\xa9\x5e\x8a\x32\x6e\x69\x52\x37\xf3\x61\x3c\x9d\xca\x0b\xc7\x91\x48\x5b\xa7\x49\x92\x1e\x28\xdb\x00\xc7\x04\x23\x21\x53\x21\xc3\x08\x39\x27\xf9\x51\x86\xe6\xcd\x9e\xe4\x84\x09\x44\x20\x90\xab\xe8\xae\x27\x1f\xb6\xc8\x74\x1c\xa5\x1c\xf0\x05\xa3\xbd\xcc\x9d\x51\xca\xa2\x7d\x9e\x23\x13\xc9\xf1\xb4\x4c\x56\x71\xc4\x53\xcc\xe3\xd3\xbd\x3b\x77\x1b\x19\xf1\x63\x07\xb2\x72\xdc\xd4\x7b\xf0\x02\xb8\x7c\xaf\x33\xc3\x1d\x0a\xe3\xc9\xba\xc2\x21\x79\x4e\x8e\x50\xd6\x0f\xda\xb1\xf5\xa3\x27\x3c\x8e\xe0\x56\xde\x60\x47\x23\xc3\x50\x2e\x71\x40\x38\x10\xa6\x59\x42\x31\x51\xe5\xb2\x95\x2c\x7d\x21\x4a\x77\x19\x89\x14\x6e\xb3\x94\x73\xda\xc8\x7b\x2a\x30\xa8\xa8\x90\xe5\x28\xc4\x11\xb6\x48\x9e\x8f\x90\xa4\xd1\x93\x0a\x0f\x32\x2b\xf1\x8c\xc8\x34\x9e\x1c\xdf\x95\x2f\xa4\x77\x49\xdc\x64\x29\x0f\xd7\x69\x2e\xab\x91\x73\x0c\x43\x16\x2b\xe5\x65\xd3\xff\x29\x13\x56\x97\x97\x83\x2e\x9c\xb4\x9a\x28\x13\x8a\xa7\x2a\xf7\x0e\xbb\xb7\xeb\xd8\x07\x7f\xfc\xe0\x5e\x57\xbc\x1e\x1c\x47\x9e\x32\x4e\xf7\xf2\x61\xb4\xc5\xe8\x49\x9d\x5f\x22\x51\x55\xe9\x7a\xcc\x9a\x72\x01\x69\x26\xe8\x8e\x72\x41\x23\x3d\xf0\xaa\x96\xb7\xca\xc3\x65\x29\xb7\x0a\xd7\xb1\x9a\x0e\x79\xba\x40\xb3\xa0\x00\x5b\x4b\x49\x4d\xbc\x95\x0f\xc7\xfe\x8d\x2a\x69\x3f\x56\x9a\xb3\x74\xec\x09\x66\xa5\x5e\xae\x2d\x7d\xef\x16\x6e\x67\x4b\xff\xa6\x99\x24\x0d\xa7\xaf\xc6\x9a\x7c\x04\xde\xad\xfe\xbb\x9e\x93\x5b\x2e\xa3\x7d\x46\x9a\x38\xcd\xc3\x9e\xf0\xde\x89\x2f\x76\x25\x5c\x7d\xdc\xb5\x51\xae\xd4\xb7\x04\xa8\x84\xfb\x78\x2a\x6b\x02\xcd\xeb\x56\x18\x11\x69\x9c\x03\x02\xc9\x15\x0f\xc3\xf5\x5a\xc6\x86\x68\x4b\xd8\x46\x5a\x48\x51\xe3\x68\x8b\x3b\x52\x37\x16\x49\x78\xaa\x6a\x0f\x0e\x7c\x9f\x63\xc9\xaa\x2a\x68\xac\x30\x49\x0f\xa0\x3a\x44\x79\xae\xa2\x0d\x03\x81\xf9\x8e\x4b\xd2\x55\xc6\x8e\xe8\xd8\x28\x36\x8a\x92\x6b\x3a\x9b\xfc\xd4\x2c\xfc\xc0\xf3\x61\x71\x3f\x9e\xbb\xb0\x7c\xbc\x91\x1e\xe2\x7e\x9e\x4c\x97\x0b\xef\x17\x17\x1e\x66\x37\xee\xc5\xb0\x71\x6c\xbb\x38\x37\xc7\x28\x65\xb1\x01\x1d\x59\x0b\xcc\x15\xf4\xfe\x49\xa8\x6a\x01\xaa\x04\xc6\x5f\x07\x55\x25\x12\x79\x51\x9d\x08\xf8\x16\x2e\xff\x63\x67\xd4\x07\x69\x80\xa0\x79\x9a\x26\x3e\xbc\x05\xf8\xcb\xe9\xb4\x79\xba\xe6\x90\xab\x8f\x70\x79\x0d\x8e\x03\x97\x0e\x65\x31\xbe\x60\xac\xa3\x37\xef\x3f\xf4\x69\x56\x70\xe2\x4c\xf2\xf7\xcb\x78\xba\x74\x17\x8d\x5e\x49\x15\x2b\x87\x4d\x81\xb4\x8f\xcd\x7c\x98\xcc\xfc\xdb\xa9\x57\x53\xf0\xcd\x0c\xfc\x59\x70\xef\xf9\x77\x56\x23\x7f\xd7\xc1\x75\xce\xd0\xfe\x2c\xe8\x35\xf6\xd8\x5b\xb8\x70\x31\x51\x0c\x52\xb2\xc3\x35\x55\xf5\xad\xcc\x66\xc5\x42\x17\x5d\x08\x74\x30\xf2\xa6\x8e\xc2\x3b\x52\x8e\xd6\x27\x8d\x07\x8d\xa4\x62\x5a\x31\xf5\x2c\x53\x64\x18\xcf\x0f\x54\xe7\xa7\x46\x21\xcf\x92\xba\x1e\xaa\xd5\xb0\x65\xbd\x83\xd8\x36\x65\x63\x6e\x29\xe1\xb0\x92\xad\x9a\x51\x33\x65\xdb\x84\xd0\xa2\x61\x9a\x8f\xd6\x23\x36\x8d\x4b\xc2\x05\x45\x1f\x4d\x53\x2e\x30\xc1\x47\x51\x26\xc3\xb5\x2a\xa2\x45\x19\x44\x84\x63\x23\x0c\x0a\x33\x83\x32\x8e\xb9\x80\x83\x8c\xb4\x2a\xa6\x9a\x1b\x65\x79\x90\x9b\xa2\x9b\x15\x22\x39\x0e\xa7\x2c\xc2\x13\x47\xe9\x4c\xa4\xeb\x42\x38\x2a\x70\x07\x24\xc9\x91\xc4\x47\x5d\x71\xf0\xfa\xe9\x4a\x25\x98\x7d\x7a\x38\x9c\xb2\x44\x15\x1c\xaa\x29\x4f\x78\xac\x07\x38\x19\xf3\xca\x67\xba\x5b\xf7\xb1\x06\x15\xf5\xa8\x64\x77\x35\x84\x2e\x7e\x9e\x2a\x0a\x77\xc8\x49\x96\x61\x0e\x24\x4f\xf7\x2c\x86\x2f\x3c\x65\xab\x10\x49\xb4\x0d\x25\xc8\x14\x71\xa5\xcf\x92\xb3\xae\x50\xc8\x00\x9f\xa7\x87\x10\xb9\xa0\x3b\x22\xd0\x72\x1c\x99\xf2\x4c\xcf\x7d\x70\xf9\x41\x29\xe3\xf2\xc3\x07\xfb\x1d\xb0\xd7\x70\x6f\xed\x3b\xf8\xc2\xb5\x28\x43\x80\xd9\x32\x50\x87\xd6\x5e\x20\xaf\xaa\x26\xb9\x6d\x95\x3c\x6b\xe1\x06\xb3\x5b\xc8\x31\x4a\xf3\xd8\x82\xf2\x9c\xc5\x3e\xd6\xa9\x7e\x0b\x2c\x82\xb9\xb4\xeb\x7c\xf6\x69\x01\x97\x1f\x4a\xae\x26\x9d\xfb\x9b\x96\x58\xd5\x83\x3e\xdd\xed\x19\x43\x5e\xa9\xac\x52\x18\x14\x0a\xfb\x6b\x3a\xd2\xeb\xeb\x17\x0e\xa1\x66\xda\x84\x1d\xd5\x1f\x1d\x3d\x10\x76\xc4\x04\x77\x0a\xfc\xff\x26\x5d\xa8\x8d\x8c\x10\x0d\x45\x9c\xec\x30\x9e\xf9\xf5\xcd\x81\xc7\xfd\x2a\xa1\x11\x8c\x1f\x3d\x0e\x6f\x9c\xf3\xea\x3e\x56\xb5\x6e\x21\xb3\xb2\x10\xea\x8a\x43\xf9\x50\x5a\x6f\x16\x74\x9b\x11\x65\xb9\x71\xa0\x49\x02\x9a\x2f\xf6\x74\xd9\xd6\x40\xc5\x7b\x5b\x0d\xaf\x71\xd0\x93\x25\x47\x4f\x8f\x01\xde\xd8\x64\xe8\x2d\x25\x6b\x0d\xa6\x93\x05\x25\xec\x6a\x25\xe5\xa8\x59\x54\xbe\x72\x90\x51\x8b\xae\x35\x0a\x68\x23\xd0\x7f\xb7\x77\x3f\xdb\x81\x79\xc7\x86\xf6\x5f\xa8\x70\xb3\xd3\xd0\x39\x5f\xeb\xbe\x66\xf5\x06\x6b\xfa\x93\xa5\xa5\x4c\xfc\x56\x95\xf9\x1d\x47\xf5\xc1\x4c\x0f\x81\x6b\x88\xa6\x39\x02\xbe\x64\xc8\xb8\x8c\x49\x8f\xd3\xc7\xbb\xc5\xcf\xd3\xea\x3c\xa6\x1d\xac\xc5\x2e\x9a\xc0\x0a\xb7\xbd\x5c\xf7\x6f\xa2\xf1\x7d\x0c\xa0\x9d\x4a\xdf\x50\xc0\xf7\xd2\x50\xfb\xd5\x14\x79\x67\x0c\x5e\x10\x33\x63\xea\x8a\x25\x41\x46\x68\xfe\x36\x83\xd3\x78\x50\xcf\xe4\x67\xf9\xdd\x59\x53\xaf\x69\xce\x45\xc1\x7f\x44\x0a\x59\x8e\xcf\xc8\x44\xd9\x1e\xd5\xef\x6f\x56\x28\xab\xcc\x3d\xc7\x18\xf6\x59\xd1\x2e\x60\x05\x5d\xa2\xba\xb7\xf4\x37\xb1\x92\x46\x0c\x78\xab\x39\x3b\xe4\xb8\xae\xb7\xd7\x4d\xa9\xc2\x78\xfd\xf5\x46\x55\x51\x13\x5e\x54\x40\x5a\x75\xd2\xe2\x2a\xfd\x5b\x8e\xf3\x81\x43\x8e\x59\x8e\x5c\x6a\xf8\x09\x8f\xe6\x63\x0c\x59\xd4\xcb\x58\xcf\x51\xc0\xe0\x80\x10\xa7\xd2\x7b\x64\xc5\xaf\x8a\x2f\xab\xa0\x8f\x94\x09\xbd\x70\x99\x1d\xf8\x3e\xcb\xd2\x5c\x00\x15\x76\xd9\xe5\xa6\xe5\x23\xcc\x21\xc3\x5c\x95\xeb\x72\x7a\x94\x53\x41\x23\x92\x80\x5e\x4d\x12\x5f\xcb\x71\x28\xd7\x6d\x6f\x65\x5c\x19\x99\x2a\x42\x1b\x25\x54\x0a\x4a\x58\xac\x5a\x04\x24\xda\x62\x2c\x9f\xe7\x78\xbe\x07\xa6\xd9\x8e\x48\xc3\x1a\xc5\x28\x99\x58\xa3\xde\xf8\xed\xf7\x4e\xc5\x11\xbf\x84\xcf\x24\xe9\xed\x5a\x37\x8a\x06\xc7\xd1\x72\x47\x24\x49\x94\xd0\xd5\xcb\x4a\x91\x16\xd9\x55\xd6\x60\x12\x52\x55\x43\xb0\xbd\x86\x3c\x9c\x8e\x5a\x34\xe6\x26\x8c\x1d\x8d\x05\x74\xfc\x1a\x48\x37\xad\xeb\x31\x47\xc2\x53\xc6\xed\xc6\x52\x51\x4a\x12\xe4\x11\x0e\x92\xa7\x6c\x94\xa5\xbc\xfd\x56\xb9\x1b\xab\xbf\x70\xe7\x87\x1f\x2e\x42\xdd\x40\x0d\x2f\x86\x80\xa3\x27\x3c\xda\xb6\x54\x41\xab\x8f\x5d\x2d\x3e\x92\xe9\xf6\xe4\xd2\x34\x1e\xa8\x45\xe4\x5a\xba\xee\xb2\xd5\x1b\xb4\x72\x42\x5f\x56\x3c\x45\x9c\x6d\xc0\xe6\x46\x53\xf7\x36\x80\xff\x9d\x79\xbd\xdf\x59\x40\xd2\x92\x4a\x16\x3b\x83\x64\xa4\x1d\x5a\x49\xa5\xbc\x39\x19\x15\x6e\x5c\x88\x68\xbd\x79\x8f\xe6\x47\x25\xc9\x53\xd6\xb3\x65\xfb\x56\xf7\x1d\x15\xc8\x99\xa3\x32\x73\xb4\xcc\xd0\x0c\x39\xcd\x39\xb5\xb3\xb4\x47\xd8\xdd\xd2\x73\x3c\x9f\x8f\x7f\xed\x60\xb8\x34\xa5\x01\xba\x54\xc2\x10\x3e\xd8\x2d\x1b\x59\x4d\x03\xeb\x40\x13\xea\x6f\x19\xfa\x8e\x04\x70\xd9\xc1\x85\x51\x80\xd9\x78\x47\x5e\xe4\x9e\xb6\xa9\x87\xf5\xee\x4d\xe5\xdb\xb0\x81\x13\xd6\x28\xdc\x52\x1a\xb5\x90\x9c\xc6\x2f\x92\xb1\x15\x47\x6f\xc5\xc9\x73\x6c\x48\xbf\x95\xe6\x55\x6c\x34\xa9\xd0\x7c\x15\xa5\x3e\xf7\xc2\x18\x56\xc7\x82\x09\xe9\x18\x62\x39\x4e\xeb\x95\x71\xc9\x3f\xaa\xf7\xbe\xba\x39\xfd\xff\xa8\xd3\x6b\xdd\x6f\xcf\x85\xac\xda\x2e\x32\x70\x29\xa1\x06\xf5\x8f\xc0\xaa\xb0\xa5\x05\xae\xc2\x56\x8b\xda\x68\x47\x4a\x57\x5f\x30\x12\x03\xbd\x20\xd9\x6c\xb4\x27\xd8\x43\xa8\xdf\x31\x2e\xda\x66\x42\xaf\x16\x6d\xdc\x2e\x39\x84\x99\xe2\xf9\xbe\x3b\x3f\xe3\x9b\xc6\x19\x69\x0c\x1f\xcb\xa9\x36\x74\x4c\xb6\x38\xf5\x1d\xc1\x4d\xaa\x72\x8a\x8a\xb5\x2b\x5c\x4b\x32\xa8\xba\xb5\xba\xf7\x5c\x7c\x9b\x61\xac\xd7\x29\x5d\xde\xff\x92\xb4\x0d\xf4\x9e\x2f\xf4\xa0\xf8\x4a\xcf\x94\xb0\xc6\x4a\x70\xf2\x95\x09\xbc\xa9\xa9\xd5\xf7\x99\xde\xa0\xdc\x7e\x68\x9a\x14\x9d\xd7\x93\xed\x11\x06\xb0\xc5\xa8\x76\x93\xaa\xf9\x7e\xf0\x9f\xf2\x4e\xf0\x44\xd5\xa6\x95\x50\xd5\x6c\xc6\x2d\x3e\x36\xd4\xff\xa7\x6b\xa3\x2e\x36\x9a\x7c\xb7\xfc\x7e\x4b\xd1\x6f\xdd\x58\xd4\xa8\xee\xd2\x89\xca\x2b\x95\xa1\x27\x81\xdb\x30\xb4\x39\x68\x2f\x3d\x51\x7f\xdb\x5d\xdb\xbc\x57\x27\x83\xb2\x0a\x55\xb3\x22\xa1\x13\x5c\x5f\xa1\x7a\x16\xfa\xaf\x15\xa6\x5a\xe0\x46\xea\xb2\x6d\xc5\x0e\x7a\x24\xb0\x9b\xe6\x79\x5b\xa0\xfe\x57\x00\x00\x00\xff\xff\x42\x3f\x47\x09\x44\x2c\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/1_base_schema.down.sql"].(os.FileInfo),
		fs["/1_base_schema.up.sql"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
