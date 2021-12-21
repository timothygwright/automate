// Generated by agg_bindings.sh. DO NOT EDIT
package assets

const BindData = `applications-load-gen BINDING_MODE strict
applications-service REQUIRED automate-pg-gateway cereal-service event-service pg-sidecar-service
applications-service BINDING_MODE strict
authn-service REQUIRED authz-service automate-dex automate-pg-gateway cereal-service pg-sidecar-service session-service teams-service
authn-service BINDING_MODE strict
authz-service REQUIRED automate-pg-gateway cereal-service pg-sidecar-service
authz-service BINDING_MODE strict
automate-ha-curator REQUIRED elasticsearch
automate-ha-curator BINDING_MODE strict
automate-ha-elasticsearch BINDING_MODE strict
automate-ha-elasticsidecar REQUIRED elasticsearch
automate-ha-elasticsidecar BINDING_MODE strict
automate-ha-haproxy OPTIONAL database pgleaderchk
automate-ha-haproxy BINDING_MODE strict
automate-ha-journalbeat REQUIRED elasticsearch
automate-ha-journalbeat BINDING_MODE strict
automate-ha-kibana REQUIRED elasticsearch
automate-ha-kibana BINDING_MODE strict
automate-ha-metricbeat REQUIRED elasticsearch
automate-ha-metricbeat OPTIONAL database
automate-ha-metricbeat BINDING_MODE strict
automate-ha-postgresql BINDING_MODE strict
automate-builder-api REQUIRED automate-builder-memcached automate-minio automate-pg-gateway pg-sidecar-service session-service
automate-builder-api BINDING_MODE strict
automate-builder-api-proxy REQUIRED automate-builder-api
automate-builder-api-proxy BINDING_MODE strict
automate-builder-memcached BINDING_MODE strict
automate-cds REQUIRED compliance-service
automate-cds BINDING_MODE strict
automate-cs-bookshelf REQUIRED automate-pg-gateway pg-sidecar-service
automate-cs-bookshelf BINDING_MODE strict
automate-cs-nginx REQUIRED automate-cs-bookshelf automate-cs-oc-erchef automate-es-gateway
automate-cs-nginx OPTIONAL automate-gateway
automate-cs-nginx BINDING_MODE relaxed
automate-cs-oc-bifrost REQUIRED automate-pg-gateway pg-sidecar-service
automate-cs-oc-bifrost BINDING_MODE strict
automate-cs-oc-erchef REQUIRED automate-cs-bookshelf automate-cs-oc-bifrost automate-es-gateway automate-pg-gateway pg-sidecar-service
automate-cs-oc-erchef OPTIONAL automate-gateway
automate-cs-oc-erchef BINDING_MODE relaxed
deployment-service BINDING_MODE strict
automate-dex REQUIRED automate-pg-gateway pg-sidecar-service
automate-dex BINDING_MODE strict
automate-elasticsearch REQUIRED backup-gateway
automate-elasticsearch BINDING_MODE strict
automate-es-gateway REQUIRED automate-elasticsearch
automate-es-gateway BINDING_MODE relaxed
automate-gateway OPTIONAL applications-service authn-service authz-service automate-cds compliance-service config-mgmt-service data-feed-service deployment-service event-feed-service infra-proxy-service ingest-service license-control-service local-user-service nodemanager-service notifications-service secrets-service teams-service user-settings-service
automate-gateway BINDING_MODE relaxed
automate-ha-pgleaderchk OPTIONAL database
automate-ha-pgleaderchk BINDING_MODE strict
automate-load-balancer OPTIONAL automate-builder-api-proxy automate-cs-nginx automate-dex automate-gateway automate-ui automate-workflow-nginx session-service
automate-load-balancer BINDING_MODE relaxed
automate-minio BINDING_MODE strict
automate-pg-gateway REQUIRED automate-postgresql
automate-pg-gateway BINDING_MODE relaxed
automate-postgresql BINDING_MODE strict
automate-postgresql13 BINDING_MODE strict
automate-prometheus OPTIONAL applications-service automate-gateway
automate-prometheus BINDING_MODE relaxed
automate-ui BINDING_MODE strict
automate-workflow-nginx REQUIRED automate-workflow-server
automate-workflow-nginx BINDING_MODE strict
automate-workflow-server REQUIRED automate-pg-gateway pg-sidecar-service
automate-workflow-server BINDING_MODE strict
backup-gateway BINDING_MODE strict
cereal-service REQUIRED automate-pg-gateway pg-sidecar-service
cereal-service BINDING_MODE strict
compliance-service REQUIRED authz-service automate-es-gateway automate-pg-gateway cereal-service event-service nodemanager-service pg-sidecar-service secrets-service
compliance-service OPTIONAL authn-service es-sidecar-service notifications-service
compliance-service BINDING_MODE strict
config-mgmt-service REQUIRED automate-es-gateway automate-pg-gateway pg-sidecar-service
config-mgmt-service BINDING_MODE strict
data-feed-service REQUIRED automate-pg-gateway cereal-service compliance-service config-mgmt-service pg-sidecar-service secrets-service
data-feed-service BINDING_MODE strict
es-sidecar-service REQUIRED automate-elasticsearch automate-es-gateway
es-sidecar-service BINDING_MODE strict
event-feed-service REQUIRED authz-service automate-es-gateway cereal-service
event-feed-service OPTIONAL es-sidecar-service
event-feed-service BINDING_MODE strict
event-gateway REQUIRED authn-service authz-service event-service
event-gateway BINDING_MODE strict
event-service BINDING_MODE strict
infra-proxy-service REQUIRED authz-service automate-pg-gateway pg-sidecar-service secrets-service
infra-proxy-service BINDING_MODE strict
ingest-service REQUIRED authz-service automate-es-gateway automate-pg-gateway cereal-service config-mgmt-service event-feed-service nodemanager-service pg-sidecar-service
ingest-service OPTIONAL es-sidecar-service
ingest-service BINDING_MODE strict
license-control-service REQUIRED automate-pg-gateway pg-sidecar-service
license-control-service BINDING_MODE strict
local-user-service REQUIRED authz-service automate-dex teams-service
local-user-service BINDING_MODE strict
nodemanager-service REQUIRED authz-service automate-pg-gateway cereal-service event-service pg-sidecar-service secrets-service
nodemanager-service BINDING_MODE strict
notifications-service REQUIRED automate-pg-gateway pg-sidecar-service secrets-service
notifications-service BINDING_MODE strict
pg-sidecar-service REQUIRED automate-pg-gateway
pg-sidecar-service BINDING_MODE strict
sample-data-service BINDING_MODE strict
secrets-service REQUIRED automate-pg-gateway pg-sidecar-service
secrets-service BINDING_MODE strict
session-service REQUIRED automate-dex automate-pg-gateway pg-sidecar-service
session-service BINDING_MODE strict
teams-service REQUIRED authz-service automate-pg-gateway cereal-service pg-sidecar-service
teams-service BINDING_MODE strict
trial-license-service BINDING_MODE strict
user-settings-service REQUIRED automate-pg-gateway pg-sidecar-service
user-settings-service BINDING_MODE strict
`
