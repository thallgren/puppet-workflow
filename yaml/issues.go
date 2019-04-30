package yaml

import "github.com/lyraproj/issue/issue"

const (
	BadParameter      = `WF_BAD_PARAMETER`
	FieldTypeMismatch = `WF_FIELD_TYPE_MISMATCH`
	NotOneDefinition  = `WF_NOT_ONE_DEFINITION`
	NotStep           = `WF_NOT_STEP`
	InvalidTypeName   = `WF_INVALID_TYPE_NAME`
)

func init() {
	issue.Hard2(FieldTypeMismatch, `%{step}: expected %{field} to be a %{expected}, got %{actual}`, issue.HF{`step`: issue.Label})
	issue.Hard2(BadParameter, `%{step}: element %{name} is not a valid %{parameterType} parameter`, issue.HF{`step`: issue.Label})
	issue.Hard(NotOneDefinition, `expected exactly one top level key (the step name)`)
	issue.Hard(NotStep, `step hash must contain workflow "steps" or a resource "state"`)
	issue.Hard(InvalidTypeName, `invalid type name '%{name}'. A type name must consist of one to many capitalized segments separated with '::'`)
}
