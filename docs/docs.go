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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles/{id}": {
            "put": {
                "description": "Updates a specific article by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Update articles by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/articles/{id}/comments": {
            "get": {
                "description": "Retrieves comments for a specific article.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Articles"
                ],
                "summary": "Get article comments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Prev",
                        "name": "prev",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Next",
                        "name": "next",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        },
                        "headers": {
                            "Cookie": {
                                "type": "string",
                                "description": "session_id"
                            }
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Register user with username,password, firstname and lastname",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterAuth"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authors/{id}": {
            "get": {
                "description": "Retrieves a specific author by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Get author by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/go-simple-rest_src_v1_authors_model.Author"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates an author",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Update a specific author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an author",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Delete a specific author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authors/{id}/articles": {
            "get": {
                "description": "Retrieves all articles written by a specific author.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Get all articles by author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new article written by a specific author.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Create a new article written by a specific author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-simple-rest_src_v1_authors_model.Article"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Response",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authors/{id}/articles/{articleId}": {
            "put": {
                "description": "Updates an article written by a specific author.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Update an article written by a specific author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Article",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AuthorArticleUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.AuthorArticleUpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an article written by a specific author.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Delete an article written by a specific author",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "articleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response",
                        "schema": {
                            "$ref": "#/definitions/model.AuthorArticleUpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "go-simple-rest_src_v1_auth_model.User": {
            "type": "object",
            "properties": {
                "articleCount": {
                    "type": "integer"
                },
                "articleLikesCount": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "go-simple-rest_src_v1_authors_model.Article": {
            "type": "object",
            "properties": {
                "authorId": {
                    "type": "string"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {},
                "likes": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "go-simple-rest_src_v1_authors_model.Author": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {},
                "lastName": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.AuthorArticleUpdateRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.AuthorArticleUpdateResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.LoginAuth": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "password",
                    "type": "string",
                    "minLength": 4
                },
                "username": {
                    "description": "username",
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/go-simple-rest_src_v1_auth_model.User"
                }
            }
        },
        "model.RegisterAuth": {
            "type": "object",
            "required": [
                "firstname",
                "lastname",
                "password",
                "username"
            ],
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
