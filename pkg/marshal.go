package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func MarshalToString(data interface{}) (string, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)

	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "")

	if err := encoder.Encode(data); err != nil {
		return "", fmt.Errorf("error marshalling data to one-line JSON: %w", err)
	}

	return string(buffer.Bytes()), nil
}

func HashData(value string) string {
	hasher := sha256.New()
	hasher.Write([]byte(value))
	hashedEmail := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedEmail)

	return hashedString
}
