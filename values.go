package slugger

// GermanCharacterMapping provides a mapping of German special characters to their ASCII equivalents.
var GermanCharacterMapping = map[string]string{
	"ä": "ae",
	"ö": "oe",
	"ü": "ue",
	"ß": "ss",
	"Ä": "Ae",
	"Ö": "Oe",
	"Ü": "Ue",
	"@": "at",
	"&": "und",
}

// EnglishCharacterMapping provides a mapping of English special characters to their ASCII equivalents.
var EnglishCharacterMapping = map[string]string{
	"@": "at",
	"&": "and",
}

// SpanishCharacterMapping provides a mapping of Spanish special characters to their ASCII equivalents.
var SpanishCharacterMapping = map[string]string{
	"á": "a",
	"é": "e",
	"í": "i",
	"ó": "o",
	"ú": "u",
	"ñ": "n",
	"Á": "A",
	"É": "E",
	"Í": "I",
	"Ó": "O",
	"Ú": "U",
	"Ñ": "N",
	"@": "at",
	"&": "y",
	"ü": "u",
	"Ü": "U",
}
