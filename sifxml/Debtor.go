package sifxml


    type Debtor struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      BilledEntity Debtor_BilledEntity `xml:"BilledEntity,omitempty" json:"BilledEntity"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      BillingName string `xml:"BillingName,omitempty" json:"BillingName"`
      BillingNote string `xml:"BillingNote,omitempty" json:"BillingNote"`
      Discount string `xml:"Discount,omitempty" json:"Discount"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Debtor_BilledEntity struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
