package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("POST", "http://webapps2.abc.utah.gov/Production/OnlineInventoryQuery/IQ/InventoryQuery.aspx", nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(&http.Cookie{Name: "__utma", Value: "54893812.1497259021.1504896406.1504896557.1506054629.2"})
	req.AddCookie(&http.Cookie{Name: "_ga", Value: "GA1.2.1497259021.1504896406"})
	req.AddCookie(&http.Cookie{Name: "ASP.NET_SessionId", Value: "yrlbwl520btrv5djnutjeabu"})
	req.Header.Add("Origin", "https://webapps2.abc.utah.gov")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("X-MicrosoftAjax", "Delta=true")
	req.Header.Add("Referer", "https://webapps2.abc.utah.gov/Production/OnlineInventoryQuery/IQ/InventoryQuery.aspx")

	fmt.Printf("%#v\n", req)

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
