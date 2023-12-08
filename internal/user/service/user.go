package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/user/handler/consumer/dto"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/user/entity"
	"github.com/murat96k/kitaptar.kz/pkg/util"
	"github.com/redis/go-redis/v9"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) (string, error) {

	id, err := m.Repository.CreateUser(ctx, u)
	if err != nil {
		return "", err
	}

	// Seed the random number generator
	//nolint
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit code
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	b, err := json.Marshal(&dto.UserCode{Code: code, Email: u.Email})
	if err != nil {
		return "", fmt.Errorf("failed to marshall UserCode err: %w", err)
	}

	m.userVerificationProducer.ProduceMessage(b)

	log.Println("Email sent successfully")

	return id, nil
}

func (m *Manager) VerifyToken(token string) (string, error) {

	claim, err := m.Token.ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("validate token err: %w", err)
	}

	return claim.UserID, nil
}

func (m *Manager) UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error {

	user, err := m.Repository.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Password != "" {
		req.Password, err = util.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user.Password = req.Password
	}

	err = m.Cache.UserCache.DeleteUser(ctx, user.Id.String())
	if err != nil {
		return err
	}

	err = m.Repository.UpdateUser(ctx, id, req)
	if err != nil {
		return err
	}

	_ = m.Cache.UserCache.SetUser(ctx, user)

	return nil
}

func (m *Manager) GetUserById(ctx context.Context, id string) (*entity.User, error) {

	user, err := m.Cache.UserCache.GetUser(ctx, id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = m.Repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	user, err := m.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) DeleteUser(ctx context.Context, id string) error {

	err := m.Cache.UserCache.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return m.Repository.DeleteUser(ctx, id)
}

func (m *Manager) SendConfirmCode(email, code string) error {

	// Set up authentication information
	auth := smtp.PlainAuth("", m.Config.SMTP.Username, m.Config.SMTP.Password, m.Config.SMTP.Host)

	// Compose the email content
	subject := "Confirmation Code for Registration"
	body := fmt.Sprintf("Your confirmation code is: %s", code)
	message := []byte("Subject: " + subject + "\r\n\r\n" + body)

	// Send the email
	err := smtp.SendMail(fmt.Sprintf("%s:%d", m.Config.SMTP.Host, m.Config.SMTP.Port), auth, m.Config.SMTP.Username, []string{email}, message)
	if err != nil {
		log.Printf("Warning: Smtp sendMail() error: %v", err)
		return err
	}

	return nil
}

func (m *Manager) ConfirmUser(ctx context.Context, userID, code string) error {

	fmt.Println("User ID: ", userID)

	user, err := m.Repository.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	dbCode, err := m.Cache.CodeCache.GetCode(ctx, user.Email)
	if err != nil {
		return err
	}

	if dbCode != code {
		return fmt.Errorf("Incorrect code error")
	}

	err = m.Repository.UpdateUser(ctx, userID, &api.UpdateUserRequest{IsVerified: true})
	if err != nil {
		return err
	}

	user.IsVerified = true

	return m.Cache.UserCache.SetUser(ctx, user)
}

func (m *Manager) GetAllUsers(ctx context.Context, userID string) ([]entity.User, error) {

	user, err := m.Repository.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Role != "admin" {
		return nil, errors.New("Insufficient permissions")
	}

	return m.Repository.GetAllUsers(ctx)
}

func (m *Manager) SetUserRoleById(ctx context.Context, userID, targetUserId, role string) error {

	user, err := m.Repository.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		return errors.New("Insufficient permissions")
	}

	err = m.Repository.SetUserRoleById(ctx, targetUserId, role)
	if err != nil {
		return err
	}

	user.Role = role
	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
