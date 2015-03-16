package main

import "text/template"
import "os"

type Guest struct {
	Honorific string
	Name      string
	DidAttend bool
	Donate    bool
	NewEvents []string
}

var myNewEvents = []string{"Minon Day", "Minion Movie Premier", "Gru's Brhday"}
var myGuests = []Guest{
	{"Mr.", "Gru", false, false, myNewEvents},
	{"Dr.", "Kevin", true, true, myNewEvents},
	{"Mr.", "Dave", true, false, myNewEvents},
	{"Mr.", "Stuart", true, true, myNewEvents},
	{"Mr.", "Carl", true, false, myNewEvents},
}

func main() {
	myGuetsTemp, err := template.New("LetterTemp").Parse("Dear {{.Honorific}} {{.Name}}:" +
		" \n{{if (.DidAttend) }}Thank you very much for attending our fund-raising event.{{if (.Donate) }}" +
		"  Your donation is greatly appreciated.{{else}}  Please consider a small monetary donation if you appreciated our event. " +
		" {{end}}{{else}}We are sorry that you could not make it to our event.{{end}} " +
		" Let get together for upcoming events:{{range $element  := .NewEvents}} {{$element}},{{end}} .\n" +
		" BestWishes,\n Harleen \n\n")
	if err != nil {
		panic(err)
	}
	for guest := range myGuests {
		err = myGuetsTemp.Execute(os.Stdout, myGuests[guest])
		if err != nil {
			panic(err)
		}
	}
}
