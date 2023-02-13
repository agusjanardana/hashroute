## Install routehash library

Install di local machine teman-teman dan perhatikan juga tags yang released terbaru

```
go install -v github.com/agusjanardana/routehash@v1.1.4
```

Cara menggunakan

```
routehash --name="name_of_route" --method="api_method" --apiGroup="api_group" --controller="your_controller"

```

Contoh 

```
routehash --name="test/asd/asd" --method="GET" --apiGroup="apiV1" --controller="c1.OauthService.CreateToken"
```

Notice 

```
Perhatikan bahwa, library ini hanya berfungsi ketika struktur folder yang digunakan sudah sesuai dengan templates golang
yang mana harus terdapat router.go pada app/routes/router.go

```