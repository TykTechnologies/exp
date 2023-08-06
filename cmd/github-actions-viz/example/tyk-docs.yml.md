# Tyk Docs sync for release

```mermaid
stateDiagram-v2
    workflow : tyk-docs.yml - Tyk Docs sync for release
    state workflow {
        sanitize: Sanitize inputs
        state sanitize {
            [*] --> step0sanitize
            step0sanitize : Sanitize JIRA input
            step0sanitize --> configs
            step0sanitize --> dashboard
            step0sanitize --> finish
            step0sanitize --> gateway
        }

        configs: Configuration docs
        state configs {
            [*] --> step0configs
            step0configs : Install Config Generator
            step0configs --> step1configs
            step1configs : Set up env
            step1configs --> step2configs
            step2configs : Sync Gateway
            step2configs --> step3configs
            step3configs : Sync Dashboard
            step3configs --> step4configs
            step4configs : Sync Pump
            step4configs --> step5configs
            step5configs : Sync MDCB
            step5configs --> step6configs
            step6configs : Store docs
            step6configs --> finish
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

        finish: Open PR against tyk-docs
        state finish {
            [*] --> step0finish
            step0finish : Restore artifacts
            step0finish --> step1finish
            step1finish : Set up env
            step1finish --> step2finish
            step2finish : Checkout Docs
            step2finish --> step3finish
            step3finish : Write out docs
            step3finish --> step4finish
            step4finish : Raise tyk-docs PR
        }
    }
```
