package metadata

import (
	"fmt"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/responses"
)

type MetadataResponse struct {
	Mailbox string
	Entries map[string]string
}

func (r *MetadataResponse) Handle(resp imap.Resp) error {
	name, fields, ok := imap.ParseNamedResp(resp)
	if !ok || name != "METADATA" {
		return responses.ErrUnhandled
	}
	if len(fields) != 2 {
		return fmt.Errorf("expected 2 fields in METADATA response")
	}

	mailbox, err := imap.ParseString(fields[0])
	if err != nil {
		return fmt.Errorf("failed to parse mailbox in METADATA response: %v", err)
	}
	if mailbox != r.Mailbox {
		return responses.ErrUnhandled
	}

	values, ok := fields[1].([]interface{})
	if !ok {
		return fmt.Errorf("expected METADATA values field to be a list")
	}
	if len(values)%2 != 0 {
		return fmt.Errorf("expected an even number of key/value items in METADATA response")
	}
	for i := 0; i < len(values); i += 2 {
		k, err := imap.ParseString(values[i])
		if err != nil {
			return fmt.Errorf("failed to parse METADATA entry name: %v", err)
		}

		v := values[i+1]
		if v == nil {
			continue
		}

		s, err := imap.ParseString(v)
		if err != nil {
			return fmt.Errorf("failed to parse METADATA entry value: %v", err)
		}

		r.Entries[k] = s
	}

	return nil
}
