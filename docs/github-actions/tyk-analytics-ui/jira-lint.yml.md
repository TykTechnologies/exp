# JIRA lint

```mermaid
stateDiagram-v2
    workflow : jira-lint.yml - JIRA lint
    state workflow {
        jira_lint: Jira lint
        state jira_lint {
            jira_lint_finish: Done
            [*] --> jira_lint_finish
        }
    }
```
