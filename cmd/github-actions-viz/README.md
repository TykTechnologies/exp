# GitHub Actions Vizualizer

The tool, if ran in a github actions folder, reads in the `*.{yml,yaml}`
files and produces a mermaidjs diagram for each workflow.

To install:

`go install github.com/TykTechnologies/exp/cmd/github-actions-viz@main`

Flags:

- `-i <folder>` - input folder (defaults to .),
- `--format <md|mermaid>` - define the output format,
- `-w` writes out files to disk.

To use:

- GitHub PRs / Markdown: https://github.blog/2022-02-14-include-diagrams-markdown-files-mermaid/
- Hugo docs / Markdown: https://discourse.gohugo.io/t/correct-way-to-embed-mermaid-js/43491/3
- Live editor: https://mermaid.live/

## Example

```mermaid
stateDiagram-v2
    workflow : tyk-docs.yml - Tyk Docs sync for release
    state workflow {
        gateway: Gateway docs
        state gateway {
            [*] --> step0gateway
            step0gateway : Checkout Gateway
            step0gateway --> step1gateway
            step1gateway : Generate docs
            step1gateway --> step2gateway
            step2gateway : Store docs
            step2gateway --> finish
        }

        dashboard: Dashboard docs
        state dashboard {
            [*] --> step0dashboard
            step0dashboard : Checkout Dashboard
            step0dashboard --> step1dashboard
            step1dashboard : Generate docs
            step1dashboard --> step2dashboard
            step2dashboard : Store docs
            step2dashboard --> finish
        }

        configs: Configuration docs
        state configs {
            [*] --> step0configs
            step0configs : Checkout Config Generator Repo
            step0configs --> step1configs
            step1configs : Generate gateway config
            step1configs --> step2configs
            step2configs : Generate dashboard config
            step2configs --> step3configs
            step3configs : Generate mdcb/pump config docs
            step3configs --> step4configs
            step4configs : Store docs
            step4configs --> finish
        }

        finish: Open PR against tyk-docs
        state finish {
            [*] --> step0finish
            step0finish : Restore artifacts
            step0finish --> step1finish
            step1finish : Print artifacts
            step1finish --> step2finish
            step2finish : Checkout Docs
            step2finish --> step3finish
            step3finish : Write out swagger schema
            step3finish --> step4finish
            step4finish : Write out markdown docs
            step4finish --> step5finish
            step5finish : Sanitize JIRA input
            step5finish --> step6finish
            step6finish : Raise docs changes PR
        }
    }
```

## Other/misc

- Docker Compose variant: https://github.com/pmsipilot/docker-compose-viz
- This project uses nektos/act - to read the github actions with their data model
