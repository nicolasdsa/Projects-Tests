package validators

import (
	"fmt"
	"strconv"
	"strings"
)

func Validate(elements map[string]interface{}, rules map[string]string) error {
	for field, rule := range rules {
		val, exists := elements[field]
		if !exists {
			return fmt.Errorf("Field '%s' is missing", field)
		}

		rules := strings.Split(rule, "|")
		for _, r := range rules {
			if err := applyRule(field, val, r); err != nil {
				return err
			}
		}
	}
	return nil
}

func applyRule(field string, value interface{}, rule string) error {
	if rule == "nullable" && value == nil {
		return nil // Se o campo for nulo e a regra for nullable, passe na validação
	}

	parts := strings.Split(rule, ":")
	switch parts[0] {
	case "required":
		if value == nil {
			return fmt.Errorf("Field '%s' is required", field)
		}
	case "int":
		if _, err := strconv.Atoi(fmt.Sprintf("%v", value)); err != nil {
			return fmt.Errorf("Field '%s' must be an integer", field)
		}
	case "min":
		min, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("Invalid rule format for 'min' on field '%s'", field)
		}
		if len(fmt.Sprintf("%v", value)) < min {
			return fmt.Errorf("Field '%s' must be at least %s characters long", field, parts[1])
		}
	case "max":
		max, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("Invalid rule format for 'max' on field '%s'", field)
		}
		if len(fmt.Sprintf("%v", value)) > max {
			return fmt.Errorf("Field '%s' must be at most %s characters long", field, parts[1])
		}
	case "string":
		if _, ok := value.(string); !ok {
			return fmt.Errorf("Field '%s' must be a string", field)
		}
	case "boolean":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("Field '%s' must be a boolean", field)
		}
	default:
		return fmt.Errorf("Unknown validation rule '%s' for field '%s'", parts[0], field)
	}
	return nil
}