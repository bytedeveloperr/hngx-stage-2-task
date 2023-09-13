package api

import (
	"strconv"

	model "github.com/bytedeveloperr/hng-stage-2/models"
)

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

func CreatePerson(c *Context) (interface{}, error) {
	name, ok := c.Body["name"].(string)
	var person = model.Person{Name: name}

	if !ok {
		return nil, Error{message: "Name must be of type string"}
	}

	result := c.DB.Create(&person)
	if result.Error != nil {
		return nil, result.Error
	}

	return person.ToResponse(), nil
}

func GetPersons(c *Context) (interface{}, error) {
	var persons []model.Person
	result := c.DB.Find(&persons)
	if result.Error != nil {
		return nil, result.Error
	}

	var res []model.PersonResponse
	for _, person := range persons {
		res = append(res, *person.ToResponse())
	}

	return res, nil
}

func GetPerson(c *Context) (interface{}, error) {
	var rawId = c.Params["id"]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		if rawId == "" {
			return nil, Error{message: "User ID is not valid"}
		}
	}

	var person model.Person
	result := c.DB.Where("id = ? OR name = ?", id, rawId).Find(&person)
	if result.Error != nil {
		return nil, result.Error
	}

	return person.ToResponse(), nil
}

func UpdatePerson(c *Context) (interface{}, error) {
	var rawId = c.Params["id"]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		if rawId == "" {
			return nil, Error{message: "User ID is not valid"}
		}
	}

	name, ok := c.Body["name"].(string)
	if !ok {
		return nil, Error{message: "Name must be of type string"}
	}

	var person model.Person
	result := c.DB.Where("id = ? OR name = ?", id, rawId).Find(&person)
	if result.Error != nil {
		return nil, result.Error
	}

	person.Name = name
	result = c.DB.Save(&person)
	if result.Error != nil {
		return nil, result.Error
	}

	return person.ToResponse(), nil
}

func DeletePerson(c *Context) (interface{}, error) {
	var rawId = c.Params["id"]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		if rawId == "" {
			return nil, Error{message: "User ID is not valids"}
		}
	}

	var person model.Person
	result := c.DB.Where("id = ? OR name = ?", id, rawId).Find(&person)
	if result.Error != nil {
		return nil, result.Error
	}

	result = c.DB.Delete(&person)
	if result.Error != nil {
		return nil, result.Error
	}

	return nil, nil
}
