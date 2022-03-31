>> go run .

--> runs the server along with all the helper functions



getData.go --> contains function to pull data from external API.
models.go --> defines go structs that act as containers for data from external API to be filled in.
controllers.go --> defines functions that handles logic and formatting of data to be displayed by endpoints.
main.go --> file which contains main(){} function. Contains gorilla mux routing as well as http.ListenAndServe() which runs the server.