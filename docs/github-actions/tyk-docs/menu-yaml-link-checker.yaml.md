# Check for missing links in menu.yaml

```mermaid
stateDiagram-v2
    workflow : menu-yaml-link-checker.yaml - Check for missing links in menu.yaml
    state workflow {
        check_menu_links: Check menu links
        state check_menu_links {
            [*] --> step0check_menu_links
            step0check_menu_links : setup Hugo
            step0check_menu_links --> step1check_menu_links
            step1check_menu_links : clone tyk-docs
            step1check_menu_links --> step2check_menu_links
            step2check_menu_links : clone tyk_libs
            step2check_menu_links --> step3check_menu_links
            step3check_menu_links : Prepare files to run script
            step3check_menu_links --> step4check_menu_links
            step4check_menu_links : setup Python
            step4check_menu_links --> step5check_menu_links
            step5check_menu_links : cache poetry install
            step5check_menu_links --> step6check_menu_links
            step6check_menu_links : install poetry
            step6check_menu_links --> step7check_menu_links
            step7check_menu_links : load cached venv
            step7check_menu_links --> step8check_menu_links
            step8check_menu_links : Install poetry dependencies except the poetry project in pyproject.toml
            step8check_menu_links --> step9check_menu_links
            step9check_menu_links : Install poetry project in pyproject.toml
            step9check_menu_links --> step10check_menu_links
            step10check_menu_links : Run script to check for missing links in menu.yaml
        }
    }
```
