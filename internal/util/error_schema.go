package util

type ErrorSchema struct {
	Code        string
	Status      int
	Origin      string
	ClientError []string
	LogError    []string
}
