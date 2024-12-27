package coinmarketcap

import (
	"context"
	"net/http"

	"golang.org/x/net/html"
)

func (s *cmcService) isDataNode(attr []html.Attribute) bool {
	for _, v := range attr {
		if v.Key == "id" {
			return v.Val == "__NEXT_DATA__"
		}
	}
	return false
}

func (s *cmcService) findData(node *html.Node) *string {
	if s.isDataNode(node.Attr) {
		return &node.FirstChild.Data
	}

	child := node.FirstChild
	for child != nil {
		data := s.findData(child)
		if data != nil {
			return data
		}

		child = child.NextSibling
	}

	return nil
}

func (s *cmcService) getHTML(ctx context.Context, fullURL string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(
		ctx, "GET", fullURL, nil,
	)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
