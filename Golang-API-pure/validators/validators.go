package validators

func CheckStringField(result map[string]interface{}, fieldName string) string {
	field, ok := result[fieldName].(string)
	if !ok || len(field) < 1 {
			panic(fieldName + " is not a string or is an empty string")
	}
	return field
}

func CheckIntField(result map[string]interface{}, fieldName string) int {
	field, ok := result[fieldName].(float64) // JSON numbers are decoded as float64
	if !ok || field < 0 || field > 200 {
			panic(fieldName + " is not a number or is out of range")
	}
	return int(field)
}