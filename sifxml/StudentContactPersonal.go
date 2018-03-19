package sifxml


    type StudentContactPersonal struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      OtherIdList OtherIdListType `xml:"OtherIdList,omitempty" json:"OtherIdList"`
      PersonInfo PersonInfoType `xml:"PersonInfo,omitempty" json:"PersonInfo"`
      EmploymentType string `xml:"EmploymentType,omitempty" json:"EmploymentType"`
      SchoolEducationalLevel EducationalLevelType `xml:"SchoolEducationalLevel,omitempty" json:"SchoolEducationalLevel"`
      NonSchoolEducation string `xml:"NonSchoolEducation,omitempty" json:"NonSchoolEducation"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    