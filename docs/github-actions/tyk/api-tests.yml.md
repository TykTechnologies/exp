# API integration Tests

```mermaid
stateDiagram-v2
    workflow : api-tests.yml - API integration Tests
    state workflow {
        test: Test
        state test {
            [*] --> step0test
            step0test : Set up Python 3.7
            step0test --> step1test
            step1test : Fix private module deps
            step1test --> step2test
            step2test : Checkout
            step2test --> step3test
            step3test : Check if test framework branch exists
            step3test --> step4test
            step4test : Checkout test repository
            step4test --> step5test
            step5test : Check if dashboard branch exists
            step5test --> step6test
            step6test : Checkout dashboard
            step6test --> step7test
            step7test : start docker compose
            step7test --> step8test
            step8test : Install test dependecies
            step8test --> step9test
            step9test : Lint with flake8
            step9test --> step10test
            step10test : Waiting for dashboard
            step10test --> step11test
            step11test : Test with pytest
            step11test --> step12test
            step12test : Archive Integration tests report
            step12test --> step13test
            step13test : Notify slack
            step13test --> step14test
            step14test : Comment on PR
            step14test --> step15test
            step15test : Xray update
            step15test --> step16test
            step16test : Getting gateway logs on failure
        }
    }
```
