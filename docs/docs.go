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
        "/api/otp": {
            "post": {
                "description": "Create an OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Create OTP",
                "parameters": [
                    {
                        "description": "OTP request body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateOtpDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/otp/key": {
            "post": {
                "description": "Generate a new OTP key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Create OTP Key",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/otp/{user-reference}": {
            "put": {
                "description": "Validate an OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Validate OTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User reference",
                        "name": "user_reference",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "OTP request body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ValidateOtpDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/otp/{user_reference}/{device_reference}": {
            "get": {
                "description": "Retrieve the last OTP for a user and device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Get the last OTP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User reference",
                        "name": "user_reference",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Device reference",
                        "name": "device_reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateOtpDto": {
            "type": "object",
            "required": [
                "channel",
                "contact",
                "device",
                "otp_type",
                "user_reference"
            ],
            "properties": {
                "channel": {
                    "type": "string"
                },
                "contact": {
                    "type": "string"
                },
                "device": {
                    "$ref": "#/definitions/dto.Device"
                },
                "otp_type": {
                    "type": "string"
                },
                "user_reference": {
                    "type": "string"
                }
            }
        },
        "dto.Device": {
            "type": "object",
            "required": [
                "brand",
                "device_reference",
                "imei",
                "model",
                "type"
            ],
            "properties": {
                "brand": {
                    "type": "string"
                },
                "device_reference": {
                    "type": "string"
                },
                "imei": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 10
                },
                "model": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.ValidateOtpDto": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "device": {
                    "$ref": "#/definitions/dto.Device"
                },
                "otp": {
                    "type": "string"
                },
                "otp_type": {
                    "type": "string"
                },
                "user_reference": {
                    "type": "string"
                }
            }
        },
        "helper.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error_reference": {
                    "type": "string"
                },
                "error_type": {
                    "type": "string"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "timestamp": {
                    "type": "string"
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