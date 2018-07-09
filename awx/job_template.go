package awx

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const jobTemplateBasePath = "api/v2/job_templates/"

// JobTemplateService is an interface for interfacing with the JobTemplate
// endpoints of the AWX API
// See: http://localhost/api/v2/job_templates/
type JobTemplateService interface {
	List(context.Context) ([]JobTemplate, *Response, error)
	Get(context.Context, int) (*JobTemplate, *Response, error)
	Create(context.Context, *JobTemplateCreateRequest) (*JobTemplate, *Response, error)
	Update(context.Context, *JobTemplateCreateRequest, int) (*Response, error)
	Delete(context.Context, int) (*Response, error)
}

// JobTemplateServiceOp handles communication with the JobTemplate related methods of the
// AWX API.
type JobTemplateServiceOp struct {
	client *Client
}

// JobTemplate represents a AWX JobTemplate
type JobTemplate struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Related struct {
		NamedURL                     string `json:"named_url"`
		CreatedBy                    string `json:"created_by"`
		ModifiedBy                   string `json:"modified_by"`
		Labels                       string `json:"labels"`
		Inventory                    string `json:"inventory"`
		ExtraCredentials             string `json:"extra_credentials"`
		Credentials                  string `json:"credentials"`
		NotificationTemplatesError   string `json:"notification_templates_error"`
		NotificationTemplatesSuccess string `json:"notification_templates_success"`
		Jobs                         string `json:"jobs"`
		ObjectRoles                  string `json:"object_roles"`
		NotificationTemplatesAny     string `json:"notification_templates_any"`
		AccessList                   string `json:"access_list"`
		Launch                       string `json:"launch"`
		InstanceGroups               string `json:"instance_groups"`
		Schedules                    string `json:"schedules"`
		Copy                         string `json:"copy"`
		ActivityStream               string `json:"activity_stream"`
		SurveySpec                   string `json:"survey_spec"`
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
			InsightsCredentialID         int    `json:"insights_credential_id"`
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
		ObjectRoles struct {
			AdminRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"admin_role"`
			ExecuteRole struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				Name        string `json:"name"`
			} `json:"execute_role"`
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
		Labels struct {
			Count   int      `json:"count"`
			Results []string `json:"results"`
		} `json:"labels"`
		RecentJobs       []int `json:"recent_jobs"`
		ExtraCredentials []int `json:"extra_credentials"`
		Credentials      []int `json:"credentials"`
	} `json:"summary_fields"`
	Created               time.Time `json:"created"`
	Modified              time.Time `json:"modified"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	JobType               string    `json:"job_type"`
	Inventory             int       `json:"inventory"`
	Project               int       `json:"project"`
	Playbook              string    `json:"playbook"`
	Forks                 int       `json:"forks"`
	Limit                 string    `json:"limit"`
	Verbosity             int       `json:"verbosity"`
	ExtraVars             string    `json:"extra_vars"`
	JobTags               string    `json:"job_tags"`
	ForceHandlers         bool      `json:"force_handlers"`
	SkipTags              string    `json:"skip_tags"`
	StartAtTask           string    `json:"start_at_task"`
	Timeout               int       `json:"timeout"`
	UseFactCache          bool      `json:"use_fact_cache"`
	LastJobRun            time.Time `json:"last_job_run"`
	LastJobFailed         bool      `json:"last_job_failed"`
	NextJobRun            time.Time `json:"next_job_run"`
	Status                string    `json:"status"`
	HostConfigKey         string    `json:"host_config_key"`
	AskDiffModeOnLaunch   bool      `json:"ask_diff_mode_on_launch"`
	AskVariablesOnLaunch  bool      `json:"ask_variables_on_launch"`
	AskLimitOnLaunch      bool      `json:"ask_limit_on_launch"`
	AskTagsOnLaunch       bool      `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch   bool      `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch    bool      `json:"ask_job_type_on_launch"`
	AskVerbosityOnLaunch  bool      `json:"ask_verbosity_on_launch"`
	AskInventoryOnLaunch  bool      `json:"ask_inventory_on_launch"`
	AskCredentialOnLaunch bool      `json:"ask_credential_on_launch"`
	SurveyEnabled         bool      `json:"survey_enabled"`
	BecomeEnabled         bool      `json:"become_enabled"`
	DiffMode              bool      `json:"diff_mode"`
	AllowSimultaneous     bool      `json:"allow_simultaneous"`
	CustomVirtualenv      string    `json:"custom_virtualenv"`
	Credential            int       `json:"credential"`
	VaultCredential       int       `json:"vault_credential"`
}

// JobTemplateCreateRequest represents a request to create a JobTemplate.
type JobTemplateCreateRequest struct {
	Name                  string `json:"name"`
	Description           string `json:"description,omitempty"`
	JobType               string `json:"job_type,omitempty"`
	Inventory             int    `json:"inventory,omitempty"`
	Project               int    `json:"project,omitempty"`
	Playbook              string `json:"playbook,omitempty"`
	Forks                 int    `json:"forks,omitempty"`
	Limit                 string `json:"limit,omitempty"`
	Verbosity             int    `json:"verbosity,omitempty"`
	ExtraVars             string `json:"extra_vars,omitempty"`
	JobTags               string `json:"job_tags,omitempty"`
	ForceHandlers         bool   `json:"force_handlers,omitempty"`
	SkipTags              string `json:"skip_tags,omitempty"`
	StartAtTask           string `json:"start_at_task,omitempty"`
	Timeout               int    `json:"timeout,omitempty"`
	UseFactCache          bool   `json:"use_fact_cache,omitempty"`
	HostConfigKey         string `json:"host_config_key,omitempty"`
	AskDiffModeOnLaunch   bool   `json:"ask_diff_mode_on_launch,omitempty"`
	AskVariablesOnLaunch  bool   `json:"ask_variables_on_launch,omitempty"`
	AskLimitOnLaunch      bool   `json:"ask_limit_on_launch,omitempty"`
	AskTagsOnLaunch       bool   `json:"ask_tags_on_launch,omitempty"`
	AskSkipTagsOnLaunch   bool   `json:"ask_skip_tags_on_launch,omitempty"`
	AskJobTypeOnLaunch    bool   `json:"ask_job_type_on_launch,omitempty"`
	AskVerbosityOnLaunch  bool   `json:"ask_verbosity_on_launch,omitempty"`
	AskInventoryOnLaunch  bool   `json:"ask_inventory_on_launch,omitempty"`
	AskCredentialOnLaunch bool   `json:"ask_credential_on_launch,omitempty"`
	SurveyEnabled         bool   `json:"survey_enabled,omitempty"`
	BecomeEnabled         bool   `json:"become_enabled,omitempty"`
	DiffMode              bool   `json:"diff_mode,omitempty"`
	AllowSimultaneous     bool   `json:"allow_simultaneous,omitempty"`
	CustomVirtualenv      string `json:"custom_virtualenv,omitempty"`
	Credential            int    `json:"credential,omitempty"`
	VaultCredential       int    `json:"vault_credential,omitempty"`
}

// jobTemplateRoot represents a JobTemplate root
type jobTemplateRoot struct {
	Count       int           `json:"count"`
	Next        int           `json:"next"`
	Previous    int           `json:"previous"`
	Results     []JobTemplate `json:"results"`
	JobTemplate *JobTemplate
}

// List all JobTemplate.
func (s *JobTemplateServiceOp) List(ctx context.Context) ([]JobTemplate, *Response, error) {
	path := jobTemplateBasePath
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(jobTemplateRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Results, resp, err
}

// Get individual JobTemplate.
func (s *JobTemplateServiceOp) Get(ctx context.Context, jobTemplateID int) (*JobTemplate, *Response, error) {
	if jobTemplateID < 1 {
		return nil, nil, NewArgError("jobTemplateID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", jobTemplateBasePath, jobTemplateID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(JobTemplate)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// Create JobTemplate
func (s *JobTemplateServiceOp) Create(ctx context.Context, createRequest *JobTemplateCreateRequest) (*JobTemplate, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := jobTemplateBasePath

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}
	root := new(JobTemplate)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	return root, resp, err
}

// Update JobTemplate.
func (s *JobTemplateServiceOp) Update(ctx context.Context, createRequest *JobTemplateCreateRequest, jobTemplateID int) (*Response, error) {
	if jobTemplateID < 1 {
		return nil, NewArgError("jobTemplateID", "cannot be less than 1")
	}
	if createRequest == nil {
		return nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s%d/", jobTemplateBasePath, jobTemplateID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, path, createRequest)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}

// Delete JobTemplate.
func (s *JobTemplateServiceOp) Delete(ctx context.Context, jobTemplateID int) (*Response, error) {
	if jobTemplateID < 1 {
		return nil, NewArgError("jobTemplateID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s%d/", jobTemplateBasePath, jobTemplateID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)

	return resp, err
}
