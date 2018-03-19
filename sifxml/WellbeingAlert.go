package sifxml


    type WellbeingAlert struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      Date string `xml:"Date,omitempty" json:"Date"`
      WellbeingAlertStartDate string `xml:"WellbeingAlertStartDate,omitempty" json:"WellbeingAlertStartDate"`
      WellbeingAlertEndDate string `xml:"WellbeingAlertEndDate,omitempty" json:"WellbeingAlertEndDate"`
      WellbeingAlertCategory string `xml:"WellbeingAlertCategory,omitempty" json:"WellbeingAlertCategory"`
      WellbeingAlertDescription string `xml:"WellbeingAlertDescription,omitempty" json:"WellbeingAlertDescription"`
      EnrolmentRestricted string `xml:"EnrolmentRestricted,omitempty" json:"EnrolmentRestricted"`
      AlertAudience string `xml:"AlertAudience,omitempty" json:"AlertAudience"`
      AlertSeverity string `xml:"AlertSeverity,omitempty" json:"AlertSeverity"`
      AlertKeyContact string `xml:"AlertKeyContact,omitempty" json:"AlertKeyContact"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    