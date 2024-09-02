Semgrep version: 1.76.0
Errors reported: 0
Path scanned: 515
Results: 20
Errors: 0


~~~yaml
- file: gateway/api_definition_test.go
  line: 1040
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    		ts := StartTest(nil)

- file: gateway/api_test.go
  line: 1718
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/auth_manager_test.go
  line: 263
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    		ts := StartTest(conf)

- file: gateway/auth_manager_test.go
  line: 279
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    		ts := StartTest(conf)

- file: gateway/cert_test.go
  line: 1605
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    		ts := StartTest(conf)

- file: gateway/gateway_test.go
  line: 645
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/mw_graphql_complexity_test.go
  line: 252
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/mw_persist_graphql_operation_test.go
  line: 46
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/mw_persist_graphql_operation_test.go
  line: 135
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/mw_persist_graphql_operation_test.go
  line: 240
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/mw_persist_graphql_operation_test.go
  line: 289
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/mw_persist_graphql_operation_test.go
  line: 392
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(nil)

- file: gateway/reverse_proxy_test.go
  line: 864
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	g := StartTest(nil)

- file: gateway/reverse_proxy_test.go
  line: 1197
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	g := StartTest(nil)

- file: gateway/rpc_storage_handler_test.go
  line: 372
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    				ts := StartTest(func(globalConf *config.Config) {
				})

- file: gateway/rpc_storage_handler_test.go
  line: 391
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    				ts := StartTest(func(globalConf *config.Config) {
					globalConf.SlaveOptions.GroupID = "group"
					globalConf.DBAppConfOptions.Tags = []string{"tag1"}
					globalConf.LivenessCheck.CheckDuration = 1000000000
					globalConf.SlaveOptions.APIKey = "apikey-test"
				})

- file: gateway/rpc_storage_handler_test.go
  line: 415
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    				ts := StartTest(func(globalConf *config.Config) {
					globalConf.SlaveOptions.GroupID = "group"
					globalConf.DBAppConfOptions.Tags = []string{"tag1"}
					globalConf.LivenessCheck.CheckDuration = 1000000000
				})

- file: gateway/rpc_storage_handler_test.go
  line: 451
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    				ts := StartTest(func(globalConf *config.Config) {
					globalConf.SlaveOptions.GroupID = "group"
					globalConf.DBAppConfOptions.Tags = []string{"tag1", "tag2"}
					globalConf.LivenessCheck.CheckDuration = 1000000000
				})

- file: gateway/rpc_storage_handler_test.go
  line: 475
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    				ts := StartTest(func(globalConf *config.Config) {
					globalConf.SlaveOptions.GroupID = "group"
					globalConf.DBAppConfOptions.Tags = []string{"tag1", "tag2"}
					globalConf.LivenessCheck.CheckDuration = 1000000000
					globalConf.DBAppConfOptions.NodeIsSegmented = true
				})

- file: gateway/server_test.go
  line: 165
  message: startest opened without corresponding close
  check: host.rules.tyk.starttest-never-closed
  example: |
    	ts := StartTest(func(globalConf *config.Config) {
		globalConf.ResourceSync.RetryAttempts = retryAttempts
		globalConf.ResourceSync.Interval = 1
	})


~~~
