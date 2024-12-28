package day12

func AddNumbers(json interface{}, part2 bool) float64 {
	total := 0.0

	switch json.(type) {
	case float64:
		total += json.(float64)

	case []interface{}:
		for _, val := range json.([]interface{}) {
			total += AddNumbers(val, part2)
		}

	case map[string]interface{}:
		for _, val := range json.(map[string]interface{}) {
			switch val.(type) {
			case string:
				if part2 && val.(string) == "red" {
					return 0.0
				}
			default:
				total += AddNumbers(val, part2)
			}
		}
	}

	return total
}
