package decoder_legacy

import "github.com/koykov/decoder"

func init() {
	decoder.RegisterCallbackFn("jsonParseAs", "jsonParse", cbJsonParse)
	decoder.RegisterCallbackFn("urlParseAs", "urlParse", cbUrlParse)
	decoder.RegisterCallbackFn("xmlParseAs", "xmlParse", cbXmlParse)
	decoder.RegisterCallbackFn("yamlParseAs", "yamlParse", cbYamlParse)
	decoder.RegisterCallbackFn("halParseAs", "halParse", cbHalParse)
}
