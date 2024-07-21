package decoder_legacy

import "github.com/koykov/decoder"

func init() {
	decoder.RegisterCallbackFn("jsonParseAs", "jsonParse", cbJsonParse).
		WithParam("data string", "Source data to parse.").
		WithParam("name string", "Name of new variable.").
		WithDescription("Parse JSON `data` and save result in context as `name`.").
		WithNote("DEPRECATED! Use `vector::parseJSON()` instead.")
	decoder.RegisterCallbackFn("urlParseAs", "urlParse", cbUrlParse).
		WithParam("data string", "Source data to parse.").
		WithParam("name string", "Name of new variable.").
		WithDescription("Parse URL `data` and save result in context as `name`.").
		WithNote("DEPRECATED! Use `vector::parseURL()` instead.")
	decoder.RegisterCallbackFn("xmlParseAs", "xmlParse", cbXmlParse).
		WithParam("data string", "Source data to parse.").
		WithParam("name string", "Name of new variable.").
		WithDescription("Parse XML `data` and save result in context as `name`.").
		WithNote("DEPRECATED! Use `vector::parseXML()` instead.")
	decoder.RegisterCallbackFn("yamlParseAs", "yamlParse", cbYamlParse).
		WithParam("data string", "Source data to parse.").
		WithParam("name string", "Name of new variable.").
		WithDescription("Parse YAML `data` and save result in context as `name`.").
		WithNote("DEPRECATED! Use `vector::parseYAML()` instead.\nCAUTION! Still not implement.")
	decoder.RegisterCallbackFn("halParseAs", "halParse", cbHalParse).
		WithParam("data string", "Source data to parse.").
		WithParam("name string", "Name of new variable.").
		WithDescription("Parse HAL `data` and save result in context as `name`.").
		WithNote("DEPRECATED! Use `vector::parseHAL()` instead.")
}
