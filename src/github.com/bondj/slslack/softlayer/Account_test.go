package softlayer

import (
	"net/http"
	"testing"
)

var user = "slack-bondj-csf"
var key = "bccfe27273aa44102d319dc547c950b02ff4fc86d2ee5854b0643928f891d806"

func TestEvents(t *testing.T) {
	var client = &http.Client{}
	GetRecentEvents(client, user, key)
}

func TestVirtual(t *testing.T) {
	var client = &http.Client{}
	GetDownVirtualGuests(client, user, key)
}
