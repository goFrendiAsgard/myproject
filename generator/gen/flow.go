package gen

import (
	"fmt"
	"github.com/state-alchemists/ayanami/generator"
	"log"
	"path/filepath"
	"strings"
)

// ExposedFlowConfig exposed ready flowConfig
type ExposedFlowConfig struct {
	MainFunctionName string
	RepoName         string
	FlowName         string
	ServiceName      string // alias to FlowName
	Packages         []string
	Events           []map[string]string
	Outputs          string
	Inputs           string
}

// FlowConfig definition
type FlowConfig struct {
	RepoName string
	FlowName string
	Inputs   []string
	Outputs  []string
	Events   []Event
	generator.IOHelper
	generator.StringHelper
}

// Validate validating config
func (c FlowConfig) Validate() bool {
	log.Printf("[INFO] VALIDATING FLOW: %s", c.FlowName)
	for _, input := range c.Inputs {
		if !c.IsMatch(input, "^[A-Za-z][a-zA-Z0-9]*$") {
			log.Printf("[ERROR] Invalid input `%s`", input)
			return false
		}
	}
	for _, output := range c.Outputs {
		if !c.IsMatch(output, "^[A-Za-z][a-zA-Z0-9]*$") {
			log.Printf("[ERROR] Invalid output `%s`", output)
			return false
		}
	}
	if !c.IsAlphaNumeric(c.FlowName) {
		log.Printf("[ERROR] Flow name should be alphanumeric, but `%s` found", c.FlowName)
		return false
	}
	if c.RepoName == "" {
		log.Printf("[ERROR] Repo name should not be empty")
		return false
	}
	for index, event := range c.Events {
		log.Printf("[INFO] Validating event %d: %v", index, event.ToMap())
		if !event.Validate() {
			return false
		}
	}
	return true
}

// Scaffold scaffolding config
func (c FlowConfig) Scaffold() error {
	log.Printf("[INFO] SCAFFOLDING FLOW: %s", c.FlowName)
	for _, event := range c.Events {
		if !event.UseFunction {
			continue
		}
		data := map[string]string{
			"FunctionPackage": event.FunctionPackage,
			"FunctionName":    event.FunctionName,
		}
		packageSourcePath := event.FunctionPackage
		functionFileName := event.GetFunctionFileName()
		// write function
		functionSourcePath := filepath.Join(packageSourcePath, functionFileName)
		if !c.IsSourceExists(functionSourcePath) {
			log.Printf("[INFO] Create %s", functionFileName)
			err := c.WriteSource(functionSourcePath, "flow.function.go", data)
			if err != nil {
				return err
			}
		} else {
			log.Printf("[INFO] %s exists", functionFileName)
		}
		// write dependencies
		for _, dependency := range event.FunctionDependencies {
			dependencySourcePath := filepath.Join(packageSourcePath, dependency)
			if !c.IsSourceExists(dependencySourcePath) {
				log.Printf("[INFO] Create %s", dependency)
				err := c.WriteSource(dependencySourcePath, "dependency.go", data)
				if err != nil {
					return err
				}
			} else {
				log.Printf("[INFO] %s exists", dependency)
			}
		}
	}
	return nil
}

