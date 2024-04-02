package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) Schema() string {
	return `{
      "type":"record",
      "name":"User",
      "fields":[
       {
          "name":"ID",
          "type":"int"
       },
       {
          "name":"Name",
          "type":"string"
       },
       {
         "name":"Email",
         "type":"string"
       }
      ]
   }`
}
