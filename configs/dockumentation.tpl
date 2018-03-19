{{define "dockument"}}
#  DOCKumentation 

Your application Dockerfile is DOCKumented below. 

## Exposed Ports
{{ $length := len .Ports }} {{ if not (eq $length 0) }}

Below are the ports the should be exposed by your application: 

{{- range $key, $dependency := .Ports }}
### Port {{.Name}}

- **Name**: {{.Name}}
- **About**: {{.About}}
- **Scheme**: {{.Scheme}}
- **Protocol**: {{.Protocol}}
{{end}}
{{end}}
## Dependencies

{{ $length := len .Dependencies }} {{ if not (eq $length 0) }}
Below are your container application dependencies: 

{{- range $key, $dependency := .Dependencies }}

### Dependency {{.Name}}

- **Name**: {{.Name}}
- **Image**: {{.Image}} (`docker pull {{.Image}}`)
- **About**: {{.About}}
- **Ports**: {{.Ports}}
- **Mandatory**: {{.Mandatory}}
{{end}}
{{end}}


## Environment Variables
{{ $length := len .Envs }} {{ if not (eq $length 0) }}
Below are some important Environment Variables:

{{- range $key, $dependency := .Envs }}

### Env {{.Name}}

- **Name**: {{.Name}}
- **About**: {{.About}}
- **Mandatory**: {{.Mandatory}}
{{end}}
{{end}}


## Resources Required
{{ $cpuLength  := len .Resources.CPU }} {{ if not (eq $cpuLength 0)}} 
{{ $memoryLength  := len .Resources.Memory }} {{ if not (eq $memoryLength 0)}} 
Resources required by your container application to run with reasonable performance: 
- **CPU**: {{.Resources.CPU}}
- **Memory**: {{.Resources.Memory}}
{{end}}
{{end}}


## Tags / Metadata
{{ $length  := len .Tags }}{{ if not (eq $length 0)}}
Metadata and container tags: 
{{- range $key, $val := .Tags}}
- **{{ $key }}**: {{ $val}}
{{end}}
{{end}}

## Misc Labels
{{ $length  := len .Others }}{{ if not (eq $length 0)}}
Non-Standard DOCKument API labels: 
{{- range $key, $val := .Others}}
- **{{ $key }}**: {{ $val}}
{{- end}}
{{end}}


{{end}}
