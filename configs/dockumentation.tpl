{{define "dockument"}}
#  DOCKumentation 

Your application Dockerfile is DOCKumented below. 

## Exposed Ports

Below are the ports the should be exposed by your application: 

{{- range $key, $dependency := .Ports }}
### Port {{.Name}}

- **Name**: {{.Name}}
- **About**: {{.About}}
- **Scheme**: {{.Scheme}}
- **Protocol**: {{.Protocol}}

{{end}}

## Dependencies

Below are your container application dependencies: 

{{- range $key, $dependency := .Dependencies }}

### Dependency {{.Name}}

- **Name**: {{.Name}}
- **Image**: {{.Image}} (`docker pull {{.Image}}`)
- **About**: {{.About}}
- **Ports**: {{.Ports}}
- **Mandatory**: {{.Mandatory}}

{{end}}

## Environment Variables

Below are some important Environment Variables:

{{- range $key, $dependency := .Envs }}

### Env {{.Name}}

- **Name**: {{.Name}}
- **About**: {{.About}}
- **Mandatory**: {{.Mandatory}}

{{end}}

## Resources Required

Resources required by your container application to run with reasonable performance: 


## Metadata & Extra Information

Container application tags and metadata: 


{{end}}