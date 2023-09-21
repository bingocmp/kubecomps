package models

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"yunion.io/x/jsonutils"
	"yunion.io/x/kubecomps/pkg/kubeserver/api"
	"yunion.io/x/kubecomps/pkg/kubeserver/k8s/common/model"
)

var (
	generalServiceManager *SGeneralServiceManager
)

func init() {
	GetGeneralServiceManager()
}

type SGeneralServiceManager struct {
	model.SK8sNamespaceResourceBaseManager
}

type SGeneralService struct {
	model.SK8sNamespaceResourceBase
	UnstructuredResourceBase
}

func GetGeneralServiceManager() *SGeneralServiceManager {
	if generalServiceManager == nil {
		generalServiceManager = &SGeneralServiceManager{
			SK8sNamespaceResourceBaseManager: model.NewK8sNamespaceResourceBaseManager(new(SGeneralService), "k8s_generalservice", "k8s_generalservices"),
		}
		generalServiceManager.SetVirtualObject(generalServiceManager)
		RegisterK8sModelManager(generalServiceManager)
	}
	return generalServiceManager
}

func (m *SGeneralServiceManager) GetK8sResourceInfo() model.K8sResourceInfo {
	return model.K8sResourceInfo{
		ResourceName: api.ResourceNameGeneralService,
		KindName:     api.KindNameGeneralService,
		Object:       &unstructured.Unstructured{},
	}
}

func (obj *SGeneralService) GetAPIObject() (*api.GeneralService, error) {
	out := new(api.GeneralService)
	if err := obj.ConvertToAPIObject(obj, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (obj *SGeneralService) FillAPIObjectBySpec(specObj jsonutils.JSONObject, out IUnstructuredOutput) error {
	ret := out.(*api.GeneralService)
	if params, err := specObj.Get("params"); err == nil {
		err = json.Unmarshal([]byte(params.String()), &ret.Params)
		if err != nil {
			return err
		}
	}
	ret.ReleaseId, _ = specObj.GetString("releaseId")
	ret.ServiceId, _ = specObj.GetString("serviceId")
	return nil
}

func (obj *SGeneralService) FillAPIObjectByStatus(statusObj jsonutils.JSONObject, out IUnstructuredOutput) error {
	ret := out.(*api.GeneralService)
	phase, _ := statusObj.GetString("phase")
	ret.GeneralServiceStatus.Phase = phase

	reason, _ := statusObj.GetString("reason")
	ret.GeneralServiceStatus.Reason = reason

	tryTimes, _ := statusObj.Int("tryTimes")
	ret.TryTimes = int32(tryTimes)

	extInfo := &ret.GeneralServiceStatus.ExternalInfo
	if extraInfoObj, err := statusObj.Get("externalInfo"); err == nil {
		extraInfoObj.Unmarshal(extInfo)
	}
	return nil
}
