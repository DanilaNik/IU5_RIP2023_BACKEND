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
        "/items": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get list of all items",
                "parameters": [
                    {
                        "type": "string",
                        "format": "text",
                        "description": "filter by search text",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingGetItemsResponse"
                        }
                    }
                }
            }
        },
        "/items/image": {
            "post": {
                "consumes": [
                    "multipart/form-data",
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Upload s3 file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "upload file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "metadata",
                        "name": "metadata",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.ImageSwagger"
                        }
                    }
                }
            }
        },
        "/items/post": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Create item",
                "parameters": [
                    {
                        "description": "Item object",
                        "name": "itemPrototype",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpmodels.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.Item"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get item by id",
                "parameters": [
                    {
                        "type": "string",
                        "format": "text",
                        "description": "item id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingGetItemByIDResponse"
                        }
                    }
                }
            }
        },
        "/items/{id}/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Delete item by id",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "text",
                        "description": "item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/items/{id}/post": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Post item to current order",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "text",
                        "description": "item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingGetDraftRequestByIDResponse"
                        }
                    }
                }
            }
        },
        "/items/{id}/put": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Change item",
                "parameters": [
                    {
                        "description": "Item object",
                        "name": "itemPrototype",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpmodels.Item"
                        }
                    },
                    {
                        "type": "integer",
                        "format": "text",
                        "description": "item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "userCreds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingLoginResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "parameters": [
                    {
                        "type": "string",
                        "default": "token=xxx",
                        "description": "token",
                        "name": "Cookie",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get list of all orders",
                "parameters": [
                    {
                        "type": "string",
                        "format": "text",
                        "description": "min date",
                        "name": "min_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "max date",
                        "name": "max_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "order status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingGetRequestsForAdminWithFiltersResponse"
                        }
                    }
                }
            }
        },
        "/orders/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete order",
                "parameters": [
                    {
                        "description": "Order id",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpmodels.RequestID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/orders/items/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete item from current order",
                "parameters": [
                    {
                        "type": "string",
                        "format": "text",
                        "description": "item id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.UserRequest"
                        }
                    }
                }
            }
        },
        "/orders/make": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Confirm current order",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get order by id",
                "parameters": [
                    {
                        "type": "string",
                        "format": "text",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.UserRequest"
                        }
                    }
                }
            }
        },
        "/orders/{id}/approve": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Approve or decline order",
                "parameters": [
                    {
                        "description": "Order status",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpmodels.RequestStatus"
                        }
                    },
                    {
                        "type": "string",
                        "format": "text",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "userPrototype",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/httpmodels.TestingRegisterResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpmodels.ImageSwagger": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                }
            }
        },
        "httpmodels.Item": {
            "type": "object",
            "properties": {
                "barcode": {
                    "type": "integer"
                },
                "depth": {
                    "type": "integer"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "httpmodels.ItemInRequest": {
            "type": "object",
            "properties": {
                "barcode": {
                    "type": "integer"
                },
                "depth": {
                    "type": "integer"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "quantityInRequest": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "httpmodels.Request": {
            "type": "object",
            "properties": {
                "completionDate": {
                    "type": "string"
                },
                "creationDate": {
                    "type": "string"
                },
                "creatorID": {
                    "type": "integer"
                },
                "formationDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "description": "status in ('draft','deleted','formed','completed','rejected')",
                    "type": "string"
                }
            }
        },
        "httpmodels.RequestID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "httpmodels.RequestStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "httpmodels.TestingGetDraftRequestByIDResponse": {
            "type": "object",
            "properties": {
                "request": {
                    "$ref": "#/definitions/httpmodels.Request"
                }
            }
        },
        "httpmodels.TestingGetItemByIDResponse": {
            "type": "object",
            "properties": {
                "item": {
                    "$ref": "#/definitions/httpmodels.Item"
                }
            }
        },
        "httpmodels.TestingGetItemsResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpmodels.Item"
                    }
                },
                "oderdID": {
                    "type": "integer"
                }
            }
        },
        "httpmodels.TestingGetRequestsForAdminWithFiltersResponse": {
            "type": "object",
            "properties": {
                "requests": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpmodels.Request"
                    }
                }
            }
        },
        "httpmodels.TestingLoginRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "httpmodels.TestingLoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "httpmodels.TestingRegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "httpmodels.TestingRegisterResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "httpmodels.UserRequest": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpmodels.ItemInRequest"
                    }
                },
                "request": {
                    "$ref": "#/definitions/httpmodels.Request"
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
