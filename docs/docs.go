// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
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
        "/guild": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a list of all guilds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guilds"
                ],
                "summary": "Get all guilds",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new guild",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guilds"
                ],
                "summary": "Create a new guild",
                "parameters": [
                    {
                        "description": "Guild object",
                        "name": "guild",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Guild"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/guild/kick": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Remove a player from a guild by player ID and guild ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guilds"
                ],
                "summary": "Remove a player from a guild",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "player_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Guild ID",
                        "name": "guild_id",
                        "in": "query",
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
        "/guild/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a guild by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guilds"
                ],
                "summary": "Get a guild by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Guild ID",
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
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update an existing guild",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guilds"
                ],
                "summary": "Update a guild",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Guild ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Guild object",
                        "name": "guild",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Guild"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a guild by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "guilds"
                ],
                "summary": "Delete a guild",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Guild ID",
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
                "description": "Authenticate user credentials",
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
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/player": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a list of all players",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get all players",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Create a new player",
                "parameters": [
                    {
                        "description": "Player object",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Player"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/player/guild/join/": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Allows a player to join a guild",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Join a player to a guild",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "player_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Guild ID",
                        "name": "guild_id",
                        "in": "query",
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
        "/player/quest/join": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Associate a player with a quest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Join a player to a quest",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "player_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quest ID",
                        "name": "quest_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Player associated with the quest successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/player/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a player by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get a player by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
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
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update an existing player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Update a player",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Player object",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UPlayer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a player by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Delete a player",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Player deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quests": {
            "get": {
                "description": "Get all quests",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quests"
                ],
                "summary": "List all quests",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Create a new quest",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quests"
                ],
                "summary": "Create a new quest",
                "parameters": [
                    {
                        "description": "Quest object",
                        "name": "quest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Quest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/quests/{id}": {
            "get": {
                "description": "Get a quest by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quests"
                ],
                "summary": "Get quest by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Quest ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "description": "Update an existing quest by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quests"
                ],
                "summary": "Update an existing quest",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Quest ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Quest object",
                        "name": "quest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Quest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "description": "Delete an existing quest by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quests"
                ],
                "summary": "Delete an existing quest",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Quest ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Credentials": {
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
        "controllers.Guild": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.Player": {
            "type": "object",
            "properties": {
                "class": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.Quest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "difficulty": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "reward": {
                    "type": "string"
                }
            }
        },
        "controllers.UPlayer": {
            "type": "object",
            "properties": {
                "class": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "guildID": {
                    "type": "integer"
                },
                "level": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}