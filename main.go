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
  "strings"
  //"strconv"
)
//------------------------------------------------------
  // globalvars
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
// read zonefile func
func readPrintZone(key string) (re []byte) {

    path := "./zones/" + key + ".zone"
  	ba ,err := readFile(path)
  	if err != nil {
  		fmt.Sprintln("Error: %s\n", err)
  	}

  return ba

}
// end read zonefile func
//------------------------------------------------------
func httpZone(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "# Github: https://github.com/7xz-ru/IP-Blacklist-Generator\n")
  fmt.Fprintf(w, "# Author: https://github.com/ZenDarkmaster\n")
  // get url keys z param
  keys, ok := r.URL.Query()["z"]

  if !ok || len(keys[0]) < 1 {
      fmt.Fprintf(w, "Url Param 'z' is missing")
      return
  }
    // Query()["key"] will return an array of items,
    // we only want the single item.
  key := keys[0]
  // end get url keys z param

  //zones list
  fmt.Fprintf(w, "# Feed zones: " + string(key) + "\n")
  //end zones list

  // string key to array words
  words := strings.Fields(key)
  // end string key to array words

  // start array cycle and print zonefile
  for i := 0; i < len(words); i++ {
      mkey := words[i]
      fmt.Fprintf(w, "# Zone - %s: \n%s", mkey, readPrintZone(mkey))
  }
  // end start array cycle and print zonefile
}
//------------------------------------------------------
func httpServer() {
  http.HandleFunc("/ipfeed", httpZone)
  http.ListenAndServe(":8080", nil)
}
//------------------------------------------------------
// for example, open the page, to add zones, use the z parameter in the url
// example: http://localhost:8080/ipfeed?z=gb+it+fr
