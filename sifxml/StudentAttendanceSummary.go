package sifxml


    type StudentAttendanceSummary struct {
        StudentAttendanceSummaryRefId IdRefType `xml:"StudentAttendanceSummaryRefId,attr" json:"-StudentAttendanceSummaryRefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      StartDate string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate string `xml:"EndDate,omitempty" json:"EndDate"`
      StartDay string `xml:"StartDay,omitempty" json:"StartDay"`
      EndDay string `xml:"EndDay,omitempty" json:"EndDay"`
      FTE string `xml:"FTE,omitempty" json:"FTE"`
      DaysAttended string `xml:"DaysAttended,omitempty" json:"DaysAttended"`
      ExcusedAbsences string `xml:"ExcusedAbsences,omitempty" json:"ExcusedAbsences"`
      UnexcusedAbsences string `xml:"UnexcusedAbsences,omitempty" json:"UnexcusedAbsences"`
      DaysTardy string `xml:"DaysTardy,omitempty" json:"DaysTardy"`
      DaysInMembership string `xml:"DaysInMembership,omitempty" json:"DaysInMembership"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    