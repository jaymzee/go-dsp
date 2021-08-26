package plot

import (
	"encoding/base64"
	"fmt"
	"os"
)

func serializeGfxCmd(cmd, payload []byte) []byte {
	img := make([]byte, 0, len(cmd)+len(payload)+6)
	img = append(append(img, '\x1b', '_', 'G'), cmd...)
	if len(payload) > 0 {
		img = append(append(img, ';'), payload...)
	}
	return append(img, '\x1b', '\\')
}

func writeChunked(cmd string, data []byte) {
	var chunk []byte
	for len(data) > 0 {
		next := min(4096, len(data))
		chunk, data = data[:next], data[next:]
		if len(data) > 0 {
			cmd += ",m=1"
		} else {
			cmd += ",m=0"
		}
		os.Stdout.Write(serializeGfxCmd([]byte(cmd), chunk))
	}
}

func WriteKitty(head string, data []byte) {
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(encoded, data)
	writeChunked(head, encoded)
}

func WriteITerm(data []byte) {
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(encoded, data)
	fmt.Printf("\033]1337;File=size=%v;inline=1:", len(data))
	os.Stdout.Write(encoded)
	os.Stdout.Write([]byte{0x07})
}
