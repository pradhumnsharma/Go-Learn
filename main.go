package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Student struct {
	PersonalDetails
	RollNo int    `json:"roll_no"`
	Branch string `json:"branch"`
}

type Teacher struct {
	PersonalDetails
	Id          int    `json:"id"`
	Designation string `json:"designation"`
	Subject     string `json:"subject"`
	Leaves      int    `json:"leaves"`
	Salary      int    `json:"salary"`
}

type PersonalDetails struct {
	FullName string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Pincode  int    `json:"pincode"`
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

var studentsList []Student
var teachersList []Teacher

func main() {
	fmt.Println("CRUD")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	r.HandleFunc("/create-student", createStudent).Methods("POST")
	r.HandleFunc("/students", getAllStudents).Methods("GET")
	r.HandleFunc("/students/{roll_no}", getStudentDetails).Methods("GET")
	r.HandleFunc("/students/{roll_no}", updateStudentDetails).Methods("PUT")
	r.HandleFunc("/students/{roll_no}", deleteStudent).Methods("DELETE")
	r.HandleFunc("/students", deleteAllStudent).Methods("DELETE")

	r.HandleFunc("/create-teacher", createTeacher).Methods("POST")
	r.HandleFunc("/teachers", getAllTeachers).Methods("GET")
	r.HandleFunc("/teachers/{id}", getTeacherDetails).Methods("GET")
	r.HandleFunc("/teachers/{id}", updateTeacherDetails).Methods("PUT")
	r.HandleFunc("/teachers/{id}", deleteTeacher).Methods("DELETE")
	r.HandleFunc("/teachers", deleteAllTeacher).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":6001", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<div><h1>Welcome to Fake Crud</h1><p>Here we will create crud with simple slices and struct without any database.<br>We will create students and teachers and perform operations on them</p></div>"))
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add all details")
		return
	}
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	if len(studentsList) > 0 {
		for _, stud := range studentsList {
			if stud.Email == student.Email {
				response := Response{0, "Student already exist", stud}
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		student.RollNo = studentsList[len(studentsList)-1].RollNo + 1
	} else {
		rand.Seed(time.Now().UnixNano())
		student.RollNo = rand.Intn(4000)
	}
	studentsList = append(studentsList, student)
	response := Response{1, "Student added successfully", student}
	json.NewEncoder(w).Encode(response)
}

func createTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add all details")
		return
	}
	var teacher Teacher
	json.NewDecoder(r.Body).Decode(&teacher)
	if len(teachersList) > 0 {
		for _, teach := range studentsList {
			if teach.Email == teacher.Email {
				response := Response{0, "Teacher already exist", teach}
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		teacher.Id = teachersList[len(teachersList)-1].Id + 1
	} else {
		rand.Seed(time.Now().UnixNano())
		teacher.Id = rand.Intn(4000)
	}
	teachersList = append(teachersList, teacher)
	response := Response{1, "Teacher added successfully", teacher}
	json.NewEncoder(w).Encode(response)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{1, "Students fetched successfully", studentsList}
	json.NewEncoder(w).Encode(response)
}

func getAllTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{1, "Teachers fetched successfully", teachersList}
	json.NewEncoder(w).Encode(response)
}

func getStudentDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, student := range studentsList {
		if strconv.Itoa(student.RollNo) == params["roll_no"] {
			response := Response{1, "Student details found", student}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response := Response{1, "Student details not found", nil}
	json.NewEncoder(w).Encode(response)
}

func getTeacherDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, teacher := range teachersList {
		if strconv.Itoa(teacher.Id) == params["id"] {
			response := Response{1, "Teacher details found", teacher}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response := Response{1, "Teacher details not found", nil}
	json.NewEncoder(w).Encode(response)
}

func updateStudentDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add all details")
		return
	}
	params := mux.Vars(r)
	for index, student := range studentsList {
		if strconv.Itoa(student.RollNo) == params["roll_no"] {
			var student Student
			json.NewDecoder(r.Body).Decode(&student)
			student.RollNo, _ = strconv.Atoi(params["roll_no"])
			studentsList = append(studentsList[:index], studentsList[index+1:]...)
			studentsList = append(studentsList, student)
			response := Response{1, "Student details updated successfully", student}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response := Response{1, "No student to update", nil}
	json.NewEncoder(w).Encode(response)
}

func updateTeacherDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add all details")
		return
	}
	params := mux.Vars(r)
	for index, teacher := range teachersList {
		if strconv.Itoa(teacher.Id) == params["id"] {
			var teacher Teacher
			json.NewDecoder(r.Body).Decode(&teacher)
			teacher.Id, _ = strconv.Atoi(params["id"])
			teachersList = append(teachersList[:index], teachersList[index+1:]...)
			teachersList = append(teachersList, teacher)
			response := Response{1, "Teacher details updated successfully", teacher}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response := Response{1, "No teacher to update", nil}
	json.NewEncoder(w).Encode(response)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, student := range studentsList {
		if strconv.Itoa(student.RollNo) == params["roll_no"] {
			studentsList = append(studentsList[:index], studentsList[index+1:]...)
			response := Response{1, "Student details deleted successfully", student}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response := Response{1, "No student to delete", nil}
	json.NewEncoder(w).Encode(response)
}

func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, teacher := range teachersList {
		if strconv.Itoa(teacher.Id) == params["id"] {
			teachersList = append(teachersList[:index], teachersList[index+1:]...)
			response := Response{1, "Teacher details deleted successfully", teacher}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response := Response{1, "No teacher to deleted", nil}
	json.NewEncoder(w).Encode(response)
}

func deleteAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	studentsList = []Student{}
	response := Response{1, "All students deleted successfully", nil}
	json.NewEncoder(w).Encode(response)
}

func deleteAllTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	teachersList = []Teacher{}
	response := Response{1, "All teachers deleted successfully", nil}
	json.NewEncoder(w).Encode(response)
}
