package tinker

const itemTemplate string = `You find {{.Prelude.Article}} {{.Prelude.Name}} {{.Name}}.
{{- range .Components -}}
{{template "component" .}}
{{- end -}}

{{define "component"}}
The {{.Name}} is {{template "property" (index .Properties 0)}}
{{- if gt (len .Properties) 1}}{{range (slice .Properties 1)}} and {{template "property" .}}{{end}}{{end -}}.
{{- end -}}

{{define "property"}}
{{- if .Countable -}}
{{- .Verb}} {{.Attribute.Article}} {{.Attribute.Description -}}
{{- else}}
{{- .Verb}} {{.Attribute.Description -}}
{{- end}}
{{- end -}}
`
