package sifxml


    type VendorInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Name string `xml:"Name,omitempty" json:"Name"`
      ContactInfo ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo"`
      CustomerId string `xml:"CustomerId,omitempty" json:"CustomerId"`
      ABN string `xml:"ABN,omitempty" json:"ABN"`
      RegisteredForGST string `xml:"RegisteredForGST,omitempty" json:"RegisteredForGST"`
      PaymentTerms string `xml:"PaymentTerms,omitempty" json:"PaymentTerms"`
      BPay string `xml:"BPay,omitempty" json:"BPay"`
      BSB string `xml:"BSB,omitempty" json:"BSB"`
      AccountNumber string `xml:"AccountNumber,omitempty" json:"AccountNumber"`
      AccountName string `xml:"AccountName,omitempty" json:"AccountName"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    