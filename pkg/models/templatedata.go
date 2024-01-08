package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap                        map[string]string
	IntMap                           map[string]int
	FloatMap                         map[string]float32
	Data                             map[string]interface{} // We use 'interface' in Go when we are not sure of the data type
	CSRFToken, Flash, Warning, Error string
}
