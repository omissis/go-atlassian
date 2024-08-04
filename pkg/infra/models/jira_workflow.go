package models

// WorkflowSearchOptions represents the search options for a workflow in Jira.
type WorkflowSearchOptions struct {
	WorkflowName []string // The names of the workflows to search for.
	Expand       []string // The fields to expand in the response.
	QueryString  string   // The query string for the search.
	OrderBy      string   // The field to order the results by.
	IsActive     bool     // Indicates if only active workflows should be returned.
}

// WorkflowPageScheme represents a page of workflows in Jira.
type WorkflowPageScheme struct {
	Self       string            `json:"self,omitempty"`       // The URL of the page.
	NextPage   string            `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int               `json:"maxResults,omitempty"` // The maximum number of results returned.
	StartAt    int               `json:"startAt,omitempty"`    // The index of the first result returned.
	Total      int               `json:"total,omitempty"`      // The total number of results available.
	IsLast     bool              `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*WorkflowScheme `json:"values,omitempty"`     // The workflows on the page.
}

// WorkflowScheme represents a workflow in Jira.
type WorkflowScheme struct {
	ID          *WorkflowPublishedIDScheme  `json:"id,omitempty"`          // The ID of the workflow.
	Transitions []*WorkflowTransitionScheme `json:"transitions,omitempty"` // The transitions of the workflow.
	Statuses    []*WorkflowStatusScheme     `json:"statuses,omitempty"`    // The statuses of the workflow.
	Description string                      `json:"description,omitempty"` // The description of the workflow.
	IsDefault   bool                        `json:"isDefault,omitempty"`   // Indicates if the workflow is the default workflow.
}

// WorkflowPublishedIDScheme represents the published ID of a workflow in Jira.
type WorkflowPublishedIDScheme struct {
	Name     string `json:"name,omitempty"`     // The name of the workflow.
	EntityID string `json:"entityId,omitempty"` // The entity ID of the workflow.
}

// WorkflowTransitionScheme represents a transition in a workflow in Jira.
// Note that a transition can have either the deprecated to/from fields or the toStatusReference/links fields,
// but never both nor a combination.
type WorkflowTransitionScheme struct {
	Actions            []*WorkflowRuleConfigurationScheme `json:"actions,omitempty"`
	Conditions         *ConditionGroupConfigurationScheme `json:"conditions,omitempty"`
	CustomIssueEventId string                             `json:"customIssueEventId,omitempty"`
	Description        string                             `json:"description,omitempty"`
	From               []*WorkflowStatusAndPortScheme     `json:"from,omitempty"` // This field is deprecated - use toStatusReference/links instead.
	ID                 string                             `json:"id,omitempty"`
	Links              []*WorkflowTransitionLinkScheme    `json:"links,omitempty"`
	Name               string                             `json:"name,omitempty"` // The name of the transition.
	To                 *WorkflowStatusAndPortScheme       `json:"to,omitempty"`   // The status to which this transition goes.
	ToStatusReference  string                             `json:"toStatusReference,omitempty"`
	TransitionScreen   *WorkflowRuleConfigurationScheme   `json:"transitionScreen,omitempty"`
	Triggers           []*WorkflowTriggerScheme           `json:"triggers,omitempty"`
	Type               string                             `json:"type,omitempty"` // The type of the transition.
	Validators         []*WorkflowRuleConfigurationScheme `json:"validators,omitempty"`
}

type WorkflowRuleConfigurationScheme struct {
	Id      string `json:"id,omitempty"`
	RuleKey string `json:"ruleKey,omitempty"`
}

type ConditionGroupConfigurationScheme struct {
	ConditionGroups []*ConditionGroupConfigurationScheme `json:"conditionGroups,omitempty"`
	Conditions      []*WorkflowRuleConfigurationScheme   `json:"conditions,omitempty"`
	Operation       string                               `json:"operation,omitempty"`
}

type WorkflowStatusAndPortScheme struct {
	Port            int    `json:"port,omitempty"`
	StatusReference string `json:"statusReference,omitempty"`
}

