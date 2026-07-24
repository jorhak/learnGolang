package domain

// Customer representa un cliente del CRM.
type Customer struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
