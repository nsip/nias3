package sifxml
import (
"bytes"
"encoding/json"
"log"
)


    type ReportPackageType AbstractContentPackageType
    type AbstractContentPackageType struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      XMLData AbstractContentPackageType_XMLData `xml:"XMLData,omitempty" json:"XMLData"`
      TextData AbstractContentPackageType_TextData `xml:"TextData,omitempty" json:"TextData"`
      BinaryData AbstractContentPackageType_BinaryData `xml:"BinaryData,omitempty" json:"BinaryData"`
      Reference AbstractContentPackageType_Reference `xml:"Reference,omitempty" json:"Reference"`
      
      }
    
    type AbstractContentElementType struct {
      XMLData AbstractContentElementType_XMLData `xml:"XMLData,omitempty" json:"XMLData"`
      TextData AbstractContentElementType_TextData `xml:"TextData,omitempty" json:"TextData"`
      BinaryData AbstractContentElementType_BinaryData `xml:"BinaryData,omitempty" json:"BinaryData"`
      Reference AbstractContentElementType_Reference `xml:"Reference,omitempty" json:"Reference"`
      
      }
    
    type MonetaryAmountType struct {
          Currency string `xml:"Currency,attr" json:"-Currency"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type ObjectNameType string
    type ServiceNameType string
    type ObjectType string
    type ReportDataObjectType string
    type URIOrBinaryType string
    type GUIDType string
    type MsgIdType GUIDType
    type RefIdType GUIDType
    type IdRefType RefIdType
    type VersionType string
    type VersionWithWildcardsType string
    type DefinedProtocolsType string
    type ExtendedContentType string
    type SelectedContentType string
    type CopyRightContainerType struct {
        Date string `xml:"Date,omitempty" json:"Date"`
      Holder string `xml:"Holder,omitempty" json:"Holder"`
      
      }
    
    type StandardsSettingBodyType struct {
        Country CountryType `xml:"Country,omitempty" json:"Country"`
      StateProvince StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince"`
      SettingBodyName string `xml:"SettingBodyName,omitempty" json:"SettingBodyName"`
      
      }
    
    type StandardHierarchyLevelType struct {
        Number string `xml:"Number,omitempty" json:"Number"`
      Description string `xml:"Description,omitempty" json:"Description"`
      
      }
    
    type StandardIdentifierType struct {
        YearCreated string `xml:"YearCreated,omitempty" json:"YearCreated"`
      ACStrandSubjectArea ACStrandSubjectAreaType `xml:"ACStrandSubjectArea,omitempty" json:"ACStrandSubjectArea"`
      StandardNumber string `xml:"StandardNumber,omitempty" json:"StandardNumber"`
      YearLevels YearLevelsType `xml:"YearLevels,omitempty" json:"YearLevels"`
      Benchmark string `xml:"Benchmark,omitempty" json:"Benchmark"`
      YearLevel YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      IndicatorNumber string `xml:"IndicatorNumber,omitempty" json:"IndicatorNumber"`
      AlternateIdentificationCodes AlternateIdentificationCodeListType `xml:"AlternateIdentificationCodes,omitempty" json:"AlternateIdentificationCodes"`
      Organization string `xml:"Organization,omitempty" json:"Organization"`
      
      }
    
    type AlternateIdentificationCodeListType struct {
        AlternateIdentificationCode []string `xml:"AlternateIdentificationCode,omitempty" json:"AlternateIdentificationCode"`
      
      }
    func (this *AlternateIdentificationCodeListType) UnmarshalJSON(data []byte) error {
        type Alias AlternateIdentificationCodeListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type RelatedLearningStandardItemRefIdListType struct {
        LearningStandardItemRefId []RelatedLearningStandardItemRefIdType `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId"`
      
      }
    func (this *RelatedLearningStandardItemRefIdListType) UnmarshalJSON(data []byte) error {
        type Alias RelatedLearningStandardItemRefIdListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type RelatedLearningStandardItemRefIdType struct {
          RelationshipType string `xml:"RelationshipType,attr" json:"-RelationshipType"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type ValidLetterMarkListType struct {
        ValidLetterMark []ValidLetterMarkType `xml:"ValidLetterMark,omitempty" json:"ValidLetterMark"`
      
      }
    func (this *ValidLetterMarkListType) UnmarshalJSON(data []byte) error {
        type Alias ValidLetterMarkListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ValidLetterMarkType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      NumericEquivalent string `xml:"NumericEquivalent,omitempty" json:"NumericEquivalent"`
      Description string `xml:"Description,omitempty" json:"Description"`
      
      }
    
    type StudentGradeMarkersListType struct {
        Marker []MarkerType `xml:"Marker,omitempty" json:"Marker"`
      
      }
    func (this *StudentGradeMarkersListType) UnmarshalJSON(data []byte) error {
        type Alias StudentGradeMarkersListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type MarkerType struct {
        StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      Role string `xml:"Role,omitempty" json:"Role"`
      
      }
    
    type GradingScoreListType struct {
        GradingAssignmentScore []AssignmentScoreType `xml:"GradingAssignmentScore,omitempty" json:"GradingAssignmentScore"`
      
      }
    func (this *GradingScoreListType) UnmarshalJSON(data []byte) error {
        type Alias GradingScoreListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AssignmentScoreType struct {
        GradingAssignmentScoreRefId string `xml:"GradingAssignmentScoreRefId,omitempty" json:"GradingAssignmentScoreRefId"`
      Weight string `xml:"Weight,omitempty" json:"Weight"`
      
      }
    
    type GradeType struct {
        Percentage string `xml:"Percentage,omitempty" json:"Percentage"`
      Numeric string `xml:"Numeric,omitempty" json:"Numeric"`
      Letter string `xml:"Letter,omitempty" json:"Letter"`
      Narrative string `xml:"Narrative,omitempty" json:"Narrative"`
      MarkInfoRefId string `xml:"MarkInfoRefId,omitempty" json:"MarkInfoRefId"`
      
      }
    
    type LearningStandardListType struct {
        LearningStandard []LearningStandardType `xml:"LearningStandard,omitempty" json:"LearningStandard"`
      
      }
    func (this *LearningStandardListType) UnmarshalJSON(data []byte) error {
        type Alias LearningStandardListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LearningStandardType struct {
        LearningStandardItemRefId []string `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId"`
      LearningStandardURL string `xml:"LearningStandardURL,omitempty" json:"LearningStandardURL"`
      LearningStandardLocalId LocalIdType `xml:"LearningStandardLocalId,omitempty" json:"LearningStandardLocalId"`
      
      }
    func (this *LearningStandardType) UnmarshalJSON(data []byte) error {
        type Alias LearningStandardType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AssignmentListType struct {
        GradingAssignmentRefId []string `xml:"GradingAssignmentRefId,omitempty" json:"GradingAssignmentRefId"`
      
      }
    func (this *AssignmentListType) UnmarshalJSON(data []byte) error {
        type Alias AssignmentListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type GenericRubricType struct {
        RubricType string `xml:"RubricType,omitempty" json:"RubricType"`
      ScoreList ScoreListType `xml:"ScoreList,omitempty" json:"ScoreList"`
      Descriptor string `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type MedicationListType struct {
        Medication []MedicationType `xml:"Medication,omitempty" json:"Medication"`
      
      }
    func (this *MedicationListType) UnmarshalJSON(data []byte) error {
        type Alias MedicationListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type MedicationType struct {
        MedicationName string `xml:"MedicationName,omitempty" json:"MedicationName"`
      Dosage string `xml:"Dosage,omitempty" json:"Dosage"`
      Frequency string `xml:"Frequency,omitempty" json:"Frequency"`
      AdministrationInformation string `xml:"AdministrationInformation,omitempty" json:"AdministrationInformation"`
      Method string `xml:"Method,omitempty" json:"Method"`
      
      }
    
    type WellbeingEventCategoryListType struct {
        WellbeingEventCategory []WellbeingEventCategoryType `xml:"WellbeingEventCategory,omitempty" json:"WellbeingEventCategory"`
      
      }
    func (this *WellbeingEventCategoryListType) UnmarshalJSON(data []byte) error {
        type Alias WellbeingEventCategoryListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type WellbeingEventCategoryType struct {
        EventCategory string `xml:"EventCategory,omitempty" json:"EventCategory"`
      WellbeingEventSubCategoryList WellbeingEventSubCategoryListType `xml:"WellbeingEventSubCategoryList,omitempty" json:"WellbeingEventSubCategoryList"`
      
      }
    
    type WellbeingEventSubCategoryListType struct {
        WellbeingEventSubCategory []string `xml:"WellbeingEventSubCategory,omitempty" json:"WellbeingEventSubCategory"`
      
      }
    func (this *WellbeingEventSubCategoryListType) UnmarshalJSON(data []byte) error {
        type Alias WellbeingEventSubCategoryListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type WellbeingEventLocationDetailsType struct {
        EventLocation string `xml:"EventLocation,omitempty" json:"EventLocation"`
      Class string `xml:"Class,omitempty" json:"Class"`
      FurtherLocationNotes string `xml:"FurtherLocationNotes,omitempty" json:"FurtherLocationNotes"`
      
      }
    
    type FollowUpActionListType struct {
        FollowUpAction []FollowUpActionType `xml:"FollowUpAction,omitempty" json:"FollowUpAction"`
      
      }
    func (this *FollowUpActionListType) UnmarshalJSON(data []byte) error {
        type Alias FollowUpActionListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type FollowUpActionType struct {
        WellbeingResponseRefId string `xml:"WellbeingResponseRefId,omitempty" json:"WellbeingResponseRefId"`
      FollowUpDetails string `xml:"FollowUpDetails,omitempty" json:"FollowUpDetails"`
      FollowUpActionCategory string `xml:"FollowUpActionCategory,omitempty" json:"FollowUpActionCategory"`
      
      }
    
    type PersonInvolvementListType struct {
        PersonInvolvement []PersonInvolvementType `xml:"PersonInvolvement,omitempty" json:"PersonInvolvement"`
      
      }
    func (this *PersonInvolvementListType) UnmarshalJSON(data []byte) error {
        type Alias PersonInvolvementListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PersonInvolvementType struct {
      PersonRefId PersonInvolvementType_PersonRefId `xml:"PersonRefId,omitempty" json:"PersonRefId"`
      ShortName string `xml:"ShortName,omitempty" json:"ShortName"`
      HowInvolved string `xml:"HowInvolved,omitempty" json:"HowInvolved"`
      
      }
    
    type WithdrawalTimeListType struct {
        Withdrawal []WithdrawalType `xml:"Withdrawal,omitempty" json:"Withdrawal"`
      
      }
    func (this *WithdrawalTimeListType) UnmarshalJSON(data []byte) error {
        type Alias WithdrawalTimeListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type WithdrawalType struct {
        WithdrawalDate string `xml:"WithdrawalDate,omitempty" json:"WithdrawalDate"`
      WithdrawalStartTime string `xml:"WithdrawalStartTime,omitempty" json:"WithdrawalStartTime"`
      WithdrawalEndTime string `xml:"WithdrawalEndTime,omitempty" json:"WithdrawalEndTime"`
      TimeTableSubjectRefId string `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      
      }
    
    type SuspensionContainerType struct {
        SuspensionCategory string `xml:"SuspensionCategory,omitempty" json:"SuspensionCategory"`
      WithdrawalTimeList WithdrawalTimeListType `xml:"WithdrawalTimeList,omitempty" json:"WithdrawalTimeList"`
      Duration string `xml:"Duration,omitempty" json:"Duration"`
      AdvisementDate string `xml:"AdvisementDate,omitempty" json:"AdvisementDate"`
      ResolutionMeetingTime string `xml:"ResolutionMeetingTime,omitempty" json:"ResolutionMeetingTime"`
      ResolutionNotes string `xml:"ResolutionNotes,omitempty" json:"ResolutionNotes"`
      EarlyReturnDate string `xml:"EarlyReturnDate,omitempty" json:"EarlyReturnDate"`
      Status string `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type DetentionContainerType struct {
        DetentionCategory string `xml:"DetentionCategory,omitempty" json:"DetentionCategory"`
      DetentionDate string `xml:"DetentionDate,omitempty" json:"DetentionDate"`
      DetentionLocation string `xml:"DetentionLocation,omitempty" json:"DetentionLocation"`
      DetentionNotes string `xml:"DetentionNotes,omitempty" json:"DetentionNotes"`
      Status string `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type PlanRequiredContainerType struct {
        PlanRequiredList PlanRequiredListType `xml:"PlanRequiredList,omitempty" json:"PlanRequiredList"`
      Status string `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type PlanRequiredListType struct {
        Plan []WellbeingPlanType `xml:"Plan,omitempty" json:"Plan"`
      
      }
    func (this *PlanRequiredListType) UnmarshalJSON(data []byte) error {
        type Alias PlanRequiredListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type WellbeingPlanType struct {
        PersonalisedPlanRefId string `xml:"PersonalisedPlanRefId,omitempty" json:"PersonalisedPlanRefId"`
      PlanNotes string `xml:"PlanNotes,omitempty" json:"PlanNotes"`
      
      }
    
    type AwardContainerType struct {
        AwardDate string `xml:"AwardDate,omitempty" json:"AwardDate"`
      AwardType string `xml:"AwardType,omitempty" json:"AwardType"`
      AwardDescription string `xml:"AwardDescription,omitempty" json:"AwardDescription"`
      AwardNotes string `xml:"AwardNotes,omitempty" json:"AwardNotes"`
      Status string `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type OtherWellbeingResponseContainerType struct {
        OtherResponseDate string `xml:"OtherResponseDate,omitempty" json:"OtherResponseDate"`
      OtherResponseType string `xml:"OtherResponseType,omitempty" json:"OtherResponseType"`
      OtherResponseDescription string `xml:"OtherResponseDescription,omitempty" json:"OtherResponseDescription"`
      OtherResponseNotes string `xml:"OtherResponseNotes,omitempty" json:"OtherResponseNotes"`
      Status string `xml:"Status,omitempty" json:"Status"`
      
      }
    
    type WellbeingDocumentListType struct {
        Document []WellbeingDocumentType `xml:"Document,omitempty" json:"Document"`
      
      }
    func (this *WellbeingDocumentListType) UnmarshalJSON(data []byte) error {
        type Alias WellbeingDocumentListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type WellbeingDocumentType struct {
        Location string `xml:"Location,omitempty" json:"Location"`
      Sensitivity string `xml:"Sensitivity,omitempty" json:"Sensitivity"`
      URL string `xml:"URL,omitempty" json:"URL"`
      DocumentType string `xml:"DocumentType,omitempty" json:"DocumentType"`
      DocumentReviewDate string `xml:"DocumentReviewDate,omitempty" json:"DocumentReviewDate"`
      DocumentDescription string `xml:"DocumentDescription,omitempty" json:"DocumentDescription"`
      
      }
    
    type NAPTestItemContentType struct {
        NAPTestItemLocalId LocalIdType `xml:"NAPTestItemLocalId,omitempty" json:"NAPTestItemLocalId"`
      ItemName string `xml:"ItemName,omitempty" json:"ItemName"`
      ItemType string `xml:"ItemType,omitempty" json:"ItemType"`
      Subdomain string `xml:"Subdomain,omitempty" json:"Subdomain"`
      WritingGenre string `xml:"WritingGenre,omitempty" json:"WritingGenre"`
      ItemDescriptor string `xml:"ItemDescriptor,omitempty" json:"ItemDescriptor"`
      ReleasedStatus string `xml:"ReleasedStatus,omitempty" json:"ReleasedStatus"`
      MarkingType string `xml:"MarkingType,omitempty" json:"MarkingType"`
      MultipleChoiceOptionCount string `xml:"MultipleChoiceOptionCount,omitempty" json:"MultipleChoiceOptionCount"`
      CorrectAnswer string `xml:"CorrectAnswer,omitempty" json:"CorrectAnswer"`
      MaximumScore string `xml:"MaximumScore,omitempty" json:"MaximumScore"`
      ItemDifficulty string `xml:"ItemDifficulty,omitempty" json:"ItemDifficulty"`
      ItemDifficultyLogit5 string `xml:"ItemDifficultyLogit5,omitempty" json:"ItemDifficultyLogit5"`
      ItemDifficultyLogit62 string `xml:"ItemDifficultyLogit62,omitempty" json:"ItemDifficultyLogit62"`
      ItemDifficultyLogit5SE string `xml:"ItemDifficultyLogit5SE,omitempty" json:"ItemDifficultyLogit5SE"`
      ItemDifficultyLogit62SE string `xml:"ItemDifficultyLogit62SE,omitempty" json:"ItemDifficultyLogit62SE"`
      ItemProficiencyBand string `xml:"ItemProficiencyBand,omitempty" json:"ItemProficiencyBand"`
      ItemProficiencyLevel string `xml:"ItemProficiencyLevel,omitempty" json:"ItemProficiencyLevel"`
      ExemplarURL string `xml:"ExemplarURL,omitempty" json:"ExemplarURL"`
      ItemSubstitutedForList SubstituteItemListType `xml:"ItemSubstitutedForList,omitempty" json:"ItemSubstitutedForList"`
      ContentDescriptionList ContentDescriptionListType `xml:"ContentDescriptionList,omitempty" json:"ContentDescriptionList"`
      StimulusList StimulusListType `xml:"StimulusList,omitempty" json:"StimulusList"`
      NAPWritingRubricList NAPWritingRubricListType `xml:"NAPWritingRubricList,omitempty" json:"NAPWritingRubricList"`
      
      }
    
    type NAPTestletContentType struct {
        NAPTestletLocalId LocalIdType `xml:"NAPTestletLocalId,omitempty" json:"NAPTestletLocalId"`
      TestletName string `xml:"TestletName,omitempty" json:"TestletName"`
      Node string `xml:"Node,omitempty" json:"Node"`
      LocationInStage string `xml:"LocationInStage,omitempty" json:"LocationInStage"`
      TestletMaximumScore string `xml:"TestletMaximumScore,omitempty" json:"TestletMaximumScore"`
      
      }
    
    type NAPTestContentType struct {
        NAPTestLocalId LocalIdType `xml:"NAPTestLocalId,omitempty" json:"NAPTestLocalId"`
      TestName string `xml:"TestName,omitempty" json:"TestName"`
      TestLevel YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel"`
      TestType string `xml:"TestType,omitempty" json:"TestType"`
      Domain string `xml:"Domain,omitempty" json:"Domain"`
      TestYear SchoolYearType `xml:"TestYear,omitempty" json:"TestYear"`
      StagesCount string `xml:"StagesCount,omitempty" json:"StagesCount"`
      DomainBands DomainBandsContainerType `xml:"DomainBands,omitempty" json:"DomainBands"`
      DomainProficiency DomainProficiencyContainerType `xml:"DomainProficiency,omitempty" json:"DomainProficiency"`
      
      }
    
    type PlausibleScaledValueListType struct {
        PlausibleScaledValue []string `xml:"PlausibleScaledValue,omitempty" json:"PlausibleScaledValue"`
      
      }
    func (this *PlausibleScaledValueListType) UnmarshalJSON(data []byte) error {
        type Alias PlausibleScaledValueListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SubstituteItemListType struct {
        SubstituteItem []SubstituteItemType `xml:"SubstituteItem,omitempty" json:"SubstituteItem"`
      
      }
    func (this *SubstituteItemListType) UnmarshalJSON(data []byte) error {
        type Alias SubstituteItemListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SubstituteItemType struct {
        SubstituteItemRefId string `xml:"SubstituteItemRefId,omitempty" json:"SubstituteItemRefId"`
      SubstituteItemLocalId LocalIdType `xml:"SubstituteItemLocalId,omitempty" json:"SubstituteItemLocalId"`
      PNPCodeList PNPCodeListType `xml:"PNPCodeList,omitempty" json:"PNPCodeList"`
      
      }
    
    type CodeFrameTestItemListType struct {
        TestItem []CodeFrameTestItemType `xml:"TestItem,omitempty" json:"TestItem"`
      
      }
    func (this *CodeFrameTestItemListType) UnmarshalJSON(data []byte) error {
        type Alias CodeFrameTestItemListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type CodeFrameTestItemType struct {
        TestItemRefId string `xml:"TestItemRefId,omitempty" json:"TestItemRefId"`
      SequenceNumber string `xml:"SequenceNumber,omitempty" json:"SequenceNumber"`
      TestItemContent NAPTestItemContentType `xml:"TestItemContent,omitempty" json:"TestItemContent"`
      
      }
    
    type StimulusLocalIdListType struct {
        StimulusLocalId []LocalIdType `xml:"StimulusLocalId,omitempty" json:"StimulusLocalId"`
      
      }
    func (this *StimulusLocalIdListType) UnmarshalJSON(data []byte) error {
        type Alias StimulusLocalIdListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type DomainBandsContainerType struct {
        Band1Lower string `xml:"Band1Lower,omitempty" json:"Band1Lower"`
      Band1Upper string `xml:"Band1Upper,omitempty" json:"Band1Upper"`
      Band2Lower string `xml:"Band2Lower,omitempty" json:"Band2Lower"`
      Band2Upper string `xml:"Band2Upper,omitempty" json:"Band2Upper"`
      Band3Lower string `xml:"Band3Lower,omitempty" json:"Band3Lower"`
      Band3Upper string `xml:"Band3Upper,omitempty" json:"Band3Upper"`
      Band4Lower string `xml:"Band4Lower,omitempty" json:"Band4Lower"`
      Band4Upper string `xml:"Band4Upper,omitempty" json:"Band4Upper"`
      Band5Lower string `xml:"Band5Lower,omitempty" json:"Band5Lower"`
      Band5Upper string `xml:"Band5Upper,omitempty" json:"Band5Upper"`
      Band6Lower string `xml:"Band6Lower,omitempty" json:"Band6Lower"`
      Band6Upper string `xml:"Band6Upper,omitempty" json:"Band6Upper"`
      Band7Lower string `xml:"Band7Lower,omitempty" json:"Band7Lower"`
      Band7Upper string `xml:"Band7Upper,omitempty" json:"Band7Upper"`
      Band8Lower string `xml:"Band8Lower,omitempty" json:"Band8Lower"`
      Band8Upper string `xml:"Band8Upper,omitempty" json:"Band8Upper"`
      Band9Lower string `xml:"Band9Lower,omitempty" json:"Band9Lower"`
      Band9Upper string `xml:"Band9Upper,omitempty" json:"Band9Upper"`
      Band10Lower string `xml:"Band10Lower,omitempty" json:"Band10Lower"`
      Band10Upper string `xml:"Band10Upper,omitempty" json:"Band10Upper"`
      
      }
    
    type DomainProficiencyContainerType struct {
        Level1Lower string `xml:"Level1Lower,omitempty" json:"Level1Lower"`
      Level1Upper string `xml:"Level1Upper,omitempty" json:"Level1Upper"`
      Level2Lower string `xml:"Level2Lower,omitempty" json:"Level2Lower"`
      Level2Upper string `xml:"Level2Upper,omitempty" json:"Level2Upper"`
      Level3Lower string `xml:"Level3Lower,omitempty" json:"Level3Lower"`
      Level3Upper string `xml:"Level3Upper,omitempty" json:"Level3Upper"`
      Level4Lower string `xml:"Level4Lower,omitempty" json:"Level4Lower"`
      Level4Upper string `xml:"Level4Upper,omitempty" json:"Level4Upper"`
      
      }
    
    type NAPTestItemListType struct {
        TestItem []NAPTestItem2Type `xml:"TestItem,omitempty" json:"TestItem"`
      
      }
    func (this *NAPTestItemListType) UnmarshalJSON(data []byte) error {
        type Alias NAPTestItemListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPTestItem2Type struct {
        TestItemRefId string `xml:"TestItemRefId,omitempty" json:"TestItemRefId"`
      TestItemLocalId LocalIdType `xml:"TestItemLocalId,omitempty" json:"TestItemLocalId"`
      SequenceNumber string `xml:"SequenceNumber,omitempty" json:"SequenceNumber"`
      
      }
    
    type NAPCodeFrameTestletListType struct {
        Testlet []NAPTestletCodeFrameType `xml:"Testlet,omitempty" json:"Testlet"`
      
      }
    func (this *NAPCodeFrameTestletListType) UnmarshalJSON(data []byte) error {
        type Alias NAPCodeFrameTestletListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPTestletCodeFrameType struct {
        NAPTestletRefId string `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId"`
      TestletContent NAPTestletContentType `xml:"TestletContent,omitempty" json:"TestletContent"`
      TestItemList CodeFrameTestItemListType `xml:"TestItemList,omitempty" json:"TestItemList"`
      
      }
    
    type NAPStudentResponseTestletListType struct {
        Testlet []NAPTestletResponseType `xml:"Testlet,omitempty" json:"Testlet"`
      
      }
    func (this *NAPStudentResponseTestletListType) UnmarshalJSON(data []byte) error {
        type Alias NAPStudentResponseTestletListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPTestletResponseType struct {
        NAPTestletRefId string `xml:"NAPTestletRefId,omitempty" json:"NAPTestletRefId"`
      NAPTestletLocalId LocalIdType `xml:"NAPTestletLocalId,omitempty" json:"NAPTestletLocalId"`
      TestletSubScore string `xml:"TestletSubScore,omitempty" json:"TestletSubScore"`
      ItemResponseList NAPTestletItemResponseListType `xml:"ItemResponseList,omitempty" json:"ItemResponseList"`
      
      }
    
    type NAPTestletItemResponseListType struct {
        ItemResponse []NAPTestletResponseItemType `xml:"ItemResponse,omitempty" json:"ItemResponse"`
      
      }
    func (this *NAPTestletItemResponseListType) UnmarshalJSON(data []byte) error {
        type Alias NAPTestletItemResponseListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPTestletResponseItemType struct {
        NAPTestItemRefId string `xml:"NAPTestItemRefId,omitempty" json:"NAPTestItemRefId"`
      NAPTestItemLocalId LocalIdType `xml:"NAPTestItemLocalId,omitempty" json:"NAPTestItemLocalId"`
      Response string `xml:"Response,omitempty" json:"Response"`
      ResponseCorrectness string `xml:"ResponseCorrectness,omitempty" json:"ResponseCorrectness"`
      Score string `xml:"Score,omitempty" json:"Score"`
      LapsedTimeItem string `xml:"LapsedTimeItem,omitempty" json:"LapsedTimeItem"`
      SequenceNumber string `xml:"SequenceNumber,omitempty" json:"SequenceNumber"`
      ItemWeight string `xml:"ItemWeight,omitempty" json:"ItemWeight"`
      SubscoreList NAPSubscoreListType `xml:"SubscoreList,omitempty" json:"SubscoreList"`
      
      }
    
    type NAPSubscoreListType struct {
        Subscore []NAPSubscoreType `xml:"Subscore,omitempty" json:"Subscore"`
      
      }
    func (this *NAPSubscoreListType) UnmarshalJSON(data []byte) error {
        type Alias NAPSubscoreListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPSubscoreType struct {
        SubscoreType string `xml:"SubscoreType,omitempty" json:"SubscoreType"`
      SubscoreValue string `xml:"SubscoreValue,omitempty" json:"SubscoreValue"`
      
      }
    
    type DomainScoreType struct {
        RawScore string `xml:"RawScore,omitempty" json:"RawScore"`
      ScaledScoreValue string `xml:"ScaledScoreValue,omitempty" json:"ScaledScoreValue"`
      ScaledScoreLogitValue string `xml:"ScaledScoreLogitValue,omitempty" json:"ScaledScoreLogitValue"`
      ScaledScoreStandardError string `xml:"ScaledScoreStandardError,omitempty" json:"ScaledScoreStandardError"`
      ScaledScoreLogitStandardError string `xml:"ScaledScoreLogitStandardError,omitempty" json:"ScaledScoreLogitStandardError"`
      StudentDomainBand string `xml:"StudentDomainBand,omitempty" json:"StudentDomainBand"`
      StudentProficiency string `xml:"StudentProficiency,omitempty" json:"StudentProficiency"`
      PlausibleScaledValueList []PlausibleScaledValueListType `xml:"PlausibleScaledValueList,omitempty" json:"PlausibleScaledValueList"`
      
      }
    func (this *DomainScoreType) UnmarshalJSON(data []byte) error {
        type Alias DomainScoreType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPWritingRubricListType struct {
        NAPWritingRubric []NAPWritingRubricType `xml:"NAPWritingRubric,omitempty" json:"NAPWritingRubric"`
      
      }
    func (this *NAPWritingRubricListType) UnmarshalJSON(data []byte) error {
        type Alias NAPWritingRubricListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type NAPWritingRubricType struct {
        RubricType string `xml:"RubricType,omitempty" json:"RubricType"`
      ScoreList ScoreListType `xml:"ScoreList,omitempty" json:"ScoreList"`
      Descriptor string `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type ScoreListType struct {
        Score []ScoreType `xml:"Score,omitempty" json:"Score"`
      
      }
    func (this *ScoreListType) UnmarshalJSON(data []byte) error {
        type Alias ScoreListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ScoreType struct {
        MaxScoreValue string `xml:"MaxScoreValue,omitempty" json:"MaxScoreValue"`
      ScoreDescriptionList ScoreDescriptionListType `xml:"ScoreDescriptionList,omitempty" json:"ScoreDescriptionList"`
      
      }
    
    type ScoreDescriptionListType struct {
        ScoreDescription []ScoreDescriptionType `xml:"ScoreDescription,omitempty" json:"ScoreDescription"`
      
      }
    func (this *ScoreDescriptionListType) UnmarshalJSON(data []byte) error {
        type Alias ScoreDescriptionListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ScoreDescriptionType struct {
        ScoreValue string `xml:"ScoreValue,omitempty" json:"ScoreValue"`
      Descriptor string `xml:"Descriptor,omitempty" json:"Descriptor"`
      
      }
    
    type StimulusListType struct {
        Stimulus []StimulusType `xml:"Stimulus,omitempty" json:"Stimulus"`
      
      }
    func (this *StimulusListType) UnmarshalJSON(data []byte) error {
        type Alias StimulusListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StimulusType struct {
        StimulusLocalId LocalIdType `xml:"StimulusLocalId,omitempty" json:"StimulusLocalId"`
      TextGenre string `xml:"TextGenre,omitempty" json:"TextGenre"`
      TextType string `xml:"TextType,omitempty" json:"TextType"`
      WordCount string `xml:"WordCount,omitempty" json:"WordCount"`
      TextDescriptor string `xml:"TextDescriptor,omitempty" json:"TextDescriptor"`
      Content string `xml:"Content,omitempty" json:"Content"`
      
      }
    
    type ContentDescriptionListType struct {
        ContentDescription []string `xml:"ContentDescription,omitempty" json:"ContentDescription"`
      
      }
    func (this *ContentDescriptionListType) UnmarshalJSON(data []byte) error {
        type Alias ContentDescriptionListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PNPCodeListType struct {
        PNPCode []string `xml:"PNPCode,omitempty" json:"PNPCode"`
      
      }
    func (this *PNPCodeListType) UnmarshalJSON(data []byte) error {
        type Alias PNPCodeListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AdjustmentContainerType struct {
        PNPCodeList PNPCodeListType `xml:"PNPCodeList,omitempty" json:"PNPCodeList"`
      BookletType string `xml:"BookletType,omitempty" json:"BookletType"`
      
      }
    
    type TestDisruptionListType struct {
        TestDisruption []TestDisruptionType `xml:"TestDisruption,omitempty" json:"TestDisruption"`
      
      }
    func (this *TestDisruptionListType) UnmarshalJSON(data []byte) error {
        type Alias TestDisruptionListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TestDisruptionType struct {
        Event string `xml:"Event,omitempty" json:"Event"`
      
      }
    
    type CalendarSummaryListType struct {
        CalendarSummaryRefId []string `xml:"CalendarSummaryRefId,omitempty" json:"CalendarSummaryRefId"`
      
      }
    func (this *CalendarSummaryListType) UnmarshalJSON(data []byte) error {
        type Alias CalendarSummaryListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type VisaSubClassType struct {
        Code VisaSubClassCodeType `xml:"Code,omitempty" json:"Code"`
      VisaExpiryDate string `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate"`
      ATEExpiryDate string `xml:"ATEExpiryDate,omitempty" json:"ATEExpiryDate"`
      ATEStartDate string `xml:"ATEStartDate,omitempty" json:"ATEStartDate"`
      VisaStatisticalCode string `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode"`
      
      }
    
    type VisaSubClassListType struct {
        VisaSubClass []VisaSubClassType `xml:"VisaSubClass,omitempty" json:"VisaSubClass"`
      
      }
    func (this *VisaSubClassListType) UnmarshalJSON(data []byte) error {
        type Alias VisaSubClassListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type VisaSubClassCodeType string
    type LanguageBaseType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      LanguageType string `xml:"LanguageType,omitempty" json:"LanguageType"`
      Dialect string `xml:"Dialect,omitempty" json:"Dialect"`
      
      }
    
    type ReligiousEventListType struct {
        ReligiousEvent []ReligiousEventType `xml:"ReligiousEvent,omitempty" json:"ReligiousEvent"`
      
      }
    func (this *ReligiousEventListType) UnmarshalJSON(data []byte) error {
        type Alias ReligiousEventListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ReligiousEventType struct {
        Type string `xml:"Type,omitempty" json:"Type"`
      Date string `xml:"Date,omitempty" json:"Date"`
      
      }
    
    type ReligionType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type DwellingArrangementType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type CountryListType struct {
        CountryOfCitizenship []CountryType `xml:"CountryOfCitizenship,omitempty" json:"CountryOfCitizenship"`
      
      }
    func (this *CountryListType) UnmarshalJSON(data []byte) error {
        type Alias CountryListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type CountryList2Type struct {
        CountryOfResidency []CountryType `xml:"CountryOfResidency,omitempty" json:"CountryOfResidency"`
      
      }
    func (this *CountryList2Type) UnmarshalJSON(data []byte) error {
        type Alias CountryList2Type
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type DebitOrCreditAmountType struct {
        MonetaryAmountType
          Type string `xml:"Type,attr" json:"-Type"`
      
      }
    
    type ScheduledActivityOverrideType struct {
          DateOfOverride string `xml:"DateOfOverride,attr" json:"-DateOfOverride"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type ActivityTimeType struct {
        CreationDate string `xml:"CreationDate,omitempty" json:"CreationDate"`
      Duration ActivityTimeType_Duration `xml:"Duration,omitempty" json:"Duration"`
      StartDate string `xml:"StartDate,omitempty" json:"StartDate"`
      FinishDate string `xml:"FinishDate,omitempty" json:"FinishDate"`
      DueDate string `xml:"DueDate,omitempty" json:"DueDate"`
      
      }
    
    type SchoolCourseInfoOverrideType struct {
        Override string `xml:"Override,attr" json:"-Override"`
      CourseCode string `xml:"CourseCode,omitempty" json:"CourseCode"`
      StateCourseCode string `xml:"StateCourseCode,omitempty" json:"StateCourseCode"`
      DistrictCourseCode string `xml:"DistrictCourseCode,omitempty" json:"DistrictCourseCode"`
      SubjectArea SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      CourseTitle string `xml:"CourseTitle,omitempty" json:"CourseTitle"`
      InstructionalLevel string `xml:"InstructionalLevel,omitempty" json:"InstructionalLevel"`
      CourseCredits string `xml:"CourseCredits,omitempty" json:"CourseCredits"`
      
      }
    
    type LocationOfInstructionType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type LanguageOfInstructionType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type MediumOfInstructionType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentActivityType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ContactFlagsType struct {
        ParentLegalGuardian string `xml:"ParentLegalGuardian,omitempty" json:"ParentLegalGuardian"`
      PickupRights string `xml:"PickupRights,omitempty" json:"PickupRights"`
      LivesWith string `xml:"LivesWith,omitempty" json:"LivesWith"`
      AccessToRecords string `xml:"AccessToRecords,omitempty" json:"AccessToRecords"`
      ReceivesAssessmentReport string `xml:"ReceivesAssessmentReport,omitempty" json:"ReceivesAssessmentReport"`
      EmergencyContact string `xml:"EmergencyContact,omitempty" json:"EmergencyContact"`
      HasCustody string `xml:"HasCustody,omitempty" json:"HasCustody"`
      DisciplinaryContact string `xml:"DisciplinaryContact,omitempty" json:"DisciplinaryContact"`
      AttendanceContact string `xml:"AttendanceContact,omitempty" json:"AttendanceContact"`
      PrimaryCareProvider string `xml:"PrimaryCareProvider,omitempty" json:"PrimaryCareProvider"`
      FeesBilling string `xml:"FeesBilling,omitempty" json:"FeesBilling"`
      FeesAccess string `xml:"FeesAccess,omitempty" json:"FeesAccess"`
      FamilyMail string `xml:"FamilyMail,omitempty" json:"FamilyMail"`
      InterventionOrder string `xml:"InterventionOrder,omitempty" json:"InterventionOrder"`
      
      }
    
    type AgencyType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type YearRangeType struct {
        Start YearLevelType `xml:"Start,omitempty" json:"Start"`
      End YearLevelType `xml:"End,omitempty" json:"End"`
      
      }
    
    type CreationUserType struct {
        Type string `xml:"Type,attr" json:"-Type"`
      UserId string `xml:"UserId,omitempty" json:"UserId"`
      
      }
    
    type AuditInfoType struct {
        CreationUser CreationUserType `xml:"CreationUser,omitempty" json:"CreationUser"`
      CreationDateTime string `xml:"CreationDateTime,omitempty" json:"CreationDateTime"`
      
      }
    
    type AttendanceInfoType struct {
        CountsTowardAttendance string `xml:"CountsTowardAttendance,omitempty" json:"CountsTowardAttendance"`
      AttendanceValue string `xml:"AttendanceValue,omitempty" json:"AttendanceValue"`
      
      }
    
    type CalendarDateInfoType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ProgramAvailabilityType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ReferralSourceType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type PromotionInfoType struct {
        PromotionStatus string `xml:"PromotionStatus,omitempty" json:"PromotionStatus"`
      
      }
    
    type CatchmentStatusContainerType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentExitStatusContainerType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentExitContainerType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentEntryContainerType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type StudentMostRecentContainerType struct {
        SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      HomeroomLocalId LocalIdType `xml:"HomeroomLocalId,omitempty" json:"HomeroomLocalId"`
      YearLevel YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      FTE string `xml:"FTE,omitempty" json:"FTE"`
      Parent1Language string `xml:"Parent1Language,omitempty" json:"Parent1Language"`
      Parent2Language string `xml:"Parent2Language,omitempty" json:"Parent2Language"`
      Parent1EmploymentType string `xml:"Parent1EmploymentType,omitempty" json:"Parent1EmploymentType"`
      Parent2EmploymentType string `xml:"Parent2EmploymentType,omitempty" json:"Parent2EmploymentType"`
      Parent1SchoolEducationLevel string `xml:"Parent1SchoolEducationLevel,omitempty" json:"Parent1SchoolEducationLevel"`
      Parent2SchoolEducationLevel string `xml:"Parent2SchoolEducationLevel,omitempty" json:"Parent2SchoolEducationLevel"`
      Parent1NonSchoolEducation string `xml:"Parent1NonSchoolEducation,omitempty" json:"Parent1NonSchoolEducation"`
      Parent2NonSchoolEducation string `xml:"Parent2NonSchoolEducation,omitempty" json:"Parent2NonSchoolEducation"`
      LocalCampusId LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      TestLevel YearLevelType `xml:"TestLevel,omitempty" json:"TestLevel"`
      Homegroup string `xml:"Homegroup,omitempty" json:"Homegroup"`
      ClassCode string `xml:"ClassCode,omitempty" json:"ClassCode"`
      MembershipType string `xml:"MembershipType,omitempty" json:"MembershipType"`
      FFPOS string `xml:"FFPOS,omitempty" json:"FFPOS"`
      ReportingSchoolId LocalIdType `xml:"ReportingSchoolId,omitempty" json:"ReportingSchoolId"`
      OtherEnrollmentSchoolACARAId LocalIdType `xml:"OtherEnrollmentSchoolACARAId,omitempty" json:"OtherEnrollmentSchoolACARAId"`
      
      }
    
    type StaffMostRecentContainerType struct {
        SchoolLocalId LocalIdType `xml:"SchoolLocalId,omitempty" json:"SchoolLocalId"`
      SchoolACARAId LocalIdType `xml:"SchoolACARAId,omitempty" json:"SchoolACARAId"`
      LocalCampusId LocalIdType `xml:"LocalCampusId,omitempty" json:"LocalCampusId"`
      NAPLANClassList NAPLANClassListType `xml:"NAPLANClassList,omitempty" json:"NAPLANClassList"`
      HomeGroup string `xml:"HomeGroup,omitempty" json:"HomeGroup"`
      
      }
    
    type StaffActivityExtensionType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type TotalEnrollmentsType struct {
        Girls string `xml:"Girls,omitempty" json:"Girls"`
      Boys string `xml:"Boys,omitempty" json:"Boys"`
      TotalStudents string `xml:"TotalStudents,omitempty" json:"TotalStudents"`
      
      }
    
    type CampusContainerType struct {
        ParentSchoolId string `xml:"ParentSchoolId,omitempty" json:"ParentSchoolId"`
      SchoolCampusId string `xml:"SchoolCampusId,omitempty" json:"SchoolCampusId"`
      CampusType string `xml:"CampusType,omitempty" json:"CampusType"`
      AdminStatus string `xml:"AdminStatus,omitempty" json:"AdminStatus"`
      
      }
    
    type HouseholdContactInfoListType struct {
        HouseholdContactInfo []HouseholdContactInfoType `xml:"HouseholdContactInfo,omitempty" json:"HouseholdContactInfo"`
      
      }
    func (this *HouseholdContactInfoListType) UnmarshalJSON(data []byte) error {
        type Alias HouseholdContactInfoListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type HouseholdContactInfoType struct {
        PreferenceNumber string `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber"`
      HouseholdContactId LocalIdType `xml:"HouseholdContactId,omitempty" json:"HouseholdContactId"`
      HouseholdSalutation string `xml:"HouseholdSalutation,omitempty" json:"HouseholdSalutation"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      EmailList EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      
      }
    
    type StatementCodesType struct {
        StatementCode []string `xml:"StatementCode,omitempty" json:"StatementCode"`
      
      }
    func (this *StatementCodesType) UnmarshalJSON(data []byte) error {
        type Alias StatementCodesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StatementsType struct {
        Statement []string `xml:"Statement,omitempty" json:"Statement"`
      
      }
    func (this *StatementsType) UnmarshalJSON(data []byte) error {
        type Alias StatementsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ProgramFundingSourcesType struct {
        ProgramFundingSource []ProgramFundingSourceType `xml:"ProgramFundingSource,omitempty" json:"ProgramFundingSource"`
      
      }
    func (this *ProgramFundingSourcesType) UnmarshalJSON(data []byte) error {
        type Alias ProgramFundingSourcesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ProgramFundingSourceType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type AttendanceTimesType struct {
        AttendanceTime []AttendanceTimeType `xml:"AttendanceTime,omitempty" json:"AttendanceTime"`
      
      }
    func (this *AttendanceTimesType) UnmarshalJSON(data []byte) error {
        type Alias AttendanceTimesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AttendanceTimeType struct {
        AttendanceType string `xml:"AttendanceType,omitempty" json:"AttendanceType"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode,omitempty" json:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus,omitempty" json:"AttendanceStatus"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime string `xml:"EndTime,omitempty" json:"EndTime"`
      DurationValue string `xml:"DurationValue,omitempty" json:"DurationValue"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      AttendanceNote string `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      
      }
    
    type PeriodAttendancesType struct {
        PeriodAttendance []PeriodAttendanceType `xml:"PeriodAttendance,omitempty" json:"PeriodAttendance"`
      
      }
    func (this *PeriodAttendancesType) UnmarshalJSON(data []byte) error {
        type Alias PeriodAttendancesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PeriodAttendanceType struct {
        AttendanceType string `xml:"AttendanceType,omitempty" json:"AttendanceType"`
      AttendanceCode AttendanceCodeType `xml:"AttendanceCode,omitempty" json:"AttendanceCode"`
      AttendanceStatus string `xml:"AttendanceStatus,omitempty" json:"AttendanceStatus"`
      Date string `xml:"Date,omitempty" json:"Date"`
      SessionInfoRefId string `xml:"SessionInfoRefId,omitempty" json:"SessionInfoRefId"`
      ScheduledActivityRefId string `xml:"ScheduledActivityRefId,omitempty" json:"ScheduledActivityRefId"`
      TimetablePeriod string `xml:"TimetablePeriod,omitempty" json:"TimetablePeriod"`
      DayId LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime string `xml:"EndTime,omitempty" json:"EndTime"`
      TimeIn string `xml:"TimeIn,omitempty" json:"TimeIn"`
      TimeOut string `xml:"TimeOut,omitempty" json:"TimeOut"`
      TimeTableCellRefId string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      TeacherList ScheduledTeacherListType `xml:"TeacherList,omitempty" json:"TeacherList"`
      RoomList RoomListType `xml:"RoomList,omitempty" json:"RoomList"`
      AttendanceNote string `xml:"AttendanceNote,omitempty" json:"AttendanceNote"`
      
      }
    
    type StaffSubjectListType struct {
        StaffSubject []StaffSubjectType `xml:"StaffSubject,omitempty" json:"StaffSubject"`
      
      }
    func (this *StaffSubjectListType) UnmarshalJSON(data []byte) error {
        type Alias StaffSubjectListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StaffSubjectType struct {
        PreferenceNumber string `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      TimeTableSubjectRefId RefIdType `xml:"TimeTableSubjectRefId,omitempty" json:"TimeTableSubjectRefId"`
      
      }
    
    type TeachingGroupListType struct {
        TeachingGroupRefId []string `xml:"TeachingGroupRefId,omitempty" json:"TeachingGroupRefId"`
      
      }
    func (this *TeachingGroupListType) UnmarshalJSON(data []byte) error {
        type Alias TeachingGroupListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ScheduledTeacherListType struct {
        TeacherCover []TeacherCoverType `xml:"TeacherCover,omitempty" json:"TeacherCover"`
      
      }
    func (this *ScheduledTeacherListType) UnmarshalJSON(data []byte) error {
        type Alias ScheduledTeacherListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TeacherCoverType struct {
        StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      StaffLocalId LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      FinishTime string `xml:"FinishTime,omitempty" json:"FinishTime"`
      Credit string `xml:"Credit,omitempty" json:"Credit"`
      Supervision string `xml:"Supervision,omitempty" json:"Supervision"`
      Weighting string `xml:"Weighting,omitempty" json:"Weighting"`
      
      }
    
    type RoomListType struct {
        RoomInfoRefId []string `xml:"RoomInfoRefId,omitempty" json:"RoomInfoRefId"`
      
      }
    func (this *RoomListType) UnmarshalJSON(data []byte) error {
        type Alias RoomListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StaffListType struct {
        StaffPersonalRefId []string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      
      }
    func (this *StaffListType) UnmarshalJSON(data []byte) error {
        type Alias StaffListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AlternateIdentificationCodes struct {
        AlternateIdentificationCode []string `xml:"AlternateIdentificationCode,omitempty" json:"AlternateIdentificationCode"`
      
      }
    func (this *AlternateIdentificationCodes) UnmarshalJSON(data []byte) error {
        type Alias AlternateIdentificationCodes
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AuthorsType struct {
        Author []string `xml:"Author,omitempty" json:"Author"`
      
      }
    func (this *AuthorsType) UnmarshalJSON(data []byte) error {
        type Alias AuthorsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type OrganizationsType struct {
        Organization []string `xml:"Organization,omitempty" json:"Organization"`
      
      }
    func (this *OrganizationsType) UnmarshalJSON(data []byte) error {
        type Alias OrganizationsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PurchasingItemsType struct {
        PurchasingItem []PurchasingItemType `xml:"PurchasingItem,omitempty" json:"PurchasingItem"`
      
      }
    func (this *PurchasingItemsType) UnmarshalJSON(data []byte) error {
        type Alias PurchasingItemsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PurchasingItemType struct {
        ItemNumber string `xml:"ItemNumber,omitempty" json:"ItemNumber"`
      ItemDescription string `xml:"ItemDescription,omitempty" json:"ItemDescription"`
      Quantity string `xml:"Quantity,omitempty" json:"Quantity"`
      UnitCost MonetaryAmountType `xml:"UnitCost,omitempty" json:"UnitCost"`
      TotalCost MonetaryAmountType `xml:"TotalCost,omitempty" json:"TotalCost"`
      QuantityDelivered string `xml:"QuantityDelivered,omitempty" json:"QuantityDelivered"`
      CancelledOrder string `xml:"CancelledOrder,omitempty" json:"CancelledOrder"`
      TaxRate string `xml:"TaxRate,omitempty" json:"TaxRate"`
      ExpenseAccounts ExpenseAccountsType `xml:"ExpenseAccounts,omitempty" json:"ExpenseAccounts"`
      
      }
    
    type ExpenseAccountsType struct {
        ExpenseAccount []ExpenseAccountType `xml:"ExpenseAccount,omitempty" json:"ExpenseAccount"`
      
      }
    func (this *ExpenseAccountsType) UnmarshalJSON(data []byte) error {
        type Alias ExpenseAccountsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ExpenseAccountType struct {
        AccountCode string `xml:"AccountCode,omitempty" json:"AccountCode"`
      Amount MonetaryAmountType `xml:"Amount,omitempty" json:"Amount"`
      FinancialAccountRefId string `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
      AccountingPeriod LocalIdType `xml:"AccountingPeriod,omitempty" json:"AccountingPeriod"`
      
      }
    
    type SchoolProgramListType struct {
        Program []SchoolProgramType `xml:"Program,omitempty" json:"Program"`
      
      }
    func (this *SchoolProgramListType) UnmarshalJSON(data []byte) error {
        type Alias SchoolProgramListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SchoolProgramType struct {
        Category string `xml:"Category,omitempty" json:"Category"`
      Type string `xml:"Type,omitempty" json:"Type"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type LearningObjectivesType struct {
        LearningObjective []string `xml:"LearningObjective,omitempty" json:"LearningObjective"`
      
      }
    func (this *LearningObjectivesType) UnmarshalJSON(data []byte) error {
        type Alias LearningObjectivesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type RecognitionListType struct {
        Recognition []string `xml:"Recognition,omitempty" json:"Recognition"`
      
      }
    func (this *RecognitionListType) UnmarshalJSON(data []byte) error {
        type Alias RecognitionListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LResourcesType struct {
        LearningResourceRefId []ResourcesType `xml:"LearningResourceRefId,omitempty" json:"LearningResourceRefId"`
      
      }
    func (this *LResourcesType) UnmarshalJSON(data []byte) error {
        type Alias LResourcesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ResourcesType struct {
          ResourceType string `xml:"ResourceType,attr" json:"-ResourceType"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type SourceObjectsType struct {
      SourceObject []SourceObjectsType_SourceObject `xml:"SourceObject,omitempty" json:"SourceObject"`
      
      }
    func (this *SourceObjectsType) UnmarshalJSON(data []byte) error {
        type Alias SourceObjectsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StudentsType struct {
        StudentPersonalRefId []string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      
      }
    func (this *StudentsType) UnmarshalJSON(data []byte) error {
        type Alias StudentsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PrerequisitesType struct {
        Prerequisite []string `xml:"Prerequisite,omitempty" json:"Prerequisite"`
      
      }
    func (this *PrerequisitesType) UnmarshalJSON(data []byte) error {
        type Alias PrerequisitesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type EssentialMaterialsType struct {
        EssentialMaterial []string `xml:"EssentialMaterial,omitempty" json:"EssentialMaterial"`
      
      }
    func (this *EssentialMaterialsType) UnmarshalJSON(data []byte) error {
        type Alias EssentialMaterialsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TechnicalRequirementsType struct {
        TechnicalRequirement string `xml:"TechnicalRequirement,omitempty" json:"TechnicalRequirement"`
      
      }
    
    type SoftwareRequirementListType struct {
        SoftwareRequirement []SoftwareRequirementType `xml:"SoftwareRequirement,omitempty" json:"SoftwareRequirement"`
      
      }
    func (this *SoftwareRequirementListType) UnmarshalJSON(data []byte) error {
        type Alias SoftwareRequirementListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SoftwareRequirementType struct {
        SoftwareTitle string `xml:"SoftwareTitle,omitempty" json:"SoftwareTitle"`
      Version string `xml:"Version,omitempty" json:"Version"`
      Vendor string `xml:"Vendor,omitempty" json:"Vendor"`
      OS string `xml:"OS,omitempty" json:"OS"`
      
      }
    
    type HouseholdListType struct {
        Household []LocalIdType `xml:"Household,omitempty" json:"Household"`
      
      }
    func (this *HouseholdListType) UnmarshalJSON(data []byte) error {
        type Alias HouseholdListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StudentSubjectChoiceListType struct {
        StudentSubjectChoice []StudentSubjectChoiceType `xml:"StudentSubjectChoice,omitempty" json:"StudentSubjectChoice"`
      
      }
    func (this *StudentSubjectChoiceListType) UnmarshalJSON(data []byte) error {
        type Alias StudentSubjectChoiceListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StudentSubjectChoiceType struct {
        PreferenceNumber string `xml:"PreferenceNumber,omitempty" json:"PreferenceNumber"`
      SubjectLocalId LocalIdType `xml:"SubjectLocalId,omitempty" json:"SubjectLocalId"`
      StudyDescription SubjectAreaType `xml:"StudyDescription,omitempty" json:"StudyDescription"`
      OtherSchoolLocalId LocalIdType `xml:"OtherSchoolLocalId,omitempty" json:"OtherSchoolLocalId"`
      
      }
    
    type IdentityAssertionsType struct {
      IdentityAssertion []IdentityAssertionsType_IdentityAssertion `xml:"IdentityAssertion,omitempty" json:"IdentityAssertion"`
      
      }
    func (this *IdentityAssertionsType) UnmarshalJSON(data []byte) error {
        type Alias IdentityAssertionsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LearningStandardsType struct {
        LearningStandardItemRefId []string `xml:"LearningStandardItemRefId,omitempty" json:"LearningStandardItemRefId"`
      
      }
    func (this *LearningStandardsType) UnmarshalJSON(data []byte) error {
        type Alias LearningStandardsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LearningResourcesType struct {
        LearningResourceRefId []string `xml:"LearningResourceRefId,omitempty" json:"LearningResourceRefId"`
      
      }
    func (this *LearningResourcesType) UnmarshalJSON(data []byte) error {
        type Alias LearningResourcesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LearningStandardsDocumentType struct {
        LearningStandardDocumentRefId []string `xml:"LearningStandardDocumentRefId,omitempty" json:"LearningStandardDocumentRefId"`
      
      }
    func (this *LearningStandardsDocumentType) UnmarshalJSON(data []byte) error {
        type Alias LearningStandardsDocumentType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ComponentsType struct {
        Component []ComponentType `xml:"Component,omitempty" json:"Component"`
      
      }
    func (this *ComponentsType) UnmarshalJSON(data []byte) error {
        type Alias ComponentsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ComponentType struct {
        Name string `xml:"Name,omitempty" json:"Name"`
      Reference string `xml:"Reference,omitempty" json:"Reference"`
      Description string `xml:"Description,omitempty" json:"Description"`
      Strategies StrategiesType `xml:"Strategies,omitempty" json:"Strategies"`
      AssociatedObjects AssociatedObjectsType `xml:"AssociatedObjects,omitempty" json:"AssociatedObjects"`
      
      }
    
    type StrategiesType struct {
        Strategy []string `xml:"Strategy,omitempty" json:"Strategy"`
      
      }
    func (this *StrategiesType) UnmarshalJSON(data []byte) error {
        type Alias StrategiesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AssociatedObjectsType struct {
      AssociatedObject []AssociatedObjectsType_AssociatedObject `xml:"AssociatedObject,omitempty" json:"AssociatedObject"`
      
      }
    func (this *AssociatedObjectsType) UnmarshalJSON(data []byte) error {
        type Alias AssociatedObjectsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type EvaluationsType struct {
        Evaluation []EvaluationType `xml:"Evaluation,omitempty" json:"Evaluation"`
      
      }
    func (this *EvaluationsType) UnmarshalJSON(data []byte) error {
        type Alias EvaluationsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type EvaluationType struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      Description string `xml:"Description,omitempty" json:"Description"`
      Date string `xml:"Date,omitempty" json:"Date"`
      Name NameType `xml:"Name,omitempty" json:"Name"`
      
      }
    
    type ApprovalsType struct {
        Approval []ApprovalType `xml:"Approval,omitempty" json:"Approval"`
      
      }
    func (this *ApprovalsType) UnmarshalJSON(data []byte) error {
        type Alias ApprovalsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ApprovalType struct {
        Organization string `xml:"Organization,omitempty" json:"Organization"`
      Date string `xml:"Date,omitempty" json:"Date"`
      
      }
    
    type MediaTypesType struct {
        MediaType []string `xml:"MediaType,omitempty" json:"MediaType"`
      
      }
    func (this *MediaTypesType) UnmarshalJSON(data []byte) error {
        type Alias MediaTypesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LEAContactListType struct {
        LEAContact []LEAContactType `xml:"LEAContact,omitempty" json:"LEAContact"`
      
      }
    func (this *LEAContactListType) UnmarshalJSON(data []byte) error {
        type Alias LEAContactListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type LEAContactType struct {
        PublishInDirectory PublishInDirectoryType `xml:"PublishInDirectory,omitempty" json:"PublishInDirectory"`
      ContactInfo ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo"`
      
      }
    
    type FinancialAccountRefIdListType struct {
        FinancialAccountRefId []string `xml:"FinancialAccountRefId,omitempty" json:"FinancialAccountRefId"`
      
      }
    func (this *FinancialAccountRefIdListType) UnmarshalJSON(data []byte) error {
        type Alias FinancialAccountRefIdListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PasswordListType struct {
      Password []PasswordListType_Password `xml:"Password,omitempty" json:"Password"`
      
      }
    func (this *PasswordListType) UnmarshalJSON(data []byte) error {
        type Alias PasswordListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ExclusionRulesType struct {
        ExclusionRule []ExclusionRuleType `xml:"ExclusionRule,omitempty" json:"ExclusionRule"`
      
      }
    func (this *ExclusionRulesType) UnmarshalJSON(data []byte) error {
        type Alias ExclusionRulesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ExclusionRuleType struct {
          Type string `xml:"Type,attr" json:"-Type"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type CharacteristicsType struct {
        AggregateCharacteristicInfoRefId []string `xml:"AggregateCharacteristicInfoRefId,omitempty" json:"AggregateCharacteristicInfoRefId"`
      
      }
    func (this *CharacteristicsType) UnmarshalJSON(data []byte) error {
        type Alias CharacteristicsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ContactsType struct {
        Contact []ContactType `xml:"Contact,omitempty" json:"Contact"`
      
      }
    func (this *ContactsType) UnmarshalJSON(data []byte) error {
        type Alias ContactsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ContactType struct {
        Name NameType `xml:"Name,omitempty" json:"Name"`
      Address AddressType `xml:"Address,omitempty" json:"Address"`
      PhoneNumber PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber"`
      Email EmailType `xml:"Email,omitempty" json:"Email"`
      
      }
    
    type TeachingGroupPeriodListType struct {
        TeachingGroupPeriod []TeachingGroupPeriodType `xml:"TeachingGroupPeriod,omitempty" json:"TeachingGroupPeriod"`
      
      }
    func (this *TeachingGroupPeriodListType) UnmarshalJSON(data []byte) error {
        type Alias TeachingGroupPeriodListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TeachingGroupPeriodType struct {
        TimeTableCellRefId string `xml:"TimeTableCellRefId,omitempty" json:"TimeTableCellRefId"`
      RoomNumber HomeroomNumberType `xml:"RoomNumber,omitempty" json:"RoomNumber"`
      StaffLocalId LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      DayId LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      PeriodId LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      CellType string `xml:"CellType,omitempty" json:"CellType"`
      
      }
    
    type TeacherListType struct {
        TeachingGroupTeacher []TeachingGroupTeacherType `xml:"TeachingGroupTeacher,omitempty" json:"TeachingGroupTeacher"`
      
      }
    func (this *TeacherListType) UnmarshalJSON(data []byte) error {
        type Alias TeacherListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TeachingGroupTeacherType struct {
        StaffPersonalRefId string `xml:"StaffPersonalRefId,omitempty" json:"StaffPersonalRefId"`
      StaffLocalId LocalIdType `xml:"StaffLocalId,omitempty" json:"StaffLocalId"`
      Name NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      Association string `xml:"Association,omitempty" json:"Association"`
      
      }
    
    type StudentListType struct {
        TeachingGroupStudent []TeachingGroupStudentType `xml:"TeachingGroupStudent,omitempty" json:"TeachingGroupStudent"`
      
      }
    func (this *StudentListType) UnmarshalJSON(data []byte) error {
        type Alias StudentListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TeachingGroupStudentType struct {
        StudentPersonalRefId string `xml:"StudentPersonalRefId,omitempty" json:"StudentPersonalRefId"`
      StudentLocalId LocalIdType `xml:"StudentLocalId,omitempty" json:"StudentLocalId"`
      Name NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      
      }
    
    type TimeTableDayListType struct {
        TimeTableDay []TimeTableDayType `xml:"TimeTableDay,omitempty" json:"TimeTableDay"`
      
      }
    func (this *TimeTableDayListType) UnmarshalJSON(data []byte) error {
        type Alias TimeTableDayListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TimeTableDayType struct {
        DayId LocalIdType `xml:"DayId,omitempty" json:"DayId"`
      DayTitle string `xml:"DayTitle,omitempty" json:"DayTitle"`
      TimeTablePeriodList TimeTablePeriodListType `xml:"TimeTablePeriodList,omitempty" json:"TimeTablePeriodList"`
      
      }
    
    type TimeTablePeriodListType struct {
        TimeTablePeriod []TimeTablePeriodType `xml:"TimeTablePeriod,omitempty" json:"TimeTablePeriod"`
      
      }
    func (this *TimeTablePeriodListType) UnmarshalJSON(data []byte) error {
        type Alias TimeTablePeriodListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type TimeTablePeriodType struct {
        PeriodId LocalIdType `xml:"PeriodId,omitempty" json:"PeriodId"`
      PeriodTitle string `xml:"PeriodTitle,omitempty" json:"PeriodTitle"`
      BellPeriod string `xml:"BellPeriod,omitempty" json:"BellPeriod"`
      StartTime string `xml:"StartTime,omitempty" json:"StartTime"`
      EndTime string `xml:"EndTime,omitempty" json:"EndTime"`
      RegularSchoolPeriod string `xml:"RegularSchoolPeriod,omitempty" json:"RegularSchoolPeriod"`
      InstructionalMinutes string `xml:"InstructionalMinutes,omitempty" json:"InstructionalMinutes"`
      UseInAttendanceCalculations string `xml:"UseInAttendanceCalculations,omitempty" json:"UseInAttendanceCalculations"`
      
      }
    
    type NAPLANClassListType struct {
        ClassCode []string `xml:"ClassCode,omitempty" json:"ClassCode"`
      
      }
    func (this *NAPLANClassListType) UnmarshalJSON(data []byte) error {
        type Alias NAPLANClassListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SchoolGroupListType struct {
        SchoolGroup []LocalIdType `xml:"SchoolGroup,omitempty" json:"SchoolGroup"`
      
      }
    func (this *SchoolGroupListType) UnmarshalJSON(data []byte) error {
        type Alias SchoolGroupListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type YearLevelEnrollmentListType struct {
        YearLevelEnrollment []YearLevelEnrollmentType `xml:"YearLevelEnrollment,omitempty" json:"YearLevelEnrollment"`
      
      }
    func (this *YearLevelEnrollmentListType) UnmarshalJSON(data []byte) error {
        type Alias YearLevelEnrollmentListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type YearLevelEnrollmentType struct {
        Year string `xml:"Year,omitempty" json:"Year"`
      Enrollment string `xml:"Enrollment,omitempty" json:"Enrollment"`
      
      }
    
    type SchoolFocusListType struct {
        SchoolFocus []string `xml:"SchoolFocus,omitempty" json:"SchoolFocus"`
      
      }
    func (this *SchoolFocusListType) UnmarshalJSON(data []byte) error {
        type Alias SchoolFocusListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AlertMessagesType struct {
        AlertMessage []AlertMessageType `xml:"AlertMessage,omitempty" json:"AlertMessage"`
      
      }
    func (this *AlertMessagesType) UnmarshalJSON(data []byte) error {
        type Alias AlertMessagesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type AlertMessageType struct {
          Type string `xml:"Type,attr" json:"-Type"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type MedicalAlertMessagesType struct {
        MedicalAlertMessage []MedicalAlertMessageType `xml:"MedicalAlertMessage,omitempty" json:"MedicalAlertMessage"`
      
      }
    func (this *MedicalAlertMessagesType) UnmarshalJSON(data []byte) error {
        type Alias MedicalAlertMessagesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type MedicalAlertMessageType struct {
          Severity string `xml:"Severity,attr" json:"-Severity"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type OtherIdListType struct {
        OtherId []OtherIdType `xml:"OtherId,omitempty" json:"OtherId"`
      
      }
    func (this *OtherIdListType) UnmarshalJSON(data []byte) error {
        type Alias OtherIdListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type OtherIdType struct {
          Type string `xml:"Type,attr" json:"-Type"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type BaseNameType struct {
        Title string `xml:"Title,omitempty" json:"Title"`
      FamilyName string `xml:"FamilyName,omitempty" json:"FamilyName"`
      GivenName string `xml:"GivenName,omitempty" json:"GivenName"`
      MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName"`
      FamilyNameFirst string `xml:"FamilyNameFirst,omitempty" json:"FamilyNameFirst"`
      PreferredFamilyName string `xml:"PreferredFamilyName,omitempty" json:"PreferredFamilyName"`
      PreferredFamilyNameFirst string `xml:"PreferredFamilyNameFirst,omitempty" json:"PreferredFamilyNameFirst"`
      PreferredGivenName string `xml:"PreferredGivenName,omitempty" json:"PreferredGivenName"`
      Suffix string `xml:"Suffix,omitempty" json:"Suffix"`
      FullName string `xml:"FullName,omitempty" json:"FullName"`
      
      }
    
    type NameOfRecordType struct {
        BaseNameType
          Type string `xml:"Type,attr" json:"-Type"`
      
      }
    
    type OtherNameType struct {
        BaseNameType
          Type string `xml:"Type,attr" json:"-Type"`
      
      }
    
    type PartialDateType string
    type LocalIdType string
    type LocationType struct {
        Type string `xml:"Type,attr" json:"-Type"`
      LocationName string `xml:"LocationName,omitempty" json:"LocationName"`
      LocationRefId LocationType_LocationRefId `xml:"LocationRefId,omitempty" json:"LocationRefId"`
      
      }
    
    type StateProvinceIdType string
    type AttendanceCodeType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type YearLevelType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      
      }
    
    type PersonInfoType struct {
        Name NameOfRecordType `xml:"Name,omitempty" json:"Name"`
      OtherNames OtherNamesType `xml:"OtherNames,omitempty" json:"OtherNames"`
      Demographics DemographicsType `xml:"Demographics,omitempty" json:"Demographics"`
      AddressList AddressListType `xml:"AddressList,omitempty" json:"AddressList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      EmailList EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      HouseholdContactInfoList HouseholdContactInfoListType `xml:"HouseholdContactInfoList,omitempty" json:"HouseholdContactInfoList"`
      
      }
    
    type YearLevelsType struct {
        YearLevel []YearLevelType `xml:"YearLevel,omitempty" json:"YearLevel"`
      
      }
    func (this *YearLevelsType) UnmarshalJSON(data []byte) error {
        type Alias YearLevelsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SchoolURLType string
    type PrincipalInfoType struct {
        ContactName NameOfRecordType `xml:"ContactName,omitempty" json:"ContactName"`
      ContactTitle string `xml:"ContactTitle,omitempty" json:"ContactTitle"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      EmailList EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      
      }
    
    type SchoolContactType struct {
        PublishInDirectory PublishInDirectoryType `xml:"PublishInDirectory,omitempty" json:"PublishInDirectory"`
      ContactInfo ContactInfoType `xml:"ContactInfo,omitempty" json:"ContactInfo"`
      
      }
    
    type SchoolContactListType struct {
        SchoolContact []SchoolContactType `xml:"SchoolContact,omitempty" json:"SchoolContact"`
      
      }
    func (this *SchoolContactListType) UnmarshalJSON(data []byte) error {
        type Alias SchoolContactListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PublishInDirectoryType string
    type ContactInfoType struct {
        Name NameType `xml:"Name,omitempty" json:"Name"`
      PositionTitle string `xml:"PositionTitle,omitempty" json:"PositionTitle"`
      Role string `xml:"Role,omitempty" json:"Role"`
      Address AddressType `xml:"Address,omitempty" json:"Address"`
      EmailList EmailListType `xml:"EmailList,omitempty" json:"EmailList"`
      PhoneNumberList PhoneNumberListType `xml:"PhoneNumberList,omitempty" json:"PhoneNumberList"`
      
      }
    
    type AddressStreetType struct {
        Line1 string `xml:"Line1,omitempty" json:"Line1"`
      Line2 string `xml:"Line2,omitempty" json:"Line2"`
      Line3 string `xml:"Line3,omitempty" json:"Line3"`
      Complex string `xml:"Complex,omitempty" json:"Complex"`
      StreetNumber string `xml:"StreetNumber,omitempty" json:"StreetNumber"`
      StreetPrefix string `xml:"StreetPrefix,omitempty" json:"StreetPrefix"`
      StreetName string `xml:"StreetName,omitempty" json:"StreetName"`
      StreetType string `xml:"StreetType,omitempty" json:"StreetType"`
      StreetSuffix string `xml:"StreetSuffix,omitempty" json:"StreetSuffix"`
      ApartmentType string `xml:"ApartmentType,omitempty" json:"ApartmentType"`
      ApartmentNumberPrefix string `xml:"ApartmentNumberPrefix,omitempty" json:"ApartmentNumberPrefix"`
      ApartmentNumber string `xml:"ApartmentNumber,omitempty" json:"ApartmentNumber"`
      ApartmentNumberSuffix string `xml:"ApartmentNumberSuffix,omitempty" json:"ApartmentNumberSuffix"`
      
      }
    
    type AddressType struct {
        Type string `xml:"Type,attr" json:"-Type"`
      Role string `xml:"Role,attr" json:"-Role"`
      EffectiveFromDate string `xml:"EffectiveFromDate,omitempty" json:"EffectiveFromDate"`
      EffectiveToDate string `xml:"EffectiveToDate,omitempty" json:"EffectiveToDate"`
      Street AddressStreetType `xml:"Street,omitempty" json:"Street"`
      City string `xml:"City,omitempty" json:"City"`
      StateProvince StateProvinceType `xml:"StateProvince,omitempty" json:"StateProvince"`
      Country CountryType `xml:"Country,omitempty" json:"Country"`
      PostalCode string `xml:"PostalCode,omitempty" json:"PostalCode"`
      GridLocation GridLocationType `xml:"GridLocation,omitempty" json:"GridLocation"`
      MapReference MapReferenceType `xml:"MapReference,omitempty" json:"MapReference"`
      RadioContact string `xml:"RadioContact,omitempty" json:"RadioContact"`
      Community string `xml:"Community,omitempty" json:"Community"`
      LocalId LocalIdType `xml:"LocalId,omitempty" json:"LocalId"`
      AddressGlobalUID GUIDType `xml:"AddressGlobalUID,omitempty" json:"AddressGlobalUID"`
      StatisticalAreas StatisticalAreasType `xml:"StatisticalAreas,omitempty" json:"StatisticalAreas"`
      
      }
    
    type MapReferenceType struct {
        Type string `xml:"Type,attr" json:"-Type"`
      XCoordinate string `xml:"XCoordinate,omitempty" json:"XCoordinate"`
      YCoordinate string `xml:"YCoordinate,omitempty" json:"YCoordinate"`
      
      }
    
    type StatisticalAreasType struct {
        StatisticalArea []StatisticalAreaType `xml:"StatisticalArea,omitempty" json:"StatisticalArea"`
      
      }
    func (this *StatisticalAreasType) UnmarshalJSON(data []byte) error {
        type Alias StatisticalAreasType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type StatisticalAreaType struct {
          SpatialUnitType string `xml:"SpatialUnitType,attr" json:"-SpatialUnitType"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type AddressListType struct {
        Address []AddressType `xml:"Address,omitempty" json:"Address"`
      
      }
    func (this *AddressListType) UnmarshalJSON(data []byte) error {
        type Alias AddressListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type EmailListType struct {
        Email []EmailType `xml:"Email,omitempty" json:"Email"`
      
      }
    func (this *EmailListType) UnmarshalJSON(data []byte) error {
        type Alias EmailListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type EmailType struct {
          Type string `xml:"Type,attr" json:"-Type"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type PhoneNumberListType struct {
        PhoneNumber []PhoneNumberType `xml:"PhoneNumber,omitempty" json:"PhoneNumber"`
      
      }
    func (this *PhoneNumberListType) UnmarshalJSON(data []byte) error {
        type Alias PhoneNumberListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type PhoneNumberType struct {
        Type string `xml:"Type,attr" json:"-Type"`
      Number string `xml:"Number,omitempty" json:"Number"`
      Extension string `xml:"Extension,omitempty" json:"Extension"`
      ListedStatus string `xml:"ListedStatus,omitempty" json:"ListedStatus"`
      Preference string `xml:"Preference,omitempty" json:"Preference"`
      
      }
    
    type CountryType string
    type GridLocationType struct {
        Latitude string `xml:"Latitude,omitempty" json:"Latitude"`
      Longitude string `xml:"Longitude,omitempty" json:"Longitude"`
      
      }
    
    type OperationalStatusType string
    type StateProvinceType string
    type SchoolYearType string
    type ElectronicIdListType struct {
        ElectronicId []ElectronicIdType `xml:"ElectronicId,omitempty" json:"ElectronicId"`
      
      }
    func (this *ElectronicIdListType) UnmarshalJSON(data []byte) error {
        type Alias ElectronicIdListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ElectronicIdType struct {
          Type string `xml:"Type,attr" json:"-Type"`
      
        Value string `xml:",chardata" json:"Value"`
      }
    
    type OtherNamesType struct {
        Name []OtherNameType `xml:"Name,omitempty" json:"Name"`
      
      }
    func (this *OtherNamesType) UnmarshalJSON(data []byte) error {
        type Alias OtherNamesType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type DemographicsType struct {
        IndigenousStatus string `xml:"IndigenousStatus,omitempty" json:"IndigenousStatus"`
      Sex string `xml:"Sex,omitempty" json:"Sex"`
      BirthDate BirthDateType `xml:"BirthDate,omitempty" json:"BirthDate"`
      DateOfDeath string `xml:"DateOfDeath,omitempty" json:"DateOfDeath"`
      BirthDateVerification string `xml:"BirthDateVerification,omitempty" json:"BirthDateVerification"`
      PlaceOfBirth string `xml:"PlaceOfBirth,omitempty" json:"PlaceOfBirth"`
      StateOfBirth StateProvinceType `xml:"StateOfBirth,omitempty" json:"StateOfBirth"`
      CountryOfBirth CountryType `xml:"CountryOfBirth,omitempty" json:"CountryOfBirth"`
      CountriesOfCitizenship CountryListType `xml:"CountriesOfCitizenship,omitempty" json:"CountriesOfCitizenship"`
      CountriesOfResidency CountryList2Type `xml:"CountriesOfResidency,omitempty" json:"CountriesOfResidency"`
      CountryArrivalDate string `xml:"CountryArrivalDate,omitempty" json:"CountryArrivalDate"`
      AustralianCitizenshipStatus string `xml:"AustralianCitizenshipStatus,omitempty" json:"AustralianCitizenshipStatus"`
      EnglishProficiency EnglishProficiencyType `xml:"EnglishProficiency,omitempty" json:"EnglishProficiency"`
      LanguageList LanguageListType `xml:"LanguageList,omitempty" json:"LanguageList"`
      DwellingArrangement DwellingArrangementType `xml:"DwellingArrangement,omitempty" json:"DwellingArrangement"`
      Religion ReligionType `xml:"Religion,omitempty" json:"Religion"`
      ReligiousEventList ReligiousEventListType `xml:"ReligiousEventList,omitempty" json:"ReligiousEventList"`
      ReligiousRegion string `xml:"ReligiousRegion,omitempty" json:"ReligiousRegion"`
      PermanentResident string `xml:"PermanentResident,omitempty" json:"PermanentResident"`
      VisaSubClass VisaSubClassCodeType `xml:"VisaSubClass,omitempty" json:"VisaSubClass"`
      VisaStatisticalCode string `xml:"VisaStatisticalCode,omitempty" json:"VisaStatisticalCode"`
      VisaExpiryDate string `xml:"VisaExpiryDate,omitempty" json:"VisaExpiryDate"`
      VisaSubClassList VisaSubClassListType `xml:"VisaSubClassList,omitempty" json:"VisaSubClassList"`
      LBOTE string `xml:"LBOTE,omitempty" json:"LBOTE"`
      ImmunisationCertificateStatus string `xml:"ImmunisationCertificateStatus,omitempty" json:"ImmunisationCertificateStatus"`
      CulturalBackground string `xml:"CulturalBackground,omitempty" json:"CulturalBackground"`
      MaritalStatus string `xml:"MaritalStatus,omitempty" json:"MaritalStatus"`
      MedicareNumber string `xml:"MedicareNumber,omitempty" json:"MedicareNumber"`
      
      }
    
    type EnglishProficiencyType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type LanguageListType struct {
        Language []LanguageBaseType `xml:"Language,omitempty" json:"Language"`
      
      }
    func (this *LanguageListType) UnmarshalJSON(data []byte) error {
        type Alias LanguageListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type BirthDateType string
    type ProjectedGraduationYearType string
    type OnTimeGraduationYearType string
    type RelationshipType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type EducationalLevelType string
    type GraduationDateType PartialDateType
    type NameType struct {
        BaseNameType
          Type string `xml:"Type,attr" json:"-Type"`
      
      }
    
    type HomeroomNumberType string
    type TimeElementType struct {
        Type string `xml:"Type,omitempty" json:"Type"`
      Code string `xml:"Code,omitempty" json:"Code"`
      Name string `xml:"Name,omitempty" json:"Name"`
      Value string `xml:"Value,omitempty" json:"Value"`
      StartDateTime string `xml:"StartDateTime,omitempty" json:"StartDateTime"`
      EndDateTime string `xml:"EndDateTime,omitempty" json:"EndDateTime"`
      SpanGaps TimeElementType_SpanGaps `xml:"SpanGaps,omitempty" json:"SpanGaps"`
      IsCurrent string `xml:"IsCurrent,omitempty" json:"IsCurrent"`
      
      }
    
    type LifeCycleType struct {
      Created LifeCycleType_Created `xml:"Created,omitempty" json:"Created"`
      ModificationHistory LifeCycleType_ModificationHistory `xml:"ModificationHistory,omitempty" json:"ModificationHistory"`
      TimeElements LifeCycleType_TimeElements `xml:"TimeElements,omitempty" json:"TimeElements"`
      
      }
    
    type OtherCodeListType struct {
      OtherCode []OtherCodeListType_OtherCode `xml:"OtherCode,omitempty" json:"OtherCode"`
      
      }
    func (this *OtherCodeListType) UnmarshalJSON(data []byte) error {
        type Alias OtherCodeListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ProgramStatusType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type SubjectAreaListType struct {
        SubjectArea []SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      
      }
    func (this *SubjectAreaListType) UnmarshalJSON(data []byte) error {
        type Alias SubjectAreaListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type ACStrandAreaListType struct {
        ACStrandSubjectArea []ACStrandSubjectAreaType `xml:"ACStrandSubjectArea,omitempty" json:"ACStrandSubjectArea"`
      
      }
    func (this *ACStrandAreaListType) UnmarshalJSON(data []byte) error {
        type Alias ACStrandAreaListType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SubjectAreaType struct {
        Code string `xml:"Code,omitempty" json:"Code"`
      OtherCodeList OtherCodeListType `xml:"OtherCodeList,omitempty" json:"OtherCodeList"`
      
      }
    
    type ACStrandSubjectAreaType struct {
        ACStrand string `xml:"ACStrand,omitempty" json:"ACStrand"`
      SubjectArea SubjectAreaType `xml:"SubjectArea,omitempty" json:"SubjectArea"`
      
      }
    
    type EducationFilterType struct {
        LearningStandardItems LearningStandardsType `xml:"LearningStandardItems,omitempty" json:"LearningStandardItems"`
      
      }
    
    type SIF_ExtendedElementsType struct {
      SIF_ExtendedElement []SIF_ExtendedElementsType_SIF_ExtendedElement `xml:"SIF_ExtendedElement,omitempty" json:"SIF_ExtendedElement"`
      
      }
    func (this *SIF_ExtendedElementsType) UnmarshalJSON(data []byte) error {
        type Alias SIF_ExtendedElementsType
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
    
    type SIF_MetadataType struct {
      TimeElements SIF_MetadataType_TimeElements `xml:"TimeElements,omitempty" json:"TimeElements"`
      LifeCycle LifeCycleType `xml:"LifeCycle,omitempty" json:"LifeCycle"`
      EducationFilter EducationFilterType `xml:"EducationFilter,omitempty" json:"EducationFilter"`
      
      }
    type AbstractContentPackageType_XMLData struct {
      Description string `xml:"Description,attr" json:"-Description"`
      Value string `xml:",chardata" json:"Value"`
}
type AbstractContentPackageType_TextData struct {
      MIMEType string `xml:"MIMEType,attr" json:"-MIMEType"`
      FileName string `xml:"FileName,attr" json:"-FileName"`
      Description string `xml:"Description,attr" json:"-Description"`
      Value string `xml:",chardata" json:"Value"`
}
type AbstractContentPackageType_BinaryData struct {
      MIMEType string `xml:"MIMEType,attr" json:"-MIMEType"`
      FileName string `xml:"FileName,attr" json:"-FileName"`
      Description string `xml:"Description,attr" json:"-Description"`
      Value string `xml:",chardata" json:"Value"`
}
type AbstractContentPackageType_Reference struct {
      MIMEType string `xml:"MIMEType,attr" json:"-MIMEType"`
      Description string `xml:"Description,attr" json:"-Description"`
       URL string `xml:"URL,omitempty" json:"URL"`
}
type AbstractContentElementType_XMLData struct {
      Description string `xml:"Description,attr" json:"-Description"`
      Value string `xml:",chardata" json:"Value"`
}
type AbstractContentElementType_TextData struct {
      MIMEType string `xml:"MIMEType,attr" json:"-MIMEType"`
      FileName string `xml:"FileName,attr" json:"-FileName"`
      Description string `xml:"Description,attr" json:"-Description"`
      Value string `xml:",chardata" json:"Value"`
}
type AbstractContentElementType_BinaryData struct {
      MIMEType string `xml:"MIMEType,attr" json:"-MIMEType"`
      FileName string `xml:"FileName,attr" json:"-FileName"`
      Description string `xml:"Description,attr" json:"-Description"`
      Value string `xml:",chardata" json:"Value"`
}
type AbstractContentElementType_Reference struct {
      MIMEType string `xml:"MIMEType,attr" json:"-MIMEType"`
      Description string `xml:"Description,attr" json:"-Description"`
       URL string `xml:"URL,omitempty" json:"URL"`
}
type PersonInvolvementType_PersonRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
type ActivityTimeType_Duration struct {
      Units string `xml:"Units,attr" json:"-Units"`
      Value string `xml:",chardata" json:"Value"`
}
type SourceObjectsType_SourceObject struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
type IdentityAssertionsType_IdentityAssertion struct {
      SchemaName string `xml:"SchemaName,attr" json:"-SchemaName"`
      Value string `xml:",chardata" json:"Value"`
}
type AssociatedObjectsType_AssociatedObject struct {
      SIF_RefObject ObjectNameType `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
type PasswordListType_Password struct {
      Algorithm string `xml:"Algorithm,attr" json:"-Algorithm"`
      KeyName string `xml:"KeyName,attr" json:"-KeyName"`
      Value string `xml:",chardata" json:"Value"`
}
type LocationType_LocationRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
type TimeElementType_SpanGaps struct {
      SpanGap []TimeElementType_SpanGap `xml:"SpanGap,omitempty" json:"SpanGap"`
}
    func (this *TimeElementType_SpanGaps) UnmarshalJSON(data []byte) error {
        type Alias TimeElementType_SpanGaps
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
type LifeCycleType_Created struct {
       DateTime string `xml:"DateTime,omitempty" json:"DateTime"`
      Creators LifeCycleType_Creators `xml:"Creators,omitempty" json:"Creators"`
}
type LifeCycleType_ModificationHistory struct {
      Modified []LifeCycleType_Modified `xml:"Modified,omitempty" json:"Modified"`
}
    func (this *LifeCycleType_ModificationHistory) UnmarshalJSON(data []byte) error {
        type Alias LifeCycleType_ModificationHistory
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
type LifeCycleType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement,omitempty" json:"TimeElement"`
}
    func (this *LifeCycleType_TimeElements) UnmarshalJSON(data []byte) error {
        type Alias LifeCycleType_TimeElements
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
type OtherCodeListType_OtherCode struct {
      Codeset string `xml:"Codeset,attr" json:"-Codeset"`
      Value string `xml:",chardata" json:"Value"`
}
type SIF_ExtendedElementsType_SIF_ExtendedElement struct {
      Name string `xml:"Name,attr" json:"-Name"`
      Type string `xml:"Type,attr" json:"-Type"`
      SIF_Action string `xml:"SIF_Action,attr" json:"-SIF_Action"`
      Value ExtendedContentType `xml:",chardata" json:"Value"`
}
type SIF_MetadataType_TimeElements struct {
       TimeElement []TimeElementType `xml:"TimeElement,omitempty" json:"TimeElement"`
}
    func (this *SIF_MetadataType_TimeElements) UnmarshalJSON(data []byte) error {
        type Alias SIF_MetadataType_TimeElements
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
type TimeElementType_SpanGap struct {
       Type string `xml:"Type,omitempty" json:"Type"`
       Code string `xml:"Code,omitempty" json:"Code"`
       Name string `xml:"Name,omitempty" json:"Name"`
       Value string `xml:"Value,omitempty" json:"Value"`
       StartDateTime string `xml:"StartDateTime,omitempty" json:"StartDateTime"`
       EndDateTime string `xml:"EndDateTime,omitempty" json:"EndDateTime"`
}
type LifeCycleType_Creators struct {
      Creator []LifeCycleType_Creator `xml:"Creator,omitempty" json:"Creator"`
}
    func (this *LifeCycleType_Creators) UnmarshalJSON(data []byte) error {
        type Alias LifeCycleType_Creators
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        if t == json.Delim('[') {
                aux := &struct{ *Alias }{Alias: (*Alias)(this)}
                if err := json.Unmarshal(data, &aux); err != nil {
                        log.Println(err)
                        return err
                }
                return nil
        }
        return this.UnmarshalJSON([]byte(re2.ReplaceAllString(re.ReplaceAllString(string(data), "$1["), "]}")))
}
type LifeCycleType_Modified struct {
       By string `xml:"By,omitempty" json:"By"`
       DateTime string `xml:"DateTime,omitempty" json:"DateTime"`
       Description string `xml:"Description,omitempty" json:"Description"`
}
type LifeCycleType_Creator struct {
       Name string `xml:"Name,omitempty" json:"Name"`
       ID string `xml:"ID,omitempty" json:"ID"`
}
