package validator

// GetIdentifier returns the identifier for feature config
func (f FeatureConfig) GetIdentifier() string {
	return f.Feature
}

// GetIdentifier returns the identifier for a segment
func (s Segment) GetIdentifier() string {
	return s.Identifier
}

// GetIdentifier returns the identifier for a target
func (t Target) GetIdentifier() string {
	return t.Identifier
}
