package main

import "net/http"


func getCat(req *http.Request) (int, any) {
    catId := req.PathValue("catId")

    cat, exists := catsDatabase[catId]
    if !exists {
        Logger.Infof("Cat '%s' not found", catId)
        return http.StatusNotFound, "Cat not found"
    }

    return http.StatusOK, cat
}


func deleteCat(req *http.Request) (int, any) {
    catId := req.PathValue("catId")

    _, exists := catsDatabase[catId]
    if !exists {
        Logger.Infof("Cat '%s' not found", catId)
        return http.StatusNotFound, "Cat not found"
    }

    delete(catsDatabase, catId)

    return http.StatusOK, map[string]string{
        "message": "Cat deleted",
    }
}
