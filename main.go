package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PredictRequest struct {
	Uid    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	RDW    float64 `json:"rdw"`
	WBC    float64 `json:"wbc"`
	RBC    float64 `json:"rbc"`
	HGB    float64 `json:"hgb"`
	HCT    float64 `json:"hct"`
	MCV    float64 `json:"mcv"`
	MCH    float64 `json:"mch"`
	MCHC   float64 `json:"mchc"`
	PLT    float64 `json:"plt"`
	NEU    float64 `json:"neu"`
	EOS    float64 `json:"eos"`
	BAS    float64 `json:"bas"`
	LYM    float64 `json:"lym"`
	MON    float64 `json:"mon"`
	SOE    float64 `json:"soe"`
	CHOL   float64 `json:"chol"`
	GLU    float64 `json:"glu"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	request := PredictRequest{
		Uid:    "web-client",
		Age:    parseInt(query.Get("age")),
		Gender: parseInt(query.Get("gender")),
		RDW:    parseFloat(query.Get("rdw")),
		WBC:    parseFloat(query.Get("wbc")),
		RBC:    parseFloat(query.Get("rbc")),
		HGB:    parseFloat(query.Get("hgb")),
		HCT:    parseFloat(query.Get("hct")),
		MCV:    parseFloat(query.Get("mcv")),
		MCH:    parseFloat(query.Get("mch")),
		MCHC:   parseFloat(query.Get("mchc")),
		PLT:    parseFloat(query.Get("plt")),
		NEU:    parseFloat(query.Get("neu")),
		EOS:    parseFloat(query.Get("eos")),
		BAS:    parseFloat(query.Get("bas")),
		LYM:    parseFloat(query.Get("lym")),
		MON:    parseFloat(query.Get("mon")),
		SOE:    parseFloat(query.Get("soe")),
		CHOL:   parseFloat(query.Get("chol")),
		GLU:    parseFloat(query.Get("glu")),
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", "https://apiml.labhub.online/api/v1/predict/hba1c", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Ошибка создания POST-запроса", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer 0l62<EJi/zJx]a?")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка при отправке POST-запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("Внешний API вернул ошибку: %d\n%s", resp.StatusCode, string(errorBody)), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения ответа от внешнего API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func parseInt(value string) int {
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}

func parseFloat(value string) float64 {
	var result float64
	if value == "" {
		return 0
	}
	fmt.Sscanf(value, "%f", &result)
	return result
}

func main() {
	http.HandleFunc("/send", handler)
	fmt.Println("Сервер запущен на http://localhost:8080/send")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
