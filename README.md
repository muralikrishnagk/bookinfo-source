# Bookinfo Microservice - Source Repository

This repository contains the source code and Helm charts for the Bookinfo microservice.

## Project Structure 
```
├── bookinfo-service/
│   ├── charts/
│   │   └── bookinfo/
│   │       ├── Chart.yaml
│   │       ├── values.yaml
│   │       └── templates/
│   │           └── deployment.yaml
│   └── src/
│       └── main.go
└── README.md
```

## Components
1. Source Code
   - Go-based microservice implementation
   - REST API endpoints
   - Business logic

2. Helm Charts
   - Bookinfo service Helm chart
   - Configuration templates
   - Values schema

## Usage
1. Clone the repository
2. Build the service
3. Package the Helm chart
4. Push to container registry

## Development
Run the service locally:
```bash
go run src/main.go
```

## Building Helm Chart
```bash
helm package charts/bookinfo
```

## License
MIT
# bookinfo-source
