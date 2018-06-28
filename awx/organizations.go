package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Organization struct {
	Id               int                       `json:"id,omitempty"`
	Type             string                    `json:"type,omitempty"`
	Url              string                    `json:"url,omitempty"`
	Related          OrganizationRelated       `json:"related,omitempty"`
	SummaryFields    OrganizationSummaryFields `json:"summary_fields,omitempty"`
	Created          string                    `json:"created,omitempty"`
	Modified         string                    `json:"modified,omitempty"`
	Name             string                    `json:"name"`
	Description      string                    `json:"description,omitempty"`
	CustomVirtualenv string                    `json:"custom_virtualenv,omitempty"`
}

type OrganizationRelated struct {
	NamedUrl                     string `json:"named_url,omitempty"`
	CreatedBy                    string `json:"created_by,omitempty"`
	ModifiedBy                   string `json:"modified_by,omitempty"`
	NotificationTemplatesError   string `json:"notification_templates_error,omitempty"`
	NotificationTemplatesSuccess string `json:"notification_templates_success,omitempty"`
	Users                        string `json:"users,omitempty"`
	NotificationTemplatesAny     string `json:"notification_templates_any,omitempty"`
	NotificationTemplates        string `json:"notification_templates,omitempty"`
	Applications                 string `json:"applications,omitempty"`
	InstanceGroups               string `json:"instance_groups,omitempty"`
	Credentials                  string `json:"credentials,omitempty"`
	Inventories                  string `json:"inventories,omitempty"`
	Projects                     string `json:"projects,omitempty"`
	WorkflowJobTemplates         string `json:"workflow_job_templates,omitempty"`
	ObjectRoles                  string `json:"object_roles,omitempty"`
	AccessList                   string `json:"access_list,omitempty"`
	Teams                        string `json:"teams,omitempty"`
	Admins                       string `json:"admins,omitempty"`
	ActivityStream               string `json:"activity_stream,omitempty"`
}

type OrganizationSummaryFields struct {
	CreatedBy struct {
		Id        int    `json:"id,omitempty"`
		Username  string `json:"username,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	} `json:"created_by"`
	ModifiedBy struct {
		Id        int    `json:"id,omitempty"`
		Username  string `json:"username,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	} `json:"modified_by"`
	ObjectRoles struct {
		AdminRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"admin_role"`
		MemberRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"member_role"`
		ExecuteRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"execute_role"`
		NotificationAdminRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"notification_admin_role"`
		WorkflowAdminRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"workflow_admin_role"`
		CredentialAdminRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"credential_admin_role"`
		ReadRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"read_role"`
		ProjectAdminTole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"project_admin_tole"`
		AuditorRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"auditor_role"`
		InventoryAdminRole struct {
			Id          int    `json:"id,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
		} `json:"inventory_admin_role"`
	} `json:"object_roles"`
	UserCapabilities struct {
		Edit   bool `json:"edit,omitempty"`
		Delete bool `json:"delete,omitempty"`
	} `json:"user_capabilities"`
	RelatedFieldCounts struct {
		JobTemplates int `json:"job_templates,omitempty"`
		Users        int `json:"users,omitempty"`
		Teams        int `json:"teams,omitempty"`
		Admins       int `json:"admins,omitempty"`
		Inventories  int `json:"inventories,omitempty"`
		Projects     int `json:"projects,omitempty"`
	} `json:"related_field_counts"`
}

func (s *Client) CreateRequestOrganization() *Organization {
	data := new(Organization)
	return data
}

func (s *Client) CreateOrganization(i *Organization) (*Organization, *http.Response, error) {
	url := fmt.Sprint(s.BaseURL + "/organizations/")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)

	req, err := http.NewRequest("POST", url, b)

	if err != nil {
		return nil, nil, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return nil, nil, err
	}

	data := new(Organization)
	json.NewDecoder(resp.Body).Decode(data)
	resp.Body.Close()

	return data, resp, nil
}

func (s *Client) ReadOrganization(id int) (*Organization, *http.Response, error) {
	url := fmt.Sprintf(s.BaseURL+"/organizations/%d/", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return nil, nil, err
	}

	data := new(Organization)
	json.NewDecoder(resp.Body).Decode(data)
	resp.Body.Close()

	return data, resp, nil
}

func (s *Client) UpdateOrganization(id int, i *Organization) (*Organization, *http.Response, error) {
	url := fmt.Sprintf(s.BaseURL+"/organizations/%d/", id)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)

	req, err := http.NewRequest("PUT", url, b)

	if err != nil {
		log.Fatal("Oops")
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return nil, nil, err
	}

	data := new(Organization)
	json.NewDecoder(resp.Body).Decode(data)
	resp.Body.Close()

	return data, resp, nil
}

func (s *Client) DeleteOrganization(id int) (*http.Response, error) {
	url := fmt.Sprintf(s.BaseURL+"/organizations/%d/", id)

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return resp, nil
}
