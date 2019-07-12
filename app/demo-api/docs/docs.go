// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-12 18:33:59.712982 +0800 CST m=+6.565439124

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/api/v1/student/add": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加学生",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学生姓名",
                        "name": "studName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "年龄",
                        "name": "studAge",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "性别",
                        "name": "studSex",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/render.JSON"
                        }
                    }
                }
            }
        },
        "/api/v1/student/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "学生列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "学生姓名",
                        "name": "studName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/render.JSON"
                        }
                    }
                }
            }
        },
        "/api/v1/student/update": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改学生",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "学生编号",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "学生姓名",
                        "name": "StudName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/render.JSON"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "render.JSON": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                },
                "ttl": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
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
