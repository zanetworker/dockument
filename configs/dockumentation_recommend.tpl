{{define "dockument_recommendation"}}
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
{{else}}
No ports to show, please insert port labels in using the right syntax, for example: 

```dockerfile
LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.scheme="tcp" \
	  api.EXPOSE.1337.protocol="http"\
      api.EXPOSE.1337.about="The main endpoint of this service."
```
{{ end }}
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
{{else}}
No dependencies to show, please insert dependency labels in using the right syntax, for example: 

```dockerfile
LABEL api.DEPENDENCY.rabbit=""\
      api.DEPENDENCY.rabbit.image="rabbit:latest"\
      api.DEPENDENCY.rabbit.port="5271"\
      api.DEPENDENCY.rabbit.about="The rabbit"\
      api.DEPENDENCY.rabbit.mandatory="true"
```
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
{{else}}
No dependencies to show, please insert dependency labels in using the right syntax, for example: 

```dockerfile
LABEL api.ENV.OPENWEATHERMAP_APIKEY="" \
      api.ENV.OPENWEATHERMAP_APIKEY.about="Access key for OpenWeatherMap. See http://openweathermap.org/appid for details." \
      api.ENV.OPENWEATHERMAP_APIKEY.mandatory="true"
```
{{end}}


## Resources Required
{{ $cpuLength  := len .Resources.CPU }} {{ if not (eq $cpuLength 0)}} 
{{ $memoryLength  := len .Resources.Memory }} {{ if not (eq $memoryLength 0)}} 
Resources required by your container application to run with reasonable performance: 
- **CPU**: {{.Resources.CPU}}
- **Memory**: {{.Resources.Memory}}
{{end}}
{{else}}
No resources to show, please insert resources labels using the right syntax, for example: 


```dockerfile
LABEL api.RESOURCES.Memory="3gb"\
      api.RESOURCES.CPU="2"
```
{{end}}


## Tags / Metadata
{{ $length  := len .Tags }}{{ if not (eq $length 0)}}
Metadata and container tags: 
{{- range $key, $val := .Tags}}
- **{{ $key }}**: {{ $val}}
{{end}}
{{else}}
No tags / metadata were defined in your Dockerfile, please insert tags using the right syntax, for example: 
```dockerfile
LABEL api.TAGS.go="1.9"
```
{{end}}


{{end}}
