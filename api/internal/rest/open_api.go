package rest

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/go-chi/chi/v5"
)

func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "ToDo API",
			Description: "REST APIs used for interacting with the ToDo Service",
			Version:     "0.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/MarioCarrion/todo-api-microservice-example",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://127.0.0.1:9234",
			},
		},
	}

	swagger.Components = &openapi3.Components{}

	swagger.Components.Schemas = openapi3.Schemas{
		"Priority": openapi3.NewSchemaRef("",
			openapi3.NewStringSchema().
				WithEnum("none", "low", "medium", "high").
				WithDefault("none")),
		"Dates": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("start", openapi3.NewStringSchema().
					WithFormat("date-time").
					WithNullable()).
				WithProperty("due", openapi3.NewStringSchema().
					WithFormat("date-time").
					WithNullable())),
		"User": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewIntegerSchema()).
				WithProperty("username", openapi3.NewStringSchema()).
				WithProperty("email", openapi3.NewStringSchema()).
				WithPropertyRef("priority", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Priority",
				}).
				WithPropertyRef("dates", &openapi3.SchemaRef{
					Ref: "#/components/schemas/Dates",
				})),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("description", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithPropertyRef("priority", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Priority",
					}).
					WithPropertyRef("dates", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Dates",
					})),
		},
		"UpdateUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("description", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("is_done", openapi3.NewBoolSchema().
						WithDefault(false)).
					WithPropertyRef("priority", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Priority",
					}).
					WithPropertyRef("dates", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Dates",
					})),
		},
		"SearchUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for searching a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("description", openapi3.NewStringSchema().
						WithMinLength(1).
						WithNullable()).
					WithProperty("is_done", openapi3.NewBoolSchema().
						WithDefault(false).
						WithNullable()).
					WithPropertyRef("priority", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Priority",
					}).WithNullable().
					WithProperty("from", openapi3.NewInt64Schema().
						WithDefault(0)).
					WithProperty("size", openapi3.NewInt64Schema().
						WithDefault(10))),
		},
	}

	swagger.Components.Responses = openapi3.ResponseBodies{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"CreateUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("user", &openapi3.SchemaRef{
						Ref: "#/components/schemas/User",
					}))),
		},
		"ReadUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one user.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("user", &openapi3.SchemaRef{
						Ref: "#/components/schemas/User",
					}))),
		},
		"SearchUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching for any user.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("users", &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "array",
							Items: &openapi3.SchemaRef{
								Ref: "#/components/schemas/User",
							},
						},
					}).
					WithProperty("total", openapi3.NewInt64Schema()))),
		},
	}

	swagger.Paths = &openapi3.Paths{
		Extensions: map[string]interface{}{
			"/login": &openapi3.PathItem{
				Post: &openapi3.Operation{
					OperationID: "Login",
					RequestBody: &openapi3.RequestBodyRef{
						Ref: "#/api/payloads/LoginRequest",
					},
					Responses: &openapi3.Responses{
						Extensions: map[string]interface{}{
							"400": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
							"500": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
							"201": &openapi3.ResponseRef{
								Ref: "#/components/responses/CreateUsersResponse",
							},
						},
					},
				},
			},
			"/users/{userId}": &openapi3.PathItem{
				Delete: &openapi3.Operation{
					OperationID: "DeleteUser",
					Parameters: []*openapi3.ParameterRef{
						{
							Value: openapi3.NewPathParameter("userId").
								WithSchema(openapi3.NewIntegerSchema()),
						},
					},
					Responses: &openapi3.Responses{
						Extensions: map[string]interface{}{
							"200": &openapi3.ResponseRef{
								Value: openapi3.NewResponse().WithDescription("User updated"),
							},
							"404": &openapi3.ResponseRef{
								Value: openapi3.NewResponse().WithDescription("User not found"),
							},
							"500": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
						},
					},
				},
				Get: &openapi3.Operation{
					OperationID: "ReadUser",
					Parameters: []*openapi3.ParameterRef{
						{
							Value: openapi3.NewPathParameter("userId").
								WithSchema(openapi3.NewIntegerSchema()),
						},
					},
					Responses: &openapi3.Responses{
						Extensions: map[string]interface{}{
							"200": &openapi3.ResponseRef{
								Ref: "#/components/responses/ReadUsersResponse",
							},
							"404": &openapi3.ResponseRef{
								Value: openapi3.NewResponse().WithDescription("User not found"),
							},
							"500": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
						},
					},
				},
				Put: &openapi3.Operation{
					OperationID: "UpdateUser",
					Parameters: []*openapi3.ParameterRef{
						{
							Value: openapi3.NewPathParameter("userId").
								WithSchema(openapi3.NewIntegerSchema()),
						},
					},
					RequestBody: &openapi3.RequestBodyRef{
						Ref: "#/components/requestBodies/UpdateUsersRequest",
					},
					Responses: &openapi3.Responses{
						Extensions: map[string]interface{}{
							"200": &openapi3.ResponseRef{
								Value: openapi3.NewResponse().WithDescription("User updated"),
							},
							"400": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
							"404": &openapi3.ResponseRef{
								Value: openapi3.NewResponse().WithDescription("User not found"),
							},
							"500": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
						},
					},
				},
			},
			"/search/users": &openapi3.PathItem{
				Post: &openapi3.Operation{
					OperationID: "SearchUser",
					RequestBody: &openapi3.RequestBodyRef{
						Ref: "#/components/requestBodies/SearchUsersRequest",
					},
					Responses: &openapi3.Responses{
						Extensions: map[string]interface{}{
							"200": &openapi3.ResponseRef{
								Ref: "#/components/responses/SearchUsersResponse",
							},
							"400": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
							"500": &openapi3.ResponseRef{
								Ref: "#/components/responses/ErrorResponse",
							},
						},
					},
				},
			},
		},
	}
	return swagger
}

func RegisterOpenAPI(router chi.Router) {
	swagger := NewOpenAPI3()

	router.Get("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, r, &swagger, http.StatusOK)
	})

	router.Get("/openapi3.yaml", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	})
}
