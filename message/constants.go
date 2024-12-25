package message

// https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#idiomas-admitidos
type LanguageCode string

const (
	English LanguageCode = "en"
	EnglishUS LanguageCode = "en_US"
	Afrikaans LanguageCode = "af"
	Albanian LanguageCode = "sq"
	Arabic LanguageCode = "ar"
	Azerbaijani LanguageCode = "az"
	Bengali LanguageCode = "bn"
	Bulgarian LanguageCode = "bg"
	Catalan LanguageCode = "ca"
	ChineseCHN LanguageCode = "zh_CN"
	ChineseHKG LanguageCode = "zh_HK"
	ChineseTAI LanguageCode = "zh_TW"
	Croatian LanguageCode = "hr"
	Czech LanguageCode = "cs"
	Danish LanguageCode = "da"
	Dutch LanguageCode = "nl"
	Estonian LanguageCode = "et"
	Filipino LanguageCode = "fil"
	Finnish LanguageCode = "fi"
	French LanguageCode = "fr"
	Georgian LanguageCode = "ka"
	German LanguageCode = "de"
	Greek LanguageCode = "el"
	Gujarati LanguageCode = "gu"
	Hausa LanguageCode = "ha"
	Hebrew LanguageCode = "he"
	Hindi LanguageCode = "hi"
	Hungarian LanguageCode = "hu"
	Indonesian LanguageCode = "id"
	Irish LanguageCode = "ga"
	Italian LanguageCode = "it"
	Japanese LanguageCode = "ja"
	Kannada LanguageCode = "kn"
	Kazakh LanguageCode = "kk"
	Kinyarwanda LanguageCode = "rw_RW"
	Korean LanguageCode = "ko"
	KyrgyzKyrgyzstan LanguageCode = "ky_KG"
	Lao LanguageCode = "lo"
	Latvian LanguageCode = "lv"
	Lithuanian LanguageCode = "lt"
	Macedonian LanguageCode = "mk"
	Malay LanguageCode = "ms"
	Malayalam LanguageCode = "ml"
	Marathi LanguageCode = "mr"
	Norwegian LanguageCode = "nb"
	Persian LanguageCode = "fa"
	Polish LanguageCode = "pl"
	PortugueseBR LanguageCode = "pt_BR"
	PortuguesePOR LanguageCode = "pt_PT"
	Punjabi LanguageCode = "pa"
	Romanian LanguageCode = "ro"
	Russian LanguageCode = "ru"
	Serbian LanguageCode = "sr"
	Slovak LanguageCode = "sk"
	Slovenian LanguageCode = "sl"
	Spanish LanguageCode = "es"
	SpanishARG LanguageCode = "es_AR"
	SpanishSPA LanguageCode = "es_ES"
	SpanishMEX LanguageCode = "es_MX"
	Swahili LanguageCode = "sw"
	Swedish LanguageCode = "sv"
	Tamil LanguageCode = "ta"
	Telugu LanguageCode = "te"
	Thai LanguageCode = "th"
	Turkish LanguageCode = "tr"
	Ukrainian LanguageCode = "uk"
	Urdu LanguageCode = "ur"
	Uzbek LanguageCode = "uz"
	Vietnamese LanguageCode = "vi"
	Zulu LanguageCode = "zu"
)

var langCodeValue = map[string]LanguageCode  {
	"en": English,
	"en_US": EnglishUS,
	"af": Afrikaans,
	"sq": Albanian,
	"ar": Arabic,
	"az": Azerbaijani,
	"bn": Bengali,
	"bg": Bulgarian,
	"ca": Catalan,
	"zh_CN": ChineseCHN,
	"zh_HK": ChineseHKG,
	"zh_TW": ChineseTAI,
	"hr": Croatian,
	"cs": Czech,
	"da": Danish,
	"nl": Dutch,
	"et": Estonian,
	"fil": Filipino,
	"fi": Finnish,
	"fr": French,
	"ka": Georgian,
	"de": German,
	"el": Greek,
	"gu": Gujarati,
	"ha": Hausa,
	"he": Hebrew,
	"hi": Hindi,
	"hu": Hungarian,
	"id": Indonesian,
	"ga": Irish,
	"it": Italian,
	"ja": Japanese,
	"kn": Kannada,
	"kk": Kazakh,
	"rw_RW": Kinyarwanda,
	"ko": Korean,
	"ky_KG": KyrgyzKyrgyzstan,
	"lo": Lao,
	"lv": Latvian,
	"lt": Lithuanian,
	"mk": Macedonian,
	"ms": Malay,
	"ml": Malayalam,
	"mr": Marathi,
	"nb": Norwegian,
	"fa": Persian,
	"pl": Polish,
	"pt_BR": PortugueseBR,
	"pt_PT": PortuguesePOR,
	"pa": Punjabi,
	"ro": Romanian,
	"ru": Russian,
	"sr": Serbian,
	"sk": Slovak,
	"sl": Slovenian,
	"es": Spanish,
	"es_AR": SpanishARG,
	"es_ES": SpanishSPA,
	"es_MX": SpanishMEX,
	"sw": Swahili,
	"sv": Swedish,
	"ta": Tamil,
	"te": Telugu,
	"th": Thai,
	"tr": Turkish,
	"uk": Ukrainian,
	"ur": Urdu,
	"uz": Uzbek,
	"vi": Vietnamese,
	"zu": Zulu,
}
