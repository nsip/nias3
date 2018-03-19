package sifxml


    type NAPTest struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      TestContent NAPTestContentType `xml:"TestContent,omitempty" json:"TestContent"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    