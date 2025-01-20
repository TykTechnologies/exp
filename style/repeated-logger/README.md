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
  line: 16
  message: Repeated calls to Logger() detected. Reduce allocations/concurrency by only invoking it once.
  check: find-repeated-logger-calls
~~~

Checks by occurence:

- 1 find-repeated-logger-calls


