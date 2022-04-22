package validator

func (f FeatureConfig) GetIdentifier() string {
	return f.Feature
}

func (s Segment) GetIdentifier() string {
	return s.Identifier
}

func (t Target) GetIdentifier() string {
	return t.Identifier
}
