package sifxml


    type TimeTableCell struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      TimeTableRefId string `xml:"TimeTableRefId,omitempty" json:"TimeTableRefId"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeachingGroupRefId string `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      RoomInfoRefId string `xml:"RoomInfoRefId,omitempty" json:"RoomInfoRefId"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      TimeTableLocalId LocalIdType `xml:"TimeTableLocalId,omitempty" json:"TimeTableLocalId"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      TeachingGroupLocalId LocalIdType `xml:"TeachingGroupLocalId,omitempty" json:"TeachingGroupLocalId"`
      RoomNumber HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      CellType string `xml:"CellType,omitempty" json:"CellType"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    