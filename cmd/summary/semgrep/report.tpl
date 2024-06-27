Semgrep version: {{ .version }}
Errors reported: {{ .errors | len }}
Path scanned: {{ .paths.scanned | len }}
Results: {{ .results | len }}
Errors: {{ .errors | len }}

{{if gt (len .results) 0}}
~~~yaml
{{ range .results -}}
- file: {{ .path }}
  line: {{ .start.line }}
  message: {{ .extra.message }}
  check: {{ .check_id }}
  example: |
    {{ .extra.lines }}

{{ end }}
~~~
{{end}}
{{if gt (len .errors) 0}}
Errors reported:
~~~yaml
{{ range .errors -}}
- file: {{with index .spans 0}}{{ .file }}{{end}}
  line: {{with index .spans 0}}{{ .start.line }}{{end}}
  message: {{ .message }}
  rule: {{ .rule_id }}
{{ end }}
~~~

{{else}}

Checks by occurence:
{{ range .checks }}
- {{ .Count }} {{ .CheckID -}}
{{ end }}

{{ end }}