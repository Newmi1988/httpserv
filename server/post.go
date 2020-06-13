package server

import(
	//"fmt"
	"net/http"
	"encoding/json"
)

// define the structure of the object in the payload
type Structure struct{
	Id string
}

// api response as stream
func PostStruct (w http.ResponseWriter, req *http.Request) {

	var u Structure
	if req.Body == nil {
		http.Error(w, "Request body empty", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	
	// modify user or implement logic for response
	u.Id += " Suffix"

	json.NewEncoder(w).Encode(&u)
}
