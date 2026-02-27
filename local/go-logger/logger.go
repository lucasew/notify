package logger

import "io"

type LogLevel int

const LogLevelDebug1 LogLevel = 1

type Logger interface {
	Debug(string, ...interface{})
	Error(...interface{})
}

type Generator struct{}

func NewGenerator(w io.Writer) *Generator { return &Generator{} }

func (g *Generator) SetDebugLevel(l LogLevel) {}

func (g *Generator) New(module string) Logger { return &mockLogger{} }

type mockLogger struct{}

func (m *mockLogger) Debug(format string, args ...interface{}) {}
func (m *mockLogger) Error(args ...interface{}) {}
