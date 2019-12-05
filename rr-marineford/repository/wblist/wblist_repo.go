package wblist

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"rr-factory.gloryholiday.com/yuetu/marineford/logger"
	"rr-factory.gloryholiday.com/yuetu/marineford/proto/marineford"
	pb "rr-factory.gloryholiday.com/yuetu/marineford/proto/marineford"
	"rr-factory.gloryholiday.com/yuetu/marineford/repository"
)

type ApplyType int32


var wbruleCollection = "wbrule"

var wbIndexKeys = []string{"displaynumber","applications.platformgroup.groupid","applications.providergroup.groupid","creator","enabled"}

type WBRuleMessage struct {
	WBRule *pb.WBRule
}

type WBRuleApplication struct {
	PlatformGroup 	*repository.OTAPlatformGroup
	ProviderGroup 	*repository.ProviderGroup
	ApplyType	   	string
	Enabled 		bool
}

type WBRuleEntity struct {
	repository.BaseInfo		`bson:",inline"`
	TripType				string
	Routes					[]*repository.LocationPair
	AbsSellDateRange		*repository.AbsDepDateRangeConstrain
	RelDepDateRange			*repository.RelDateRangeConstraint
	AbsDepDateRange 		[]*repository.AbsDepDateRangeConstrain
	Applications 			[]*WBRuleApplication
	Remark 					string
	ChangeHistory 			*repository.ChangeHistoryEntity
}


func (e *WBRuleEntity) MarshalMessage() (proto.Message,error) {
	return &pb.WBRule{
		BaseInfo:			repository.MarshalBaseInfoMessage(&e.BaseInfo),
		TripType:			pb.TripType(pb.TripType_value[e.TripType]),
		Routes:				repository.MarshalRouteMessage(e.Routes),
		AbsSellDateConstraint:repository.MarshalAbsDateRangeConstraintMessage(e.AbsSellDateRange)
		RelDepDateConstraint:  repository.MarshalRelDateRangeConstraintMessage(e.RelDepDateRange),
		AbsDepDateConstraints: repository.MarshalAbsDateRangeConstraintMessages(e.AbsDepDateRanges),
		ApplyTo:               marshalWBRuleApplicationMessages(e.Applications),
		Remark:                e.Remark,	
	},nil
}

func (wm *WBRuleMessage) MarshalEntity() (repository.Entity,error) {
	p := wm.WBRule
	return &WBRuleEntity{
		BaseInfo:         repository.MarshalBaseInfoEntity(p.BaseInfo),
		TripType:         p.TripType.String(),
		Routes:           repository.MarshalRouteEntities(p.Routes), // airport is not supported
		AbsSellDateRange: repository.MarshalAbsDateRangeConstraintEntity(p.AbsSellDateConstraint),
		RelDepDateRange:  repository.MarshalRelDateRangeConstraintEntity(p.RelDepDateConstraint),
		AbsDepDateRanges: repository.MarshalAbsDateRangeConstraintEntities(p.AbsDepDateConstraints),
		Applications:     marshalWBRuleApplicationEntities(p.ApplyTo),
		Remark:           p.Remark,
		ChangeHistory:    &repository.ChangeHistoryEntity{OperationLogs: []*repository.OperationLogEntity{}},
	},nil	
}
func (e *WBRuleEntity) GetBaseInfo() *repository.BaseInfo {
	return &e.BaseInfo
}


func (e *WBRuleEntity) GetPolicyPrefix() sting {
	return "WB"
}

func (e *WBRuleEntity) SetBaseInfo(baseinfo *repository.BaseInfo) {
	e.BaseInfo = *BaseInfo
}

func marshalWBRuleApplicationMessages(applications []*WBRuleApplication []*pb.WBRuleApplication) {
		var res []*pb.WBRuleApplication
	for _, app := range {
		res = append(res,marshalWBRuleApplicationMessage(app))
	}
	return res
}

func marshalWBRuleApplicationMessage(){
	pass
}