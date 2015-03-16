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

var myNewEvents = []string{"Minion Day", "Minion Movie Premier", "Gru's Birthday"}
var myGuests = []Guest{
	{"Mr.", "Gru", false, false, myNewEvents},
	{"Dr.", "Kevin", true, true, myNewEvents},
	{"Mr.", "Dave", true, false, myNewEvents},
	{"Mr.", "Stuart", true, true, myNewEvents},
	{"Mr.", "Carl", true, false, myNewEvents},
}

func main() {
	myGuetsTemp, err := template.New("LetterTemp").Parse("Dear {{.Honorific}} {{.Name}}:" +
		" \n{{if (.DidAttend) }}Thank you for attending our event.{{if (.Donate) }}" +
		"  Your donation is much appreciated." +
		" {{end}}{{else}}You were missed at the event.{{end}} " +
		" Lets get together for upcoming events:{{range $element  := .NewEvents}} {{$element}},{{end}} .\n" +
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
