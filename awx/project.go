package awx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const projectBasePath = "api/v2/projects/"

// ProjectService is an interface for interfacing with the Project
// endpoints of the AWX API
// See: http://localhost/api/v2/projects/
type ProjectService interface {
	List(context.Context) ([]Project, *Response, error)
	Get(context.Context, int) (*Project, *Response, error)
	Create(context.Context, *ProjectCreateRequest) (*Project, *Response, error)
	Delete(context.Context, int) (*Response, error)
}

// ProjectsServiceOp handles communication with the Project related methods of the
// AWX API.
type ProjectServiceOp struct {
	client *Client
}

// Project represents a AWX Project
type Project struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		NamedURL                     string `json:"named_url"`
		CreatedBy                    string `json:"created_by"`
		LastJob                      string `json:"last_job"`
		NotificationTemplatesError   string `json:"notification_templates_error"`
		NotificationTemplatesSuccess string `json:"notification_templates_success"`
		ObjectRoles                  string `json:"object_roles"`
		NotificationTemplatesAny     string `json:"notification_templates_any"`
		Copy                         string `json:"copy"`
		ProjectUpdates               string `json:"project_updates"`
		Update                       string `json:"update"`
		AccessList                   string `json:"access_list"`
		Teams                        string `json:"teams"`
		ScmInventorySources          string `json:"scm_inventory_sources"`
		InventoryFiles               string `json:"inventory_files"`
		Schedules                    string `json:"schedules"`
		Playbooks                    string `json:"playbooks"`
		ActivityStream               string `json:"activity_stream"`
		Organization                 string `json:"organization"`
		LastUpdate                   string `json:"last_update"`
	} `json:"related"`
	SummaryFields struct {
		LastJob struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			Finished    time.Time `json:"finished"`
			Status      string    `json:"status"`
			Failed      bool      `json:"failed"`
		} `json:"last_job"`
		LastUpdate struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			Failed      bool   `json:"failed"`
		} `json:"last_update"`
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
		ObjectRoles struct {
			AdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"admin_role"`
			UseRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"use_role"`
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
			Edit     bool `json:"edit"`
			Start    bool `json:"start"`
			Copy     bool `json:"copy"`
			Schedule bool `json:"schedule"`
			Delete   bool `json:"delete"`
		} `json:"user_capabilities"`
	} `json:"summary_fields"`
	Created               time.Time `json:"created"`
	Modified              time.Time `json:"modified"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	LocalPath             string    `json:"local_path"`
	ScmType               string    `json:"scm_type"`
	ScmURL                string    `json:"scm_url"`
	ScmBranch             string    `json:"scm_branch"`
	ScmClean              bool      `json:"scm_clean"`
	ScmDeleteOnUpdate     bool      `json:"scm_delete_on_update"`
	Credential            int       `json:"credential"`
	Timeout               int       `json:"timeout"`
	LastJobRun            time.Time `json:"last_job_run"`
	LastJobFailed         bool      `json:"last_job_failed"`
	NextJobRun            time.Time `json:"next_job_run"`
	Status                string    `json:"status"`
	Organization          int       `json:"organization"`
	ScmDeleteOnNextUpdate bool      `json:"scm_delete_on_next_update"`
	ScmUpdateOnLaunch     bool      `json:"scm_update_on_launch"`
	ScmUpdateCacheTimeout int       `json:"scm_update_cache_timeout"`
	ScmRevision           string    `json:"scm_revision"`
	CustomVirtualenv      string    `json:"custom_virtualenv"`
	LastUpdateFailed      bool      `json:"last_update_failed"`
	LastUpdated           time.Time `json:"last_updated"`
}

// ProjectCreateRequest represents a request to create a Project.
type ProjectCreateRequest struct {
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	LocalPath             string `json:"local_path,omitempty"`
	ScmType               string `json:"scm_type,omitempty"`
	ScmURL                string `json:"scm_url,omitempty"`
	ScmBranch             string `json:"scm_branch,omitempty"`
	ScmClean              bool   `json:"scm_clean,omitempty"`
	ScmDeleteOnUpdate     bool   `json:"scm_delete_on_update,omitempty"`
	Credential            int    `json:"credential,omitempty"`
	Timeout               int    `json:"timeout,omitempty"`
	Organization          int    `json:"organization,omitempty"`
	ScmUpdateOnLaunch     bool   `json:"scm_update_on_launch,omitempty"`
	ScmUpdateCacheTimeout int    `json:"scm_update_cache_timeout,omitempty"`
	CustomVirtualenv      string `json:"custom_virtualenv,omitempty"`
}

// projectyRoot represents a Project root
type projectRoot struct {
	Count    int       `json:"count"`
	Next     int       `json:"next"`
	Previous int       `json:"previous"`
	Results  []Project `json:"results"`
	Project  *Project
}

// List all Projects.
func (s *ProjectServiceOp) List(ctx context.Context) ([]Project, *Response, error) {
	path := projectBasePath
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(projectRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Results, resp, err
}

// Get individual Project.
func (s *ProjectServiceOp) Get(ctx context.Context, projectID int) (*Project, *Response, error) {
	if projectID < 1 {
		return nil, nil, NewArgError("projectID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", projectBasePath, projectID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Project)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create Project
func (s *ProjectServiceOp) Create(ctx context.Context, createRequest *ProjectCreateRequest) (*Project, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := projectBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}
	root := new(projectRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	return root.Project, resp, err
}

// Delete Project.
func (s *ProjectServiceOp) Delete(ctx context.Context, projectID int) (*Response, error) {
	if projectID < 1 {
		return nil, NewArgError("projectID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", projectBasePath, projectID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
