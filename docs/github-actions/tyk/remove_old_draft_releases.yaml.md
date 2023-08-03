# Delete Draft Releases older than 5 days

```mermaid
stateDiagram-v2
    workflow : remove_old_draft_releases.yaml - Delete Draft Releases older than 5 days
    state workflow {
        build: Delete drafts
        state build {
            [*] --> step0build
            step0build : Delete drafts
        }
    }
```
