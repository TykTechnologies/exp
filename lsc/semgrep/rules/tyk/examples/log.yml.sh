#!/bin/bash
set -e

# Run this script to regenerate the various log rules.

function rewriteLog() {
	echo "---"
	echo "rules:"

	local prefix=$1
	local log=$2
	sed -ze "s/:prefix:/$prefix/g;s/:log:/$log/g" log.yml.tpl
}

function rewriteLogs() {
	# For each prefixed logger with a global variable...

	rewriteLog "main" "mainLog" > ../log-mainLog.yml
	rewriteLog "certs" "certLog" > ../log-certLog.yml
	rewriteLog "pub-sub" "pubSubLog" > ../log-pubSubLog.yml
	rewriteLog "dashboard" "dashLog" > ../log-dashLog.yml

	rewriteLog "api" "apiLog" > ../log-apiLog.yml
	rewriteLog "host-check-mgr" "hostCheckLog" > ../log-hostCheckLog.yml
	rewriteLog "coprocess" "coprocessLog" > ../log-coprocessLog.yml
	rewriteLog "python" "pythonLog" > ../log-pythonLog.yml
	rewriteLog "webhooks" "webhookLog" > ../log-webhookLog.yml
}

rewriteLogs