{{define "dockument"}}

#  DOCKumentation 

Your application Dockerfile is DOCKumented below. 

## Exposed Ports

Below are the ports the should be exposed by your application: 

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

## Resources Required

Resources required by your container application to run with reasonable performance: 


## Metadata & Extra Information

Container application tags and metadata: 


{{end}}