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
        "contact": {
            "name": "mirpo",
            "url": "https://github.com/mirpo"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/receipts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "List uploaded receipts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListFileResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.FileResponse": {
            "type": "object",
            "properties": {
                "file": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "widths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "handler.ListFileResponse": {
            "type": "object",
            "properties": {
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.FileResponse"
                    }
                },
                "widths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "uploader",
	Description:      "upload receipt REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
