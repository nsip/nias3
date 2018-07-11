curl -i -X GET http://localhost:1492/sifxml/TimeTableSubjects
curl -i -X GET "http://localhost:1492/sifxml/kla2student?kla=French&yrlvl=10"
curl -i -H "Accept: application/json" -X GET "http://localhost:1492/sifxml/kla2student?kla=French&yrlvl=10"
curl -i -X GET "http://localhost:1492/sifxml/kla2staff?kla=French&yrlvl=10"
curl -i -H "Accept: application/json" -X GET "http://localhost:1492/sifxml/kla2staff?kla=French&yrlvl=10"

