// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-18 16:03:25.645574 +0800 CST m=+0.076351220

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "This is a simple trying RESTful-API",
        "title": "blog-api",
        "termsOfService": "https://github.com/Mictrlan/blog-api",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "https://github.com/Mictrlan/blog-api/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/add/article": {
            "post": {
                "description": "Add a new article",
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new article to the store",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "false",
                        "description": "Desc",
                        "name": "description",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Content",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/add/auth": {
            "post": {
                "description": "Add a new user for verification",
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/add/tag": {
            "post": {
                "description": "Add a new tag",
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new tag to the store",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "CreatedBy",
                        "name": "created_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/delete/article": {
            "delete": {
                "description": "delete article by id",
                "produces": [
                    "application/json"
                ],
                "summary": "softdelete article by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/delete/tag": {
            "delete": {
                "description": "delete tag by id",
                "produces": [
                    "application/json"
                ],
                "summary": "softdelete tag by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get/article": {
            "post": {
                "description": "article tag by id",
                "produces": [
                    "application/json"
                ],
                "summary": "query an article information by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get/articles": {
            "post": {
                "description": "get articles by tag_id",
                "produces": [
                    "application/json"
                ],
                "summary": "query articles information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get/tag": {
            "post": {
                "description": "tag tag by id",
                "produces": [
                    "application/json"
                ],
                "summary": "query an tag information by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/get/tags": {
            "get": {
                "description": "get tags information",
                "produces": [
                    "application/json"
                ],
                "summary": "query tags information",
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "User login",
                "produces": [
                    "application/json"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/modifyPwd": {
            "put": {
                "description": "Modify user password",
                "produces": [
                    "application/json"
                ],
                "summary": "Modify user password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PwdNew",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Confirm",
                        "name": "confirm",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/remove/article": {
            "delete": {
                "description": "remove article by id",
                "produces": [
                    "application/json"
                ],
                "summary": "harddelete article by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/remove/tag": {
            "delete": {
                "description": "remove tag by id",
                "produces": [
                    "application/json"
                ],
                "summary": "harddelete tag by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/update/article": {
            "put": {
                "description": "update article by id",
                "produces": [
                    "application/json"
                ],
                "summary": "update article by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "TagID",
                        "name": "tag_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Title",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "false",
                        "description": "Desc",
                        "name": "description",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Content",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UpdateBy",
                        "name": "updated_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/update/tag": {
            "put": {
                "description": "update tag by id",
                "produces": [
                    "application/json"
                ],
                "summary": "update tag by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "UpdateBy",
                        "name": "updated_by",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/upload": {
            "post": {
                "description": "upload files",
                "produces": [
                    "application/json"
                ],
                "summary": "Sort and upload files",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ArticleID",
                        "name": "articles_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"stauts\":200,\"message\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
