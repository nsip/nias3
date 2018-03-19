package sifxml


    type StudentPeriodAttendance struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      Date string `xml:"Date,omitempty" json:"Date"`
      SessionInfoRefId IdRefType `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId"`
      ScheduledActivityRefId IdRefType `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimetablePeriod string `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod"`
      TimeIn string `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut string `xml:"TimeOut,omitempty" json:"TimeOut"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode,omitempty" json:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus,omitempty" json:"AttendanceStatus"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      AuditInfo AuditInfoType `xml:"AuditInfo,omitempty" json:"AuditInfo"`
      AttendanceComment string `xml:"AttendanceComment,omitempty" json:"AttendanceComment"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    