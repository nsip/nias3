package sifxml


    type StudentGrade struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId IdRefType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      Homegroup string `xml:"Homegroup,omitempty" json:"Homegroup"`
      YearLevel YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      TeachingGroupShortName string `xml:"TeachingGroupShortName,omitempty" json:"TeachingGroupShortName"`
      TeachingGroupRefId IdRefType `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      StaffPersonalRefId IdRefType `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      Markers StudentGradeMarkersListType `xml:"Markers,omitempty" json:"Markers"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      TermInfoRefId IdRefType `xml:"TermInfoRefId,omitempty" json:"TermInfoRefId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      LearningArea ACStrandSubjectAreaType `xml:"LearningArea,omitempty" json:"LearningArea"`
      LearningStandardList LearningStandardListType `xml:"LearningStandardList,omitempty" json:"LearningStandardList"`
      GradingScoreList GradingScoreListType `xml:"GradingScoreList,omitempty" json:"GradingScoreList"`
      Grade GradeType `xml:"Grade,omitempty" json:"Grade"`
      TeacherJudgement string `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    