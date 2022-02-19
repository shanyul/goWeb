package service

import (
	"designer-api/internal/models"
)

type ConfigService struct {
	ConfigModel models.ConfigModel
}

type Config struct {
	models.Config
}

func (service *ConfigService) ExistByName(key string) (bool, error) {
	return service.ConfigModel.IsExist(key)
}

func (service *ConfigService) Add(config *Config) error {

	configData := models.Config{}
	configData.Key = config.Key
	configData.Value = config.Value

	if err := service.ConfigModel.AddConfig(&configData); err != nil {
		return err
	}

	return nil
}

func (service *ConfigService) Edit(config *Config) error {

	configData := models.Config{}
	configData.Key = config.Key
	configData.Value = config.Value

	if err := service.ConfigModel.EditConfig(config.Key, &configData); err != nil {
		return err
	}

	return nil
}

func (service *ConfigService) GetAll(config *Config) ([]models.Config, error) {
	data, err := service.ConfigModel.GetConfigList(service.getMaps(config))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (service *ConfigService) Get(key string) (models.Config, error) {

	config, err := service.ConfigModel.GetConfig(key)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (service *ConfigService) Delete(id int) error {
	return service.ConfigModel.DeleteConfig(id)
}

func (service *ConfigService) Count(config *Config) (int64, error) {
	return service.ConfigModel.GetConfigTotal(service.getMaps(config))
}

func (service *ConfigService) getMaps(config *Config) map[string]interface{} {
	maps := make(map[string]interface{})
	if config.Key != "" {
		maps["key"] = config.Key
	}
	if config.Value != "" {
		maps["value"] = config.Value
	}

	return maps
}
