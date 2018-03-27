package sifxml


    type StaffAssignment struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId string `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      SchoolYear SchoolYearType `xml:"SchoolYear,omitempty" json:"SchoolYear"`
      StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      PrimaryAssignment string `xml:"PrimaryAssignment,omitempty" json:"PrimaryAssignment"`
      JobStartDate string `xml:"JobStartDate,omitempty" json:"JobStartDate"`
      JobEndDate string `xml:"JobEndDate,omitempty" json:"JobEndDate"`
      JobFTE string `xml:"JobFTE,omitempty" json:"JobFTE"`
      JobFunction string `xml:"JobFunction,omitempty" json:"JobFunction"`
      EmploymentStatus string `xml:"EmploymentStatus,omitempty" json:"EmploymentStatus"`
      StaffSubjectList StaffSubjectListType `xml:"StaffSubjectList,omitempty" json:"StaffSubjectList"`
      StaffActivity StaffActivityExtensionType `xml:"StaffActivity,omitempty" json:"StaffActivity"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      CasualReliefTeacher string `xml:"CasualReliefTeacher,omitempty" json:"CasualReliefTeacher"`
      Homegroup string `xml:"Homegroup,omitempty" json:"Homegroup"`
      House string `xml:"House,omitempty" json:"House"`
      CalendarSummaryList CalendarSummaryListType `xml:"CalendarSummaryList,omitempty" json:"CalendarSummaryList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    