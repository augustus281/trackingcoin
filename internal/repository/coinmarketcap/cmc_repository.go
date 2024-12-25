package repository

type ICMCRepository interface{}

type cmcRepo struct{}

func NewCMCRepo() ICMCRepository {
	return &cmcRepo{}
}
