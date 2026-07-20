package main

import "fmt"

type NotificationSender interface {
    Send(message string)
}

type EmailNotification struct {}
type SMSNotification struct {}
type PushNotification struct {}

func (e EmailNotification) Send(message string) {
    fmt.Printf(
        "Nofificacion correo: %s\n",
        message,
    )
}

func (s SMSNotification) Send(message string) {
    fmt.Printf(
        "Notificacion SMS: %s\n",
         message,
    )
}

func (p PushNotification) Send(message string) {
    fmt.Printf(
        "Nofificacion Push: %s\n",
         message,
    )
}

type NotificationService struct {
    notification NotificationSender
}

func NewNotificationService(
    notification NotificationSender,
) *NotificationService {
   return &NotificationService{
         notification: notification,
   }
}

func (n NotificationService) Process(message string) {
    n.notification.Send(message)
}

func main() {
    service := NewNotificationService(
              EmailNotification{},
    )
    service.Process("Hola como estas")
    service = NewNotificationService(
             SMSNotification{},
    )
    service.Process("Su saldo es insuficiente")
    service = NewNotificationService(
             PushNotification{},
    )
    service.Process("Verificacion Multifactor")
}
