// https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components/

package whatsapp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandleEventFunc[T any] func(e *T)

type Webhook struct {
     verifyToken      string

     // Notifications handlers
     onStatusChange HandleEventFunc[Status]
     onNewMessage   HandleEventFunc[Message]
}

// Whatsapp send us this
type NotificationPayload struct { 
    // The specific webhook. The webhook is whatsapp_business_account
    Object      string   `json:"object"`
    Entries     []Entry  `json:"entry"`
}

type Entry struct {
    ID      string   `json:"id"`
    Changes []Change `json:"changes"`
}

type Change struct {
    Value Value  `json:"value"`
    // Notification Type. Value will be messages
    Field string `json:"field"`
}

type Value struct {
    // Product used to send the message. Value is always whatsapp.    
    MessagingProduct string   `json:"messaging_product"` 
    // A metadata object describing the business subscribed to the webhook.
    Metadata         Metadata `json:"metadata"`
    Contacts         []Contact `json:"contacts"`
    Errors           []Error   `json:"errors"`
    Messages         []Message `json:"messages"`
    Statuses         []Status  `json:"statuses"`
}

// Message represents a single message received through the webhook.
type Message struct {
    From       string      `json:"from"`
    ID         string      `json:"id"`
    Timestamp  string      `json:"timestamp"`
    Errors     []Error     `json:"errors,omitempty"`
    // Only included when a user replies or interacts with one of your messages.
    Context    *Context    `json:"context,omitempty"`
    // Webhook is triggered when a customer's phone number or profile information has been updated
    Identity   *Identity   `json:"identity,omitempty"`

    Type       MessageType `json:"type"`
    // Message Type
    Audio      *Audio      `json:"audio,omitempty"`
    Button     *Button     `json:"button,omitempty"`
    Document   *Document   `json:"document,omitempty"`
    Image      *Image      `json:"image,omitempty"`
    Interactive *Interactive `json:"interactive,omitempty"`
    Order      *Order      `json:"order,omitempty"`
    Referral   *Referral   `json:"referral,omitempty"`
    Sticker    *Sticker    `json:"sticker,omitempty"`
    System     *System     `json:"system,omitempty"`
    Text       *Text       `json:"text,omitempty"`
    Video      *Video      `json:"video,omitempty"`
}

// Audio object for audio messages.
type Audio struct {
    ID        string `json:"id"`
    MimeType string `json:"mime_type"`
}

// Button object for button messages.
type Button struct {
    Payload string `json:"payload"`
    Text    string `json:"text"`
}

// Context object for context information.
type Context struct {
    Forwarded          bool           `json:"forwarded"`
    FrequentlyForwarded bool           `json:"frequently_forwarded"`
    From               string          `json:"from"`
    ID                 string          `json:"id"`
    ReferredProduct    *ReferredProduct `json:"referred_product,omitempty"`
}

// ReferredProduct object within Context.
type ReferredProduct struct {
    CatalogID        string `json:"catalog_id"`
    ProductRetailerID string `json:"product_retailer_id"`
}

// Document object for document messages.
type Document struct {
    Caption   string `json:"caption"`
    Filename  string `json:"filename"`
    Sha256    string `json:"sha256"`
    MimeType string `json:"mime_type"`
    ID        string `json:"id"`
}

// Identity object for customer profile updates.
type Identity struct {
    Acknowledged    bool   `json:"acknowledged"`
    CreatedTimestamp string `json:"created_timestamp"`
    Hash            string `json:"hash"`
}

// Image object for image messages.
type Image struct {
    Caption   string `json:"caption"`
    Sha256    string `json:"sha256"`
    ID        string `json:"id"`
    MimeType string `json:"mime_type"`
}

// ButtonReply object within InteractiveType.
type ButtonReply struct {
    ID    string `json:"id"`
    Title string `json:"title"`
}

// ListReply object within InteractiveType.
type ListReply struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

// Order object for order messages.
type Order struct {
    CatalogID   string       `json:"catalog_id"`
    Text        string       `json:"text"`
    ProductItems []ProductItem `json:"product_items"`
}

// ProductItem object within Order.
type ProductItem struct {
    ProductRetailerID string `json:"product_retailer_id"`
    Quantity          string `json:"quantity"`
    ItemPrice         string `json:"item_price"`
    Currency          string `json:"currency"`
}

// Referral object for referral messages.
type Referral struct {
    SourceURL     string `json:"source_url"`
    SourceType    string `json:"source_type"`
    SourceID      string `json:"source_id"`
    Headline      string `json:"headline"`
    Body          string `json:"body"`
    MediaType     string `json:"media_type"`
    ImageURL      string `json:"image_url"`
    VideoURL      string `json:"video_url"`
    ThumbnailURL string `json:"thumbnail_url"`
    CtwaClid      string `json:"ctwa_clid"`
}

// Sticker object for sticker messages.
type Sticker struct {
    MimeType string `json:"mime_type"`
    Sha256    string `json:"sha256"`
    ID        string `json:"id"`
    Animated bool   `json:"animated"`
}

// System object for system messages.
type System struct {
    Body         string `json:"body"`
    Identity     string `json:"identity"`
    NewWaID      string `json:"new_wa_id"` // Deprecated, use WaID for v12.0 and later
    WaID         string `json:"wa_id"`
    Type         string `json:"type"`
    Customer     string `json:"customer"`
}

// Video object for video messages.
type Video struct {
    Caption   string `json:"caption"`
    Sha256    string `json:"sha256"`
    ID        string `json:"id"`
    MimeType string `json:"mime_type"`
}

// Text object for text messages.
type Text struct {
    Body    string  `json:"body"`
}

type Interactive struct {
    ButtonReply *ButtonReply `json:"button_reply,omitempty"`
    ListReply   *ListReply   `json:"list_reply,omitempty"`
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

type Contact struct {
	WaID    string `json:"wa_id"`
	Profile struct {
        Name    string `json:"name"`
    }`json:"profile"`
}

type Status struct {
    Id          string  `json:"id"`
    RecipientId string  `json:"recipient_id"` 
    Errors      []Error `json:"errors"`
    Status      DeliveryStatus  `json:"status"`
    Timestamp   string  `json:"timestamp"`
    //TODO: Conversation
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

func NewWebhook(verifyToken string) *Webhook {
    return &Webhook {
        verifyToken: verifyToken,
    }
}

func (wh *Webhook) OnStatusChange(callback HandleEventFunc[Status]){
    wh.onStatusChange = callback
}

func (wh *Webhook) OnNewMessage(callback HandleEventFunc[Message]){
    wh.onNewMessage = callback
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

// implements the http.Handler interface
// this way you can do http.ListenAndServe(":8080", webhook)
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

func (wh *Webhook) handleNotification(noti NotificationPayload) {
    for _, entry := range noti.Entries {
        for _, change := range entry.Changes {
            for _, status := range change.Value.Statuses {
                if wh.onStatusChange != nil {
                    go wh.onStatusChange(&status)
                }
            }
            for _, message := range change.Value.Messages {
                if wh.onNewMessage != nil {
                    go wh.onNewMessage(&message)
                }
            }
            // TODO:
            // for _, status := range change.Value.Contacts {
            //     go wh.OnNewMessage(status)
            // }
        }
    }
}

func (wh *Webhook) verify(w http.ResponseWriter, r *http.Request) {
    challenge := r.URL.Query().Get("hub.challenge")
    verifyToken := r.URL.Query().Get("hub.verify_token")
    mode := r.URL.Query().Get("hub.mode")

    if mode == "subscribe" && verifyToken == wh.verifyToken {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(challenge))
    } else {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("400 - Bad request"))
    }
}
