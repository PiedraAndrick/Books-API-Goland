package service

import (
	"api-go/internal/model"
	"api-go/internal/store"
	"errors"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) ObtenerTodosLosLibros() ([]*model.Libro, error) {
	libros, err := s.store.GetAll()
	if err != nil {
		return nil, err
	}
	return libros, nil
}
func (s *Service) ObtenerLibroPorId(id int) (*model.Libro, error) {
	return s.store.GetByID(id)
}
func (s *Service) CrearLibro(libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("necesitamos el titulo")
	}
	return s.store.Create(&libro)
}
func (s *Service) ActualizarLibro(id int, libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("necesitamos el titulo")
	}

	return s.store.Update(id, &libro)
}
func (s *Service) EliminarLibro(id int) error {
	return s.store.Delete(id)
}
