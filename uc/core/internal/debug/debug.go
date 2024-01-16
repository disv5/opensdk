package debug

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
)

// PrintError debug print error
func PrintError(err error, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [ERROR]", err)
}

// PrintStringResponse debug print string response
func PrintStringResponse(str string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [RESPONSE]", str)
}

// PrintGetRequest debug print get request
func PrintGetRequest(url string, debug bool) {
	if !debug {
		return
	}
	log.Println("[DEBUG] [API] GET", url)
}

// PrintPostJSONRequest debug print post json request
func PrintPostJSONRequest(url string, body []byte, debug bool) {
	if !debug {
		return
	}
	const format = "[DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}

// PrintPostMultipartRequest debug print post multipart request
func PrintPostMultipartRequest(url string, mp map[string]string, debug bool) {
	if !debug {
		return
	}
	body, _ := json.Marshal(mp)
	const format = "[DEBUG] [API] multipart/form-data POST %s\n" +
		"http request body:\n%s\n"

	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}

// DecodeJSONHttpResponse decode http json response with debug
func DecodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) error {
	if !debug {
		return json.NewDecoder(r).Decode(v)
	}
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}

// 将对应格式文本转换成utf-8
func DecodeEncoding(r io.Reader) *transform.Reader {
	e := determineEncodeing(r)
	return transform.NewReader(r, e.NewDecoder())
}

// 判断传输来的文本的字符集格式是什么
func determineEncodeing(r io.Reader) encoding.Encoding {
	peek, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	determineEncoding, _, _ := charset.DetermineEncoding(peek, "")
	return determineEncoding
}


func DecodeStingHttpResponse(r io.Reader, resp *[][]string, debug bool) error {
	body, err := ioutil.ReadAll(r)
	if err !=nil {
		return err
	}

	// 转码
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	a := csv.NewReader(reader)
	body2, err := a.ReadAll()
	if err !=nil {
		return err
	}
	*resp = body2
	if debug {
		log.Printf("[DEBUG] [API] http response body:\n%s\n", body2)
    }

	return nil
}
