package sifxml


    type SystemRole struct {
        RefId RefIdType `xml:"RefId,attr" json:"-RefId"`
      SIF_RefId SystemRole_SIF_RefId `xml:"SIF_RefId,omitempty" json:"SIF_RefId"`
      SystemContextList SystemRole_SystemContextList `xml:"SystemContextList,omitempty" json:"SystemContextList"`
      SIF_Metadata string `xml:"SIF_Metadata,omitempty" json:"SIF_Metadata"`
      SIF_ExtendedElements string `xml:"SIF_ExtendedElements,omitempty" json:"SIF_ExtendedElements"`
      
      }
    type SystemRole_SIF_RefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value IdRefType `xml:",chardata" json:"Value"`
}
type SystemRole_SystemContextList struct {
      SystemContext []SystemRole_SystemContext `xml:"SystemContext,omitempty" json:"SystemContext"`
}
    func (this *SystemRole_SystemContextList) UnmarshalJSON(data []byte) error {
        type Alias SystemRole_SystemContextList
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        log.Println(t)
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
type SystemRole_SystemContext struct {
      SystemId string `xml:"SystemId,attr" json:"-SystemId"`
      RoleList SystemRole_RoleList `xml:"RoleList,omitempty" json:"RoleList"`
}
type SystemRole_RoleList struct {
      Role []SystemRole_Role `xml:"Role,omitempty" json:"Role"`
}
    func (this *SystemRole_RoleList) UnmarshalJSON(data []byte) error {
        type Alias SystemRole_RoleList
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        log.Println(t)
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
type SystemRole_Role struct {
      RoleId string `xml:"RoleId,attr" json:"-RoleId"`
      RoleScopeList SystemRole_RoleScopeList `xml:"RoleScopeList,omitempty" json:"RoleScopeList"`
}
type SystemRole_RoleScopeList struct {
      RoleScope []SystemRole_RoleScope `xml:"RoleScope,omitempty" json:"RoleScope"`
}
    func (this *SystemRole_RoleScopeList) UnmarshalJSON(data []byte) error {
        type Alias SystemRole_RoleScopeList
        d := json.NewDecoder(bytes.NewBuffer(data))
        t, err := d.Token()
        t, err = d.Token()
        t, err = d.Token()
        if err != nil {
                return err
        }
        log.Println(t)
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
type SystemRole_RoleScope struct {
       RoleScopeName string `xml:"RoleScopeName,omitempty" json:"RoleScopeName"`
      RoleScopeRefId SystemRole_RoleScopeRefId `xml:"RoleScopeRefId,omitempty" json:"RoleScopeRefId"`
}
type SystemRole_RoleScopeRefId struct {
      SIF_RefObject string `xml:"SIF_RefObject,attr" json:"-SIF_RefObject"`
      Value string `xml:",chardata" json:"Value"`
}
