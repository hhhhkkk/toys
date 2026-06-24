package entity

type AwardTarget int

const (
	AwardTargetInvitee AwardTarget = iota // 被邀请人（固定）
	AwardTargetInviter                    // 邀请人（阶梯）
)

type AwardLevel int

const (
	AwardLevelBase AwardLevel = iota // 无
	AwardLevel3                      // 拉满3人
	AwardLevel6                      // 拉满6人
	AwardLevel10                     // 拉满10人
)

func (a AwardTarget) GetAwardLevel(inviteCount int) AwardLevel {
	if a == AwardTargetInvitee {
		return AwardLevelBase
	}
	ret := AwardLevelBase
	switch true {
	case inviteCount >= 10:
		ret = AwardLevel10
		fallthrough
	case inviteCount >= 6:
		ret = AwardLevel6
		fallthrough
	case inviteCount >= 3:
		ret = AwardLevel3
	}
	return ret
}

func (a AwardLevel) String() string {
	switch a {
	case AwardLevelBase:
		return "无"
	case AwardLevel3:
		return "拉满3人"
	case AwardLevel6:
		return "拉满6人"
	case AwardLevel10:
		return "拉满10人"
	default:
		return "未知"
	}
}
