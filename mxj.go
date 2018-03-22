package main

import (
	"encoding/json"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/tidwall/sjson"
	"log"
	"regexp"
	"strconv"
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

func changeJSONTags(j string) string {
	return strings.Replace(j, `"#text":`, `"Value":`, -1)
}

func Map2SIFXML(m mxj.Map) ([]byte, error) {
	root, err := m.Root()
	if err != nil {
		return nil, err
	}
	//log.Println(root)
	m02 := m[root].(map[string]interface{})
	// log.Printf("m02\n%+v\n", m02)
	j, err := json.Marshal(m02)
	if err != nil {
		return nil, err
	}
	// log.Println(string(j))
	return Root2SIF(root, []byte(changeJSONTags(string(j))))
}

func put_triple(batch *leveldb.Batch, triple Triple) {
	log.Printf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))
	batch.Put([]byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))), []byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))))
	batch.Put([]byte(fmt.Sprintf("s:%s o:%s p:%s", strconv.Quote(triple.s), strconv.Quote(triple.o), strconv.Quote(triple.p))), []byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))))
	batch.Put([]byte(fmt.Sprintf("p:%s s:%s o:%s", strconv.Quote(triple.p), strconv.Quote(triple.s), strconv.Quote(triple.o))), []byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))))
	batch.Put([]byte(fmt.Sprintf("p:%s o:%s s:%s", strconv.Quote(triple.p), strconv.Quote(triple.o), strconv.Quote(triple.s))), []byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))))
	batch.Put([]byte(fmt.Sprintf("o:%s p:%s s:%s", strconv.Quote(triple.o), strconv.Quote(triple.p), strconv.Quote(triple.s))), []byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))))
	batch.Put([]byte(fmt.Sprintf("o:%s s:%s p:%s", strconv.Quote(triple.o), strconv.Quote(triple.s), strconv.Quote(triple.p))), []byte(fmt.Sprintf("s:%s p:%s o:%s", strconv.Quote(triple.s), strconv.Quote(triple.p), strconv.Quote(triple.o))))
}

func storeXMLasDBtriples(s []byte) (string, error) {
	db := GetDB()
	batch := new(leveldb.Batch)
	m, err := mxj.NewMapXml(s)
	if err != nil {
		return "", err
	}
	refid, err := m.ValueForPath("*.-RefId")
	if err != nil {
		return "", err
	}
	for _, n := range m.LeafNodes() {
		put_triple(batch, Triple{s: fmt.Sprintf("%v", refid), p: n.Path, o: fmt.Sprintf("%v", n.Value)})
	}
	batcherr := db.Write(batch, nil)
	if batcherr != nil {
		return "", err
	}
	batch.Reset()
	return refid.(string), nil
}

var mxj2sjsonPathRe1 = regexp.MustCompile(`\[(\d+)\]`)
var mxj2sjsonPathRe2 = regexp.MustCompile(`\.#text$`)

func mxj2sjsonPath(p string) string {
	return mxj2sjsonPathRe1.ReplaceAllString(
		mxj2sjsonPathRe2.ReplaceAllString(p, ".Value"), ".$1")
}

func dbTriples2XML(refid string) ([]byte, error) {
	triple_strings := getIdentifiers(fmt.Sprintf("s:%s ", strconv.Quote(fmt.Sprintf("%v", refid))))
	triples := parseTriples(triple_strings)
	json := ""
	var err error
	for _, t := range triples {
		//log.Printf("%s %s %s\n", t.s, t.p, t.o)
		//log.Printf("%s %s %s\n", t.s, mxj2sjsonPath(t.p), t.o)
		json, err = sjson.Set(json, mxj2sjsonPath(t.p), t.o)
		if err != nil {
			return nil, err
		}
	}
	//log.Printf("%+v\n", json)
	mm, err := mxj.NewMapJson([]byte(json))
	if err != nil {
		return nil, err
	}
	// log.Printf("%+v\n", mm)
	return Map2SIFXML(mm)
}

func main() {
	refid, _ := storeXMLasDBtriples(staff_personal)
	x, _ := dbTriples2XML(refid)
	log.Printf("Map2SIFXML\n%+v\n", string(x))
}
