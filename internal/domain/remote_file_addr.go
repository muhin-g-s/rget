package domain

import (
	"errors"
	"fmt"
	"net/url"
)

type RemoteFileAddr struct {
	value string
}

var (
	ErrRemoteFileEmptyStr            = errors.New("cannot be an empty string")
	ErrRemoteFileAddrInvalidProtocol = errors.New("invalid protocol")
	ErrRemoteFileAddrNotValidUrl     = errors.New("not valid url")
)

func NewRemoteFileAddr(str string) (RemoteFileAddr, error) {
	err := validate(str)
	if err != nil {
		return RemoteFileAddr{}, err
	}

	return RemoteFileAddr{
		value: str,
	}, nil
}

func ParseRemoteFileAddrs(strs []string) ([]RemoteFileAddr, error) {
	result := make([]RemoteFileAddr, 0, len(strs))

	for _, str := range strs {
		validated, err := NewRemoteFileAddr(str)
		if err != nil {
			return nil, fmt.Errorf("%w %s", err, str)
		}

		result = append(result, validated)
	}

	return result, nil
}

func (vo RemoteFileAddr) Value() string {
	return vo.value
}

func validate(str string) error {
	if str == "" {
		return ErrRemoteFileEmptyStr
	}

	u, err := url.ParseRequestURI(str)
	if err != nil {
		return ErrRemoteFileAddrNotValidUrl
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return ErrRemoteFileAddrInvalidProtocol
	}

	if u.Host == "" {
		return ErrRemoteFileAddrNotValidUrl
	}

	return nil
}
