package actionlint

import (
	"regexp"
)

// BrandingColors is a set of colors allowed at branding.color in action.yaml.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#brandingcolor
var BrandingColors = map[string]struct{}{
	"white":     {},
	"black":     {},
	"yellow":    {},
	"blue":      {},
	"green":     {},
	"orange":    {},
	"red":       {},
	"purple":    {},
	"gray-dark": {},
}

// BrandingIcons is a set of icon names allowed at branding.icon in action.yaml.
// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#brandingicon
var BrandingIcons = map[string]struct{}{
	"activity":           {},
	"airplay":            {},
	"alert-circle":       {},
	"alert-octagon":      {},
	"alert-triangle":     {},
	"align-center":       {},
	"align-justify":      {},
	"align-left":         {},
	"align-right":        {},
	"anchor":             {},
	"aperture":           {},
	"archive":            {},
	"arrow-down-circle":  {},
	"arrow-down-left":    {},
	"arrow-down-right":   {},
	"arrow-down":         {},
	"arrow-left-circle":  {},
	"arrow-left":         {},
	"arrow-right-circle": {},
	"arrow-right":        {},
	"arrow-up-circle":    {},
	"arrow-up-left":      {},
	"arrow-up-right":     {},
	"arrow-up":           {},
	"at-sign":            {},
	"award":              {},
	"bar-chart-2":        {},
	"bar-chart":          {},
	"battery-charging":   {},
	"battery":            {},
	"bell-off":           {},
	"bell":               {},
	"bluetooth":          {},
	"bold":               {},
	"book-open":          {},
	"book":               {},
	"bookmark":           {},
	"box":                {},
	"briefcase":          {},
	"calendar":           {},
	"camera-off":         {},
	"camera":             {},
	"cast":               {},
	"check-circle":       {},
	"check-square":       {},
	"check":              {},
	"chevron-down":       {},
	"chevron-left":       {},
	"chevron-right":      {},
	"chevron-up":         {},
	"chevrons-down":      {},
	"chevrons-left":      {},
	"chevrons-right":     {},
	"chevrons-up":        {},
	"circle":             {},
	"clipboard":          {},
	"clock":              {},
	"cloud-drizzle":      {},
	"cloud-lightning":    {},
	"cloud-off":          {},
	"cloud-rain":         {},
	"cloud-snow":         {},
	"cloud":              {},
	"code":               {},
	"command":            {},
	"compass":            {},
	"copy":               {},
	"corner-down-left":   {},
	"corner-down-right":  {},
	"corner-left-down":   {},
	"corner-left-up":     {},
	"corner-right-down":  {},
	"corner-right-up":    {},
	"corner-up-left":     {},
	"corner-up-right":    {},
	"cpu":                {},
	"credit-card":        {},
	"crop":               {},
	"crosshair":          {},
	"database":           {},
	"delete":             {},
	"disc":               {},
	"dollar-sign":        {},
	"download-cloud":     {},
	"download":           {},
	"droplet":            {},
	"edit-2":             {},
	"edit-3":             {},
	"edit":               {},
	"external-link":      {},
	"eye-off":            {},
	"eye":                {},
	"fast-forward":       {},
	"feather":            {},
	"file-minus":         {},
	"file-plus":          {},
	"file-text":          {},
	"file":               {},
	"film":               {},
	"filter":             {},
	"flag":               {},
	"folder-minus":       {},
	"folder-plus":        {},
	"folder":             {},
	"gift":               {},
	"git-branch":         {},
	"git-commit":         {},
	"git-merge":          {},
	"git-pull-request":   {},
	"globe":              {},
	"grid":               {},
	"hard-drive":         {},
	"hash":               {},
	"headphones":         {},
	"heart":              {},
	"help-circle":        {},
	"home":               {},
	"image":              {},
	"inbox":              {},
	"info":               {},
	"italic":             {},
	"layers":             {},
	"layout":             {},
	"life-buoy":          {},
	"link-2":             {},
	"link":               {},
	"list":               {},
	"loader":             {},
	"lock":               {},
	"log-in":             {},
	"log-out":            {},
	"mail":               {},
	"map-pin":            {},
	"map":                {},
	"maximize-2":         {},
	"maximize":           {},
	"menu":               {},
	"message-circle":     {},
	"message-square":     {},
	"mic-off":            {},
	"mic":                {},
	"minimize-2":         {},
	"minimize":           {},
	"minus-circle":       {},
	"minus-square":       {},
	"minus":              {},
	"monitor":            {},
	"moon":               {},
	"more-horizontal":    {},
	"more-vertical":      {},
	"move":               {},
	"music":              {},
	"navigation-2":       {},
	"navigation":         {},
	"octagon":            {},
	"package":            {},
	"paperclip":          {},
	"pause-circle":       {},
	"pause":              {},
	"percent":            {},
	"phone-call":         {},
	"phone-forwarded":    {},
	"phone-incoming":     {},
	"phone-missed":       {},
	"phone-off":          {},
	"phone-outgoing":     {},
	"phone":              {},
	"pie-chart":          {},
	"play-circle":        {},
	"play":               {},
	"plus-circle":        {},
	"plus-square":        {},
	"plus":               {},
	"pocket":             {},
	"power":              {},
	"printer":            {},
	"radio":              {},
	"refresh-ccw":        {},
	"refresh-cw":         {},
	"repeat":             {},
	"rewind":             {},
	"rotate-ccw":         {},
	"rotate-cw":          {},
	"rss":                {},
	"save":               {},
	"scissors":           {},
	"search":             {},
	"send":               {},
	"server":             {},
	"settings":           {},
	"share-2":            {},
	"share":              {},
	"shield-off":         {},
	"shield":             {},
	"shopping-bag":       {},
	"shopping-cart":      {},
	"shuffle":            {},
	"sidebar":            {},
	"skip-back":          {},
	"skip-forward":       {},
	"slash":              {},
	"sliders":            {},
	"smartphone":         {},
	"speaker":            {},
	"square":             {},
	"star":               {},
	"stop-circle":        {},
	"sun":                {},
	"sunrise":            {},
	"sunset":             {},
	"table":              {},
	"tablet":             {},
	"tag":                {},
	"target":             {},
	"terminal":           {},
	"thermometer":        {},
	"thumbs-down":        {},
	"thumbs-up":          {},
	"toggle-left":        {},
	"toggle-right":       {},
	"trash-2":            {},
	"trash":              {},
	"trending-down":      {},
	"trending-up":        {},
	"triangle":           {},
	"truck":              {},
	"tv":                 {},
	"type":               {},
	"umbrella":           {},
	"underline":          {},
	"unlock":             {},
	"upload-cloud":       {},
	"upload":             {},
	"user-check":         {},
	"user-minus":         {},
	"user-plus":          {},
	"user-x":             {},
	"user":               {},
	"users":              {},
	"video-off":          {},
	"video":              {},
	"voicemail":          {},
	"volume-1":           {},
	"volume-2":           {},
	"volume-x":           {},
	"volume":             {},
	"watch":              {},
	"wifi-off":           {},
	"wifi":               {},
	"wind":               {},
	"x-circle":           {},
	"x-square":           {},
	"x":                  {},
	"zap-off":            {},
	"zap":                {},
	"zoom-in":            {},
	"zoom-out":           {},
}

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#runsimage
func isImageOnDockerRegistry(image string) bool { _ = "STUB: not implemented"; return false }

