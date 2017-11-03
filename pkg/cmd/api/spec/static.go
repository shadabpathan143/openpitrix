// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Code generated by "makestatic"; DO NOT EDIT.

package static

var Files = map[string]string{
	"app.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "app.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/apps": {
      "get": {
        "operationId": "GetAppList",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixAppListResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "page_size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "page_number",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "AppService"
        ]
      },
      "post": {
        "operationId": "CreateApp",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixApp"
            }
          }
        ],
        "tags": [
          "AppService"
        ]
      }
    },
    "/v1/apps/{id}": {
      "get": {
        "operationId": "GetApp",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixApp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "AppService"
        ]
      },
      "delete": {
        "operationId": "DeleteApp",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "AppService"
        ]
      },
      "post": {
        "operationId": "UpdateApp",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixApp"
            }
          }
        ],
        "tags": [
          "AppService"
        ]
      }
    }
  },
  "definitions": {
    "openpitrixApp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "repo_id": {
          "type": "string"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "last_modified": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "openpitrixAppListResponse": {
      "type": "object",
      "properties": {
        "total_items": {
          "type": "integer",
          "format": "int32"
        },
        "total_pages": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "current_page": {
          "type": "integer",
          "format": "int32"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openpitrixApp"
          }
        }
      }
    },
    "protobufEmpty": {
      "type": "object",
      "description": "service Foo {\n      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);\n    }\n\nThe JSON representation for ` + "`" + `Empty` + "`" + ` is empty JSON object ` + "`" + `{}` + "`" + `.",
      "title": "A generic empty message that you can re-use to avoid defining duplicated\nempty messages in your APIs. A typical example is to use it as the request\nor the response type of an API method. For instance:"
    }
  }
}
`,

	"app_runtime.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "app_runtime.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/appruntimes": {
      "get": {
        "summary": "Returns a list containing all app runtimes.",
        "operationId": "GetAppRuntimeList",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixAppRuntimeListResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "page_size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "page_number",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "AppRuntimeService"
        ]
      },
      "post": {
        "operationId": "CreateAppRuntime",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixAppRuntime"
            }
          }
        ],
        "tags": [
          "AppRuntimeService"
        ]
      }
    },
    "/v1/appruntimes/{id}": {
      "get": {
        "operationId": "GetAppRuntime",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixAppRuntime"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "AppRuntimeService"
        ]
      },
      "delete": {
        "operationId": "DeleteAppRuntime",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "AppRuntimeService"
        ]
      },
      "post": {
        "operationId": "UpdateAppRuntime",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixAppRuntime"
            }
          }
        ],
        "tags": [
          "AppRuntimeService"
        ]
      }
    }
  },
  "definitions": {
    "openpitrixAppRuntime": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "last_modified": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "openpitrixAppRuntimeListResponse": {
      "type": "object",
      "properties": {
        "total_items": {
          "type": "integer",
          "format": "int32"
        },
        "total_pages": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "current_page": {
          "type": "integer",
          "format": "int32"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openpitrixAppRuntime"
          }
        }
      }
    },
    "protobufEmpty": {
      "type": "object",
      "description": "service Foo {\n      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);\n    }\n\nThe JSON representation for ` + "`" + `Empty` + "`" + ` is empty JSON object ` + "`" + `{}` + "`" + `.",
      "title": "A generic empty message that you can re-use to avoid defining duplicated\nempty messages in your APIs. A typical example is to use it as the request\nor the response type of an API method. For instance:"
    }
  }
}
`,

	"cluster.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "cluster.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/clusters": {
      "get": {
        "operationId": "GetClusterList",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixClusterListResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "page_size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "page_number",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "ClusterService"
        ]
      },
      "post": {
        "operationId": "CreateCluster",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixCluster"
            }
          }
        ],
        "tags": [
          "ClusterService"
        ]
      }
    },
    "/v1/clusters/{id}": {
      "get": {
        "operationId": "GetCluster",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixCluster"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ClusterService"
        ]
      },
      "delete": {
        "operationId": "DeleteCluster",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ClusterService"
        ]
      },
      "post": {
        "operationId": "UpdateCluster",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixCluster"
            }
          }
        ],
        "tags": [
          "ClusterService"
        ]
      }
    }
  },
  "definitions": {
    "openpitrixCluster": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "app_id": {
          "type": "string"
        },
        "app_version": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "transition_status": {
          "type": "string"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "last_modified": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "openpitrixClusterListResponse": {
      "type": "object",
      "properties": {
        "total_items": {
          "type": "integer",
          "format": "int32"
        },
        "total_pages": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "current_page": {
          "type": "integer",
          "format": "int32"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openpitrixCluster"
          }
        }
      }
    },
    "protobufEmpty": {
      "type": "object",
      "description": "service Foo {\n      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);\n    }\n\nThe JSON representation for ` + "`" + `Empty` + "`" + ` is empty JSON object ` + "`" + `{}` + "`" + `.",
      "title": "A generic empty message that you can re-use to avoid defining duplicated\nempty messages in your APIs. A typical example is to use it as the request\nor the response type of an API method. For instance:"
    }
  }
}
`,

	"repo.swagger.json": `{
  "swagger": "2.0",
  "info": {
    "title": "repo.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/repos": {
      "get": {
        "operationId": "GetRepoList",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixRepoListResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "page_size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          },
          {
            "name": "page_number",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int32"
          }
        ],
        "tags": [
          "RepoService"
        ]
      },
      "post": {
        "operationId": "CreateRepo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixRepo"
            }
          }
        ],
        "tags": [
          "RepoService"
        ]
      }
    },
    "/v1/repos/{id}": {
      "get": {
        "operationId": "GetRepo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/openpitrixRepo"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "RepoService"
        ]
      },
      "delete": {
        "operationId": "DeleteRepo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "RepoService"
        ]
      },
      "post": {
        "operationId": "UpdateRepo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protobufEmpty"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/openpitrixRepo"
            }
          }
        ],
        "tags": [
          "RepoService"
        ]
      }
    }
  },
  "definitions": {
    "openpitrixRepo": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "last_modified": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "openpitrixRepoListResponse": {
      "type": "object",
      "properties": {
        "total_items": {
          "type": "integer",
          "format": "int32"
        },
        "total_pages": {
          "type": "integer",
          "format": "int32"
        },
        "page_size": {
          "type": "integer",
          "format": "int32"
        },
        "current_page": {
          "type": "integer",
          "format": "int32"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/openpitrixRepo"
          }
        }
      }
    },
    "protobufEmpty": {
      "type": "object",
      "description": "service Foo {\n      rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);\n    }\n\nThe JSON representation for ` + "`" + `Empty` + "`" + ` is empty JSON object ` + "`" + `{}` + "`" + `.",
      "title": "A generic empty message that you can re-use to avoid defining duplicated\nempty messages in your APIs. A typical example is to use it as the request\nor the response type of an API method. For instance:"
    }
  }
}
`,
}
