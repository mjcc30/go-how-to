package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// typed function
type TransformFunc func(string) string

// server structure with typed function
type Server struct {
	filenameTransformFunc TransformFunc // filename transform function
}

// handler function for server
func (s *Server) handleRequest(filename string) string {
	newFilename := s.filenameTransformFunc(filename) // call the transform function
	
	return string(newFilename) // return the new transformed filename
}

// hashFilename returns the SHA256 hash of the given filename.
func hashFilename(filename string) string {
	hash := newSHA256([]byte(filename)) // hash the filename using SHA256
	newFilename := hex.EncodeToString(hash[:]) // convert the hash to a hex string
	return newFilename // return the hex string
}

// newSHA256 returns the SHA256 hash of the given byte slice.
func newSHA256(data []byte) []byte {
	hash := sha256.Sum256(data) // hash the data using SHA256
	return hash[:] // return the hash as a byte slice
}

// prefixFilename transforms the filename by prefixing it with a string
func prefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename // return the transformed filename
	}
}

// transforms the filename by prefixing or hashing it
func main() {
	prefixServer := &Server{ // create a new server
		filenameTransformFunc: prefixFilename("BOB"), // prefix the filename with "BOB"
	}

	prefixRes := prefixServer.handleRequest("cool_picture.jpg")
	fmt.Println(prefixRes) // => "BOBcool_picture.jpg"

	hashServer := &Server{ // create a new server
		filenameTransformFunc: hashFilename, // hash the filename
	}
	hashRes := hashServer.handleRequest("cool_picture.jpg")
	fmt.Println(hashRes) // => "4e710fb390942d8870691282e709246a14964a7cd291028ae89f0c4ccd930c11"

}
