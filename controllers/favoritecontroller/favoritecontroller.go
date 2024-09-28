package favoritecontroller

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

func GetAll(w http.ResponseWriter, r *http.Request) {
	var favorite []models.Favorite

	//Ambil data dari database
	if err := config.DB.Find(&favorite).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	if len(favorite) == 0 {
		helper.ResponseJson(w, http.StatusNotFound, "No Favorite found for this user")
		return
	}

	helper.ResponseJson(w, http.StatusOK, favorite)

}

func GetByID(w http.ResponseWriter, r *http.Request) {

	// Ambil user_id dari context
	userID := r.Context().Value("user_id").(int)

	var favorite []models.Favorite
	if err := config.DB.Where("user_id = ?", userID).Find(&favorite).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "Favorite not found"}
			helper.ResponseJson(w, http.StatusNotFound, response)
			return
		}
		response := map[string]string{"message": "Error finding Favorite"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJson(w, http.StatusOK, favorite)

}

func Create(w http.ResponseWriter, r *http.Request) {
	var favorite models.Favorite

	// Decode request body menjadi struct Favorite
	err := json.NewDecoder(r.Body).Decode(&favorite)
	if err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		return
	}

	// Ambil user_id dari context
	userID := r.Context().Value("user_id").(int)

	favorite.UserID = userID

	// Cek apakah Airdrop ID ada (valid)
	var airdrop models.Airdrop
	if err := config.DB.First(&airdrop, favorite.AirdropID).Error; err != nil {
		helper.ResponseJson(w, http.StatusNotFound, map[string]string{"message": "Airdrop not found"})
		return
	}

	// Simpan data favorite ke database
	if err := config.DB.Create(&favorite).Error; err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, map[string]string{"message": "Failed to add favorite"})
		return
	}

	response := map[string]string{"Message": "Berhasi input Favorite"}
	helper.ResponseJson(w, http.StatusOK, response)
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

	var favorite models.Favorite
	if err := config.DB.First(&favorite, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "Favorite not found"}
			helper.ResponseJson(w, http.StatusNotFound, response)
			return
		}
		response := map[string]string{"message": "Error finding Favorite"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// Hapus dari database
	if err := config.DB.Delete(&favorite).Error; err != nil {
		response := map[string]string{"message": "Error deleting Favorite"}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// Kirim respons berhasil
	response := map[string]string{"message": "Favorite deleted successfully"}
	helper.ResponseJson(w, http.StatusOK, response)

}
