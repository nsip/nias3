package sifxml


    type Activity struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Title string `xml:"Title,omitempty" json:"Title"`
      Preamble string `xml:"Preamble,omitempty" json:"Preamble"`
      TechnicalRequirements TechnicalRequirementsType `xml:"TechnicalRequirements,omitempty" json:"TechnicalRequirements"`
      SoftwareRequirementList SoftwareRequirementListType `xml:"SoftwareRequirementList,omitempty" json:"SoftwareRequirementList"`
      EssentialMaterials EssentialMaterialsType `xml:"EssentialMaterials,omitempty" json:"EssentialMaterials"`
      LearningObjectives LearningObjectivesType `xml:"LearningObjectives,omitempty" json:"LearningObjectives"`
      LearningStandards LearningStandardsType `xml:"LearningStandards,omitempty" json:"LearningStandards"`
      SubjectArea SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      Prerequisites PrerequisitesType `xml:"Prerequisites,omitempty" json:"Prerequisites"`
      Students StudentsType `xml:"Students,omitempty" json:"Students"`
      SourceObjects SourceObjectsType `xml:"SourceObjects,omitempty" json:"SourceObjects"`
      Points string `xml:"Points,omitempty" json:"Points"`
      ActivityTime ActivityTimeType `xml:"ActivityTime,omitempty" json:"ActivityTime"`
      AssessmentRefId IdRefType `xml:"AssessmentRefId,omitempty" json:"AssessmentRefId"`
      MaxAttemptsAllowed string `xml:"MaxAttemptsAllowed,omitempty" json:"MaxAttemptsAllowed"`
      ActivityWeight string `xml:"ActivityWeight,omitempty" json:"ActivityWeight"`
      Evaluation Activity_Evaluation `xml:"Evaluation,omitempty" json:"Evaluation"`
      LearningResources LearningResourcesType `xml:"LearningResources,omitempty" json:"LearningResources"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Activity_Evaluation struct {
      EvaluationType string `xml:"EvaluationType,attr" json:"-EvaluationType"`
       Description string `xml:"Description,omitempty" json:"Description"`
}
