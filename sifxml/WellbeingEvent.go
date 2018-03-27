package sifxml


    type WellbeingEvent struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      EventId LocalIdType `xml:"EventId,omitempty" json:"EventId"`
      WellbeingEventNotes string `xml:"WellbeingEventNotes,omitempty" json:"WellbeingEventNotes"`
      WellbeingEventCategoryClass string `xml:"WellbeingEventCategoryClass,omitempty" json:"WellbeingEventCategoryClass"`
      WellbeingEventCategoryList WellbeingEventCategoryListType `xml:"WellbeingEventCategoryList,omitempty" json:"WellbeingEventCategoryList"`
      ReportingStaffRefId string `xml:"ReportingStaffRefId,omitempty" json:"ReportingStaffRefId"`
      WellbeingEventLocationDetails WellbeingEventLocationDetailsType `xml:"WellbeingEventLocationDetails,omitempty" json:"WellbeingEventLocationDetails"`
      WellbeingEventCreationTimeStamp string `xml:"WellbeingEventCreationTimeStamp,omitempty" json:"WellbeingEventCreationTimeStamp"`
      WellbeingEventDate string `xml:"WellbeingEventDate,omitempty" json:"WellbeingEventDate"`
      WellbeingEventTime string `xml:"WellbeingEventTime,omitempty" json:"WellbeingEventTime"`
      WellbeingEventDescription string `xml:"WellbeingEventDescription,omitempty" json:"WellbeingEventDescription"`
      WellbeingEventTimePeriod string `xml:"WellbeingEventTimePeriod,omitempty" json:"WellbeingEventTimePeriod"`
      ConfidentialFlag string `xml:"ConfidentialFlag,omitempty" json:"ConfidentialFlag"`
      PersonInvolvementList PersonInvolvementListType `xml:"PersonInvolvementList,omitempty" json:"PersonInvolvementList"`
      FollowUpActionList FollowUpActionListType `xml:"FollowUpActionList,omitempty" json:"FollowUpActionList"`
      Status string `xml:"Status,omitempty" json:"Status"`
      DocumentList WellbeingDocumentListType `xml:"DocumentList,omitempty" json:"DocumentList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    