# Boilerplate nestjs-gql-rest-advanced

✨ includes

- Graphql + RestAPI support
- SOLID principle style coding orientations
- Abstracted repository CRUD with database connection to perform atomic transactions
- Authentication with passport via JWT/Local strategy setup
- Kafka setup with dead letter queue in MongoDB to handle missed/not-consumed message topics
- Dockerized application with multi-stage build setup
- kubectl(k8s), minikube, helm (deployment packet manager) setup
- Logging-Interceptor setup application level
- prometheus scraper to scrape application logs from logging-interceptor
- grafana+minikube+helm+promotheus setup for application metrics and monitoring reff: deployment.yaml
- keda setup for event driven autoscaling setup wired with k8s & minikube pods & running local cluster inspecting on prometheus reff: scaledobject.yaml
