package main

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	chartPath := "/home/tsai/Desktop/golangProjects/goBackendApi/helmSDK/helmChartFiles/singleUserHelm"
	namespace := "user-resource"
	releaseName := "fafe2a0b-2b2d-46df-b24a-4bdc4d1b16cc"
	vals := map[string]interface{}{
		"jupyter": map[string]interface{}{
			"user": releaseName,
		},
		"grafana": map[string]interface{}{
			"podLabels": map[string]interface{}{
				"user": releaseName,
			},
		},
		"global": map[string]interface{}{
			"user": releaseName,
		},
	}

	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace,
		os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
	chart, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}

	client := action.NewInstall(actionConfig)
	client.Namespace = namespace
	client.ReleaseName = releaseName
	rel, err := client.Run(chart, vals)
	if err != nil {
		panic(err)
	}

	log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
	// this will confirm the values set during installation
	log.Println(rel.Config)
}
