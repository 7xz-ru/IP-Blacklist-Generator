package main
//------------------------------------------------------
import (
  //"net/http"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"io"
	"path/filepath"
  "strings"
  //"strconv"
  "flag"
  "github.com/valyala/fasthttp"
  "log"
)
//------------------------------------------------------
  // vars/globalvars
  var addr = flag.String("addr", "127.0.0.1:8080",
  	"TCP address to listen to for incoming connections")
//------------------------------------------------------
func main() {

	flag.Parse()

	s := fasthttp.Server{
		Handler: handler,
	}

	err := s.ListenAndServe(*addr)
	if err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}

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
//------------------------------------------------------
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
// convert byte data to string data
func BytesToString(data []byte) string {
  return string(data[:])
}
// end convert byte data to string data
//------------------------------------------------------
func httpZone(ctx *fasthttp.RequestCtx) {
  // author
  ctx.WriteString("# Github: https://github.com/7xz-ru/IP-Blacklist-Generator\n")
  ctx.WriteString("# Author: https://github.com/ZenDarkmaster\n")
  // end author

  // get url keys z param
  keys := ctx.QueryArgs().Peek("z")
  key := BytesToString(keys)
  // end get url keys z param

  //zones list
  ctx.WriteString("# Feed zones: " + string(key) + "\n")
  //end zones list

  // string key to array words
  words := strings.Fields(key)
  // end string key to array words

  // start array cycle and print zonefile
  for i := 0; i < len(words); i++ {
      mkey := words[i]
      fmt.Fprintf(ctx, "# Zone - %s: \n%s", mkey, readPrintZone(mkey))
  }
  // end start array cycle and print zonefile
}
//------------------------------------------------------
func handler(ctx *fasthttp.RequestCtx) {
	httpZone(ctx)
}
//------------------------------------------------------
// For example, open the page, to add zones, use the z parameter in the url
// Example: http://localhost:8080/?z=gb+it+fr
