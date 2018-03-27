package sifxml


    type LearningStandardDocument struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Title string `xml:"Title,omitempty" json:"Title"`
      Description string `xml:"Description,omitempty" json:"Description"`
      RichDescription AbstractContentElementType `xml:"RichDescription,omitempty" json:"RichDescription"`
      Source string `xml:"Source,omitempty" json:"Source"`
      Organizations OrganizationsType `xml:"Organizations,omitempty" json:"Organizations"`
      Authors AuthorsType `xml:"Authors,omitempty" json:"Authors"`
      OrganizationContactPoint string `xml:"OrganizationContactPoint,omitempty" json:"OrganizationContactPoint"`
      SubjectAreas ACStrandAreaListType `xml:"SubjectAreas,omitempty" json:"SubjectAreas"`
      DocumentStatus string `xml:"DocumentStatus,omitempty" json:"DocumentStatus"`
      DocumentDate string `xml:"DocumentDate,omitempty" json:"DocumentDate"`
      LocalAdoptionDate string `xml:"LocalAdoptionDate,omitempty" json:"LocalAdoptionDate"`
      LocalArchiveDate string `xml:"LocalArchiveDate,omitempty" json:"LocalArchiveDate"`
      EndOfLifeDate string `xml:"EndOfLifeDate,omitempty" json:"EndOfLifeDate"`
      Copyright CopyRightContainerType `xml:"Copyright,omitempty" json:"Copyright"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      RepositoryDate string `xml:"RepositoryDate,omitempty" json:"RepositoryDate"`
      LearningStandardItemRefId string `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId"`
      RelatedLearningStandards LearningStandardsDocumentType `xml:"RelatedLearningStandards,omitempty" json:"RelatedLearningStandards"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    