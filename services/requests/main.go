package requests

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Get(rawURL string, params map[string]string) ([]byte, error) {
	var result []byte

	endpoint, err := url.Parse(rawURL)
	if err != nil {
		log.Println(err)
		return result, err
	}

	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	endpoint.RawQuery = values.Encode()

	resp, err := http.Get(endpoint.String())
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return result, err
	}

	return respBody, nil
}
