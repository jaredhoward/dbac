package list

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func GetListFromBytes(input []byte) (*List, error) {
	return parseAlphaPriceListPDF(bytes.NewReader(input))
}

func GetAlphaPriceList(loc string) ([]byte, error) {
	if loc == "" {
		loc = DefaultInputAlphaPriceList
	}

	resp, err := http.Get(loc)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
