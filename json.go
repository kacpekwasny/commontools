package commontools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// return list with missing keys
func CheckKeys(m map[string]string, keys ...string) []string {
	missing := []string{}
	for _, k := range keys {
		_, ok := m[k]
		if !ok {
			missing = append(missing, k)
		}
	}
	return missing
}

// json from Request to map[string]string
func Json2mapSS(r *http.Request) (map[string]string, error) {
	// read data that will be used to create an account
	reqBody, _ := ioutil.ReadAll(r.Body)
	m := &map[string]string{}
	err := json.Unmarshal(reqBody, m)
	return *m, err
}
