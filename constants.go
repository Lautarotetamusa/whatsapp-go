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

type InteractionType string
const (
	//TODO: cta_url its not present in the references but in the Examples
	CallToActionType InteractionType = "cta_url"
	ButtonType       InteractionType = "button"
	CatalogType      InteractionType = "catalog_message"
	ListType         InteractionType = "list"
	ProductType      InteractionType = "product"
    ProductList      InteractionType = "product_list"
	FlowType         InteractionType = "flow"
)

// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#interactive-obj
// Action object requires specefic fields.
// requiredFields gives the name of the field required by an
// Specific InteractionType
type ActionRequiredField string
const (
    ButtonsField    ActionRequiredField = "buttons"
    SectionsField   ActionRequiredField = "sections"
    CTAField        ActionRequiredField = "cta_url"
    ParametersField ActionRequiredField = "parameters"
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
