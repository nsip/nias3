package main

import (
	"encoding/json"
	//"encoding/xml"
	//"fmt"
	"github.com/clbanning/mxj"
	//"github.com/nsip/nias3/sifxml"
	"log"
	"strings"
)

var purchase_order = []byte(`
<PurchaseOrder RefId="ED12345F-DA84-9727-5BC2-8AA349DD3721">
  <FormNumber>00342</FormNumber>
  <VendorInfoRefId>BD12345F-DA84-9727-5BC2-8AA349DD3723</VendorInfoRefId>
  <ChargedLocationInfoRefId>ED12345F-DA84-9727-5BC2-8AA349DD3722</ChargedLocationInfoRefId>
  <EmployeePersonalRefId>AD12345F-DA84-9727-5BC2-8AA349DD3721</EmployeePersonalRefId>
  <PurchasingItems>
    <PurchasingItem>
      <ItemNumber>154486</ItemNumber>
      <ItemDescription>Floor Wax</ItemDescription>
      <Quantity>10</Quantity>
      <UnitCost Currency="AUD">7.00</UnitCost>
      <ExpenseAccounts>
        <ExpenseAccount>
          <AccountCode>10-1100-610</AccountCode>
          <Amount Currency="AUD">70.00</Amount>
        </ExpenseAccount>
      </ExpenseAccounts>
    </PurchasingItem>
  </PurchasingItems>
</PurchaseOrder>
`)

var purchase_order2 = []byte(`
<PurchaseOrder RefId="ED12345F-DA84-9727-5BC2-8AA349DD3721">
  <FormNumber>00342</FormNumber>
  <VendorInfoRefId>BD12345F-DA84-9727-5BC2-8AA349DD3723</VendorInfoRefId>
  <ChargedLocationInfoRefId>ED12345F-DA84-9727-5BC2-8AA349DD3722</ChargedLocationInfoRefId>
  <EmployeePersonalRefId>AD12345F-DA84-9727-5BC2-8AA349DD3721</EmployeePersonalRefId>
  <PurchasingItems>
    <PurchasingItem>
      <ItemNumber>154486</ItemNumber>
      <ItemDescription>Floor Wax</ItemDescription>
      <Quantity>10</Quantity>
      <UnitCost Currency="AUD">7.00</UnitCost>
      <ExpenseAccounts>
        <ExpenseAccount>
          <AccountCode>10-1100-610</AccountCode>
          <Amount Currency="AUD">70.00</Amount>
        </ExpenseAccount>
      </ExpenseAccounts>
    </PurchasingItem>
        <PurchasingItem>
      <ItemNumber>154487</ItemNumber>
      <ItemDescription>Ear Wax</ItemDescription>
      <Quantity>12</Quantity>
      <UnitCost Currency="AUD">8.00</UnitCost>
      <ExpenseAccounts>
        <ExpenseAccount>
          <AccountCode>10-1100-611</AccountCode>
          <Amount Currency="AUD">72.00</Amount>
        </ExpenseAccount>
      </ExpenseAccounts>
    </PurchasingItem>
  </PurchasingItems>
</PurchaseOrder>
`)

