package bookcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aldeny/go-crud-library/helper"
	"github.com/aldeny/go-crud-library/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var responseSusccess = helper.ResponseSuccess
var responseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := models.DB.Find(&books).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseSusccess(w, http.StatusOK, books)

}

func Show(w http.ResponseWriter, r *http.Request) {
	getr := mux.Vars(r)
	id, err := strconv.ParseInt(getr["id"], 10, 64)

	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			responseError(w, http.StatusNotFound, "Buku tidak ditemukan")
			return
		default:
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseSusccess(w, http.StatusOK, book)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&book).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseSusccess(w, http.StatusCreated, book)
}

func Update(w http.ResponseWriter, r *http.Request) {
	getr := mux.Vars(r)
	id, err := strconv.ParseInt(getr["id"], 10, 64)

	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book models.Book

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Where("id = ?", id).Updates(&book).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	book.Id = id

	responseSusccess(w, http.StatusOK, book)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	getr := mux.Vars(r)
	id, err := strconv.ParseInt(getr["id"], 10, 64)

	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var book models.Book
	if err := models.DB.Where("id = ?", id).Delete(&book).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseSusccess(w, http.StatusOK, "Buku berhasil di hapus!")
}
