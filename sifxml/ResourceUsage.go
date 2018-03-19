package sifxml
import (
"bytes"
"encoding/json"
"log"
)


    type ResourceUsage struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SchoolInfoRefId IdRefType `xml:"SchoolInfoRefId,omitempty" json:"SchoolInfoRefId"`
      ResourceUsageContentType ResourceUsage_ResourceUsageContentType `xml:"ResourceUsageContentType,omitempty" json:"ResourceUsageContentType"`
      ResourceReportColumnList ResourceUsage_ResourceReportColumnList `xml:"ResourceReportColumnList,omitempty" json:"ResourceReportColumnList"`
      ResourceReportLineList ResourceUsage_ResourceReportLineList `xml:"ResourceReportLineList,omitempty" json:"ResourceReportLineList"`
      SIF_Metadata SIF_MetadataType `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements SIF_ExtendedElementsType `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type ResourceUsage_ResourceUsageContentType struct {
       Code string `xml:"Code,omitempty" json:"Code"`
       LocalDescription string `xml:"LocalDescription,omitempty" json:"LocalDescription"`
}
type ResourceUsage_ResourceReportColumnList struct {
      ResourceReportColumn []ResourceUsage_ResourceReportColumn `xml:"ResourceReportColumn,omitempty" json:"ResourceReportColumn"`
}
    func (this *ResourceUsage_ResourceReportColumnList) UnmarshalJSON(data []byte) error {
        type Alias ResourceUsage_ResourceReportColumnList
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
type ResourceUsage_ResourceReportLineList struct {
      ResourceReportLine []ResourceUsage_ResourceReportLine `xml:"ResourceReportLine,omitempty" json:"ResourceReportLine"`
}
    func (this *ResourceUsage_ResourceReportLineList) UnmarshalJSON(data []byte) error {
        type Alias ResourceUsage_ResourceReportLineList
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
type ResourceUsage_ResourceReportColumn struct {
       ColumnName string `xml:"ColumnName,omitempty" json:"ColumnName"`
       ColumnDescription string `xml:"ColumnDescription,omitempty" json:"ColumnDescription"`
       ColumnDelimiter string `xml:"ColumnDelimiter,omitempty" json:"ColumnDelimiter"`
}
type ResourceUsage_ResourceReportLine struct {
      SIF_RefId ResourceUsage_SIF_RefId `xml:"SIF_RefId,omitempty" json:"SIF_RefId"`
       StartDate string `xml:"StartDate,omitempty" json:"StartDate"`
       EndDate string `xml:"EndDate,omitempty" json:"EndDate"`
       CurrentCost MonetaryAmountType `xml:"CurrentCost,omitempty" json:"CurrentCost"`
       ReportRow string `xml:"ReportRow,omitempty" json:"ReportRow"`
}
type ResourceUsage_SIF_RefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
