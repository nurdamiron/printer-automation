package service

import (
    "context"
    "fmt"

    "github.com/google/uuid"
    "github.com/nurdamiron/printer-automation/internal/domain"
    "github.com/nurdamiron/printer-automation/internal/repository"
)

type PrintJobService interface {
    CreateJob(ctx context.Context, userID string, status string) (*domain.PrintJob, error)
    GetJob(ctx context.Context, id string) (*domain.PrintJob, error)
}

type printJobService struct {
    jobRepo repository.PrintJobRepository
}

func NewPrintJobService(jobRepo repository.PrintJobRepository) PrintJobService {
    return &printJobService{jobRepo: jobRepo}
}

func (s *printJobService) CreateJob(ctx context.Context, userID string, status string) (*domain.PrintJob, error) {
    jobUUID := uuid.New()
    userUUID, err := uuid.Parse(userID)
    if err != nil {
        return nil, fmt.Errorf("invalid user ID: %w", err)
    }

    job := &domain.PrintJob{
        ID:     jobUUID,
        UserID: userUUID,
        Status: status,
    }
    if err := s.jobRepo.Create(ctx, job); err != nil {
        return nil, err
    }
    return job, nil
}

func (s *printJobService) GetJob(ctx context.Context, id string) (*domain.PrintJob, error) {
    return s.jobRepo.GetByID(ctx, id)
}
