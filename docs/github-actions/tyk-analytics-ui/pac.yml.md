# Policy as Code

```mermaid
stateDiagram-v2
    workflow : pac.yml - Policy as Code
    state workflow {
        terraform: Terraform
        state terraform {
            [*] --> step0terraform
            step0terraform : Checkout
            step0terraform --> step3terraform
            step3terraform : Terraform Init
            step3terraform --> step4terraform
            step4terraform : Terraform Validate
            step4terraform --> step5terraform
            step5terraform : Terraform Plan
            step5terraform --> step6terraform
            step6terraform : Update Pull Request
            step6terraform --> step7terraform
            step7terraform : Terraform Plan Status
        }
    }
```
