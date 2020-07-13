// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/nilorg",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/nilorg/naas",
            "email": "862860000@qq.com"
        },
        "license": {
            "name": "GNU General Public License v3.0",
            "url": "https://github.com/nilorg/naas/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/login": {
            "post": {
                "description": "后台管理员登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员（已弃用）"
                ],
                "summary": "管理员登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        },
        "/oauth2/clients": {
            "get": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "查询客户端翻页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth2"
                ],
                "summary": "client",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "当前页",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/model.TableListData"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "$ref": "#/definitions/service.ResultClientInfo"
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "创建客户端",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth2"
                ],
                "summary": "client",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.OAuth2ClientEditModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        },
        "/oauth2/clients/{client_id}": {
            "put": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "根据客户端ID,修改客户端信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth2"
                ],
                "summary": "client",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client id",
                        "name": "client_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "客户端信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.OAuth2ClientEditModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        },
        "/oauth2/clients/{client_id}/scopes": {
            "get": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "查询客户端scope",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth2"
                ],
                "summary": "scope",
                "parameters": [
                    {
                        "type": "string",
                        "description": "客户端ID",
                        "name": "client_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.OAuth2ClientScope"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/oauth2/scopes": {
            "get": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "查询所有scope",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OAuth2"
                ],
                "summary": "scope",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.OAuth2Scope"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/resources/{resource_id}/web_routes": {
            "post": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Resource（资源）"
                ],
                "summary": "添加web路由",
                "parameters": [
                    {
                        "type": "string",
                        "description": "资源ID",
                        "name": "resource_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ResourceAddWebRouteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        },
        "/roles": {
            "get": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "recursive:递归获取所有角色\nlist:查询列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role（角色）"
                ],
                "summary": "查询角色",
                "parameters": [
                    {
                        "enum": [
                            "recursive",
                            "list"
                        ],
                        "type": "string",
                        "description": "查询参数",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Role"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/roles/{role_code}/resource_web_route/{resource_web_route_id}": {
            "post": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role（角色）"
                ],
                "summary": "添加资源web路由",
                "parameters": [
                    {
                        "type": "string",
                        "description": "角色Code",
                        "name": "role_code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "资源web路由ID",
                        "name": "resource_web_route_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ResourceAddWebRouteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "查询用户翻页数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User（用户）"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "当前页",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.TableListData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User（用户）"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.userCreateModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "根据用户ID,获取一个用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User（用户）"
                ],
                "summary": "获取一个用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "根据用户ID,修改一个用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User（用户）"
                ],
                "summary": "修改一个用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户需要修改的信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserUpdateModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "OAuth2AccessCode": []
                    }
                ],
                "description": "根据用户ID,删除一个用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User（用户）"
                ],
                "summary": "删除一个用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Result": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "object",
                    "$ref": "#/definitions/api.ResultError"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "api.ResultError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "api.userCreateModel": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "model.OAuth2ClientScope": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "oauth2_client_id": {
                    "type": "integer"
                },
                "scope_code": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.OAuth2Scope": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "description": "basic,",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Pagination": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.Role": {
            "type": "object",
            "properties": {
                "child_roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Role"
                    }
                },
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent_code": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.TableListData": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "object"
                },
                "pagination": {
                    "type": "object",
                    "$ref": "#/definitions/model.Pagination"
                }
            }
        },
        "service.OAuth2ClientEditModel": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "redirect_uri": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "service.ResourceAddWebRouteRequest": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "service.ResultClientInfo": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "redirect_uri": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "service.UserUpdateModel": {
            "type": "object",
            "properties": {
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
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "http://localhost:8080/oauth2/authorize",
            "tokenUrl": "http://localhost:8080/oauth2/token",
            "scopes": {
                "email": " 用户emial",
                "openid": " 用户openid",
                "phone": " 用户手机号",
                "profile": " 用户资料"
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
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "NilOrg认证授权服务",
	Description: "NilOrg认证授权服务Api详情.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
