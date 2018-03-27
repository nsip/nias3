package sifxml


    type SchoolCourseInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      TermInfoRefId string `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      CourseCode string `xml:"CourseCode,omitempty" json:"CourseCode"`
      StateCourseCode string `xml:"StateCourseCode,omitempty" json:"StateCourseCode"`
      DistrictCourseCode string `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode"`
      SubjectAreaList SubjectAreaListType `xml:"SubjectAreaList,omitempty" json:"SubjectAreaList"`
      CourseTitle string `xml:"CourseTitle,omitempty" json:"CourseTitle"`
      Description string `xml:"Description,omitempty" json:"Description"`
      InstructionalLevel string `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel"`
      CourseCredits string `xml:"CourseCredits,omitempty" json:"CourseCredits"`
      CoreAcademicCourse string `xml:"CoreAcademicCourse,omitempty" json:"CoreAcademicCourse"`
      GraduationRequirement string `xml:"GraduationRequirement,omitempty" json:"GraduationRequirement"`
      Department string `xml:"Department,omitempty" json:"Department"`
      CourseContent string `xml:"CourseContent,omitempty" json:"CourseContent"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    