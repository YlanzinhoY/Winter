package dependencyorganization

func DeveloperTools() map[string]string {

	return map[string]string{

		"GraalVm Native Support":         "native",
		"GraphQL DGS Code Generation":    "dgs-codegen",
		"Spring Boot DevTools":           "devtools",
		"Lombok":                         "lombok",
		"Spring Configuration Processor": "configuration-processor",
		"Docker Compose Support":         "docker-compose",
		"Spring Modulith":                "modulith",
	}

}

func Web() map[string]string {
	return map[string]string{
		"Spring Web":                     "web",
		"Spring Reactive Web":            "webflux",
		"Spring for GraphQL":             "graphql",
		"Spring Session":                 "session",
		"Rest Repositories HAL Explorer": "data-rest-explorer",
		"Spring HATEOAS":                 "hateoas",
		"Spring Web Services":            "web-services",
		"Jersey":                         "jersey",
		"Vaadin":                         "vaadin",
		"Netflix DGS":                    "netflix-dgs",
		"htmx":                           "htmx",
	}

}

func TemplateEngines() map[string]string {

	return map[string]string{

		"Thymeleaf":         "thymeleaf",
		"Apache Freemarker": "freemarker",
		"Mustache":          "mustache",
		"Groovy Templates":  "groovy-templates",
		"JTE":               "jte",
	}

}

func Security() map[string]string {

	return map[string]string{
		"Spring Security":             "security",
		"OAuth2 Resource Server":      "oauth2-resource-server",
		"OAuth2 Authorization Server": "oauth2-authorization-server",
		"OAuthClient":                 "oauth2-client",
		"Spring LDAP":                 "data-ldap",
		"Okta":                        "okta",
	}

}

func Sql() map[string]string {

	return map[string]string{

		"JDBC":                  "jdbc",
		"Spring Data JPA":       "data-jpa",
		"Spring Data JDBC":      "data-jdbc",
		"Spring Data R2DBC":     "data-r2dbc",
		"MyBatis":               "mybatis",
		"Liquibase Migration":   "liquibase",
		"Flyway Migration":      "flyway",
		"JOOQ Access Layer":     "jooq",
		"IBM DB2 Driver":        "db2",
		"Apache Derby Database": "derby",
		"H2 Database":           "h2",
		"HyperSQL Database":     "hsql",
		"MariaDB Driver":        "mariadb",
		"MS SQL Server Driver":  "sqlserver",
		"MySQL":                 "mysql",
		"Oracle Driver":         "oracle",
		"PostgresSQL":           "postgresql",
	}

}

func NoSql() map[string]string {
	return map[string]string{
		"Spring Data Redis (Acecss+Driver)":         "data-redis",
		"Spring Data Reactive Redis":                "data-redis-reactive",
		"Spring Data MongoDB":                       "data-mongodb",
		"Spring Data Reactive MongoDB":              "data-mongodb-reactive",
		"Spring Data Elasticsearch(Access+Driver)":  "data-elasticsearch",
		"Spring Data for Apache Cassandra":          "data-cassandra",
		"Spring Data Reactive for Apache Cassandra": "data-cassandra-reactive",
		"Spring Data Couchbase":                     "data-couchbase",
		"Spring Data Reactive Couchbase":            "data-couchbase-reactive",
		"Spring Data Neo4j":                         "data-neo4j",
	}

}

func Messaging() map[string]string {
	return map[string]string{

		"Spring Integration":                  "integration",
		"Spring for RabbitMQ":                 "amqp",
		"Spring for RabbitMQ Streams":         "amqp-streams",
		"Spring for Apache Kafka":             "kafka",
		"Spring for Apache Kafka Streams":     "kafka-streams",
		"Spring for Apache ActiveMQ 5":        "activemq",
		"Spring for Apache ActiveMQ Artemis":  "artemis",
		"Spring for Apache Pulsar":            "pulsar",
		"Spring for Apache Pulsar (Reactive)": "pulsar-reactive",
		"WebSocket":                           "websocket",
		"RSocket":                             "rsocket",
		"Apache Camel":                        "camel",
		"Solace PubSub+":                      "solace",
	}

}

func Io() map[string]string {
	return map[string]string{

		"Spring Batch":             "batch",
		"Validation":               "validation",
		"Java Mail Sender":         "mail",
		"Quartz Scheduler":         "quartz",
		"Spring Cache Abstraction": "cache",
		"Spring Shell":             "spring-shell",
	}

}

func Ops() map[string]string {
	return map[string]string{
		"Spring Boot Actuator":                    "actuator",
		"CycloneDX SBOM support":                  "sbom-cyclone-dx",
		"codecentric's (Client)":                  "codecentric-spring-boot-admin-client",
		"codecentic's Spring Boot Admin (Server)": "codecentric-spring-boot-admin-server",
	}

}

func Observability() map[string]string {
	return map[string]string{

		"Datadog":             "datadog",
		"Dynatrace":           "dynatrace",
		"Influx":              "influx",
		"Graphite":            "graphite",
		"New Relic":           "new-relic",
		"OTLP for metrics":    "otlp-metrics",
		"Prometheus":          "prometheus",
		"Distributed Tracing": "distributed-tracing",
		"Wavefront":           "wavefront",
		"Zipkin":              "zipkin",
	}

}

