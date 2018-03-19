package sifxml


    type NAPCodeFrame struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      NAPTestRefId IdRefType `xml:"NAPTestRefId,omitempty" json:"NAPTestRefId"`
      TestContent NAPTestContentType `xml:"TestContent,omitempty" json:"TestContent"`
      TestletList NAPCodeFrameTestletListType `xml:"TestletList,omitempty" json:"TestletList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    