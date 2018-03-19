package sifxml


    type NAPTestScoreSummary struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      NAPTestRefId IdRefType `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      NAPTestLocalId LocalIdType `xml:"NAPTestLocalId,omitempty" json:"NAPTestLocalId"`
      DomainNationalAverage string `xml:"DomainNationalAverage,omitempty" json:"DomainNationalAverage"`
      DomainSchoolAverage string `xml:"DomainSchoolAverage,omitempty" json:"DomainSchoolAverage"`
      DomainJurisdictionAverage string `xml:"DomainJurisdictionAverage,omitempty" json:"DomainJurisdictionAverage"`
      DomainTopNational60Percent string `xml:"DomainTopNational60Percent,omitempty" json:"DomainTopNational60Percent"`
      DomainBottomNational60Percent string `xml:"DomainBottomNational60Percent,omitempty" json:"DomainBottomNational60Percent"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    