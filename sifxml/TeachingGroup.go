package sifxml


    type TeachingGroup struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      ShortName string `xml:"ShortName,omitempty" json:"ShortName"`
      LongName string `xml:"LongName,omitempty" json:"LongName"`
      GroupType string `xml:"GroupType,omitempty" json:"GroupType"`
      Set string `xml:"Set,omitempty" json:"Set"`
      Block string `xml:"Block,omitempty" json:"Block"`
      CurriculumLevel string `xml:"CurriculumLevel,omitempty" json:"CurriculumLevel"`
      SchoolInfoRefId RefIdType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolCourseInfoRefId RefIdType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId"`
      SchoolCourseLocalId LocalIdType `xml:"SchoolCourseLocalId,omitempty" json:"SchoolCourseLocalId"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TimeTableSubjectLocalId LocalIdType `xml:"TimeTableSubjectLocalId,omitempty" json:"TimeTableSubjectLocalId"`
      Semester string `xml:"Semester,omitempty" json:"Semester"`
      StudentList StudentListType `xml:"StudentList,omitempty" json:"StudentList"`
      TeacherList TeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      MinClassSize string `xml:"MinClassSize,omitempty" json:"MinClassSize"`
      MaxClassSize string `xml:"MaxClassSize,omitempty" json:"MaxClassSize"`
      TeachingGroupPeriodList TeachingGroupPeriodListType `xml:"TeachingGroupPeriodList,omitempty" json:"TeachingGroupPeriodList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    