package arona

import (
	"context"
	"fmt"

	"github.com/arisu-archive/arona-protos/protos"
)

type ArenaService service

type ArenaRankListRequestWrapper struct {
	*protos.ArenaRankListRequest
}

func (w ArenaRankListRequestWrapper) Packet() *protos.RequestPacket {
	return &w.RequestPacket
}

func (s *ArenaService) GetRanks(
	ctx context.Context,
	session *UserSession,
	rank int32,
	count int32,
) (*protos.ArenaRankListResponse, error) {
	param := ArenaRankListRequestWrapper{
		ArenaRankListRequest: &protos.ArenaRankListRequest{
			StartIndex: rank,
			Count:      count,
		},
	}
	req, err := s.client.R().WithSession(session).Game(ctx, protos.Protocol_Arena_RankList, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create arena rank list request: %w", err)
	}
	result := new(protos.ArenaRankListResponse)
	_, err = s.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("failed to get arena rank list response: %w", err)
	}
	return result, nil
}

type ArenaDailyRewardRequestWrapper struct {
	*protos.ArenaDailyRewardRequest
}

func (w ArenaDailyRewardRequestWrapper) Packet() *protos.RequestPacket {
	return &w.RequestPacket
}

func (s *ArenaService) DailyReward(ctx context.Context, session *UserSession) (*protos.ArenaDailyRewardResponse, error) {
	param := ArenaDailyRewardRequestWrapper{
		ArenaDailyRewardRequest: &protos.ArenaDailyRewardRequest{},
	}
	req, err := s.client.R().WithSession(session).Game(ctx, protos.Protocol_Arena_DailyReward, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create arena daily reward request: %w", err)
	}
	result := new(protos.ArenaDailyRewardResponse)
	_, err = s.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("failed to get arena daily reward response: %w", err)
	}
	return result, nil
}

type ArenaCumulativeTimeRewardRequestWrapper struct {
	*protos.ArenaCumulativeTimeRewardRequest
}

func (w ArenaCumulativeTimeRewardRequestWrapper) Packet() *protos.RequestPacket {
	return &w.RequestPacket
}

func (s *ArenaService) CumulativeTimeReward(ctx context.Context, session *UserSession) (*protos.ArenaCumulativeTimeRewardResponse, error) {
	param := ArenaCumulativeTimeRewardRequestWrapper{
		ArenaCumulativeTimeRewardRequest: &protos.ArenaCumulativeTimeRewardRequest{},
	}
	req, err := s.client.R().WithSession(session).Game(ctx, protos.Protocol_Arena_CumulativeTimeReward, param)
	if err != nil {
		return nil, fmt.Errorf("failed to create arena cumulative time reward request: %w", err)
	}
	result := new(protos.ArenaCumulativeTimeRewardResponse)
	_, err = s.client.Do(ctx, req, result)
	if err != nil {
		return nil, fmt.Errorf("failed to get arena cumulative time reward response: %w", err)
	}
	return result, nil
}
