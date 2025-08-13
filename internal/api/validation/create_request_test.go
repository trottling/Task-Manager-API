package validation

import (
	"testing"

	"server/internal/api/dto"
)

func TestValidateCreateTaskRequest(t *testing.T) {
	ok := dto.CreateTaskRequest{
		Name:   "Do something",
		Status: "new",
	}
	if err := ValidateCreateTaskRequest(ok); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cases := []struct {
		name string
		req  dto.CreateTaskRequest
	}{
		{"empty name", dto.CreateTaskRequest{Name: "", Status: "new"}},
		{"invalid status", dto.CreateTaskRequest{Name: "X", Status: "weird"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := ValidateCreateTaskRequest(tc.req); err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
