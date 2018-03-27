package sifxml


    type CalendarSummary struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      DaysInSession string `xml:"DaysInSession,omitempty" json:"DaysInSession"`
      StartDate string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate string `xml:"EndDate,omitempty" json:"EndDate"`
      FirstInstructionDate string `xml:"FirstInstructionDate,omitempty" json:"FirstInstructionDate"`
      LastInstructionDate string `xml:"LastInstructionDate,omitempty" json:"LastInstructionDate"`
      GraduationDate GraduationDateType `xml:"GraduationDate,omitempty" json:"GraduationDate"`
      InstructionalMinutes string `xml:"InstructionalMinutes,omitempty" json:"InstructionalMinutes"`
      MinutesPerDay string `xml:"MinutesPerDay,omitempty" json:"MinutesPerDay"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    