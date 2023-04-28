package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// encodeToBase64 will take a value that is JSON encodable and encode it with base64.
func encodeToBase64(v any) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	err = encoder.Close()
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// decodeFromBase64 can take a v struct and decode a base64 encoded string into it.
func decodeFromBase64(v any, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}
