// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/client/logs.go

package client

import (
	"log"
	"net/http"
	"time"
)

// API request/response payloads may contain generated server credentials.
// Keep diagnostic metadata only; URLs, headers and bodies are deliberately omitted.
func (c *Client) logRequest(method, url string, body []byte, headers http.Header) {
	log.Printf("[DEBUG] HTTP request method=%s", method)
}
func (c *Client) logResponse(method, url, status string, body []byte, headers http.Header, duration time.Duration) {
	log.Printf("[DEBUG] HTTP response method=%s status=%s duration=%v", method, status, duration)
}
func (c *Client) logError(method, url string, err error, duration time.Duration) {
	log.Printf("[ERROR] HTTP request failed method=%s duration=%v", method, duration)
}
