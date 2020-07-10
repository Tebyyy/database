package database

type Language string

const (
	English Language = "en"
	French  Language = "fr"
	Spanish Language = "es"
	German  Language = "de"
)

var Flags = map[Language]string{
	English: "🇬🇧",
	French:  "🇫🇷",
	Spanish: "🇪🇸",
	German:  "🇩🇪",
}

// https://discord.com/developers/docs/dispatch/field-values
var Locales = map[string]Language{
	"en-US": English,
	"en-GB": English,
	"zh-CN": English, // Chinese (China)
	"zh-TW": English, // Chinese (Taiwan)
	"cs":    English, // Czech
	"da":    English, // Danish
	"nl":    English, // Dutch
	"fr":    French,
	"de":    German,
	"hu":    English, // Hungarian
	"it":    English, // Italian
	"ja":    English, // Japanese
	"ko":    English, // Korean
	"no":    English, // Norwegian
	"pl":    English, // Polish
	"pt-BR": English, // Portuguese (Brazil)
	"ru":    English, // Russian
	"es-ES": Spanish,
	"sv-SE": English, // Swedish
	"tr":    English, // Turkish
	"bg":    English, // Bulgarian
	"uk":    English, // Ukrainian
	"fi":    English, // Finnish
	"hr":    English, // Croatian
	"ro":    English, // Romanian
	"lt":    English, // Lithuanian
}