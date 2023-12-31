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
        "/api/voucher/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "CreateVoucher.",
                "parameters": [
                    {
                        "description": "Voucher Code",
                        "name": "voucher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VoucherRequestModel"
                        }
                    }
                ],
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
        "/api/voucher/redeem": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "RedeemVoucher.",
                "parameters": [
                    {
                        "description": "Voucher Code",
                        "name": "voucher",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RedeemVoucherRequest"
                        }
                    }
                ],
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
        "/api/voucher/{voucherCode}/used": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetVoucherCodeUsed"
                ],
                "summary": "GetVoucherCodeUsed",
                "parameters": [
                    {
                        "type": "string",
                        "description": "voucherCode",
                        "name": "voucherCode",
                        "in": "path",
                        "required": true
                    }
                ],
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
        }
    },
    "definitions": {
        "models.RedeemVoucherRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.VoucherRequestModel": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "code": {
                    "type": "string"
                },
                "usable": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
