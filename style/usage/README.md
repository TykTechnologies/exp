# Semgrep scan report

| **Key**             | **Value**                           |
|----------------------|-------------------------------------|
| **Semgrep version**  | `1.102.0`                   |
| **Errors reported**  | `0`              |
| **Path scanned**     | `1` |
| **Results**          | `2`             |
| **Errors**           | `0`             |


~~~yaml
- file: example.go
  line: 9
  message: Usage of atomic.Value detected. Consider using atomic.Pointer for better type safety.
  check: find-atomic-value-usage
- file: example.go
  line: 14
  message: Usage of atomic.Value detected. Consider using atomic.Pointer for better type safety.
  check: find-atomic-value-usage
~~~

Checks by occurence:

- 2 find-atomic-value-usage


