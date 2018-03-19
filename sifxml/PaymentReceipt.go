package sifxml


    type PaymentReceipt struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      TransactionType string `xml:"TransactionType,omitempty" json:"TransactionType"`
      InvoiceRefId IdRefType `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      VendorInfoRefId IdRefType `xml:"VendorInfoRefId,omitempty" json:"VendorInfoRefId"`
      DebtorRefId IdRefType `xml:"DebtorRefId,omitempty" json:"DebtorRefId"`
      ChargedLocationInfoRefId IdRefType `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      TransactionDate string `xml:"TransactionDate,omitempty" json:"TransactionDate"`
      TransactionAmount DebitOrCreditAmountType `xml:"TransactionAmount,omitempty" json:"TransactionAmount"`
      ReceivedTransactionId string `xml:"ReceivedTransactionId,omitempty" json:"ReceivedTransactionId"`
      FinancialAccountRefIdList FinancialAccountRefIdListType `xml:"FinancialAccountRefIdList,omitempty" json:"FinancialAccountRefIdList"`
      TransactionDescription string `xml:"TransactionDescription,omitempty" json:"TransactionDescription"`
      TaxRate string `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      TransactionMethod string `xml:"TransactionMethod,omitempty" json:"TransactionMethod"`
      ChequeNumber string `xml:"ChequeNumber,omitempty" json:"ChequeNumber"`
      TransactionNote string `xml:"TransactionNote,omitempty" json:"TransactionNote"`
      AccountingPeriod LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    