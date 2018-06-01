package model

import (
	"bytes"
	"fmt"
)

/**
 * Details about the Speech Synthesis Markup Language (SSML) can be found on this page:
 * https://dueros.baidu.com/didp/doc/dueros-bot-platform/dbp-custom/ssml_markdown
 */

// Helper Types

type SSMLTextBuilder struct {
	buffer *bytes.Buffer
}

func NewSSMLTextBuilder() *SSMLTextBuilder {
	return &SSMLTextBuilder{bytes.NewBufferString("")}
}

func (this *SSMLTextBuilder) AppendPlainSpeech(text string) *SSMLTextBuilder {

	this.buffer.WriteString(text)

	return this
}

func (this *SSMLTextBuilder) AppendAudio(src string) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<audio src=\"%s\"></audio>", src))

	return this
}

/*
func (this *SSMLTextBuilder) AppendBreak(strength, time string) *SSMLTextBuilder {

	if strength == "" {
		// The default strength is medium
		strength = "medium"
	}

	this.buffer.WriteString(fmt.Sprintf("<break strength=\"%s\" time=\"%s\"/>", strength, time))

	return this
}

func (this *SSMLTextBuilder) AppendEmphasis(text, level string) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<emphasis level=\"%s\">%s</emphasis>", level, text))

	return this
}

func (this *SSMLTextBuilder) AppendParagraph(text string) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<p>%s</p>", text))

	return this
}

func (this *SSMLTextBuilder) AppendProsody(text, rate, pitch, volume string) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<prosody rate=\"%s\" pitch=\"%s\" volume=\"%s\">%s</prosody>", rate, pitch, volume, text))

	return this
}

func (this *SSMLTextBuilder) AppendSentence(text string) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<s>%s</s>", text))

	return this
}
*/

func (this *SSMLTextBuilder) AppendSilence(time int) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<silence time=\"%s\"></silence>", time))

	return this
}

func (this *SSMLTextBuilder) AppendSubstitution(text, alias string) *SSMLTextBuilder {

	this.buffer.WriteString(fmt.Sprintf("<sub alias=\"%s\">%s</sub>", alias, text))

	return this
}

func (this *SSMLTextBuilder) AppendBackground(text string, src string, repeat bool) *SSMLTextBuilder {
	repeatAttr := "yes"
	if !repeat {
		repeatAttr = ""
	}

	this.buffer.WriteString(fmt.Sprintf("<background src=\"%s\" repeat=\"%s\">%s</background>", src, repeatAttr, text))
	return this
}

func (this *SSMLTextBuilder) ApplyBackground(src string, repeat bool) *SSMLTextBuilder {
	repeatAttr := "yes"
	if !repeat {
		repeatAttr = ""
	}

	buffer := bytes.NewBufferString("")
	buffer.WriteString(fmt.Sprintf("<background src=\"%s\" repeat=\"%s\">%s</background>", src, repeatAttr, this.buffer.String()))

	this.buffer = buffer
	return this
}

func (this *SSMLTextBuilder) Build() string {
	return fmt.Sprintf("<speak>%s</speak>", this.buffer.String())
}
