package sifxml


    type LearningResource struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Name string `xml:"Name,omitempty" json:"Name"`
      Author string `xml:"Author,omitempty" json:"Author"`
      Contacts ContactsType `xml:"Contacts,omitempty" json:"Contacts"`
      Location LearningResource_Location `xml:"Location,omitempty" json:"Location"`
      Status string `xml:"Status,omitempty" json:"Status"`
      Description string `xml:"Description,omitempty" json:"Description"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      SubjectAreas ACStrandAreaListType `xml:"SubjectAreas,omitempty" json:"SubjectAreas"`
      MediaTypes MediaTypesType `xml:"MediaTypes,omitempty" json:"MediaTypes"`
      UseAgreement string `xml:"UseAgreement,omitempty" json:"UseAgreement"`
      AgreementDate string `xml:"AgreementDate,omitempty" json:"AgreementDate"`
      Approvals ApprovalsType `xml:"Approvals,omitempty" json:"Approvals"`
      Evaluations EvaluationsType `xml:"Evaluations,omitempty" json:"Evaluations"`
      Components ComponentsType `xml:"Components,omitempty" json:"Components"`
      LearningStandards LearningStandardsType `xml:"LearningStandards,omitempty" json:"LearningStandards"`
      LearningResourcePackageRefId IdRefType `xml:"LearningResourcePackageRefId,omitempty" json:"LearningResourcePackageRefId"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type LearningResource_Location struct {
      ReferenceType string `xml:"ReferenceType,attr" json:"-ReferenceType"`
      Value string `xml:",chardata" json:"Value"`
}
