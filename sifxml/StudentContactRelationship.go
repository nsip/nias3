package sifxml


    type StudentContactRelationship struct {
        StudentContactRelationshipRefId string `xml:"StudentContactRelationshipRefId,attr" json:"-StudentContactRelationshipRefId"`
      StudentPersonalRefId RefIdType `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentContactPersonalRefId RefIdType `xml:"StudentContactPersonalRefId,omitempty" json:"StudentContactPersonalRefId"`
      Relationship RelationshipType `xml:"Relationship,omitempty" json:"Relationship"`
      ParentRelationshipStatus string `xml:"ParentRelationshipStatus,omitempty" json:"ParentRelationshipStatus"`
      HouseholdList HouseholdListType `xml:"HouseholdList,omitempty" json:"HouseholdList"`
      ContactFlags ContactFlagsType `xml:"ContactFlags,omitempty" json:"ContactFlags"`
      MainlySpeaksEnglishAtHome string `xml:"MainlySpeaksEnglishAtHome,omitempty" json:"MainlySpeaksEnglishAtHome"`
      ContactSequence string `xml:"ContactSequence,omitempty" json:"ContactSequence"`
      ContactSequenceSource string `xml:"ContactSequenceSource,omitempty" json:"ContactSequenceSource"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    