package sifxml


    type StudentSchoolEnrollment struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      MembershipType string `xml:"MembershipType,omitempty" json:"MembershipType"`
      TimeFrame string `xml:"TimeFrame,omitempty" json:"TimeFrame"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      EntryDate string `xml:"EntryDate,omitempty" json:"EntryDate"`
      EntryType StudentEntryContainerType `xml:"EntryType,omitempty" json:"EntryType"`
      YearLevel YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      Homeroom StudentSchoolEnrollment_Homeroom `xml:"Homeroom,omitempty" json:"Homeroom"`
      Advisor StudentSchoolEnrollment_Advisor `xml:"Advisor,omitempty" json:"Advisor"`
      Counselor StudentSchoolEnrollment_Counselor `xml:"Counselor,omitempty" json:"Counselor"`
      Homegroup string `xml:"Homegroup,omitempty" json:"Homegroup"`
      ACARASchoolId LocalIdType `xml:"ACARASchoolId,omitempty" json:"ACARASchoolId"`
      ClassCode string `xml:"ClassCode,omitempty" json:"ClassCode"`
      TestLevel YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel"`
      ReportingSchool string `xml:"ReportingSchool,omitempty" json:"ReportingSchool"`
      House string `xml:"House,omitempty" json:"House"`
      IndividualLearningPlan string `xml:"IndividualLearningPlan,omitempty" json:"IndividualLearningPlan"`
      Calendar StudentSchoolEnrollment_Calendar `xml:"Calendar,omitempty" json:"Calendar"`
      ExitDate string `xml:"ExitDate,omitempty" json:"ExitDate"`
      ExitStatus StudentExitStatusContainerType `xml:"ExitStatus,omitempty" json:"ExitStatus"`
      ExitType StudentExitContainerType `xml:"ExitType,omitempty" json:"ExitType"`
      FTE string `xml:"FTE,omitempty" json:"FTE"`
      FTPTStatus string `xml:"FTPTStatus,omitempty" json:"FTPTStatus"`
      FFPOS string `xml:"FFPOS,omitempty" json:"FFPOS"`
      CatchmentStatus CatchmentStatusContainerType `xml:"CatchmentStatus,omitempty" json:"CatchmentStatus"`
      RecordClosureReason string `xml:"RecordClosureReason,omitempty" json:"RecordClosureReason"`
      PromotionInfo PromotionInfoType `xml:"PromotionInfo,omitempty" json:"PromotionInfo"`
      PreviousSchool LocalIdType `xml:"PreviousSchool,omitempty" json:"PreviousSchool"`
      DestinationSchool LocalIdType `xml:"DestinationSchool,omitempty" json:"DestinationSchool"`
      StudentSubjectChoiceList StudentSubjectChoiceListType `xml:"StudentSubjectChoiceList,omitempty" json:"StudentSubjectChoiceList"`
      StartedAtSchoolDate string `xml:"StartedAtSchoolDate,omitempty" json:"StartedAtSchoolDate"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type StudentSchoolEnrollment_Homeroom struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
type StudentSchoolEnrollment_Advisor struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
type StudentSchoolEnrollment_Counselor struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
type StudentSchoolEnrollment_Calendar struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
