package api

type Client struct {
    BaseURL string
}

func NewClient(baseURL string) *Client {
    return &Client{BaseURL: baseURL}
}

func (c *Client) UploadDiff(diff string) error {
    // Placeholder for uploading diff
    return nil
}

