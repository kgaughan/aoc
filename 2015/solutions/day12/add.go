package day12

func AddNumbers(json interface{}, part2 bool) float64 {
	total := 0.0

	switch v := json.(type) {
	case float64:
		total += v

	case []interface{}:
		for _, val := range v {
			total += AddNumbers(val, part2)
		}

	case map[string]interface{}:
		for _, val := range v {
			switch v := val.(type) {
			case string:
				if part2 && v == "red" {
					return 0.0
				}
			default:
				total += AddNumbers(val, part2)
			}
		}
	}

	return total
}
