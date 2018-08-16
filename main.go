package main

import (
	"github.com/nsip/nias3/webserver"
)

func main() {
	//webserver.SIFGetManyToDataStore("http://hits.nsip.edu.au/SIF3InfraREST/hits/requests/SchoolInfos?navigationPage=1&navigationPageSize=5&access_token=ZmZhODMzNjEtMGExOC00NDk5LTgyNjMtYjMwNjI4MGRjZDRlOmYxYzA1NjNhOWIzZTQyMGJiMDdkYTJkOTBkYjQ3OWVm&authenticationMethod=Basic")
	//webserver.SendXmlToDataStore("nswdig.xml")
	webserver.SendXmlToDataStore("staffpersonals.xml")
	webserver.SendJsonToDataStore("xapi.json", "xAPI")
	webserver.Webserver()
}
