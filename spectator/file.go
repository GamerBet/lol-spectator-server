package spectator

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func UnBufferedRead(filepath string) []byte {
	b, _ := ioutil.ReadFile(filepath)

	return b
}

func BufferedRead(filepath string) ([]byte, error) {
	fi, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(fi)
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	result := []byte{}
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}

		// send a chunk
		result = append(result, buf[:n]...)
	}

	return result, nil
}
