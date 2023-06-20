package global

{{- if .HasGlobal }}

import "github.com/pemako/gva/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}