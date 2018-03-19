package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/clbanning/mxj"
	"github.com/mitchellh/mapstructure"
	"github.com/nsip/nias3/sifxml"
	"log"
	"reflect"
	"regexp"
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

type PurchaseOrder1 struct {
	RefId                      sifxml.RefIdType                `xml:"RefId,attr" json:"-RefId"`
	FormNumber                 string                          `xml:"FormNumber,omitempty" json:"FormNumber"`
	VendorInfoRefId            sifxml.IdRefType                `xml:"VendorInfoRefId,omitempty" json:"VendorInfoRefId"`
	ChargedLocationInfoRefId   sifxml.IdRefType                `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
	EmployeePersonalRefId      sifxml.IdRefType                `xml:"EmployeePersonalRefId,omitempty" json:"EmployeePersonalRefId"`
	PurchasingItems            PurchasingItemsType             `xml:"PurchasingItems,omitempty" json:"PurchasingItems"`
	CreationDate               string                          `xml:"CreationDate,omitempty" json:"CreationDate"`
	TaxRate                    string                          `xml:"TaxRate,omitempty" json:"TaxRate"`
	TaxAmount                  MonetaryAmountType              `xml:"TaxAmount,omitempty" json:"TaxAmount"`
	TotalAmount                MonetaryAmountType              `xml:"TotalAmount,omitempty" json:"TotalAmount"`
	UpdateDate                 string                          `xml:"UpdateDate,omitempty" json:"UpdateDate"`
	FullyDelivered             string                          `xml:"FullyDelivered,omitempty" json:"FullyDelivered"`
	OriginalPurchaseOrderRefId sifxml.IdRefType                `xml:"OriginalPurchaseOrderRefId,omitempty" json:"OriginalPurchaseOrderRefId"`
	SIF_Metadata               SIF_MetadataType                `xml:"SIF_Metadata,omitempty"`
	SIF_ExtendedElements       sifxml.SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty"`
}

type PurchasingItemsType struct {
	PurchasingItem []PurchasingItemType `xml:"PurchasingItem" json:"PurchasingItem"`
}