type WorkflowTransitionLinkScheme struct {
	FromPort            int    `json:"fromPort,omitempty"`
	FromStatusReference string `json:"fromStatusReference,omitempty"`
	ToPort              int    `json:"toPort,omitempty"`
}

type WorkflowTriggerScheme struct {
	Id      string `json:"id,omitempty"`
	RuleKey string `json:"ruleKey,omitempty"`
}

// WorkflowTransitionScreenScheme represents a screen associated with a transition in a workflow in Jira.
type WorkflowTransitionScreenScheme struct {
	ID         string      `json:"id,omitempty"`         // The ID of the screen.
	Properties interface{} `json:"properties,omitempty"` // The properties of the screen.
}

// WorkflowTransitionRulesScheme represents the rules of a transition in a workflow in Jira.
type WorkflowTransitionRulesScheme struct {
	Conditions    []*WorkflowTransitionRuleScheme `json:"conditions,omitempty"`    // The conditions of the transition.
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`    // The validators of the transition.
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"` // The post functions of the transition.
}

// WorkflowTransitionRuleScheme represents a rule of a transition in a workflow in Jira.
type WorkflowTransitionRuleScheme struct {
	Type          string      `json:"type,omitempty"`          // The type of the rule.
	Configuration interface{} `json:"configuration,omitempty"` // The configuration of the rule.
}

// WorkflowStatusScheme represents a status in a workflow in Jira.
type WorkflowStatusScheme struct {
	ID         string                          `json:"id,omitempty"`         // The ID of the status.
	Name       string                          `json:"name,omitempty"`       // The name of the status.
	Properties *WorkflowStatusPropertiesScheme `json:"properties,omitempty"` // The properties of the status.
}

// WorkflowStatusPropertiesScheme represents the properties of a status in a workflow in Jira.
type WorkflowStatusPropertiesScheme struct {
	IssueEditable bool `json:"issueEditable,omitempty"` // Indicates if the issue is editable.
}

// WorkflowCreatedResponseScheme represents the response after a workflow is created in Jira.
type WorkflowCreatedResponseScheme struct {
	Name     string `json:"name,omitempty"`     // The name of the created workflow.
	EntityID string `json:"entityId,omitempty"` // The entity ID of the created workflow.
}

// WorkflowPayloadScheme represents the payload for creating a workflow in Jira.
type WorkflowPayloadScheme struct {
	Name        string                             `json:"name,omitempty"`        // The name of the workflow.
	Description string                             `json:"description,omitempty"` // The description of the workflow.
	Statuses    []*WorkflowTransitionScreenScheme  `json:"statuses,omitempty"`    // The statuses of the workflow.
	Transitions []*WorkflowTransitionPayloadScheme `json:"transitions,omitempty"` // The transitions of the workflow.
}

// WorkflowTransitionPayloadScheme represents the payload for a transition in a workflow in Jira.
type WorkflowTransitionPayloadScheme struct {
	Name        string                                 `json:"name,omitempty"`        // The name of the transition.
	Description string                                 `json:"description,omitempty"` // The description of the transition.
	From        []string                               `json:"from,omitempty"`        // The statuses from which this transition can be executed.
	To          string                                 `json:"to,omitempty"`          // The status to which this transition goes.
	Type        string                                 `json:"type,omitempty"`        // The type of the transition.
	Rules       *WorkflowTransitionRulePayloadScheme   `json:"rules,omitempty"`       // The rules of the transition.
	Screen      *WorkflowTransitionScreenPayloadScheme `json:"screen,omitempty"`      // The screen associated with the transition.
	Properties  interface{}                            `json:"properties,omitempty"`  // The properties of the transition.
}

// WorkflowTransitionScreenPayloadScheme represents the payload for a screen associated with a transition in a workflow in Jira.
type WorkflowTransitionScreenPayloadScheme struct {
	ID string `json:"id"` // The ID of the screen.
}

