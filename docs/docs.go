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
        "/accounts": {
            "post": {
                "description": "Create new account with initial balance. Account ID must be unique.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Create new account",
                "parameters": [
                    {
                        "description": "Account creation data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Account created successfully"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Account already exists",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{account_id}": {
            "get": {
                "description": "Returns a single account by ID. Account ID must be a positive integer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get account by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID (positive integer)",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account details",
                        "schema": {
                            "$ref": "#/definitions/dto.AccountDetail"
                        }
                    },
                    "400": {
                        "description": "Invalid account ID format",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Account doesn't exists",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "Transfers specified amount from source to destination account. Amount must be positive with up to 2 decimal places.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Submit a transaction between accounts",
                "parameters": [
                    {
                        "description": "Transaction details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SubmitTransaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AccountDetail": {
            "type": "object",
            "required": [
                "account_id",
                "balance"
            ],
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "balance": {
                    "type": "string"
                }
            }
        },
        "dto.CreateAccountRequest": {
            "type": "object",
            "required": [
                "account_id",
                "initial_balance"
            ],
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "initial_balance": {
                    "type": "string",
                    "minLength": 0
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {},
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.SubmitTransaction": {
            "type": "object",
            "required": [
                "amount",
                "destination_account_id",
                "source_account_id"
            ],
            "properties": {
                "amount": {
                    "type": "string",
                    "minLength": 0
                },
                "destination_account_id": {
                    "type": "integer"
                },
                "source_account_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
