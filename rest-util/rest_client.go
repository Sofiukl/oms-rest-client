package restutil

import (
	"fmt"
	"net/http"
)

//Get calls the get rest call for url specified
func Get(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Fail to get %s", url)
		err := fmt.Errorf("fail to get %s", url)
		return nil, err
	}

	fmt.Println(resp)
	return resp, nil
}
