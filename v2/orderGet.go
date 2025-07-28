package v2

import (
	"context"
	"fmt"
	"net/http"
)

func (c *clientImpl) OrderGet(ctx context.Context, cdekNumber *int64, imNumber *string) (*Response, error) {
	uri := "/v2/orders"

	if cdekNumber != nil || imNumber != nil {
		uri += "?"
		if cdekNumber != nil {
			uri += fmt.Sprintf("&cdek_number=%d", *cdekNumber)
		}
		if imNumber != nil {
			uri += fmt.Sprintf("&im_number=%s", *imNumber)
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.buildUri(uri, nil),
		nil,
	)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	accessToken, err := c.getAccessToken(ctx)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	return jsonReq[Response](req)
}
