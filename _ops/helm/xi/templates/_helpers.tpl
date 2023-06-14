## loadbalancer.ip deployment
{{- define "loadbalancer.ip" -}}
{{- if .Loadbalancer }}
{{- if .loadbalancer.ip }}
loadBalancerIP: {{ .loadbalancer.ip }}
{{- end }}
{{- end }}
{{- end -}}

## release.name state
{{- define "release.name" -}}
xi-{{ .Release.Name }}
{{- end -}}

## release.name.sql state
{{- define "release.name.sql" -}}
{{ include "release.name" . }}-sql
{{- end -}}
