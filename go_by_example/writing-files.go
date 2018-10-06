// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write_check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("This is a test data.\n1234567890\nqwertyuiopasdfghjklzxcvbnm\nQWERTYUIOPASDFGHJKLZXCVBNM")
	err := ioutil.WriteFile("/Users/dubby/Desktop/test/test.txt", d1, 0644)
	write_check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/Users/dubby/Desktop/test/test1.txt")
	write_check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	write_check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()

}
