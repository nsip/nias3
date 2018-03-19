package sifxml


    type PersonPicture struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      ParentObjectRefId PersonPicture_ParentObjectRefId `xml:"ParentObjectRefId,omitempty" json:"ParentObjectRefId"`
      SchoolYear string `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      PictureSource PersonPicture_PictureSource `xml:"PictureSource,omitempty" json:"PictureSource"`
      OKToPublish string `xml:"OKToPublish,omitempty" json:"OKToPublish"`
      SIF_Metadata string `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements string `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type PersonPicture_ParentObjectRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
type PersonPicture_PictureSource struct {
      Type string `xml:"Type,attr" json:"-Type"`
      Value URIOrBinaryType `xml:",chardata" json:"Value"`
}
