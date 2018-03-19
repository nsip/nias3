package sifxml


    type LEAInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId string `xml:"CommonwealthId,omitempty" json:"CommonwealthId"`
      LEAName string `xml:"LEAName,omitempty" json:"LEAName"`
      LEAURL string `xml:"LEAURL,omitempty" json:"LEAURL"`
      EducationAgencyType AgencyType `xml:"EducationAgencyType,omitempty" json:"EducationAgencyType"`
      LEAContactList LEAContactListType `xml:"LEAContactList,omitempty" json:"LEAContactList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      OperationalStatus OperationalStatusType `xml:"OperationalStatus,omitempty" json:"OperationalStatus"`
      JurisdictionLowerHouse string `xml:"JurisdictionLowerHouse,omitempty" json:"JurisdictionLowerHouse"`
      SLA string `xml:"SLA,omitempty" json:"SLA"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    