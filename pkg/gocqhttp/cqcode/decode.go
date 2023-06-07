package cqcode

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func Decode(code string) any {
	// Return empty string if code is not a CQ code
	if !strings.HasPrefix(code, "[CQ:") {
		log.Debug("[CqCode] Error compiling CQ code: " + code)
		return ""
	}

	// Trim prefix and suffix then split key-value pairs
	trimPrefix := strings.TrimPrefix(code, "[CQ:")
	trimSuffix := strings.TrimSuffix(trimPrefix, "]")
	splitKV := strings.Split(trimSuffix, ",")
	var codeMap = make(map[string]string)

	// Read type and store in map
	if strings.Contains(splitKV[0], "=") {
		log.Debug("[CqCode] Error reading CQ code type: " + code)
		return ""
	}
	codeMap["code-type"] = splitKV[0] // type

	// Read key-value pairs and store in map
	for _, v := range splitKV {
		if !strings.Contains(v, "=") {
			continue
		}
		split := strings.Split(v, "=")
		if len(split) != 2 {
			log.Debug("[CqCode] Error splitting CQ code key-value pairs: " + code)
			return ""
		}
		codeMap[split[0]] = split[1]
	}

	// Return map
	return codeMap
}
