package sifxml


    type ReportAuthorityInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      AuthorityName string `xml:"AuthorityName,omitempty" json:"AuthorityName"`
      AuthorityId string `xml:"AuthorityId,omitempty" json:"AuthorityId"`
      AuthorityDepartment string `xml:"AuthorityDepartment,omitempty" json:"AuthorityDepartment"`
      AuthorityLevel string `xml:"AuthorityLevel,omitempty" json:"AuthorityLevel"`
      ContactInfo ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo"`
      Address AddressType `xml:"Address,omitempty" json:"Address"`
      PhoneNumber PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    