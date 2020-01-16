package template_test

import (
	"logur.dev/logur"

	templateintegration "logur.dev/integration/template"
)

func ExampleNew() {
	logger := templateintegration.New(logur.NoopLogger{})

	// Output:
	_ = logger
}
