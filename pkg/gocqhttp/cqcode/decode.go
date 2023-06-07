package cqcode

import (
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func Find(msg string) []string {
	regex := regexp.MustCompile(`\[CQ:.*?]`)
	if regex == nil {
		log.Debug("[CqCode] Error compiling CQ code regex")
		return nil
	}
	result := regex.FindAllString(msg, -1)
	return result
}

func Decode(code string) map[string]string {
	// Return nil if code is not a CQ code
	if !strings.HasPrefix(code, "[CQ:") {
		log.Debug("[CqCode] Error compiling CQ code: " + code)
		return nil
	}

	// Trim prefix and suffix then split key-value pairs
	trimPrefix := strings.TrimPrefix(code, "[CQ:")
	trimSuffix := strings.TrimSuffix(trimPrefix, "]")
	splitKV := strings.Split(trimSuffix, ",")
	var codeMap = make(map[string]string)

	// Read type and store in map
	if strings.Contains(splitKV[0], "=") {
		log.Debug("[CqCode] Error reading CQ code type: " + code)
		return nil
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
			return nil
		}
		codeMap[split[0]] = split[1]
	}

	// Return map
	return codeMap
}
