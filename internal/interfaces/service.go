package interfaces

type PredictionBuilder interface {
	Build(params map[string]string) map[string]interface{}
}
