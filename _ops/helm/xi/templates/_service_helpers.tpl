## service.ports deployment state
{{- define "service.port" -}}
- name: {{ (index . 0) }}
  {{- toYaml (index . 1) | nindent 2 }}
{{- end -}}