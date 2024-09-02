# TT-11965

This test setup does the following:

- Dockerfile, uses a ubuntu:16.04 image (RHEL8 env)
- Compiles plugin using `v5.3.1-alpha2`
- Loads the plugin into gateway verifying no glibc error

Run `test` to verify, trimmed example output below:

```
task: [default] docker build --rm --progress=plain -t temp .
#0 building with "default" instance using docker driver

#1 [internal] load .dockerignore
#1 transferring context: 2B done
#1 DONE 0.0s

#2 [internal] load build definition from Dockerfile
#2 transferring dockerfile: 354B done
#2 DONE 0.0s

#3 [internal] load metadata for docker.io/library/ubuntu:16.04
#3 DONE 1.7s

#4 [1/5] FROM docker.io/library/ubuntu:16.04@sha256:1f1a2d56de1d604801a9671f301190704c25d604a416f59e03c04f5c6ffee0d6
#4 DONE 11.8s

#5 [2/5] RUN apt-get update && apt -yyy install curl
#5 DONE 15.3s

#6 [3/5] RUN curl -sL "https://packagecloud.io/tyk/tyk-gateway-unstable/packages/ubuntu/focal/tyk-gateway_5.3.1~alpha2_amd64.deb/download.deb?distro_version_id=210" -o ./tyk-gateway.deb
#6 DONE 13.9s

#7 [4/5] RUN dpkg -i ./tyk-gateway.deb
#7 0.293 Selecting previously unselected package tyk-gateway.
#7 0.296 (Reading database ... 5283 files and directories currently installed.)
#7 0.296 Preparing to unpack ./tyk-gateway.deb ...
#7 0.298 Creating user and group...
#7 0.332 Unpacking tyk-gateway (5.3.1~alpha2) ...
#7 0.685 Setting up tyk-gateway (5.3.1~alpha2) ...
#7 0.690 [32m Post Install of the install directory ownership and permissions[0m
#7 0.690 [32m Post Install of an clean install[0m
#7 0.690 [32m Reload the service unit from disk[0m
#7 0.691 Failed to connect to bus: No such file or directory
#7 0.691 [32m Unmask the service[0m
#7 0.692 [32m Set the preset flag for the service unit[0m
#7 0.693 Created symlink /etc/systemd/system/multi-user.target.wants/tyk-gateway.service, pointing to /lib/systemd/system/tyk-gateway.service.
#7 0.693 [32m Set the enabled flag for the service unit[0m
#7 0.694 Synchronizing state of tyk-gateway.service with SysV init with /lib/systemd/systemd-sysv-install...
#7 0.694 Executing /lib/systemd/systemd-sysv-install enable tyk-gateway
#7 0.713 Failed to connect to bus: No such file or directory
#7 0.714 Processing triggers for systemd (229-4ubuntu21.31) ...
#7 DONE 0.8s

#8 [5/5] RUN /opt/tyk-gateway/tyk version
#8 0.453 Release version: 5.3.1-alpha2
#8 0.453 Built by:        goreleaser
#8 0.453 Build date:      2024-04-19T17:51:49Z
#8 0.453 Commit:          dff399f805ffbf6eefad05712ecfaeca34061d7c
#8 0.453 Go version:      go1.21.8
#8 0.453 OS/Arch:         linux/amd64
#8 0.453 
#8 DONE 0.5s

#9 exporting to image
#9 exporting layers
#9 exporting layers 0.4s done
#9 writing image sha256:128de40629d503e173c911452b8f01ee31b5da40b91bcfa010fbc85efb37739d done
#9 naming to docker.io/library/temp done
#9 DONE 0.4s
task: [test] rm -f ./basic-plugin/*.so
```

Compile plugin:

```
task: [test] docker run --rm -v ./basic-plugin:/plugin-source -w /plugin-source tykio/tyk-plugin-compiler:v5.3.1-alpha2 plugin.so
PLUGIN_BUILD_PATH: /go/src/github.com/TykTechnologies/plugin_plugin
PLUGIN_SOURCE_PATH: /plugin-source
INFO: No plugin id provided, keeping go.mod as is
task: [test] cp -f ./basic-plugin/*.so ./basic-plugin/plugin.so
```

Load plugin into gateway:

```
task: [test] docker run --rm -v ./basic-plugin:/plugin-source -w /plugin-source --entrypoint=/opt/tyk-gateway/tyk temp plugin load -f plugin.so -s MyPluginPre
[file=plugin.so, symbol=MyPluginPre] loaded ok, got 0x7f37ab10c520
```

This verifies successful load of a plugin and referencing a symbol from the plugin.

Some debug tail output:

```
task: [test] strings ./basic-plugin/plugin.so | grep main.go
github.com/jinzhu/now@v1.1.2/main.go
task: [clean] docker rmi -f temp
Untagged: temp:latest
Deleted: sha256:128de40629d503e173c911452b8f01ee31b5da40b91bcfa010fbc85efb37739d
```
