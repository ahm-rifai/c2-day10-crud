package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	connection.DatabaseConnect()
	
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")

	route.HandleFunc("/project-detail/{id}", projectDetail).Methods("GET")
	
	route.HandleFunc("/add-project", addProject).Methods("GET")	
	route.HandleFunc("/create-project", createProject).Methods("POST") // CREATE PROJECT

	route.HandleFunc("/edit-project/{id}", editProject).Methods("GET") 
	route.HandleFunc("/update-project/{id}", updateProject).Methods("POST") // UPDATE PROJECT
	
	route.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET") // DELETE PROJECT

	fmt.Println("Server running on port 3000")
	http.ListenAndServe("localhost:3000", route)

}



type Project struct {
	ID 				int
	ProjectName 	string
	StartDate 		time.Time
	EndDate 		time.Time
	FormatStartDate	string
	FormatEndDate	string
	Duration 		string
	Description 	string
	Technologies 	[]string
	Image 			string
}



func home(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	var dataProject []Project
	item := Project{}

	data, _ := connection.Conn.Query(context.Background(), `SELECT "ID", "ProjectName", "StartDate", "EndDate", "Description", "Technologies", "Image" FROM tb_projects`)
	
	for data.Next() {

		err := data.Scan(&item.ID, &item.ProjectName, &item.StartDate, &item.EndDate, &item.Description, &item.Technologies, &item.Image)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		item := Project{
			ID: 			item.ID,
			ProjectName: 	item.ProjectName,
			Duration: 		getDuration(item.StartDate, item.EndDate),
			Description: 	item.Description,
			Technologies: 	item.Technologies,
			Image: 			item.Image,
		}

		dataProject = append(dataProject, item)
	}

	response := map[string]interface{}{
		"DataProject": dataProject,
	}
	tmpl.Execute(w, response)
}

func contact(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func projectDetail(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	ID, _ := strconv.Atoi(mux.Vars(r)["id"])
	
	renderProjectDetail := Project{}

	err = connection.Conn.QueryRow(context.Background(), `SELECT "ID", "ProjectName", "StartDate", "EndDate", "Description", "Technologies", "Image" FROM public.tb_projects WHERE "ID" = $1`, ID).Scan(&renderProjectDetail.ID, &renderProjectDetail.ProjectName, &renderProjectDetail.StartDate, &renderProjectDetail.EndDate, &renderProjectDetail.Description, &renderProjectDetail.Technologies, &renderProjectDetail.Image)
	
	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
	}

	renderProjectDetail = Project{
		ID: 				renderProjectDetail.ID,
		ProjectName: 		renderProjectDetail.ProjectName,
		FormatStartDate: 	formatDate(renderProjectDetail.StartDate),
		FormatEndDate: 		formatDate(renderProjectDetail.EndDate),
		Duration: 			getDuration(renderProjectDetail.StartDate, renderProjectDetail.EndDate),
		Description: 		renderProjectDetail.Description,
		Technologies: 		renderProjectDetail.Technologies,
		Image: 				renderProjectDetail.Image,
	}
	

	response := map[string]interface{}{
		"RenderProjectDetail": renderProjectDetail,
	}
	tmpl.Execute(w, response)
}

func addProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func createProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	ProjectName := r.PostForm.Get("input-project")
	StartDate  := r.PostForm.Get("input-start")
	EndDate  := r.PostForm.Get("input-end")
	Description := r.PostForm.Get("input-desc")
	Technologies := []string{r.PostForm.Get("node"), r.PostForm.Get("react"), r.PostForm.Get("next"), r.PostForm.Get("type")}
	Image := r.PostForm.Get("input-img")

	_, err = connection.Conn.Exec(context.Background(), `INSERT INTO public.tb_projects("ProjectName", "StartDate", "EndDate", "Description", "Technologies", "Image") VALUES ( $1, $2, $3, $4, $5, $6)`, ProjectName, StartDate, EndDate, Description, Technologies, Image)

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editProject(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "Text/html; charset=utp-8")
	var tmpl, err = template.ParseFiles("views/edit-project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	ID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var changeProject = Project{}

	err = connection.Conn.QueryRow(context.Background(), `SELECT "ID", "ProjectName", "StartDate", "EndDate", "Description", "Technologies", "Image" FROM public.tb_projects WHERE "ID" = $1`, ID).Scan(&changeProject.ID, &changeProject.ProjectName, &changeProject.StartDate, &changeProject.EndDate, &changeProject.Description, &changeProject.Technologies, &changeProject.Image)
	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}

	changeProject = Project{
		ID: 				changeProject.ID,
		ProjectName: 		changeProject.ProjectName,
		FormatStartDate: 	returnDate(changeProject.StartDate),
		FormatEndDate: 		returnDate(changeProject.EndDate),
		Description: 		changeProject.Description,
		Technologies: 		changeProject.Technologies,
		Image: 				changeProject.Image,
	}

	response := map[string]interface{} {
		"ChangeProject" : changeProject,
	}

	tmpl.Execute(w, response)
}

func updateProject(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	ID, _ := strconv.Atoi(mux.Vars(r)["id"])
	ProjectName := r.PostForm.Get("input-project")
	StartDate  := r.PostForm.Get("input-start")
	EndDate  := r.PostForm.Get("input-end")
	Description := r.PostForm.Get("input-desc")
	Technologies := []string{r.PostForm.Get("node"), r.PostForm.Get("react"), r.PostForm.Get("next"), r.PostForm.Get("type")}
	Image := r.PostForm.Get("input-img")

	_, err = connection.Conn.Exec(context.Background(), `UPDATE public.tb_projects SET "ProjectName"=$1, "StartDate"=$2, "EndDate"=$3, "Description"=$4, "Technologies"=$5, "Image"=$6 WHERE "ID"=$7`, ProjectName, StartDate, EndDate, Description,  Technologies, Image, ID)
	
	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request)  {
	ID, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := connection.Conn.Exec(context.Background(), `DELETE FROM public.tb_projects WHERE "ID" = $1`, ID)

	if err != nil {
		w.Write([]byte("message : " + err.Error()))
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}










func getDuration(startDate time.Time, endDate time.Time) string {

	distance := endDate.Sub(startDate).Hours() / 24
	var duration string

	if distance >= 30 {
		if (distance / 30) == 1 {
			duration = "1 Month"
		} else {
			duration = strconv.Itoa(int(distance/30)) + " Months"
		}
	} else {
		if distance <= 1 {
			duration = "1 Day"
		} else {
			duration = strconv.Itoa(int(distance)) + " Days"
		}
	}

	return duration
}

func formatDate(InputDate time.Time) string {

	formated := InputDate.Format("02 January 2006")

	return formated
}

func returnDate(InputDate time.Time) string {

	formated := InputDate.Format("2006-01-02")

	return formated
}
