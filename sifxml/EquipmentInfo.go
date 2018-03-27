package sifxml


    type EquipmentInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Name string `xml:"Name,omitempty" json:"Name"`
      Description string `xml:"Description,omitempty" json:"Description"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      AssetNumber LocalIdType `xml:"AssetNumber,omitempty" json:"AssetNumber"`
      InvoiceRefId string `xml:"InvoiceRefId,omitempty" json:"InvoiceRefId"`
      PurchaseOrderRefId string `xml:"PurchaseOrderRefId,omitempty" json:"PurchaseOrderRefId"`
      EquipmentType string `xml:"EquipmentType,omitempty" json:"EquipmentType"`
      SIF_RefId EquipmentInfo_SIF_RefId `xml:"SIF_RefId,omitempty" json:"SIF_RefId"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type EquipmentInfo_SIF_RefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
