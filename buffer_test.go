package linkedbuffer

import (
	"bytes"
	"crypto/rand"
	. "github.com/itsabgr/go-handy"
	"io"
	"io/ioutil"
	"testing"
)

func writeRandomBytes(writer io.Writer, n uint) [][]byte {
	testBytes := make([][]byte, n)
	for _, b := range testBytes {
		_, err := rand.Read(b)
		Throw(err)
		_, err = writer.Write(b)
		Throw(err)
	}
	return testBytes
}
func TestBuffer_ReadWrite(t *testing.T) {
	defer Catch(func(rec interface{}) {
		t.Fatal(rec)
	})
	buf := &Buffer{}
	tobeBytes := bytes.Join(writeRandomBytes(buf, 10), []byte{})
	if buf.Len() != uint(len(tobeBytes)) {
		t.Fatal("wrong Len")
	}
	readBytes, err := ioutil.ReadAll(buf)
	Throw(err)
	if !bytes.Equal(readBytes, tobeBytes) {
		t.Fatal("broken Read or Write")
	}
}
