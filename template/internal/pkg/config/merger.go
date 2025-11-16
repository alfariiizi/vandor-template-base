package config

import "reflect"

// DeepMerge merges src map into dst map recursively
// Later values override earlier values
func DeepMerge(dst, src map[string]interface{}) map[string]interface{} {
	if dst == nil {
		dst = make(map[string]interface{})
	}

	for key, srcVal := range src {
		if dstVal, exists := dst[key]; exists {
			// Both values exist, check if they're both maps
			srcMap, srcIsMap := srcVal.(map[string]interface{})
			dstMap, dstIsMap := dstVal.(map[string]interface{})

			if srcIsMap && dstIsMap {
				// Both are maps, merge recursively
				dst[key] = DeepMerge(dstMap, srcMap)
			} else {
				// Not both maps, src overrides dst
				dst[key] = srcVal
			}
		} else {
			// Key doesn't exist in dst, add it
			dst[key] = srcVal
		}
	}

	return dst
}

// MergeMultiple merges multiple maps in order (later overrides earlier)
func MergeMultiple(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, m := range maps {
		result = DeepMerge(result, m)
	}

	return result
}

// ToMap converts any struct to map[string]interface{} using reflection
func ToMap(v interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(v)

	// Handle pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result, nil
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// Get mapstructure tag or use field name
		key := field.Tag.Get("mapstructure")
		if key == "" || key == "-" {
			key = field.Name
		}

		// Convert field value to interface
		result[key] = fieldVal.Interface()
	}

	return result, nil
}
