~~//postgre DB dockerla kaldıralım~~
//config ler lazım db config

// db katmanı --> db open yapıcak connectionu sağlıcak ve repoya paslıcak.
// repository todo --> getTodosFromDB,addTodoToDB
// service todo --> getTodos , addTodo --> repoyla konuşup transactionları başlatıcak
// handlers --> getTodos , addTodo endpointlere gelen isteği ilgili service fonks.
// maplice
// server --> handlers[] --> server config --> handlerları bu app e register edicez
// serveri ayağı kaldırıacz.

//gorm model --> todo --> id, task

// db_test


im wrting something