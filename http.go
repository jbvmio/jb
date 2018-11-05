package jb

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
)

// HClient here:
type HClient struct {
	url     string
	headers map[string]string
	body    []byte

	client  *http.Client
	request *http.Request
}

// HTTPClient returns a new HClient
func HTTPClient(url string, headers map[string]string, skipTLS bool) (*HClient, error) {
	var hc HClient
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTLS},
	}
	httpClient := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &hc, err
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	} else {
		headers = make(map[string]string)
	}
	hc.client = httpClient
	hc.headers = headers
	hc.request = req
	hc.url = url
	return &hc, nil
}

// GET is a convenience wrapper returning the response of a GET request
func (hc *HClient) GET(addr string) ([]byte, error) {
	var (
		retBody []byte
		err     error
	)
	hc.url = addr
	u, err := url.Parse(hc.url)
	if err != nil {
		return retBody, err
	}
	hc.request.URL = u
	hc.SetMethod("GET")
	resp, err := hc.client.Do(hc.request)
	if err != nil {
		return retBody, err
	}
	defer resp.Body.Close()
	retBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return retBody, err
	}
	return retBody, err

}

// Do here
func (hc *HClient) Do() ([]byte, error) {
	var (
		retBody []byte
		err     error
	)
	u, err := url.Parse(hc.url)
	if err != nil {
		return retBody, err
	}
	hc.request.URL = u
	hc.request.ContentLength = int64(len(hc.body))
	resp, err := hc.client.Do(hc.request)
	if err != nil {
		return retBody, err
	}
	defer resp.Body.Close()
	retBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return retBody, err
	}
	return retBody, err
}

// SetHeaders here
func (hc *HClient) SetHeaders(headers map[string]string) {
	if headers != nil {
		hc.headers = headers
		for k, v := range hc.headers {
			hc.request.Header.Add(k, v)
		}
	}
}

// GetHeaders here
func (hc *HClient) GetHeaders() map[string]string {
	return hc.headers
}

// SetMethod here
func (hc *HClient) SetMethod(method string) {
	hc.request.Method = method
}

// GetMethod here
func (hc *HClient) GetMethod() string {
	return hc.request.Method
}

// SetURL here
func (hc *HClient) SetURL(url string) {
	hc.url = url
}

// SetBody here
func (hc *HClient) SetBody(body string) {
	hc.body = []byte(body)
	hc.request.Body = ioutil.NopCloser(bytes.NewReader(hc.body))
	//hc.request.ContentLength = int64(len(hc.body))
}

// SetBodyBytes here
func (hc *HClient) SetBodyBytes(body []byte) {
	hc.body = body
	hc.request.Body = ioutil.NopCloser(bytes.NewReader(hc.body))
	//hc.request.ContentLength = int64(len(hc.body))
}

// GetBody here
func (hc *HClient) GetBody() string {
	return string(hc.body)
}

// GetBodyBytes here
func (hc *HClient) GetBodyBytes() []byte {
	return hc.body
}
