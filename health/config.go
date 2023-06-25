package health

import (
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
	"unicode/utf8"
)

func starGen(str any) string {
	// Get string length
	var strLen = 8
	switch str.(type) {
	case string:
		strLen = utf8.RuneCountInString(str.(string))
	case int:
		for i := str.(int); i > 0; i /= 10 {
			strLen += 1
		}
	}

	// Generate stars
	var res string
	for i := 0; i < strLen; i++ {
		res += "*"
	}
	return res
}

func convert(i any) any {
	switch x := i.(type) {
	case map[any]any:
		m2 := map[string]any{}
		for k, v := range x {
			reg := regexp.MustCompile(`.*name|pass|uri|host|dir|_db.*`)
			if reg == nil {
				return nil
			}
			if reg.MatchString(k.(string)) {
				m2[k.(string)] = starGen(v)
				continue
			}
			m2[k.(string)] = convert(v)
		}
		return m2
	case []any:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func Config() any {
	// Read config file
	cfg, err := os.ReadFile("config.yaml")
	if err != nil {
		return err.Error()
	}

	// Convert yaml to interface{}
	var body any
	if err := yaml.Unmarshal(cfg, &body); err != nil {
		return nil
	}
	body = convert(body)

	return body
}
