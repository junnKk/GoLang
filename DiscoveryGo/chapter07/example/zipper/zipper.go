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
	"https://cdn.mkhealth.co.kr/news/photo/202102/52163_52859_5928.jpg",
	"http://image.dongascience.com/Photo/2020/03/5bddba7b6574b95d37b6079c199d7101.jpg",
	"https://scontent-ssn1-1.cdninstagram.com/v/t51.2885-19/s320x320/241410622_242587151207694_3391323637909543013_n.jpg?_nc_ht=scontent-ssn1-1.cdninstagram.com&_nc_ohc=b2xRRUV4cVQAX-NksJJ&edm=ABfd0MgBAAAA&ccb=7-4&oh=554ca2e8f3269aa95bca9195721ac449&oe=61599398&_nc_sid=7bff83",
}

func main() {
	var wg sync.WaitGroup // 기본값 0으로 맞춰져 카운터가 들어가 있음
	// wg.Add(len(urls)) // 호출될 때마다 숫자를 더함
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done() // 사실상 wg.Add(-1)이라고 보면 됨
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
