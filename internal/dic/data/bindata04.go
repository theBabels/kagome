package data

import(
	"os"
	"time"
)

var _dicIpaUnkDic = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xfd\xdf\xcd\xc4\xc8\xf4\xbf\x87\x81\xf1\x7f\x17\x03\x83\xe9\xff\x4e\x66\x46\x46\x56\xdf\xfc\xa2\x82\x0c\x90\x00\x23\x33\x23\x9b\x4f\x6a\x5a\x89\xa7\x0b\x23\x0b\x03\x23\x7b\x50\x66\x7a\x06\x94\xcd\x16\x9e\x0a\xe2\x00\x99\x0c\x0c\xff\x18\x7f\x00\x0d\xd0\x60\xfc\xc7\xc5\x05\xc1\x19\x6b\x18\x80\x0c\x29\x08\x9e\xa4\x02\x54\x0d\x84\xff\xca\xd9\x40\xa2\x02\x10\x5c\x55\x04\xe2\x08\x41\x70\x9f\x04\x88\x23\x01\xc1\x39\x62\x48\x32\x49\x7d\x48\x32\x4e\x11\x48\xc6\x26\x15\x30\x20\x2c\xb4\xb3\x41\x32\xda\x57\x84\x81\x91\x0b\x08\xff\xa9\x3a\x21\x29\xb1\x9d\x81\x6c\xd2\x2f\x24\x3b\x92\x8d\x90\x34\xfb\x45\x21\xd9\x91\x2c\x84\xc4\xe9\xde\x82\xa4\xac\xb8\x0d\xc9\x80\x5e\x09\x24\x7b\xd2\xa4\x40\x1c\x36\x08\x9e\xfb\x0f\xea\xf7\x96\x5d\x48\xb6\x17\x07\x21\x29\x69\x17\x41\x92\x49\xba\x84\x64\x52\xa4\x0b\x92\x85\xe5\x75\x48\x16\xb6\x57\x20\xb9\xab\x73\x1a\x88\x23\x07\xc1\xd7\x16\x21\x29\xcb\x92\x42\x32\xcd\xf3\x15\x92\x9e\x2c\x64\xe7\x84\x86\x20\xd9\x13\x64\x04\x75\x71\x5e\x09\x92\xb1\x57\xbf\x30\x30\x0a\x01\xe1\x3f\xd7\x3e\x24\xa7\x77\x26\x31\xf0\xfd\x6f\x65\x01\x26\x1f\x60\x68\xb0\x80\x92\x82\x14\x90\xc5\x2d\xc2\xc3\x20\x26\x24\xc1\xa3\xc4\x62\xc0\x67\xc7\xe1\xc5\xc5\xc0\xe6\xc7\xe4\x23\x60\x23\x00\x94\x62\x13\xe1\x10\xe2\xe0\xe1\x61\xe1\xe2\xe3\xe2\xe2\xe2\xfd\xdf\x0c\x4a\x79\x2d\xc0\x94\xd7\xc4\xc0\xc0\xf3\xbf\x11\xc4\x6b\x62\x60\xe4\x01\x26\x28\x56\x33\xa0\xb8\x06\x3b\xdb\xd3\x09\xbd\x2f\x56\xce\x63\x7b\xb2\xa3\xe1\x45\xc7\x1a\x46\x2d\x18\x84\x49\xf0\x3c\x9d\xbd\xeb\xd9\x9c\x4e\xa8\xaa\xa7\x73\x36\x3c\x9d\x3f\x1f\x45\x31\x3b\xe7\xb3\x96\xf9\x4f\xbb\xa7\x02\xa5\x11\xba\x71\xe9\x47\xb5\x05\x87\x9a\x5d\xbb\x80\x2c\x54\x3b\xb0\x2a\x7c\xbe\xb5\xe5\xf9\xce\x29\x74\x34\x0c\x9b\xef\x49\x0b\x3f\x74\xff\xbf\x58\x31\xe3\x69\xff\x76\x9c\x9a\x89\x34\x95\xaa\x01\x41\x4c\x14\xe1\x09\x08\x32\x15\x52\x98\x30\x70\x07\xd4\xe3\xa6\xad\x4f\x97\x74\x3e\xeb\x5b\xfa\x7c\xdb\x2c\x24\x69\x52\x12\x2d\x8e\xe0\xc5\x61\x32\x09\x46\x90\x99\x68\x28\x8a\x5e\x3c\x31\xc2\xfc\x6c\xea\x06\x82\x6e\x21\x27\x0e\xc8\x4b\x12\x44\xa5\x6a\xd4\xa0\xc1\x1f\xab\xe8\xfe\x83\x66\xbe\xe7\x2b\x77\x3d\x9f\xb9\x97\xe8\xa4\x03\x08\x00\x00\xff\xff\x08\x05\xd2\xa0\xce\x07\x00\x00"

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

	info := bindataFileInfo{name: "dic/ipa/unk.dic", size: 1998, mode: os.FileMode(420), modTime: time.Unix(1443284311, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

