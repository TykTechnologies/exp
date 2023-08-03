# Release

```mermaid
stateDiagram-v2
    workflow : release.yml - Release
    state workflow {
        ci: ci
        state ci {
            [*] --> step0ci
            step0ci : Shallow checkout of tyk-analytics
            step0ci --> step2ci
            step2ci : Login to Amazon ECR
            step2ci --> step4ci
            step4ci : Docker metadata
            step4ci --> step7ci
            step7ci : CI build
            step7ci --> sbom
            step7ci --> tat
        }

        goreleaser: ${{ matrix.golang_cross }}
        state goreleaser {
            [*] --> step0goreleaser
            step0goreleaser : Fix private module deps
            step0goreleaser --> step1goreleaser
            step1goreleaser : Checkout of tyk-analytics
            step1goreleaser --> step2goreleaser
            step2goreleaser : Add Git safe.directory
            step2goreleaser --> step5goreleaser
            step5goreleaser : Login to DockerHub
            step5goreleaser --> step6goreleaser
            step6goreleaser : Login to Cloudsmith
            step6goreleaser --> step7goreleaser
            step7goreleaser : Unlock agent and set tag
            step7goreleaser --> step8goreleaser
            step8goreleaser : Delete old release assets
            step8goreleaser --> ci
            step8goreleaser --> smoke-tests
            step8goreleaser --> upgrade-deb
            step8goreleaser --> upgrade-rpm
        }

        sbom: sbom
        smoke-tests: smoke-tests
        state smoke-tests {
            [*] --> step1smoke-tests
            step1smoke-tests : Run tests
        }

        tat: tat
        upgrade-deb: upgrade-deb
        state upgrade-deb {
            [*] --> step4upgrade-deb
            step4upgrade-deb : generate dockerfile
            step4upgrade-deb --> step5upgrade-deb
            step5upgrade-deb : install on ${{ matrix.distro }}
        }

        upgrade-rpm: upgrade-rpm
        state upgrade-rpm {
            [*] --> step3upgrade-rpm
            step3upgrade-rpm : generate dockerfile
            step3upgrade-rpm --> step4upgrade-rpm
            step4upgrade-rpm : install on ${{ matrix.distro }}
        }
    }
```
