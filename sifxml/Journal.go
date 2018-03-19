package sifxml


    type Journal struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      DebitFinancialAccountRefId IdRefType `xml:"DebitFinancialAccountRefId,omitempty" json:"DebitFinancialAccountRefId"`
      CreditFinancialAccountRefId IdRefType `xml:"CreditFinancialAccountRefId,omitempty" json:"CreditFinancialAccountRefId"`
      OriginatingTransactionRefId Journal_OriginatingTransactionRefId `xml:"OriginatingTransactionRefId,omitempty" json:"OriginatingTransactionRefId"`
      Amount MonetaryAmountType `xml:"Amount,omitempty" json:"Amount"`
      GSTCodeOriginal string `xml:"GSTCodeOriginal,omitempty" json:"GSTCodeOriginal"`
      GSTCodeReplacement string `xml:"GSTCodeReplacement,omitempty" json:"GSTCodeReplacement"`
      Note string `xml:"Note,omitempty" json:"Note"`
      CreatedDate string `xml:"CreatedDate,omitempty" json:"CreatedDate"`
      ApprovedDate string `xml:"ApprovedDate,omitempty" json:"ApprovedDate"`
      CreatedBy string `xml:"CreatedBy,omitempty" json:"CreatedBy"`
      ApprovedBy string `xml:"ApprovedBy,omitempty" json:"ApprovedBy"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Journal_OriginatingTransactionRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
