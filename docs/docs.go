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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hello": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "init test function",
                "responses": {}
            }
        },
        "/test": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "used to test function",
                "responses": {}
            }
        },
        "/test/upload": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "used to test the file upload function",
                "parameters": [
                    {
                        "type": "file",
                        "description": "the avatar image file selected by the user",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/user/avatarUpdate": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to upload image and replace user's avatar as the uploaded image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "the avatar image file selected by the user",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/user/detail": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to get the target user's detailed information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/follow": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to follow other people",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/followerList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to get the user's follower list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "used to authorize user and return jwt token, id",
                "parameters": [
                    {
                        "description": "the passed-in parameter of login function",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLoginReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/modify": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to modify the user's personal information",
                "parameters": [
                    {
                        "description": "the passed-in parameter of modify function",
                        "name": "UserModifyReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserModifyReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/passwordChange": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to modify the user's password",
                "parameters": [
                    {
                        "description": "old password and new password",
                        "name": "UserPasswordChangeReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserPasswordChangeReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/register": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "used to register new account",
                "parameters": [
                    {
                        "description": "the passed-in parameter of register function",
                        "name": "UserRegisterReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegisterReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/subscribedList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "used to get the user's subscribed list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/tagList": {
            "get": {
                "tags": [
                    "User"
                ],
                "summary": "used to get all the tag",
                "responses": {}
            }
        }
    },
    "definitions": {
        "request.UserLoginReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UserModifyReq": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "request.UserPasswordChangeReq": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "request.UserRegisterReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
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
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Intelligent Terminal Backend API",
	Description:      "This is the API document of Intelligent Terminal Backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
