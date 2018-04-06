package main

import (
	"github.com/nsip/nias3/webserver"
	//"github.com/nsip/nias3/xml2triples"
	//"log"
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
  <!--
  <ElectronicIdList>
    <ElectronicId Type="01">206655</ElectronicId>
  </ElectronicIdList>
  -->
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

func main() {
	/*
		refid := "D3E34F41-9D75-101A-8C3D-00AA001A1652"
		xml2triples.StoreXMLasDBtriples(staff_personal, refid)
		x, _ := xml2triples.DbTriples2XML(refid)
		log.Printf("Map2SIFXML\n%+v\n", string(x))
	*/
	webserver.SIFGetToDataStore("http://hits.nsip.edu.au/SIF3InfraREST/hits/requests/SchoolInfos?navigationPage=1&navigationPageSize=5&access_token=ZmZhODMzNjEtMGExOC00NDk5LTgyNjMtYjMwNjI4MGRjZDRlOmYxYzA1NjNhOWIzZTQyMGJiMDdkYTJkOTBkYjQ3OWVm&authenticationMethod=Basic")
	webserver.Webserver()
}
