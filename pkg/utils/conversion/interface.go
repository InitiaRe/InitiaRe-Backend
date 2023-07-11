package conversion

// Get value from interface. Return default value if key not exists
func GetFromInterface(src map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value, exists := src[key]
	if !exists {
		return defaultValue
	}
	return value
}


