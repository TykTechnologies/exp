# Tyk API documentation (OAS) sync

```mermaid
stateDiagram-v2
    workflow : swagger-update.yaml - Tyk API documentation (OAS) sync
    state workflow {
        swagger-spec-update: swagger-spec-update
        state swagger-spec-update {
            [*] --> step0swagger-spec-update
            step0swagger-spec-update : Checkout
            step0swagger-spec-update --> step1swagger-spec-update
            step1swagger-spec-update : Checkout gateway
            step1swagger-spec-update --> step2swagger-spec-update
            step2swagger-spec-update : Copy Swagger
            step2swagger-spec-update --> step3swagger-spec-update
            step3swagger-spec-update : Raise gateway swagger changes Pull Request
            step3swagger-spec-update --> step4swagger-spec-update
            step4swagger-spec-update : Checkout dashboard
            step4swagger-spec-update --> step5swagger-spec-update
            step5swagger-spec-update : Copy Swagger
            step5swagger-spec-update --> step6swagger-spec-update
            step6swagger-spec-update : Raise dashboard swagger changes Pull Request
        }
    }
```
