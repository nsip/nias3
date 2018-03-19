package sifxml


    type TermInfo struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      StartDate string `xml:"StartDate,omitempty" json:"StartDate"`
      EndDate string `xml:"EndDate,omitempty" json:"EndDate"`
      Description string `xml:"Description,omitempty" json:"Description"`
      RelativeDuration string `xml:"RelativeDuration,omitempty" json:"RelativeDuration"`
      TermCode string `xml:"TermCode,omitempty" json:"TermCode"`
      Track string `xml:"Track,omitempty" json:"Track"`
      TermSpan string `xml:"TermSpan,omitempty" json:"TermSpan"`
      MarkingTerm string `xml:"MarkingTerm,omitempty" json:"MarkingTerm"`
      SchedulingTerm string `xml:"SchedulingTerm,omitempty" json:"SchedulingTerm"`
      AttendanceTerm string `xml:"AttendanceTerm,omitempty" json:"AttendanceTerm"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    