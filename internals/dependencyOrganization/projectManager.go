package dependencyorganization

func PackageManager() map[string]string {
	return map[string]string{
		"Gradle - Groovy": "gradle-project",
		"Gradle - Kotlin": "gradle-project-kotlin",
		"Maven":           "maven-project",
	}
}
