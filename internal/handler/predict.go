package handler

import (
	"bytes"
	"encoding/json"
	"go-api-practice/config"
	"go-api-practice/internal/interfaces"
	"go-api-practice/internal/services"
	"go-api-practice/pkg/logger"
	"io"
	"net/http"
)

var supportedModels = map[string]bool{
	"hba1c": true,
	"ldll":  true,
	"ferr":  true,
	"ldl":   true,
	"tg":    true,
}

var modelBuilders = map[string]interfaces.PredictionBuilder{
	"hba1c": &services.Hba1cBuilder{},
	"ldll":  &services.LdllBuilder{},
	"ferr":  &services.FerrBuilder{},
	"ldl":   &services.LdlBuilder{},
	"tg":    &services.TgBuilder{},
	"hdl":   &services.HdlBuilder{}, // для проверки отключения
}

func PredictHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	model := r.URL.Query().Get("model")
	if model == "" {
		http.Error(w, "Missing 'model' parameter", http.StatusBadRequest)
		return
	}

	// Проверка на отключенную модель
	if model == "hdl" {
		http.Error(w, "Model 'hdl' is temporarily disabled", http.StatusServiceUnavailable)
		return
	}

	endpoint := config.Cfg.PredictAPIURL + "/" + model

	if _, ok := supportedModels[model]; !ok {
		http.Error(w, "Unsupported model: "+model, http.StatusBadRequest)
		return
	}

	// Соберем JSON-данные из query-параметров
	params := make(map[string]string)
	for key, values := range r.URL.Query() {
		if key == "model" {
			continue
		}
		params[key] = values[0]
	}

	builder, ok := modelBuilders[model]
	if !ok {
		http.Error(w, "Unsupported model: "+model, http.StatusBadRequest)
		return
	}

	// Отдельная проверка для hdl — временно отключена
	if model == "hdl" {
		http.Error(w, "Model 'hdl' is temporarily disabled", http.StatusServiceUnavailable)
		return
	}

	// Генерация payload-а
	payload := builder.Build(params)

	for key, values := range r.URL.Query() {
		if key == "model" {
			continue
		}
		payload[key] = values[0] // берём только первое значение
	}

	// Преобразуем в JSON
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	// Подготовим POST-запрос
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", config.Cfg.AuthToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Отправляем
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from external API", http.StatusInternalServerError)
		logger.Log.Printf("Error reading response: %v", err)
		return
	}

	if resp.StatusCode >= 400 {
		// Логируем ошибку
		logger.Log.Printf("External API returned error for model '%s': %s", model, string(body))

		// Возвращаем ошибку клиенту
		http.Error(w, "External API error: "+string(body), http.StatusBadGateway)
		return
	}

	// Возвращаем клиенту ответ от стороннего API
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)

	// Логируем
	logger.Log.Printf("Model: %s | Sent to: %s | Status: %d\n", model, endpoint, resp.StatusCode)

}
