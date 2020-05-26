package res

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _res_doc_sample_txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x90\xcd\x6a\xf2\x40\x18\x85\xf7\xdf\x55\xbc\x8b\x6f\x9b\x94\x2c\xdc\xf6\x5e\x8a\x19\xa1\x60\x9a\x82\x23\xc4\xae\x42\x30\x6a\x21\xa2\xa5\x56\x8a\x85\x6a\x21\x56\x37\x99\x5a\x0a\xf5\x67\x12\x7b\x31\x99\x3f\x57\xb9\x85\x92\xa6\x08\xda\xae\x3b\x8b\x59\xcd\x39\xe7\x79\x06\x00\xfe\x9f\x5e\x99\x3a\x72\x10\x68\x35\x10\x93\xb5\xe8\x12\xf9\x46\x25\x1d\x8b\x5e\x5f\x86\x9b\x2c\x0e\x14\xd9\xca\x84\xf0\xe5\x94\xfb\x4b\xf5\x12\xa9\xb9\xcb\x5f\x1f\x85\x3b\x4b\x5d\xef\xdf\x41\xbe\x01\x26\xb2\xec\x13\x8c\x6a\x58\x6f\x9c\x59\x55\xd0\xca\x60\x94\x40\xab\x9c\xa3\xaa\x09\x5f\xb7\x01\x9a\x5d\xbc\xb2\xeb\xf8\xb2\x8e\x75\xec\x60\xd0\x2a\x80\x91\x83\xf3\xb2\xfd\x11\xd7\x33\xf5\x14\x88\xbb\x45\xce\x33\x18\x8b\x4e\x9f\xdf\x24\x8c\x86\xa9\xeb\xb1\xe4\x43\x0e\xe6\x3b\xbf\x2b\x13\x22\x86\x6d\x46\xdf\x0f\x67\xb3\x38\x28\x12\x46\x69\x5f\x91\xc5\x01\xa3\xd3\x7c\x45\x4c\x62\x1e\xf7\xe4\xa8\xa9\xb6\xb7\xbc\xbd\xe1\x9d\xc5\x11\x4e\x51\xc9\x56\xd1\x0f\x3f\xb3\x20\x2f\xdb\x96\x65\x5f\x7c\x1b\xfe\xa9\xf3\x2f\xb6\xea\xb9\x25\x1f\x86\x3b\x7a\xaf\x48\x58\x7c\xc9\x31\x23\x5b\x45\x72\xd4\xe4\x2d\x9f\x93\x75\xea\x7a\x9f\x01\x00\x00\xff\xff\xe0\x8d\x55\x32\xf2\x01\x00\x00")

func res_doc_sample_txt() ([]byte, error) {
	return bindata_read(
		_res_doc_sample_txt,
		"res/doc/sample.txt",
	)
}

var _res_doc_usage_txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x4d\x53\xda\x40\x18\xc7\xef\xfd\x14\x7b\xe9\x71\xe5\xd2\xe9\xa1\xdf\x06\x25\x14\xdb\x40\xac\x59\xeb\xcb\x29\x20\x2f\xa2\xc1\xd8\x0a\xbe\x50\x54\x70\x40\x69\x67\x92\xd0\xe9\x54\x42\x02\xf8\x61\xd8\xe7\xd9\x70\xe2\x2b\x74\x96\x58\xa1\x4e\x73\x4a\x32\xbf\xe7\xf7\xfc\xff\x3b\x4b\x08\xa1\x29\x42\x69\x4a\x51\x37\x08\xc1\x9b\x8e\x68\x1c\x81\xe7\xc0\xe1\x77\xfe\xd8\xc2\xac\x3b\x31\x72\xaf\x24\xa3\x13\x4a\x75\x85\x11\x42\x42\x67\x2c\x46\x0e\xf4\x3b\x50\xe8\x87\xae\x1d\x76\x0d\xf8\x79\x8d\xc6\xfd\xc4\xc8\x89\x6a\x17\x0f\xfa\xe0\x0e\x60\x74\x0a\xe5\x8a\xf8\xe6\xc0\xa8\x36\x6d\x18\xe1\x5d\x16\x1b\x65\x28\x5e\xe2\xd5\xfe\xf4\xf2\x44\x3a\xa5\x74\x1a\x5c\x84\x4e\x1b\x9b\x03\xac\x38\xe0\xd4\xf9\xa0\x8c\x67\x25\x1e\x3c\x88\xea\x0d\x1e\x9c\x60\xad\x87\x15\xe7\x69\xff\xfc\xa1\x09\x42\x69\x42\x49\xc6\xb7\x54\xf6\x77\x7a\x79\x6e\x01\xee\x12\x4a\x77\xe3\x69\x75\xfe\x85\x66\x09\x9c\xfa\xff\xc1\x35\x42\xe9\x9a\xb6\x95\x61\x0b\x70\x79\xbb\xa8\xe7\xf1\xaa\x85\xb5\xde\xf3\x84\x1c\xa2\xc9\x75\x45\x4d\xc8\xd7\xa8\x5b\x38\x3e\x85\x92\x2f\xea\x79\xb0\xcf\xd1\xf9\x3d\x1b\x9a\xa2\xda\x9d\x1a\xe7\x60\xf5\xe1\xa0\x38\xad\x57\x27\x46\x0e\x1b\x3f\x22\x3d\x9e\x3f\xcc\x86\x66\x14\x3e\x1a\x9c\x27\x8d\x82\x71\xcf\x96\x2b\xcb\x86\x3c\xae\xb9\xec\x9f\xfe\x1a\xa1\x54\xdb\x62\x64\xa9\x56\xa4\x78\x59\x2b\x29\x43\x6a\x9b\xe9\xf8\x9c\x7d\x62\x9a\x43\x18\x5a\xb3\xa1\x89\x55\x17\xcd\x2c\x53\x76\xd8\xc4\xc8\x7e\xd0\xb5\xcc\xc4\xc8\xee\xa4\x55\xf8\x6a\xea\x9f\xd4\xd0\xb5\xc1\xea\x3c\x27\xe4\x9e\x1f\x81\x4b\x29\x18\xa1\x94\xc5\x57\x55\xe5\xa5\x5b\xc2\xf2\x7f\x54\x51\x54\xbb\xdc\x3f\x8e\x52\xae\x67\x74\x65\x93\x45\x6e\x51\xcf\x87\xad\x2e\x9c\x54\x16\x46\xfa\x59\xd9\x5c\xd5\xf4\xb9\x10\x6e\xaf\xe1\xb0\x89\x76\x1b\xac\x1c\xd6\x7a\x32\xef\xc5\x58\xb4\xfd\xd0\xbd\x13\x41\x31\xec\xed\x87\x6e\x6d\x71\x35\xa5\x83\x7b\x3e\xf7\x8b\x78\x36\xe0\xe3\x47\x68\x74\x79\x10\x60\xde\x8a\xee\x1e\xf7\x6c\x2c\xdf\x87\x2d\xf3\xb5\x0e\x5f\x46\x3c\x68\xcf\x86\x26\x04\x7e\xe8\x38\xd0\x2b\x42\xe1\x01\x0e\x9b\x50\xe8\x88\x63\x17\x6e\xf7\xc1\xba\x98\x96\x2c\xee\xd9\x92\x29\xf4\xf9\xe8\x14\xcf\x06\xf8\xab\x06\x56\x4e\x14\xee\x79\xd0\xe1\xde\x11\x34\x7a\x70\x65\x3c\x9d\x46\x8a\xb1\x0d\xfd\x5d\x2c\xb6\xbd\xbd\xbd\xb2\xc7\x14\x9d\xad\x67\xde\xaf\x64\x14\x16\x5b\xd5\xb4\x8f\xb1\x3d\x96\xa4\x09\x6d\x2d\x16\x4f\x24\x28\xd3\xe8\x46\x9c\xa5\xe8\x9b\xb7\x2b\x29\x96\x56\xff\x04\x00\x00\xff\xff\x3a\x89\xa5\xcd\x74\x03\x00\x00")

