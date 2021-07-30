package elk

import (
	"entgo.io/ent/entc/gen"
	"github.com/masseelch/elk/internal"
	"github.com/stoewer/go-strcase"
	"text/template"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -modtime=1 ./template/...

const (
	actionCreate = "create"
	actionRead   = "read"
	actionUpdate = "update"
	actionDelete = "delete"
	actionList   = "list"
)

var (
	// HTTPTemplates holds all templates for generating http handlers.
	HTTPTemplates = []*gen.Template{
		parse("template/http/create.tmpl"),
		parse("template/http/delete.tmpl"),
		parse("template/http/handler.tmpl"),
		parse("template/http/helpers.tmpl"),
		parse("template/http/list.tmpl"),
		parse("template/http/read.tmpl"),
		parse("template/http/relations.tmpl"),
		parse("template/http/response.tmpl"),
		parse("template/http/update.tmpl"),
	}
	// TemplateFuncs contains the extra template functions used by elk.
	TemplateFuncs = template.FuncMap{
		"edgesToLoad":        edgesToLoad,
		"kebab":              strcase.KebabCase,
		"needsSerialization": needsSerialization,
		"needsValidation":    needsValidation,
		"stringSlice":        stringSlice,
		"validationTags":     validationTags,
	}
)

func parse(path string) *gen.Template {
	return gen.MustParse(gen.NewTemplate(path).
		Funcs(TemplateFuncs).
		Parse(string(internal.MustAsset(path))))
}

// validationTags extracts the validation tags to use for the given action / method.
func validationTags(a interface{}, m string) string {
	if a == nil {
		return ""
	}
	an := Annotation{}
	if err := an.Decode(a); err != nil {
		return ""
	}
	if m == "create" && an.CreateValidation != "" {
		return an.CreateValidation
	}
	if m == "update" && an.UpdateValidation != "" {
		return an.UpdateValidation
	}
	return an.Validation
}

// needsValidation returns if a type needs validation for a given request type.
func needsValidation(n *gen.Type, m string) bool {
	an := Annotation{}.Name()
	for _, f := range n.Fields {
		if validationTags(f.Annotations[an], m) != "" {
			return true
		}
	}
	for _, e := range n.Edges {
		if validationTags(e.Annotations[an], m) != "" {
			return true
		}
	}
	return false
}

// stringSlice casts a given []interface{} to []string.
func stringSlice(is []interface{}) []string {
	ss := make([]string, len(is))
	for i, v := range is {
		ss[i] = v.(string)
	}
	return ss
}

// needsSerialization checks if a given field or edge is to be serialized according to its annotations and the requested
// groups.
func needsSerialization(a interface{}, g groups) (bool, error) {
	// If no groups are given on the field default is to include it in the output.
	if a == nil {
		return false, nil
	}

	// If there are groups given check if the groups match the requested ones.
	an := Annotation{}
	if err := an.Decode(a); err != nil {
		return false, err
	}
	return g.Match(an.Groups), nil
}

// func groupPermutations(n *gen.Type) map[*gen.Type][]groups {
// 	gs := make([]groups, 0)
//
// }
