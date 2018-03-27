package sifxml


    type ChargedLocationInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      LocationType string `xml:"LocationType,omitempty" json:"LocationType"`
      SiteCategory string `xml:"SiteCategory,omitempty" json:"SiteCategory"`
      Name string `xml:"Name,omitempty" json:"Name"`
      Description string `xml:"Description,omitempty" json:"Description"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      ParentChargedLocationInfoRefId string `xml:"ParentChargedLocationInfoRefId,omitempty" json:"ParentChargedLocationInfoRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    