func res_doc_usage_txt() ([]byte, error) {
	return bindata_read(
		_res_doc_usage_txt,
		"res/doc/usage.txt",
	)
}

var _res_messages_en_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x41\x6f\xdb\x3e\x0c\xc5\xef\xf9\x14\x84\x81\xa0\x97\xfc\x03\xfc\xb7\x9b\x8f\x2b\xba\x61\xc0\x2e\x6b\xb7\xd3\x3a\x08\xaa\x4c\x3b\x02\x1c\x2a\x11\xe9\x25\x41\xd1\xef\x3e\x50\xb6\xb7\x62\x88\x6c\x6f\x97\x1c\xac\x9f\xde\x23\x1f\xc5\x3c\xaf\x00\x8a\xd6\x52\xd3\xd9\x06\x8b\x12\x0a\xa4\xff\xbe\x3e\x14\x1b\xfd\xbc\x47\x66\xdb\x20\x17\x25\x7c\x5b\x01\x00\x3c\xa7\x5f\x80\xc2\x57\x8a\xd6\x21\x1a\x3c\xdb\xfd\xa1\xc5\x74\x21\x1d\x49\xb4\xc4\xad\x15\x1f\x28\xc9\x6d\x9b\x2d\xac\xb9\x48\xc7\x2f\x9b\x6b\x32\x12\xc4\xb6\x59\x81\x2f\xe9\x74\xbc\x7e\xed\xfe\x13\x36\x9e\x8c\x0b\x54\xfb\x26\x2b\xf3\x4e\x21\x90\x00\x3d\xb7\x01\xb9\x1c\x10\x6e\xf0\xec\xe5\x46\x3f\x1f\x3b\x2f\xdb\xa9\x2a\x91\x04\xa3\xf9\x95\x54\xce\xe7\x4e\x31\x90\x1d\xc2\x88\xc2\x25\x74\x70\xb2\x24\x6a\xd3\x31\x96\x8f\xf4\xff\x16\xee\xa8\x69\x3d\xef\x60\xcd\x8f\xf4\x66\x0b\xb7\x3b\x4f\xc8\x38\x93\x94\xeb\x62\x44\x92\xb9\x5e\x6f\x7b\x6c\x68\xb5\x9c\x0c\x8f\x3b\xe7\x90\xd9\x48\x30\x11\x6d\x95\xd5\x1c\x38\xed\x41\x39\xa8\x7d\x3b\x57\x6d\x6d\x7d\x3b\xea\x1a\xe5\xb3\xe2\xef\xad\x6f\xff\x54\x9e\x1c\xc6\x28\x7d\xb0\x91\x71\x99\x76\x42\xd3\x60\x5c\x20\xd1\x78\x42\xfd\x0f\x5e\xd1\x92\x6e\x44\xce\x6d\xcd\xe0\x19\x28\x08\xdc\x23\xdf\x27\x16\xea\x10\xf7\x56\x36\x20\xf1\xf2\xbb\x10\xcb\x4a\x7c\x24\x96\x57\xee\xd7\xec\x1b\x24\x8c\x56\xd0\x44\x74\x21\x56\x79\xeb\x0f\x03\x08\xeb\x0a\x06\x56\xed\x86\x16\x37\xc0\x07\xa4\x4a\x0f\x19\xdd\xb2\x8e\x5d\x20\x42\x27\x86\x8f\xad\x97\x7c\xc2\xf5\x90\x70\x38\x20\xc1\xc3\xe7\x4f\x5e\x7a\x43\x8c\x31\xc4\x72\x69\xbc\x78\x46\x67\x8e\x1d\xc6\xcb\xac\x91\xa2\x90\xd0\xbf\x9c\x5d\x38\xcd\x6a\xf7\xd3\xe9\x3b\x7e\x0b\x31\x9c\x96\x5a\xb0\xfd\x81\x06\xcf\x0e\xf3\x7f\x65\xa3\x87\xa2\x90\xd0\x12\x16\xcd\x21\xa2\x3e\x00\xb1\x4f\x13\xef\x7c\x14\xef\x61\x48\xf0\x32\x79\x4f\x8c\x51\x4c\x65\xc5\xce\xaa\xf7\x2c\x28\x3b\x23\xee\x5b\x34\x6e\xa7\x1b\x60\xc4\xef\x27\xea\xd6\xf7\xd9\x83\xa0\xa0\xee\xcf\xb2\x85\xb8\xd8\x7d\x3e\xeb\xd7\xdb\xa0\x60\xda\x03\x06\x4f\xf9\x5d\x58\x01\x7c\x5f\xbd\xfc\x0c\x00\x00\xff\xff\xe4\x6f\x10\x7d\x11\x07\x00\x00")

