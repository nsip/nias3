package sifxml


    type TimeTableSubject struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      AcademicYear YearLevelType `xml:"AcademicYear,omitempty" json:"AcademicYear"`
      AcademicYearRange YearRangeType `xml:"AcademicYearRange,omitempty" json:"AcademicYearRange"`
      CourseLocalId LocalIdType `xml:"CourseLocalId,omitempty" json:"CourseLocalId"`
      SchoolCourseInfoRefId RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId"`
      Faculty string `xml:"Faculty,omitempty" json:"Faculty"`
      SubjectShortName string `xml:"SubjectShortName,omitempty" json:"SubjectShortName"`
      SubjectLongName string `xml:"SubjectLongName,omitempty" json:"SubjectLongName"`
      SubjectType string `xml:"SubjectType,omitempty" json:"SubjectType"`
      ProposedMaxClassSize string `xml:"ProposedMaxClassSize,omitempty" json:"ProposedMaxClassSize"`
      ProposedMinClassSize string `xml:"ProposedMinClassSize,omitempty" json:"ProposedMinClassSize"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      Semester string `xml:"Semester,omitempty" json:"Semester"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    