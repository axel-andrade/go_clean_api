package repositories_impl

import (
	"context"
	"go_clean_api/api/entities"
	"go_clean_api/api/infra/database"
	"time"
)

var ctx = context.Background()

type TokenManagerRepositoryImpl struct{}

/**
** FetchAuth()aceita o AccessDetailsda ExtractTokenMetadatafunção e procura-o no redis. Se o
** registro não for encontrado, pode significar que o token expirou, portanto, um erro é gerado.
**/

func (tmri *TokenManagerRepositoryImpl) GetAuth(auth *entities.AccessDetails) (string, error) {

	rdb := database.GetRedisDB()
	userid, err := rdb.Get(ctx, auth.AccessUUID).Result()
	if err != nil {
		return "", err
	}

	return userid, nil
}

/**
** Passamos no TokenDetailsque contém informações sobre o tempo de expiração dos JWTs e os
** uuids usados ​​na criação dos JWTs. Se o tempo de expiração for atingido para o token de
** atualização ou para o token de acesso , o JWT será excluído automaticamente do Redis.
**/

func (tmri *TokenManagerRepositoryImpl) CreateAuth(userid string, td *entities.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	rdb := database.GetRedisDB()
	errAccess := rdb.Set(ctx, td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := rdb.Set(ctx, td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}
