## service.loadBalancerIP deployment
{{- define "service.loadBalancerIP" -}}
{{ if .loadBalancerIP -}}
loadBalancerIP: {{ .loadBalancerIP }}
{{- end }}
{{- end -}}

## service.ports deployment state
{{- define "service.port" -}}
{{- $ingress := coalesce .port .targetPort -}}
{{- $portName := printf "%s-%v" (lower .protocol) $ingress -}}
- name: {{ coalesce .name $portName }}
  port: {{ coalesce $ingress }}
  targetPort: {{ coalesce .targetPort .port }}

  {{ if .protocol }}protocol: {{ .protocol }}{{ end }}
{{- end -}}

## service.ports deployment state
{{- define "service.ports" -}}
{{- range $_, $port := . }}
{{ include "service.port" $port }}
{{- end }}
{{- end -}}