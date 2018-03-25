package labels

import (
	"strings"
)

func parseDependencies(dependencyLabels map[string]string) *Dependencies {
	dependencies := Dependencies{}
	for _, depMap := range divideMapByKey(dependencyLabels) {
		dependency := Dependency{}
		for dependencyLabel, value := range depMap {
			dependencyLabelStrings := strings.Split(dependencyLabel, ".")
			if len(dependencyLabelStrings) > 3 {
				dependencyParam := dependencyLabelStrings[3]
				switch dependencyParam {
				case "image":
					dependency.Image = value
				case "port":
					//TODO parse ports correctly
					dependency.Ports = []string{value}
				case "about":
					dependency.About = value
				case "mandatory":
					dependency.Mandatory = value
				}
			} else {
				dependency.Name = dependencyLabelStrings[2]
			}
		}
		dependencies = append(dependencies, dependency)
	}
	return &dependencies
}

func parseCommandTests(commandTestLabels map[string]string) *CommandTests {
	commandTests := CommandTests{}
	for _, commandTestMap := range divideMapByKey(commandTestLabels) {
		commandTest := CommandTest{}
		for commandTestsLabel, value := range commandTestMap {
			commandTestsLabelStrings := strings.Split(commandTestsLabel, ".")
			if len(commandTestsLabelStrings) > 3 {
				commandTestParam := commandTestsLabelStrings[3]
				switch commandTestParam {
				case "name":
					commandTest.Name = value
				case "command":
					commandTest.Command = value
				case "args":
					//TODO parse ports correctly
					commandTest.Args = []string{value}
				case "expectedOutput":
					commandTest.ExpectedOutput = value
				case "expectedError":
					commandTest.ExpectedError = value
				case "excludedOutput":
					commandTest.ExcludedOutput = value
				case "excludedError":
					commandTest.ExcludedError = value
				}
			}
		}
		commandTests = append(commandTests, commandTest)
	}
	return &commandTests
}

func parseFileExistenceTests(fileExistenceTestsLabels map[string]string) *FileExistenceTests {
	fileExistenceTests := FileExistenceTests{}
	for _, fileExistenceTestsMap := range divideMapByKey(fileExistenceTestsLabels) {
		fileExistenceTest := FileExistenceTest{}
		for fileExistenceTestsLabel, value := range fileExistenceTestsMap {
			fileExistenceTestsLabelStrings := strings.Split(fileExistenceTestsLabel, ".")
			if len(fileExistenceTestsLabelStrings) > 3 {
				fileExistenceTestParam := fileExistenceTestsLabelStrings[3]
				switch fileExistenceTestParam {
				case "name":
					fileExistenceTest.Name = value
				case "path":
					fileExistenceTest.Path = value
				case "shouldExist":
					fileExistenceTest.ShouldExist = value
				case "permissions":
					fileExistenceTest.Permissions = value
				}
			}
		}
		fileExistenceTests = append(fileExistenceTests, fileExistenceTest)
	}
	return &fileExistenceTests
}

func parseEnvs(envLabels map[string]string) *Envs {
	envs := Envs{}
	for _, envMap := range divideMapByKey(envLabels) {
		env := Env{}
		for envLabel, value := range envMap {
			envLabelStrings := strings.Split(envLabel, ".")
			if len(envLabelStrings) > 3 {
				envParam := envLabelStrings[3]
				switch envParam {
				case "about":
					env.About = value
				case "mandatory":
					env.Mandatory = value
				}
			} else {
				env.Name = envLabelStrings[2]
			}
		}
		envs = append(envs, env)
	}
	return &envs
}

func parsePorts(portLabels map[string]string) *Ports {
	ports := Ports{}
	for _, portMap := range divideMapByKey(portLabels) {
		port := Port{}
		for portLabel, value := range portMap {
			portLabelStrings := strings.Split(portLabel, ".")
			if len(portLabelStrings) > 3 {
				portParam := portLabelStrings[3]
				switch portParam {
				case "about":
					port.About = value
				case "scheme":
					port.Scheme = value
				case "protocol":
					port.Protocol = value
				}
			} else {
				port.Name = portLabelStrings[2]
			}
		}
		ports = append(ports, port)
	}
	return &ports
}
func parseResources(resourceLabels map[string]string) *Resources {
	resources := Resources{}
	for resourceLabel, value := range resourceLabels {
		resourceLabelString := strings.Split(resourceLabel, ".")
		resourceParam := resourceLabelString[2]

		switch resourceParam {
		case "CPU":
			resources.CPU = value
		case "Memory":
			resources.Memory = value
		}
	}
	return &resources
}

//Tags have dynamic keys
func parseTags(tagLabels map[string]string) *Tags {
	tags := Tags{}
	for tagLabel, value := range tagLabels {
		tagLabelSting := strings.Split(tagLabel, ".")
		tagParam := tagLabelSting[2]
		tags[tagParam] = value
	}
	return &tags
}
