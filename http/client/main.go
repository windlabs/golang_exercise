package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	httpRequest()
}

func httpRequest() {
	urlObj, _ := url.Parse("http://127.0.0.1:8080/test")
	data := url.Values{}
	data.Set("name", "jq")
	data.Set("abc[0]", "def")
	data.Set("abc[1]", "ghi")
	params := data.Encode()
	urlObj.RawQuery = params
	req, err := http.NewRequest("GET", urlObj.String(), nil)

	if err != nil {
		fmt.Printf("http request failed, err:%v.", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("http client request failed, err:%v.", err)
		return
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read failed, err:%v.", err)
		return
	}
	fmt.Printf(string(res))
}

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8080/test?abc[]=def&abc[]=efg&name=jq")
	if err != nil {
		fmt.Printf("get Url Content failed,err:%v.", err)
		return
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read response failed, err:%v.", err)
		return
	}
	fmt.Println(string(res))
}
