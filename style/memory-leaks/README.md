# Semgrep scan report

| **Key**             | **Value**                           |
|----------------------|-------------------------------------|
| **Semgrep version**  | `1.102.0`                   |
| **Errors reported**  | `1`              |
| **Path scanned**     | `1` |
| **Results**          | `1`             |
| **Errors**           | `1`             |


~~~yaml
- file: example.go
  line: 10
  message: Map value type is not a pointer. Consider using a pointer for efficiency if the value type is large.
  check: find-non-pointer-map-values
~~~

Errors reported:

~~~yaml
- message: Internal matching error when running find-non-pointer-map-values on example.go:
 An error occurred while invoking the Semgrep engine. Please help us fix this by creating an issue at https://github.com/semgrep/semgrep

metavariable-pattern failed because $VALUE does not bind to a sub-program, please check your rule
~~~


