package awx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const organizationBasePath = "api/v2/organizations/"

// OrganizationService is an interface for interfacing with the Organization
// endpoints of the AWX API
// See: http://localhost/api/v2/organizations/
type OrganizationService interface {
	List(context.Context) ([]Organization, *Response, error)
	Get(context.Context, int) (*Organization, *Response, error)
	Create(context.Context, *OrganizationCreateRequest) (*Organization, *Response, error)
	Update(context.Context, *OrganizationCreateRequest, int) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

// OrganizationServiceOp handles communication with the Organization related methods of the
// AWX API.
type OrganizationServiceOp struct {
	client *Client
}

// Organization represents a AWX Organization
type Organization struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		NamedURL                     string `json:"named_url"`
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
	} `json:"related"`
	SummaryFields struct {
		CreatedBy struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"created_by"`
		ModifiedBy struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"modified_by"`
		ObjectRoles struct {
			AdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"admin_role"`
			MemberRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"member_role"`
			ExecuteRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"execute_role"`
			NotificationAdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"notification_admin_role"`
			WorkflowAdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"workflow_admin_role"`
			CredentialAdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"credential_admin_role"`
			ReadRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"read_role"`
			ProjectAdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"project_admin_role"`
			AuditorRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"auditor_role"`
			InventoryAdminRole struct {
				ID          int    `json:"id"`
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
	} `json:"summary_fields"`
	Created          time.Time `json:"created"`
	Modified         time.Time `json:"modified"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CustomVirtualenv string    `json:"custom_virtualenv"`
}

// OrganizationCreateRequest represents a request to create a Organization.
type OrganizationCreateRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	CustomVirtualenv string `json:"custom_virtualenv,omitempty"`
}

// OrganizationRoot represents a Organization root
type organizationRoot struct {
	Count        int            `json:"count"`
	Next         int            `json:"next"`
	Previous     int            `json:"previous"`
	Results      []Organization `json:"results"`
	Organization *Organization
}

// List all Organizations.
func (s *OrganizationServiceOp) List(ctx context.Context) ([]Organization, *Response, error) {
	path := organizationBasePath
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(organizationRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Results, resp, err
}

// Get individual Organization.
func (s *OrganizationServiceOp) Get(ctx context.Context, organizationID int) (*Organization, *Response, error) {
	if organizationID < 1 {
		return nil, nil, NewArgError("organizationID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", organizationBasePath, organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organization)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create Organization
func (s *OrganizationServiceOp) Create(ctx context.Context, createRequest *OrganizationCreateRequest) (*Organization, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := organizationBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(Organization)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update Organization.
func (s *OrganizationServiceOp) Update(ctx context.Context, createRequest *OrganizationCreateRequest, organizationID int) (*Response, error) {
	if organizationID < 1 {
		return nil, NewArgError("organizationID", "cannot be less than 1")
	}
	if createRequest == nil {
		return nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", organizationBasePath, organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, path, createRequest)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Delete Organization.
func (s *OrganizationServiceOp) Delete(ctx context.Context, organizationID int) (*Response, error) {
	if organizationID < 1 {
		return nil, NewArgError("OrganizationID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", organizationBasePath, organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
