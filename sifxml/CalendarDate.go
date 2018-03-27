package sifxml


    type CalendarDate struct {
        CalendarDateRefId string `xml:"CalendarDateRefId,attr" json:"-CalendarDateRefId"`
      Date string `xml:"Date,omitempty" json:"Date"`
      CalendarSummaryRefId string `xml:"CalendarSummaryRefId,omitempty" json:"CalendarSummaryRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      CalendarDateType CalendarDateInfoType `xml:"CalendarDateType,omitempty" json:"CalendarDateType"`
      CalendarDateNumber string `xml:"CalendarDateNumber,omitempty" json:"CalendarDateNumber"`
      StudentAttendance AttendanceInfoType `xml:"StudentAttendance,omitempty" json:"StudentAttendance"`
      TeacherAttendance AttendanceInfoType `xml:"TeacherAttendance,omitempty" json:"TeacherAttendance"`
      AdministratorAttendance AttendanceInfoType `xml:"AdministratorAttendance,omitempty" json:"AdministratorAttendance"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    