// WorkflowTransitionRulePayloadScheme represents the payload for the rules of a transition in a workflow in Jira.
type WorkflowTransitionRulePayloadScheme struct {
	Conditions    *WorkflowConditionScheme        `json:"conditions,omitempty"`    // The conditions of the transition.
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"` // The post functions of the transition.
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`    // The validators of the transition.
}

// WorkflowConditionScheme represents a condition in a workflow in Jira.
type WorkflowConditionScheme struct {
	Conditions    []*WorkflowConditionScheme `json:"conditions,omitempty"`    // The conditions of the workflow.
	Configuration interface{}                `json:"configuration,omitempty"` // The configuration of the condition.
	Operator      string                     `json:"operator,omitempty"`      // The operator of the condition.
	Type          string                     `json:"type,omitempty"`          // The type of the condition.
}

type WorkflowSearchCriteria struct {
	ProjectAndIssueTypes []*WorkflowSearchProjectIssueTypeMapping `json:"projectAndIssueTypes,omitempty"`
	WorkflowIds          []string                                 `json:"workflowIds,omitempty"`
	WorkflowNames        []string                                 `json:"workflowNames,omitempty"`
}

type WorkflowSearchProjectIssueTypeMapping struct {
	IssueTypeId string `json:"issueTypeId,omitempty"`
	ProjectId   string `json:"projectId,omitempty"`
}

type WorkflowReadResponseScheme struct {
	Statuses  []*WorkflowStatusDetailScheme `json:"statuses,omitempty"`
	Workflows []*JiraWorkflowScheme         `json:"workflows,omitempty"`
}

type JiraWorkflowScheme struct {
	Description      string                           `json:"description,omitempty"`
	Id               string                           `json:"id,omitempty"`
	IsEditable       bool                             `json:"isEditable,omitempty"`
	Name             string                           `json:"name,omitempty"`
	Scope            *WorkflowStatusScopeScheme       `json:"scope,omitempty"`
	StartPointLayout *WorkflowLayoutScheme            `json:"startPointLayout,omitempty"`
	Statuses         []*WorkflowReferenceStatusScheme `json:"statuses,omitempty"`
	TaskId           string                           `json:"taskId,omitempty"`
	Transitions      []*WorkflowTransitionScheme      `json:"transitions,omitempty"`
	Usages           []*ProjectIssueTypesScheme       `json:"usages,omitempty"`
	Version          *WorkflowDocumentVersionScheme   `json:"version,omitempty"`
}

type WorkflowLayoutScheme struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type WorkflowReferenceStatusScheme struct {
	Deprecated      bool                  `json:"deprecated,omitempty"`
	Layout          *WorkflowLayoutScheme `json:"layout,omitempty"`
	StatusReference string                `json:"statusReference,omitempty"`
}

type WorkflowDocumentVersionScheme struct {
	ID            string `json:"id,omitempty"`
	VersionNumber int    `json:"versionNumber,omitempty"`
}

type WorkflowCapabilitiesScheme struct {
	ConnectRules []*AvailableWorkflowConnectRuleScheme `json:"connectRules,omitempty"`
	EditorScope  string                                `json:"editorScope,omitempty"`
	ForgeRules   []*AvailableWorkflowForgeRuleScheme   `json:"forgeRules,omitempty"`
	ProjectTypes []string                              `json:"projectTypes,omitempty"`
	SystemRules  []*AvailableWorkflowSystemRuleScheme  `json:"systemRules,omitempty"`
	TriggerRules []*AvailableWorkflowTriggers          `json:"triggerRules,omitempty"`
}

type AvailableWorkflowConnectRuleScheme struct {
	AddonKey    string `json:"addonKey,omitempty"`
	CreateUrl   string `json:"createUrl,omitempty"`
	Description string `json:"description,omitempty"`
	EditUrl     string `json:"editUrl,omitempty"`
	ModuleKey   string `json:"moduleKey,omitempty"`
	Name        string `json:"name,omitempty"`
	RuleKey     string `json:"ruleKey,omitempty"`
	RuleType    string `json:"ruleType,omitempty"`
	ViewUrl     string `json:"viewUrl,omitempty"`
}

type AvailableWorkflowForgeRuleScheme struct {
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	RuleKey     string `json:"ruleKey,omitempty"`
	RuleType    string `json:"ruleType,omitempty"`
}

type AvailableWorkflowSystemRuleScheme struct {
	Description                     string   `json:"description,omitempty"`
	IncompatibleRuleKeys            []string `json:"incompatibleRuleKeys,omitempty"`
	IsAvailableForInitialTransition bool     `json:"isAvailableForInitialTransition,omitempty"`
	IsVisible                       bool     `json:"isVisible,omitempty"`
	Name                            string   `json:"name,omitempty"`
	RuleKey                         string   `json:"ruleKey,omitempty"`
	RuleType                        string   `json:"ruleType,omitempty"`
}

type AvailableWorkflowTriggers struct {
	AvailableTypes []*AvailableWorkflowTriggerTypeScheme `json:"availableTypes,omitempty"`
	RuleKey        string                                `json:"ruleKey,omitempty"`
}

type AvailableWorkflowTriggerTypeScheme struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type WorkflowCreatesPayload struct {
	Scope     *WorkflowScopeScheme          `json:"scope,omitempty"`
	Statuses  []*WorkflowStatusUpdateScheme `json:"statuses,omitempty"`
	Workflows []*WorkflowCreateScheme       `json:"workflows,omitempty"`
}

func (w *WorkflowCreatesPayload) AddStatus(status *WorkflowStatusUpdateScheme, layout *WorkflowLayoutScheme, transition *TransitionUpdateDTOScheme) {
	w.Statuses = append(w.Statuses, status)

	for _, workflow := range w.Workflows {
		workflow.Statuses = append(workflow.Statuses, &StatusLayoutUpdateScheme{
			StatusReference: status.StatusReference,
			Layout:          layout,
		})
		workflow.Transitions = append(workflow.Transitions, transition)
	}
}

type WorkflowScopeScheme struct {
	Project *WorkflowScopeProjectScheme `json:"project,omitempty"`
	Type    string                      `json:"type,omitempty"`
}

type WorkflowScopeProjectScheme struct {
	ID string `json:"id,omitempty"`
}

type WorkflowStatusUpdateScheme struct {
	Description     string `json:"description,omitempty"`
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	StatusCategory  string `json:"statusCategory,omitempty"`
	StatusReference string `json:"statusReference,omitempty"`
}

type WorkflowCreateScheme struct {
	Description      string                       `json:"description,omitempty"`
	Name             string                       `json:"name,omitempty"`
	StartPointLayout *WorkflowLayoutScheme        `json:"startPointLayout,omitempty"`
	Statuses         []*StatusLayoutUpdateScheme  `json:"statuses,omitempty"`
	Transitions      []*TransitionUpdateDTOScheme `json:"transitions,omitempty"`
}

type StatusLayoutUpdateScheme struct {
	Layout          *WorkflowLayoutScheme `json:"layout,omitempty"`
	StatusReference string                `json:"statusReference"`
}

type TransitionUpdateDTOScheme struct {
	Actions            []*WorkflowRuleConfigurationScheme `json:"actions,omitempty"`
	Conditions         *ConditionGroupUpdateScheme        `json:"conditions,omitempty"`
	CustomIssueEventId string                             `json:"customIssueEventId,omitempty"`
	Description        string                             `json:"description,omitempty"`
	From               []*StatusReferenceAndPortScheme    `json:"from,omitempty"`
	ID                 string                             `json:"id,omitempty"`
	Links              []*WorkflowTransitionLinkScheme    `json:"links,omitempty"`
	Name               string                             `json:"name,omitempty"`
	To                 *StatusReferenceAndPortScheme      `json:"to,omitempty"`
	ToStatusReference  string                             `json:"toStatusReference,omitempty"`
	TransitionScreen   *WorkflowRuleConfigurationScheme   `json:"transitionScreen,omitempty"`
	Triggers           []*WorkflowTriggerScheme           `json:"triggers,omitempty"`
	Type               string                             `json:"type,omitempty"`
	Validators         []*WorkflowRuleConfigurationScheme `json:"validators,omitempty"`
}

type ConditionGroupUpdateScheme struct {
	ConditionGroups []*ConditionGroupUpdateScheme      `json:"conditionGroups,omitempty"`
	Conditions      []*WorkflowRuleConfigurationScheme `json:"conditions,omitempty"`
	Operation       string                             `json:"operation,omitempty"`
}

type StatusReferenceAndPortScheme struct {
	Port            int    `json:"port,omitempty"`
	StatusReference string `json:"statusReference,omitempty"`
}

type WorkflowCreateResponseScheme struct {
	Statuses  []*JiraWorkflowStatusScheme `json:"statuses,omitempty"`
	Workflows []*JiraWorkflowScheme       `json:"workflows,omitempty"`
}

type JiraWorkflowStatusScheme struct {
	Description     string                     `json:"description,omitempty"`
	ID              string                     `json:"id,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Scope           *WorkflowScopeScheme       `json:"scope,omitempty"`
	StatusCategory  string                     `json:"statusCategory,omitempty"`
	StatusReference string                     `json:"statusReference,omitempty"`
	Usages          []*ProjectIssueTypesScheme `json:"usages,omitempty"`
}

