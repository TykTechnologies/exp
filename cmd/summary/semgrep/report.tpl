Semgrep version: {{ .version }}
Errors reported: {{ .errors | len }}
Path scanned: {{ .paths.scanned | len }}
Results: {{ .results | len }}

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

Checks by occurence:
{{ range .checks }}
- {{ .Count }} {{ .CheckID -}}
{{ end }}
