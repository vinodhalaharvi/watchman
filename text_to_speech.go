package main

import (
	htgotts "github.com/hegedustibor/htgo-tts"
)

func speak(what string) {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak(what)
}
