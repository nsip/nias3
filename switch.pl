print << "END";
package main
import (
          "github.com/nsip/nias3/sifxml"
          "encoding/xml"
          "encoding/json"
          "fmt"
)

func Root2SIF(r string, j []byte) ([]byte, error) {
  switch r {
END


foreach $l (split /\n/, `ls -1 sifxml`) {
    $l =~ s/\.go//;
    next if $l eq "regexp";
    next if $l eq "DataModel";
    print << "END";
    case "$l":
    res := sifxml.$l\{\}
json.Unmarshal(j, &res)
return xml.MarshalIndent(res, "", "  ")
END
}
print << "END";
default:
  return nil, fmt.Errorf("Type %s is not recognised as a SIF type", r) ;
  }
}
END

