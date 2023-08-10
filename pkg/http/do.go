package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pvwnthem/gopwd/constants"
)

func (c *Client) Do(req *http.Request) ([]string, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if resErr := resp.Body.Close(); err == nil {
			err = resErr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(constants.ErrUnexpectedApiResponse, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.ReplaceAll(string(body), "\r\n", "\n"), "\n"), err
}
