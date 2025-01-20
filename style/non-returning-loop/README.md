# Semgrep scan report

| **Key**             | **Value**                           |
|----------------------|-------------------------------------|
| **Semgrep version**  | `1.102.0`                   |
| **Errors reported**  | `0`              |
| **Path scanned**     | `1` |
| **Results**          | `1`             |
| **Errors**           | `0`             |


~~~yaml
- file: example.go
  line: 4
  message: Potential infinite loop detected. Ensure there is a condition to exit the loop, or include a break or return statement.
  check: find-non-returning-loops
~~~

Checks by occurence:

- 1 find-non-returning-loops


