# CI tests

```mermaid
stateDiagram-v2
    workflow : ci-tests.yml - CI tests
    state workflow {
        golangci-lint: golangci-lint
        state golangci-lint {
            [*] --> step0golangci-lint
            step0golangci-lint : use gh token
            step0golangci-lint --> step1golangci-lint
            step1golangci-lint : Checkout Tyk Analytics
            step1golangci-lint --> step2golangci-lint
            step2golangci-lint : Download golangci-lint
            step2golangci-lint --> step3golangci-lint
            step3golangci-lint : golangci-lint
            step3golangci-lint --> sonar-cloud-analysis
        }

        sonar-cloud-analysis: sonar-cloud-analysis
        state sonar-cloud-analysis {
            [*] --> step0sonar-cloud-analysis
            step0sonar-cloud-analysis : Checkout Tyk Analytics
            step0sonar-cloud-analysis --> step1sonar-cloud-analysis
            step1sonar-cloud-analysis : Fetch base branch
            step1sonar-cloud-analysis --> step2sonar-cloud-analysis
            step2sonar-cloud-analysis : Setup Golang
            step2sonar-cloud-analysis --> step3sonar-cloud-analysis
            step3sonar-cloud-analysis : Download coverage artifacts
            step3sonar-cloud-analysis --> step4sonar-cloud-analysis
            step4sonar-cloud-analysis : Download golangcilint artifacts
            step4sonar-cloud-analysis --> step5sonar-cloud-analysis
            step5sonar-cloud-analysis : Check reports existence
            step5sonar-cloud-analysis --> step6sonar-cloud-analysis
            step6sonar-cloud-analysis : Install Dependencies
            step6sonar-cloud-analysis --> step7sonar-cloud-analysis
            step7sonar-cloud-analysis : merge reports
            step7sonar-cloud-analysis --> step8sonar-cloud-analysis
            step8sonar-cloud-analysis : SonarCloud Scan
        }

        test: ${{ matrix.databases }}
        state test {
            [*] --> step0test
            step0test : Checkout Tyk Analytics
            step0test --> step1test
            step1test : Checkout Tyk Analytics UI
            step1test --> step2test
            step2test : Setup Golang
            step2test --> step3test
            step3test : Cache
            step3test --> step4test
            step4test : Install Dependencies
            step4test --> step5test
            step5test : Fetch base branch
            step5test --> step6test
            step6test : ignore vendor and use go mod
            step6test --> step7test
            step7test : Start MongoDB
            step7test --> step8test
            step8test : Run Dashboard Tests
            step8test --> sonar-cloud-analysis
        }
    }
```
