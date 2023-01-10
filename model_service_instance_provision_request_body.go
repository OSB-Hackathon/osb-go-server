/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServiceInstanceProvisionRequestBody struct {
	ServiceId string `json:"service_id"`
	PlanId string `json:"plan_id"`
	Context *Context `json:"context,omitempty"`
	OrganizationGuid string `json:"organization_guid"`
	SpaceGuid string `json:"space_guid"`
	Parameters *Object `json:"parameters,omitempty"`
}
