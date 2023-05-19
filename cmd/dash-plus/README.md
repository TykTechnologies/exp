# Dash plus

*Dashboard and Gateway, together*

This program imports dashboard and gateway, running the services on their
individual ports. However, if dashboard would define a `http.Handler`, we
would not need a secondary HTTP server for the dashboard, but could
register that handler into the Gateway.

The idea is, that Gateway has all this functionality which can be used in
front of the Dashboard as well. The dashboard provides an openapi spec,
and could itself have an api definition for the gateway to use
implemented functionality.

`go get github.com/TykTechnologies/exp/cmd/dash-plus@main`

