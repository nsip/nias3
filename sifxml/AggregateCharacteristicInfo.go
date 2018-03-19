package sifxml


    type AggregateCharacteristicInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      Definition string `xml:"Definition,omitempty" json:"Definition"`
      ElementName string `xml:"ElementName,omitempty" json:"ElementName"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    