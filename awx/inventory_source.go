package awx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const inventorySourceBasePath = "api/v2/inventory_sources/"

// InventorySourceService is an interface for interfacing with the InventorySource
// endpoints of the AWX API
// See: http://localhost/api/v2/inventories/
type InventorySourceService interface {
	List(context.Context) ([]InventorySource, *Response, error)
	Get(context.Context, int) (*InventorySource, *Response, error)
	Create(context.Context, *InventorySourceCreateRequest) (*InventorySource, *Response, error)
	Update(context.Context, *InventorySourceCreateRequest, int) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

// InventorySourceServiceOp handles communication with the InventorySource related methods of the
// AWX API.
type InventorySourceServiceOp struct {
	client *Client
}

// InventorySource represents a AWX InventorySource
type InventorySource struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		NamedURL                     string `json:"named_url"`
		CreatedBy                    string `json:"created_by"`
		ModifiedBy                   string `json:"modified_by"`
		NotificationTemplatesError   string `json:"notification_templates_error"`
		NotificationTemplatesSuccess string `json:"notification_templates_success"`
		NotificationTemplatesAny     string `json:"notification_templates_any"`
		InventoryUpdates             string `json:"inventory_updates"`
		Update                       string `json:"update"`
		Hosts                        string `json:"hosts"`
		Groups                       string `json:"groups"`
		Schedules                    string `json:"schedules"`
		Credentials                  string `json:"credentials"`
		ActivityStream               string `json:"activity_stream"`
		Inventory                    string `json:"inventory"`
	} `json:"related"`
	SummaryFields struct {
		Inventory struct {
			ID                           int    `json:"id"`
			Name                         string `json:"name"`
			Description                  string `json:"description"`
			HasActiveFailures            bool   `json:"has_active_failures"`
			TotalHosts                   int    `json:"total_hosts"`
			HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
			TotalGroups                  int    `json:"total_groups"`
			GroupsWithActiveFailures     int    `json:"groups_with_active_failures"`
			HasInventorySources          bool   `json:"has_inventory_sources"`
			TotalInventorySources        int    `json:"total_inventory_sources"`
			InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
			OrganizationID               int    `json:"organization_id"`
			Kind                         string `json:"kind"`
		} `json:"inventory"`
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
		UserCapabilities struct {
			Edit     bool `json:"edit"`
			Start    bool `json:"start"`
			Schedule bool `json:"schedule"`
			Delete   bool `json:"delete"`
		} `json:"user_capabilities"`
	} `json:"summary_fields"`
	Created               time.Time `json:"created"`
	Modified              time.Time `json:"modified"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	Source                string    `json:"source"`
	SourcePath            string    `json:"source_path"`
	SourceScript          string    `json:"source_script"`
	SourceVars            string    `json:"source_vars"`
	Credential            int       `json:"credential"`
	SourceRegions         string    `json:"source_regions"`
	InstanceFilters       string    `json:"instance_filters"`
	GroupBy               string    `json:"group_by"`
	Overwrite             bool      `json:"overwrite"`
	OverwriteVars         bool      `json:"overwrite_vars"`
	Timeout               int       `json:"timeout"`
	Verbosity             int       `json:"verbosity"`
	LastJobRun            time.Time `json:"last_job_run"`
	LastJobFailed         bool      `json:"last_job_failed"`
	NextJobRun            time.Time `json:"next_job_run"`
	Status                string    `json:"status"`
	Inventory             int       `json:"inventory"`
	UpdateOnLaunch        bool      `json:"update_on_launch"`
	UpdateCacheTimeout    int       `json:"update_cache_timeout"`
	SourceProject         int       `json:"source_project"`
	UpdateOnProjectUpdate bool      `json:"update_on_project_update"`
	LastUpdateFailed      bool      `json:"last_update_failed"`
	LastUpdated           time.Time `json:"last_updated"`
}

// InventorySourceCreateRequest represents a request to create a InventorySource.
type InventorySourceCreateRequest struct {
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	Source                string `json:"source,omitempty"`
	SourcePath            string `json:"source_path,omitempty"`
	SourceScript          string `json:"source_script,omitempty"`
	SourceVars            string `json:"source_vars,omitempty"`
	Credential            int    `json:"credential,omitempty"`
	SourceRegions         string `json:"source_regions,omitempty"`
	InstanceFilters       string `json:"instance_filters,omitempty"`
	GroupBy               string `json:"group_by,omitempty"`
	Overwrite             bool   `json:"overwrite,omitempty"`
	OverwriteVars         bool   `json:"overwrite_vars,omitempty"`
	Timeout               int    `json:"timeout,omitempty"`
	Verbosity             int    `json:"verbosity,omitempty"`
	Inventory             int    `json:"inventory,omitempty"`
	UpdateOnLaunch        bool   `json:"update_on_launch,omitempty"`
	UpdateCacheTimeout    int    `json:"update_cache_timeout,omitempty"`
	SourceProject         int    `json:"source_project,omitempty"`
	UpdateOnProjectUpdate bool   `json:"update_on_project_update,omitempty"`
}

// InventorySourceRoot represents a InventorySource root
type inventorySourceRoot struct {
	Count     int               `json:"count"`
	Next      int               `json:"next"`
	Previous  int               `json:"previous"`
	Results   []InventorySource `json:"results"`
	Inventory *InventorySource
}

// List all Inventories.
func (s *InventorySourceServiceOp) List(ctx context.Context) ([]InventorySource, *Response, error) {
	path := inventorySourceBasePath
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(inventorySourceRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Results, resp, err
}

// Get individual InventorySource.
func (s *InventorySourceServiceOp) Get(ctx context.Context, inventorySourceID int) (*InventorySource, *Response, error) {
	if inventorySourceID < 1 {
		return nil, nil, NewArgError("inventorySourceID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", inventorySourceBasePath, inventorySourceID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(InventorySource)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create InventorySource
func (s *InventorySourceServiceOp) Create(ctx context.Context, createRequest *InventorySourceCreateRequest) (*InventorySource, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := inventorySourceBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(InventorySource)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Update InventorySource.
func (s *InventorySourceServiceOp) Update(ctx context.Context, createRequest *InventorySourceCreateRequest, inventorySourceID int) (*Response, error) {
	if inventorySourceID < 1 {
		return nil, NewArgError("inventorySourceID", "cannot be less than 1")
	}
	if createRequest == nil {
		return nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", inventorySourceBasePath, inventorySourceID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, path, createRequest)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Delete InventorySource.
func (s *InventorySourceServiceOp) Delete(ctx context.Context, inventorySourceID int) (*Response, error) {
	if inventorySourceID < 1 {
		return nil, NewArgError("inventoryID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", inventorySourceBasePath, inventorySourceID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
