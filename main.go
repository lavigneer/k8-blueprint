package main

import (
	_ "embed"
	"io/fs"
	"log"
	"os"
	"slices"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	sprig "github.com/go-task/slim-sprig/v3"
)

//go:embed templates/Chart.tmpl.yaml
var chartTemplate string

//go:embed templates/Makefile.tmpl
var makefileTemplate string

//go:embed templates/jenkins/values.tmpl.yaml
var jenkinsValuesTemplate string

//go:embed templates/prometheus/values.tmpl.yaml
var prometheusValuesTemplate string

//go:embed templates/loki/values.tmpl.yaml
var lokiValuesTemplate string

type FormData struct {
	dirFs fs.FS
	Apps  []string
}

func writeMainChart(formData FormData) {
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(chartTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("fake-repo/Chart.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, formData)
	if err != nil {
		log.Fatal(err)
	}
}
func writeMakefile(formData FormData) {
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(makefileTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("fake-repo/Makefile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, formData)
	if err != nil {
		log.Fatal(err)
	}
}

func writeJenkinsValues(formData FormData) {
	if !slices.Contains(formData.Apps, "jenkins") {
		return
	}
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(jenkinsValuesTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("fake-repo/jenkins.values.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, formData)
	if err != nil {
		log.Fatal(err)
	}
}

func writePrometheusValues(formData FormData) {
	if !slices.Contains(formData.Apps, "prometheus") {
		return
	}
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(prometheusValuesTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("fake-repo/prometheus.values.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, formData)
	if err != nil {
		log.Fatal(err)
	}
}

func writeLokiValues(formData FormData) {
	if !slices.Contains(formData.Apps, "loki") {
		return
	}
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(lokiValuesTemplate)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("fake-repo/loki.values.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, formData)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var formData FormData
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("What applications do you want to install?").
				Options(
					huh.NewOption("Jenkins", "jenkins"),
					huh.NewOption("Prometheus", "prometheus"),
					huh.NewOption("Loki", "loki"),
				).
				Value(&formData.Apps),
		),
	).WithProgramOptions(tea.WithAltScreen())
	formData.dirFs = os.DirFS("fake-repo")
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	writeMainChart(formData)
	writeMakefile(formData)
	writeJenkinsValues(formData)
	writePrometheusValues(formData)
	writeLokiValues(formData)
}
