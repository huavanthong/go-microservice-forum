// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/forgotpassword": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User forgot password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Forgot password",
                "parameters": [
                    {
                        "description": "Confirm forget password",
                        "name": "forget",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ForgotPasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User sign in to service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign In User",
                "parameters": [
                    {
                        "description": "Authenticate user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.UserLoginSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete all cookie in session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Log out user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Refresh access token after the specific time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh access token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.UserRefreshTokenSuccess"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Register a new user for service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/payload.UserRegisterSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/auth/resetpassword/{resetToken}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Validate the reset token and update the user’s password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Reset password",
                "parameters": [
                    {
                        "description": "New password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ResetPasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/auth/verifyemail/{verificationCode}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Verify email user that sign up to service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verify email user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Verification Code",
                        "name": "verificationCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "209": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/sessions/oauth/google": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Sign in a new user by Google OAtuth2, then save a new user to DB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in a new user by Google OAuth2",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "307": {
                        "description": "Temporary Redirect",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/payload.Response"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the current user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get the current user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payload.UserRegisterSuccess"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ForgotPasswordInput": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "models.ResetPasswordInput": {
            "type": "object",
            "required": [
                "password",
                "passwordConfirm"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "passwordConfirm": {
                    "type": "string"
                }
            }
        },
        "models.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "models.SignUpInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "passwordConfirm"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password123"
                },
                "passwordConfirm": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "example": "johndoe@gmail.com"
                },
                "id": {
                    "type": "string",
                    "example": "5bbdadf782ebac06a695a8e7"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "photo": {
                    "type": "string",
                    "example": "http://www.golangprograms.com/skin/frontend/base/default/logo.png"
                },
                "provider": {
                    "type": "string",
                    "example": "google oauth2"
                },
                "role": {
                    "type": "string",
                    "example": "user"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "payload.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Error message"
                },
                "status": {
                    "type": "string",
                    "example": "failed"
                }
            }
        },
        "payload.UserLoginSuccess": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0NTg2NzcsImlhdCI6MTY1NzQ1Nzc3NywibmJmIjoxNjU3NDU3Nzc3LCJzdWIiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAifQ.WbRHMAdggCfHR06XKpmbFCu3DNjPkjOPYs9b8TuvBZym1d_TD7JCMadmNCq_Sim9bOzhMi8muDZBb1wRBkli2A"
                },
                "message": {
                    "type": "string",
                    "example": "Generate token success"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "payload.UserRefreshTokenSuccess": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc0NTg2NzcsImlhdCI6MTY1NzQ1Nzc3NywibmJmIjoxNjU3NDU3Nzc3LCJzdWIiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAifQ.WbRHMAdggCfHR06XKpmbFCu3DNjPkjOPYs9b8TuvBZym1d_TD7JCMadmNCq_Sim9bOzhMi8muDZBb1wRBkli2A"
                },
                "message": {
                    "type": "string",
                    "example": "Refresh token success"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "payload.UserRegisterSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 201
                },
                "data": {
                    "$ref": "#/definitions/models.UserResponse"
                },
                "message": {
                    "type": "string",
                    "example": "Register a new user successfully"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v3",
	Schemes:          []string{},
	Title:            "UserManagement Service API Document",
	Description:      "List APIs of UserManagement Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
