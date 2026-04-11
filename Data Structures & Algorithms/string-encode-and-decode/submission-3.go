import (
    "encoding/base64"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "empty"
	}

	jsonMarshal, _ := json.Marshal(strs)
	encodedString := base64.StdEncoding.EncodeToString(jsonMarshal)

	return encodedString
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "empty" {
		return []string{}
	}

	decodedBytes, _ := base64.StdEncoding.DecodeString(encoded)

    decodedMap := []string{}
    _ = json.Unmarshal(decodedBytes, &decodedMap)
	return decodedMap 
}
