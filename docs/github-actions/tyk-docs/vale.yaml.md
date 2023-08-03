# Linting

```mermaid
stateDiagram-v2
    workflow : vale.yaml - Linting
    state workflow {
        prose: prose
        state prose {
            [*] --> step0prose
            step0prose : Checkout
            step0prose --> step1prose
            step1prose : Vale
        }
    }
```
