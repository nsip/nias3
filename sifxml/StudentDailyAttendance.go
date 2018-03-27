package sifxml


    type StudentDailyAttendance struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      Date string `xml:"Date,omitempty" json:"Date"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      DayValue string `xml:"DayValue,omitempty" json:"DayValue"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode,omitempty" json:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus,omitempty" json:"AttendanceStatus"`
      TimeIn string `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut string `xml:"TimeOut,omitempty" json:"TimeOut"`
      AbsenceValue string `xml:"AbsenceValue,omitempty" json:"AbsenceValue"`
      AttendanceNote string `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    