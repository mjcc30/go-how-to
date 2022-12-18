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
	return &HashReaderImpl{ // create a new hash reader
		Reader: bytes.NewReader(b), // create a new reader
		buf:    bytes.NewBuffer(b), // create a new buffer
	}
}

// create hash function wich return an encoded string of the hash
func (h *HashReaderImpl) Hash() string {
	return hex.EncodeToString(h.buf.Bytes()) // return the hash of the reader
}

// create a broadcast function wich take  and returns nil
func Broadcast(r io.Reader) (string,error) {
	b, err := ioutil.ReadAll(r) // read all from the reader
	if err != nil {
		return "", err
	}
	fmt.Println("string of the bytes: ",string(b))
	return string(b), nil // return the string of the bytes
}

// hash and broadcast function
func HashAndBroadcast(r HashReader) (string,error) {
	hash := r.Hash() // get the hash of the reader
	fmt.Println("hash:", hash)

	return Broadcast(r) // return the broadcasted bytes
}

// the main function
func main() {
	var test string = "hello high value software enginer" // the string to hash
	payload := []byte(test) // the bytes to hash
	h := NewHashReader(payload) // the hash reader
	HashAndBroadcast(h) // hash and broadcast
}