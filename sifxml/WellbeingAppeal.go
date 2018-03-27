package sifxml


    type WellbeingAppeal struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      WellbeingResponseRefId string `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId"`
      LocalAppealId LocalIdType `xml:"LocalAppealId,omitempty" json:"LocalAppealId"`
      AppealStatusCode string `xml:"AppealStatusCode,omitempty" json:"AppealStatusCode"`
      Date string `xml:"Date,omitempty" json:"Date"`
      AppealNotes string `xml:"AppealNotes,omitempty" json:"AppealNotes"`
      AppealOutcome string `xml:"AppealOutcome,omitempty" json:"AppealOutcome"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    