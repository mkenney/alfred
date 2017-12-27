package alfred

import (
	"net/http"
	"testing"
)

func TestServeComponent(t *testing.T) {
	tasks := make(map[string]Task)
	tasks["http.serve"] = Task{
		Serve: "8088",
	}
	context := InitialContext([]string{})
	go serve(tasks["http.serve"], context, tasks)
	response, err := http.Get("http://localhost:8088/serve_test.go")
	if err != nil {
		t.Fatalf("Expected a proper 200 response from localhost")
	}
	if response.StatusCode != 200 {
		t.Fatalf("Status code 200 expected")
	}
}
