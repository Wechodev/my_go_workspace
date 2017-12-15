package main

import "net/http"

func main()  {
	client := &http.Client{
		CheckRedirect: redirectPoilcyFunc,
	}

	resp, err := client.Get("http://xx.com")
	req, err := http.NewRequest("GET", "http://xxx.com", nil)

	req.Header.Add("User-Agent", "Our Custom User-Agent")

	resp, err := client.Do(req)
}
