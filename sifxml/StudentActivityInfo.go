package sifxml


    type StudentActivityInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Title string `xml:"Title,omitempty" json:"Title"`
      Description string `xml:"Description,omitempty" json:"Description"`
      StudentActivityType StudentActivityType `xml:"StudentActivityType,omitempty" json:"StudentActivityType"`
      StudentActivityLevel string `xml:"StudentActivityLevel,omitempty" json:"StudentActivityLevel"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      CurricularStatus string `xml:"CurricularStatus,omitempty" json:"CurricularStatus"`
      Location LocationType `xml:"Location,omitempty" json:"Location"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    