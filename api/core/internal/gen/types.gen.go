// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package gen

// HealthCheck HealthCheckのレスポンス
type HealthCheck struct {
	// Message メッセージ。通常はok
	Message string `json:"message"`
}
