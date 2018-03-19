package sifxml


    type FinancialAccount struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      ParentAccountRefId IdRefType `xml:"ParentAccountRefId,omitempty" json:"ParentAccountRefId"`
      ChargedLocationInfoRefId IdRefType `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      AccountNumber string `xml:"AccountNumber,omitempty" json:"AccountNumber"`
      Name string `xml:"Name,omitempty" json:"Name"`
      Description string `xml:"Description,omitempty" json:"Description"`
      ClassType string `xml:"ClassType,omitempty" json:"ClassType"`
      CreationDate string `xml:"CreationDate,omitempty" json:"CreationDate"`
      CreationTime string `xml:"CreationTime,omitempty" json:"CreationTime"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    