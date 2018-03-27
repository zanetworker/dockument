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

## Tests

## Command Tests
{{ $length := len .CommandTests }} {{ if not (eq $length 0) }}
Below are your container application dependencies: 

{{- range $key, $commandTest := .CommandTests }}

### Command Test {{.Name}}

{{ $length := len .Name }} {{ if not (eq $length 0) }}- **Name**: {{.Name}}{{ end}}
{{ $length := len .Setup }} {{ if not (eq $length 0) }}- **Setup**: {{.Setup}} {{end}}
{{ $length := len .Command }} {{ if not (eq $length 0) }}- **Command**: {{.Command}}{{end}}
{{ $length := len .Args }} {{ if not (eq $length 0) }}- **Args**: {{.Args}}{{end}}
{{ $length := len .ExpectedOutput }} {{ if not (eq $length 0) }}- **Expected Output**: {{.ExpectedOutput}}{{end}}

{{end}}
{{end}}

## File Content Tests
{{ $length := len .FileContentTests }} {{ if not (eq $length 0) }}
Below are your container application dependencies: 

{{- range $key, $fileContentTest := .FileContentTests }}

### File Content Test {{.Name}}

{{ $length := len .Name }} {{ if not (eq $length 0) }}- **Name**: {{.Name}}{{ end}}
{{ $length := len .Path }} {{ if not (eq $length 0) }}- **Path**: {{.Path}} {{end}}
{{ $length := len .ExpectedContents }} {{ if not (eq $length 0) }}- **Expected Contents**: {{.ExpectedContents}}{{end}}
{{ $length := len .ExcludedContents }} {{ if not (eq $length 0) }}- **Excluded Content**: {{.ExcludedContents}}{{end}}
{{end}}
{{end}}

## File Existence Tests
{{ $length := len .FileExistenceTests }} {{ if not (eq $length 0) }}
Below are your container application dependencies: 

{{- range $key, $fileExistenceTest := .FileExistenceTests }}

### File Existence Test {{.Name}}

{{ $length := len .Name }} {{ if not (eq $length 0) }}- **Name**: {{.Name}}{{ end}}
{{ $length := len .Path }} {{ if not (eq $length 0) }}- **Path**: {{.Path}} {{end}}
{{ $length := len .ShouldExist }} {{ if not (eq $length 0) }}- **Should Exist**: {{.ShouldExist}}{{end}}
{{ $length := len .Permissions }} {{ if not (eq $length 0) }}- **With Permissions**: {{.Permissions}}{{end}}
{{end}}
{{end}}


## Meta Data Tests
{{ $length := len .MetadataTests }} {{ if not (eq $length 0) }}
Below are your container application dependencies: 

{{- range $key, $metaDataTests := .MetadataTests }}

### Meta Data Test
{{ $length := len .Env }} {{ if not (eq $length 0) }}- **Environment Variables**: {{.Env}}{{end}}
{{ $length := len .ExposedPorts }} {{ if not (eq $length 0) }}- **Exposed Ports**: {{.ExposedPorts}}{{end}}
{{ $length := len .EntryPoint }} {{ if not (eq $length 0) }}- **Entry Point**: {{.EntryPoint}}{{end}}
{{ $length := len .Cmd }} {{ if not (eq $length 0) }}- **Cmd**: {{.Cmd}}{{end}}
{{ $length := len .Workdir }} {{ if not (eq $length 0) }}- **Workdir**: {{.Workdir}}{{end}}
{{ $length := len .Volumes }} {{ if not (eq $length 0) }}- **Volumes**: {{.Volumes}}{{end}}


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
