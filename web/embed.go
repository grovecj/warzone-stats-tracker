//go:build !embed_dist

package web

import "embed"

// DistFS is empty when the frontend hasn't been built.
// The main.go code gracefully handles this (no static files served).
var DistFS embed.FS
