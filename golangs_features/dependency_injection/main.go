package main

import (
	"context"
	"fmt"
)


type UserRepository interface {
	GetEmail(ctx context.Context, userID int) (string, error)
}

type Mailer interface {
	Send(ctx context.Context, email string, message string) (error)
}

type UserLogger interface {
	Info(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}

type UserService struct {
	repo UserRepository
	mailer Mailer
	logger UserLogger
}

func NewUserService(userRepo UserRepository, userMailer Mailer, userLogger UserLogger) *UserService {
    return &UserService{
		repo: userRepo,
		mailer: userMailer,
		logger: userLogger,
	}
}

func (s *UserService) NotifyUser(ctx context.Context, userID int, message string) error {
	email, err := s.repo.GetEmail(ctx, userID)
	if err != nil {
		s.logger.Error(ctx, "failed to get email, error:", err.Error())
		return fmt.Errorf("failed to get email with user id: %d, error: %w", userID, err)
	}

	if err = s.mailer.Send(ctx, email, message); err != nil {
		s.logger.Error(ctx, "failed to send email, error:", err.Error())
		return fmt.Errorf("failed to send message: %w", err)
	}
	
	s.logger.Info(ctx, "successful sending email")

    return nil
}

func main() {
    // Здесь не нужно писать реализацию БД или API, 
    // просто покажи структуру кода.
}