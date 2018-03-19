package sifxml


    type SectionInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolCourseInfoRefId IdRefType `xml:"SchoolCourseInfoRefId,omitempty" json:"SchoolCourseInfoRefId"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      TermInfoRefId IdRefType `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      MediumOfInstruction MediumOfInstructionType `xml:"MediumOfInstruction,omitempty" json:"MediumOfInstruction"`
      LanguageOfInstruction LanguageOfInstructionType `xml:"LanguageOfInstruction,omitempty" json:"LanguageOfInstruction"`
      LocationOfInstruction LocationOfInstructionType `xml:"LocationOfInstruction,omitempty" json:"LocationOfInstruction"`
      SummerSchool string `xml:"SummerSchool,omitempty" json:"SummerSchool"`
      SchoolCourseInfoOverride SchoolCourseInfoOverrideType `xml:"SchoolCourseInfoOverride,omitempty" json:"SchoolCourseInfoOverride"`
      CourseSectionCode string `xml:"CourseSectionCode,omitempty" json:"CourseSectionCode"`
      SectionCode string `xml:"SectionCode,omitempty" json:"SectionCode"`
      CountForAttendance string `xml:"CountForAttendance,omitempty" json:"CountForAttendance"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    