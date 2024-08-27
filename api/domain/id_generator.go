package domain

//go:generate moq -fmt goimports -out ./id_generator_mock.go -rm . IDGenerator

type IDGenerator interface {
	Generate() string
}
