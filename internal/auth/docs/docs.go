// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/login": {
            "post": {
                "description": "Sign in to the app with auth service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login to app",
                "parameters": [
                    {
                        "description": "req body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Refresh access and refresh token by refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "req body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_murat96k_kitaptar_kz_internal_auth_handler_dto.RefreshInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "req body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "api.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "github_com_murat96k_kitaptar_kz_internal_auth_handler_dto.RefreshInput": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "kitaptar.kz Auth service",
	Description:      "Service for authorization for users",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}