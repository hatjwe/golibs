package systemd

import (
	"errors"
	"fmt"
	log "github.com/hatjwe/golibs/logs"
	"github.com/kardianos/service"
	"go.uber.org/zap"
)

type SystemService struct{}

func (ss *SystemService) Initialization(Name string, DisplayName string, Description string) (service.Service, error) {
	svcConfig := &service.Config{
		Name:        Name,
		DisplayName: DisplayName,
		Description: Description,
	}
	systemService := &SystemService{}
	s, err := service.New(systemService, svcConfig)
	if err != nil {
		fmt.Printf("service New failed, err: %v\n", err)
		return nil, errors.New("初始化service失败," + err.Error())
	}

	return s, nil

}
func (ss *SystemService) Install(s service.Service) error {
	err := s.Install()
	if err != nil {
		log.Logger.Error("安装系统服务错误", zap.Error(err))
	}
	return err
}
func (ss *SystemService) Start(s service.Service) error {
	var startSystemctl bool
	go func() {
		err := s.Start()
		if err != nil {
			startSystemctl = true
			log.Logger.Error("启动错误", zap.Error(err))
		}
		log.Logger.Info("启动成功")
	}()
	if startSystemctl {
		return errors.New("启动错误")
	}
	return nil
}

func (ss *SystemService) run() error {
	err := ss.run()
	if err != nil {
		log.Logger.Error("运行错误", zap.Error(err))
		return errors.New("运行错误" + err.Error())
	}
	log.Logger.Info("运行服务成功")
	return nil
}

func (ss *SystemService) Stop(s service.Service) error {
	err := ss.Stop(s)
	if err != nil {
		log.Logger.Error("停止错误", zap.Error(err))
		return errors.New("停止错误" + err.Error())
	}
	log.Logger.Info("停止服务成功")
	return nil
}
