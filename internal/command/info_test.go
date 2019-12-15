package command

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-pg/pg"

	"github.com/sepuka/chat/internal/domain"
	"github.com/stretchr/testify/assert"

	"github.com/sepuka/chat/internal/context"
	"go.uber.org/zap"

	"github.com/sepuka/chat/internal/repository/mocks"
)

var (
	clientRepo = &mocks.ClientRepository{}
	hostsRepo  = &mocks.VirtualHostRepository{}
	logger     = zap.NewNop()
)

const (
	firstUsersContainerId = `8e80e6003b7a`
	unknownContainerId    = `FakeFakeFake`
	firstExistsUser       = `firstUser`
	secondExistsUser      = `secondUser`
	unknownUser           = `stranger`
)

func TestExecInfo(t *testing.T) {
	var (
		firstClient = &domain.Client{Login: firstExistsUser}
		host        = &domain.VirtualHost{
			Id:        777,
			Container: firstUsersContainerId,
			CreatedAt: time.Date(2019, 12, 15, 21, 56, 01, 0, time.UTC),
			Client:    firstClient,
		}
	)

	clientRepo.On(`GetByLogin`, firstExistsUser).Return(firstClient, nil)
	clientRepo.On(`GetByLogin`, secondExistsUser).Return(&domain.Client{}, nil)
	clientRepo.On(`GetByLogin`, unknownUser).Return(nil, pg.ErrNoRows)
	hostsRepo.On(`GetByContainerId`, unknownContainerId).Return(nil, pg.ErrNoRows)
	hostsRepo.On(`GetByContainerId`, firstUsersContainerId).Return(host, nil)

	var testCases = map[string]struct {
		request *context.Request
		result  *Result
		err     error
	}{
		`without any container id`: {
			context.NewRequest(firstExistsUser, domain.Manual, `info`),
			nil,
			NoContainerIdError,
		},
		`wrong container hash format`: {
			context.NewRequest(firstExistsUser, domain.Manual, `info`, `fake container hash`),
			nil,
			WrongContainerIdFormat,
		},
		`unknown user`: {
			context.NewRequest(unknownUser, domain.Manual, `info`, firstUsersContainerId),
			&Result{
				Response: []byte(`you have not any hosts`),
			},
			nil,
		},
		`unknown container`: {
			context.NewRequest(firstExistsUser, domain.Manual, `info`, unknownContainerId),
			nil,
			WrongContainerIdFormat,
		},
		`container does not belong at user`: {
			context.NewRequest(secondExistsUser, domain.Manual, `info`, firstUsersContainerId),
			nil,
			NoHostsByContainerId,
		},
		`information`: {
			context.NewRequest(firstExistsUser, domain.Manual, `info`, firstUsersContainerId),
			&Result{
				Response: []byte("#777\t8e80e6003b7a\tcreated at 15 Dec 19 21:56 UTC"),
			},
			nil,
		},
	}
	var (
		info = NewInfo(clientRepo, hostsRepo, logger.Sugar())
	)

	for testName, testCase := range testCases {
		result, err := info.Exec(testCase.request)
		assert.Equal(t, testCase.result, result, fmt.Sprintf(`unexpected result at case "%s"`, testName))
		if testCase.err != nil {
			assert.Error(t, testCase.err, err, fmt.Sprintf(`unexpected error at case "%s"`, testName))
		} else {
			assert.NoError(t, testCase.err, err, fmt.Sprintf(`unexpecred error at case "%s"`, testName))
		}
	}
}
