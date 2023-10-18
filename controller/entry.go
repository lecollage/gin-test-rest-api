package controller

import (
	"net/http"
	"test_api/helper"
	"test_api/model"
	"test_api/resources"

	"github.com/gin-gonic/gin"
)

func AddEntry(context *gin.Context) {
	var input model.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID
	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

func FindEntryById(id string) (model.Entry, error) {
	var entry model.Entry
	err := resources.Database.Model(&model.Entry{}).Where("id=?", id).Find(&entry).Error

	return entry, err
}

func GetTestEntryById(context *gin.Context) {
	// get entry data
	id := context.Param("id")
	entry, errEntry := FindEntryById(id)

	if errEntry != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errEntry.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": entry})
}

func GetEntryById(context *gin.Context) {
	// get user data
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get entry data
	id := context.Param("id")
	entry, errEntry := FindEntryById(id)

	if errEntry != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errEntry.Error()})
		return
	}

	if user.ID != entry.UserID {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Entry doesn't belong to this user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": entry})
}

func UpdateEntry(context *gin.Context) {
	// get user data
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get entry data
	id := context.Param("id")
	_, errFindEntry := FindEntryById(id)

	if errFindEntry != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errFindEntry.Error()})
		return
	}

	var input model.Entry

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	errEntry := resources.Database.Model(&model.Entry{}).Where("id=?", id).Omit("user_id").Updates(input).Error

	if errEntry != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errEntry.Error()})
		return
	}

	var updatedEntry model.Entry
	errEntry = resources.Database.Model(&model.Entry{}).Where("id=?", id).Find(&updatedEntry).Error

	if errEntry != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errEntry.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedEntry})
}

func DeleteEntry(context *gin.Context) {
	// get user data
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get entry data
	id := context.Param("id")

	var entry model.Entry

	errEntry := resources.Database.Model(&model.Entry{}).Where("id=?", id).Delete(&entry).Error

	if errEntry != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errEntry.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "deleted"})
}
