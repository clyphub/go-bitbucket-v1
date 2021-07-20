/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucketv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// APIResponse contains generic data from API Response
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
	Values  map[string]interface{}
}

type SelfLink struct {
	Href string `json:"href"`
}

type CloneLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type Links struct {
	Self []SelfLink `json:"self,omitempty"`
}

type Project struct {
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Type        string `json:"type"`
	Links       Links  `json:"links"`
}

// Repository contains data from a BitBucket Repository
type Repository struct {
	Slug          string   `json:"slug,omitempty"`
	ID            int      `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	ScmID         string   `json:"scmId,omitempty"`
	State         string   `json:"state,omitempty"`
	StatusMessage string   `json:"statusMessage,omitempty"`
	Forkable      bool     `json:"forkable,omitempty"`
	Project       *Project `json:"project,omitempty"`
	Public        bool     `json:"public,omitempty"`
	Links         *struct {
		Clone []CloneLink `json:"clone,omitempty"`
		Self  []SelfLink  `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Owner *struct {
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
		ID           int    `json:"id"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		Slug         string `json:"slug"`
		Type         string `json:"type"`
		AvatarURL    string `json:"avatarUrl"`
	} `json:"owner,omitempty"`
	Origin *Repository `json:"origin,omitempty"`
}

type UserWithNameEmail struct {
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
}

type UserWithLinks struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	ID           int    `json:"id,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Type         string `json:"type,omitempty"`
	Links        Links  `json:"links,omitempty"`
}

type User struct {
	Name                        string `json:"name"`
	EmailAddress                string `json:"emailAddress"`
	ID                          int    `json:"id"`
	DisplayName                 string `json:"displayName"`
	Active                      bool   `json:"active"`
	Slug                        string `json:"slug"`
	Type                        string `json:"type"`
	DirectoryName               string `json:"directoryName"`
	Deletable                   bool   `json:"deletable"`
	LastAuthenticationTimestamp int64  `json:"lastAuthenticationTimestamp"`
	MutableDetails              bool   `json:"mutableDetails"`
	MutableGroups               bool   `json:"mutableGroups"`
}

type UserWithMetadata struct {
	User               UserWithLinks `json:"user,omitempty"`
	Role               string        `json:"role,omitempty"`
	Approved           bool          `json:"approved,omitempty"`
	Status             string        `json:"status,omitempty"`
	LastReviewedCommit string        `json:"lastReviewedCommit,omitempty"`
}

// PermissionGlobal are global permissions
type PermissionGlobal string

const (
	// PermissionGlobalLicensedUser represents the ability to log into the system
	PermissionGlobalLicensedUser PermissionGlobal = "LICENSED_USER"
	// PermissionGlobalProjectCreate allows project creation
	PermissionGlobalProjectCreate PermissionGlobal = "PROJECT_CREATE"
	// PermissionGlobalAdmin represents an administrator
	PermissionGlobalAdmin PermissionGlobal = "ADMIN"
	// PermissionGlobalSysAdmin represents a system administrator
	PermissionGlobalSysAdmin PermissionGlobal = "SYS_ADMIN"
)

// PermissionProject are project level permissions
type PermissionProject string

const (
	// PermissionProjectAdmin grants admin priviledges
	PermissionProjectAdmin PermissionProject = "PROJECT_ADMIN"
	// PermissionProjectRead grants read priviledges
	PermissionProjectRead PermissionProject = "PROJECT_READ"
	// PermissionProjectWrite grants write priviledges
	PermissionProjectWrite PermissionProject = "PROJECT_WRITE"
)

// PermissionRepository are repository level permissions
type PermissionRepository string

const (
	// PermissionRepositoryAdmin grants admin priviledges
	PermissionRepositoryAdmin PermissionRepository = "REPO_ADMIN"
	// PermissionRepositoryRead grants read priviledges
	PermissionRepositoryRead PermissionRepository = "REPO_READ"
	// PermissionRepositoryWrite grants write priviledges
	PermissionRepositoryWrite PermissionRepository = "REPO_WRITE"
)

// UserPermission contains a user with its permission
type UserPermission struct {
	User       User   `json:"user"`
	Permission string `json:"permission"`
}

// Group represents a user group
type Group struct {
	Name string `json:"name"`
}

// GroupPermission contains a group with its permission
type GroupPermission struct {
	Group      Group  `json:"group"`
	Permission string `json:"permission"`
}

type MergeResult struct {
	Outcome string `json:"outcome"`
	Current bool   `json:"current"`
}

type MergeGetResponse struct {
	CanMerge   bool                     `json:"canMerge"`
	Conflicted bool                     `json:"conflicted"`
	Outcome    string                   `json:"outcome"`
	Vetoes     []map[string]interface{} `json:"vetoes"`
}

type PullRequestRef struct {
	ID           string     `json:"id"`
	DisplayID    string     `json:"displayId"`
	LatestCommit string     `json:"latestCommit"`
	Repository   Repository `json:"repository"`
}

type PullRequest struct {
	ID           int                `json:"id"`
	Version      int32              `json:"version"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	State        string             `json:"state"`
	Open         bool               `json:"open"`
	Closed       bool               `json:"closed"`
	CreatedDate  int64              `json:"createdDate"`
	UpdatedDate  int64              `json:"updatedDate"`
	FromRef      PullRequestRef     `json:"fromRef"`
	ToRef        PullRequestRef     `json:"toRef"`
	Locked       bool               `json:"locked"`
	Author       *UserWithMetadata  `json:"author,omitempty"`
	Reviewers    []UserWithMetadata `json:"reviewers"`
	Participants []UserWithMetadata `json:"participants,omitempty"`
	Properties   struct {
		MergeResult       MergeResult `json:"mergeResult"`
		ResolvedTaskCount int         `json:"resolvedTaskCount"`
		OpenTaskCount     int         `json:"openTaskCount"`
	} `json:"properties"`
	Links Links `json:"links"`
}

// SSHKey contains data from a SSHKey in the BitBucket Server
type SSHKey struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Label string `json:"label"`
}

// Commit contains data from a commit in BitBucket
type Commit struct {
	ID                 string `json:"id"`
	DisplayID          string `json:"displayId"`
	Author             User   `json:"author"`
	AuthorTimestamp    int64  `json:"authorTimestamp"`
	Committer          User   `json:"committer"`
	CommitterTimestamp int64  `json:"committerTimestamp"`
	Message            string `json:"message"`
	Parents            []struct {
		ID        string `json:"id"`
		DisplayID string `json:"displayId"`
	} `json:"parents"`
}

type BuildStatus struct {
	State       string `json:"state"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	DateAdded   int64  `json:"dateAdded"`
}

type Diff struct {
	Diffs []struct {
		Source struct {
			Components []string `json:"components"`
			Parent     string   `json:"parent"`
			Name       string   `json:"name"`
			Extension  string   `json:"extension"`
			ToString   string   `json:"toString"`
		} `json:"source"`
		Destination struct {
			Components []string `json:"components"`
			Parent     string   `json:"parent"`
			Name       string   `json:"name"`
			Extension  string   `json:"extension"`
			ToString   string   `json:"toString"`
		} `json:"destination"`
		Hunks []struct {
			SourceLine      int `json:"sourceLine"`
			SourceSpan      int `json:"sourceSpan"`
			DestinationLine int `json:"destinationLine"`
			DestinationSpan int `json:"destinationSpan"`
			Segments        []struct {
				Type  string `json:"type"`
				Lines []struct {
					Destination int    `json:"destination"`
					Source      int    `json:"source"`
					Line        string `json:"line"`
					Truncated   bool   `json:"truncated"`
				} `json:"lines"`
				Truncated bool `json:"truncated"`
			} `json:"segments"`
			Truncated bool `json:"truncated"`
		} `json:"hunks"`
		Truncated    bool    `json:"truncated"`
		ContextLines float64 `json:"contextLines"`
		FromHash     string  `json:"fromHash"`
		ToHash       string  `json:"toHash"`
		WhiteSpace   string  `json:"whiteSpace"`
	} `json:"diffs"`
	Truncated    bool    `json:"truncated"`
	ContextLines float64 `json:"contextLines"`
	FromHash     string  `json:"fromHash"`
	ToHash       string  `json:"toHash"`
	WhiteSpace   string  `json:"whiteSpace"`
}

// Tag contaings git Tag information
type Tag struct {
	ID              string `json:"id"`
	DisplayID       string `json:"displayId"`
	Type            string `json:"type"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	Hash            string `json:"hash"`
}

// Branch contains git Branch information
type Branch struct {
	ID              string `json:"id"`
	DisplayID       string `json:"displayId"`
	Type            string `json:"type"`
	LatestCommit    string `json:"latestCommit"`
	LatestChangeset string `json:"latestChangeset"`
	IsDefault       bool   `json:"isDefault"`
}

// Content contains repository content information (and files content)
type Content struct {
	Children struct {
		IsLastPage bool `json:"isLastPage"`
		Limit      int  `json:"limit"`
		Size       int  `json:"size"`
		Start      int  `json:"start"`
		Values     []struct {
			ContentID string `json:"contentId,omitempty"`
			Path      struct {
				Components []string `json:"components"`
				Extension  string   `json:"extension"`
				Name       string   `json:"name"`
				Parent     string   `json:"parent"`
				ToString   string   `json:"toString"`
			} `json:"path"`
			Size int    `json:"size,omitempty"`
			Type string `json:"type"`
			Node string `json:"node,omitempty"`
		} `json:"values"`
	} `json:"children"`
	Path struct {
		Components []string `json:"components"`
		Name       string   `json:"name"`
		ToString   string   `json:"toString"`
	} `json:"path"`
	Revision string `json:"revision"`
}

type WebhookConfiguration struct {
	Secret string `json:"secret"`
}

type Webhook struct {
	ID            int                  `json:"id"`
	Name          string               `json:"name"`
	Events        []string             `json:"events"`
	Url           string               `json:"url"`
	Active        bool                 `json:"active"`
	Configuration WebhookConfiguration `json:"configuration"`
}

type Code struct {
	Category   string `json:"category,omitempty"`
	IsLastPage bool   `json:"isLastPage,omitempty"`
	Count      int    `json:"count,omitempty"`
	Start      int    `json:"start,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	NextStart  int    `json:"nextStart,omitempty"`
	Values     []struct {
		Repository  *Repository `json:"repository,omitempty"`
		File        string      `json:"file"`
		HitContexts [][]struct {
			Line int    `json:"line"`
			Text string `json:"text"`
		} `json:"hitContexts"`
		HitCount int `json:"hitCount"`
	} `json:"values,omitempty"`
}

type Scope struct {
	Type string `json:"type"`
}

type Query struct {
	Substituted bool `json:"substituted"`
}

type SearchResult struct {
	Scope        Scope `json:"scope"`
	Code         *Code `json:"code,omitempty"`
	Repositories *Code `json:"repositories,omitempty"`
	Query        Query `json:"query"`
}

type Limits struct {
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
}

type SearchQuery struct {
	Query    string `json:"query"`
	Entities struct {
		Code Code `json:"code"`
	} `json:"entities"`
	Limits Limits `json:"limits"`
}

type Parent struct {
	ID int `json:"id"`
}

type DiffType string

const (
	DiffTypeEffective DiffType = "EFFECTIVE"
	DiffTypeCommit    DiffType = "COMMIT"
	DiffTypeRange     DiffType = "RANGE"
)

type LineType string

const (
	LineTypeAdded   LineType = "ADDED"
	LineTypeRemoved LineType = "REMOVED"
	LineTypeContext LineType = "CONTEXT"
)

type FileType string

const (
	FileTypeFrom FileType = "FROM"
	FileTypeTo   FileType = "TO"
)

type Anchor struct {
	DiffType DiffType `json:"diffType,omitempty"`
	Line     int      `json:"line,omitempty"`
	LineType LineType `json:"lineType,omitempty"`
	FileType FileType `json:"fileType,omitempty"`
	FromHash string   `json:"fromHash,omitempty"`
	Path     string   `json:"path,omitempty"`
	SrcPath  string   `json:"srcPath,omitempty"`
	ToHash   string   `json:"toHash,omitempty"`
}

type Comment struct {
	Text   string  `json:"text"`
	Parent *Parent `json:"parent,omitempty"`
	Anchor *Anchor `json:"anchor,omitempty"`
}

type Properties struct {
	Key string `json:"key"`
}

type PermittedOperations struct {
	Editable  bool `json:"editable"`
	Deletable bool `json:"deletable"`
}

type ActivityComment struct {
	Properties          Properties          `json:"properties"`
	ID                  int                 `json:"id"`
	Version             int                 `json:"version"`
	Text                string              `json:"text"`
	Author              User                `json:"author"`
	CreatedDate         int64               `json:"createdDate"`
	UpdatedDate         int64               `json:"updatedDate"`
	Comments            []ActivityComment   `json:"comments"`
	PermittedOperations PermittedOperations `json:"permittedOperations"`
}

type CommitsStats struct {
	Commits []Commit `json:"commits"`
	Total   int      `json:"total"`
}

type Action string

const (
	ActionCommented Action = "COMMENTED"
	ActionRescoped  Action = "RESCOPED"
	ActionMerged    Action = "MERGED"
	ActionApproved  Action = "APPROVED"
	ActionDeclined  Action = "DECLINED"
	ActionOpened    Action = "OPENED"
)

type Activity struct {
	ID               int             `json:"id"`
	CreatedDate      int             `json:"createdDate"`
	User             User            `json:"user"`
	Action           Action          `json:"action"`
	CommentAction    string          `json:"commentAction"`
	Comment          ActivityComment `json:"comment"`
	CommentAnchor    Anchor          `json:"commentAnchor"`
	FromHash         string          `json:"fromHash,omitempty"`
	PreviousFromHash string          `json:"previousFromHash,omitempty"`
	PreviousToHash   string          `json:"previousToHash,omitempty"`
	ToHash           string          `json:"toHash,omitempty"`
	Added            CommitsStats    `json:"added"`
	Removed          CommitsStats    `json:"removed"`
}

type Activities struct {
	NextPageStart int        `json:"nextPageStart"`
	IsLastPage    bool       `json:"isLastPage"`
	Limit         int        `json:"limit"`
	Size          int        `json:"size"`
	Start         int        `json:"start"`
	Values        []Activity `json:"values"`
}

// String converts global permission to its string representation
func (p PermissionGlobal) String() string {
	return string(p)
}

// String converts project permission to its string representation
func (p PermissionProject) String() string {
	return string(p)
}

// String converts repository permission to its string representation
func (p PermissionRepository) String() string {
	return string(p)
}

func (k *SSHKey) String() string {
	parts := make([]string, 1, 2)
	parts[0] = strings.TrimSpace(k.Text)
	return strings.Join(parts, " ")
}

// GetCommitsResponse cast Commits into structure
func GetCommitsResponse(r *APIResponse) ([]Commit, error) {
	var m []Commit
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetTagsResponse cast Tags into structure
func GetTagsResponse(r *APIResponse) ([]Tag, error) {
	var m []Tag
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// CreateTagResponse cast Tag into structure
func CreateTagResponse(r *APIResponse) (Tag, error) {
	var m Tag
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetBranchesResponse cast Tags into structure
func GetBranchesResponse(r *APIResponse) ([]Branch, error) {
	var m []Branch
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetRepositoriesResponse cast Repositories into structure
func GetRepositoriesResponse(r *APIResponse) ([]Repository, error) {
	var m []Repository
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetRepositoryResponse cast Repositories into structure
func GetRepositoryResponse(r *APIResponse) (Repository, error) {
	var m Repository
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetDiffResponse cast Diff into structure
func GetDiffResponse(r *APIResponse) (Diff, error) {
	var m Diff
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetSSHKeysResponse cast SSHKeys into structure
func GetSSHKeysResponse(r *APIResponse) ([]SSHKey, error) {
	var m []SSHKey
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetBuildStatusResponse cast BuildStatus into structure
func GetBuildStatusesResponse(r *APIResponse) ([]BuildStatus, error) {
	var m []BuildStatus
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetPullRequestResponse cast PullRequest into structure
func GetPullRequestResponse(r *APIResponse) (PullRequest, error) {
	var m PullRequest
	err := mapstructure.Decode(r.Values, &m)
	return m, err
}

// GetPullRequestResponse PullRequests into structure
func GetPullRequestsResponse(r *APIResponse) ([]PullRequest, error) {
	var m []PullRequest
	err := mapstructure.Decode(r.Values["values"], &m)
	return m, err
}

// GetContentResponse cast Content into structure
func GetContentResponse(r *APIResponse) (Content, error) {
	var c Content
	err := mapstructure.Decode(r.Values, &c)
	return c, err
}

// GetUserWithMetadataResponse casts users into structure
func GetUserWithMetadataResponse(r *APIResponse) (UserWithMetadata, error) {
	var c UserWithMetadata
	err := mapstructure.Decode(r.Values, &c)
	return c, err
}

// GetUsersResponse casts users into structure
func GetUsersResponse(r *APIResponse) ([]User, error) {
	var c []User
	err := mapstructure.Decode(r.Values["values"], &c)
	return c, err
}

// GetUsersPermissionResponse casts user permissions into structure
func GetUsersPermissionResponse(r *APIResponse) ([]UserPermission, error) {
	var c []UserPermission
	err := mapstructure.Decode(r.Values["values"], &c)
	return c, err
}

// GetGroupsPermissionResponse casts group permissions into structure
func GetGroupsPermissionResponse(r *APIResponse) ([]GroupPermission, error) {
	var c []GroupPermission
	err := mapstructure.Decode(r.Values["values"], &c)
	return c, err
}

// GetWebhooksResponse cast Webhooks into structure
func GetWebhooksResponse(r *APIResponse) ([]Webhook, error) {
	var h []Webhook
	err := mapstructure.Decode(r.Values["values"], &h)
	return h, err
}

// GetSearchResultResponse cast Search results into structure
func GetSearchResultResponse(r *APIResponse) (SearchResult, error) {
	var h SearchResult
	err := mapstructure.Decode(r.Values, &h)
	return h, err
}

// GetActivitiesResponse cast Activities results into structure
func GetActivitiesResponse(r *APIResponse) (Activities, error) {
	var h Activities
	err := json.Unmarshal(r.Payload, &h)
	return h, err
}

// NewAPIResponse create new APIResponse from http.Response
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError create new erroneous API response from http.response and error
func NewAPIResponseWithError(r *http.Response, bodyBytes []byte, err error) (*APIResponse, error) {
	response := &APIResponse{Response: r, Message: strings.Replace(err.Error(), "\"", "", -1)}
	if bodyBytes != nil {
		parseErr := json.Unmarshal(bodyBytes, &response.Values)
		if parseErr != nil {
			err = fmt.Errorf("%s ParseError: %v", err.Error(), parseErr)
		}
	}
	return response, err
}

// NewBitbucketAPIResponse create new API response from http.response
func NewBitbucketAPIResponse(r *http.Response) (*APIResponse, error) {
	response := &APIResponse{Response: r}
	err := json.NewDecoder(r.Body).Decode(&response.Values)
	if err != nil {
		return nil, err
	}
	return response, err
}

// NewRawAPIResponse create new API response from http.response with raw data
func NewRawAPIResponse(r *http.Response) (*APIResponse, error) {
	response := &APIResponse{Response: r}
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	response.Payload = raw
	return response, nil
}

// HasNextPage returns if response is paged and has next page and where it does start
func HasNextPage(response *APIResponse) (isLastPage bool, nextPageStart int) {
	isLastPage, ok := response.Values["isLastPage"].(bool)
	if ok && !isLastPage {
		floatStart, ok := response.Values["nextPageStart"].(float64)
		if ok {
			nextPageStart = int(floatStart)
		}
	} else if !ok {
		isLastPage = true
	}
	return !isLastPage, nextPageStart
}
