package sifxml


    type AggregateStatisticFact struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      AggregateStatisticInfoRefId string `xml:"AggregateStatisticInfoRefId,omitempty" json:"AggregateStatisticInfoRefId"`
      Characteristics CharacteristicsType `xml:"Characteristics,omitempty" json:"Characteristics"`
      Excluded string `xml:"Excluded,omitempty" json:"Excluded"`
      Value string `xml:"Value,omitempty" json:"Value"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    