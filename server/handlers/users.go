package handlers

import (
	"encoding/json"
	"fmt"
	authdto "housy/dto/auth"
	dto "housy/dto/result"
	usersdto "housy/dto/users"
	"housy/models"
	"housy/pkg/bcrypt"
	"housy/repositories"
	"net/http"
	"strconv"
	"os"

	// "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range users {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		users[i].Image = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Image = os.Getenv("PATH_FILE") + user.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(usersdto.RequestUser)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	user := models.User{
// 		ID: ,
// 		Fullname:   request.Fullname,
// 		Username:   request.Username,
// 		Email:      request.Email,
// 		Password:   request.Password,
// 		ListAsRole: request.ListAsRole,
// 		Gender:     request.Gender,
// 		Phone:      request.Phone,
// 		Address:    request.Address,
// 		// Image:      request.Image,
// 	}

// 	data, err := h.UserRepository.CreateUser(user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	id, _ := strconv.Atoi((mux.Vars(r)["id"]))
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// password, err := bcrypt.HashingPassword(request.Password)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// }

	// if request.Fullname != "" {
	// 	user.Fullname = request.Fullname
	// }

	// if request.Username != "" {
	// 	user.Username = request.Username
	// }

	// if request.Email != "" {
	// 	user.Email = request.Email
	// }

	// if password != "" {
	// 	user.Password = password
	// }

	// if request.ListAsRole != "" {
	// 	user.ListAsRole = request.ListAsRole
	// }

	// if request.Gender != "" {
	// 	user.Gender = request.Gender
	// }

	// if request.Phone != "" {
	// 	user.Phone = request.Phone
	// }

	// if request.Address != "" {
	// 	user.Address = request.Address
	// }

	if filename != "false" {
		user.Image = filename
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:         u.ID,
		Fullname:   u.Fullname,
		Username:   u.Username,
		Email:      u.Email,
		Password:   u.Password,
		ListAsRole: u.ListAsRole,
		Gender:     u.Gender,
		Phone:      u.Phone,
		Address:    u.Address,
	}
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applocation/json")

	request := new(authdto.ChangePasswordRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Print("CHANGE_PASSWORD")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		fmt.Print("satu")
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.OldPassword, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "YOUR OLD PASSWORD DOESN'T MATCH!!!"}
		json.NewEncoder(w).Encode(response)
		fmt.Print("dua")
		return
	}

	newPassword, err := bcrypt.HashingPassword(request.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user.Password = newPassword

	data, err := h.UserRepository.ChangePassword(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Print("__CHANGE_PASSWORD_SUCCESS", data)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
