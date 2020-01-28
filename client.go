package metadata

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// Client is a METADATA client.
type Client struct {
	c *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{c}
}

func (c *Client) SupportMetadata() (bool, error) {
	return c.c.Support(Capability)
}

func (c *Client) GetMetadata(mbox string, entries []string, options *GetMetadataOptions) (map[string]string, error) {
	if s := c.c.State(); s != imap.AuthenticatedState && s != imap.SelectedState {
		return nil, client.ErrNotLoggedIn
	}

	cmd := &GetMetadataCommand{
		Mailbox: mbox,
		Entries: entries,
		Options: options,
	}
	res := &MetadataResponse{
		Mailbox: mbox,
		Entries: make(map[string]string, len(entries)),
	}
	status, err := c.c.Execute(cmd, res)
	if err != nil {
		return nil, err
	}

	return res.Entries, status.Err()
}

func (c *Client) SetMetadata(mbox string, entries map[string]string) error {
	if s := c.c.State(); s != imap.AuthenticatedState && s != imap.SelectedState {
		return client.ErrNotLoggedIn
	}

	cmd := &SetMetadataCommand{
		Mailbox: mbox,
		Entries: entries,
	}
	status, err := c.c.Execute(cmd, nil)
	if err != nil {
		return err
	}

	return status.Err()
}