type ValidationOptionsForCreateScheme struct {
	Payload *WorkflowCreatesPayload       `json:"payload,omitempty"`
	Options *ValidationOptionsLevelScheme `json:"validationOptions"`
}

type ValidationOptionsLevelScheme struct {
	Levels []string `json:"levels,omitempty"`
}

type WorkflowValidationErrorListScheme struct {
	Errors []*WorkflowValidationErrorScheme `json:"errors,omitempty"`
}

type WorkflowValidationErrorScheme struct {
	Code             string                          `json:"code,omitempty"`
	ElementReference *WorkflowElementReferenceScheme `json:"elementReference,omitempty"`
	Level            string                          `json:"level,omitempty"`
	Message          string                          `json:"message,omitempty"`
	Type             string                          `json:"type,omitempty"`
}

type WorkflowElementReferenceScheme struct {
	PropertyKey            string                         `json:"propertyKey,omitempty"`
	RuleID                 string                         `json:"ruleId,omitempty"`
	StatusMappingReference *ProjectAndIssueTypePairScheme `json:"statusMappingReference,omitempty"`
	StatusReference        string                         `json:"statusReference,omitempty"`
	TransitionID           string                         `json:"transitionId,omitempty"`
}

type ProjectAndIssueTypePairScheme struct {
	IssueTypeID string `json:"issueTypeId,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
}

type WorkflowUpdatesPayloadScheme struct {
	Statuses  []*WorkflowStatusUpdateScheme `json:"statuses,omitempty"`
	Workflows []*WorkflowUpdateScheme       `json:"workflows,omitempty"`
}

type WorkflowUpdateScheme struct {
	DefaultStatusMappings []*StatusMigrationScheme       `json:"defaultStatusMappings,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ID                    string                         `json:"id,omitempty"`
	StartPointLayout      *WorkflowLayoutScheme          `json:"startPointLayout,omitempty"`
	StatusMappings        []*StatusMappingDTOScheme      `json:"statusMappings,omitempty"`
	Version               *WorkflowDocumentVersionScheme `json:"version,omitempty"`
}

type StatusMigrationScheme struct {
	NewStatusReference string `json:"newStatusReference,omitempty"`
	OldStatusReference string `json:"oldStatusReference,omitempty"`
}

type StatusMappingDTOScheme struct {
	IssueTypeID      string                       `json:"issueTypeId,omitempty"`
	ProjectID        string                       `json:"projectId,omitempty"`
	StatusMigrations []*StatusMigrationScheme     `json:"statusMigrations,omitempty"`
	Statuses         []*StatusLayoutUpdateScheme  `json:"statuses,omitempty"`
	Transitions      []*TransitionUpdateDTOScheme `json:"transitions,omitempty"`
}

type WorkflowUpdateResponseScheme struct {
	Statuses  []*JiraWorkflowStatusScheme `json:"statuses,omitempty"`
	TaskID    string                      `json:"taskId,omitempty"`
	Workflows []*JiraWorkflowScheme       `json:"workflows,omitempty"`
}

type ValidationOptionsForUpdateScheme struct {
	Payload *WorkflowUpdatesPayloadScheme `json:"payload,omitempty"`
	Options *ValidationOptionsLevelScheme `json:"validationOptions,omitempty"`
}
