package sifxml


    type PurchaseOrder struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      FormNumber string `xml:"FormNumber,omitempty" json:"FormNumber"`
      VendorInfoRefId IdRefType `xml:"VendorInfoRefId,omitempty" json:"VendorInfoRefId"`
      ChargedLocationInfoRefId IdRefType `xml:"ChargedLocationInfoRefId,omitempty" json:"ChargedLocationInfoRefId"`
      EmployeePersonalRefId IdRefType `xml:"EmployeePersonalRefId,omitempty" json:"EmployeePersonalRefId"`
      PurchasingItems PurchasingItemsType `xml:"PurchasingItems,omitempty" json:"PurchasingItems"`
      CreationDate string `xml:"CreationDate,omitempty" json:"CreationDate"`
      TaxRate string `xml:"TaxRate,omitempty" json:"TaxRate"`
      TaxAmount MonetaryAmountType `xml:"TaxAmount,omitempty" json:"TaxAmount"`
      TotalAmount MonetaryAmountType `xml:"TotalAmount,omitempty" json:"TotalAmount"`
      UpdateDate string `xml:"UpdateDate,omitempty" json:"UpdateDate"`
      FullyDelivered string `xml:"FullyDelivered,omitempty" json:"FullyDelivered"`
      OriginalPurchaseOrderRefId IdRefType `xml:"OriginalPurchaseOrderRefId,omitempty" json:"OriginalPurchaseOrderRefId"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    