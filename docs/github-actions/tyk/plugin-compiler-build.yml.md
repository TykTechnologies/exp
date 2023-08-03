# Plugin-compiler build

```mermaid
stateDiagram-v2
    workflow : plugin-compiler-build.yml - Plugin-compiler build
    state workflow {
        docker-build: 
        state docker-build {
            [*] --> step0docker-build
            step0docker-build : Checkout
            step0docker-build --> step1docker-build
            step1docker-build : Configure AWS Credentials
            step1docker-build --> step2docker-build
            step2docker-build : Login to AWS ECR
            step2docker-build --> step3docker-build
            step3docker-build : Set docker metadata
            step3docker-build --> step4docker-build
            step4docker-build : Login to Dockerhub
            step4docker-build --> step5docker-build
            step5docker-build : Build and push to dockerhub/ECR
        }
    }

```
