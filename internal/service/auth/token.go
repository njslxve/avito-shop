package auth

import "log/slog"

func (a *Auth) Token(userID string) (string, error) {
	const op = "auth.Token"

	token, err := a.GenerateToken(userID)
	if err != nil {
		a.logger.Error("failed to generate token",
			slog.String("operation", op),
			slog.String("error", err.Error()),
		)
		return "", err
	}

	return token, nil
}
