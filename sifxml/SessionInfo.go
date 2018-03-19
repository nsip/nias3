package sifxml


    type SessionInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      TimeTableCellRefId IdRefType `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      TimeTableSubjectLocalId LocalIdType `xml:"TimeTableSubjectLocalId,omitempty" json:"TimeTableSubjectLocalId"`
      TeachingGroupLocalId LocalIdType `xml:"TeachingGroupLocalId,omitempty" json:"TeachingGroupLocalId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      StaffPersonalLocalId LocalIdType `xml:"StaffPersonalLocalId,omitempty" json:"StaffPersonalLocalId"`
      RoomNumber string `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      DayId LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      SessionDate string `xml:"SessionDate,omitempty" json:"SessionDate"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      FinishTime string `xml:"FinishTime,omitempty" json:"FinishTime"`
      RollMarked string `xml:"RollMarked,omitempty" json:"RollMarked"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    