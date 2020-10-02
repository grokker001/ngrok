// +build !release,!autoupdate

package client

import (
	"github.com/grokker001/ngrok/src/ngrok/client/mvc"
)

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
