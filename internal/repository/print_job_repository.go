package repository

import (
    "context"

    "github.com/nurdamiron/printer-automation/internal/domain"
    "gorm.io/gorm"
)

type PrintJobRepository interface {
    Create(ctx context.Context, job *domain.PrintJob) error
    GetByID(ctx context.Context, id string) (*domain.PrintJob, error)
    // При необходимости: Update, Delete, List...
}

type printJobRepository struct {
    db *gorm.DB
}

func NewPrintJobRepository(db *gorm.DB) PrintJobRepository {
    return &printJobRepository{db: db}
}

func (r *printJobRepository) Create(ctx context.Context, job *domain.PrintJob) error {
    return r.db.WithContext(ctx).Create(job).Error
}

func (r *printJobRepository) GetByID(ctx context.Context, id string) (*domain.PrintJob, error) {
    var job domain.PrintJob
    if err := r.db.WithContext(ctx).First(&job, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &job, nil
}
