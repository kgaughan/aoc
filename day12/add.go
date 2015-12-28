package day12

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
			total += AddNumbers(val)
		}
	}

	return total
}

