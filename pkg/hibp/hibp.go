package hibp

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gopwdhttp "github.com/pvwnthem/gopwd/pkg/http"
)

func Check(p string) (bool, error) {
	// Full disclaimer: I stole half of this and the other half is really shitty :); I'll fix it later (maybe)
	if len(p) == 0 {
		return false, nil
	}

	hashed := hash(p)
	prefix := hashed[:5]
	suffix := hashed[5:]

	url, err := url.Parse("https://api.pwnedpasswords.com/")
	if err != nil {
		return false, err
	}

	httpClient := gopwdhttp.New(http.DefaultClient)

	request, err := httpClient.Request("GET", url.String()+fmt.Sprintf("range/%s", prefix), nil)
	if err != nil {
		return false, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return false, err
	}

	for _, target := range response {
		if len(target) < 36 {
			continue
		}

		if target[:35] == suffix {
			if _, err = strconv.ParseInt(target[36:], 10, 64); err != nil {
				return false, err
			}

			return true, err
		}
	}

	return false, err

}

func hash(v string) string {
	algo := sha1.New()

	algo.Write([]byte(v))
	return strings.ToUpper(hex.EncodeToString(algo.Sum(nil)))
}
