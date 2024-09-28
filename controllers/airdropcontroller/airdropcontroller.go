package airdropcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabillahsakti/coindropedia/config"
	"github.com/sabillahsakti/coindropedia/helper"
	"github.com/sabillahsakti/coindropedia/models"
	"gorm.io/gorm"
)

func Show(w http.ResponseWriter, r *http.Request) {
	var airdropInput []models.Airdrop

	//Ambil data dari database
	if err := config.DB.Find(&airdropInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	if len(airdropInput) == 0 {
		helper.ResponseJson(w, http.StatusNotFound, "No chats found for this user")
		return
	}

	helper.ResponseJson(w, http.StatusOK, airdropInput)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	//Mengambil ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := map[string]string{"message": "Invalid ID"}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	var airdrop models.Airdrop
	if err := config.DB.First(&airdrop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "Airdrop not found"}
			helper.ResponseJson(w, http.StatusNotFound, response)
			return
		}
		response := map[string]string{"message": "Error finding airdrop"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// Hapus dari database
	if err := config.DB.Delete(&airdrop).Error; err != nil {
		response := map[string]string{"message": "Error deleting airdrop"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// Kirim respons berhasil
	response := map[string]string{"message": "Airdrop deleted successfully"}
	helper.ResponseJson(w, http.StatusOK, response)

}

func Create(w http.ResponseWriter, r *http.Request) {
	//Mengambil inputan json dari front end

	var airdropInput models.Airdrop
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&airdropInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	// Insert ke database
	if err := config.DB.Create(&airdropInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"Message": "Berhasi input airdropl"}
	helper.ResponseJson(w, http.StatusOK, response)
}
