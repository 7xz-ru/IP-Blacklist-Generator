package main
//------------------------------------------------------
import (
  "net/http"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"io"
	"path/filepath"
)
//------------------------------------------------------
var Zone string = "test"
//------------------------------------------------------
func main() {
  httpServer()
}
//------------------------------------------------------
func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	pullPath := filepath.Join(parentPath, path)
	file, err := os.Open(pullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	return read(file)
}

func read(fd_r io.Reader) ([]byte, error) {
	br := bufio.NewReader(fd_r)
	var buf bytes.Buffer

	for {
		ba, isPrefix, err := br.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}

	}
	return buf.Bytes(), nil
}
//------------------------------------------------------
func httpZone(w http.ResponseWriter, r *http.Request) {
  //fmt.Fprintf(w, "Github: ZenDarkmaster")

  path := "./zones/" + Zone + ".zone"
	ba ,err := readFile(path)
	if err != nil {
		fmt.Fprintln(w, "Error: %s\n", err)
	}
	fmt.Fprintf(w, "# Country ZONE - '%s' : \n%s\n", Zone, ba)

}
//------------------------------------------------------
func httpServer() {
  http.HandleFunc("/", httpZone)
  http.ListenAndServe(":8080", nil)
}
//------------------------------------------------------
