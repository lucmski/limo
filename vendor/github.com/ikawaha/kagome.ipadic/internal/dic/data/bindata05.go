package data

import(
	"os"
	"time"
)

var _dicIpaUnkDic = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xfd\xdf\xc9\xc4\xc8\xf4\xbf\x8b\x81\xf1\x7f\x07\x03\x83\xe9\xff\x76\x66\x46\x46\x56\xdf\xfc\xa2\x82\x0c\x90\x00\x23\x33\x23\x9b\x4f\x6a\x5a\x89\xa7\x0b\x23\x0b\x03\x23\x7b\x50\x66\x7a\x06\x94\xcd\x16\x9e\x0a\xe2\x00\x99\x0c\x0c\xff\x18\x7f\x00\x0d\xd0\x60\xfc\xc7\xc5\x05\xc1\x19\x6b\x18\x80\x0c\x29\x08\x9e\xa4\x02\x54\x0d\x84\xff\xca\xd9\x40\xa2\x02\x10\x5c\x55\x04\xe2\x08\x41\x70\x9f\x04\x88\x23\x01\xc1\x39\x62\x48\x32\x49\x7d\x48\x32\x4e\x11\x48\xc6\x26\x15\x30\x20\x2c\xb4\xb3\x41\x32\xda\x57\x84\x81\x91\x0b\x08\xff\xa9\x3a\x21\x29\xb1\x9d\x81\x6c\xd2\x2f\x24\x3b\x92\x8d\x90\x34\xfb\x45\x21\xd9\x91\x2c\x84\xc4\xe9\xde\x82\xa4\xac\xb8\x0d\xc9\x80\x5e\x09\x24\x7b\xd2\xa4\x40\x1c\x36\x08\x9e\xfb\x0f\xea\xf7\x96\x5d\x48\xb6\x17\x07\x21\x29\x69\x17\x41\x92\x49\xba\x84\x64\x52\xa4\x0b\x92\x85\xe5\x75\x48\x16\xb6\x57\x20\xb9\xab\x73\x1a\x88\x23\x07\xc1\xd7\x16\x21\x29\xcb\x92\x42\x32\xcd\xf3\x15\x92\x9e\x2c\x64\xe7\x84\x86\x20\xd9\x13\x64\x04\x75\x71\x5e\x09\x92\xb1\x57\xbf\x30\x30\x0a\x01\xe1\x3f\xd7\x3e\x24\xa7\x77\x26\x31\xf0\xfd\x6f\x64\x01\x26\x9f\x26\x06\x46\x16\x50\x52\x90\x02\xb2\xb8\x45\x78\x84\x24\x58\x0c\x04\x6c\xf8\xec\x38\xbc\x98\x7c\xd8\xfc\xb8\x18\x18\xc4\x78\x94\x04\x80\x52\x6c\x5c\x5c\x22\x1c\x42\x1c\x3c\x3c\x2c\x5c\x7c\x5c\xbc\xff\xbb\x41\x29\xaf\x07\x98\xf2\x5a\x18\x18\x78\xfe\x37\x83\x78\x2d\x0c\x8c\x3c\xc0\x04\xc5\x6a\x06\x14\xd7\x60\x67\x7b\x3a\xa1\xf7\xc5\xca\x79\x6c\x4f\x76\x34\xbc\xe8\x58\xc3\xa8\x05\x83\x30\x09\x9e\xa7\xb3\x77\x3d\x9b\xd3\x09\x55\xf5\x74\xce\x86\xa7\xf3\xe7\xa3\x28\x66\xe7\x7c\xd6\x32\xff\x69\xf7\x54\xa0\x34\x42\x37\x2e\xfd\xa8\xb6\xe0\x50\xb3\x6b\x17\x90\x85\x6a\x07\x56\x85\xcf\xb7\xb6\x3c\xdf\x39\x85\x8e\x86\x61\xf3\x3d\x69\xe1\x87\xee\xff\x17\x2b\x66\x3c\xed\xdf\x8e\x53\x33\x91\xa6\x52\x35\x20\x88\x89\x22\x3c\x01\x41\xa6\x42\x0a\x13\x06\xee\x80\x7a\xdc\xb4\xf5\xe9\x92\xce\x67\x7d\x4b\x9f\x6f\x9b\x85\x24\x4d\x4a\xa2\xc5\x11\xbc\x38\x4c\x26\xc1\x08\x32\x13\x0d\x45\xd1\x8b\x27\x46\x98\x9f\x4d\xdd\x40\xd0\x2d\xe4\xc4\x01\x79\x49\x82\xa8\x54\x8d\x1a\x34\xf8\x63\x15\xdd\x7f\xd0\xcc\xf7\x7c\xe5\xae\xe7\x33\xf7\x12\x9d\x74\x00\x01\x00\x00\xff\xff\x98\x66\x95\xae\xce\x07\x00\x00"

func dicIpaUnkDicBytes() ([]byte, error) {
	return bindataRead(
		_dicIpaUnkDic,
		"dic/ipa/unk.dic",
	)
}

func dicIpaUnkDic() (*asset, error) {
	bytes, err := dicIpaUnkDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/ipa/unk.dic", size: 1998, mode: os.FileMode(420), modTime: time.Unix(1452316192, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
