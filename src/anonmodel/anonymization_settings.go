package anonmodel

import "fmt"

// AnonymizationSettings stores the settings about the dataset
type AnonymizationSettings struct {
	E         int    `json:"e" bson:"e"`     // Epsilon (waits for K + E intents)
	Max       int    `json:"max" bson:"max"` // Maximum EC size before split
	K         int    `json:"k" bson:"k"`
	Algorithm string `json:"algorithm" bson:"algorithm"`
	Mode      string `json:"mode" bson:"mode"`
}

func (settings *AnonymizationSettings) validate() error {
	if settings.K < 2 {
		return fmt.Errorf("The 'k' value should be at least 2, got: %v", settings.K)
	}

	if settings.Algorithm != "mondrian" && settings.Algorithm != "client-side" && settings.Algorithm != "client-side-custom" {
		return fmt.Errorf("The only currently supported anonymizations are 'mondrian', 'client-side' or 'client-side-custom', got '%v'", settings.Algorithm)
	}

	if settings.Mode != "single" && settings.Mode != "continuous" {
		return fmt.Errorf("Anonymization mode should be 'single' or 'continuous', got '%v'", settings.Mode)
	}

	return nil
}