func res_messages_en_json() ([]byte, error) {
	return bindata_read(
		_res_messages_en_json,
		"res/messages_en.json",
	)
}

var _res_messages_zh_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xcf\x4e\xdb\x4c\x14\xc5\xf7\x79\x8a\x91\x25\x8b\x4d\x84\xf4\x7d\x4b\xf6\x5d\x54\xea\xa6\x74\x59\x2a\x6b\x70\x6e\x82\x25\x67\x0c\x1e\xa7\x0d\x45\x48\x16\x2a\x22\x94\xa4\xa0\x26\x21\x94\x3f\x0d\xb4\x4d\x41\x95\x12\x07\x09\xd1\x10\x1c\x78\x98\x7a\x26\x93\x55\x5e\xa1\xb2\x03\x08\xaa\xd8\x89\x60\x63\xc9\xf6\xf5\xf9\xcd\xb9\xf7\x5c\x2f\xc5\x10\x92\x74\x4c\x52\x19\x9c\x02\x69\x0a\x49\xef\xe7\xa4\xb8\xff\x2c\x0d\x94\xe2\x14\x50\x69\x0a\xbd\x8e\x21\x84\xd0\x52\x70\x45\x48\xd2\x12\x7e\x5d\xd2\x30\x15\xc8\xe2\xf4\xbc\x0e\xc1\x07\xc1\x2b\xcb\xc4\x84\xea\xd8\xd2\x0c\xe2\xd7\x78\x57\x1b\xec\xe7\x0a\x92\xa9\x14\x14\x2c\xc7\x87\x09\xcd\x42\x4a\x23\x8a\x6a\x90\xa4\x96\x0a\x55\x62\xae\xcd\x8e\x37\x7a\xab\x85\x6e\xa7\x11\x47\xe2\xaa\xc8\x56\x6b\x13\x90\xd5\xac\x89\x9e\x6d\xb3\xb5\xf6\x1f\x7b\x25\x8a\x01\xc4\x02\x53\xb9\xb3\x19\x46\x11\xce\xef\x81\xb4\xd7\x39\xe4\xfb\x55\xbe\xbf\xe7\x75\xae\xbb\xa5\x93\xee\xee\x07\xe1\xd4\xc5\x89\x3d\x35\x43\xfe\x9b\x44\x62\xe3\x94\x6f\xaf\x21\x99\xce\x90\xff\x27\x91\xd7\xaa\x0f\xee\xa2\x0e\xa0\x66\x4c\x13\x88\x35\xd2\x66\xa7\xc8\xd6\x0b\x03\x9b\x7d\x77\xf7\x4e\x71\x98\x24\xcd\xa8\x2a\x50\xaa\x58\x86\x62\x02\x4e\x44\x98\xba\x64\x9b\xdb\x7c\x7b\xcd\xbb\x3c\x97\x29\xcf\x6d\xb1\x8f\xd5\x11\xed\x4a\x62\x4d\xbf\xd5\x55\x92\x5a\xc4\x84\x1f\x8a\xb3\x1f\xa7\xe2\xac\x36\xa6\xf8\x3c\x36\x29\x8c\x50\x3f\xfe\xce\xbf\x6e\x3d\x41\xdd\xc4\xc4\x8f\x70\x98\xbe\x4c\xbd\x56\x81\xef\x38\xd3\x40\xa7\x83\x4a\x7e\xe8\x32\x77\xb3\xef\xe6\x59\xf3\x40\x38\xe5\x01\xdf\x6b\xb5\xa7\x81\x3e\x27\xd4\xa2\xcc\xb9\x10\xa7\x47\xf7\x8f\x30\xec\x0c\x29\x20\x60\x62\x0b\x14\x13\x54\xc3\x4c\x84\xf3\xbb\xa5\x2a\xcf\x6d\xc9\x09\xaf\xf5\x4b\x34\x9a\xac\x53\x66\xb9\xe6\xad\xdb\xbe\x9b\x17\x76\x85\x57\xce\xe5\x44\xf7\xf8\xf3\x98\xae\x55\x83\x10\x50\x2d\x85\x2e\xe8\x9a\x15\xde\x57\xbe\x5e\x64\xae\xfd\xea\xe5\x0b\xcd\x02\x5e\x6e\xf2\x42\x83\xb5\x8b\xb7\x0d\xee\xbb\xf9\x5e\xe9\x8b\x70\x9c\xbe\xbb\x2b\xd3\x31\xc1\x90\x05\x55\x59\xc8\x80\xb9\x18\x0a\xbd\xc1\x55\x6b\xc2\xf9\x76\x47\x12\x4e\x9d\x6d\xd6\x02\xd2\xa3\xb8\x37\x63\x36\xde\x8d\xc0\x8a\xa3\xfc\x60\x98\xff\x78\x1c\x1b\x44\xf1\x5b\x50\x20\xab\x82\x1e\xfe\xb7\xbb\x3e\x60\xf5\x9d\x67\x7e\xcd\x13\x3a\xa9\x9a\xe0\x47\xc7\xc2\xb3\x11\x8b\xc1\x72\x7b\xec\xb2\x2d\x8e\x4e\x1e\x37\x34\x4d\x07\x45\x9d\xf3\x13\xaf\x58\x5a\x3a\x22\x27\x41\x16\xbd\xeb\x06\x2f\x5d\xf0\xca\x79\xaf\x72\xe6\xb5\xda\x0f\xe5\x23\xf3\xbf\x88\xd3\xe1\xed\xba\x17\x7e\xbf\xee\x86\xd5\xfe\xd4\xdd\x6b\xb0\x4e\xd9\x5f\xcd\x7a\x1c\x0d\x5d\x81\x18\x42\x6f\x62\xcb\x7f\x03\x00\x00\xff\xff\xf5\xca\x37\x06\xba\x06\x00\x00")

