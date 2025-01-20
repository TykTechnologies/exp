# Semgrep scan report

| **Key**             | **Value**                           |
|----------------------|-------------------------------------|
| **Semgrep version**  | `{{ .version }}`                   |
| **Errors reported**  | `{{ .errors | len }}`              |
| **Path scanned**     | `{{ with .paths }}{{ .scanned | len }}{{else}}0{{end}}` |
| **Results**          | `{{ .results | len }}`             |
| **Errors**           | `{{ .errors | len }}`             |

{{if gt (len .results) 0}}
~~~yaml
{{ range .results -}}
- file: {{ .path }}
  line: {{ .start.line }}
  message: {{ .extra.message }}
  check: {{ .check_id }}
{{if ne .extra.lines "requires login" }}
  example: |
    {{ .extra.lines }}
{{ end }}{{ end -}}
~~~
{{end}}{{if gt (len .errors) 0}}
Errors reported:

~~~yaml
{{ range .errors -}}
- message: {{ .message }}
{{if .spans }}{{ with index .spans 0 }}
  file: {{ or .file "N/A" }}
  line: {{ or .start.line "N/A" }}
{{end}}
  rule: {{ or .rule_id "N/A" -}}{{ end -}}{{ end -}}
~~~

{{else}}
Checks by occurence:
{{ range .checks }}
- {{ .Count }} {{ .CheckID -}}
{{ end }}

{{ end }}
