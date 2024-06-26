package internal

import (
	"github.com/open-component-model/ocm/pkg/contexts/credentials"
)

type AccessSpecInfo struct {
	Short      string                       `json:"short"`
	MediaType  string                       `json:"mediaType"`
	Hint       string                       `json:"hint"`
	ConsumerId credentials.ConsumerIdentity `json:"consumerId"`
}