// Build building config
func (c FlowConfig) Build() error {
	log.Printf("[INFO] BUILDING FLOW: %s", c.FlowName)
	depPath := fmt.Sprintf("flow-%s", c.FlowName)
	repoName := c.RepoName
	mainFunctionName := "main"
	// create program
	err := c.CreateProgram(depPath, repoName, mainFunctionName)
	if err != nil {
		return err
	}
	// write common things
	for _, templateName := range []string{"go.mod", "Makefile", ".gitignore"} {
		log.Printf("[INFO] Create %s", templateName)
		goModPath := filepath.Join(depPath, templateName)
		err := c.WriteDep(goModPath, templateName, c)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateProgram create main.go and others
func (c FlowConfig) CreateProgram(depPath, repoName, mainFunctionName string) error {
	// write functions and dependencies
	for _, event := range c.Events {
		if !event.UseFunction {
			continue
		}
		packageSourcePath := event.FunctionPackage
		packageDepPath := filepath.Join(depPath, event.FunctionPackage)
		functionFileName := event.GetFunctionFileName()
		// copy function
		log.Printf("[INFO] Create %s", functionFileName)
		functionSourcePath := filepath.Join(packageSourcePath, functionFileName)
		functionDepPath := filepath.Join(packageDepPath, functionFileName)
		err := c.CopySourceToDep(functionSourcePath, functionDepPath)
		if err != nil {
			return err
		}
		// copy dependencies
		for _, dependency := range event.FunctionDependencies {
			log.Printf("[INFO] Create %s", dependency)
			dependencySourcePath := filepath.Join(packageSourcePath, dependency)
			dependencyDepPath := filepath.Join(packageDepPath, dependency)
			err := c.CopySourceToDep(dependencySourcePath, dependencyDepPath)
			if err != nil {
				return err
			}
		}
	}
	// write main file
	mainFileName := fmt.Sprintf("%s.go", strings.ToLower(mainFunctionName))
	log.Printf("[INFO] Create %s", mainFileName)
	mainPath := filepath.Join(depPath, mainFileName)
	err := c.WriteDep(mainPath, "flow.main.go", c.toExposed(repoName, mainFunctionName))
	return err
}

// AppendEvent add new Event object
func (c *FlowConfig) AppendEvent(event Event) {
	c.Events = append(c.Events, event)
}

// AddEvent add new Event
func (c *FlowConfig) AddEvent(inputEventName, outputEventName, varName string) {
	c.AppendEvent(NewEvent(inputEventName, outputEventName, varName))
}

// AddEventVal add new Event with value
func (c *FlowConfig) AddEventVal(inputEventName, outputEventName, varName string, value interface{}) {
	c.AppendEvent(NewEventVal(inputEventName, outputEventName, varName, value))
}

// AddEventFunc add new Event with function
func (c *FlowConfig) AddEventFunc(inputEventName, outputEventName, varName, functionPackage, functionName string, functionDependencies []string) {
	c.AppendEvent(NewEventFunc(inputEventName, outputEventName, varName, functionPackage, functionName, functionDependencies))
}

// AddInputEvent add new InputEvent
func (c *FlowConfig) AddInputEvent(eventName, varName string) {
	c.AppendEvent(NewInputEvent(eventName, varName))
}

// AddInputEventVal add new InputEvent with value
func (c *FlowConfig) AddInputEventVal(eventName, varName string, value interface{}) {
	c.AppendEvent(NewInputEventVal(eventName, varName, value))
}

// AddInputEventFunc add new InputEvent with function
func (c *FlowConfig) AddInputEventFunc(eventName, varName, functionPackage, functionName string, functionDependencies []string) {
	c.AppendEvent(NewInputEventFunc(eventName, varName, functionPackage, functionName, functionDependencies))
}

// AddOutputEvent create new OutputEvent
func (c *FlowConfig) AddOutputEvent(eventName, varName string) {
	c.AppendEvent(NewOutputEvent(eventName, varName))
}

// AddOutputEventVal create new OutputEvent with value
func (c *FlowConfig) AddOutputEventVal(eventName, varName string, value interface{}) {
	c.AppendEvent(NewOutputEventVal(eventName, varName, value))
}

// AddOutputEventFunc create new OutputEvent with function
func (c *FlowConfig) AddOutputEventFunc(eventName, varName, functionPackage, functionName string, functionDependencies []string) {
	c.AppendEvent(NewOutputEventFunc(eventName, varName, functionPackage, functionName, functionDependencies))
}

func (c *FlowConfig) toExposed(repoName, mainFunctionName string) ExposedFlowConfig {
	return ExposedFlowConfig{
		RepoName:         repoName,
		MainFunctionName: mainFunctionName,
		FlowName:         c.FlowName,
		ServiceName:      c.FlowName,
		Packages:         c.getPackagesForExposed(),
		Events:           c.getEventsForExposed(),
		Outputs:          c.QuoteArrayAndJoin(c.Outputs, ", "),
		Inputs:           c.QuoteArrayAndJoin(c.Inputs, ", "),
	}
}

func (c *FlowConfig) getEventsForExposed() []map[string]string {
	var events []map[string]string
	for _, event := range c.Events {
		events = append(events, event.ToIndentedMap())
	}
	return events
}

func (c *FlowConfig) getPackagesForExposed() []string {
	var packages []string
	for _, event := range c.Events {
		if !event.UseFunction {
			continue
		}
		packages = append(packages, event.FunctionPackage)
	}
	return packages
}

// NewFlow create new flow
func NewFlow(g *generator.Generator, repoName, flowName string, inputs, outputs []string, events []Event) FlowConfig {
	return FlowConfig{
		RepoName: repoName,
		FlowName: flowName,
		Inputs:   inputs,
		Outputs:  outputs,
		Events:   events,
		IOHelper: g.IOHelper,
	}
}

// NewEmptyFlow create new empty flow
func NewEmptyFlow(g *generator.Generator, repoName, flowName string, inputs, outputs []string) FlowConfig {
	return NewFlow(g, repoName, flowName, inputs, outputs, []Event{})
}
