package Handler

import (
	"encoding/json"
	"mengdiao/ICP-Finder/Service"
	"net/http"
)

func GetICPInfoHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	feature := r.FormValue("feature")

	if url == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}
	if feature == "disabled" {
		info, err := Service.GetSourceICPInfo(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//设置Header
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		//将info编码为JSON格式兵将其发送到http响应中
		json.NewEncoder(w).Encode(info)
	} else if feature == "enabled" || feature == "" {
		info, err := Service.GetICPInfo(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		json.NewEncoder(w).Encode(info)
	} else {
		http.Error(w, "Invalid value for parameter", http.StatusBadRequest)
		return
	}
}
