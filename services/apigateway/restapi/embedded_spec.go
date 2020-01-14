// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/github.com.roomflow.backend.v1+json"
  ],
  "produces": [
    "application/github.com.roomflow.backend.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An application to help students locate ideal study space",
    "title": "RoomFlow",
    "version": "1.0.0"
  },
  "paths": {
    "/login": {
      "post": {
        "tags": [
          "login"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/token"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "register"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "registration successful",
            "schema": {
              "$ref": "#/definitions/token"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "credentials": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "minLength": 1
        },
        "username": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "token": {
      "type": "object",
      "required": [
        "accesstoken"
      ],
      "properties": {
        "accesstoken": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/github.com.roomflow.backend.v1+json"
  ],
  "produces": [
    "application/github.com.roomflow.backend.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An application to help students locate ideal study space",
    "title": "RoomFlow",
    "version": "1.0.0"
  },
  "paths": {
    "/login": {
      "post": {
        "tags": [
          "login"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/token"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "register"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "registration successful",
            "schema": {
              "$ref": "#/definitions/token"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "credentials": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "minLength": 1
        },
        "username": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "token": {
      "type": "object",
      "required": [
        "accesstoken"
      ],
      "properties": {
        "accesstoken": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
