package metadata

import (
	"strings"

	"github.com/emersion/go-imap"
)

type GetMetadataCommand struct {
	Mailbox string
	Entries []string
	Options *GetMetadataOptions
}

func (cmd *GetMetadataCommand) Command() *imap.Command {
	args := []interface{}{
		imap.FormatMailboxName(cmd.Mailbox),
		imap.FormatStringList(cmd.Entries),
	}

	return &imap.Command{
		Name:      "GETMETADATA",
		Arguments: args,
	}
}

type SetMetadataCommand struct {
	Mailbox string
	Entries map[string]string
}

func (cmd *SetMetadataCommand) Command() *imap.Command {
	var entries []interface{}
	for k, v := range cmd.Entries {
		entries = append(entries, k, strings.NewReader(v))
	}

	args := []interface{}{
		imap.FormatMailboxName(cmd.Mailbox),
		entries,
	}

	return &imap.Command{
		Name:      "SETMETADATA",
		Arguments: args,
	}
}
