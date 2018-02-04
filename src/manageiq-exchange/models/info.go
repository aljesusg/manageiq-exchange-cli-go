package models

import(
  "fmt"
)
type Info struct {
	Version   string
  Providers map[string]Provider
}

type Provider struct{
  Type string          `json:"type"`
  Enabled bool         `json:"enabled"`
  ApplicationId string `json:"id_application"`
  Server string        `json:"server"`
  Version string       `json:"version"`
  Verify bool          `json:"verify"`
}

func (a *Info) Init(data map[string]interface{}){
  a.Version = data["version"].(string)
  a.Providers = make(map[string]Provider)
  for k, v := range data["providers"].(map[string]interface{}) {
    var pr Provider
    CreateFromMap(v.(map[string]interface{}), &pr)
    a.Providers[k] = pr
  }
}

func (a *Info) Print() string{
  var result string
  result = fmt.Sprintf("%s: %s\n\n",PrintColor("Version","Red"),a.Version)
  result += fmt.Sprintf("Providers: \n")
  for k,v := range a.Providers {
    result += fmt.Sprintf("    %s: \n",k)
    result += fmt.Sprintf("        %s: %s\n","Server",v.Server)
    result += fmt.Sprintf("        %s: %s\n","ApplicationId",v.ApplicationId)
    result += fmt.Sprintf("        %s: %s\n","Version",v.Version)
  }
  return result
}
