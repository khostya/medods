// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/access": {
            "get": {
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "userID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/model.Tokens"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/resp.RespError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/resp.RespError"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": " ",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Refresh"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "refreshed",
                        "schema": {
                            "$ref": "#/definitions/model.Tokens"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/resp.RespError"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/resp.RespError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.Refresh": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "model.Tokens": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                },
                "refresh": {
                    "type": "string"
                }
            }
        },
        "resp.RespError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "validationError": {
                    "$ref": "#/definitions/validator.ValidationError"
                }
            }
        },
        "validator.ValidationError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "param": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "medods",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}