package sifxml


    type SchoolInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      CommonwealthId string `xml:"CommonwealthId,omitempty" json:"CommonwealthId"`
      ACARAId string `xml:"ACARAId,omitempty" json:"ACARAId"`
      OtherIdList OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      SchoolName string `xml:"SchoolName,omitempty" json:"SchoolName"`
      LEAInfoRefId RefIdType `xml:"LEAInfoRefId,omitempty" json:"LEAInfoRefId"`
      OtherLEA SchoolInfo_OtherLEA `xml:"OtherLEA,omitempty" json:"OtherLEA"`
      SchoolDistrict string `xml:"SchoolDistrict,omitempty" json:"SchoolDistrict"`
      SchoolDistrictLocalId LocalIdType `xml:"SchoolDistrictLocalId,omitempty" json:"SchoolDistrictLocalId"`
      SchoolType string `xml:"SchoolType,omitempty" json:"SchoolType"`
      SchoolFocusList SchoolFocusListType `xml:"SchoolFocusList,omitempty" json:"SchoolFocusList"`
      SchoolURL SchoolURLType `xml:"SchoolURL,omitempty" json:"SchoolURL"`
      SchoolEmailList EmailListType `xml:"SchoolEmailList,omitempty" json:"SchoolEmailList"`
      PrincipalInfo PrincipalInfoType `xml:"PrincipalInfo,omitempty" json:"PrincipalInfo"`
      SchoolContactList SchoolContactListType `xml:"SchoolContactList,omitempty" json:"SchoolContactList"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      SessionType string `xml:"SessionType,omitempty" json:"SessionType"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      ARIA string `xml:"ARIA,omitempty" json:"ARIA"`
      OperationalStatus OperationalStatusType `xml:"OperationalStatus,omitempty" json:"OperationalStatus"`
      FederalElectorate string `xml:"FederalElectorate,omitempty" json:"FederalElectorate"`
      Campus CampusContainerType `xml:"Campus,omitempty" json:"Campus"`
      SchoolSector string `xml:"SchoolSector,omitempty" json:"SchoolSector"`
      IndependentSchool string `xml:"IndependentSchool,omitempty" json:"IndependentSchool"`
      NonGovSystemicStatus string `xml:"NonGovSystemicStatus,omitempty" json:"NonGovSystemicStatus"`
      System string `xml:"System,omitempty" json:"System"`
      ReligiousAffiliation string `xml:"ReligiousAffiliation,omitempty" json:"ReligiousAffiliation"`
      SchoolGeographicLocation string `xml:"SchoolGeographicLocation,omitempty" json:"SchoolGeographicLocation"`
      LocalGovernmentArea string `xml:"LocalGovernmentArea,omitempty" json:"LocalGovernmentArea"`
      JurisdictionLowerHouse string `xml:"JurisdictionLowerHouse,omitempty" json:"JurisdictionLowerHouse"`
      SLA string `xml:"SLA,omitempty" json:"SLA"`
      SchoolCoEdStatus string `xml:"SchoolCoEdStatus,omitempty" json:"SchoolCoEdStatus"`
      BoardingSchoolStatus string `xml:"BoardingSchoolStatus,omitempty" json:"BoardingSchoolStatus"`
      YearLevelEnrollmentList YearLevelEnrollmentListType `xml:"YearLevelEnrollmentList,omitempty" json:"YearLevelEnrollmentList"`
      TotalEnrollments TotalEnrollmentsType `xml:"TotalEnrollments,omitempty" json:"TotalEnrollments"`
      Entity_Open string `xml:"Entity_Open,omitempty" json:"Entity_Open"`
      Entity_Close string `xml:"Entity_Close,omitempty" json:"Entity_Close"`
      SchoolGroupList SchoolGroupListType `xml:"SchoolGroupList,omitempty" json:"SchoolGroupList"`
      SchoolTimeZone string `xml:"SchoolTimeZone,omitempty" json:"SchoolTimeZone"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type SchoolInfo_OtherLEA struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value RefIdType `xml:",chardata" json:"Value"`
}
