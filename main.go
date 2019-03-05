package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// YTURL ...
const YTURL = "http://www.youtube.com/get_video_info?video_id=%s&el=embedded&ps=default&eurl=&gl=US&hl=en"

func main() {
	port := ":" + os.Getenv("PORT")

	// main handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if u, ok := q["url"]; ok {
			if len(u) == 0 {
				http.Error(w, "empty url", 400)
				return
			}

			info, err := getYTInfo(u[0])
			if err != nil {
				http.Error(w, err.Error(), 400)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(info))
			return
		}

		// show index page
		fmt.Fprint(w, "YT - Youtube video link grabber")
	})

	log.Fatal(http.ListenAndServe(port, nil))
}

func getYTInfo(link string) (d []byte, err error) {
	ytUrl, err := getYTEndpoint(link)
	if err != nil {
		return
	}

	resp, err := http.Get(ytUrl)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	// key value pair chunks/slice
	chunks := strings.Split(string(body), "&")
	// info mapping
	info := map[string]string{}
	for _, kv := range chunks {
		kva := strings.Split(kv, "=")
		info[kva[0]], _ = url.QueryUnescape(kva[1])
	}

	videos := []map[string]string{}
	// comma-separated stream map
	cssm := strings.Split(info["url_encoded_fmt_stream_map"], ",")
	// range over through the list
	for _, sm := range cssm {
		video := map[string]string{}
		for _, um := range strings.Split(sm, "&") {
			kv := strings.Split(um, "=")
			video[kv[0]], _ = url.QueryUnescape(kv[1])
		}
		videos = append(videos, video)
	}

	d, err = json.MarshalIndent(videos, "  ", "  ")
	return
}

func getYTEndpoint(link string) (string, error) {
	uri := strings.Split(link, "watch?v=")
	if len(uri) < 2 {
		return "", errors.New("Invalid youtube URL")
	}

	ytID := uri[1]
	// remove extra params if passed in url includes other key/value pairs
	ytID = regexp.MustCompile("&.+").ReplaceAllString(ytID, "")

	return fmt.Sprintf(YTURL, ytID), nil
}
