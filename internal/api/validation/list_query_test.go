package validation

import (
	"testing"

	"server/internal/api/dto"
)

func TestProcessListQuery_Status(t *testing.T) {
	ok := "new"
	bad := "old"

	t.Run("nil status is ok", func(t *testing.T) {
		q := dto.ListTasksQuery{Limit: 10, Offset: 0}
		if err := ProcessListQuery(&q, 50, 200); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("valid status", func(t *testing.T) {
		q := dto.ListTasksQuery{Status: &ok, Limit: 10, Offset: 0}
		if err := ProcessListQuery(&q, 50, 200); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		q := dto.ListTasksQuery{Status: &bad, Limit: 10, Offset: 0}
		if err := ProcessListQuery(&q, 50, 200); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestProcessListQuery_Pagination(t *testing.T) {
	defLimit, maxLimit := 50, 200

	t.Run("limit <= 0 replaced by default", func(t *testing.T) {
		q := dto.ListTasksQuery{Limit: 0, Offset: 5}
		if err := ProcessListQuery(&q, defLimit, maxLimit); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if q.Limit != defLimit {
			t.Fatalf("limit = %d; want %d", q.Limit, defLimit)
		}
	})

	t.Run("limit > max_limit clamped", func(t *testing.T) {
		q := dto.ListTasksQuery{Limit: 999, Offset: 0}
		if err := ProcessListQuery(&q, defLimit, maxLimit); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if q.Limit != maxLimit {
			t.Fatalf("limit = %d; want %d", q.Limit, maxLimit)
		}
	})

	t.Run("negative offset becomes zero", func(t *testing.T) {
		q := dto.ListTasksQuery{Limit: 10, Offset: -10}
		if err := ProcessListQuery(&q, defLimit, maxLimit); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if q.Offset != 0 {
			t.Fatalf("offset = %d; want 0", q.Offset)
		}
	})
}
