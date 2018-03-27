package sifxml


    type GradingAssignmentScore struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentPersonalLocalId LocalIdType `xml:"StudentPersonalLocalId,omitempty" json:"StudentPersonalLocalId"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      GradingAssignmentRefId string `xml:"GradingAssignmentRefId,omitempty" json:"GradingAssignmentRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      DateGraded string `xml:"DateGraded,omitempty" json:"DateGraded"`
      ExpectedScore string `xml:"ExpectedScore,omitempty" json:"ExpectedScore"`
      ScorePoints string `xml:"ScorePoints,omitempty" json:"ScorePoints"`
      ScorePercent string `xml:"ScorePercent,omitempty" json:"ScorePercent"`
      ScoreLetter string `xml:"ScoreLetter,omitempty" json:"ScoreLetter"`
      ScoreDescription string `xml:"ScoreDescription,omitempty" json:"ScoreDescription"`
      SubscoreList NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList"`
      TeacherJudgement string `xml:"TeacherJudgement,omitempty" json:"TeacherJudgement"`
      MarkInfoRefId string `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId"`
      AssignmentScoreIteration string `xml:"AssignmentScoreIteration,omitempty" json:"AssignmentScoreIteration"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    