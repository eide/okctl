// Package tempo provides a Helm chart for installing:
// - https://github.com/grafana/helm-charts/tree/main/charts/tempo
// - https://grafana.com/oss/tempo/
package tempo

import (
	"bytes"
	"text/template"

	"github.com/oslokommune/okctl/pkg/config"

	"github.com/oslokommune/okctl/pkg/helm"
)

// New returns an initialised Helm chart for installing cluster-tempo
func New(values *Values) *helm.Chart {
	return &helm.Chart{
		RepositoryName: "grafana",
		RepositoryURL:  "https://grafana.github.io/helm-charts",
		ReleaseName:    "tempo",
		Version:        "0.6.3",
		Chart:          "tempo",
		Namespace:      "monitoring",
		Timeout:        config.DefaultChartApplyTimeout,
		Values:         values,
	}
}

// NewDefaultValues returns the mapped values.yml containing
// the default values
func NewDefaultValues() *Values {
	return &Values{}
}

// Values contains the required inputs for generating the values.yml
// One of those cases where there really isn't much to change, but
// I will leave these structures here nonetheless.
// We should configuring and setting up a S3 storage backend:
// - https://github.com/grafana/tempo/blob/master/docs/tempo/website/configuration/s3.md
type Values struct{}

// RawYAML implements the raw marshaller interface in the Helm package
func (v *Values) RawYAML() ([]byte, error) {
	tmpl, err := template.New("values").Parse(valuesTemplate)
	if err != nil {
		return nil, err
	}

	var buff bytes.Buffer

	err = tmpl.Execute(&buff, *v)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// nolint: lll
const valuesTemplate = `# -- Overrides the chart's name
nameOverride: ""

# -- Overrides the chart's computed fullname
fullnameOverride: ""

replicas: 1
tempo:
  repository: grafana/tempo
  tag: 0.6.0
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ##
  # pullSecrets:
  #   - myRegistryKeySecretName

  updateStrategy: RollingUpdate
  resources: {}
  #  requests:
  #    cpu: 1000m
  #    memory: 4Gi
  #  limits:
  #    cpu: 2000m
  #    memory: 6Gi

  memBallastSizeMbs: 1024
  authEnabled: false
  ingester: {}
  retention: 24h
  server:
    httpListenPort: 3100
  # tempo storage backend
  # refer https://github.com/grafana/tempo/tree/master/docs/tempo/website/configuration
  ## Use s3 for example
  # backend: s3                                         # store traces in s3
  #  s3:
  #    bucket: tempo                                   # store traces in this bucket
  #    endpoint: s3.dualstack.us-east-2.amazonaws.com  # api endpoint
  #    access_key: ...                                 # optional. access key when using static credentials.
  #    secret_key: ...                                 # optional. secret key when using static credentials.
  #    insecure: false                                 # optional. enable if endpoint is http
  ## end
  storage:
    trace:
      backend: local
      local:
        path: /tmp/tempo/traces
      wal:
        path: /var/tempo/wal
  # this configuration will listen on all ports and protocols that tempo is capable of.
  # the receives all come from the OpenTelemetry collector.  more configuration information can
  # be found there: https://github.com/open-telemetry/opentelemetry-collector/tree/master/receiver
  receivers:
    jaeger:
      protocols:
        grpc:
          endpoint: 0.0.0.0:14250
        thrift_binary:
          endpoint: 0.0.0.0:6832
        thrift_compact:
          endpoint: 0.0.0.0:6831
        thrift_http:
          endpoint: 0.0.0.0:14268
  ## Additional container arguments
  extraArgs: {}
  # -- Environment variables to add
  extraEnv: []
  # -- Volume mounts to add
  extraVolumeMounts: []
  # - name: extra-volume
  #   mountPath: /mnt/volume
  #   readOnly: true
  #   existingClaim: volume-claim

tempoQuery:
  repository: grafana/tempo-query
  tag: latest
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ##
  # pullSecrets:
  #   - myRegistryKeySecretName
  ## Additional container arguments
  extraArgs: {}
  # -- Environment variables to add
  extraEnv: []
  # -- Volume mounts to add
  extraVolumeMounts: []
  # - name: extra-volume
  #   mountPath: /mnt/volume
  #   readOnly: true
  #   existingClaim: volume-claim

serviceAccount:
  # -- Specifies whether a ServiceAccount should be created
  create: true
  # -- The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name: null
  # -- Image pull secrets for the service account
  imagePullSecrets: []
  # -- Annotations for the service account
  annotations: {}

service:
  type: ClusterIP
  annotations: {}
  labels: {}

persistence:
  enabled: false
  # storageClassName: local-path
  accessModes:
    - ReadWriteOnce
  size: 10Gi

## Pod Annotations
podAnnotations: []

# -- Volumes to add
extraVolumes: []

## Node labels for pod assignment
## ref: https://kubernetes.io/docs/user-guide/node-selection/
#
nodeSelector: {}

## Tolerations for pod assignment
## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
##
tolerations: []

## Affinity for pod assignment
## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
##
affinity: {}
`
