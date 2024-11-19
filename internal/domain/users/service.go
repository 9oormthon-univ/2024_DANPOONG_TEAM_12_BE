package users

import "github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"

type usersService struct {
	usersRepository *UsersRepository
}

func SetUsersService(repository *UsersRepository) types.UsersService {
	u := &usersService{
		usersRepository: repository,
	}
	return u
}
