package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var timer = time.NewTicker(500 * time.Microsecond)


func Fetch(url string) ([]byte, error) {
	<- timer.C
	response, err := http.Get(url)
	if err !=  nil {
		return nil, fmt.Errorf("获取网页失败: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("获取网页失败：%v", response.StatusCode)
	}

	responseRead := bufio.NewReader(response.Body)

	encode, ecer := determineEncoding(responseRead)
	if ecer != nil {
		return nil, fmt.Errorf("%v", ecer)
	}

	transReader := transform.NewReader(responseRead, encode.NewDecoder())
	content, _ := ioutil.ReadAll(transReader)
	return content, nil
}

func determineEncoding(reader *bufio.Reader) (encoding.Encoding, error)  {
	peek, err := reader.Peek(1024)
	if err != nil {
		log.Printf("")
		return unicode.UTF8, nil
	}

	e, _, _ := charset.DetermineEncoding(peek, "")
	return e, nil
}