func Testing() map[string]string {
	return map[string]string{

		"Spring Rest Docs":     "restdocs",
		"Testcontainers":       "testcontainers",
		"Contract Verifier":    "cloud-contract-verifier",
		"Contract Stub Runner": "cloud-contract-stub-runner",
		"Embedded LDAP Server": "unboundid-ldap",
	}

}

func SpringCloud() map[string]string {
	return map[string]string{
		"Cloud Bootstrap": "cloud-starter",
		"Function":        "cloud-function",
		"Task":            "cloud-task",
	}

}

func SpringCloudConfig() map[string]string {
	return map[string]string{

		"Config Clinet":                  "cloud-config-client",
		"Config Server":                  "cloud-config-server",
		"Vault Configuration":            "cloud-starter-vault-config",
		"Apache Zookeeper Configuration": "cloud-starter-zookeeper-config",
		"Consul Configuration":           "cloud-starter-consul-config",
	}

}

func SpringCloudDiscovery() map[string]string {
	return map[string]string{
		"Eureka Discovery Client":    "cloud-eureka",
		"Eureka Server":              "cloud-eureka-server",
		"Apache Zookeeper Descovery": "cloud-starter-zookeeper-discovery",
		"Consul":                     "cloud-starter-consul-config",
		"Discovery":                  "cloud-starter-consul-discovery",
	}
}

func SpringCloudRouting() map[string]string {
	return map[string]string{
		"Gateway":            "cloud-gateway",
		"Reactive Gateway":   "cloud-gateway-reactive",
		"OpenFeign":          "cloud-feign",
		"Cloud LoadBalancer": "cloud-loadbalancer",
	}
}

func SpringCloudCircuitBreaker() map[string]string {
	return map[string]string{
		"Resilience4J": "cloud-resilience4j",
	}
}

func SpringCloudMessaging() map[string]string {
	return map[string]string{
		"Cloud Bus":    "cloud-bus",
		"Cloud Stream": "cloud-stream",
	}
}

func VmWareTanzuApplicationService() map[string]string {
	return map[string]string{
		"Config Client (TAS)":    "scs-config-client",
		"Service Registry (TAS)": "scs-service-registry",
	}

}

func MicrosoftAzure() map[string]string {
	return map[string]string{
		"Azure Support":          "azure-support",
		"Azure Active Directory": "azure-active-directory",
		"Azure Cosmos DB":        "azure-cosmos-db",
		"Azure Key Vault":        "azure-keyvault",
		"Azure Storage":          "azure-storage",
	}

}

func GoogleCloud() map[string]string {

	return map[string]string{
		"Google Cloud Support":   "cloud-gcp",
		"Google Cloud Messaging": "cloud-gcp-pubsub",
		"Google Cloud Storage":   "cloud-gcp-storage",
	}

}

func Ai() map[string]string {
	return map[string]string{
		"Anthropic Claude":                       "spring-ai-anthropic,spring-ai-azure-openai",
		"Azure OpenAI":                           "spring-ai-vectordb-azure",
		"Azure AI Search":                        "spring-ai-bedrock",
		"Amazon Bedrock":                         "spring-ai-vectordb-cassandra",
		"Apache Cassandra Vector Database":       "spring-ai-vectordb-chroma",
		"Chroma Vector Database":                 "spring-ai-vectordb-elasticsearch",
		"Elasticsearch Vector Database":          "spring-ai-vectordb-milvus",
		"Milvus Vector Database":                 "spring-ai-mistral",
		"Mistral AI":                             "spring-ai-vectordb-mongodb-atlas",
		"MongoDB Atlas Vector Database":          "spring-ai-vectordb-neo4j",
		"Neo4j Vector Database":                  "spring-ai-ollama,spring-ai-openai",
		"Ollama":                                 "spring-ai-vectordb-oracle",
		"OpenAI":                                 "spring-ai-vectordb-pgvector",
		"PGvector Vector Database":               "spring-ai-vectordb-pinecone",
		"PostgresML":                             "spring-ai-postgresml",
		"Redis Search and Query Vector Database": "spring-ai-vectordb-redis",
		"Stability AI":                           "spring-ai-stabilityai",
		"Transformers (ONNX) Embeddings":         "spring-ai-transformers",
		"Vertex AI PaLM2":                        "spring-ai-vertexai-palm2",
		"Vertex AI Gemini":                       "spring-ai-vertexai-gemini",
		"Vertex AI Embeddings":                   "spring-ai-vertexai-embeddings",
		"Qdrant Vector Database":                 "spring-ai-vectordb-qdrant",
		"Typesense Vector Database":              "spring-ai-vectordb-typesense",
		"Weaviate Vector Database":               "spring-ai-vectordb-weaviate",
		"Markdown Document Reader":               "spring-ai-markdown-document-reader",
		"Tika Document Reader":                   "spring-ai-tika-document-reader",
		"PDF Document Reader":                    "spring-ai-pdf-document-reader",
		"Timefold Solver":                        "timefold-solver",
	}

}
