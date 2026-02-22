```markdown
# Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    SHOP ||--o{ USER : employs
    SHOP ||--o{ PRODUCT : owns
    SHOP ||--o{ TRANSACTION : records
    PRODUCT ||--o{ TRANSACTION : involves

    SHOP {
        int id PK
        string name
        boolean active
        string whatsapp_number
    }
    USER {
        int id PK
        string name
        string email
        string password
        string role
        int shop_id FK
    }
    PRODUCT {
        int id PK
        string name
        float purchase_price
        float selling_price
        int stock
        int shop_id FK
    }
    TRANSACTION {
        int id PK
        string type
        int quantity
        float amount
        int product_id FK
        int shop_id FK
    }