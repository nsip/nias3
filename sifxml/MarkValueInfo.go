package sifxml


    type MarkValueInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Name string `xml:"Name,omitempty" json:"Name"`
      PercentageMinimum string `xml:"PercentageMinimum,omitempty" json:"PercentageMinimum"`
      PercentageMaximum string `xml:"PercentageMaximum,omitempty" json:"PercentageMaximum"`
      PercentagePassingGrade string `xml:"PercentagePassingGrade,omitempty" json:"PercentagePassingGrade"`
      NumericPrecision string `xml:"NumericPrecision,omitempty" json:"NumericPrecision"`
      NumericScale string `xml:"NumericScale,omitempty" json:"NumericScale"`
      NumericLow string `xml:"NumericLow,omitempty" json:"NumericLow"`
      NumericHigh string `xml:"NumericHigh,omitempty" json:"NumericHigh"`
      NumericPassingGrade string `xml:"NumericPassingGrade,omitempty" json:"NumericPassingGrade"`
      ValidLetterMarkList ValidLetterMarkListType `xml:"ValidLetterMarkList,omitempty" json:"ValidLetterMarkList"`
      Narrative string `xml:"Narrative,omitempty" json:"Narrative"`
      NarrativeMaximumSize string `xml:"NarrativeMaximumSize,omitempty" json:"NarrativeMaximumSize"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    