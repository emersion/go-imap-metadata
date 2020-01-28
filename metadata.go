// Package metadata implements the IMAP METADATA extension.
//
// IMAP METADATA is defined in RFC 5464.
package metadata

// Capability is the IMAP METADATA capbility.
const Capability = "METADATA"

// TODO: add MaxSize and Depth to GetMetadataOptions

type GetMetadataOptions struct{}