// RuleAction is a rule to check running action in steps of jobs.
// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#jobsjob_idstepsuses
type RuleAction struct {
	RuleBase
	cache *LocalActionsCache
}

// NewRuleAction creates new RuleAction instance.
func NewRuleAction(cache *LocalActionsCache) *RuleAction { _ = "STUB: not implemented"; return nil }

// VisitStep is callback when visiting Step node.
func (rule *RuleAction) VisitStep(n *Step) error { _ = "STUB: not implemented"; return nil }

// Cannot parse specification made with interpolation. Give up

// Relative to repository root

// Parse {owner}/{repo}@{ref} or {owner}/{repo}/{path}@{ref}
func (rule *RuleAction) checkRepoAction(spec string, exec *ExecAction) {
	_ = "STUB: not implemented"
	return
}

// remove {ref}

// eat {owner}

func (rule *RuleAction) invalidActionFormat(pos *Pos, spec string, why string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleAction) missingRunsProp(pos *Pos, prop, ty, action, path string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleAction) checkInvalidRunsProps(pos *Pos, r *ActionMetadataRuns, ty, action, path string, props []string) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleAction) checkRunsFileExists(file, dir, prop, name string, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#runs-for-docker-container-actions
func (rule *RuleAction) checkLocalDockerActionRuns(r *ActionMetadataRuns, dir, name string, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#runs-for-composite-actions
func (rule *RuleAction) checkLocalCompositeActionRuns(r *ActionMetadataRuns, dir, name string, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#runs-for-javascript-actions
func (rule *RuleAction) checkLocalJavaScriptActionRuns(r *ActionMetadataRuns, dir, name string, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

func (rule *RuleAction) checkLocalActionInputs(meta *ActionMetadata, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions#runs
func (rule *RuleAction) checkLocalActionRuns(meta *ActionMetadata, pos *Pos) {
	_ = "STUB: not implemented"
	return
}

// Probably invalid version of Node.js runner. Assume it is JavaScript action to find as many errors as possible

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#example-using-the-github-packages-container-registry
func (rule *RuleAction) checkDockerAction(uri string, exec *ExecAction) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/creating-actions/metadata-syntax-for-github-actions
func (rule *RuleAction) checkLocalActionMetadata(meta *ActionMetadata, action *ExecAction) {
	_ = "STUB: not implemented"
	return
}

// https://docs.github.com/en/actions/learn-github-actions/workflow-syntax-for-github-actions#example-using-action-in-the-same-repository-as-the-workflow
func (rule *RuleAction) checkLocalAction(spec string, action *ExecAction) {
	_ = "STUB: not implemented"
	return
}

var reNewlineWithIndent = regexp.MustCompile(`\s*\r?\n\s*`)

func (rule *RuleAction) checkAction(meta *ActionMetadata, exec *ExecAction, describe func(*ActionMetadata) string) {
	_ = "STUB: not implemented"
	// Check specified inputs are defined in action's inputs spec
	return
}

// Note: Using required inputs cannot be avoided. So we don't report it as error (though this should not
// happen normally).

// Check mandatory inputs are specified
