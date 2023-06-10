package Service

func GetICPInfo(url string) (map[string]interface{}, error) {
	oldMap, err := GetSourceICPInfo(url)
	if err != nil {
		return nil, err
	}
	//新建一个newMap
	newMap := make(map[string]interface{})
	//获取map中isDomainICPOk的值
	isDomainICPOk := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["isDomainICPOk"]
	//判断是否备案
	if isDomainICPOk == float64(1) {
		icpSerial := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["ICPSerial"]
		//腾讯那的api写的就是Orgnization 所以有拼写错误
		organization := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["Orgnization"]
		detectTime := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["detect_time"]
		url := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["url"]
		newMap["isDomainICPOk"] = "True"
		newMap["ICPSerial"] = icpSerial
		newMap["Organization"] = organization
		newMap["detect_time"] = detectTime
		newMap["url"] = url
		return newMap, nil
	} else {
		detectTime := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["detect_time"]
		url := oldMap["data"].(map[string]interface{})["results"].(map[string]interface{})["url"]
		newMap["isDomainICPOk"] = "False"
		newMap["detect_time"] = detectTime
		newMap["url"] = url
		return newMap, nil

	}
}
