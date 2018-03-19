package sifxml


    type StudentPersonal struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      AlertMessages AlertMessagesType `xml:"AlertMessages,omitempty" json:"AlertMessages"`
      MedicalAlertMessages MedicalAlertMessagesType `xml:"MedicalAlertMessages,omitempty" json:"MedicalAlertMessages"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      StateProvinceId StateProvinceIdType `xml:"StateProvinceId,omitempty" json:"StateProvinceId"`
      ElectronicIdList ElectronicIdListType `xml:"ElectronicIdList,omitempty" json:"ElectronicIdList"`
      OtherIdList OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      PersonInfo PersonInfoType `xml:"PersonInfo,omitempty" json:"PersonInfo"`
      ProjectedGraduationYear ProjectedGraduationYearType `xml:"ProjectedGraduationYear,omitempty" json:"ProjectedGraduationYear"`
      OnTimeGraduationYear OnTimeGraduationYearType `xml:"OnTimeGraduationYear,omitempty" json:"OnTimeGraduationYear"`
      GraduationDate GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate"`
      MostRecent StudentMostRecentContainerType `xml:"MostRecent,omitempty" json:"MostRecent"`
      AcceptableUsePolicy string `xml:"AcceptableUsePolicy,omitempty" json:"AcceptableUsePolicy"`
      GiftedTalented string `xml:"GiftedTalented,omitempty" json:"GiftedTalented"`
      EconomicDisadvantage string `xml:"EconomicDisadvantage,omitempty" json:"EconomicDisadvantage"`
      ESL string `xml:"ESL,omitempty" json:"ESL"`
      ESLDateAssessed string `xml:"ESLDateAssessed,omitempty" json:"ESLDateAssessed"`
      YoungCarersRole string `xml:"YoungCarersRole,omitempty" json:"YoungCarersRole"`
      Disability string `xml:"Disability,omitempty" json:"Disability"`
      IntegrationAide string `xml:"IntegrationAide,omitempty" json:"IntegrationAide"`
      EducationSupport string `xml:"EducationSupport,omitempty" json:"EducationSupport"`
      HomeSchooledStudent string `xml:"HomeSchooledStudent,omitempty" json:"HomeSchooledStudent"`
      Sensitive string `xml:"Sensitive,omitempty" json:"Sensitive"`
      OfflineDelivery string `xml:"OfflineDelivery,omitempty" json:"OfflineDelivery"`
      PrePrimaryEducation string `xml:"PrePrimaryEducation,omitempty" json:"PrePrimaryEducation"`
      FirstAUSchoolEnrollment string `xml:"FirstAUSchoolEnrollment,omitempty" json:"FirstAUSchoolEnrollment"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    