package pipeline

import "fmt"

type AndroidLint struct {
	Pattern             string
	HasAdvancedSettings bool
	FailedTotalNormal   float64
	FailedTotalHigh     float64
}

func (androidLint *AndroidLint) parseJSON(jsonString map[string]interface{}) error {
	for key, value := range jsonString {
		switch key {
		case "pattern":
			androidLint.Pattern = value.(string)
		case "failedTotalNormal":
			androidLint.HasAdvancedSettings = true
			androidLint.FailedTotalNormal = value.(float64)
		case "failedTotalHigh":
			androidLint.HasAdvancedSettings = true
			androidLint.FailedTotalHigh = value.(float64)
		default:
			return fmt.Errorf("Unknown key for AndroidLint plugin: got %#v for key %s", value, key)
		}
	}

	return nil
}

type Findbugs struct {
	Pattern string
}

func (findbugs *Findbugs) parseJSON(jsonString map[string]interface{}) error {
	for key, value := range jsonString {
		switch key {
		case "pattern":
			findbugs.Pattern = value.(string)
		default:
			return fmt.Errorf("Unknown key for Findbugs plugin: got %#v for key %s", value, key)
		}
	}

	return nil
}

type Pmd struct {
	Pattern string
}

func (pmd *Pmd) parseJSON(jsonString map[string]interface{}) error {
	for key, value := range jsonString {
		switch key {
		case "pattern":
			pmd.Pattern = value.(string)
		default:
			return fmt.Errorf("Unknown key for Pmd plugin: got %#v for key %s", value, key)
		}
	}

	return nil
}

type TaskPublisher struct {
	Pattern        string
	ExcludePattern string
}

func (taskPublisher *TaskPublisher) parseJSON(jsonString map[string]interface{}) error {
	for key, value := range jsonString {
		switch key {
		case "pattern":
			taskPublisher.Pattern = value.(string)
		case "excludePattern":
			taskPublisher.ExcludePattern = value.(string)
		default:
			return fmt.Errorf("Unknown key for TaskPublisher plugin: got %#v for key %s", value, key)
		}
	}

	return nil
}

type Violations struct {
	Checkstyle Checkstyle
}

type Checkstyle struct {
	Pattern string
}

func (violations *Violations) parseJSON(jsonString map[string]interface{}) error {
	for key, value := range jsonString {
		switch key {
		case "checkstyle":
			violations.Checkstyle.Pattern = (value.(map[string]interface{}))["pattern"].(string)
		default:
			return fmt.Errorf("Unknown key for TaskPublisher plugin: got %#v for key %s", value, key)
		}
	}

	return nil
}
