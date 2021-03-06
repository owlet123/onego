package services

import (
	"context"

	"github.com/onego-project/onego/resources"
)

// TokenService to service OpenNebula User login token.
type TokenService struct {
	Service
}

const nonExpiringToken = -1
const resetToken = 0

const noGID = -1

// manageToken generates login token for the given user (expiring or non-expiring).
func (ts TokenService) manageToken(ctx context.Context, userName string,
	period int, effectiveGroupID int) (string, error) {
	resArr, err := ts.call(ctx, "one.user.login", userName, "", period, effectiveGroupID)
	if err != nil {
		return "", err
	}

	return resArr[resultIndex].ResultString(), err
}

// GenerateToken generates login token for the given user.
func (ts TokenService) GenerateToken(ctx context.Context, userName string, period int,
	effectiveGroup resources.Group) (string, error) {
	effectiveGroupID, err := effectiveGroup.ID()
	if err != nil {
		return "", err
	}

	return ts.manageToken(ctx, userName, period, effectiveGroupID)
}

// GenerateInfiniteToken generates non-expiring login token for the given user.
func (ts TokenService) GenerateInfiniteToken(ctx context.Context, userName string,
	effectiveGroup resources.Group) (string, error) {
	effectiveGroupID, err := effectiveGroup.ID()
	if err != nil {
		return "", err
	}

	return ts.manageToken(ctx, userName, nonExpiringToken, effectiveGroupID)
}

// GenerateUnscopedToken generates login token for the given user.
func (ts TokenService) GenerateUnscopedToken(ctx context.Context, userName string, period int) (string, error) {
	return ts.manageToken(ctx, userName, period, noGID)
}

// GenerateInfiniteUnscopedToken generates non-expiring login token for the given user.
func (ts TokenService) GenerateInfiniteUnscopedToken(ctx context.Context, userName string) (string, error) {
	return ts.manageToken(ctx, userName, nonExpiringToken, noGID)
}

// RevokeToken resets login token for the given user.
func (ts TokenService) RevokeToken(ctx context.Context, userName, token string) error {
	_, err := ts.call(ctx, "one.user.login", userName, token, resetToken)

	return err
}

// RevokeAllTokens resets all login tokens for the given user.
func (ts TokenService) RevokeAllTokens(ctx context.Context, userName string) error {
	_, err := ts.call(ctx, "one.user.login", userName, "", resetToken)

	return err
}
