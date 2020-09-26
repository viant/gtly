package gtly

//Option represents field option
type Option func(field *Field)

//DateLayoutOpt field with data layout option
func DateLayoutOpt(layout string) Option {
	return func(field *Field) {
		field.DataLayout = layout
	}
}

//ProviderOpt return a field provider option
func ProviderOpt(provider *Provider) Option {
	return func(field *Field) {
		field.provider = provider
	}
}

//OmitEmptyOpt returns a field omit empty option
func OmitEmptyOpt(omitEmpty bool) Option {
	return func(field *Field) {
		field.OmitEmpty = &omitEmpty
	}
}

//ComponentTypeOpt return a field component type option
func ComponentTypeOpt(componentType string) Option {
	return func(field *Field) {
		field.ComponentType = componentType
	}
}
