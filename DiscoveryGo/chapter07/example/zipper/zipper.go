package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

var urls = []string{
	"http://images.freeimages.com/images/previews/587/disco-ball-1421094.jpg",
	"http://images.freeimages.com/images/previews/2c2/carnival-1434122.jpg",
	"http://images.freeimages.com/images/previews/43b/orange-smoothie-1-1381411.jpg",
}

func main() {
	var wg sync.WaitGroup // 기본값 0으로 맞춰져 카운터가 들어가 있음
	// wg.Add(len(urls)) // 호출될 때마다 숫자를 더함
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg") // 파일들의 리스트를 뽑는데 사용. 현재 디렉터리의 .jpg 파일들
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("images.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}

// download downloads url and returns the contents and error.
func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

// urlToFilename returns the filename part from the rawurl.
func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil // 경로에서 가장 마지막 부분을 반환
}

// writeZip writes a zip archive file.
func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}
