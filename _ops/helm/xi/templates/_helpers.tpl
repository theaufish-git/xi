## release.name state
{{- define "release.name" -}}
xi-{{ .Release.Name }}
{{- end -}}

## release.name.sql state
{{- define "release.name.sql" -}}
{{ include "release.name" . }}-sql
{{- end -}}

## sql.port
{{- define "sql.port" -}}
{{ with (index . 0) -}}{{- coalesce .targetPort .port -}}{{- end }}
{{- end -}}