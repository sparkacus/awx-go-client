package awx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const inventoryBasePath = "api/v2/inventories/"

// InventoryService is an interface for interfacing with the Inventory
// endpoints of the AWX API
// See: http://localhost/api/v2/inventories/
type InventoryService interface {
	List(context.Context) ([]Inventory, *Response, error)
	Get(context.Context, int) (*Inventory, *Response, error)
	Create(context.Context, *InventoryCreateRequest) (*Inventory, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// DropletsServiceOp handles communication with the Inventory related methods of the
// AWX API.
type InventoryServiceOp struct {
	client *Client
}

// Inventory represents a AWX Inventory
type Inventory struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		NamedURL               string `json:"named_url"`
		CreatedBy              string `json:"created_by"`
		ModifiedBy             string `json:"modified_by"`
		JobTemplates           string `json:"job_templates"`
		VariableData           string `json:"variable_data"`
		RootGroups             string `json:"root_groups"`
		ObjectRoles            string `json:"object_roles"`
		AdHocCommands          string `json:"ad_hoc_commands"`
		Script                 string `json:"script"`
		Tree                   string `json:"tree"`
		AccessList             string `json:"access_list"`
		ActivityStream         string `json:"activity_stream"`
		InstanceGroups         string `json:"instance_groups"`
		Hosts                  string `json:"hosts"`
		Groups                 string `json:"groups"`
		Copy                   string `json:"copy"`
		UpdateInventorySources string `json:"update_inventory_sources"`
		InventorySources       string `json:"inventory_sources"`
		InsightsCredential     string `json:"insights_credential"`
		Organization           string `json:"organization"`
	} `json:"related"`
	SummaryFields struct {
		InsightsCredential struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"insights_credential"`
		Organization struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"organization"`
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
			UseRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"use_role"`
			AdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"admin_role"`
			AdhocRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"adhoc_role"`
			UpdateRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"update_role"`
			ReadRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"read_role"`
		} `json:"object_roles"`
		UserCapabilities struct {
			Edit   bool `json:"edit"`
			Copy   bool `json:"copy"`
			Adhoc  bool `json:"adhoc"`
			Delete bool `json:"delete"`
		} `json:"user_capabilities"`
	} `json:"summary_fields"`
	Created                      time.Time   `json:"created"`
	Modified                     time.Time   `json:"modified"`
	Name                         string      `json:"name"`
	Description                  string      `json:"description"`
	Organization                 int         `json:"organization"`
	Kind                         string      `json:"kind"`
	HostFilter                   interface{} `json:"host_filter"`
	Variables                    string      `json:"variables"`
	HasActiveFailures            bool        `json:"has_active_failures"`
	TotalHosts                   int         `json:"total_hosts"`
	HostsWithActiveFailures      int         `json:"hosts_with_active_failures"`
	TotalGroups                  int         `json:"total_groups"`
	GroupsWithActiveFailures     int         `json:"groups_with_active_failures"`
	HasInventorySources          bool        `json:"has_inventory_sources"`
	TotalInventorySources        int         `json:"total_inventory_sources"`
	InventorySourcesWithFailures int         `json:"inventory_sources_with_failures"`
	InsightsCredential           int         `json:"insights_credential"`
	PendingDeletion              bool        `json:"pending_deletion"`
}

// InventoryCreateRequest represents a request to create a Inventory.
type InventoryCreateRequest struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Organization       int    `json:"organization"`
	Kind               string `json:"kind,omitempty"`
	HostFilter         string `json:"host_filter,omitempty"`
	Variables          string `json:"variables,omitempty"`
	InsightsCredential int    `json:"insights_credential,omitempty"`
}

// InventoryRoot represents a Inventory root
type inventoryRoot struct {
	Count     int         `json:"count"`
	Next      int         `json:"next"`
	Previous  int         `json:"previous"`
	Results   []Inventory `json:"results"`
	Inventory *Inventory
}

// List all Inventories.
func (s *InventoryServiceOp) List(ctx context.Context) ([]Inventory, *Response, error) {
	path := inventoryBasePath
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(inventoryRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Results, resp, err
}

// Get individual Inventory.
func (s *InventoryServiceOp) Get(ctx context.Context, inventoryID int) (*Inventory, *Response, error) {
	if inventoryID < 1 {
		return nil, nil, NewArgError("inventoryID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", inventoryBasePath, inventoryID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Inventory)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create Inventory
func (s *InventoryServiceOp) Create(ctx context.Context, createRequest *InventoryCreateRequest) (*Inventory, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := inventoryBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(inventoryRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Inventory, resp, err
}

// Delete Inventory.
func (s *InventoryServiceOp) Delete(ctx context.Context, inventoryID int) (*Response, error) {
	if inventoryID < 1 {
		return nil, NewArgError("inventoryID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", inventoryBasePath, inventoryID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
