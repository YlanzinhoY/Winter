package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	dependencyorganization "github.com/spring_initializr/ylanzinhoy/internals/dependencyOrganization"
	"github.com/spring_initializr/ylanzinhoy/internals/services"
)

func ApiExecuteController() {
	skipMessage := "if you dont need this, please press enter to skip!"
	springVersion := "3.3.5"

	packageManager := inputs(dependencyorganization.PackageManager())
	language := inputs(dependencyorganization.Langueges())

	javaVersion := inputs(dependencyorganization.JavaVersion())

	fmt.Println("Name: ")
	var name string
	fmt.Scan(&name)

	fmt.Println("Group: (EX: com.github)")
	var groupId string
	fmt.Scan(&groupId)

	fmt.Println("description: ")
	var description string = " "
	fmt.Scan(&description)

	packageName := groupId + "." + name
	artifactId := name

	var dependencies []string

	developerTools := dependenciesInput(dependencyorganization.DeveloperTools(), "developer tools "+skipMessage)
	web := dependenciesInput(dependencyorganization.Web(), "Web "+skipMessage)
	sql := dependenciesInput(dependencyorganization.Sql(), "SQL "+skipMessage)
	security := dependenciesInput(dependencyorganization.Security(), "Security "+skipMessage)
	noSql := dependenciesInput(dependencyorganization.Sql(), "NoSql "+skipMessage)
	io := dependenciesInput(dependencyorganization.Io(), "i/o")

	dependencies = append(append(append(append(append(append([]string{}, io...), developerTools...), web...), sql...), security...), noSql...)
	dependenciesFormatter := strings.Join(dependencies, ",")

	fmt.Println(dependenciesFormatter)

	mountUrl2 := "https://start.spring.io/starter.zip?type=" + packageManager + "&language=" + language + "&platformVersion=" + springVersion + "&groupId=" + groupId + "&artifactId" + artifactId + "&name=" + name + "&description=" + description + "%20project%20foe%20Spring%Boot" + "&packageName=" + packageName + "&packaging=jar&javaVersion=" + javaVersion + "&dependencies=" + dependenciesFormatter

	fmt.Println(mountUrl2)

	generateSpringProject(mountUrl2, name)
	buildingScaffold(name, language, packageName)

}

func dependenciesInput(dependencies map[string]string, topicName string) []string {
	var keys []string
	var values []string

	for k, v := range dependencies {

		keys = append(keys, k)
		values = append(values, v)
	}

	choicesDependencies := dependencyorganization.NewdependencyStructureOrganization().ChoiceDependencies(keys, values, topicName)

	var selectedValues []string

	for _, choice := range choicesDependencies {
		if val, ok := dependencies[choice]; ok {
			selectedValues = append(selectedValues, val)
		}
	}

	return selectedValues
}

func inputs(dependencies map[string]string) string {

	var keys []string
	var values []string

	for k, v := range dependencies {

		keys = append(keys, k)
		values = append(values, v)
	}
	catchChoice := dependencyorganization.NewdependencyStructureOrganization().ChoiceOne(keys)

	for i, keysValue := range keys {

		if catchChoice == keysValue {
			catchChoice = values[i]
		}
	}

	return catchChoice

}

func generateSpringProject(mountUrl string, name string) {
	zipName := name + ".zip"
	res, err := http.Get(mountUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res.StatusCode)

	defer res.Body.Close()

	out, err := os.Create(zipName)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer out.Close()

	_, err = io.Copy(out, res.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	services.ExtractZip(zipName, name)
}

func buildingScaffold(DirName string, lang string, groupId string) {

	groupPath := strings.ReplaceAll(groupId, ".", string(os.PathSeparator))

	path := filepath.Join(DirName, "src", "main", lang, groupPath)

	scaffold := services.Scaffold{
		Path: path,
		Lang: lang,
	}

	scaffold.Scaffold()

}
