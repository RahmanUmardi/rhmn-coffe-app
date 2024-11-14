package usecase

import (
	"rhmn-coffe/entity"
	"rhmn-coffe/entity/dto"
	"rhmn-coffe/shared/service"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Register(payload dto.AuthRequestDto) (entity.User, error)
}

type authUseCase struct {
	useCase    UserUsecase
	jwtService service.JwtService
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {

	user, err := a.useCase.FindByUsernamePassword(payload.Username, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtService.CreateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	message := "Success Login"

	response := dto.AuthResponseDto{
		Message: message,
		Token:   token.Token,
	}

	return response, nil
}

func (a *authUseCase) Register(payload dto.AuthRequestDto) (entity.User, error) {
	return a.useCase.Register(entity.User{Username: payload.Username, Password: payload.Password})
}

func NewAuthUseCase(uc UserUsecase, jwtService service.JwtService) AuthUseCase {
	return &authUseCase{useCase: uc, jwtService: jwtService}
}
