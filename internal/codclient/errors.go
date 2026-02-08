package codclient

import "errors"

var (
	ErrPrivateProfile = errors.New("player profile is set to private")
	ErrPlayerNotFound = errors.New("player not found")
	ErrRateLimited    = errors.New("rate limited by CoD API")
	ErrAPIUnavailable = errors.New("CoD API is unavailable")
	ErrTokenExpired   = errors.New("SSO token has expired")
)
