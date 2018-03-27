package sifxml


    type Invoice struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      InvoicedEntity Invoice_InvoicedEntity `xml:"InvoicedEntity,omitempty" json:"InvoicedEntity"`
      FormNumber LocalIdType `xml:"FormNumber,omitempty" json:"FormNumber"`
      BillingDate string `xml:"BillingDate,omitempty" json:"BillingDate"`
      TransactionDescription string `xml:"TransactionDescription,omitempty" json:"TransactionDescription"`
      BilledAmount DebitOrCreditAmountType `xml:"BilledAmount,omitempty" json:"BilledAmount"`
      Ledger string `xml:"Ledger,omitempty" json:"Ledger"`
      ChargedLocationInfoRefId string `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      NetAmount MonetaryAmountType `xml:"NetAmount,omitempty" json:"NetAmount"`
      TaxRate string `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxType string `xml:"TaxType,omitempty" json:"TaxType"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      CreatedBy string `xml:"CreatedBy,omitempty" json:"CreatedBy"`
      ApprovedBy string `xml:"ApprovedBy,omitempty" json:"ApprovedBy"`
      ItemDetail string `xml:"ItemDetail,omitempty" json:"ItemDetail"`
      DueDate string `xml:"DueDate,omitempty" json:"DueDate"`
      FinancialAccountRefIdList FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList"`
      AccountingPeriod LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      RelatedPurchaseOrderRefId string `xml:"RelatedPurchaseOrderRefId,omitempty" json:"RelatedPurchaseOrderRefId"`
      PurchasingItems PurchasingItemsType `xml:"PurchasingItems,omitempty" json:"PurchasingItems"`
      Voluntary string `xml:"Voluntary,omitempty" json:"Voluntary"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type Invoice_InvoicedEntity struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
