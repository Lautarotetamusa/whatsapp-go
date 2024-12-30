package whatsapp

type MessageType string

const (
	TextType        MessageType = "text"
	ImageType       MessageType = "image"
	AudioType       MessageType = "audio"
	StickerType     MessageType = "sticker"
	VideoType       MessageType = "video"
	DocumentType    MessageType = "document"
	TemplateType    MessageType = "template"
	ContactsType    MessageType = "contacts"
	InteractiveType MessageType = "interactive"
)

// button: se usa para los botones de respuesta.
// catalog_message: se usa para los mensajes de catálogo.
// list: se usa para los mensajes de lista.
// product: se usa para los mensajes sobre un solo producto.
// product_list: se usa para los mensajes sobre varios productos.
// flow: se usa para los mensajes de Flows.
type InteractionType string

const (
	//TODO: cta_url its not present in the references but in the Examples
	CallToActionType InteractionType = "cta_url"
	ButtonType       InteractionType = "button"
	ButtonsType      InteractionType = "buttons"
	CatalogType      InteractionType = "catalog_message"
	List             InteractionType = "list"
	Product          InteractionType = "product"
	Product_list     InteractionType = "product_list"
	Flow             InteractionType = "flow"
)

// "header" or "body"
type ComponentType string

const (
	ComponentTypeHeader ComponentType = "header"
	ComponentTypeBody   ComponentType = "body"
)

// https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#idiomas-admitidos
type LanguageCode string

const (
	English          LanguageCode = "en"
	EnglishUS        LanguageCode = "en_US"
	Afrikaans        LanguageCode = "af"
	Albanian         LanguageCode = "sq"
	Arabic           LanguageCode = "ar"
	Azerbaijani      LanguageCode = "az"
	Bengali          LanguageCode = "bn"
	Bulgarian        LanguageCode = "bg"
	Catalan          LanguageCode = "ca"
	ChineseCHN       LanguageCode = "zh_CN"
	ChineseHKG       LanguageCode = "zh_HK"
	ChineseTAI       LanguageCode = "zh_TW"
	Croatian         LanguageCode = "hr"
	Czech            LanguageCode = "cs"
	Danish           LanguageCode = "da"
	Dutch            LanguageCode = "nl"
	Estonian         LanguageCode = "et"
	Filipino         LanguageCode = "fil"
	Finnish          LanguageCode = "fi"
	French           LanguageCode = "fr"
	Georgian         LanguageCode = "ka"
	German           LanguageCode = "de"
	Greek            LanguageCode = "el"
	Gujarati         LanguageCode = "gu"
	Hausa            LanguageCode = "ha"
	Hebrew           LanguageCode = "he"
	Hindi            LanguageCode = "hi"
	Hungarian        LanguageCode = "hu"
	Indonesian       LanguageCode = "id"
	Irish            LanguageCode = "ga"
	Italian          LanguageCode = "it"
	Japanese         LanguageCode = "ja"
	Kannada          LanguageCode = "kn"
	Kazakh           LanguageCode = "kk"
	Kinyarwanda      LanguageCode = "rw_RW"
	Korean           LanguageCode = "ko"
	KyrgyzKyrgyzstan LanguageCode = "ky_KG"
	Lao              LanguageCode = "lo"
	Latvian          LanguageCode = "lv"
	Lithuanian       LanguageCode = "lt"
	Macedonian       LanguageCode = "mk"
	Malay            LanguageCode = "ms"
	Malayalam        LanguageCode = "ml"
	Marathi          LanguageCode = "mr"
	Norwegian        LanguageCode = "nb"
	Persian          LanguageCode = "fa"
	Polish           LanguageCode = "pl"
	PortugueseBR     LanguageCode = "pt_BR"
	PortuguesePOR    LanguageCode = "pt_PT"
	Punjabi          LanguageCode = "pa"
	Romanian         LanguageCode = "ro"
	Russian          LanguageCode = "ru"
	Serbian          LanguageCode = "sr"
	Slovak           LanguageCode = "sk"
	Slovenian        LanguageCode = "sl"
	Spanish          LanguageCode = "es"
	SpanishARG       LanguageCode = "es_AR"
	SpanishSPA       LanguageCode = "es_ES"
	SpanishMEX       LanguageCode = "es_MX"
	Swahili          LanguageCode = "sw"
	Swedish          LanguageCode = "sv"
	Tamil            LanguageCode = "ta"
	Telugu           LanguageCode = "te"
	Thai             LanguageCode = "th"
	Turkish          LanguageCode = "tr"
	Ukrainian        LanguageCode = "uk"
	Urdu             LanguageCode = "ur"
	Uzbek            LanguageCode = "uz"
	Vietnamese       LanguageCode = "vi"
	Zulu             LanguageCode = "zu"
)
