package entity

import (
	"fmt"
	"time"
	"strconv"
	"errors"
	"regexp"
)

type grade int
const (
	B1 grade = iota
	B2
	B3
	B4
	M1
	M2
	Error
)

var (
	ErrInvalidId = errors.New("id is required or invalid")
	ErrInvalidName = errors.New("name is required or invalid")
	ErrInvalidYear = errors.New("invalid year")
)

type participantEntity struct {
	id string
	name string
}

type ParticipantEntity interface {
	GetID() string
	GetName() string
	DistinguishGrade() (grade, error)
}

func NewParticipantEntity(id string, name string) (ParticipantEntity ,error) {

	match, err := regexp.MatchString("^[0-9]{2}(T|G)[0-9]{3}$", id);
	if err != nil {
		return nil, fmt.Errorf("calling regrep.MatchString: %v", err)
	}
	if !match {
		return nil, ErrInvalidId
	}

	match, err = regexp.MatchString(".+", name);
	if err != nil {
		return nil, fmt.Errorf("calling regrep.MatchString: %v", err)
	}
	if !match {
		return nil, ErrInvalidName
	}

	return &participantEntity{id: id, name: name}, nil
}

func (participantEntity *participantEntity) GetID() string {
	return participantEntity.id
}

func (participantEntity *participantEntity) GetName() string {
	return participantEntity.name
}

func (participantEntity *participantEntity) DistinguishGrade() (grade, error) {
	t := time.Now()
	nowYear := t.Year()
	id := participantEntity.GetID();
	admissionYear, err := strconv.Atoi(fmt.Sprintf("%d%s", 20, id[0:2]))
	if err != nil {
		return Error, ErrInvalidYear
	}

	switch nowYear-admissionYear {
	case 0:
		if id[2] == uint8('G') {
			return M1, nil
		}
		return B1, nil
	case 1:
		if id[2] == uint8('G') {
			return M2, nil
		}
		return B2, nil
	case 2:
		return B3, nil
	case 3:
		return B4, nil
	default:
		return Error, ErrInvalidYear
	}
}