package auth

import "log/slog"

func (a *Auth) Token(userID string) (string, error) {
	token, err := a.GenerateToken(userID)
	if err != nil {
		a.logger.Error("failed to generate token",
			slog.String("error", err.Error()),
		)
		return "", err //TODO
	}

	return token, nil
}
