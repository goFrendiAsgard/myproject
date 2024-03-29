package gen

import (
	"fmt"
	"github.com/state-alchemists/ayanami/generator"
	"log"
	"path/filepath"
	"strings"
)

// ExposedGomonolithProc GomonolithProc for template
type ExposedGomonolithProc struct {
	Functions []string
}

// GoMonolithProc procedureuration to generate GoMonolith
type GoMonolithProc struct {
	DirName  string
	RepoName string
	generator.IOHelper
}

// Validate validating procedure
func (p GoMonolithProc) Validate(configs generator.Configs) bool {
	return true
}

// Scaffold scaffolding procedure
func (p GoMonolithProc) Scaffold(configs generator.Configs) error {
	return nil
}

// Build building procedure
func (p GoMonolithProc) Build(configs generator.Configs) error {
	log.Printf("[INFO] BUILDING MONOLITH: %s", p.DirName)
	depPath := p.DirName
	var mainFunctionList []string
	for _, config := range configs {
		switch config.(type) {
		case CmdConfig:
			c := config.(CmdConfig)
			mainFunctionName := fmt.Sprintf("Service%s", strings.Title(c.ServiceName))
			mainFunctionList = append(mainFunctionList, mainFunctionName)
			err := c.CreateProgram(depPath, p.RepoName, mainFunctionName)
			if err != nil {
				return err
			}
		case GoServiceConfig:
			c := config.(GoServiceConfig)
			mainFunctionName := fmt.Sprintf("Service%s", strings.Title(c.ServiceName))
			mainFunctionList = append(mainFunctionList, mainFunctionName)
			err := c.CreateProgram(depPath, p.RepoName, mainFunctionName)
			if err != nil {
				return err
			}
		case GatewayConfig:
			c := config.(GatewayConfig)
			mainFunctionName := fmt.Sprintf("Gateway%s", strings.Title(c.ServiceName))
			mainFunctionList = append(mainFunctionList, mainFunctionName)
			err := c.CreateProgram(depPath, p.RepoName, mainFunctionName)
			if err != nil {
				return err
			}
		case FlowConfig:
			c := config.(FlowConfig)
			mainFunctionName := fmt.Sprintf("Flow%s", strings.Title(c.FlowName))
			mainFunctionList = append(mainFunctionList, mainFunctionName)
			err := c.CreateProgram(depPath, p.RepoName, mainFunctionName)
			if err != nil {
				return err
			}
		}
	}
	// write main.go
	data := ExposedGomonolithProc{Functions: mainFunctionList}
	mainPath := filepath.Join(depPath, "main.go")
	err := p.WriteDep(mainPath, "gomonolith.main.go", data)
	if err != nil {
		return err
	}
	// write common things
	for _, templateName := range []string{"go.mod", "Makefile", ".gitignore"} {
		log.Printf("[INFO] Create %s", templateName)
		goModPath := filepath.Join(depPath, templateName)
		err := p.WriteDep(goModPath, templateName, p)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewGoMonolith make monolithic app
func NewGoMonolith(g *generator.Generator, dirName, repoName string) GoMonolithProc {
	return GoMonolithProc{
		DirName:  dirName,
		RepoName: repoName,
		IOHelper: g.IOHelper,
	}
}
