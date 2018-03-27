package sifxml


    type NAPStudentResponseSet struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      ReportExclusionFlag string `xml:"ReportExclusionFlag,omitempty" json:"ReportExclusionFlag"`
      CalibrationSampleFlag string `xml:"CalibrationSampleFlag,omitempty" json:"CalibrationSampleFlag"`
      EquatingSampleFlag string `xml:"EquatingSampleFlag,omitempty" json:"EquatingSampleFlag"`
      PathTakenForDomain string `xml:"PathTakenForDomain,omitempty" json:"PathTakenForDomain"`
      ParallelTest string `xml:"ParallelTest,omitempty" json:"ParallelTest"`
      StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      PlatformStudentIdentifier LocalIdType `xml:"PlatformStudentIdentifier,omitempty" json:"PlatformStudentIdentifier"`
      NAPTestRefId string `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId,omitempty" json:"NAPTestLocalId"`
      DomainScore DomainScoreType `xml:"DomainScore,omitempty" json:"DomainScore"`
      TestletList NAPStudentResponseTestletListType `xml:"TestletList,omitempty" json:"TestletList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    