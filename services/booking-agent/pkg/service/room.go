package service

import (
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/pkg/repository"
)

type RoomListService struct {
	repo repository.RoomList
}

func NewRoomListService(repo repository.RoomList) *RoomListService {
	return &RoomListService{repo: repo}
}

func (s *RoomListService) UpdateRooms(rooms []models.Room) (int, error) {
	return s.repo.UpdateRooms(rooms)
}

func (s *RoomListService) GetAllRooms() ([]models.Room, error) {
	return s.repo.GetAllRooms()
}

func (s *RoomListService) GetRoomById(roomId int) (models.Room, error) {
	return s.repo.GetRoomById(roomId)
}

func (s *RoomListService) GetAllRoomsBySearch(params models.SearchParams) ([]models.Room, error) {
	return s.repo.GetAllRoomsBySearch(params)
}
