package sifxml


    type WellbeingResponse struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      Date string `xml:"Date,omitempty" json:"Date"`
      WellbeingResponseStartDate string `xml:"WellbeingResponseStartDate,omitempty" json:"WellbeingResponseStartDate"`
      WellbeingResponseEndDate string `xml:"WellbeingResponseEndDate,omitempty" json:"WellbeingResponseEndDate"`
      WellbeingResponseCategory string `xml:"WellbeingResponseCategory,omitempty" json:"WellbeingResponseCategory"`
      WellbeingResponseNotes string `xml:"WellbeingResponseNotes,omitempty" json:"WellbeingResponseNotes"`
      PersonInvolvementList PersonInvolvementListType `xml:"PersonInvolvementList,omitempty" json:"PersonInvolvementList"`
      Suspension SuspensionContainerType `xml:"Suspension,omitempty" json:"Suspension"`
      Detention DetentionContainerType `xml:"Detention,omitempty" json:"Detention"`
      PlanRequired PlanRequiredContainerType `xml:"PlanRequired,omitempty" json:"PlanRequired"`
      Award AwardContainerType `xml:"Award,omitempty" json:"Award"`
      OtherResponse OtherWellbeingResponseContainerType `xml:"OtherResponse,omitempty" json:"OtherResponse"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    