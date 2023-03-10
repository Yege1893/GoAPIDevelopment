/*
 * Todo app OAS
 *
 * OpenApi specification for a todo application
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateBody struct {

	Task string `json:"task,omitempty"`

	Status string `json:"status,omitempty"`

	Assignee string `json:"assignee,omitempty"`

	Priority string `json:"priority,omitempty"`
}
