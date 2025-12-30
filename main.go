package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/taber-miyauchi/notifier-core"
	"github.com/taber-miyauchi/notifier-email"
)

func main() {
	notifier := email.NewEmailNotifier("smtp.example.com", "noreply@example.com")

	handler := NewNotificationHandler(notifier)

	fmt.Println("Starting notification service on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// NotificationHandler handles HTTP requests for sending notifications.
type NotificationHandler struct {
	notifier core.Notifier
}

// NewNotificationHandler creates a handler with the given notifier.
func NewNotificationHandler(n core.Notifier) *NotificationHandler {
	return &NotificationHandler{notifier: n}
}

// ServeHTTP processes notification requests.
func (h *NotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	msg := core.NewMessage(
		r.FormValue("to"),
		r.FormValue("subject"),
		r.FormValue("body"),
	)

	if err := h.notifier.Send(r.Context(), msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Notification sent!")
}
