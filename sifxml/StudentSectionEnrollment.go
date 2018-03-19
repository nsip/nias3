package sifxml


    type StudentSectionEnrollment struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SectionInfoRefId IdRefType `xml:"SectionInfoRefId,omitempty" json:"SectionInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      EntryDate string `xml:"EntryDate,omitempty" json:"EntryDate"`
      ExitDate string `xml:"ExitDate,omitempty" json:"ExitDate"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    