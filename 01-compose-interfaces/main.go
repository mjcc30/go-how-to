package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

// create a Hash reader interface
type HashReader interface {
	io.Reader
	Hash() string
}

// create a hash reader struct
type HashReaderImpl struct {
	*bytes.Reader
	buf *bytes.Buffer
}

// create a new hash reader function wich take a byte and returns the hash of the reader
func NewHashReader(b []byte) *HashReaderImpl {
	return &HashReaderImpl{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

// create hash function wich return an encoded string of the hash
func (h *HashReaderImpl) Hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

// create a broadcast function wich take  and returns nil
func Broadcast(r io.Reader) (string,error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	fmt.Println("string of the bytes: ",string(b))
	return string(b), nil
}

// hash and broadcast function
func HashAndBroadcast(r HashReader) (string,error) {
	hash := r.Hash()
	fmt.Println("hash:", hash)

	return Broadcast(r)
}

// the main function
func main() {
	var test string = "hello high value software enginer"
	payload := []byte(test)
	h := NewHashReader(payload)
	HashAndBroadcast(h)
}