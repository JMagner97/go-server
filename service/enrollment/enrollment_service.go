package enrollment

import (
	"context"
	"go-server/Model"
)

type EnrollmentService interface {
	Create(ctx context.Context, enrollment *Model.Enrollment)
	FindAll(ctx context.Context) ([]*Model.Enrollment, error)
	Delete(ctx context.Context, enrollment *Model.Enrollment)
}
