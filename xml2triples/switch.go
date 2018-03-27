package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/nsip/nias3/sifxml"
)

func Root2SIF(r string, j []byte) ([]byte, error) {
	switch r {
	case "Activity":
		res := sifxml.Activity{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "AggregateCharacteristicInfo":
		res := sifxml.AggregateCharacteristicInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "AggregateStatisticFact":
		res := sifxml.AggregateStatisticFact{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "AggregateStatisticInfo":
		res := sifxml.AggregateStatisticInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "CalendarDate":
		res := sifxml.CalendarDate{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "CalendarSummary":
		res := sifxml.CalendarSummary{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "ChargedLocationInfo":
		res := sifxml.ChargedLocationInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "Debtor":
		res := sifxml.Debtor{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "EquipmentInfo":
		res := sifxml.EquipmentInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "FinancialAccount":
		res := sifxml.FinancialAccount{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "GradingAssignment":
		res := sifxml.GradingAssignment{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "GradingAssignmentScore":
		res := sifxml.GradingAssignmentScore{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "Identity":
		res := sifxml.Identity{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "Invoice":
		res := sifxml.Invoice{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "Journal":
		res := sifxml.Journal{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "LEAInfo":
		res := sifxml.LEAInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "LearningResource":
		res := sifxml.LearningResource{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "LearningResourcePackage":
		res := sifxml.LearningResourcePackage{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "LearningStandardDocument":
		res := sifxml.LearningStandardDocument{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "LearningStandardItem":
		res := sifxml.LearningStandardItem{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "MarkValueInfo":
		res := sifxml.MarkValueInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPCodeFrame":
		res := sifxml.NAPCodeFrame{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPEventStudentLink":
		res := sifxml.NAPEventStudentLink{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPStudentResponseSet":
		res := sifxml.NAPStudentResponseSet{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPTest":
		res := sifxml.NAPTest{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPTestItem":
		res := sifxml.NAPTestItem{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPTestScoreSummary":
		res := sifxml.NAPTestScoreSummary{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "NAPTestlet":
		res := sifxml.NAPTestlet{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "PaymentReceipt":
		res := sifxml.PaymentReceipt{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "PersonPicture":
		res := sifxml.PersonPicture{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "PersonalisedPlan":
		res := sifxml.PersonalisedPlan{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "PurchaseOrder":
		res := sifxml.PurchaseOrder{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "ReportAuthorityInfo":
		res := sifxml.ReportAuthorityInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "ResourceBooking":
		res := sifxml.ResourceBooking{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "ResourceUsage":
		res := sifxml.ResourceUsage{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "RoomInfo":
		res := sifxml.RoomInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "ScheduledActivity":
		res := sifxml.ScheduledActivity{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "SchoolCourseInfo":
		res := sifxml.SchoolCourseInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "SchoolInfo":
		res := sifxml.SchoolInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "SchoolPrograms":
		res := sifxml.SchoolPrograms{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "SectionInfo":
		res := sifxml.SectionInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "SessionInfo":
		res := sifxml.SessionInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StaffAssignment":
		res := sifxml.StaffAssignment{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StaffPersonal":
		res := sifxml.StaffPersonal{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentActivityInfo":
		res := sifxml.StudentActivityInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentActivityParticipation":
		res := sifxml.StudentActivityParticipation{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentAttendanceSummary":
		res := sifxml.StudentAttendanceSummary{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentAttendanceTimeList":
		res := sifxml.StudentAttendanceTimeList{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentContactPersonal":
		res := sifxml.StudentContactPersonal{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentContactRelationship":
		res := sifxml.StudentContactRelationship{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentDailyAttendance":
		res := sifxml.StudentDailyAttendance{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentGrade":
		res := sifxml.StudentGrade{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentParticipation":
		res := sifxml.StudentParticipation{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentPeriodAttendance":
		res := sifxml.StudentPeriodAttendance{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentPersonal":
		res := sifxml.StudentPersonal{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentSchoolEnrollment":
		res := sifxml.StudentSchoolEnrollment{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "StudentSectionEnrollment":
		res := sifxml.StudentSectionEnrollment{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "SystemRole":
		res := sifxml.SystemRole{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "TeachingGroup":
		res := sifxml.TeachingGroup{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "TermInfo":
		res := sifxml.TermInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "TimeTable":
		res := sifxml.TimeTable{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "TimeTableCell":
		res := sifxml.TimeTableCell{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "TimeTableSubject":
		res := sifxml.TimeTableSubject{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "VendorInfo":
		res := sifxml.VendorInfo{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "WellbeingAlert":
		res := sifxml.WellbeingAlert{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "WellbeingAppeal":
		res := sifxml.WellbeingAppeal{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "WellbeingCharacteristic":
		res := sifxml.WellbeingCharacteristic{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "WellbeingEvent":
		res := sifxml.WellbeingEvent{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	case "WellbeingResponse":
		res := sifxml.WellbeingResponse{}
		json.Unmarshal(j, &res)
		return xml.MarshalIndent(res, "", "  ")
	default:
		return nil, fmt.Errorf("Type %s is not recognised as a SIF type", r)
	}
}