var staff_personal = []byte(`
<StaffPersonal RefId="D3E34F41-9D75-101A-8C3D-00AA001A1652">
  <LocalId>946379881</LocalId>
  <StateProvinceId>C2345681</StateProvinceId>
  <ElectronicIdList>
    <ElectronicId Type="01">206655</ElectronicId>
  </ElectronicIdList>
  <OtherIdList>
    <OtherId Type="0004">333333333</OtherId>
  </OtherIdList>
  <PersonInfo>
    <Name Type="LGL">
      <FamilyName>Smith</FamilyName>
      <GivenName>Fred</GivenName>
      <FullName>Fred Smith</FullName>
    </Name>
    <OtherNames>
      <Name Type="AKA">
        <FamilyName>Anderson</FamilyName>
        <GivenName>Samuel</GivenName>
        <FullName>Samuel Anderson</FullName>
      </Name>
      <Name Type="PRF">
        <FamilyName>Rowinski</FamilyName>
        <GivenName>Sam</GivenName>
        <FullName>Sam Rowinski </FullName>
      </Name>
    </OtherNames>
    <Demographics>
      <IndigenousStatus>3</IndigenousStatus>
      <Sex>1</Sex>
      <BirthDate>1990-09-26</BirthDate>
      <BirthDateVerification>1004</BirthDateVerification>
      <PlaceOfBirth>Clayton</PlaceOfBirth>
      <StateOfBirth>VIC</StateOfBirth>
      <CountryOfBirth>1101</CountryOfBirth>
      <CountriesOfCitizenship>
        <CountryOfCitizenship>8104</CountryOfCitizenship>
        <CountryOfCitizenship>1101</CountryOfCitizenship>
      </CountriesOfCitizenship>
      <CountriesOfResidency>
        <CountryOfResidency>8104</CountryOfResidency>
        <CountryOfResidency>1101</CountryOfResidency>
      </CountriesOfResidency>
      <CountryArrivalDate>1990-09-26</CountryArrivalDate>
      <AustralianCitizenshipStatus>1</AustralianCitizenshipStatus>
      <EnglishProficiency>
        <Code>1</Code>
      </EnglishProficiency>
      <LanguageList>
        <Language>
          <Code>0001</Code>
          <LanguageType>1</LanguageType>
        </Language>
      </LanguageList>
      <DwellingArrangement>
        <Code>1671</Code>
      </DwellingArrangement>
      <Religion>
        <Code>2013</Code>
      </Religion>
      <ReligiousEventList>
        <ReligiousEvent>
          <Type>Baptism</Type>
          <Date>2000-09-01</Date>
        </ReligiousEvent>
        <ReligiousEvent>
          <Type>Christmas</Type>
          <Date>2009-12-24</Date>
        </ReligiousEvent>
      </ReligiousEventList>
      <ReligiousRegion>The Religion Region</ReligiousRegion>
      <PermanentResident>P</PermanentResident>
      <VisaSubClass>101</VisaSubClass>
      <VisaStatisticalCode>05</VisaStatisticalCode>
    </Demographics>
    <AddressList>
      <Address Type="0123" Role="012A">
        <Street>
          <Line1>Unit1/10</Line1>
          <Line2>Barkley Street</Line2>
        </Street>
        <City>Yarra Glenn</City>
        <StateProvince>VIC</StateProvince>
        <Country>1101</Country>
        <PostalCode>9999</PostalCode>
      </Address>
      <Address Type="0123A" Role="1073">
        <Street>
          <Line1>34 Term Address Street</Line1>
        </Street>
        <City>Home Town</City>
        <StateProvince>WA</StateProvince>
        <Country>1101</Country>
        <PostalCode>9999</PostalCode>
      </Address>
    </AddressList>
    <PhoneNumberList>
      <PhoneNumber Type="0096">
        <Number>03 9637-2289</Number>
        <Extension>72289</Extension>
        <ListedStatus>Y</ListedStatus>
      </PhoneNumber>
      <PhoneNumber Type="0888">
        <Number>0437-765-234</Number>
        <ListedStatus>N</ListedStatus>
      </PhoneNumber>
    </PhoneNumberList>
    <EmailList>
      <Email Type="01">fsmith@yahoo.com</Email>
      <Email Type="02">freddy@gmail.com</Email>
    </EmailList>
  </PersonInfo>
  <Title>Principal</Title>
  <EmploymentStatus>A</EmploymentStatus>
  <MostRecent>
    <SchoolLocalId>S12345</SchoolLocalId>
    <LocalCampusId>D</LocalCampusId>
    <!--<SchoolACARAId>VIC687</SchoolACARAId>-->
    <NAPLANClassList>
      <ClassCode>9</ClassCode>
      <ClassCode>7</ClassCode>
    </NAPLANClassList>
    <HomeGroup>9G</HomeGroup>
  </MostRecent>
</StaffPersonal>
`)

func changeJSONTags(j []byte) []byte {
	return []byte(strings.Replace(string(j), `"#text":`, `"Value":`, -1))
}

func Map2SIFXML(m mxj.Map) ([]byte, error) {
	root, err := m.Root()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(root)
	m02 := m[root].(map[string]interface{})
	log.Printf("m02\n%+v\n", m02)
	j, err := json.Marshal(m02)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(string(j))
	return Root2SIF(root, changeJSONTags(j))
}

func main() {
	/*
		p := sifxml.PurchaseOrder{RefId: "A", FormNumber: "B"}
		log.Printf("%+v\n", p)
		pj, _ := json.Marshal(p)
		log.Printf("%+v\n", string(pj))

		fmt.Println(string(purchase_order))
		m, err := mxj.NewMapXml(purchase_order)
		if err != nil {
			log.Fatal("err:", err.Error())
		}
		log.Printf("Purchase order 1, NewMapXml\n%+v\n", m)
		xmlValue, err := m.XmlIndent("", "  ")
		log.Printf("Purchase order 1, NewMapXml XmlIndent\n%+v\n", string(xmlValue))
		var result sifxml.PurchaseOrder
		log.Printf("Purchase order 1, Empty struct\n%+v\n", result)
		m0 := m["PurchaseOrder"].(map[string]interface{})
		j, err := json.Marshal(m0)
		log.Printf("Purchase order 1, NewMapXml JSON\n%+v\n", string(j))
		j1 := changeJSONTags(j)
		json.Unmarshal(j1, &result)
		log.Printf("Purchase order 1, NewMapXml JSON UnmarshalJSON\n%+v\n", result)
		x, _ := xml.MarshalIndent(result, "", "  ")
		log.Printf("Purchase order 1, NewMapXml JSON UnmarshalJSON MarshalXML\n%+v\n", string(x))
	*/
	/*
		var result1 sifxml.PurchaseOrder
		m01, _ := mxj.NewMapXml(purchase_order2)
		root0, _ := m01.Root()
		j0, err := json.Marshal(m01[root0].(map[string]interface{}))
		//m02 := m01["PurchaseOrder"].(map[string]interface{})
		//j, err = json.Marshal(m02)
		//j1 = changeJSONTags(j)
		json.Unmarshal(changeJSONTags(j0), &result1)
		//log.Printf("Purchase order 2, NewMapXML JSON\n%+v\n", string(j))
		//log.Printf("Purchase order 2, NewMapXML JSON changeJSONTages\n%+v\n", string(j1))
		log.Printf("Purchase order 2, NewMapXML JSON UnmarshalJSON struct\n%+v\n", result1)
		x, _ := xml.MarshalIndent(result1, "", "  ")
		log.Printf("Purchase order 2, NewMapXML JSON UnmarshalJSON MarshalXML\n%+v\n", string(x))
	*/

	m, _ := mxj.NewMapXml(staff_personal)
	x, _ := Map2SIFXML(m)
	log.Printf("Map2SIFXML\n%+v\n", string(x))
}
