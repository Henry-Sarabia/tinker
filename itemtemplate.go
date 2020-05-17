package tinker

const itemTemplate string = `You find {{getArticle .Prelude.Name}} {{.Prelude.Name}} {{.Name}}.
{{- range .Components -}}
{{template "component" .}}
{{- end -}}

{{define "component"}}
The {{.Name}} is {{template "property" (index .Properties 0)}}
{{- if gt (len .Properties) 1}}{{range (slice .Properties 1)}} and {{template "property" .}}{{end}}{{end -}}.
{{- end -}}

{{define "property"}}
{{- if .Countable -}}
{{- .Verb}} {{getArticle .Attribute.Description}} {{.Attribute.Description -}}
{{- else}}
{{- .Verb}} {{.Attribute.Description -}}
{{- end}}
{{- end -}}
`
