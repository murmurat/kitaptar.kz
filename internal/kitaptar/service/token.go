package service

import (
	"fmt"
)

func (m *Manager) VerifyToken(token string) (string, error) {

	claim, err := m.Token.ValidateToken(token)
	if err != nil {
		return "", fmt.Errorf("validate token err: %w", err)
	}

	return claim.UserID, nil
}
