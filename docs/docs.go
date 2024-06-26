// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Indomobil",
            "url": "https://github.com/IMSIDevOps",
            "email": "dev.ops@indomobil.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/IMSIDevOps/-/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/binning/getAll": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Binning List By Header",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Binning"
                ],
                "summary": "Show An Binning List",
                "parameters": [
                    {
                        "description": "Insert Header Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/request.BinningHeaderRequest"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.BinningHeaderResponses"
                            }
                        }
                    }
                }
            }
        },
        "/auth/loginAuth": {
            "post": {
                "description": "Login With User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login With User",
                "parameters": [
                    {
                        "description": "Insert Header Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payloads.LoginRequestPayloads"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payloads.Respons"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "REST API User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "description": "Insert Register Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payloads.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payloads.Respons"
                        }
                    }
                }
            }
        },
        "/user/username/{username}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "REST API User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Controller"
                ],
                "summary": "Find User By ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payloads.Respons"
                        }
                    }
                }
            }
        },
        "/user/{user_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "REST API User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Controller"
                ],
                "summary": "Find User By ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/payloads.Respons"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "payloads.LoginRequestPayloads": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "payloads.RegisterRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_email": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "user_role_id": {
                    "type": "integer"
                }
            }
        },
        "payloads.Respons": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "request.BinningHeaderRequest": {
            "type": "object",
            "properties": {
                "COMPANY_CODE": {
                    "type": "string"
                },
                "PO_DOC_NO": {
                    "type": "string"
                }
            }
        },
        "response.BinningDetailResponses": {
            "type": "object",
            "properties": {
                "binDocNo": {
                    "type": "string"
                },
                "binLineNo": {
                    "type": "string"
                },
                "caseNo": {
                    "type": "string"
                },
                "grpoQty": {
                    "type": "integer"
                },
                "itemCode": {
                    "type": "string"
                },
                "locCode": {
                    "type": "string"
                },
                "poLineNo": {
                    "type": "string"
                }
            }
        },
        "response.BinningHeaderResponses": {
            "type": "object",
            "properties": {
                "companyCode": {
                    "type": "string"
                },
                "item": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.BinningDetailResponses"
                    }
                },
                "poDocNo": {
                    "type": "string"
                },
                "whscode": {
                    "type": "string"
                },
                "whsgroup": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "DMS User Service",
	Description:      "DMS User Service Architecture",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
