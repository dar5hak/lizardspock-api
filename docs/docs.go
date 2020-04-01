// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-01 19:35:05.86582178 +0530 IST m=+0.028796810

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
        "contact": {},
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/play/1p": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Play against the system",
                "parameters": [
                    {
                        "type": "string",
                        "description": "your choice: one of rock, paper, scissors, lizard or spock",
                        "name": "choice",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/game.OnePlayerGameResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/game.OnePlayerGameError"
                        }
                    }
                }
            }
        },
        "/play/2p": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Play a two-player game",
                "parameters": [
                    {
                        "type": "string",
                        "description": "player 1's choice: one of rock, paper, scissors, lizard or spock",
                        "name": "p1Choice",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "player 2's choice: one of rock, paper, scissors, lizard or spock",
                        "name": "p2Choice",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/game.TwoPlayerGameResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/game.TwoPlayerGameError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "game.OnePlayerGameError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "userChoice": {
                    "type": "string"
                }
            }
        },
        "game.OnePlayerGameResult": {
            "type": "object",
            "properties": {
                "didYouWin": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "systemChoice": {
                    "type": "string"
                },
                "userChoice": {
                    "type": "string"
                }
            }
        },
        "game.TwoPlayerGameError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "p1Choice": {
                    "type": "string"
                },
                "p2Choice": {
                    "type": "string"
                }
            }
        },
        "game.TwoPlayerGameResult": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "p1Choice": {
                    "type": "string"
                },
                "p2Choice": {
                    "type": "string"
                },
                "winner": {
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
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "lizardspock API",
	Description: "",
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
