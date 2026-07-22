package main

import(
   "log"
   "crm-api/internal/server"
)

func main() {
    app := server.New()
    log.Println("CRM API inicada")
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
