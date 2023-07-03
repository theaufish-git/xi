## container.env.default
{{- define "container.env.default" -}}
- name: XI_SQL_HOST
  valueFrom:
    fieldRef:
      fieldPath: status.podIP
- name: POD_IP
  valueFrom:
    fieldRef:
      fieldPath: status.podIP
- name: POD_NAME
  valueFrom:
    fieldRef:
      fieldPath: metadata.name
- name: NODE_NAME
  valueFrom:
    fieldRef:
      fieldPath: spec.nodeName
{{- end -}}

## container.env.secrets container
{{- define "container.env.secrets" -}}
{{- range $secret := . }}
{{- range $key, $var := $secret.values }}
- name: {{ if $secret.prefix }}{{ upper $secret.prefix }}_{{ end }}{{ upper $key }}
  valueFrom:
    secretKeyRef:
      name: {{ $secret.from }}
      key: {{ $var }}
{{- end }}
{{- end }}
{{- end -}}

## container.ports $port
{{- define "container.port" -}}
- containerPort: {{ coalesce .targetPort .port }}
{{- end -}}

## container.ports $ports
{{- define "container.ports" -}}
{{- range $_, $port := . }}
- containerPort: {{ coalesce $port.targetPort $port.port }}
{{- end }}
{{- end -}}
