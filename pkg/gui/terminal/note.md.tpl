# {{.Title}}

{{range .Synopsis}}
{{.}}
{{end}}

{{range .Headlines}}
## {{.Title}}


{{range .ActionItems}}
- {{.Title}}
`{{if .Done}} Done {{else}} In-progress {{end}}`
{{end}}

{{end}}

# Discussions
{{range .Dialogs}}
- `{{.Timestamp}}`: {{.Person}}: {{.Message}}
{{end}}
