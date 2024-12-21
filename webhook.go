package whatsapp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler interface {
    OnStatusChange(s *Status)
    OnNewMessage(m *Message)
}

type HandleEventFunc[T any] func(e *T)

type Webhook struct {
     token      string

     // Notifications handlers
     OnStatusChangeCallback HandleEventFunc[Status]
}

// Whatsapp send us this
type NotificationPayload struct {
    Object      string   `json:"object"`
    Entries     []Entry  `json:"entry"`
}

type Entry struct {
    ID      string   `json:"id"`
    Changes []Change `json:"changes"`
}

type Change struct {
    Value Value  `json:"value"`
    Field string `json:"field"`
}

type Value struct {
    //MessagingProduct string   `json:"messaging_product"`
    Metadata         Metadata `json:"metadata"`
    Contacts         []Contact `json:"contacts"`
    Errors           []Error   `json:"errors"`
    Messages         []Message `json:"messages"`
    Statuses         []Status  `json:"statuses"`
}

type Message struct {
    From        string  `json:"from"`
    Id          string  `json:"id"`
    Timestamp   string  `json:"timestamp"`
    Type        MessageType  `json:"type"`

    Text        *Text           `json:"text,omitempty"`
    Interactive *Interactive    `json:"interactive,omitempty"`
    Button      *Button         `json:"button,omitempty"`
}

type Text struct {
    Body    string  `json:"body"`
}

type Button struct {
    Text string `json:"text"`
    Payload string `json:"payload"`
}

type Interactive struct {
    ButtonReply *ButtonReply `json:"button_reply,omitempty"`
    ListReply   *ListReply   `json:"list_reply,omitempty"`
}

type ButtonReply struct {
    ID    string `json:"id"`
    Title string `json:"title"`
}

type ListReply struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type Metadata struct {
    DisplayPhoneNumber string `json:"display_phone_number"`
    PhoneNumberID      string `json:"phone_number_id"`
}

type Error struct {
	// Common fields for both v15 and v16+
	Code  uint32    `json:"code"`
	Title string    `json:"title"`

	// Additional fields for v16+
	Message  string             `json:"message,omitempty"`
	ErrorData map[string]string `json:"error_data,omitempty"`
}

type Status struct {
    Id          string  `json:"id"`
    RecipientId string  `json:"recipient_id"` 
    Errors      []Error `json:"errors"`
    Status      DeliveryStatus  `json:"status"`
    Timestamp   string  `json:"timestamp"`
}

type DeliveryStatus string
const (
	DeliveryStatusDelivered DeliveryStatus = "delivered"
	DeliveryStatusRead      DeliveryStatus = "read"
	DeliveryStatusSent      DeliveryStatus = "sent"
)
    
var deliveryStatusValue = map[string]DeliveryStatus {
    "delivered": DeliveryStatusDelivered,
    "read": DeliveryStatusRead,
    "sent": DeliveryStatusSent, 
}

func (s *DeliveryStatus) UnmarshalJSON(data []byte) (err error) {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}
	if *s, err = parseDeliveryStatus(status); err != nil {
		return err
	}
	return nil
}

func parseDeliveryStatus(s string) (DeliveryStatus, error) {
    value, ok := deliveryStatusValue[s]
	if !ok {
		return DeliveryStatus(""), fmt.Errorf("%q is not a valid delivery status", s)
	}
	return DeliveryStatus(value), nil
}

type Contact struct {
	WaID    string `json:"wa_id"`
	Profile struct {
        Name    string `json:"name"`
    }`json:"profile"`
}

func NewWebhook(token string) *Webhook {
    return &Webhook {
        token: token,
    }
}

func (wh *Webhook) reciveNotificaction(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    var notification NotificationPayload
    if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    wh.handleNotification(notification)

    w.WriteHeader(http.StatusOK)
}

func (wh *Webhook) OnStatusChange(callback HandleEventFunc[Status]){
    wh.OnStatusChangeCallback = callback
}

func (wh *Webhook) handleNotification(noti NotificationPayload) {
    for _, entry := range noti.Entries {
        for _, change := range entry.Changes {
            for _, status := range change.Value.Statuses {
                go wh.OnStatusChangeCallback(&status)
            }
            // TODO:
            // for _, status := range change.Value.Messages {
            //     go wh.OnNewMessage(status)
            // }
            // for _, status := range change.Value.Contacts {
            //     go wh.OnNewMessage(status)
            // }
        }
    }
}

func (wh *Webhook) ServeHTTP(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodPost:
        wh.reciveNotificaction(w, r)
        return
    case http.MethodGet:
        wh.verify(w, r)
        return
    }
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (wh *Webhook) verify(w http.ResponseWriter, r *http.Request) {
    challenge := r.URL.Query().Get("hub.challenge")
    verify_token := r.URL.Query().Get("hub.verify_token")
    mode := r.URL.Query().Get("hub.mode")

    if mode == "subscribe" && verify_token == wh.token {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(challenge))
    } else {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("400 - Bad request"))
    }
}
