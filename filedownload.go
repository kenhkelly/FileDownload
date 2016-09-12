package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const version string = "v0.3"

var dirPath string

func main() {
	dirPathPtr := flag.String("dir", "", "[optional] Set directory on run")
	portPtr := flag.String("port", "8080", "[optional] Set port to run on (default 8080)")
	flag.Parse()
	dirPath = *dirPathPtr
	port := *portPtr

	if len(dirPath) > 0 {
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			panic(fmt.Sprintf("Specified directory '%s' does not exist", dirPath))
		}
	} else {
		dirPath, _ = os.Getwd()
	}

	fmt.Println("Directory set at: " + dirPath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Usage: /get?url=<url>")
	})

	http.HandleFunc("/get", get)
	fmt.Println("FileDownload " + version)
	fmt.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func get(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	getUrl := r.URL.Query().Get("url")

	if len(getUrl) == 0 {
		fmt.Fprintf(w, "URL missing")
		return
	}

	u, err := url.Parse(getUrl)
	ufrags := strings.Split(u.Path, "/")
	fname := ufrags[len(ufrags)-1]
	filename := fmt.Sprintf("%s/%s", dirPath, fname)
	i := 0
	for {
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			filename = fmt.Sprintf("%s/%v.%s", dirPath, i, fname)
		} else {
			break
		}
		i++
	}

	c := http.Client{}
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		panic("Err: " + err.Error())
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7")
	resp, err := c.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error getting url: %s", err)
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(w, "Error writing file: %s", err)
	}

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error writing file: %s", err)
	}

	end := time.Now()
	msg := fmt.Sprintf("Wrote %v bytes", size)

	fmt.Fprint(w, msg)

	msg = fmt.Sprintf("%s to %s in %v", msg, filename, end.Sub(start))
	fmt.Println(msg)
}
