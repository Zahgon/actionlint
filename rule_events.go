package actionlint

//go:generate go run ./scripts/generate-webhook-events ./all_webhooks.go

// RuleEvents is a rule to check 'on' field in workflow.
// https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows
type RuleEvents struct {
	RuleBase
}

// NewRuleEvents creates new RuleEvents instance.
func NewRuleEvents() *RuleEvents { _ = "STUB: not implemented"; return nil }

// VisitWorkflowPre is callback when visiting Workflow node before visiting its children.
func (rule *RuleEvents) VisitWorkflowPre(n *Workflow) error { _ = "STUB: not implemented"; return nil }

func (rule *RuleEvents) checkEvent(event Event) { _ = "STUB: not implemented"; return }

// Nothing to do

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#onschedule
func (rule *RuleEvents) checkCron(spec *String) { _ = "STUB: not implemented"; return }

// (#14) https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#scheduled-events
//
// > The shortest interval you can run scheduled workflows is once every 5 minutes.

// https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax#onschedule
func (rule *RuleEvents) checkTimezone(tz *String) {
	_ = "STUB: not implemented"
	// `time.LoadLocation` accepts special values "", "UTC", and "Local" but they are not correct IANA timezone names.
	return
}

func (rule *RuleEvents) filterNotAvailable(pos *Pos, filter, hook string, available []string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleEvents) checkExclusiveFilters(filter, ignore *WebhookEventFilter, hook string, available []string) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#webhook-events
func (rule *RuleEvents) checkWebhookEvent(event *WebhookEvent) { _ = "STUB: not implemented"; return }

// Some filters are available with specific events and exclusive
// - on.merge_group.<branches|branches-ignore>
// - on.<push|pull_request|pull_request_target>.<paths|paths-ignore>
// - on.push.<branches|tags|branches-ignore|tags-ignore>
// - on.<pull_request|pull_request_target>.<branches|branches-ignore>
// - on.workflow_run.<branches|branches-ignore>

func (rule *RuleEvents) checkTypes(hook *String, types []*String, expected []string) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/learn-github-actions/reusing-workflows
func (rule *RuleEvents) checkWorkflowCallEvent(event *WorkflowCallEvent) {
	_ = "STUB: not implemented"
	return
}

// ${{ }} is available in the default value

// https://github.blog/changelog/2021-11-10-github-actions-input-types-for-manual-workflows/
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#onworkflow_dispatchinputs
func (rule *RuleEvents) checkWorkflowDispatchEvent(event *WorkflowDispatchEvent) {
	_ = "STUB: not implemented"
	return
}

// TODO: Can some check be done for WorkflowDispatchEventInputTypeEnvironment?
// What is suitable for default value of the type? (Or is a default value never suitable?)

// Maximum number of inputs is 25
// https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#providing-inputs
// https://github.blog/changelog/2025-12-04-actions-workflow-dispatch-workflows-now-support-25-inputs

// https://docs.github.com/en/actions/reference/workflows-and-actions/events-that-trigger-workflows#image_version_ready
func (rule *RuleEvents) checkImageVersionEvent(event *ImageVersionEvent) {
	_ = "STUB: not implemented"
	// Do nothing
	return
}