func res_messages_zh_json() ([]byte, error) {
	return bindata_read(
		_res_messages_zh_json,
		"res/messages_zh.json",
	)
}

var _res_res_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x99\x4d\x6f\x1d\xc7\xd1\x85\xd7\xf7\xfe\x8a\xfb\x12\xf0\x0b\x12\x10\xe8\xf9\xee\x19\x03\xda\xd8\xc8\x36\x8b\x6c\xd3\x01\x31\x33\xdd\xa3\x30\x31\x49\xe5\x5e\xd2\x69\x2b\xf0\x7f\x0f\x9e\xaa\x33\x22\x63\x2b\x01\x22\x64\x41\x89\x9c\x8f\xee\xea\xaa\x53\xa7\x4e\xd5\x7c\x9c\xd7\xbf\xce\x1f\xf2\xe9\x9c\x2f\xc7\xe3\xfd\xc3\xc7\xa7\xf3\xf3\xe9\xfa\x78\xb8\x5a\x7e\x7e\xce\x97\xab\xe3\xe1\x6a\x7d\x7a\xf8\x78\xce\x97\xcb\xb7\x1f\x3e\xdd\x7f\xe4\xc2\xf6\xf0\xcc\x7f\xf7\x4f\xfc\x7b\x79\x3e\xdf\x3f\x7e\xb8\x5c\x1d\x6f\x8e\xc7\xed\xe5\x71\x3d\x2d\xf7\x8f\x69\x7e\x9e\xef\xce\x79\x4e\xd7\xfc\x76\xfa\xe3\x9f\x58\xeb\xdd\xe9\x71\x7e\xc8\x27\x7f\xfe\xe6\x74\xbd\x5f\xcd\xe7\xf3\xd3\xf9\xe6\xf4\x8f\xe3\xe1\xc3\x27\xfb\xeb\xf4\xdd\xfb\x13\x5b\xdd\xfe\x3e\xff\xfd\x0f\x79\x4e\xf9\x7c\x6d\xb6\xf0\xf7\xf7\x2f\xdb\x96\xcf\xb6\xec\xcd\xcd\xf1\x70\xbf\xd9\x0b\xff\xf7\xfe\xf4\x78\xff\x23\x4b\x1c\xce\xf9\xf9\xe5\xfc\xc8\x9f\xef\x4e\xdb\xc3\xf3\xed\xef\x58\x7d\xbb\xbe\x62\xa1\xd3\x37\x7f\xfb\xee\xf4\xcd\x4f\x57\x6e\x89\xed\x75\x73\x3c\xfc\x72\x3c\x1e\x7e\x9a\xcf\xa7\xe5\x65\x3b\xf9\x3e\xbe\xc9\xf1\x70\xe7\xe6\xbc\x3f\xdd\x3f\xdd\xfe\xf0\xf4\xf1\xe7\xeb\xff\x5f\x5e\xb6\x77\xa7\x0f\x9f\x6e\xb0\xf5\xf6\x87\x1f\x9f\x2e\xf9\xfa\xe6\xf8\xbf\x33\x43\x6f\x2d\x2f\xdb\xed\xf7\x58\x72\x7d\xf3\x8e\x25\x8e\xbf\x1c\x8f\x58\x78\x77\xce\x97\xbb\xf4\xb4\xde\x5d\xe6\x87\x8f\x3f\xe6\xbb\xe7\xf2\x7c\x7a\x2f\xef\x5e\x5f\xc5\x52\x6f\xb1\x8c\x4b\x2c\xd5\x18\x4b\x55\x7d\xf9\x67\xdb\x62\x59\xd6\x58\xa6\x2a\x96\x35\xc5\x32\xcc\xb1\x6c\x4d\x2c\x5d\x15\x4b\x3d\xc6\x32\xf6\xb1\x6c\x21\x96\xb4\xc5\xd2\xf7\xfe\x2c\x6b\x0e\x5b\x2c\xd3\x12\xcb\xd4\xc5\xd2\xac\xb1\xa4\x35\x96\x6d\x88\xa5\xcf\xb1\x8c\x73\x2c\xf5\x14\xcb\x5c\xc7\x32\x54\xb1\x4c\x73\x2c\x63\x13\x4b\xd3\xc6\xb2\x76\xb1\xcc\x39\x96\xae\x89\xa5\xad\x7c\xbf\xa6\x8e\x65\x6e\x62\x99\xfb\x58\xfa\xc1\xdf\x67\xdf\xfd\x1e\xd7\xda\x10\xcb\x34\xc5\xd2\xcf\xb1\x54\xd8\xc8\xfd\x10\x4b\xdd\xc4\x12\x96\x58\xda\xda\xef\xb7\xd8\x19\x62\x59\x26\x5f\x63\x62\xdd\xc1\x7d\x90\x66\xdf\xbb\x5d\xfc\x0c\xfd\xe4\x67\x6e\xa7\x58\x72\x88\x25\x4c\xb1\x54\x83\xfc\x92\x63\x99\x36\x3f\x8f\xad\x3b\xc7\x12\x9a\x58\x6a\x6c\x1e\x63\x69\x7b\xff\x7d\x6a\x63\x59\xfa\x58\xf2\xe8\xb6\x6c\x53\x2c\xdd\x10\x4b\xc3\xfd\xa4\x7d\x72\x2c\x3d\xb1\x18\xdc\x67\xf8\xab\xca\xb1\xd4\x5d\x2c\x09\x1b\x38\x6f\x17\xcb\x56\xc5\x92\x7b\xf7\xe9\xb6\xc4\xd2\x2d\x7e\x4e\xd6\x9d\x27\x3f\xd3\xba\xf8\x5a\x16\xdb\xde\xcf\xc2\x73\x7d\x8a\x25\x6f\x1e\xa7\xae\x8e\x65\xc9\xb1\x54\x75\x2c\xcd\x10\xcb\xd2\xc4\x92\xd7\x58\xea\x36\x96\x71\x75\xbf\xf6\xa3\x62\xb8\xba\x1f\x88\x6d\xaa\xdc\x16\x8b\x59\xe7\x18\x98\x17\x7f\x66\x6e\x63\x99\xf1\xfd\xe4\xfb\x2f\xc1\xd7\x27\xb6\xec\x0d\x2e\x6c\x1f\x6c\x1e\xfd\xf7\x31\xc7\x12\x7a\xdf\x9b\x35\x59\xbf\xe1\xac\xac\x5f\xc7\x32\xb6\xb1\x6c\xad\x3f\xbb\xa5\x58\xea\x3a\x96\x14\x62\x69\x5b\x9d\xbb\x8b\x65\x64\x2d\xce\xd8\xc7\xb2\x66\xbf\x67\xb8\xe4\x9a\x62\xc4\x99\x9b\xce\xcf\x86\x8f\xf1\x15\x76\x2c\x75\x2c\x99\x73\xb4\xee\xef\x3c\xb8\xbf\x16\x9e\x5f\xdc\x1f\x4d\xe3\xef\x0c\xc9\xe3\xc6\x5a\xd5\xe6\xd8\x5a\xda\x58\xda\x31\x96\x46\xb1\xe5\xfe\xa0\xf3\x63\xff\xda\xbb\x0f\xf0\x4d\x6a\x63\x09\xab\xdb\xd9\x71\x5e\xb0\x92\x3d\x77\xb0\x61\x1e\x15\xc3\xc1\x7d\x67\xfe\x22\x66\xf8\x30\xf9\x5a\x9c\x9f\x33\xf5\xac\x3d\x39\xee\x53\xed\xf6\x80\x6b\xec\x69\xf0\x21\x67\x26\x2e\xd8\xd3\xbb\x3d\xec\x5d\x2f\x8e\x5b\xf6\xc1\xaf\x3c\xc7\x7e\x79\x76\xec\x18\x26\x85\x43\xfc\x10\xc8\x07\xf0\x32\x3a\x1e\x58\x83\x7d\xc9\x25\x72\xb5\x97\xdf\xc1\x7d\xea\xfc\x1c\x4d\x72\x9f\x83\x7b\x8b\xed\xec\xeb\x70\x0d\x5f\xbc\xe5\x15\x7e\x72\x15\xcb\x98\x1c\x5f\x6d\xe3\xdc\xf2\xfa\xdc\xd5\x5e\x24\x7e\xcb\x64\xd7\x5f\x2c\x09\x3b\x23\xbe\xad\x29\xc7\xc3\xe1\x0b\x4c\xf8\xee\x78\x38\x5c\x9d\xf3\xe5\xdb\xf4\xb4\x7e\xeb\x97\x6f\x9f\xcb\xf3\xd5\xbb\xe3\xe1\xe6\x37\x04\xfa\x72\x99\x3f\x7c\x3d\x7f\x0e\xab\xf3\x4c\xc7\x39\x5b\xe7\x99\x9d\x3f\xd7\xe0\x79\x69\xf8\xee\x9c\xab\xf2\x14\x4b\xa8\x3d\xd7\x53\xe3\x7f\xc3\x95\x86\xbd\x9d\x3b\x3a\x8f\xb1\xe5\xa0\x72\x14\x4c\x93\xff\xcd\xf4\x8a\x03\x78\x73\xe5\xdd\xd1\x79\x91\xbc\xef\xab\x58\x7a\xf6\xaa\xfc\x7d\x30\x0b\x9e\xb1\x91\x1c\x64\x3f\xee\xc3\xc1\x55\xe3\xf9\x3a\x68\x0d\x78\x10\x4e\xe2\xdd\xdc\x78\x9e\x84\xce\xb1\x06\x46\x66\x61\x34\xb4\xb1\x74\xb3\xc7\x94\x7c\xe2\x3d\xb0\xbe\xad\xee\x93\x9d\x97\x2a\xbd\x83\xcd\xec\x87\x2d\xbc\x07\xc6\xe1\x74\xee\x63\x3f\xbc\x08\x5f\xc0\xad\x35\x79\x5c\xfb\x9a\x6b\xe5\x39\x12\x82\x63\x1b\x1b\xd7\xc6\x7d\xd2\x8a\xf7\xc1\xe7\x2c\x0e\x20\x17\x6b\xd9\x06\x3e\xc7\x3d\xaf\x1a\xed\x1f\x9c\x7f\xc9\x65\x38\x61\xeb\x7c\x1d\x7c\x06\x87\x8f\xc1\x6b\x04\xdc\x1d\x78\x26\xb9\x7f\xa8\x0d\xe0\x78\x1d\x3c\x96\xd4\xb1\x95\xe7\x27\xe7\xd3\x3a\x78\x9e\x6e\xb3\x63\x9e\xb5\xe1\x3c\x7c\x47\x5c\x88\x35\x75\x8d\xb5\xe0\x21\xce\x05\x27\xb4\xe2\x1e\xb0\xc2\x39\xe1\xd2\x7a\xf0\x5c\x26\xb7\xe1\x1e\xea\x07\xb1\x48\xbd\xfb\x00\x9b\xad\x46\x77\x5e\x97\xa8\x9b\xec\x53\xe3\xd7\xd5\x6b\x09\x31\x62\x0d\xfc\x5a\xb5\xf2\xd7\xe8\xb6\x91\xcb\xd4\xa9\xb9\x52\x3d\x08\xc2\x1c\xfc\xba\x3a\xe7\x92\xd7\x9c\x9d\x6b\x93\xf8\x7b\x4e\xee\x9f\xba\xf7\xd8\x10\x4b\xcb\xf3\xd5\x63\x5c\xbd\x8d\x31\xf6\x4d\xee\x2f\xf8\x0e\xbc\xa1\x0f\x88\x23\x7c\x81\x3f\x07\xd5\xa8\x9c\xc5\xaf\xb3\xdf\xcf\xad\xaf\x11\x74\x5e\xd3\x02\x83\xaf\x8f\x2d\xf8\x98\x7d\xc1\x0d\x75\xb8\x53\x0e\x51\x8b\xe0\xbb\xa9\x77\x3c\x57\x8b\xe3\x98\xbd\xa8\x21\xb3\xf0\xbd\xd5\x5e\xcb\x58\x97\xda\x9d\xb2\xf3\x25\x7e\x03\x7b\x96\x53\xda\xdf\xf8\x3c\x79\xac\xa9\x4d\xac\x01\x2f\xe2\xcb\x36\xbf\x6a\x0c\x72\x31\xeb\x5c\x0b\x75\x74\x73\xee\xc6\xcf\x6d\x52\x2c\x66\x5f\x1b\x7e\x80\xf3\x89\x97\xf9\xb1\xf2\x7a\x47\xfc\x89\x09\x6b\xe3\x6b\x74\x4c\x13\x5e\xb1\xca\x1a\xd4\x01\x6a\x16\x6b\x4e\xd9\xff\x5e\x55\xf7\xf0\x11\xbc\x81\x5d\xac\x8d\xaf\x78\x1f\xdc\xa2\xbf\xe0\x1c\xec\x22\x97\x1b\x71\x89\xd5\x82\xd5\x31\x4f\x8d\xa0\x56\xc3\xe5\x9b\xd6\x20\xae\xbd\xb8\x88\x9a\x36\x74\x5e\x5f\xa8\x4d\x33\x1c\x33\x3a\x37\x35\xe2\x25\x6a\x2d\xf9\x30\x89\xe7\xc0\x3b\x76\x84\xe4\xef\xe3\x83\x4e\x35\x1c\x2d\x47\x4d\x63\x8f\x51\x1a\x84\x3c\x42\x93\x91\xa3\x70\x29\x39\x68\x39\xaf\x7c\x0b\xd9\x39\x8c\xb8\x71\xee\xfd\x3a\x18\xc2\x1e\xd6\x30\x3e\x9b\x3d\x1e\xa6\x0f\xba\x37\xef\xd4\x7e\xbd\x95\x6f\xa8\x69\xf8\xd1\xb8\xbe\x76\xce\xe2\x1c\xb5\x78\x0b\x1d\x44\x6d\x26\x16\xac\x4d\x1e\x9b\x4e\x5d\x9c\x87\xc8\xc3\xb0\xb9\x8f\x7a\xf9\x0a\xac\xa1\x87\x89\x0d\x9c\xd6\x37\xee\x5f\xf2\x8c\x3c\x25\xaf\xa9\x9f\x60\x6b\xd8\xeb\xfd\x26\xee\xe1\xec\xca\x39\xd6\x82\x0b\x88\x3f\x98\xc7\xcf\x70\x33\x39\x6e\x5c\xd0\x79\xcc\x4c\x93\x66\xe7\x40\x30\x04\xfe\xf0\x67\x10\x8f\x90\xfb\xe0\x3a\x0d\x9e\x7b\x70\x36\x35\xc9\xce\x85\xbe\xea\x9c\xd3\xcc\x47\x49\xfc\x59\x3b\xb7\x82\x87\xac\xbc\x67\x0f\xb8\x2a\x28\xef\xf0\x03\x1a\xce\xf4\xb7\xb4\x0e\xf5\x32\x88\xd7\xe0\x1a\x72\xb9\x0b\xce\xeb\x9c\x3d\xec\xf6\x56\x9e\x77\x70\xa2\xe5\xbf\xde\xa7\x96\x5a\x3f\x21\x7f\xa0\x37\x36\x69\xec\x4a\x3a\xba\x13\x66\x59\x73\x95\x9e\x41\xe7\x56\x9d\xe3\x03\x1e\x27\x87\x4c\x6f\x2e\xe2\x87\x5a\xba\x23\x3b\x36\x3f\xf3\x3d\x58\x69\x1d\x77\x83\x6c\x00\x23\x0b\xd8\x19\x95\x2b\xab\xdb\x88\xef\x2d\xd7\x27\x71\xf0\x24\xdd\xdf\xb9\x1d\xd4\x6e\xb0\x87\x66\xe6\x77\x72\xa4\x93\xce\xe7\x7d\xce\x9e\x54\xcf\x38\x3b\x75\xa9\xed\xc4\x87\x95\xe3\x82\x73\x4d\xd2\xa2\xf8\x06\x0d\x6b\x35\x28\x79\x6d\xc0\x16\xd3\x90\xe9\xcd\x4f\xe3\x3a\x83\x7d\x78\x17\x0c\x81\xb5\x24\x4c\x90\xb3\xb5\xb0\x04\x6e\xc1\x0e\x31\x1f\x37\x5f\xdf\xb8\x44\xf9\x0c\xc7\x12\xe3\x46\x75\xa8\x53\x5d\xa5\x0e\xa1\x71\x39\x2f\xb6\x59\x4f\xa0\xbe\x07\x7f\xc2\xf1\x3b\x07\x4c\x3a\x3b\x1c\x4d\x4c\x7e\xad\x09\xa9\x59\xe0\x93\xf7\xc9\x75\xab\x91\xed\xbf\xd5\x84\x9f\xc5\xd9\xd7\x49\xc2\xcf\xaf\xff\x8b\x22\xb4\xab\x5f\x16\x84\x0f\xf9\xc2\xcd\xcb\x5d\x7e\xbc\xfb\xcb\xe5\xe9\xf1\xab\x34\x21\xfe\xb1\x1e\xaa\xf6\x3e\x2b\x29\x47\xaa\xd5\xf3\xce\x34\xe1\xa4\x1e\xa7\x73\xee\x81\xfb\x27\x69\x28\xfc\xc1\xff\xf8\xd5\xb4\x91\x7a\x16\x70\x49\x6d\x43\x3f\xc0\x13\x83\x7c\x4f\x6c\xf0\xab\xe9\xad\xd9\x7b\x10\x74\x0d\xfa\x8e\xdc\xa7\xef\x02\x6f\x59\x7d\x00\x76\x51\x9f\xb0\xa3\x95\x66\xdc\x7b\x13\x78\xda\xf2\x62\x55\x8f\x18\x54\x17\x55\x1b\xc0\x15\xfd\x41\x2d\x0e\xb1\xfa\xb1\xf9\xd9\xc1\x2c\x3c\x03\xe7\x61\x53\x52\x4f\x8d\xcf\xa8\xf1\x68\x21\x70\x66\x73\x87\xac\x5a\xd5\xa9\x77\x51\x1f\x69\x7c\x31\xb8\xcf\xd0\xb9\xb5\xf4\x09\xfd\x09\xf8\xdd\xfb\x07\xdb\x37\xa8\xb7\xae\x9c\x8f\xe1\xe9\xdd\x86\xa6\xf6\x1a\xc6\x73\x49\xeb\xcf\xea\xdb\xc0\x9c\xdd\x4f\xae\x53\xc8\x07\x9b\x4d\x24\xd7\x37\x9c\x0d\xcc\xaf\xca\x07\xeb\xdb\x93\xfb\x00\xce\xed\xa4\xed\x1b\xcd\x41\x88\x03\xbc\xca\x19\x59\x07\x5f\xc2\xd7\xc4\xd5\xb4\xc1\xe4\x7b\x72\x3e\xe3\xa8\xa4\x1a\xdb\x3a\x57\x52\x83\xd0\x62\x68\x96\x5e\xfd\x7f\xa3\x3a\x0d\xff\xc1\x23\xd8\x3b\x09\x63\xe4\xec\xde\x83\x5b\x0d\x58\x5f\xeb\x80\xcd\x11\x56\xaf\x55\xc6\x95\x9a\x17\x80\x87\xbd\x77\xa4\x86\x83\x09\xfa\xd8\x4a\xba\x19\x3c\xa2\x9d\xd8\x9f\x5e\x1b\x6d\xc2\xbe\xad\xf6\xc5\x57\xab\xb0\xc1\xff\x8d\xfc\x08\x77\xc1\x41\xf0\x04\x75\xd7\xb0\x58\x3b\x0f\x80\xdb\x4d\xf1\x86\x4f\xac\x4e\x8f\xae\xef\x5a\xcd\x2c\x06\xcd\x9e\xf6\x67\x4d\xeb\x09\xd7\xd6\x47\x6c\x5e\xe3\xe1\x74\xea\x26\x79\x95\xa5\x47\xd0\xc0\x70\x32\x3e\x40\x27\x83\x57\xea\xd1\xa0\xfa\x40\xdd\x20\x17\xc9\xbb\x51\x3e\xe0\x5c\x9c\x27\x29\x8e\x6d\xe5\xb9\x61\x71\xee\x9d\x1f\xeb\x5d\x2f\x27\xd5\xc2\xc9\xed\x46\xc7\x58\xbf\x10\xfc\x79\xab\x87\xba\x86\xae\xa8\xa5\xf5\xd1\x35\x59\x75\xd6\xf8\x58\xdc\xd7\xef\x7a\x63\x75\x6d\xdc\x68\xcd\x45\x33\x91\xa0\x39\x16\xb5\xbb\x53\x0e\xe1\x1f\x6a\x3e\xb9\x40\xce\x52\x5f\x3b\xcd\x4c\x2a\xcd\xa8\xb8\x8e\x1f\x88\x7f\xa7\xf8\x61\x1b\x7e\xa5\x0e\xd7\x93\xef\x65\x73\xb0\xd5\xf3\xd7\xe6\x26\x41\x71\x99\x1d\x3f\xd4\x3f\x9b\xb7\x0d\xbe\x1f\x5a\x19\x8d\x83\xa6\xe1\x3e\x9a\xd1\x74\xfa\xea\xf1\x36\xae\x0f\xaa\x8f\x41\x73\xa6\xe5\x75\x1e\x63\x33\x8c\xca\x6d\xda\x73\x8e\x6b\x41\xfd\x2c\xf6\x1b\xf7\x8f\xfa\x7b\x7e\xad\x9d\xa6\x3b\x06\xb7\xaf\x55\x2f\x58\x05\xe7\x0f\x6c\xe0\x5d\xce\x95\xd2\xeb\xbc\xc9\x6a\x64\xe5\xfe\x31\xed\x2b\x0c\x12\x03\x78\xcc\x66\x1c\x8d\xeb\x6d\x30\x10\xa4\xb7\x79\x87\xba\xbf\xb6\xc2\xbd\xfc\x1a\xa4\x57\x39\x0b\x98\x26\xb7\x83\x7a\x1c\xcb\x65\xd5\x43\xea\xed\xa8\xbe\x91\xb8\xae\xd2\x5d\x93\x66\x9b\x93\xe6\x5f\xcb\x3e\x73\x4b\xbe\xef\xe7\x99\xab\x38\x1c\x9e\xb0\xd9\x4c\xf6\x1c\xb0\x3c\x5a\xe5\xe7\xd9\x71\x42\xdc\xe1\x15\xfc\x64\xba\x43\xfd\x0b\x3a\x14\x0d\x41\x2e\x9b\x96\x7c\xb3\x2f\xbc\x84\x1f\x4d\x5b\xcb\x26\xab\x47\x95\x6b\x00\xce\x0b\x17\xe5\x7d\x96\xd4\x29\x37\x6b\xcf\x11\x38\x05\x7e\x23\x6e\x3c\x6b\xb1\xad\x35\xeb\xea\x9d\x53\xc1\xdc\xac\xb3\xa0\x53\xc0\x9b\x71\x4e\xeb\xf8\xb0\xbe\xa9\x93\x56\x0d\xea\x71\x2b\xf7\x69\x16\x16\xb3\xb8\x7b\xd3\x1c\x02\x7d\x49\x5d\xe4\x87\x35\x38\xaf\xc5\x9a\x1c\x1b\xbd\x37\xb0\xfc\xcb\x8e\x1f\x30\x68\xfc\xb6\x6b\xcc\x59\xf8\x92\x96\x81\xe3\xf0\x3d\xba\x17\xee\xb4\x9a\xb2\xb9\x8e\xda\xc4\x13\xbf\x99\x61\x75\x9e\xbf\xe0\x3f\x68\x4e\x09\x0e\xbf\xa0\x57\x7e\xad\x1d\xfe\x7b\xc9\xf2\xeb\x15\x3e\xab\x96\x37\x37\x6e\xb9\xf1\x1f\x84\xcb\xa7\x3f\x7f\xbd\x70\x21\xb9\x26\x09\x5a\xc8\xdf\x86\x50\x02\xc4\xaa\x8f\x01\x36\x0c\x51\xf1\xa0\x08\xd8\x30\x36\x69\x88\xdd\xb9\x93\x6c\x78\x3d\xc8\xd9\x9d\x07\x76\x1e\x9c\x04\x7a\x15\x9f\x41\x0d\x3a\x80\x18\x55\x50\x48\x24\x82\x50\x0b\x24\x88\x5f\x6b\xc2\x47\x07\x30\xef\xd1\x00\x63\x23\xb6\x12\xe8\x46\xc5\x9e\x6b\x14\x3b\xde\xa1\xa0\x77\x22\xef\x49\x43\x74\x0b\xda\xe4\x64\x49\x30\x6b\x25\x30\x42\x9f\x44\x67\x1d\x0a\x0c\x8d\x1e\xc9\x8f\x58\x00\x6c\x3b\x41\x01\xee\xb4\x0f\x7d\x2a\x6f\x1c\x10\x08\x9c\x13\x42\xd9\xf4\x41\x81\x62\x64\x44\xba\x7a\xc1\x5d\xf7\x8f\x05\x8d\x9f\x1f\x7f\x22\x90\x6a\x0d\x34\x2a\x11\x0a\xe4\x87\x18\xc9\x1a\xa8\x71\x3e\x40\x4e\x5c\xec\xa3\x41\x52\xf1\x6b\x3c\x1e\x7d\xf5\xfa\x1e\x00\x1e\x25\x12\x46\x35\xbc\xec\x61\x43\x40\xf9\x2f\x49\xa0\x04\x11\x65\xd2\x47\x18\x04\x8f\x35\xd1\x8d\x9a\x4a\x7d\xa0\x58\x05\x74\xfe\xdf\x07\x06\xab\x04\x6b\x96\x2f\x58\x83\x78\xd9\x80\x2a\x68\x20\xa0\xe4\x36\x11\x12\xdc\xbe\x49\x1f\x31\x6a\x35\x23\x9c\x07\x01\x35\x2a\x26\xf8\xad\xd3\x10\x08\xc2\x00\x5f\x36\x48\xa8\x84\xc9\xc1\xd7\x21\x36\x83\x1a\xea\x55\x67\x6e\x35\x54\xa2\xf1\x1a\xfb\x37\x04\x53\xbb\x98\xc4\x36\x62\xc7\x19\x29\x2c\x83\x86\xff\x46\x8e\xbd\x13\x35\x31\x9d\x54\x34\x17\x35\x9c\xc4\xb5\x96\xa0\x68\x34\xc0\xab\x54\x04\x88\x59\xdf\x7a\xac\x79\xbe\xd6\x87\x1f\xf6\x34\xf1\xa4\x81\x02\x38\xb4\x8f\x1a\x93\x13\x09\xc2\xce\x3e\xa4\xe5\xd7\xc1\x69\xe8\x35\x80\x59\xbc\x29\x9a\x44\x98\xfc\xe0\x7f\x6b\x18\x95\x7f\xa3\x84\x93\x0d\xbe\x5a\x7d\xa8\x5a\xbc\xc0\xe2\x97\x2c\x81\x04\x61\xd9\x50\x54\x82\x02\x9b\x26\xe5\x38\x44\x3b\xe9\x63\xcb\xac\x42\x6a\xc2\xac\xf1\xb8\x1b\x79\x0e\xca\xfb\xd6\x05\x5f\xab\x61\x35\xd7\xf1\xad\x0d\xc4\x24\xd4\xc8\x0f\x6c\x20\xef\x83\x06\xc1\x36\x24\xae\x3d\xb7\x10\x03\xc4\x0a\x01\x45\xfe\x82\xdb\x59\x1f\xa1\x6c\x70\xa5\x41\x74\x27\xf1\xc2\x3b\xd6\xe8\x74\xbe\x2e\x1c\x04\x51\x13\x07\xcb\xaf\x45\x8d\x72\xf0\x86\x13\xae\xb2\xb8\xd5\x5e\xec\x93\x86\xbe\x26\xde\x66\x09\xc5\xe4\xfe\xe6\x5e\xaf\xa1\x1d\x05\x75\xc7\xb3\xf1\xa3\x06\x52\xfc\x3e\xef\xb8\x59\x5d\x0c\xcd\x6a\xee\x5b\xfd\xcd\xbe\x56\x04\x2b\xe7\x1c\xe2\xb0\x2a\xa6\xe0\xaf\x15\xf7\xa4\x7d\x58\xac\x86\x9b\x7c\xb4\xf3\x65\xf7\x2f\x79\x6b\x1f\x3d\xf7\x66\x4c\x8d\xb6\x0d\xfe\x83\x17\x40\x13\x5c\xfb\x60\xa3\x71\x1f\x2f\x2a\xd6\x59\x8d\xc6\xa6\x77\x4c\xb4\xea\xc3\x14\x02\xcf\x1a\xc4\x5e\x6b\xd5\x5e\x28\x89\x3f\x78\xb2\x41\xe7\xfa\x3a\x44\xe2\x99\xee\x9f\x01\x00\x00\xff\xff\x6d\xc2\x26\x99\x00\x20\x00\x00")

func res_res_go() ([]byte, error) {
	return bindata_read(
		_res_res_go,
		"res/res.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"res/doc/sample.txt": res_doc_sample_txt,
	"res/doc/usage.txt": res_doc_usage_txt,
	"res/messages_en.json": res_messages_en_json,
	"res/messages_zh.json": res_messages_zh_json,
	"res/res.go": res_res_go,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"res": &_bintree_t{nil, map[string]*_bintree_t{
		"doc": &_bintree_t{nil, map[string]*_bintree_t{
			"sample.txt": &_bintree_t{res_doc_sample_txt, map[string]*_bintree_t{
			}},
			"usage.txt": &_bintree_t{res_doc_usage_txt, map[string]*_bintree_t{
			}},
		}},
		"messages_en.json": &_bintree_t{res_messages_en_json, map[string]*_bintree_t{
		}},
		"messages_zh.json": &_bintree_t{res_messages_zh_json, map[string]*_bintree_t{
		}},
		"res.go": &_bintree_t{res_res_go, map[string]*_bintree_t{
		}},
	}},
}}
