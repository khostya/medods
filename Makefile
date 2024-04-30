swag:
	swag init -q -g ./httpserver/handlers/router.go -dir ./internal,./pkg --output ./docs
.PHONY: swag