# Diagramme d'Architecture

```mermaid
graph TD
    Client[Client Web / Postman] -->|HTTP REST| API[Fiber API Go]
    API -->|Validation| Middleware[Auth / Role / Tenant]
    Middleware --> Handlers[Logique Métier]
    Handlers -->|GORM| DB[(PostgreSQL 14)]
    
    subgraph Docker Infrastructure
        API
        DB
    end