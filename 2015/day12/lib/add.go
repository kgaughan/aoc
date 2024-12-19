package lib

func AddNumbers(json interface{}) float64 {
	total := 0.0

	switch json.(type) {
	case float64:
		total += json.(float64)

	case []interface{}:
		for _, val := range json.([]interface{}) {
			total += AddNumbers(val)
		}

	case map[string]interface{}:
		for _, val := range json.(map[string]interface{}) {
			switch val.(type) {
			case string:
				if val.(string) == "red" {
					return 0.0
				}
			default:
				total += AddNumbers(val)
			}
		}
	}

	return total
}
