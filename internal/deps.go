package internal

//go:generate mockery --name=Logger --output=./mocks --case=underscore

type Logger interface {
	Print(i ...interface{})
	Debug(i ...interface{})
	Fatal(i ...interface{})
}
