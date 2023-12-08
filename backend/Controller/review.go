package controller


import (

"net/http"


"github.com/gin-gonic/gin"

"github.com/n6teen/sa-66-example/Entity"

)

func ListReview(c *gin.Context) {
    var reviews []entity.Review

    if err := entity.DB().Preload("Rating").Preload("Genre").Preload("User").Preload("Movie").Raw("SELECT * FROM reviews").Find(&reviews).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": reviews})
}

func CreateReview(c *gin.Context) {
    var review entity.Review

    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := entity.DB().Create(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": review})
}

func GetReviewByMovieID(c *gin.Context) {
    var review entity.Review

    MovieID := c.Param("MovieID")

    if err := entity.DB().Preload("Rating").Preload("Genre").Preload("User").Preload("Movie").Raw("SELECT * FROM reviews WHERE  = ?", MovieID).Find(&review).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": review})
}

func UpdateReview(c *gin.Context) {

    var review entity.Review
    if err := c.ShouldBindJSON(&review); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
    
    }
    
    
    if tx := entity.DB().Where("id = ?", review.ID).First(&review); tx.RowsAffected == 0 { 
    c.JSON(http.StatusBadRequest, gin.H{"error": "review not found"})
    return
    
    }
    
    
    if err := entity.DB().Save(&review).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
    }
    c.JSON(http.StatusOK, gin.H{"data": review})
    
    }

    func DeleteReview(c *gin.Context) {

        id := c.Param("id")
        if tx := entity.DB().Exec("DELETE FROM reviews WHERE id = ?", id); tx.RowsAffected == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
        return
        }
        c.JSON(http.StatusOK, gin.H{"data": id})
        
    }

    func ListRating(c *gin.Context) {

        var ratings []entity.Rating
        
        if err := entity.DB().Raw("SELECT * FROM ratings").Scan(&ratings).Error; err != nil {
        
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }
        
        c.JSON(http.StatusOK, gin.H{"data": ratings})
        
    }

    func ListGenre(c *gin.Context) {

        var genres []entity.Genre
        
        if err := entity.DB().Raw("SELECT * FROM genres").Scan(&genres).Error; err != nil {
        
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }
        
        c.JSON(http.StatusOK, gin.H{"data": genres})
        
    }

    func ListMovie(c *gin.Context) {
        var movies []entity.Movie
    
        if err := entity.DB().Raw("SELECT * FROM movies").Scan(&movies).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
    
        c.JSON(http.StatusOK, gin.H{"data": movies})
    }

   