// http://choly.ca/post/go-json-marshalling/
func (this *PurchasingItemsType) UnmarshalJSON(data []byte) error {
	type Alias PurchasingItemsType
	d := json.NewDecoder(bytes.NewBuffer(data))
	// "PurchasingItem":[

	t, err := d.Token()
	t, err = d.Token()
	t, err = d.Token()
	if err != nil {
		return err
	}
	log.Println(t)
	if t == json.Delim('[') {
		aux := &struct{ *Alias }{Alias: (*Alias)(this)}
		if err := json.Unmarshal(data, &aux); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	re := regexp.MustCompile(`^(\s*\{"[^"]+"\s*:)`)
	re2 := regexp.MustCompile(`\}\s*$`)
	return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}

type PurchasingItemType struct {
	ItemNumber        string              `xml:"ItemNumber,omitempty" json:"ItemNumber"`
	ItemDescription   string              `xml:"ItemDescription,omitempty" json:"ItemDescription" `
	Quantity          string              `xml:"Quantity,omitempty" json:"Quantity"`
	UnitCost          MonetaryAmountType  `xml:"UnitCost,omitempty" json:"UnitCost"`
	TotalCost         MonetaryAmountType  `xml:"TotalCost,omitempty" json:"TotalCost"`
	QuantityDelivered string              `xml:"QuantityDelivered,omitempty" json:"QuantityDelivered"`
	CancelledOrder    string              `xml:"CancelledOrder,omitempty" json:"CancelledOrder"`
	TaxRate           string              `xml:"TaxRate,omitempty" json:"TaxRate"`
	ExpenseAccounts   ExpenseAccountsType `xml:"ExpenseAccounts,omitempty" json:"ExpenseAccounts"`
}

type MonetaryAmountType struct {
	Currency string `xml:"Currency,attr" json:"-Currency"`
	Value    string `xml:",chardata", json:"Value"`
}

type SIF_MetadataType struct {
	TimeElements    sifxml.SIF_MetadataType_TimeElements `xml:"TimeElements"`
	LifeCycle       sifxml.LifeCycleType                 `xml:"LifeCycle"`
	EducationFilter sifxml.EducationFilterType           `xml:"EducationFilter"`
}

type ExpenseAccountsType struct {
	ExpenseAccount []ExpenseAccountType `xml:"ExpenseAccount" json:"ExpenseAccount"`
}

var re = regexp.MustCompile(`^(\s*\{"[^"]+"\s*:)`)
var re2 = regexp.MustCompile(`\}\s*$`)

func (this *ExpenseAccountsType) UnmarshalJSON(data []byte) error {
	type Alias ExpenseAccountsType
	d := json.NewDecoder(bytes.NewBuffer(data))
	t, err := d.Token()
	t, err = d.Token()
	t, err = d.Token()
	if err != nil {
		return err
	}
	log.Println(t)
	if t == json.Delim('[') {
		aux := &struct{ *Alias }{Alias: (*Alias)(this)}
		if err := json.Unmarshal(data, &aux); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}

type ExpenseAccountType struct {
	AccountCode           string             `xml:"AccountCode,omitempty" json:"AccountCode"`
	Amount                MonetaryAmountType `xml:"Amount,omitempty" json:"Amount"`
	FinancialAccountRefId sifxml.IdRefType   `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
	AccountingPeriod      sifxml.LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
}

func mxj2struct(m mxj.Map, result interface{}) error {
	o := rename_map(m)
	log.Printf("%+v\n", o)
	return mapstructure.Decode(m, &result)
}

func rename_map(m mxj.Map) interface{} {
	o := mxj.Map(m)
	for k, v := range m {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			v = rename_map(v.(mxj.Map))
		}
		if strings.HasPrefix(k, "-") {
			o[strings.TrimPrefix(k, "-")] = v
			delete(o, k)
		}
		if k == "#text" {
			o["Value"] = v
			delete(o, k)
		}
	}
	return o
}

func changeJSONTags(j []byte) []byte {
	return []byte(strings.Replace(string(j), `"#text":`, `"Value":`, -1))
}

func main() {
	p := PurchaseOrder1{RefId: "A", FormNumber: "B"}
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
	var result PurchaseOrder1
	log.Printf("Purchase order 1, Empty struct\n%+v\n", result)
	/*
		err = mxj2struct(m, &result)
		if err != nil {
			log.Fatal("err:", err.Error())
		}
		log.Printf("%+v\n", result)
	*/
	m0 := m["PurchaseOrder"].(map[string]interface{})
	j, err := json.Marshal(m0)
	log.Printf("Purchase order 1, NewMapXml JSON\n%+v\n", string(j))
	j1 := changeJSONTags(j)
	json.Unmarshal(j1, &result)
	/*
		err = mxj.Map(m["PurchaseOrder"].(map[string]interface{})).Struct(&result)
		if err != nil {
			log.Fatal("err:", err.Error())
		}
	*/
	log.Printf("Purchase order 1, NewMapXml JSON UnmarshalJSON\n%+v\n", result)
	x, _ := xml.MarshalIndent(result, "", "  ")
	log.Printf("Purchase order 1, NewMapXml JSON UnmarshalJSON MarshalXML\n%+v\n", string(x))

	m01, err := mxj.NewMapXml(purchase_order2)
	root, err := m01.Root()
	log.Println(root)
	//j, err = json.Marshal(m01[root].(map[string]interface{}))
	m02 := m01["PurchaseOrder"].(map[string]interface{})
	j, err = json.Marshal(m02)
	j1 = changeJSONTags(j)
	json.Unmarshal(j1, &result)
	log.Printf("Purchase order 2, NewMapXML JSON\n%+v\n", string(j))
	log.Printf("Purchase order 2, NewMapXML JSON changeJSONTages\n%+v\n", string(j1))
	log.Printf("Purchase order 2, NewMapXML JSON UnmarshalJSON struct\n%+v\n", result)
	x, _ = xml.MarshalIndent(result, "", "  ")
	log.Printf("Purchase order 2, NewMapXML JSON UnmarshalJSON MarshalXML\n%+v\n", string(x))

}
