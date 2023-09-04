// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Someone here",
            "url": "contact.here",
            "email": "email@here.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles": {
            "get": {
                "description": "Get list article with paging and filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Get list article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "Page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size",
                        "name": "Size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.ListPaging"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Create article",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.Response"
                        }
                    }
                }
            }
        },
        "/articles/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get detail article",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Get detail article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update an existing article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Update article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.Response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login and return token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserWithToken"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create new user, returns user and token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_auth_models.Response"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get list category with paging and filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get list category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "Page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size",
                        "name": "Size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_category_models.ListPaging"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create category",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_category_models.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_category_models.Response"
                        }
                    }
                }
            }
        },
        "/categories/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_category_models.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_category_models.Response"
                        }
                    }
                }
            }
        },
        "/storage/media/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload media file (pdf, docs, images, videos, etc.)",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Storage"
                ],
                "summary": "Upload media",
                "parameters": [
                    {
                        "type": "file",
                        "description": "binary file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_storage_models.Response"
                        }
                    }
                }
            }
        },
        "/todos": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get list todo with paging and filter",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get list todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "Page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size",
                        "name": "Size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.ListPaging"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Create todo",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response"
                        }
                    }
                }
            }
        },
        "/todos/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get list current user todo by token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get list current user todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "Page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size",
                        "name": "Size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get detail todo",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get detail todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Update todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Delete todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response"
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
                "description": "Get user info by token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_auth_models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Ho-Minh_InitiaRe-website_internal_article_models.CreateRequest": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "pre_publish_content": {
                    "type": "string"
                },
                "publish_date": {
                    "type": "string"
                },
                "short_brief": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_article_models.ListPaging": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_article_models.Response"
                    }
                },
                "size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_article_models.Response": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "pre_publish_content": {
                    "type": "string"
                },
                "publish_date": {
                    "type": "string"
                },
                "short_brief": {
                    "type": "string"
                },
                "status_id": {
                    "type": "integer"
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_article_models.UpdateRequest": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "pre_publish_content": {
                    "type": "string"
                },
                "publish_date": {
                    "type": "string"
                },
                "short_brief": {
                    "type": "string"
                },
                "thumbnail": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_auth_models.Response": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "school": {
                    "type": "string"
                },
                "status": {
                    "description": "Custom fields",
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_category_models.CreateRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_category_models.ListPaging": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_category_models.Response"
                    }
                },
                "size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_category_models.Response": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "integer"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_category_models.UpdateRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_storage_models.Response": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "download_url": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_todo_models.CreateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_todo_models.ListPaging": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response"
                    }
                },
                "size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_todo_models.Response": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "update_by": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "github_com_Ho-Minh_InitiaRe-website_internal_todo_models.UpdateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "models.LoginRequest": {
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
        "models.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "school": {
                    "type": "string"
                }
            }
        },
        "models.UserWithToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/github_com_Ho-Minh_InitiaRe-website_internal_auth_models.Response"
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
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "InitiaRe API",
	Description:      "InitiaRe REST API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
