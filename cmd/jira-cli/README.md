# jira-cli

A Go utility to validate Jira ticket states against allowed statuses (`In Dev`, `In Progress`, etc.) using the `jira-cli` library.

## Install

```
go install github.com/TykTechnologies/exp/cmd/jira-lint@main
```

## Configure

Set the following env vars:

- JIRA_API_TOKEN
- JIRA_API_EMAIL
- JIRA_API_URL=https://tyktech.atlassian.net

You can generate a PAT via [Atlassian UI for API Tokens](https://id.atlassian.com/manage-profile/security/api-tokens).
For GitHub actions, the token should be passed as a secret.

## Usage

```
root@carbon:~/tyk/exp/cmd/jira-cli# ./jira-cli SYSE-311 ; echo $?
2025/01/16 21:07:39 Validation failed: ticket 'SYSE-311' is in an invalid state: 'Done'. Allowed states are: [In Dev In Code Review Ready for Testing In Test In Progress In Review]
1
root@carbon:~/tyk/exp/cmd/jira-cli# ./jira-cli TT-11909 ; echo $?
2025/01/16 21:07:43 Validation succeeded: Ticket is in a valid state
0
```

If ticket is not in an allowed state, jira cli exits with a non-zero exit code.

## Allowed States

- `In Dev`
- `In Code Review`
- `Ready for Testing`
- `In Test`
- `In Progress`
- `In Review`
