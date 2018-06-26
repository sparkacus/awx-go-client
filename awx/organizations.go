package awx

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Organization struct {
	Id               int                       `json:"id"`
	Type             string                    `json:"type"`
	Url              string                    `json:"url"`
	Related          OrganizationRelated       `json:"related"`
	SummaryFields    OrganizationSummaryFields `json:"summary_fields"`
	Created          string                    `json:"created"`
	Modified         string                    `json:"modified"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description"`
	CustomVirtualenv string                    `json:"custom_virtualenv"`
}

type OrganizationRelated struct {
	NamedUrl                     string `json:"named_url"`
	CreatedBy                    string `json:"created_by"`
	ModifiedBy                   string `json:"modified_by"`
	NotificationTemplatesError   string `json:"notification_templates_error"`
	NotificationTemplatesSuccess string `json:"notification_templates_success"`
	Users                        string `json:"users"`
	NotificationTemplatesAny     string `json:"notification_templates_any"`
	NotificationTemplates        string `json:"notification_templates"`
	Applications                 string `json:"applications"`
	InstanceGroups               string `json:"instance_groups"`
	Credentials                  string `json:"credentials"`
	Inventories                  string `json:"inventories"`
	Projects                     string `json:"projects"`
	WorkflowJobTemplates         string `json:"workflow_job_templates"`
	ObjectRoles                  string `json:"object_roles"`
	AccessList                   string `json:"access_list"`
	Teams                        string `json:"teams"`
	Admins                       string `json:"admins"`
	ActivityStream               string `json:"activity_stream"`
}

type OrganizationSummaryFields struct {
	CreatedBy struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"created_by"`
	ModifiedBy struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"modified_by"`
	ObjectRoles struct {
		AdminRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"admin_role"`
		MemberRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"member_role"`
		ExecuteRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"execute_role"`
		NotificationAdminRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"notification_admin_role"`
		WorkflowAdminRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"workflow_admin_role"`
		CredentialAdminRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"credential_admin_role"`
		ReadRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"read_role"`
		ProjectAdminTole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"project_admin_tole"`
		AuditorRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"auditor_role"`
		InventoryAdminRole struct {
			Id          int    `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"inventory_admin_role"`
	} `json:"object_roles"`
	UserCapabilities struct {
		Edit   bool `json:"edit"`
		Delete bool `json:"delete"`
	} `json:"user_capabilities"`
	RelatedFieldCounts struct {
		JobTemplates int `json:"job_templates"`
		Users        int `json:"users"`
		Teams        int `json:"teams"`
		Admins       int `json:"admins"`
		Inventories  int `json:"inventories"`
		Projects     int `json:"projects"`
	} `json:"related_field_counts"`
}

func (s *Client) GetOrganization() (*Organization, *http.Response, error) {
	url := fmt.Sprint(s.BaseURL + "/organizations/1/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Oops")
	}

	resp, err := s.doRequest(req)
	if err != nil {
		log.Fatal("Oops")
	}

	data := new(Organization)
	json.NewDecoder(resp.Body).Decode(data)
	fmt.Println(resp.Body)
	resp.Body.Close()

	return data, resp, nil
}
